package backend

import (
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/distillation/pkg/distillation"
	"github.com/a-clap/distillation/pkg/process"
	"github.com/a-clap/embedded/pkg/gpio"
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

func (*Models) GPIOPhaseConfig() process.GPIOConfig {
	return process.GPIOConfig{}
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

func (*Models) MoveToNextConfig() process.MoveToNextConfig {
	return process.MoveToNextConfig{}
}

func (*Models) GPIOActiveLevel() gpio.ActiveLevel {
	return gpio.High
}

func (*Models) TemperatureErrorCodeWrongID() int {
	return int(distillation.ErrorCodeWrongID)
}

func (*Models) TemperatureErrorCodeEmptyBuffer() int {
	return int(distillation.ErrorCodeEmptyBuffer)
}

func (*Models) TemperatureErrorCodeInternal() int {
	return int(distillation.ErrorCodeInternal)
}
