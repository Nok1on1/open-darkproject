package main

import (
	"fmt"
	"image/color"
	"math"
)

type UsbPayload struct {
	brightness byte
	colors     [3]byte
	mode       Mode
	multicolor byte
}

type Mode byte

const (
	SOLID_MODE Mode = iota + 1
	TRIGGER
	FOOTPRINT
	STAR
	RAIN
	RANDOM
	//more here
)

func newPayload(brightness int, color color.RGBA, mode Mode, multicolor bool) (UsbPayload, error) {
	if brightness > 100 || brightness < 0 {
		return UsbPayload{}, fmt.Errorf("brightness should be between 0-100")
	}
	brightnessByte := byte(math.Ceil(float64(brightness) / 20))

	return UsbPayload{
		brightness: brightnessByte,
		colors:     [3]byte{color.R, color.G, color.B},
		mode:       mode,
		multicolor: 0,
	}, nil
}

func (u *UsbPayload) GeneratePayload() []byte {
	return []byte{
		0xaa, 0x23, 0x38, 0x00, 0x00, 0x00, 0x00, 0x00, byte(u.mode), u.colors[0], u.colors[1], u.colors[2], 0xff, 0x00, 0x00, 0x00,
		0x00, u.brightness, 0x05, 0x00, 0x00, 0x00, 0xaa, 0x55, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
}
