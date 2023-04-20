package backendmock

import (
	"errors"

	"github.com/a-clap/distillation-gui/backend/phases"
	"github.com/a-clap/iot/pkg/distillation"
	"github.com/a-clap/iot/pkg/distillation/process"
)

var (
	_ phases.Client = (*PhasesClient)(nil)
)

type PhasesClient struct {
	Config  process.Config
	Process distillation.ProcessConfig
	Stats   distillation.ProcessStatus
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
	p.Config.PhaseNumber = count.PhaseNumber
	for i := 0; i < len(p.Config.Phases); i++ {
		p.Config.Phases = append(p.Config.Phases, process.PhaseConfig{})
	}
	return count, nil
}

// ConfigureProcess implements phases.Client
func (p *PhasesClient) ConfigureProcess(cfg distillation.ProcessConfig) (distillation.ProcessConfig, error) {
	p.Process = cfg
	return cfg, nil
}

// GetPhaseConfig implements phases.Client
func (p *PhasesClient) GetPhaseConfig(phaseNumber int) (distillation.ProcessPhaseConfig, error) {
	if p.Config.PhaseNumber < phaseNumber {
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
