package backend

import (
	"github.com/a-clap/distillation-gui/backend/ds"
	"github.com/a-clap/distillation-gui/backend/gpio"
	"github.com/a-clap/distillation-gui/backend/heater"
	"github.com/a-clap/distillation-gui/backend/phases"
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
		b.dsChan = make(chan error, 10)
		if err := ds.Init(c, b.dsChan, b.interval); err != nil {
			return err
		}
		ds.Run()
		return nil
	}
}

func WithPTClient(c pt.Client) Option {
	return func(b *Backend) error {
		b.ptChan = make(chan error, 10)
		if err := pt.Init(c, b.ptChan, b.interval); err != nil {
			return err
		}
		pt.Run()
		return nil
	}
}

func WithGPIOClient(c gpio.Client) Option {
	return func(b *Backend) error {
		return gpio.Init(c)
	}
}

func WithWifi(c wifi.Client) Option {
	return func(b *Backend) error {
		wifi.Init(c)
		return nil
	}
}

func WithPhaseClient(c phases.Client) Option {
	return func(b *Backend) error {
		b.phaseChan = make(chan error, 10)
		phases.Init(c, b.phaseChan, b.interval)
		return nil
	}
}
