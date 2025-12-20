package nats

import (
	"github.com/mananapr/jetea/internal/config"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nats-io/nats.go"
)

type Client struct {
	conn *nats.Conn
}

func Connect(serverCfg config.NATSConfig) tea.Cmd {
	return func() tea.Msg {
		opts := []nats.Option{
			nats.Name("jetea"),
		}

		if serverCfg.TLS {
			opts = append(opts, nats.Secure())
		}

		if serverCfg.User != nil {
			opts = append(opts, nats.UserInfo(*serverCfg.User, *serverCfg.Password))
		}

		conn, err := nats.Connect(serverCfg.Hostname, opts...)
		if err != nil {
			return ConnectionErrorMsg{
				ServerName: serverCfg.Name,
				Err:        err,
			}
		}

		return ConnectedMsg{
			ServerName: serverCfg.Name,
			Conn:       conn,
		}
	}
}

func Disconnect(conn *nats.Conn) tea.Cmd {
	return func() tea.Msg {
		conn.Close()
		return DisconnectedMsg{}
	}
}
