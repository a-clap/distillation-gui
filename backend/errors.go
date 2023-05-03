package backend

const (
	ErrDSSetCorrection = iota + 10
	ErrDSEnable
	ErrDSSetSamples
	ErrDSSetResolution
)
const (
	ErrPTSetCorrection = iota + 20
	ErrPTEnable
	ErrPTSetSamples
)
const (
	ErrGPIOSetActiveLevel = iota + 30
	ErrGPIOSetState
)

const (
	ErrWIFIAPList = iota + 40
)

const (
	ErrPhaseGetCount = iota + 50
	ErrPhaseGetPhaseConfigs
	ErrPhaseGetGlobalConfig
	ErrPhasesSetPhaseCount
	ErrPhasesSetConfig
	ErrPhasesValidateConfig
	ErrPhasesEnable
	ErrPhasesDisable
	ErrPhasesMoveToNext
	ErrPhasesSetGlobalGPIO
)
