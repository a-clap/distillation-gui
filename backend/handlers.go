package backend

import (
	"context"
	"time"

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
	go func() {
		for {
			e.OnHeaterChange(parameters.Heater{})
			<-time.After(1 * time.Second)
		}

	}()
}

// OnHeaterChange implements heater.Listener
func (e *eventEmitter) OnHeaterChange(heater parameters.Heater) {
	runtime.EventsEmit(e.ctx, NotifyHeaters)
}
