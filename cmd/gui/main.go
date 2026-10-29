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
	"fyne.io/fyne/v2/widget"
)

type appState int

const (
	stateOff appState = iota
	stateConnecting
	stateOn
)

type UI struct {
	state        appState
	mu           sync.Mutex
	circle       *canvas.Circle
	statusLabel  *widget.Label
	hintLabel    *widget.Label
	mainWindow   fyne.Window
	coreConfig   *config.Config
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

	// لیبل وضعیت
	ui.statusLabel = widget.NewLabelWithStyle("خاموش", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	ui.statusLabel.TextSize = 22

	// راهنما
	ui.hintLabel = widget.NewLabelWithStyle("برای اتصال، روی خورشید بزن", fyne.TextAlignCenter, fyne.TextStyle{Italic: true})

	// دکمه‌ی نامرئی روی دایره
	tapBtn := widget.NewButton("", func() { ui.toggle() })
	tapBtn.Importance = widget.LowImportance
	tapStack := container.NewStack(circleWrap, container.NewGridWrap(fyne.NewSize(220, 220), tapBtn))

	title := widget.NewLabelWithStyle("azitavpn", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	title.TextSize = 28

	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		widget.NewSeparator(),
		layout.NewSpacer(),
		tapStack,
		layout.NewSpacer(),
		ui.statusLabel,
		ui.hintLabel,
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
		ui.statusLabel.SetText("خاموش")
		ui.hintLabel.SetText("برای اتصال، روی خورشید بزن")
	case stateConnecting:
		ui.circle.FillColor = colYellow
		ui.statusLabel.SetText("در حال اتصال...")
		ui.hintLabel.SetText("لطفاً صبر کن")
	case stateOn:
		ui.circle.FillColor = colGreen
		ui.statusLabel.SetText("متصل ✓")
		ui.hintLabel.SetText("برای قطع، دوباره بزن")
	}
	ui.circle.Refresh()
}

func (ui *UI) connect() {
	// خواندن config.json از کنار فایل اجرایی یا دایرکتوری جاری
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
