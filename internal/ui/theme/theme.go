package theme

import (
	"github.com/charmbracelet/lipgloss"
)

type Theme struct {
	SelectedBackground lipgloss.AdaptiveColor
	PrimaryBorder      lipgloss.AdaptiveColor
	FaintBorder        lipgloss.AdaptiveColor
	SecondaryBorder    lipgloss.AdaptiveColor
	FaintText          lipgloss.AdaptiveColor
	PrimaryText        lipgloss.AdaptiveColor
	SecondaryText      lipgloss.AdaptiveColor
	InvertedText       lipgloss.AdaptiveColor
	SuccessText        lipgloss.AdaptiveColor
	WarningText        lipgloss.AdaptiveColor
	ErrorText          lipgloss.AdaptiveColor
}

var DefaultTheme = Theme{
	PrimaryBorder:      lipgloss.AdaptiveColor{Light: "13", Dark: "8"},
	SecondaryBorder:    lipgloss.AdaptiveColor{Light: "8", Dark: "11"},
	SelectedBackground: lipgloss.AdaptiveColor{Light: "6", Dark: "8"},
	FaintBorder:        lipgloss.AdaptiveColor{Light: "7", Dark: "8"},
	PrimaryText:        lipgloss.AdaptiveColor{Light: "0", Dark: "15"},
	SecondaryText:      lipgloss.AdaptiveColor{Light: "8", Dark: "7"},
	FaintText:          lipgloss.AdaptiveColor{Light: "7", Dark: "7"},
	InvertedText:       lipgloss.AdaptiveColor{Light: "15", Dark: "0"},
	SuccessText:        lipgloss.AdaptiveColor{Light: "2", Dark: "2"},
	WarningText:        lipgloss.AdaptiveColor{Light: "3", Dark: "3"},
	ErrorText:          lipgloss.AdaptiveColor{Light: "1", Dark: "1"},
}
