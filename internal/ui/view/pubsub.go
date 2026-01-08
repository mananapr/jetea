package view

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/nats"
	"github.com/mananapr/jetea/internal/ui/context"
	"github.com/mananapr/jetea/internal/ui/keys"
)

type msgItem struct {
	index int
	body  string
}

func (m msgItem) Title() string {
	return fmt.Sprintf("Msg %d", m.index)
}

func (m msgItem) Description() string {
	if len(m.body) > 40 {
		return m.body[:40] + "…"
	}
	return m.body
}

func (m msgItem) FilterValue() string { return "" }

type PubSubView struct {
	ctx      *context.AppContext
	input    textinput.Model
	msgList  list.Model
	selected string
}

func NewPubSubView(ctx *context.AppContext) *PubSubView {
	ti := textinput.New()
	ti.Placeholder = "subject (e.g. foo.bar)"

	d := list.NewDefaultDelegate()
	l := list.New(nil, d, 1, 1)
	l.SetFilteringEnabled(false)
	l.SetShowTitle(false)
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)
	l.Styles.NoItems = lipgloss.NewStyle()

	return &PubSubView{ctx: ctx, input: ti, msgList: l}
}

func (m *PubSubView) Type() config.ViewType {
	return config.PubSubView
}

func (m *PubSubView) Init() tea.Cmd {
	return nil
}

func (m *PubSubView) Update(msg tea.Msg) (View, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.input.Width = m.ctx.ContentWidth
		m.msgList.SetSize(m.ctx.ContentWidth/2, m.ctx.ContentHeight-3)
		log.Info("set input contentwidth", "width", m.ctx.ContentWidth)

	case nats.SubscribedMsg:
		m.ctx.SubStatus = context.SubscriptionStatus.SUBSCRIBED

	case nats.UnsubscribedMsg:
		m.ctx.SubStatus = context.SubscriptionStatus.UNSUBSCRIBED

	case nats.MsgReceived:
		log.Info("msg received", "sub", m.input.Value(), "msg", msg.Body)
		m.ctx.Messages = append(m.ctx.Messages, msg.Body)
		var itms []list.Item
		for i, nm := range m.ctx.Messages {
			itms = append(itms, msgItem{index: i, body: nm})
		}
		m.msgList.SetItems(itms)

	case nats.ConnectionErrorMsg:
		m.ctx.SubStatus = context.SubscriptionStatus.UNSUBSCRIBED
		m.ctx.Error = msg.Err
		log.Info("sub error", "error", msg.Err.Error())

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.PubSubKeys.Submit):
			m.input.Blur()
			m.ctx.ActiveInput = false
			sub := m.input.Value()
			if sub == "" {
				m.ctx.Error = fmt.Errorf("Subject cannot be empty!")
				return m, nil
			} else {
				m.ctx.SubStatus = context.SubscriptionStatus.SUBSCRIBING
				log.Info("subscribing", "subject", m.input.Value())
				return m, nats.Subscribe(m.ctx, m.input.Value())
			}

		case m.ctx.ActiveInput:
			var cmd tea.Cmd
			m.input, cmd = m.input.Update(msg)
			return m, cmd

		case key.Matches(msg, keys.PubSubKeys.Subscribe):
			m.input.Focus()
			m.ctx.ActiveInput = true
			log.Info("sub input focus")
			return m, nil

		case key.Matches(msg, keys.PubSubKeys.Unsubscribe):
			log.Info("unsubscribing", "subject", m.input.Value())
			m.input.SetValue("")
			return m, nats.Unsubscribe(m.ctx)
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	m.msgList, cmd = m.msgList.Update(msg)

	if it, ok := m.msgList.SelectedItem().(msgItem); ok {
		m.selected = it.body
	}

	if m.ctx.ActiveInput {
		m.input.TextStyle = lipgloss.NewStyle().Foreground(m.ctx.Theme.PrimaryText)
		m.input.PlaceholderStyle = lipgloss.NewStyle().Foreground(m.ctx.Theme.FaintBorder)
		m.input.PromptStyle = lipgloss.NewStyle().Foreground(m.ctx.Theme.SuccessText)
	} else if m.ctx.SubStatus == context.SubscriptionStatus.SUBSCRIBED {
		m.input.TextStyle = lipgloss.NewStyle().Bold(true).Foreground(m.ctx.Theme.SuccessText)
		m.input.PromptStyle = lipgloss.NewStyle().Bold(true).Foreground(m.ctx.Theme.SuccessText)
	} else {
		m.input.TextStyle = lipgloss.NewStyle().Foreground(m.ctx.Theme.FaintBorder)
		m.input.PlaceholderStyle = lipgloss.NewStyle().Foreground(m.ctx.Theme.FaintBorder)
		m.input.PromptStyle = lipgloss.NewStyle().Foreground(m.ctx.Theme.FaintBorder)
	}

	return m, cmd
}

func (m *PubSubView) View() string {
	left := lipgloss.NewStyle().
		Width(m.ctx.ContentWidth / 2).
		Height(m.ctx.ContentHeight - 3).
		Render(m.msgList.View())

	if len(m.msgList.Items()) == 0 {
		left = lipgloss.NewStyle().
			Width(m.ctx.ContentWidth / 2).
			Height(m.ctx.ContentHeight - 3).
			Render("No messages.")
	}

	right := lipgloss.NewStyle().
		Width(m.ctx.ContentWidth / 2).
		Height(m.ctx.ContentHeight - 3).
		Padding(1).
		Render(m.selected)

	inputPrompt := lipgloss.NewStyle().Foreground(m.ctx.Theme.FaintBorder).Render("Subscribe")
	if m.ctx.ActiveInput {
		inputPrompt = lipgloss.NewStyle().Foreground(m.ctx.Theme.SuccessText).Render("Subscribe")
	} else if m.ctx.SubStatus == context.SubscriptionStatus.SUBSCRIBED {
		inputPrompt = lipgloss.NewStyle().Bold(true).Foreground(m.ctx.Theme.SuccessText).Render("Subscribe")
	}

	input := lipgloss.NewStyle().
		Width(m.ctx.ContentWidth).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		Foreground(m.ctx.Theme.SecondaryText).
		Render(fmt.Sprintf("%s %s", inputPrompt, m.input.View()))

	return lipgloss.JoinVertical(
		lipgloss.Top,
		input,
		lipgloss.JoinHorizontal(lipgloss.Top, left, right),
	)
}
