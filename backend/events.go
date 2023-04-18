package backend

// Events is just way to return variables in UI
// Each method
type Events struct{}

const (
	NotifyHeaters       = "rcv:heaters"
	NotifyDSConfig      = "rcv:dscfg"
	NotifyDSTemperature = "rcv:dstmp"
)

func (Events) NotifyHeaters() string {
	return NotifyHeaters
}

func (Events) NotifyDSConfig() string {
	return NotifyDSConfig
}

func (Events) NotifyDSTemperature() string {
	return NotifyDSTemperature
}
