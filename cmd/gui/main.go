package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/bepass-org/bepass/config"
	"github.com/bepass-org/bepass/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type appState int

const (
	stateOff appState = iota
	stateConnecting
	stateOn
)

type UI struct {
	state      appState
	mu         sync.Mutex
	circle     *canvas.Circle
	statusText *canvas.Text
	hintText   *canvas.Text
	mainWindow fyne.Window
	coreConfig *config.Config
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(aziTheme{})

	if icon := loadIconFromConfigDir(); icon != nil {
		myApp.SetIcon(icon)
	}

	myWindow := myApp.NewWindow("azitavpn")
	myWindow.Resize(fyne.NewSize(420, 560))

	ui := &UI{mainWindow: myWindow}

	// دایره وضعیت
	ui.circle = canvas.NewCircle(colGray)
	ui.circle.Resize(fyne.NewSize(200, 200))
	circleWrap := container.NewGridWrap(fyne.NewSize(220, 220), ui.circle)

	// متن وضعیت
	ui.statusText = canvas.NewText("خاموش", colFg)
	ui.statusText.TextSize = 24
	ui.statusText.TextStyle = fyne.TextStyle{Bold: true}
	ui.statusText.Alignment = fyne.TextAlignCenter

	// راهنما
	ui.hintText = canvas.NewText("برای اتصال، روی دایره بزن", colGray)
	ui.hintText.TextSize = 14
	ui.hintText.Alignment = fyne.TextAlignCenter

	// دکمه‌ی نامرئی روی دایره
	tapBtn := widget.NewButton("", func() { ui.toggle() })
	tapBtn.Importance = widget.LowImportance
	tapStack := container.NewMax(circleWrap, container.NewGridWrap(fyne.NewSize(220, 220), tapBtn))

	title := widget.NewRichText(&widget.TextSegment{
		Style: widget.RichTextStyle{
			SizeName:  theme.SizeNameHeadingText,
			Alignment: fyne.TextAlignCenter,
			TextStyle: fyne.TextStyle{Bold: true},
		},
		Text: "azitavpn",
	})

	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		widget.NewSeparator(),
		layout.NewSpacer(),
		container.NewCenter(tapStack),
		layout.NewSpacer(),
		container.NewCenter(ui.statusText),
		container.NewCenter(ui.hintText),
		layout.NewSpacer(),
	)

	myWindow.SetContent(container.NewPadded(content))
	myWindow.ShowAndRun()
}

func (ui *UI) toggle() {
	ui.mu.Lock()
	defer ui.mu.Unlock()

	switch ui.state {
	case stateOff:
		ui.setState(stateConnecting)
		go ui.connect()
	case stateOn:
		ui.setState(stateOff)
		go ui.disconnect()
	}
}

func (ui *UI) setState(s appState) {
	ui.state = s
	switch s {
	case stateOff:
		ui.circle.FillColor = colGray
		ui.statusText.Text = "خاموش"
		ui.statusText.Color = colFg
		ui.hintText.Text = "برای اتصال، روی دایره بزن"
	case stateConnecting:
		ui.circle.FillColor = colYellow
		ui.statusText.Text = "در حال اتصال..."
		ui.statusText.Color = colFg
		ui.hintText.Text = "لطفاً صبر کن"
	case stateOn:
		ui.circle.FillColor = colGreen
		ui.statusText.Text = "متصل ✓"
		ui.statusText.Color = colGreen
		ui.hintText.Text = "برای قطع، دوباره بزن"
	}
	ui.circle.Refresh()
	ui.statusText.Refresh()
	ui.hintText.Refresh()
}

func (ui *UI) connect() {
	cfgPath := ""
	if exe, err := os.Executable(); err == nil {
		p := filepath.Join(filepath.Dir(exe), "config.json")
		if _, err := os.Stat(p); err == nil {
			cfgPath = p
		}
	}
	if cfgPath == "" {
		if cwd, err := os.Getwd(); err == nil {
			p := filepath.Join(cwd, "config.json")
			if _, err := os.Stat(p); err == nil {
				cfgPath = p
			}
		}
	}
	if cfgPath == "" {
		ui.showError(fmt.Errorf("config.json پیدا نشد"))
		ui.setState(stateOff)
		return
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		ui.showError(err)
		ui.setState(stateOff)
		return
	}

	cfg := &config.Config{}
	if err := json.Unmarshal(data, cfg); err != nil {
		ui.showError(fmt.Errorf("config.json معتبر نیست: %v", err))
		ui.setState(stateOff)
		return
	}
	ui.coreConfig = cfg

	if err := server.Run(true); err != nil {
		ui.showError(err)
		ui.setState(stateOff)
		return
	}

	ui.setState(stateOn)
}

func (ui *UI) disconnect() {
	if err := server.ShutDown(); err != nil {
		ui.showError(err)
	}
	ui.setState(stateOff)
}

func (ui *UI) showError(err error) {
	dialog.ShowError(err, ui.mainWindow)
}
