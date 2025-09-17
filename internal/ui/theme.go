package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Theme struct {
	colors map[string]tcell.Color
}

func NewTheme() Theme {
	theme := Theme{
		colors: map[string]tcell.Color{
			"background": tcell.GetColor("#282a36"),
			"foreground": tcell.GetColor("#f8f8f2"),
			"selection":  tcell.GetColor("#44475a"),
			"comment":    tcell.GetColor("#6272a4"),
			"purple":     tcell.GetColor("#bd93f9"),
			"pink":       tcell.GetColor("#ff79c6"),
			"green":      tcell.GetColor("#50fa7b"),
			"orange":     tcell.GetColor("#ffb86c"),
			"red":        tcell.GetColor("#ff5555"),
			"yellow":     tcell.GetColor("#f1fa8c"),
			"cyan":       tcell.GetColor("#8be9fd"),
		},
	}
	return theme
}

func (t Theme) GetColor(color string) tcell.Color {
	return t.colors[color]
}

func (t Theme) GetTheme() tview.Theme {
	theme := tview.Theme{
		PrimitiveBackgroundColor:    t.GetColor("background"),
		ContrastBackgroundColor:     t.GetColor("selection"),
		MoreContrastBackgroundColor: t.GetColor("background"),
		BorderColor:                 t.GetColor("purple"),
		TitleColor:                  t.GetColor("cyan"),
		GraphicsColor:               t.GetColor("purple"),
		PrimaryTextColor:            t.GetColor("foreground"),
		SecondaryTextColor:          t.GetColor("green"),
		TertiaryTextColor:           t.GetColor("yellow"),
		InverseTextColor:            t.GetColor("background"),
		ContrastSecondaryTextColor:  t.GetColor("pink"),
	}

	return theme
}
