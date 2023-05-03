/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package parameters

import (
	"github.com/a-clap/distillation/pkg/distillation"
	"github.com/a-clap/distillation/pkg/process"
)

type Heater struct {
	ID      string `json:"ID"`
	Enabled bool   `json:"enabled"`
}

type DS struct {
	distillation.DSConfig
}

type Temperature struct {
	distillation.Temperature
}

type PT struct {
	distillation.PTConfig
}

type GPIO struct {
	distillation.GPIOConfig
}

type GUI struct {
	Heaters []Heater       `json:"heaters"`
	DS      []DS           `json:"DS"`
	PT      []PT           `json:"PT"`
	GPIO    []GPIO         `json:"GPIO"`
	Process process.Config `json:"process"`
}
