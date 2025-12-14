package view

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/ui/context"
	"github.com/mananapr/jetea/internal/ui/keys"
)

type RequestReplyView struct {
	ctx *context.AppContext
}

func NewRequestReplyView(ctx *context.AppContext) *RequestReplyView {
	return &RequestReplyView{ctx: ctx}
}

func (m *RequestReplyView) Type() config.ViewType {
	return config.RequestReplyView
}

func (m *RequestReplyView) Init() tea.Cmd {
	return nil
}

func (m *RequestReplyView) Update(msg tea.Msg) (View, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.RequestReplyKeys.Select):
			m.ctx.View = config.PubSubView
		}
	}
	return m, nil
}

func (m *RequestReplyView) View() string {
	return "Request/Reply View"
}
