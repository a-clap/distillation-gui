package backend

import (
	"github.com/a-clap/distillation-gui/backend/wifi"
	"github.com/a-clap/logging"
)

type WifiConnected struct {
	Connected bool   `json:"connected"`
	AP        string `json:"AP"`
}

func (b *Backend) WifiAPList() []string {
	logger.Debug("WifiAPList")

	aps, err := wifi.AP()

	if err != nil {
		logger.Error("WifiAPList", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrWIFIAPList)

		return nil
	}

	return aps
}

func (b *Backend) WifiIsConnected() WifiConnected {
	logger.Debug("WifiIsConnected")

	conn, ap, err := wifi.Connected()

	if err != nil {
		logger.Error("WifiIsConnected", logging.String("error", err.Error()))
		b.eventEmitter.OnError(ErrWifiIsConnected)

		return WifiConnected{}
	}

	return WifiConnected{
		Connected: conn,
		AP:        ap,
	}
}
