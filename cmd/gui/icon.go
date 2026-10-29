package main

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
)

// loadIconFromConfigDir آیکون را از کنار فایل اجرایی یا config.json می‌خواند
func loadIconFromConfigDir() fyne.Resource {
	// ۱. کنار فایل اجرایی
	if exe, err := os.Executable(); err == nil {
		p := filepath.Join(filepath.Dir(exe), "icon.png")
		if data, err := os.ReadFile(p); err == nil {
			return fyne.NewStaticResource("icon.png", data)
		}
	}

	// ۲. کنار config.json (دایرکتوری جاری)
	if cwd, err := os.Getwd(); err == nil {
		p := filepath.Join(cwd, "icon.png")
		if data, err := os.ReadFile(p); err == nil {
			return fyne.NewStaticResource("icon.png", data)
		}
	}

	// ۳. داخل APK (assets) — fallback
	if data, err := os.ReadFile("icon.png"); err == nil {
		return fyne.NewStaticResource("icon.png", data)
	}

	return nil
}
