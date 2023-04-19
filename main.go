package main

import (
	"embed"
	"log"

	"github.com/a-clap/distillation-gui/backend"
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/distillation-gui/backendmock"
	"github.com/a-clap/iot/pkg/distillation"
	"github.com/a-clap/iot/pkg/ds18b20"
	"github.com/a-clap/iot/pkg/embedded"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
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
		distillation.DSConfig{DSSensorConfig: embedded.DSSensorConfig{
			SensorConfig: ds18b20.SensorConfig{
				Name:         "DS_1",
				ID:           "1",
				Resolution:   ds18b20.Resolution9Bit,
				Correction:   0,
				PollInterval: 500,
				Samples:      1,
			},
		}},
		distillation.DSConfig{DSSensorConfig: embedded.DSSensorConfig{
			SensorConfig: ds18b20.SensorConfig{
				Name:         "DS_2",
				ID:           "2",
				Resolution:   ds18b20.Resolution10Bit,
				Correction:   0,
				PollInterval: 1,
				Samples:      2,
			},
		}},
		distillation.DSConfig{DSSensorConfig: embedded.DSSensorConfig{
			SensorConfig: ds18b20.SensorConfig{
				Name:         "DS_3",
				ID:           "3",
				Resolution:   ds18b20.Resolution11Bit,
				Correction:   0,
				PollInterval: 1,
				Samples:      3,
			},
		}},
		distillation.DSConfig{DSSensorConfig: embedded.DSSensorConfig{
			SensorConfig: ds18b20.SensorConfig{
				Name:         "DS_4",
				ID:           "4",
				Resolution:   ds18b20.Resolution12Bit,
				Correction:   0,
				PollInterval: 1,
				Samples:      4,
			},
		}},
	)

	// Create backend
	b, err := backend.New(
		backend.WithHeaterClient(&heaterClient),
		backend.WithDSClient(&dsClient),
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "gui",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        b.Startup,
		Bind: []interface{}{
			b,
			&backend.Events{},
			&backend.Models{},
			&parameters.GUI{},
		},
	})

	if err != nil {
		log.Fatalln("Error:", err.Error())
	}
}
