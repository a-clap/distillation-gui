package backend

const (
	ErrDSSetCorrection = iota + 10
	ErrDSEnable
	ErrDSSetSamples
	ErrDSSetResolution
	ErrDSSetName
)
const (
	ErrPTSetCorrection = iota + 20
	ErrPTEnable
	ErrPTSetSamples
	ErrPTSetName
)
const (
	ErrGPIOSetActiveLevel = iota + 30
	ErrGPIOSetState
)

const (
	ErrWIFIAPList = iota + 40
	ErrWifiIsConnected
)

const (
	ErrSetNTP = iota + 45
	ErrSetTime
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

const (
	ErrSave = iota + 70
	ErrLoad
)

const (
	ErrConnDS = iota + 80
	ErrConnPT
	ErrConnPhase
)
