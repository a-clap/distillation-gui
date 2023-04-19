package backendmock

import (
	"errors"

	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/iot/pkg/distillation"
	"golang.org/x/exp/slices"
)

var (
	_ ds.Client = (*DSClient)(nil)
)

type DSClient struct {
	DS []distillation.DSConfig
}

// Configure implements ds.Client
func (d *DSClient) Configure(sensor distillation.DSConfig) (distillation.DSConfig, error) {
	idx := slices.IndexFunc(d.DS, func(ds distillation.DSConfig) bool {
		return ds.ID == sensor.ID
	})

	if idx == -1 {
		return distillation.DSConfig{}, errors.New("no such ds")
	}
	d.DS[idx] = sensor
	return sensor, nil
}

// GetSensors implements ds.Client
func (d *DSClient) GetSensors() ([]distillation.DSConfig, error) {
	return d.DS, nil
}

// Temperatures implements ds.Client
func (d *DSClient) Temperatures() ([]distillation.DSTemperature, error) {
	logger.Debug("Temperatures")
	var temps []distillation.DSTemperature
	for _, elem := range d.DS {
		if elem.Enabled {
			temps = append(temps, distillation.DSTemperature{
				ID:          elem.ID,
				Temperature: randomTemperature(0, 75),
			})
		}
	}
	return temps, nil
}
