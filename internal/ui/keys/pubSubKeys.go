package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type PubSubKeyMap struct {
	Submit      key.Binding
	Subscribe   key.Binding
	Unsubscribe key.Binding
}

var PubSubKeys = &PubSubKeyMap{
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "submit subject for sub"),
	),
	Subscribe: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "subscribe to a subject"),
	),
	Unsubscribe: key.NewBinding(
		key.WithKeys("u"),
		key.WithHelp("u", "unsubscribe"),
	),
}

func PubSubFullHelp() []key.Binding {

	return []key.Binding{
		PubSubKeys.Submit,
		PubSubKeys.Subscribe,
		PubSubKeys.Unsubscribe,
	}
}
