/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package ds

import (
	"sync/atomic"
	"time"

	"github.com/a-clap/distillation-gui/backend/parameters"
	"github.com/a-clap/distillation/pkg/distillation"
	"github.com/a-clap/embedded/pkg/ds18b20"
	"golang.org/x/exp/slices"
)

type Client interface {
	Get() ([]distillation.DSConfig, error)
	Configure(sensor distillation.DSConfig) (distillation.DSConfig, error)
	Temperatures() ([]distillation.DSTemperature, error)
}

type Listener interface {
	OnDSConfigChange(config parameters.DS)
	OnDSTemperatureChange(temperature parameters.Temperature)
}

type config struct {
	parameters.DS
}

type dsHandler struct {
	client    Client
	listeners []Listener
	running   atomic.Bool
	sensors   map[string]*config
	interval  time.Duration
	finish    chan struct{}
	err       chan<- error
}

var (
	handler = &dsHandler{
		client:    nil,
		listeners: make([]Listener, 0),
		running:   atomic.Bool{},
		sensors:   make(map[string]*config),
		interval:  1 * time.Second,
		finish:    nil,
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
		handler.sensors[s.ID] = &config{
			DS: parameters.DS{DSConfig: s},
		}
	}
	return nil
}

// Stop stops Run
func Stop() {
	if handler.running.Load() {
		close(handler.finish)
		handler.running.Store(false)
	}
}

// Run provides temperature updates
func Run() {
	if handler.running.Load() == false {
		handler.finish = make(chan struct{})
		handler.running.Store(true)
		update()
	}
}

func Apply(config []parameters.DS) []error {
	var errs []error
	for _, c := range config {
		// SetCorrection
		if err := SetCorrection(c.ID, c.Correction); err != nil {
			errs = append(errs, err)
		}

		// SetResolution
		if err := SetResolution(c.ID, c.Resolution); err != nil {
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
	}
	return errs
}

// Refresh read every possible data from Client and serves them to Listener
func Refresh() {
	configs := Get()
	for _, conf := range configs {
		notifyConfig(conf)
	}
}

// AddListener add new listener
func AddListener(listener Listener) {
	handler.listeners = append(handler.listeners, listener)
}

// Get returns all available sensors. It doesn't call notify as it returns value
func Get() []parameters.DS {
	sensors := make([]parameters.DS, 0, len(handler.sensors))
	for _, s := range handler.sensors {
		sensors = append(sensors, s.DS)

	}
	slices.SortFunc(sensors, func(i, j parameters.DS) bool {
		return i.ID < j.ID
	})
	return sensors
}

func Enable(id string, enable bool) error {
	cfg, ok := handler.sensors[id]
	if !ok {
		err := &Error{ID: id, Op: "Enable", Err: ErrIDNotFound.Error()}
		return err
	}

	if cfg.Enabled == enable {
		return nil
	}

	setCfg := cfg.DSConfig
	setCfg.Enabled = enable
	newCfg, err := handler.client.Configure(setCfg)
	if err != nil {
		err = &Error{ID: id, Op: "Enable.Configure", Err: err.Error()}
	} else {
		cfg.DSConfig.Enabled = newCfg.Enabled
	}
	notifyConfig(cfg.DS)
	return err
}

func SetCorrection(id string, correction float64) error {
	cfg, ok := handler.sensors[id]
	if !ok {
		err := &Error{ID: id, Op: "SetCorrection", Err: ErrIDNotFound.Error()}
		return err
	}

	if cfg.Correction == correction {
		return nil
	}

	setCfg := cfg.DSConfig
	setCfg.Correction = correction
	newCfg, err := handler.client.Configure(setCfg)
	if err != nil {
		err = &Error{ID: id, Op: "SetCorrection.Configure", Err: ErrIDNotFound.Error()}
	} else {
		cfg.DSConfig.Correction = newCfg.Correction
	}
	notifyConfig(cfg.DS)
	return err
}

func SetSamples(id string, samples uint) error {
	cfg, ok := handler.sensors[id]
	if !ok {
		err := &Error{ID: id, Op: "SetSamples", Err: ErrIDNotFound.Error()}
		return err
	}

	if cfg.Samples == samples {
		return nil
	}

	setCfg := cfg.DSConfig
	setCfg.Samples = samples
	newCfg, err := handler.client.Configure(setCfg)
	if err != nil {
		err = &Error{ID: id, Op: "SetSamples.Configure", Err: err.Error()}
	} else {
		cfg.DSConfig.Samples = newCfg.Samples
	}

	notifyConfig(cfg.DS)
	return err
}

func SetResolution(id string, resolution ds18b20.Resolution) error {
	cfg, ok := handler.sensors[id]
	if !ok {
		err := &Error{ID: id, Op: "SetResolution", Err: ErrIDNotFound.Error()}
		return err
	}

	if cfg.Resolution == resolution {
		return nil
	}

	setCfg := cfg.DSConfig
	setCfg.Resolution = resolution
	newCfg, err := handler.client.Configure(setCfg)
	if err != nil {
		err = &Error{ID: id, Op: "SetResolution.Configure", Err: err.Error()}
	} else {
		cfg.DSConfig.Resolution = newCfg.Resolution
	}

	notifyConfig(cfg.DS)
	return err
}

func update() {
	go func() {
		for handler.running.Load() {
			select {
			case <-handler.finish:
				handler.running.Store(false)
				break
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
						ID:          temp.ID,
						Temperature: temp.Temperature,
					})
				}
			}
		}
	}()
}
