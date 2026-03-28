package keys

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"

	"github.com/mananapr/jetea/internal/config"
)

type KeyMap struct {
	viewType config.ViewType
	Up       key.Binding
	Down     key.Binding
	NextTab  key.Binding
	PrevTab  key.Binding
	Refresh  key.Binding
	Help     key.Binding
	Quit     key.Binding
}

var Keys = &KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	NextTab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("⭾", "next tab"),
	),
	PrevTab: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+⭾", "previous tab"),
	),
	Refresh: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refresh"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func CreateKeyMapForView(viewType config.ViewType) help.KeyMap {
	Keys.viewType = viewType
	return Keys
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	var additionalKeys []key.Binding

	switch k.viewType {
	case config.ServerSelectionView:
		additionalKeys = ServerFullHelp()
	case config.PubSubView:
		additionalKeys = PubSubFullHelp()
	case config.JetstreamView:
		additionalKeys = JetstreamFullHelp()
	case config.RequestReplyView:
		additionalKeys = RequestReplyFullHelp()
	}

	sections := [][]key.Binding{
		k.NavigationKeys(),
		k.AppKeys(),
		additionalKeys,
	}

	return sections
}

func (k KeyMap) NavigationKeys() []key.Binding {
	return []key.Binding{
		k.Up,
		k.Down,
		k.Help,
		k.Quit,
	}
}

func (k KeyMap) AppKeys() []key.Binding {
	return []key.Binding{
		k.Refresh,
		k.NextTab,
		k.PrevTab,
	}
}
