package footer

import (
	"strings"

	bHelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"

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
		Background(m.ctx.Theme.SecondaryBorder).
		Foreground(m.ctx.Theme.PrimaryText).
		Padding(0, 1).
		Render("? help")

	connectedIndicator := lipgloss.NewStyle().
		Background(m.ctx.Theme.SuccessText).
		Foreground(m.ctx.Theme.FaintBorder).
		Padding(0, 1)

	connectingIndicator := lipgloss.NewStyle().
		Background(m.ctx.Theme.WarningText).
		Foreground(m.ctx.Theme.FaintBorder).
		Padding(0, 1)

	serverNameIndicator := lipgloss.NewStyle().
		Background(m.ctx.Theme.SecondaryBorder).
		Foreground(m.ctx.Theme.PrimaryText).
		Padding(0, 1)

	disconnectedIndicator := lipgloss.NewStyle().
		Background(m.ctx.Theme.ErrorText).
		Foreground(m.ctx.Theme.FaintBorder).
		Padding(0, 1)

	leftSection := ""
	if m.ctx.ConnectedServer != nil {
		switch m.ctx.ConnectionStatus {
		case context.NATSConnectionStatus.CONNECTED:
			m.leftSection = util.StringPtr(connectedIndicator.Render("CONNECTED") + serverNameIndicator.Render(*m.ctx.ConnectedServer))
		case context.NATSConnectionStatus.CONNECTING:
			m.leftSection = util.StringPtr(connectingIndicator.Render("CONNECTING") + serverNameIndicator.Render(*m.ctx.ConnectedServer))
		}
		log.Info("connection status", "status", m.ctx.ConnectionStatus)
	} else {
		m.leftSection = util.StringPtr(disconnectedIndicator.Render("DISCONNECTED"))
	}
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
