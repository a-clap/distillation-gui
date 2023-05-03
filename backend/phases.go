package backend

import (
	"github.com/a-clap/distillation-gui/backend/phases"
	"github.com/a-clap/distillation/pkg/distillation"
	"github.com/a-clap/distillation/pkg/process"
	"github.com/a-clap/logging"
)

func (b *Backend) PhasesGetPhaseCount() *distillation.ProcessPhaseCount {
	logger.Debug("PhasesGetPhaseCount")

	count, err := phases.GetPhaseCount()
	if err != nil {
		logger.Error("error on PhasesGetPhaseCount", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrPhaseGetCount)
		return nil
	}
	return &count

}
func (b *Backend) PhasesGetPhaseConfigs() []distillation.ProcessPhaseConfig {
	logger.Debug("PhasesGetPhaseConfigs")

	confs, err := phases.GetPhaseConfigs()
	if err != nil {
		logger.Error("error on PhasesGetPhaseConfigs", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrPhaseGetPhaseConfigs)
		return nil
	}
	return confs
}
func (b *Backend) PhasesGetGlobalConfig() *process.Config {
	logger.Debug("PhasesGetGlobalConfig")

	conf, err := phases.GetGlobalConfig()
	if err != nil {
		logger.Error("error on PhasesGetGlobalConfig", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrPhaseGetGlobalConfig)
		return nil
	}
	return &conf
}

func (b *Backend) PhasesSetPhaseCount(count int) {
	logger.Debug("PhasesSetPhaseCount")

	if err := phases.SetPhaseCount(uint(count)); err != nil {
		logger.Error("error on SetPhaseCount", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrPhasesSetPhaseCount)
	}
}
func (b *Backend) PhasesSetConfig(number int, cfg distillation.ProcessPhaseConfig) {
	logger.Debug("PhasesSetConfig")
	if err := phases.SetConfig(number, cfg); err != nil {
		logger.Error("error on SetConfig", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrPhasesSetConfig)
	}
}
func (b *Backend) PhasesValidateConfig() {
	logger.Debug("PhasesValidateConfig")
	if err := phases.ValidateConfig(); err != nil {
		logger.Error("error on ValidateConfig", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrPhasesValidateConfig)
	}
}
func (b *Backend) PhasesEnable() {
	logger.Debug("PhasesEnable")
	if err := phases.Enable(); err != nil {
		logger.Error("error on Enable", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrPhasesEnable)
	}
}
func (b *Backend) PhasesDisable() {
	logger.Debug("PhasesDisable")
	if err := phases.Disable(); err != nil {
		logger.Error("error on Disable", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrPhasesDisable)
	}
}
func (b *Backend) PhasesMoveToNext() {
	logger.Debug("PhasesMoveToNext")
	if err := phases.MoveToNext(); err != nil {
		logger.Error("error on MoveToNext", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrPhasesMoveToNext)
	}
}
