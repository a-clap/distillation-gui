package backendmock

import (
	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/iot/pkg/distillation"
)

var (
	_ heater.Client = (*HeaterClient)(nil)
)

type HeaterClient struct {
	Heaters []distillation.HeaterConfigGlobal
}

// Configure implements heater.Client
func (*HeaterClient) Configure(setConfig distillation.HeaterConfig) (distillation.HeaterConfig, error) {
	panic("unimplemented")
}

// Enable implements heater.Client
func (*HeaterClient) Enable(setConfig distillation.HeaterConfigGlobal) (distillation.HeaterConfigGlobal, error) {
	panic("unimplemented")
}

// GetAll implements heater.Client
func (*HeaterClient) GetAll() ([]distillation.HeaterConfigGlobal, error) {
	return
}

// GetEnabled implements heater.Client
func (*HeaterClient) GetEnabled() ([]distillation.HeaterConfig, error) {
	panic("unimplemented")
}
