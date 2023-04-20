package backend

import (
	"context"
	"time"

	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/distillation-gui/backend/gpio"
	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/distillation-gui/backend/pt"
	"github.com/a-clap/distillation-gui/backend/wifi"
	"github.com/a-clap/iot/pkg/ds18b20"
	embeddedgpio "github.com/a-clap/iot/pkg/embedded/gpio"
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
	b.handleErrors()
}

func (b *Backend) handleErrors() {
	if b.dsChan != nil {
		go func() {
			for err := range b.dsChan {
				logger.Warn("Error from DS", logging.String("error", err.Error()))
				// TODO: How to get ID based on error?
				b.eventEmitter.OnError(0)
			}
		}()
	}
	if b.ptChan != nil {
		go func() {
			for err := range b.ptChan {
				logger.Warn("Error from PT", logging.String("error", err.Error()))
				// TODO: How to get ID based on error?
				b.eventEmitter.OnError(0)
			}
		}()
	}

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
		b.eventEmitter.OnError(ErrDSSetCorrection)
	}
}

func (b *Backend) DSEnable(id string, ena bool) {
	if err := ds.Enable(id, ena); err != nil {
		logger.Error("DSEnable error ", logging.String("ID", id), logging.Bool("enable", ena))
		b.eventEmitter.OnError(ErrDSEnable)
	}
}

func (b *Backend) DSSetSamples(id string, samples uint) {
	if err := ds.SetSamples(id, samples); err != nil {
		logger.Error("SetSamples error ", logging.String("ID", id), logging.Uint("samples", samples))
		b.eventEmitter.OnError(ErrDSSetSamples)
	}

}
func (b *Backend) DSSetResolution(id string, res uint) {
	logger.Debug("SetResolution", logging.String("ID", id), logging.Uint("resolution", res))
	if err := ds.SetResolution(id, ds18b20.Resolution(res)); err != nil {
		logger.Error("SetResolution error ", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrDSSetResolution)
	}
}

func (b *Backend) PTGet() []parameters.PT {
	return pt.Get()
}

func (b *Backend) PTSetCorrection(id string, correction float64) {
	logger.Debug("SetCorrection ", logging.String("ID", id), logging.Float64("correction", correction))
	if err := pt.SetCorrection(id, correction); err != nil {
		logger.Error("SetCorrection error ", logging.String("ID", id), logging.Float64("correction", correction))
		b.eventEmitter.OnError(ErrPTSetCorrection)
	}
}

func (b *Backend) PTEnable(id string, ena bool) {
	logger.Debug("PTEnable ", logging.String("ID", id), logging.Bool("enable", ena))
	if err := pt.Enable(id, ena); err != nil {
		logger.Error("PTEnable error ", logging.String("ID", id), logging.Bool("enable", ena))
		b.eventEmitter.OnError(ErrPTEnable)
	}
}

func (b *Backend) PTSetSamples(id string, samples uint) {
	logger.Debug("SetSamples ", logging.String("ID", id), logging.Uint("samples", samples))
	if err := pt.SetSamples(id, samples); err != nil {
		logger.Error("SetSamples error ", logging.String("ID", id), logging.Uint("samples", samples))
		b.eventEmitter.OnError(ErrPTSetSamples)
	}
}

func (b *Backend) WifiAPList() []string {
	logger.Debug("WifiAPList")
	aps, err := wifi.AP()
	if err != nil {
		logger.Error("WifiAPList", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrWIFIAPList)
		return nil
	}
	return aps
}

func (b *Backend) GPIOGet() []parameters.GPIO {
	logger.Debug("GPIOGet")
	return gpio.Get()
}

func (b *Backend) GPIOSetActiveLevel(id string, active embeddedgpio.ActiveLevel) {
	logger.Debug("GPIOSetActiveLevel", logging.String("id", id), logging.Int("active", int(active)))
	if err := gpio.SetActiveLevel(id, active); err != nil {
		logger.Error("GPIOSetActiveLevel", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrGPIOSetActiveLevel)
	}
}

func (b *Backend) GPIOSetState(id string, value bool) {
	logger.Debug("GPIOSetState", logging.String("id", id), logging.Bool("value", value))
	if err := gpio.SetState(id, value); err != nil {
		logger.Error("GPIOSetState", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrGPIOSetState)
	}
}
