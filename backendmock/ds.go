package backendmock

import (
	"errors"
	"time"

	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/distillation/pkg/distillation"
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
func (d *DSClient) Get() ([]distillation.DSConfig, error) {
	return d.DS, nil
}

// Temperatures implements ds.Client
func (d *DSClient) Temperatures() ([]distillation.Temperature, error) {
	var temps []distillation.Temperature
	for _, elem := range d.DS {
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
