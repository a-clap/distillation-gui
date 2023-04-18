package backend

import (
	"context"
	"fmt"
	"time"

	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/iot/pkg/ds18b20"
	"github.com/a-clap/logging"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Backend struct {
	ctx          context.Context
	eventEmitter *eventEmitter
	dsChan       chan error
	interval     time.Duration
}

func New(opts ...Option) (*Backend, error) {
	b := &Backend{
		eventEmitter: newEventEmitter(),
		dsChan:       make(chan error, 5),
		interval:     time.Second,
	}
	for _, opt := range opts {
		if err := opt(b); err != nil {
			return nil, err
		}
	}

	return b, nil

}

// Startup is called by Wails on application startup
func (b *Backend) Startup(ctx context.Context) {
	b.ctx = ctx
	b.eventEmitter.init(ctx)

	go func() {
		for {
			<-time.After(1 * time.Second)
			fmt.Println("emitting...")
			runtime.EventsEmit(b.ctx, "args", "blah", "adam", ds18b20.SensorConfig{})
		}
	}()
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

func (b *Backend) DSGet() []parameters.DS {
	return ds.Get()
}

func (b *Backend) DSSetCorrection(id string, correction float64) {
	if err := ds.SetCorrection(id, correction); err != nil {
		log.Error("SetCorrection error ", logging.String("ID", id), logging.Float64("correction", correction))
	}
}

func (b *Backend) DSEnable(id string, ena bool) {
	if err := ds.Enable(id, ena); err != nil {
		log.Error("DSEnable error ", logging.String("ID", id), logging.Bool("enable", ena))
	}
}

func (b *Backend) DSSetSamples(id string, samples uint) {
	if err := ds.SetSamples(id, samples); err != nil {
		log.Error("SetSamples error ", logging.String("ID", id), logging.Uint("correction", samples))
	}

}
func (b *Backend) DSSetResolution(id string, res uint) {
	if err := ds.SetResolution(id, ds18b20.Resolution(res)); err != nil {
		log.Error("SetResolution error ", logging.String("ID", id), logging.Uint("resolution", res))
	}
}
