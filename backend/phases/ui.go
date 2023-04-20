/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package phases

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/a-clap/iot/pkg/distillation"
	"github.com/a-clap/iot/pkg/distillation/process"
)

// Client is an interface to read/set listed configs
type Client interface {
	GetPhaseCount() (distillation.ProcessPhaseCount, error)
	GetPhaseConfig(phaseNumber int) (distillation.ProcessPhaseConfig, error)
	ConfigurePhaseCount(count distillation.ProcessPhaseCount) (distillation.ProcessPhaseCount, error)
	ConfigurePhase(phaseNumber int, setConfig distillation.ProcessPhaseConfig) (distillation.ProcessPhaseConfig, error)
	ValidateConfig() (distillation.ProcessConfigValidation, error)
	ConfigureProcess(cfg distillation.ProcessConfig) (distillation.ProcessConfig, error)
	Status() (distillation.ProcessStatus, error)
}

// Listener allows to be notified about changes in listed configs
type Listener interface {
	OnPhasesCountChange(count distillation.ProcessPhaseCount)
	OnPhaseConfigChange(phaseNumber int, cfg distillation.ProcessPhaseConfig)
	OnConfigValidate(validation distillation.ProcessConfigValidation)
	OnStatusChange(status distillation.ProcessStatus)
	OnConfigChange(c distillation.ProcessConfig)
}

type processHandler struct {
	client    Client
	count     distillation.ProcessPhaseCount
	phases    map[int]*distillation.ProcessPhaseConfig
	listeners []Listener
	status    distillation.ProcessStatus
	running   atomic.Bool
	enabled   atomic.Bool
	interval  time.Duration
	err       chan<- error
	finish    chan struct{}
}

var (
	handler = &processHandler{
		client:    nil,
		count:     distillation.ProcessPhaseCount{},
		phases:    make(map[int]*distillation.ProcessPhaseConfig, 10),
		listeners: make([]Listener, 0),
		status:    distillation.ProcessStatus{},
		running:   atomic.Bool{},
		interval:  1 * time.Second,
		err:       nil,
	}
)

// Init prepare package to handle various requests
func Init(c Client, err chan<- error, interval time.Duration) {
	handler.client = c
	handler.err = err
	handler.interval = interval
}

func Apply(config process.Config) []error {
	var errs []error

	err := SetPhaseCount(config.PhaseNumber)
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	for i, c := range config.Phases {
		if err := SetConfig(i, distillation.ProcessPhaseConfig{PhaseConfig: c}); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// Refresh read every possible data from Client and serves them to Listener
func Refresh() {
	configs, err := GetPhaseConfigs()
	if err == nil {
		notifyProcessCount(distillation.ProcessPhaseCount{PhaseNumber: len(configs)})
		for i, conf := range configs {
			notifyConfigChange(i, conf)
		}
	}

	_ = ValidateConfig()
	s, err := handler.client.Status()
	if err == nil {
		notifyStatus(s)
	}
}

func Stop() {
	if handler.running.Load() {
		handler.running.Store(false)
		close(handler.finish)
	}
}

func Run() {
	if !handler.running.Load() {
		handler.finish = make(chan struct{})
		handler.enabled.Store(true)
		update()
	}
}

// AddListener adds listener. Each listener is called after config changes
func AddListener(listener Listener) {
	handler.listeners = append(handler.listeners, listener)
}

// GetPhaseCount returns current PhaseCount. It doesn't call any notify, as it return value
func GetPhaseCount() (distillation.ProcessPhaseCount, error) {
	c, err := handler.client.GetPhaseCount()
	if err != nil {
		return c, err
	}

	handler.count = c
	return handler.count, err
}

// GetPhaseConfigs returns slice of current configs. It doesn't call any notify, as it return value
func GetPhaseConfigs() ([]distillation.ProcessPhaseConfig, error) {
	c, err := GetPhaseCount()
	if err != nil {
		return nil, err
	}

	var configs []distillation.ProcessPhaseConfig
	for i := 0; i < c.PhaseNumber; i++ {
		cfg, err := handler.client.GetPhaseConfig(i)
		if err != nil {
			return nil, err
		}
		handler.phases[i] = &cfg
		configs = append(configs, cfg)
	}

	return configs, nil
}

// SetPhaseCount sets distillation.ProcessPhaseCount and notify listeners about change
func SetPhaseCount(count int) error {
	log.Println("setphasecount")
	c := distillation.ProcessPhaseCount{PhaseNumber: count}
	c, err := handler.client.ConfigurePhaseCount(c)
	if err != nil {
		log.Println("err")
		return err
	}
	notifyProcessCount(c)
	return nil
}

func SetConfig(number int, cfg distillation.ProcessPhaseConfig) error {
	c, err := handler.client.ConfigurePhase(number, cfg)
	if err != nil {
		err := &Error{Op: "SetConfig.ConfigurePhase", Err: err.Error()}
		if c, ok := handler.phases[number]; ok {
			notifyConfigChange(number, *c)
		}
		return err
	}
	notifyConfigChange(number, c)
	return nil
}

func ValidateConfig() error {
	v, err := handler.client.ValidateConfig()
	if err != nil {
		return &Error{Op: "ValidateConfig", Err: err.Error()}
	}
	notifyValidate(v)
	return nil
}

func Enable() error {
	if handler.status.Running {
		return ErrRunning
	}

	config := distillation.ProcessConfig{
		Enable:     true,
		MoveToNext: false,
		Disable:    false,
	}
	var err error
	if config, err = handler.client.ConfigureProcess(config); err != nil {
		return &Error{Op: "process.Enable", Err: err.Error()}
	}
	handler.enabled.Store(true)
	notifyConfig(config)
	return nil
}

func Disable() error {
	if !handler.status.Running {
		return ErrDisabled
	}

	config := distillation.ProcessConfig{
		Enable:     false,
		MoveToNext: false,
		Disable:    true,
	}

	var err error
	if config, err = handler.client.ConfigureProcess(config); err != nil {
		return &Error{Op: "process.Disable", Err: err.Error()}
	}
	notifyConfig(config)
	return nil
}

func MoveToNext() error {
	if !handler.status.Running {
		return ErrRunning
	}

	config := distillation.ProcessConfig{
		Enable:     false,
		MoveToNext: true,
		Disable:    false,
	}

	var err error
	if config, err = handler.client.ConfigureProcess(config); err != nil {
		return &Error{Op: "process.Disable", Err: err.Error()}
	}
	notifyConfig(config)
	return nil
}

func update() {
	handler.running.Store(true)
	go func() {
		for handler.running.Load() {
			select {
			case <-handler.finish:
				handler.running.Store(false)
			case <-time.After(handler.interval):
				if handler.enabled.Load() {
					s, err := handler.client.Status()
					if err != nil {
						err := &Error{Op: "process.Status", Err: err.Error()}
						notifyError(err)
						continue
					}
					handler.status = s
					notifyStatus(s)
					if !handler.status.Running {
						handler.enabled.Store(false)
					}
				}
			}
		}
	}()

}