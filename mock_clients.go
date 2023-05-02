package main

import (
	"fmt"
	"log"
	"time"

	"github.com/a-clap/distillation-gui/backend"
	"github.com/a-clap/distillation-gui/backendmock"
	"github.com/a-clap/distillation/pkg/distillation"
	"github.com/a-clap/distillation/pkg/process"
	"github.com/a-clap/distillation/pkg/wifi"
	"github.com/a-clap/embedded/pkg/ds18b20"
	"github.com/a-clap/embedded/pkg/embedded"
	"github.com/a-clap/embedded/pkg/gpio"
	"github.com/a-clap/embedded/pkg/max31865"
)

func mockClients() []backend.Option {
	// HeaterClient - Mock
	heaterClient := backendmock.HeaterClient{}
	heaterClient.Globals = append(heaterClient.Globals,
		distillation.HeaterConfigGlobal{ID: "heater_1", Enabled: false},
		distillation.HeaterConfigGlobal{ID: "heater_2", Enabled: true},
		distillation.HeaterConfigGlobal{ID: "heater_3", Enabled: false},
		distillation.HeaterConfigGlobal{ID: "heater_4", Enabled: false},
	)

	dsClient := backendmock.DSClient{}
	dsClient.DS = append(dsClient.DS,
		distillation.DSConfig{DSSensorConfig: embedded.DSSensorConfig{SensorConfig: ds18b20.SensorConfig{
			Name:         "DS_1",
			ID:           "1",
			Resolution:   ds18b20.Resolution9Bit,
			Correction:   0,
			PollInterval: 500,
			Samples:      1,
		}}},
		distillation.DSConfig{DSSensorConfig: embedded.DSSensorConfig{SensorConfig: ds18b20.SensorConfig{
			Name:         "DS_2",
			ID:           "2",
			Resolution:   ds18b20.Resolution10Bit,
			Correction:   0,
			PollInterval: 1,
			Samples:      2,
		}}},
		distillation.DSConfig{DSSensorConfig: embedded.DSSensorConfig{SensorConfig: ds18b20.SensorConfig{
			Name:         "DS_3",
			ID:           "3",
			Resolution:   ds18b20.Resolution11Bit,
			Correction:   0,
			PollInterval: 1,
			Samples:      3,
		}}},
		distillation.DSConfig{DSSensorConfig: embedded.DSSensorConfig{SensorConfig: ds18b20.SensorConfig{
			Name:         "DS_4",
			ID:           "4",
			Resolution:   ds18b20.Resolution12Bit,
			Correction:   0,
			PollInterval: 1,
			Samples:      4,
		}}},
	)

	ptClient := backendmock.PTClient{}
	ptClient.PT = append(ptClient.PT,
		distillation.PTConfig{PTSensorConfig: embedded.PTSensorConfig{
			Enabled: false,
			SensorConfig: max31865.SensorConfig{
				Name:         "PT_1",
				ID:           "id_1",
				Correction:   0,
				ASyncPoll:    false,
				PollInterval: 1 * time.Second,
				Samples:      3,
			}}},
		distillation.PTConfig{PTSensorConfig: embedded.PTSensorConfig{
			Enabled: false,
			SensorConfig: max31865.SensorConfig{
				Name:         "PT_2",
				ID:           "id_2",
				Correction:   10.0,
				ASyncPoll:    true,
				PollInterval: 1 * time.Second,
				Samples:      7,
			}}},
		distillation.PTConfig{PTSensorConfig: embedded.PTSensorConfig{
			Enabled: false,
			SensorConfig: max31865.SensorConfig{
				Name:         "PT_3",
				ID:           "id_3",
				Correction:   12.0,
				ASyncPoll:    true,
				PollInterval: 1 * time.Second,
				Samples:      13,
			}}},
	)

	gpioClient := backendmock.GPIOClient{}
	gpioClient.GPIO = append(gpioClient.GPIO,
		distillation.GPIOConfig{GPIOConfig: embedded.GPIOConfig{Config: gpio.Config{
			ID:          "gpio_1",
			Direction:   gpio.DirOutput,
			ActiveLevel: gpio.Low,
			Value:       false,
		}}},
		distillation.GPIOConfig{GPIOConfig: embedded.GPIOConfig{Config: gpio.Config{
			ID:          "gpio_2",
			Direction:   gpio.DirOutput,
			ActiveLevel: gpio.High,
			Value:       true,
		}}},
		distillation.GPIOConfig{GPIOConfig: embedded.GPIOConfig{Config: gpio.Config{
			ID:          "gpio_3",
			Direction:   gpio.DirOutput,
			ActiveLevel: gpio.Low,
			Value:       true,
		}}},
	)

	phaseClient := backendmock.PhasesClient{}
	phaseClient.Config = process.Config{PhaseNumber: 3, Phases: make([]process.PhaseConfig, 3)}
	for i := range phaseClient.Config.Phases {
		gpio := make([]process.GPIOConfig, 3)
		for j := range gpio {
			gpio[j] = process.GPIOConfig{
				ID:         fmt.Sprint("gpio", j),
				SensorID:   fmt.Sprint("sensor ", j),
				TLow:       0,
				THigh:      0,
				Hysteresis: 0,
				Inverted:   false,
			}
		}

		phaseClient.Config.Phases[i].GPIO = gpio

		heaters := make([]process.HeaterPhaseConfig, 3)
		for j := range heaters {
			heaters[j] = process.HeaterPhaseConfig{
				ID:    fmt.Sprint("heater", j),
				Power: 0,
			}
		}
		phaseClient.Config.Phases[i].Heaters = heaters
	}

	w, err := wifi.New()
	if err != nil {
		log.Fatalln(err)
	}

	return []backend.Option{
		backend.WithHeaterClient(&heaterClient),
		backend.WithDSClient(&dsClient),
		backend.WithPTClient(&ptClient),
		backend.WithGPIOClient(&gpioClient),
		backend.WithPhaseClient(&phaseClient),
		backend.WithWifi(w),
	}
}
