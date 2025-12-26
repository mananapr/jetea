package view

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/ui/context"
	"github.com/mananapr/jetea/internal/ui/keys"
)

type PubSubView struct {
	ctx *context.AppContext
}

func NewPubSubView(ctx *context.AppContext) *PubSubView {
	return &PubSubView{ctx: ctx}
}

func (m *PubSubView) Type() config.ViewType {
	return config.PubSubView
}

func (m *PubSubView) Init() tea.Cmd {
	return nil
}

func (m *PubSubView) Update(msg tea.Msg) (View, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.PubSubKeys.Select):
			m.ctx.View = config.PubSubView
		}
	}
	return m, nil
}

func (m *PubSubView) View() string {
	return lipgloss.NewStyle().Render("PubSub View")
}
