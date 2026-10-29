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
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type UI struct {
	mu         sync.Mutex
	connected  bool
	btn        *widget.Button
	status     *widget.Label
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
	myWindow.Resize(fyne.NewSize(400, 480))

	ui := &UI{mainWindow: myWindow}

	ui.status = widget.NewLabelWithStyle("خاموش", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	ui.btn = widget.NewButton("اتصال", func() { ui.toggle() })

	title := widget.NewLabelWithStyle("azitavpn", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		widget.NewSeparator(),
		layout.NewSpacer(),
		container.NewCenter(ui.status),
		layout.NewSpacer(),
		ui.btn,
		layout.NewSpacer(),
	)

	myWindow.SetContent(container.NewPadded(content))
	myWindow.ShowAndRun()
}

func (ui *UI) toggle() {
	ui.mu.Lock()
	if ui.connected {
		ui.mu.Unlock()
		go ui.disconnect()
	} else {
		ui.mu.Unlock()
		go ui.connect()
	}
}

func (ui *UI) setStatus(text string) {
	ui.status.SetText(text)
}

func (ui *UI) connect() {
	ui.btn.Disable()
	ui.setStatus("در حال اتصال...")

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
		dialog.ShowError(fmt.Errorf("config.json پیدا نشد"), ui.mainWindow)
		ui.setStatus("خاموش")
		ui.btn.Enable()
		return
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		dialog.ShowError(err, ui.mainWindow)
		ui.setStatus("خاموش")
		ui.btn.Enable()
		return
	}

	cfg := &config.Config{}
	if err := json.Unmarshal(data, cfg); err != nil {
		dialog.ShowError(fmt.Errorf("config.json معتبر نیست: %v", err), ui.mainWindow)
		ui.setStatus("خاموش")
		ui.btn.Enable()
		return
	}
	ui.coreConfig = cfg

	if err := server.Run(true); err != nil {
		dialog.ShowError(err, ui.mainWindow)
		ui.setStatus("خاموش")
		ui.btn.Enable()
		return
	}

	ui.mu.Lock()
	ui.connected = true
	ui.mu.Unlock()
	ui.setStatus("متصل ✓")
	ui.btn.SetText("قطع اتصال")
	ui.btn.Enable()
}

func (ui *UI) disconnect() {
	ui.btn.Disable()
	ui.setStatus("در حال قطع...")

	if err := server.ShutDown(); err != nil {
		dialog.ShowError(err, ui.mainWindow)
	}

	ui.mu.Lock()
	ui.connected = false
	ui.mu.Unlock()
	ui.setStatus("خاموش")
	ui.btn.SetText("اتصال")
	ui.btn.Enable()
}
