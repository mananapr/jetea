package view

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/ui/context"
	"github.com/mananapr/jetea/internal/ui/keys"
)

type JetstreamView struct {
	ctx *context.AppContext
}

func NewJetstreamView(ctx *context.AppContext) *JetstreamView {
	return &JetstreamView{ctx: ctx}
}

func (m *JetstreamView) Type() config.ViewType {
	return config.JetstreamView
}

func (m *JetstreamView) Init() tea.Cmd {
	return nil
}

func (m *JetstreamView) Update(msg tea.Msg) (View, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.JestreamKeys.Select):
			m.ctx.View = config.PubSubView
		}
	}
	return m, nil
}

func (m *JetstreamView) View() string {
	return "Jetstream View"
}
