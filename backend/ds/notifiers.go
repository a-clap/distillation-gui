/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package ds

import (
	"github.com/a-clap/distillation-gui/backend/parameters"
)

func notifyConfig(config parameters.DS) {
	for _, listener := range handler.listeners {
		listener.OnDSConfigChange(config)
	}
}

func notifyTemperature(temperature parameters.Temperature) {
	for _, listener := range handler.listeners {
		listener.OnDSTemperatureChange(temperature)
	}
}

// notifyError notifies user about error, which can't be return via return
func notifyError(e error) {
	if handler.err != nil {
		// error is discarded, if channel is nil
		handler.err <- e
	}
}
