/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package gpio

import (
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/iot/pkg/distillation"
	"github.com/a-clap/iot/pkg/embedded/gpio"
	"golang.org/x/exp/slices"
)

type Client interface {
	Get() ([]distillation.GPIOConfig, error)
	Configure(setConfig distillation.GPIOConfig) (distillation.GPIOConfig, error)
}
type Listener interface {
	OnGPIOChange(config parameters.GPIO)
}

type gpioHandler struct {
	client    Client
	listeners []Listener
	gpio      map[string]*parameters.GPIO
}

var (
	handler = &gpioHandler{
		client:    nil,
		listeners: make([]Listener, 0),
		gpio:      make(map[string]*parameters.GPIO),
	}
)

// Init prepare package to handle various requests
func Init(c Client) error {
	handler.client = c
	return initHandler()
}

func initHandler() error {
	ios, err := handler.client.Get()
	if err != nil {
		return err
	}

	for _, io := range ios {
		handler.gpio[io.ID] = &parameters.GPIO{GPIOConfig: io}
	}

	return nil
}

func Apply(config []parameters.GPIO) []error {
	var errs []error
	for _, c := range config {
		// SetCorrection
		if err := SetActiveLevel(c.ID, c.ActiveLevel); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func AddListener(listener Listener) {
	handler.listeners = append(handler.listeners, listener)
}

// Refresh read every possible data from Client and serves them to Listener
func Refresh() {
	configs := Get()
	for _, conf := range configs {
		notifyConfig(conf)
	}
}

func Get() []parameters.GPIO {
	sensors := make([]parameters.GPIO, 0, len(handler.gpio))
	for _, s := range handler.gpio {
		sensors = append(sensors, *s)

	}
	slices.SortFunc(sensors, func(i, j parameters.GPIO) bool {
		return i.ID < j.ID
	})
	return sensors
}

func SetActiveLevel(id string, level gpio.ActiveLevel) error {
	cfg, ok := handler.gpio[id]
	if !ok {
		err := &Error{ID: id, Op: "SetActiveLevel", Err: ErrIDNotFound.Error()}
		return err
	}

	if cfg.ActiveLevel == level {
		return nil
	}

	setConfig := cfg.GPIOConfig
	setConfig.ActiveLevel = level
	newCfg, err := handler.client.Configure(setConfig)
	if err != nil {
		err = &Error{ID: id, Op: "SetActiveLevel.Configure", Err: err.Error()}
	} else {
		cfg.GPIOConfig = newCfg
	}
	notifyConfig(*cfg)
	return err
}

func SetState(id string, value bool) error {
	cfg, ok := handler.gpio[id]
	if !ok {
		err := &Error{ID: id, Op: "SetState", Err: ErrIDNotFound.Error()}
		return err
	}

	if cfg.Value == value {
		return nil
	}

	setConfig := cfg.GPIOConfig
	setConfig.Value = value
	newCfg, err := handler.client.Configure(setConfig)
	if err != nil {
		err = &Error{ID: id, Op: "SetState.Configure", Err: err.Error()}
	} else {
		cfg.GPIOConfig = newCfg
	}
	notifyConfig(*cfg)
	return err
}
