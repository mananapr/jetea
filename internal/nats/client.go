package nats

import (
	"fmt"

	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/ui/context"

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

func Subscribe(ctx *context.AppContext, subject string) tea.Cmd {
	return func() tea.Msg {
		if ctx.NatsConnection == nil {
			ctx.Error = fmt.Errorf("Not connected to any server!")
			return nil
		}

		if ctx.ActiveSub != nil {
			_ = ctx.ActiveSub.Unsubscribe()
			ctx.ActiveSub = nil
		}

		sub, err := ctx.NatsConnection.Subscribe(subject, func(m *nats.Msg) {
			ctx.Program.Send(MsgReceived{
				Body: string(m.Data),
			},
			)
		})

		if err != nil {
			return ConnectionErrorMsg{
				ServerName: subject,
				Err:        err,
			}
		}

		ctx.ActiveSub = sub

		return SubscribedMsg{Subject: subject}
	}
}

func Unsubscribe(ctx *context.AppContext) tea.Cmd {
	return func() tea.Msg {
		if ctx.ActiveSub != nil {
			_ = ctx.ActiveSub.Unsubscribe()
			ctx.ActiveSub = nil
		}
		return UnsubscribedMsg{}
	}
}
