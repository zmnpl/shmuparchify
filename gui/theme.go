package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	pink   = &color.NRGBA{R: 238, G: 111, B: 248, A: 255}
	purple = &color.NRGBA{R: 128, G: 0, B: 128, A: 255}
	orange = &color.NRGBA{R: 198, G: 123, B: 0, A: 255}
	grey   = &color.Gray{Y: 123}
)

// customTheme is a simple demonstration of a bespoke theme loaded by a Fyne app.
type customTheme struct {
}

func (customTheme) Color(c fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNamePrimary, theme.ColorNameHover, theme.ColorNameFocus:
		return pink
	default:
		return theme.DefaultTheme().Color(c, variant)
	}
}

func (customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (customTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (customTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}

func newCustomTheme() fyne.Theme {
	return &customTheme{}
}
