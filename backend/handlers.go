package backend

import (
	"context"

	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type eventEmitter struct {
	ctx context.Context
}

func newEventEmitter() *eventEmitter {
	return &eventEmitter{}
}

func (e *eventEmitter) init(ctx context.Context) {
	e.ctx = ctx
	heater.AddListener(e)
	ds.AddListener(e)
}

// OnHeaterChange implements heater.Listener
func (e *eventEmitter) OnHeaterChange(parameters.Heater) {
	runtime.EventsEmit(e.ctx, NotifyHeaters)
}

// OnDSConfigChange implements ds.Listener
func (e *eventEmitter) OnDSConfigChange(parameters.DS) {
	runtime.EventsEmit(e.ctx, NotifyDSConfig)
}

// OnDSTemperatureChange implements ds.Listener
func (e *eventEmitter) OnDSTemperatureChange(parameters.Temperature) {
	runtime.EventsEmit(e.ctx, NotifyDSTemperature)
}
