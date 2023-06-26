package backend

import (
	"fmt"
	"time"
)

func (b *Backend) SetSystemTime(time.Time) error {
	return fmt.Errorf("not supported on windows")
}
