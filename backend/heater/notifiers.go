/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package heater

import "github.com/a-clap/distillation-gui/backend/parameters"

func notify(config parameters.Heater) {
	for _, listener := range handler.listeners {
		listener.OnHeaterChange(config)
	}
}
