package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// پالت رنگی Suppersun
var (
	colSunYellow = color.NRGBA{R: 0xFF, G: 0xD4, B: 0x47, A: 0xFF} // زرد خورشیدی
	colSunDeep   = color.NRGBA{R: 0xE8, G: 0x8B, B: 0x1A, A: 0xFF} // نارنجی طلایی
	colCream     = color.NRGBA{R: 0xF3, G: 0xE6, B: 0xC4, A: 0xFF} // کرم
	colTurquoise = color.NRGBA{R: 0x3F, G: 0xD0, B: 0xC9, A: 0xFF} // فیروزه‌ای روشن
	colTurqDeep  = color.NRGBA{R: 0x0E, G: 0x3B, B: 0x3A, A: 0xFF} // فیروزه‌ای تیره
	colBg        = color.NRGBA{R: 0x07, G: 0x24, B: 0x23, A: 0xFF} // پس‌زمینه
	colFg        = color.NRGBA{R: 0xFF, G: 0xF6, B: 0xC9, A: 0xFF} // متن
	colHover     = color.NRGBA{R: 0x1A, G: 0x5A, B: 0x57, A: 0xFF}
	colFocus     = color.NRGBA{R: 0xFF, G: 0xD4, B: 0x47, A: 0xFF}
	colDisabled  = color.NRGBA{R: 0x66, G: 0x66, B: 0x66, A: 0xFF}
)

type sunTheme struct{}

func (sunTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		return colBg
	case theme.ColorNameButton:
		return colSunDeep
	case theme.ColorNamePrimary:
		return colSunYellow
	case theme.ColorNameForeground:
		return colFg
	case theme.ColorNameHover:
		return colHover
	case theme.ColorNameFocus:
		return colFocus
	case theme.ColorNameDisabled:
		return colDisabled
	case theme.ColorNameInputBackground:
		return colTurqDeep
	case theme.ColorNameInputBorder:
		return colTurquoise
	case theme.ColorNamePlaceHolder:
		return colCream
	case theme.ColorNameSelection:
		return colTurquoise
	case theme.ColorNameSeparator:
		return colTurqDeep
	case theme.ColorNameOverlayBackground:
		return colTurqDeep
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x99}
	}
	return theme.DefaultTheme().Color(n, v)
}

func (sunTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (sunTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (sunTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNamePadding:
		return 6
	case theme.SizeNameText:
		return 14
	}
	return theme.DefaultTheme().Size(n)
}
