package view

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mananapr/jetea/internal/config"
)

type View interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (View, tea.Cmd)
	View() string
	Type() config.ViewType
}
