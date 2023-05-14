/*
 * Copyright (c) 2023 a-clahandler. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package pt

import (
	"sync/atomic"
	"time"

	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/distillation/pkg/distillation"
	"golang.org/x/exp/slices"
)

type Client interface {
	Get() ([]distillation.PTConfig, error)
	Configure(sensor distillation.PTConfig) (distillation.PTConfig, error)
	Temperatures() ([]distillation.Temperature, error)
}

type Listener interface {
	OnPTConfigChange(parameters.PT)
	OnPTTemperatureChange(temperature parameters.Temperature)
}

type ptHandler struct {
	client    Client
	listeners []Listener
	running   atomic.Bool
	sensors   map[string]*parameters.PT
	interval  time.Duration
	finish    chan struct{}
	err       chan<- error
}

var (
	handler = &ptHandler{
		client:    nil,
		listeners: make([]Listener, 0),
		running:   atomic.Bool{},
		sensors:   make(map[string]*parameters.PT),
		interval:  1 * time.Second,
		err:       nil,
	}
)

// Init prepare package to handle various requests
func Init(c Client, err chan<- error, interval time.Duration) error {
	handler.client = c
	handler.err = err
	handler.interval = interval
	return initHandler()
}

func initHandler() error {
	sensors, err := handler.client.Get()
	if err != nil {
		return err
	}

	for _, s := range sensors {
		handler.sensors[s.ID] = &parameters.PT{
			PTConfig: s,
		}
	}
	return nil
}

func Apply(config []parameters.PT) []error {
	var errs []error
	for _, c := range config {
		// SetCorrection
		if err := SetCorrection(c.ID, c.Correction); err != nil {
			errs = append(errs, err)
		}

		// SetSamples
		if err := SetSamples(c.ID, c.Samples); err != nil {
			errs = append(errs, err)
		}

		// Enable
		if err := Enable(c.ID, c.Enabled); err != nil {
			errs = append(errs, err)
		}

		// Enable
		if err := SetName(c.ID, c.Name); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func Stop() {
	if handler.running.Load() {
		close(handler.finish)
		handler.running.Store(false)
	}
}

func Run() {
	if !handler.running.Load() {
		handler.finish = make(chan struct{})
		handler.running.Store(true)
		update()
	}
}

// Refresh read every possible data from Client and serves them to Listener
func Refresh() {
	configs := Get()
	for _, conf := range configs {
		notifyConfig(conf)
	}
}

func AddListener(listener Listener) {
	handler.listeners = append(handler.listeners, listener)
}

func Get() []parameters.PT {
	sensors := make([]parameters.PT, 0, len(handler.sensors))
	for _, s := range handler.sensors {
		sensors = append(sensors, *s)

	}
	slices.SortFunc(sensors, func(i, j parameters.PT) bool {
		return i.ID < j.ID
	})
	return sensors
}
func SetName(id string, name string) error {
	cfg, ok := handler.sensors[id]
	if !ok {
		return &Error{ID: id, Op: "SetName", Err: ErrIDNotFound.Error()}
	}

	setCfg := cfg.PTConfig
	setCfg.Name = name
	newCfg, err := handler.client.Configure(setCfg)
	if err != nil {
		err = &Error{ID: id, Op: "SetName.Configure", Err: err.Error()}
	} else {
		cfg.Name = newCfg.Name
	}

	notifyConfig(*cfg)
	return err
}
func Enable(id string, enable bool) error {
	cfg, ok := handler.sensors[id]
	if !ok {
		return &Error{ID: id, Op: "Enable", Err: ErrIDNotFound.Error()}
	}

	// Nothing to do
	if cfg.Enabled == enable {
		return nil
	}

	setCfg := cfg.PTConfig
	setCfg.Enabled = enable
	newCfg, err := handler.client.Configure(setCfg)
	if err != nil {
		err = &Error{ID: id, Op: "Enable.Configure", Err: err.Error()}
	} else {
		cfg.Enabled = newCfg.Enabled
	}

	notifyConfig(*cfg)
	return err
}

func SetCorrection(id string, correction float64) error {
	cfg, ok := handler.sensors[id]
	if !ok {
		return &Error{ID: id, Op: "SetCorrection", Err: ErrIDNotFound.Error()}
	}

	if cfg.Correction == correction {
		return nil
	}

	setCfg := cfg.PTConfig
	setCfg.Correction = correction

	newCfg, err := handler.client.Configure(setCfg)
	if err != nil {
		err = &Error{ID: id, Op: "SetCorrection.Configure", Err: err.Error()}
	} else {
		cfg.Correction = newCfg.Correction
	}
	notifyConfig(*cfg)
	return err
}

func SetSamples(id string, samples uint) error {
	cfg, ok := handler.sensors[id]
	if !ok {
		return &Error{ID: id, Op: "SetSamples", Err: ErrIDNotFound.Error()}
	}

	if cfg.Samples == samples {
		return nil
	}

	setCfg := cfg.PTConfig
	setCfg.Samples = samples
	newCfg, err := handler.client.Configure(setCfg)
	if err != nil {
		err = &Error{ID: id, Op: "SetSamples.Configure", Err: err.Error()}
	} else {
		cfg.Samples = newCfg.Samples
	}
	notifyConfig(*cfg)
	return err
}

func update() {
	go func() {
		for handler.running.Load() {
			select {
			case <-handler.finish:
				handler.running.Store(false)
			case <-time.After(handler.interval):
				temps, err := handler.client.Temperatures()
				if err != nil {
					err := &Error{ID: "", Op: "Temperatures", Err: err.Error()}
					notifyError(err)
					continue
				}

				for _, temp := range temps {
					if _, ok := handler.sensors[temp.ID]; !ok {
						err := &Error{ID: temp.ID, Op: "Temperatures", Err: ErrIDNotFound.Error()}
						notifyError(err)
						continue
					}

					notifyTemperature(parameters.Temperature{
						Temperature: distillation.Temperature{
							ID:          temp.ID,
							Temperature: temp.Temperature,
							Stamp:       temp.Stamp,
							Error:       temp.Error,
						},
					})
				}
			}
		}
	}()
}
