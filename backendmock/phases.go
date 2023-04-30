package backendmock

import (
	"errors"
	"fmt"

	"github.com/a-clap/distillation-gui/backend/phases"
	"github.com/a-clap/distillation/pkg/distillation"
	"github.com/a-clap/distillation/pkg/distillation/process"
)

var (
	_ phases.Client = (*PhasesClient)(nil)
)

type PhasesClient struct {
	Config  process.Config
	Process distillation.ProcessConfig
	Stats   distillation.ProcessStatus
}

func (p *PhasesClient) Init(count int) {
	p.Config.PhaseNumber = count
	p.Config.Phases = make([]process.PhaseConfig, count)
	for i := range p.Config.Phases {
		p.Config.Phases[i].Next = process.MoveToNextConfig{
			Type:                   process.ByTime,
			SensorID:               "BLAH",
			SensorThreshold:        1.23,
			TemperatureHoldSeconds: 13,
			SecondsToMove:          10,
		}
		p.Config.Phases[i].GPIO = make([]process.GPIOPhaseConfig, 3)
		for j := range p.Config.Phases[i].GPIO {
			p.Config.Phases[i].GPIO[j] = process.GPIOPhaseConfig{
				ID:         fmt.Sprintf("gpio %v", j),
				SensorID:   "sensor",
				TLow:       1.23,
				THigh:      4.56,
				Hysteresis: 3.17,
				Inverted:   false,
			}
		}

		p.Config.Phases[i].Heaters = make([]process.HeaterPhaseConfig, 3)
		for j := range p.Config.Phases[i].Heaters {
			p.Config.Phases[i].Heaters[j] = process.HeaterPhaseConfig{
				ID:    fmt.Sprintf("heater %v", j),
				Power: j,
			}
		}
	}

}

// ConfigurePhase implements phases.Client
func (p *PhasesClient) ConfigurePhase(phaseNumber int, setConfig distillation.ProcessPhaseConfig) (distillation.ProcessPhaseConfig, error) {
	if p.Config.PhaseNumber < phaseNumber {
		return distillation.ProcessPhaseConfig{}, errors.New("no such phase")
	}
	p.Config.Phases[phaseNumber] = setConfig.PhaseConfig
	return setConfig, nil
}

// ConfigurePhaseCount implements phases.Client
func (p *PhasesClient) ConfigurePhaseCount(count distillation.ProcessPhaseCount) (distillation.ProcessPhaseCount, error) {
	p.Init(count.PhaseNumber)
	return count, nil
}

// ConfigureProcess implements phases.Client
func (p *PhasesClient) ConfigureProcess(cfg distillation.ProcessConfig) (distillation.ProcessConfig, error) {
	p.Process = cfg
	return cfg, nil
}

// GetPhaseConfig implements phases.Client
func (p *PhasesClient) GetPhaseConfig(phaseNumber int) (distillation.ProcessPhaseConfig, error) {
	if p.Config.PhaseNumber <= phaseNumber {
		return distillation.ProcessPhaseConfig{}, errors.New("overflow")
	}
	c := distillation.ProcessPhaseConfig{PhaseConfig: p.Config.Phases[phaseNumber]}
	return c, nil

}

// GetPhaseCount implements phases.Client
func (p *PhasesClient) GetPhaseCount() (distillation.ProcessPhaseCount, error) {
	return distillation.ProcessPhaseCount{PhaseNumber: p.Config.PhaseNumber}, nil
}

// Status implements phases.Client
func (p *PhasesClient) Status() (distillation.ProcessStatus, error) {
	return p.Stats, nil
}

// ValidateConfig implements phases.Client
func (p *PhasesClient) ValidateConfig() (distillation.ProcessConfigValidation, error) {
	return distillation.ProcessConfigValidation{Valid: true}, nil
}
