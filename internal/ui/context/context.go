package context

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/ui/style"
	"github.com/mananapr/jetea/internal/ui/theme"
	"github.com/nats-io/nats.go"
)

type NATSConnectionStatusType string
type SubscriptionStatusType int

var NATSConnectionStatus = struct {
	CONNECTED    NATSConnectionStatusType
	CONNECTING   NATSConnectionStatusType
	DISCONNECTED NATSConnectionStatusType
}{
	CONNECTED:    "connected",
	CONNECTING:   "connecting",
	DISCONNECTED: "disconnected",
}

var SubscriptionStatus = struct {
	UNSUBSCRIBED SubscriptionStatusType
	SUBSCRIBING  SubscriptionStatusType
	SUBSCRIBED   SubscriptionStatusType
}{
	UNSUBSCRIBED: 0,
	SUBSCRIBING:  1,
	SUBSCRIBED:   2,
}

type AppContext struct {
	Program          *tea.Program
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
	ActiveSub        *nats.Subscription
	ActiveInput      bool
	ConnectedServer  *string
	ConnectionStatus NATSConnectionStatusType
	SubStatus        SubscriptionStatusType
	SubSubject       string
	Messages         []string
}
