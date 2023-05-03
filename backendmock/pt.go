package backendmock

import (
	"errors"
	"time"

	"github.com/a-clap/distillation-gui/backend/pt"
	"github.com/a-clap/distillation/pkg/distillation"
	"golang.org/x/exp/slices"
)

var (
	_ pt.Client = (*PTClient)(nil)
)

type PTClient struct {
	PT []distillation.PTConfig
}

// Configure implements ds.Client
func (p *PTClient) Configure(sensor distillation.PTConfig) (distillation.PTConfig, error) {
	logger.Debug("Configure")
	idx := slices.IndexFunc(p.PT, func(ds distillation.PTConfig) bool {
		return ds.ID == sensor.ID
	})

	if idx == -1 {
		return distillation.PTConfig{}, errors.New("no such pt")
	}
	p.PT[idx] = sensor
	return sensor, nil
}

// GetSensors implements ds.Client
func (p *PTClient) Get() ([]distillation.PTConfig, error) {
	logger.Debug("GetSensors")
	return p.PT, nil
}

// Temperatures implements ds.Client
func (p *PTClient) Temperatures() ([]distillation.Temperature, error) {
	var temps []distillation.Temperature
	for _, elem := range p.PT {
		if elem.Enabled {
			temps = append(temps, distillation.Temperature{
				ID:          elem.ID,
				Temperature: randomTemperature(70, 75) + elem.Correction,
				Stamp:       time.Now().Unix(),
				Error:       0,
			})
		}
	}
	return temps, nil
}
