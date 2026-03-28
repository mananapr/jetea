package keys

import (
	"github.com/charmbracelet/bubbles/key"
)

type ServerKeyMap struct {
	Select key.Binding
}

var ServerKeys = &ServerKeyMap{
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select server"),
	),
}

func ServerFullHelp() []key.Binding {

	return []key.Binding{
		ServerKeys.Select,
	}
}
