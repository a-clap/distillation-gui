/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package gpio

import (
	"github.com/a-clap/distillation-gui/backend/parameters"
)

func notifyConfig(config parameters.GPIO) {
	for _, listener := range handler.listeners {
		listener.OnGPIOChange(config)
	}
}
