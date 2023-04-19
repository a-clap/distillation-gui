package backend

import (
	"context"

	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/distillation-gui/backend/pt"
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
	pt.AddListener(e)
}

// OnHeaterChange implements heater.Listener
func (e *eventEmitter) OnHeaterChange(h parameters.Heater) {
	runtime.EventsEmit(e.ctx, NotifyHeaters, h)
}

// OnDSConfigChange implements ds.Listener
func (e *eventEmitter) OnDSConfigChange(c parameters.DS) {
	runtime.EventsEmit(e.ctx, NotifyDSConfig, c)
}

// OnDSTemperatureChange implements ds.Listener
func (e *eventEmitter) OnDSTemperatureChange(t parameters.Temperature) {
	runtime.EventsEmit(e.ctx, NotifyDSTemperature, t)
}

// OnPTConfigChange implements pt.Listener
func (e *eventEmitter) OnPTConfigChange(p parameters.PT) {
	runtime.EventsEmit(e.ctx, NotifyPTConfig, p)
}

// OnPTTemperatureChange implements pt.Listener
func (e *eventEmitter) OnPTTemperatureChange(t parameters.Temperature) {
	runtime.EventsEmit(e.ctx, NotifyPTTemperature, t)
}
