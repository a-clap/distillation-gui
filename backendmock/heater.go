package backendmock

import (
	"errors"

	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/iot/pkg/distillation"
	"github.com/a-clap/logging"
	"golang.org/x/exp/slices"
)

var (
	_ heater.Client = (*HeaterClient)(nil)
)

type HeaterClient struct {
	Globals []distillation.HeaterConfigGlobal
	Enabled []distillation.HeaterConfig
}

// Configure implements heater.Client
func (h *HeaterClient) Configure(setConfig distillation.HeaterConfig) (distillation.HeaterConfig, error) {
	logger.Debug("configure")
	idx := slices.IndexFunc(h.Enabled, func(c distillation.HeaterConfig) bool {
		return setConfig.ID == c.ID
	})
	if idx == -1 {
		logger.Debug("no such heater", logging.String("ID", setConfig.ID))
		return distillation.HeaterConfig{}, errors.New("no such heater")
	}
	h.Enabled[idx] = setConfig
	return h.Enabled[idx], nil
}

// Enable implements heater.Client
func (h *HeaterClient) Enable(setConfig distillation.HeaterConfigGlobal) (distillation.HeaterConfigGlobal, error) {
	logger.Debug("enable", logging.String("ID", setConfig.ID))
	idx := slices.IndexFunc(h.Globals, func(c distillation.HeaterConfigGlobal) bool {
		return setConfig.ID == c.ID
	})
	if idx == -1 {
		logger.Debug("no such heater", logging.String("ID", setConfig.ID))
		return distillation.HeaterConfigGlobal{}, errors.New("no such heater")
	}
	h.Globals[idx] = setConfig
	return h.Globals[idx], nil
}

// GetAll implements heater.Client
func (h *HeaterClient) GetAll() ([]distillation.HeaterConfigGlobal, error) {
	logger.Debug("GetAll")
	return h.Globals, nil
}

// GetEnabled implements heater.Client
func (h *HeaterClient) GetEnabled() ([]distillation.HeaterConfig, error) {
	logger.Debug("GetEnabled")
	return h.Enabled, nil
}
