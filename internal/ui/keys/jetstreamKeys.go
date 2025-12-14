package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type JetstreamKeyMap struct {
	Select key.Binding
}

var JestreamKeys = &JetstreamKeyMap{
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select stream"),
	),
}

func JetstreamFullHelp() []key.Binding {

	return []key.Binding{
		JestreamKeys.Select,
	}
}
