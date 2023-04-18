/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package ds

import (
	"errors"
)

var (
	ErrIDNotFound = errors.New("id not found")
)

// Error is common error returned by this package
type Error struct {
	ID  string `json:"ID"`
	Op  string `json:"op"`
	Err string `json:"error"`
}

func (e *Error) Error() string {
	if e.Err == "" {
		return "<nil>"
	}
	s := e.Op
	if e.ID != "" {
		s += ":" + e.ID
	}
	s += ": " + e.Err
	return s
}
