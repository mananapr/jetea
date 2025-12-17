package view

import (
	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/ui/context"
	"github.com/mananapr/jetea/internal/ui/keys"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type item struct {
	title, hostname string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.hostname }
func (i item) FilterValue() string { return i.title }

type ServerView struct {
	ctx        *context.AppContext
	serverList list.Model
}

func NewServerView(ctx *context.AppContext) *ServerView {
	l := list.New(nil, list.NewDefaultDelegate(), 1, 1)

	l.KeyMap.ShowFullHelp.Unbind()
	l.KeyMap.CloseFullHelp.Unbind()
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)

	return &ServerView{ctx: ctx, serverList: l}
}

func (m *ServerView) InitServerList() {
	var items []list.Item

	for _, s := range m.ctx.Config.NATSServers {
		items = append(items, item{
			title:    *s.Name,
			hostname: s.Hostname,
		})
	}

	m.serverList.SetItems(items)

	log.Info("server list init", "items", len(items))
}

func (m *ServerView) Type() config.ViewType {
	return config.ServerSelectionView
}

func (m *ServerView) Init() tea.Cmd {
	return nil
}

func (m *ServerView) Update(msg tea.Msg) (View, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.serverList.SetSize(m.ctx.ContentWidth, m.ctx.ContentHeight)
		log.Info(
			"list render",
			"w", m.serverList.Width(),
			"h", m.serverList.Height(),
			"items", len(m.serverList.Items()),
		)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.ServerKeys.Select):
			m.ctx.View = config.PubSubView
		}
	}

	var cmd tea.Cmd
	m.serverList, cmd = m.serverList.Update(msg)

	return m, cmd
}

func (m *ServerView) View() string {
	return lipgloss.NewStyle().Render(m.serverList.View())
}
