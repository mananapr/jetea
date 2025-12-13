package ui

import (
	"os"
	"time"

	"github.com/mananapr/jetea/internal/config"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	log "github.com/charmbracelet/log"
)

type initMsg struct{}
type cfgErrorMsg struct {
	err error
}

type Model struct {
	cfg string
	msg string
}

func NewModel(cfgFlag string) Model {
	m := Model{}
	m.msg = "Hello World!"
	m.cfg = cfgFlag

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
	cfg, cfgErr := config.LoadConfig(m.cfg)
	if cfgErr != nil {
		return cfgErrorMsg{err: cfgErr}
	}
	log.Debug("config fetched", "cfg", m.cfg, "pCfg", cfg)

	return initMsg{}
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

		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		} else {
			m.msg = "Key Pressed!!"
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := m.msg
	return s
}
