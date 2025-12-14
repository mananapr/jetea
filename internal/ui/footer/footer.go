package footer

import (
	"strings"

	bHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"

	"github.com/mananapr/jetea/internal/ui/context"
	"github.com/mananapr/jetea/internal/ui/keys"
	"github.com/mananapr/jetea/internal/util"
)

type Model struct {
	ctx          *context.AppContext
	leftSection  *string
	rightSection *string
	help         bHelp.Model
	ShowAll      bool
}

func NewModel(ctx *context.AppContext) Model {
	help := bHelp.New()
	help.ShowAll = true
	l := ""
	r := ""
	return Model{
		ctx:          ctx,
		help:         help,
		leftSection:  &l,
		rightSection: &r,
	}
}

func (m Model) View() string {
	var footer string

	helpIndicator := lipgloss.NewStyle().
		Background(m.ctx.Theme.FaintText).
		Foreground(m.ctx.Theme.SelectedBackground).
		Padding(0, 1).
		Render("? help")
	leftSection := ""
	if m.leftSection != nil {
		leftSection = *m.leftSection
	}
	rightSection := ""
	if m.rightSection != nil {
		rightSection = *m.rightSection
	}
	spacing := lipgloss.NewStyle().
		Background(m.ctx.Theme.SelectedBackground).
		Render(
			strings.Repeat(
				" ",
				util.Max(0, m.ctx.ScreenWidth-lipgloss.Width(leftSection)-lipgloss.Width(rightSection)-lipgloss.Width(helpIndicator)),
			),
		)

	footer = m.ctx.Styles.FooterStyle.Render(lipgloss.JoinHorizontal(lipgloss.Top, leftSection, spacing, rightSection, helpIndicator))

	if m.ShowAll {
		keymap := keys.CreateKeyMapForView(m.ctx.View)
		fullHelp := m.help.View(keymap)
		return lipgloss.JoinVertical(lipgloss.Top, footer, fullHelp)
	}

	return footer
}

func (m *Model) SetWidth(width int) {
	m.help.Width = width
}

func (m *Model) UpdateProgramContext(ctx *context.AppContext) {
	m.ctx = ctx
	m.help.Styles = ctx.Styles.Help.BubbleStyle
}

func (m *Model) SetLeftSection(leftSection string) {
	*m.leftSection = leftSection
}

func (m *Model) SetRightSection(rightSection string) {
	*m.rightSection = rightSection
}
