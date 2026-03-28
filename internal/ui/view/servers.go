package view

import (
	"fmt"

	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/nats"
	"github.com/mananapr/jetea/internal/ui/context"
	"github.com/mananapr/jetea/internal/ui/keys"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type item struct {
	title, description string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type ServerView struct {
	ctx        *context.AppContext
	serverList list.Model
}

func NewServerView(ctx *context.AppContext) *ServerView {
	d := list.NewDefaultDelegate()
	d.Styles.SelectedTitle = d.Styles.SelectedTitle.Foreground(ctx.Theme.WarningText).BorderLeftForeground(ctx.Theme.WarningText)
	d.Styles.SelectedDesc = d.Styles.SelectedDesc.Foreground(ctx.Theme.WarningText).BorderLeftForeground(ctx.Theme.WarningText)
	l := list.New(nil, d, 1, 1)

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
		desc := s.Hostname
		if s.TLS {
			desc = fmt.Sprintf("%s | TLS", desc)
		}
		if s.User != nil {
			desc = fmt.Sprintf("%s | BasicAuth", desc)
		}
		items = append(items, item{
			title:       s.Name,
			description: desc,
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
		m.serverList.SetSize(m.ctx.ContentWidth, m.ctx.ContentHeight-1)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.ServerKeys.Select):
			selectedItem := m.serverList.SelectedItem()
			if i, ok := selectedItem.(item); ok {
				for _, s := range m.ctx.Config.NATSServers {
					if s.Name == i.title {
						m.ctx.ConnectionStatus = context.NATSConnectionStatus.CONNECTING
						m.ctx.ConnectedServer = &s.Name
						return m, nats.Connect(s)
					}
				}
			}
		}
	}

	var cmd tea.Cmd
	m.serverList, cmd = m.serverList.Update(msg)

	return m, cmd
}

func (m *ServerView) View() string {
	return lipgloss.NewStyle().Render(m.serverList.View())
}
