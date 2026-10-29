package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// پالت رنگی azitavpn — دخترونه، شاد، لوکس
var (
	colPinkSoft   = color.NRGBA{R: 0xFF, G: 0xC2, B: 0xD9, A: 0xFF} // صورتی پاستلی
	colPinkDeep   = color.NRGBA{R: 0xFF, G: 0x6F, B: 0xA8, A: 0xFF} // صورتی پررنگ
	colYellow     = color.NRGBA{R: 0xFF, G: 0xE0, B: 0x6B, A: 0xFF} // زرد خورشیدی
	colTurquoise  = color.NRGBA{R: 0x6F, G: 0xE7, B: 0xD8, A: 0xFF} // فیروزه‌ای روشن
	colCream      = color.NRGBA{R: 0xFF, G: 0xF5, B: 0xE6, A: 0xFF} // کرم
	colBg         = color.NRGBA{R: 0xFF, G: 0xF7, B: 0xFA, A: 0xFF} // پس‌زمینه صورتی خیلی روشن
	colFg         = color.NRGBA{R: 0x6B, G: 0x2C, B: 0x4C, A: 0xFF} // متن شرابی تیره
	colGray       = color.NRGBA{R: 0xC8, G: 0xBF, B: 0xC8, A: 0xFF} // خاکستری خاموش
	colGreen      = color.NRGBA{R: 0x4C, G: 0xD9, B: 0x9C, A: 0xFF} // سبز متصل
	colDisabled   = color.NRGBA{R: 0xB0, G: 0xA0, B: 0xB0, A: 0xFF}
)

type aziTheme struct{}

func (aziTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		return colBg
	case theme.ColorNameButton:
		return colPinkDeep
	case theme.ColorNamePrimary:
		return colPinkSoft
	case theme.ColorNameForeground:
		return colFg
	case theme.ColorNameHover:
		return colPinkSoft
	case theme.ColorNameFocus:
		return colYellow
	case theme.ColorNameDisabled:
		return colDisabled
	case theme.ColorNameInputBackground:
		return colCream
	case theme.ColorNameInputBorder:
		return colPinkSoft
	case theme.ColorNamePlaceHolder:
		return colGray
	case theme.ColorNameSelection:
		return colTurquoise
	case theme.ColorNameSeparator:
		return colPinkSoft
	case theme.ColorNameOverlayBackground:
		return colCream
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0x6B, G: 0x2C, B: 0x4C, A: 0x33}
	}
	return theme.DefaultTheme().Color(n, v)
}

func (aziTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (aziTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (aziTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNamePadding:
		return 8
	case theme.SizeNameText:
		return 15
	}
	return theme.DefaultTheme().Size(n)
}
