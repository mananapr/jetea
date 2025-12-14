package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type PubSubKeyMap struct {
	Select key.Binding
}

var PubSubKeys = &PubSubKeyMap{
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select pub/sub"),
	),
}

func PubSubFullHelp() []key.Binding {

	return []key.Binding{
		PubSubKeys.Select,
	}
}
