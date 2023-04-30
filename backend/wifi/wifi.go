/*
 * Copyright (c) 2023 a-clap. All rights reserved.
 * Use of this source code is governed by a MIT-style license that can be found in the LICENSE file.
 */

package wifi

import (
	"errors"

	"github.com/a-clap/distillation/pkg/wifi"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	handler           = new(wifiHandler)
)

type Client interface {
	APs() ([]wifi.AP, error)
	Connected() (wifi.Status, error)
	Disconnect() error
	Connect(n wifi.Network) error
}

type Listener interface {
	OnWifiConfigChange()
}

type wifiHandler struct {
	Client Client
}

func Init(client Client) {
	handler.Client = client
}

func IsImplemented() bool {
	return handler.Client != nil
}

func isNotImplemented() bool {
	return !IsImplemented()
}

func Connect(ap, psk string) error {
	if isNotImplemented() {
		return ErrNotImplemented
	}
	w := wifi.Network{
		AP: wifi.AP{
			ID:   0,
			SSID: ap,
		},
		Password: "psk",
	}

	return handler.Client.Connect(w)
}

func Connected() (connected bool, ap string, err error) {
	if isNotImplemented() {
		err = ErrNotImplemented
		return
	}
	s, err := handler.Client.Connected()
	if err != nil {
		return
	}
	return s.Connected, s.SSID, nil
}

func AP() ([]string, error) {
	if isNotImplemented() {
		return nil, ErrNotImplemented
	}
	nets, err := handler.Client.APs()
	if err != nil {
		return nil, err
	}
	aps := make([]string, len(nets))
	for i, net := range nets {
		aps[i] = net.SSID
	}
	return aps, nil
}

func Disconnect() error {
	if isNotImplemented() {
		return ErrNotImplemented
	}
	return handler.Client.Disconnect()
}
