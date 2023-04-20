package backend

// Events is just way to return constants in UI
type Events struct{}

const (
	NotifyHeaters       = "rcv:heaters"
	NotifyDSConfig      = "rcv:dscfg"
	NotifyDSTemperature = "rcv:dstmp"
	NotifyPTConfig      = "rcv:ptcfg"
	NotifyPTTemperature = "rcv:pttmp"
	NotifyGPIO          = "rcv:gpio"
	NotifyError         = "rcv:error"
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

func (Events) NotifyPTConfig() string {
	return NotifyPTConfig
}

func (Events) NotifyPTTemperature() string {
	return NotifyPTTemperature
}

func (Events) NotifyGPIO() string {
	return NotifyGPIO
}

func (Events) NotifyError() string {
	return NotifyError
}
