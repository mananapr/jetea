package context

import (
	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/ui/style"
	"github.com/mananapr/jetea/internal/ui/theme"
	"github.com/nats-io/nats.go"
)

type NATSConnectionStatusType string

var NATSConnectionStatus = struct {
	CONNECTED    NATSConnectionStatusType
	CONNECTING   NATSConnectionStatusType
	DISCONNECTED NATSConnectionStatusType
}{
	CONNECTED:    "connected",
	CONNECTING:   "connecting",
	DISCONNECTED: "disconnected",
}

type AppContext struct {
	ScreenHeight     int
	ScreenWidth      int
	ContentWidth     int
	ContentHeight    int
	ConfigFlag       string
	Config           config.Config
	View             config.ViewType
	Theme            theme.Theme
	Styles           style.AppStyles
	Error            error
	NatsConnection   *nats.Conn
	ConnectedServer  *string
	ConnectionStatus NATSConnectionStatusType
}
