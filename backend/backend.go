package backend

import (
	"context"
	"time"

	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/distillation-gui/backend/pt"
	"github.com/a-clap/distillation-gui/backend/wifi"
	"github.com/a-clap/iot/pkg/ds18b20"
	"github.com/a-clap/logging"
)

var (
	logger = logging.GetLogger()
)

type Backend struct {
	ctx          context.Context
	eventEmitter *eventEmitter
	dsChan       chan error
	ptChan       chan error
	interval     time.Duration
}

func New(opts ...Option) (*Backend, error) {
	b := &Backend{
		eventEmitter: newEventEmitter(),
		dsChan:       make(chan error, 5),
		ptChan:       make(chan error, 5),
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
		logger.Error("SetCorrection error ", logging.String("ID", id), logging.Float64("correction", correction))
	}
}

func (b *Backend) DSEnable(id string, ena bool) {
	if err := ds.Enable(id, ena); err != nil {
		logger.Error("DSEnable error ", logging.String("ID", id), logging.Bool("enable", ena))
	}
}

func (b *Backend) DSSetSamples(id string, samples uint) {
	if err := ds.SetSamples(id, samples); err != nil {
		logger.Error("SetSamples error ", logging.String("ID", id), logging.Uint("samples", samples))
	}

}
func (b *Backend) DSSetResolution(id string, res uint) {
	logger.Debug("SetResolution", logging.String("ID", id), logging.Uint("resolution", res))
	if err := ds.SetResolution(id, ds18b20.Resolution(res)); err != nil {
		logger.Error("SetResolution error ", logging.String("error", err.Error()))
	}
}

func (b *Backend) PTGet() []parameters.PT {
	return pt.Get()
}

func (b *Backend) PTSetCorrection(id string, correction float64) {
	logger.Debug("SetCorrection ", logging.String("ID", id), logging.Float64("correction", correction))
	if err := pt.SetCorrection(id, correction); err != nil {
		logger.Error("SetCorrection error ", logging.String("ID", id), logging.Float64("correction", correction))
	}
}

func (b *Backend) PTEnable(id string, ena bool) {
	logger.Debug("PTEnable ", logging.String("ID", id), logging.Bool("enable", ena))
	if err := pt.Enable(id, ena); err != nil {
		logger.Error("PTEnable error ", logging.String("ID", id), logging.Bool("enable", ena))
	}
}

func (b *Backend) PTSetSamples(id string, samples uint) {
	logger.Debug("SetSamples ", logging.String("ID", id), logging.Uint("samples", samples))
	if err := pt.SetSamples(id, samples); err != nil {
		logger.Error("SetSamples error ", logging.String("ID", id), logging.Uint("samples", samples))
	}
}

func (b *Backend) WifiAPList() []string {
	logger.Debug("WifiAPList")
	aps, err := wifi.AP()
	if err != nil {
		logger.Error("Failed to get ap list", logging.String("error", err.Error()))
		return nil
	}
	return aps
}
