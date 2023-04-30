package backendmock

import (
	"errors"

	"github.com/a-clap/distillation-gui/backend/gpio"
	"github.com/a-clap/distillation/pkg/distillation"
	"github.com/a-clap/logging"
	"golang.org/x/exp/slices"
)

var (
	_ gpio.Client = (*GPIOClient)(nil)
)

type GPIOClient struct {
	GPIO []distillation.GPIOConfig
}

// Configure implements gpio.Client
func (g *GPIOClient) Configure(setConfig distillation.GPIOConfig) (distillation.GPIOConfig, error) {
	logger.Debug("configure")
	idx := slices.IndexFunc(g.GPIO, func(c distillation.GPIOConfig) bool {
		return setConfig.ID == c.ID
	})

	if idx == -1 {
		logger.Debug("no such GPIO", logging.String("ID", setConfig.ID))
		return distillation.GPIOConfig{}, errors.New("no such GPIO")
	}

	g.GPIO[idx] = setConfig
	return setConfig, nil

}

// Get implements gpio.Client
func (g *GPIOClient) Get() ([]distillation.GPIOConfig, error) {
	return g.GPIO, nil
}
