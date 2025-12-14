package ui

import (
	"os"
	"time"

	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/ui/context"
	"github.com/mananapr/jetea/internal/ui/footer"
	"github.com/mananapr/jetea/internal/ui/keys"
	"github.com/mananapr/jetea/internal/ui/style"
	"github.com/mananapr/jetea/internal/ui/theme"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	log "github.com/charmbracelet/log"
)

type initMsg struct{}
type cfgErrorMsg struct {
	err error
}

type Model struct {
	ctx    *context.AppContext
	keys   *keys.KeyMap
	footer footer.Model
}

func NewModel(cfgFlag string) Model {
	m := Model{}

	m.keys = keys.Keys

	m.ctx = &context.AppContext{
		ConfigFlag: cfgFlag,
		View:       config.ServerSelectionView,
		Theme:      theme.DefaultTheme,
		Styles:     style.BuildStyles(theme.DefaultTheme),
	}

	m.footer = footer.NewModel(m.ctx)

	return m
}

func showConfigError(err error) {
	styles := log.DefaultStyles()
	styles.Key = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true)
	styles.Separator = lipgloss.NewStyle()

	logger := log.New(os.Stderr)
	logger.SetStyles(styles)
	logger.SetTimeFormat(time.Kitchen)
	logger.SetPrefix("config")
	logger.SetReportCaller(true)

	logger.Error("failed to parse config", "error", err)
}

func (m *Model) initProgram() tea.Msg {
	cfg, cfgErr := config.LoadConfig(m.ctx.ConfigFlag)
	if cfgErr != nil {
		return cfgErrorMsg{err: cfgErr}
	}
	m.ctx.Config = cfg

	return initMsg{}
}

func (m *Model) handleWindowResize(msg tea.WindowSizeMsg) {
	log.Info("window resized", "width", msg.Width, "height", msg.Height)
	m.footer.SetWidth(msg.Width)
	m.ctx.ScreenWidth = msg.Width
	m.ctx.ScreenHeight = msg.Height
	if m.footer.ShowAll {
		m.ctx.ContentHeight = msg.Height - style.ExpandedHelpHeight
	} else {
		m.ctx.ContentHeight = msg.Height - style.FooterHeight
	}
	m.ctx.ContentWidth = msg.Width
	log.Info("content resized", "width", m.ctx.ContentWidth, "height", m.ctx.ContentHeight)
}

func (m Model) Init() tea.Cmd {
	return m.initProgram
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case cfgErrorMsg:
		return m, tea.Sequence(
			tea.ExitAltScreen,
			func() tea.Msg {
				showConfigError(msg.err)
				return nil
			},
			tea.Quit,
		)

	case tea.KeyMsg:
		log.Debug("Key pressed", "key", msg.String())

		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.footer.ShowAll = !m.footer.ShowAll
		}

	case tea.WindowSizeMsg:
		m.handleWindowResize(msg)

	}

	return m, nil
}

func (m Model) View() string {
	content := lipgloss.Place(m.ctx.ContentWidth, m.ctx.ContentHeight, lipgloss.Center, lipgloss.Center, "main content")
	footer := m.footer.View()

	return lipgloss.JoinVertical(lipgloss.Top, content, footer)
}
