package backend

import (
	"context"

	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/distillation-gui/backend/parameters"
)

type Backend struct {
	ctx          context.Context
	eventEmitter *eventEmitter
}

func New() *Backend {
	b := &Backend{
		eventEmitter: newEventEmitter(),
	}
	return b

}

// Startup is called by Wails on application startup
func (b *Backend) Startup(ctx context.Context) {
	b.ctx = ctx
	b.eventEmitter.init(ctx)
}

// Events returns Event structure - wails need to generate binding for Events methods
func (b *Backend) Events() Events {
	return Events{}
}

func (b *Backend) HeaterEnable(id string, enable bool) {
	heater.Enable(id, enable)
}

func (b *Backend) HeatersGet() []parameters.Heater {
	return heater.Get()
}
