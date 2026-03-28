package nats

import (
	"github.com/nats-io/nats.go"
)

type ConnectMsg struct {
	ServerName string
}

type ConnectedMsg struct {
	ServerName string
	Conn       *nats.Conn
}

type ConnectionErrorMsg struct {
	ServerName string
	Err        error
}

type DisconnectedMsg struct {
	ServerName string
}

type MsgReceived struct {
	Body string
}

type SubscribedMsg struct {
	Subject string
}

type UnsubscribedMsg struct{}
