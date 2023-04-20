/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package phases

import (
	"errors"
)

var (
	ErrRunning  = errors.New("already enabled")
	ErrDisabled = errors.New("already disabled")
)

// Error is common error returned by this package
type Error struct {
	Op  string `json:"op"`
	Err string `json:"error"`
}

func (e *Error) Error() string {
	if e.Err == "" {
		return "<nil>"
	}
	s := e.Op + ": " + e.Err
	return s
}
