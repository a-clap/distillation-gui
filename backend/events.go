package backend

// Events is just way to return variables in UI
// Each method
type Events struct{}

const (
	NotifyHeaters = "rcv:heaters"
)

func (Events) NotifyHeaters() string {
	return NotifyHeaters
}
