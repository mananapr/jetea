package style

import (
	bHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/mananapr/jetea/internal/ui/theme"
)

var (
	FooterHeight       = 1
	ExpandedHelpHeight = 10
	TabHeight          = 1
)

type AppStyles struct {
	MainTextStyle    lipgloss.Style
	FaintTextStyle   lipgloss.Style
	FooterStyle      lipgloss.Style
	ActiveTabStyle   lipgloss.Style
	InactiveTabStyle lipgloss.Style
	ErrorStyle       lipgloss.Style

	Help struct {
		Text        lipgloss.Style
		KeyText     lipgloss.Style
		BubbleStyle bHelp.Styles
	}
}

func BuildStyles(theme theme.Theme) AppStyles {
	var style AppStyles

	style.MainTextStyle = lipgloss.NewStyle().Foreground(theme.PrimaryText).Bold(true)
	style.FaintTextStyle = lipgloss.NewStyle().Foreground(theme.FaintText)
	style.FooterStyle = lipgloss.NewStyle().Background(theme.SelectedBackground).Height(FooterHeight)
	style.ActiveTabStyle = lipgloss.NewStyle().Foreground(theme.PrimaryText).Background(theme.SelectedBackground).Bold(true).Padding(0, 2).Height(TabHeight)
	style.InactiveTabStyle = lipgloss.NewStyle().Faint(true).Padding(0, 2).Height(TabHeight)

	style.ErrorStyle = style.FooterStyle.Foreground(theme.ErrorText).MaxHeight(FooterHeight)

	style.Help.Text = lipgloss.NewStyle().Foreground(theme.SecondaryText)
	style.Help.KeyText = lipgloss.NewStyle().Foreground(theme.PrimaryText)
	style.Help.BubbleStyle = bHelp.Styles{
		ShortDesc:      style.Help.Text.Foreground(theme.FaintText),
		FullDesc:       style.Help.Text.Foreground(theme.FaintText),
		ShortSeparator: style.Help.Text.Foreground(theme.SecondaryBorder),
		FullSeparator:  style.Help.Text,
		FullKey:        style.Help.KeyText,
		ShortKey:       style.Help.KeyText,
		Ellipsis:       style.Help.Text,
	}

	return style
}
