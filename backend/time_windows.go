package backend

import (
	"errors"
)

const (
	ErrNotImplemented = errors.New("not implemented")
)

func (b *Backend) NTPGet() (bool, error) {
	return false, ErrNotImplemented
}

func (b *Backend) NTPSet(enable bool) error {
	return ErrNotImplemented
}

func (b *Backend) TimeSet(ts int64) error {
	return ErrNotImplemented
}
