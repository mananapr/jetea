package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type RequestReplyKeyMap struct {
	Select key.Binding
}

var RequestReplyKeys = &RequestReplyKeyMap{
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select req/reply"),
	),
}

func RequestReplyFullHelp() []key.Binding {

	return []key.Binding{
		RequestReplyKeys.Select,
	}
}
