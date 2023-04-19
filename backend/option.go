package backend

import (
	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/distillation-gui/backend/pt"
	"github.com/a-clap/distillation-gui/backend/wifi"
)

type Option func(b *Backend) error

func WithHeaterClient(c heater.Client) Option {
	return func(b *Backend) error {
		return heater.Init(c)
	}
}

func WithDSClient(c ds.Client) Option {
	return func(b *Backend) error {
		if err := ds.Init(c, b.dsChan, b.interval); err != nil {
			return err
		}
		ds.Run()
		return nil
	}
}

func WithPTClient(c pt.Client) Option {
	return func(b *Backend) error {
		if err := pt.Init(c, b.ptChan, b.interval); err != nil {
			return err
		}
		pt.Run()
		return nil
	}
}

func WithWifi(c wifi.Client) Option {
	return func(b *Backend) error {
		wifi.Init(c)
		return nil
	}
}
