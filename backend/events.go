package backend

// Events is just way to return constants in UI
type Events struct{}

const (
	NotifyHeaters           = "rcv:heaters"
	NotifyDSConfig          = "rcv:dscfg"
	NotifyDSTemperature     = "rcv:dstmp"
	NotifyPTConfig          = "rcv:ptcfg"
	NotifyPTTemperature     = "rcv:pttmp"
	NotifyGPIO              = "rcv:gpio"
	NotifyError             = "rcv:error"
	NotifyPhasesConfig      = "rcv:phase_config"
	NotifyPhasesValidate    = "rcv:phase_validate"
	NotifyPhasesPhaseConfig = "rcv:phase_phase_config"
	NotifyPhasesPhaseCount  = "rcv:phase_count"
	NotifyPhasesStatus      = "rcv:phase_status"
	NotifyGlobalConfig      = "rcv:global_config"
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

func (Events) NotifyPhasesConfig() string {
	return NotifyPhasesConfig
}
func (Events) NotifyPhasesValidate() string {
	return NotifyPhasesValidate
}
func (Events) NotifyPhasesPhaseConfig() string {
	return NotifyPhasesPhaseConfig
}
func (Events) NotifyPhasesPhaseCount() string {
	return NotifyPhasesPhaseCount
}
func (Events) NotifyPhasesStatus() string {
	return NotifyPhasesStatus
}
func (Events) NotifyGlobalConfig() string {
	return NotifyGlobalConfig
}
