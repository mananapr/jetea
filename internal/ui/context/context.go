package context

import (
	"github.com/mananapr/jetea/internal/config"
	"github.com/mananapr/jetea/internal/ui/style"
	"github.com/mananapr/jetea/internal/ui/theme"
)

type AppContext struct {
	ScreenHeight  int
	ScreenWidth   int
	ContentWidth  int
	ContentHeight int
	ConfigFlag    string
	Config        config.Config
	View          config.ViewType
	Theme         theme.Theme
	Styles        style.AppStyles
	Error         error
}
