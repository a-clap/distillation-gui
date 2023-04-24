/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package heater

import (
	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/iot/pkg/distillation"
)

type Client interface {
	GetAll() ([]distillation.HeaterConfigGlobal, error)
	Enable(setConfig distillation.HeaterConfigGlobal) (distillation.HeaterConfigGlobal, error)
}

type Listener interface {
	OnHeaterChange(heater parameters.Heater)
}

type heaterHandler struct {
	client    Client
	listeners []Listener
	heaters   map[string]*parameters.Heater
}

var (
	handler = &heaterHandler{
		client:    nil,
		listeners: make([]Listener, 0),
		heaters:   make(map[string]*parameters.Heater),
	}
)

// Init prepare package to handle various requests
func Init(c Client) error {
	handler.client = c
	return initHandler()
}

func initHandler() error {
	globals, err := handler.client.GetAll()
	if err != nil {
		return err
	}

	for _, global := range globals {
		heater := &parameters.Heater{
			ID:      global.ID,
			Enabled: global.Enabled,
		}
		handler.heaters[global.ID] = heater
	}

	return nil
}
func Apply(config []parameters.Heater) []error {
	var errs []error
	for _, c := range config {
		// SetCorrection
		if err := Enable(c.ID, c.Enabled); err != nil {
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
		notify(conf)
	}
}

func Get() []parameters.Heater {
	heaters := make([]parameters.Heater, 0, len(handler.heaters))
	for _, v := range handler.heaters {
		heaters = append(heaters, *v)
	}

	return heaters
}

func Enable(id string, enable bool) error {
	cfg, ok := handler.heaters[id]
	if !ok {
		err := &Error{ID: id, Op: "EnableGlobal", Err: ErrIDNotFound.Error()}
		return err
	}

	// Nothing to do
	if cfg.Enabled == enable {
		return nil
	}

	setCfg := distillation.HeaterConfigGlobal{
		ID:      cfg.ID,
		Enabled: enable,
	}

	newCfg, err := handler.client.Enable(setCfg)
	if err != nil {
		err = &Error{ID: id, Op: "EnableGlobal.Enable", Err: err.Error()}
	} else {
		cfg.Enabled = newCfg.Enabled
	}
	notify(*cfg)
	return err
}
