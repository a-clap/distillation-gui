/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package phases

import (
	"log"

	"github.com/a-clap/distillation/pkg/distillation"
	"github.com/a-clap/distillation/pkg/process"
	"golang.org/x/exp/slices"
)

// notifyProcessCount notifies listeners about distillation.ProcessPhaseCount change
func notifyProcessCount(count distillation.ProcessPhaseCount) {
	log.Println("notifypcount")
	for _, listener := range handler.listeners {
		listener.OnPhasesCountChange(count)
	}
}

func notifyConfigChange(phaseNumber int, cfg distillation.ProcessPhaseConfig) {
	for _, listener := range handler.listeners {
		listener.OnPhaseConfigChange(phaseNumber, cfg)
	}
}
func notifyValidate(v distillation.ProcessConfigValidation) {
	for _, listener := range handler.listeners {
		listener.OnConfigValidate(v)
	}
}
func notifyStatus(s distillation.ProcessStatus) {
	// Sort sensors
	slices.SortFunc(s.Temperature, func(i, j process.TemperaturePhaseStatus) bool {
		return i.ID < j.ID
	})
	// Sort GPIO
	slices.SortFunc(s.GPIO, func(i, j process.GPIOPhaseStatus) bool {
		return i.ID < j.ID
	})
	// And heaters
	slices.SortFunc(s.Heaters, func(i, j process.HeaterPhaseStatus) bool {
		return i.ID < j.ID
	})
	for _, listener := range handler.listeners {
		listener.OnStatusChange(s)
	}
}
func notifyConfig(c distillation.ProcessConfig) {
	for _, listener := range handler.listeners {
		listener.OnConfigChange(c)
	}
}

// notifyError notifies user about error, which can't be return via return
func notifyError(e error) {
	if handler.err != nil {
		// error is discarded, if channel is nil
		handler.err <- e
	}
}
