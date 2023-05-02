package backend

import (
	"github.com/a-clap/distillation-gui/backend/phases"
	"github.com/a-clap/distillation/pkg/distillation"
	"github.com/a-clap/distillation/pkg/process"
)

func (b *Backend) PhasesGetPhaseCount() (distillation.ProcessPhaseCount, error) {
	logger.Debug("PhasesGetPhaseCount")
	return phases.GetPhaseCount()
}
func (b *Backend) PhasesGetPhaseConfigs() ([]distillation.ProcessPhaseConfig, error) {
	logger.Debug("PhasesGetPhaseConfigs")
	return phases.GetPhaseConfigs()
}
func (b *Backend) PhasesGetGlobalConfig() (process.Config, error) {
	logger.Debug("PhasesGetGlobalConfig")
	return phases.GetGlobalConfig()
}

func (b *Backend) PhasesSetPhaseCount(count int) error {
	logger.Debug("PhasesSetPhaseCount")
	return phases.SetPhaseCount(uint(count))
}
func (b *Backend) PhasesSetConfig(number int, cfg distillation.ProcessPhaseConfig) error {
	logger.Debug("PhasesSetConfig")
	return phases.SetConfig(number, cfg)
}
func (b *Backend) PhasesValidateConfig() error {
	logger.Debug("PhasesValidateConfig")
	return phases.ValidateConfig()
}
func (b *Backend) PhasesEnable() error {
	logger.Debug("PhasesEnable")
	return phases.Enable()
}
func (b *Backend) PhasesDisable() error {
	logger.Debug("PhasesDisable")
	return phases.Disable()
}
func (b *Backend) PhasesMoveToNext() error {
	logger.Debug("PhasesMoveToNext")
	return phases.MoveToNext()
}
