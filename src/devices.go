package main

import "github.com/google/gousb"

type Device struct {
	Name         string
	VID          gousb.ID
	PID          gousb.ID
	ConfigNum    byte
	InterfaceNum byte
	AltSetting   byte
	OutEndpoint  byte
}

var TERRA_NOSTRA_WIRED = Device{
	Name:         "Terra Nostra Wired",
	VID:          0x0c45,
	PID:          0x8006,
	ConfigNum:    1,
	InterfaceNum: 2,
	AltSetting:   0,
	OutEndpoint:  3,
}
