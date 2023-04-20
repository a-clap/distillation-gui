package backend

import (
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/iot/pkg/distillation"
	"github.com/a-clap/iot/pkg/distillation/process"
)

// Models allows us to create models.ts in frontend with needed structures
type Models struct {
}

func (*Models) Temperature() parameters.Temperature {
	return parameters.Temperature{}
}

func (*Models) HeaterPhaseConfig() process.HeaterPhaseConfig {
	return process.HeaterPhaseConfig{}
}

func (*Models) GPIOPhaseConfig() process.GPIOPhaseConfig {
	return process.GPIOPhaseConfig{}
}

func (*Models) ProcessConfigValidation() distillation.ProcessConfigValidation {
	return distillation.ProcessConfigValidation{}
}

func (*Models) ProcessStatus() distillation.ProcessStatus {
	return distillation.ProcessStatus{}
}

func (*Models) HeaterPhaseStatus() process.HeaterPhaseStatus {
	return process.HeaterPhaseStatus{}
}

func (*Models) TemperaturePhaseStatus() process.TemperaturePhaseStatus {
	return process.TemperaturePhaseStatus{}
}

func (*Models) GPIOPhaseStatus() process.GPIOPhaseStatus {
	return process.GPIOPhaseStatus{}
}
