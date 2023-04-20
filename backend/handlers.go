package backend

import (
	"context"

	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/distillation-gui/backend/gpio"
	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/distillation-gui/backend/phases"
	"github.com/a-clap/distillation-gui/backend/pt"
	"github.com/a-clap/iot/pkg/distillation"
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
	gpio.AddListener(e)
	phases.AddListener(e)
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

// OnGPIOChange implements gpio.Listener
func (e *eventEmitter) OnGPIOChange(config parameters.GPIO) {
	runtime.EventsEmit(e.ctx, NotifyGPIO, config)
}

func (e *eventEmitter) OnError(errID int) {
	runtime.EventsEmit(e.ctx, NotifyError, errID)
}

// OnConfigChange implements phases.Listener
func (e *eventEmitter) OnConfigChange(c distillation.ProcessConfig) {
	logger.Debug("OnConfigChange")
	runtime.EventsEmit(e.ctx, NotifyPhasesConfig, c)
}

// OnConfigValidate implements phases.Listener
func (e *eventEmitter) OnConfigValidate(validation distillation.ProcessConfigValidation) {
	logger.Debug("OnConfigValidate")
	runtime.EventsEmit(e.ctx, NotifyPhasesValidate, validation)
}

// OnPhaseConfigChange implements phases.Listener
func (e *eventEmitter) OnPhaseConfigChange(phaseNumber int, cfg distillation.ProcessPhaseConfig) {
	logger.Debug("OnPhaseConfigChange")
	runtime.EventsEmit(e.ctx, NotifyPhasesPhaseConfig, phaseNumber, cfg)
}

// OnPhasesCountChange implements phases.Listener
func (e *eventEmitter) OnPhasesCountChange(count distillation.ProcessPhaseCount) {
	logger.Debug("OnPhasesCountChange")
	runtime.EventsEmit(e.ctx, NotifyPhasesPhaseCount, count)
}

// OnStatusChange implements phases.Listener
func (e *eventEmitter) OnStatusChange(status distillation.ProcessStatus) {
	logger.Debug("OnStatusChange")
	runtime.EventsEmit(e.ctx, NotifyPhasesStatus, status)
}
