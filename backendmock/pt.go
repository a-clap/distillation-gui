package backendmock

import (
	"errors"

	"github.com/a-clap/distillation-gui/backend/pt"
	"github.com/a-clap/iot/pkg/distillation"
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
func (p *PTClient) GetSensors() ([]distillation.PTConfig, error) {
	logger.Debug("GetSensors")
	return p.PT, nil
}

// Temperatures implements ds.Client
func (p *PTClient) Temperatures() ([]distillation.PTTemperature, error) {
	var temps []distillation.PTTemperature
	for _, elem := range p.PT {
		if elem.Enabled {
			temps = append(temps, distillation.PTTemperature{
				ID:          elem.ID,
				Temperature: randomTemperature(70, 75) + elem.Correction,
			})
		}
	}
	return temps, nil
}
