package view

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/ui/context"
	"github.com/mananapr/jetea/internal/ui/keys"
)

type ServerView struct {
	ctx *context.AppContext
}

func NewServerView(ctx *context.AppContext) *ServerView {
	return &ServerView{ctx: ctx}
}

func (m *ServerView) Type() config.ViewType {
	return config.ServerSelectionView
}

func (m *ServerView) Init() tea.Cmd {
	return nil
}

func (m *ServerView) Update(msg tea.Msg) (View, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.ServerKeys.Select):
			m.ctx.View = config.PubSubView
		}
	}
	return m, nil
}

func (m *ServerView) View() string {
	return "Server Selection View"
}
