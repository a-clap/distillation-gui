package backend

import (
	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/distillation-gui/backend/heater"
)

type Option func(b *Backend) error

func WithHeaterClient(c heater.Client) Option {
	return func(b *Backend) error {
		return heater.Init(c)
	}
}
func WithDSClient(c ds.Client) Option {
	return func(b *Backend) error {
		return ds.Init(c, b.dsChan, b.interval)
	}
}
