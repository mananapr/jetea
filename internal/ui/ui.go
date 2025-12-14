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
	"github.com/mananapr/jetea/internal/ui/view"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	log "github.com/charmbracelet/log"
)

type initMsg struct{}
type cfgErrorMsg struct {
	err error
}

type Tab struct {
	View  config.ViewType
	Title string
}

type Model struct {
	ctx              *context.AppContext
	keys             *keys.KeyMap
	footer           footer.Model
	views            []view.View
	tabs             []Tab
	currentSelection int
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
	m.views = []view.View{
		view.NewServerView(m.ctx),
		view.NewPubSubView(m.ctx),
		view.NewJetstreamView(m.ctx),
		view.NewRequestReplyView(m.ctx),
	}
	m.tabs = []Tab{
		{View: config.ServerSelectionView, Title: "Servers"},
		{View: config.PubSubView, Title: "Pub/Sub"},
		{View: config.JetstreamView, Title: "JetStream"},
		{View: config.RequestReplyView, Title: "Req/Reply"},
	}
	m.currentSelection = 0
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

	for _, v := range m.views {
		if sv, ok := v.(*view.ServerView); ok {
			sv.InitServerList()
		}
	}

	return initMsg{}
}

func (m *Model) recalcLayout() {
	h := m.ctx.ScreenHeight

	if m.footer.ShowAll {
		m.ctx.ContentHeight = h - style.ExpandedHelpHeight - style.TabHeight
	} else {
		m.ctx.ContentHeight = h - style.FooterHeight - style.TabHeight
	}
}

func (m *Model) resizeActiveView() tea.Cmd {
	return func() tea.Msg {
		return tea.WindowSizeMsg{
			Width:  m.ctx.ScreenWidth,
			Height: m.ctx.ScreenHeight,
		}
	}
}

func (m Model) tabBarView() string {
	var out []string

	for i, tab := range m.tabs {
		if i == m.currentSelection {
			out = append(out, m.ctx.Styles.ActiveTabStyle.Render(tab.Title))
		} else {
			out = append(out, m.ctx.Styles.InactiveTabStyle.Render(tab.Title))
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, out...)
}

func (m *Model) handleWindowResize(msg tea.WindowSizeMsg) {
	log.Info("window resized", "width", msg.Width, "height", msg.Height)
	m.footer.SetWidth(msg.Width)
	m.ctx.ScreenWidth = msg.Width
	m.ctx.ScreenHeight = msg.Height
	if m.footer.ShowAll {
		m.ctx.ContentHeight = msg.Height - style.ExpandedHelpHeight - style.TabHeight
	} else {
		m.ctx.ContentHeight = msg.Height - style.FooterHeight - style.TabHeight
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
			m.recalcLayout()
			return m, m.resizeActiveView()
		case key.Matches(msg, m.keys.NextTab):
			viewCount := len(m.views)
			if m.currentSelection == viewCount-1 {
				m.currentSelection = 0
			} else {
				m.currentSelection += 1
			}
			m.ctx.View = m.views[m.currentSelection].Type()
		case key.Matches(msg, m.keys.PrevTab):
			viewCount := len(m.views)
			if m.currentSelection == 0 {
				m.currentSelection = viewCount - 1
			} else {
				m.currentSelection -= 1
			}
			m.ctx.View = m.views[m.currentSelection].Type()
		}

	case tea.WindowSizeMsg:
		m.handleWindowResize(msg)

	}

	var cmd tea.Cmd
	_, cmd = m.views[m.currentSelection].Update(msg)
	for i, v := range m.views {
		if m.ctx.View == v.Type() {
			m.currentSelection = i
			break
		}
	}

	return m, cmd
}

func (m Model) View() string {
	tabs := m.tabBarView()
	content := lipgloss.Place(m.ctx.ContentWidth, m.ctx.ContentHeight, lipgloss.Top, lipgloss.Top, m.views[m.currentSelection].View())
	footer := m.footer.View()

	return lipgloss.JoinVertical(lipgloss.Top, tabs, content, footer)
}
