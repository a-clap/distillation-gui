package loadSaver

import (
	"errors"
	"log"

	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/distillation-gui/backend/gpio"
	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/distillation-gui/backend/phases"
	"github.com/a-clap/distillation-gui/backend/pt"
	"github.com/a-clap/distillation/pkg/distillation"
	"github.com/a-clap/distillation/pkg/process"
	"gopkg.in/yaml.v3"
)

type LoadSaver interface {
	Save(key string, data []byte) error
	Load(key string) (data []byte, err error)
}

type params struct {
	heaterSettings  map[string]*parameters.Heater
	dsSettings      map[string]*parameters.DS
	ptSettings      map[string]*parameters.PT
	gpioSettings    map[string]*parameters.GPIO
	processSettings process.Config
}

type handlerType struct {
	LoadSaver LoadSaver
	params    params
}

const (
	paramsKey = "params"
)

var (
	handler = &handlerType{
		LoadSaver: nil,
		params: params{
			heaterSettings: make(map[string]*parameters.Heater, 0),
			dsSettings:     make(map[string]*parameters.DS, 0),
			ptSettings:     make(map[string]*parameters.PT, 0),
			gpioSettings:   make(map[string]*parameters.GPIO, 0),
			processSettings: process.Config{
				PhaseNumber: 0,
				Phases:      make([]process.PhaseConfig, 0),
			},
		}}
)

func Init(saver LoadSaver) {
	handler.LoadSaver = saver
}

func Run() error {
	if handler.LoadSaver == nil {
		return errors.New("no load saver interface")
	}

	// Add handler to listeners, so he will always have the newest values
	heater.AddListener(handler)
	ds.AddListener(handler)
	gpio.AddListener(handler)
	pt.AddListener(handler)
	phases.AddListener(handler)
	return nil
}

func Load() []error {
	data, err := handler.LoadSaver.Load(paramsKey)
	if err != nil {
		return []error{err}
	}
	params := parameters.GUI{}

	if err := yaml.Unmarshal(data, &params); err != nil {
		return []error{err}
	}
	var errs []error
	if err := ds.Apply(params.DS); err != nil {
		errs = append(errs, err...)
	}
	if err := pt.Apply(params.PT); err != nil {
		errs = append(errs, err...)
	}
	if err := gpio.Apply(params.GPIO); err != nil {
		errs = append(errs, err...)
	}
	if err := heater.Apply(params.Heaters); err != nil {
		errs = append(errs, err...)
	}
	if err := phases.Apply(params.Process); err != nil {
		errs = append(errs, err...)
	}

	return errs
}

func Save() error {
	params := parameters.GUI{}
	for _, elem := range handler.params.heaterSettings {
		params.Heaters = append(params.Heaters, *elem)
	}
	for _, elem := range handler.params.gpioSettings {
		params.GPIO = append(params.GPIO, *elem)
	}
	for _, elem := range handler.params.dsSettings {
		params.DS = append(params.DS, *elem)
	}
	for _, elem := range handler.params.ptSettings {
		params.PT = append(params.PT, *elem)
	}
	params.Process = handler.params.processSettings

	y, err := yaml.Marshal(params)
	if err != nil {
		log.Println(err)
		return err
	}
	return handler.LoadSaver.Save(paramsKey, y)
}

func (h *handlerType) OnHeaterChange(config parameters.Heater) {
	// If there is no such heater, create it
	if _, ok := h.params.heaterSettings[config.ID]; !ok {
		h.params.heaterSettings[config.ID] = &parameters.Heater{
			ID:      config.ID,
			Enabled: false,
		}
	}
	h.params.heaterSettings[config.ID].Enabled = config.Enabled
}

func (h *handlerType) OnDSConfigChange(config parameters.DS) {
	// If there is no such DS, create it
	if _, ok := h.params.dsSettings[config.ID]; !ok {
		h.params.dsSettings[config.ID] = &parameters.DS{}
	}
	h.params.dsSettings[config.ID] = &config
}

func (h *handlerType) OnDSTemperatureChange(parameters.Temperature) {
	// don't care
}

func (h *handlerType) OnGPIOChange(config parameters.GPIO) {
	// If there is no such gpio, create it
	if _, ok := h.params.gpioSettings[config.ID]; !ok {
		h.params.gpioSettings[config.ID] = &parameters.GPIO{}
	}
	h.params.gpioSettings[config.ID] = &config
}

func (h *handlerType) OnPTConfigChange(config parameters.PT) {
	// If there is no such DS, create it
	if _, ok := h.params.ptSettings[config.ID]; !ok {
		h.params.ptSettings[config.ID] = &parameters.PT{}
	}
	h.params.ptSettings[config.ID] = &config

}

func (h *handlerType) OnPTTemperatureChange(parameters.Temperature) {
}

func (h *handlerType) OnPhasesCountChange(c distillation.ProcessPhaseCount) {
	h.params.processSettings.PhaseNumber = c.PhaseNumber
}

func (h *handlerType) OnPhaseConfigChange(phaseNumber int, cfg distillation.ProcessPhaseConfig) {
	h.growPhases(phaseNumber, cfg)
	h.params.processSettings.Phases[phaseNumber] = cfg.PhaseConfig
}

func (h *handlerType) growPhases(min int, cfg distillation.ProcessPhaseConfig) {
	currentPhases := len(h.params.processSettings.Phases)
	for i := 0; i <= (min - currentPhases); i++ {
		h.params.processSettings.Phases = append(h.params.processSettings.Phases, cfg.PhaseConfig)
	}
}

func (h *handlerType) OnGlobalConfig(config process.Config) {
	h.params.processSettings = config
}

func (h *handlerType) OnConfigValidate(distillation.ProcessConfigValidation) {
}

func (h *handlerType) OnStatusChange(distillation.ProcessStatus) {
}

func (h *handlerType) OnConfigChange(distillation.ProcessConfig) {
}
