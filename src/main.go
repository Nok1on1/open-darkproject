package main

import (
	"flag"
	"image/color"
	"log"

	"github.com/google/gousb"
)

func main() {
	brightness := flag.Int("b", 100, "brightness")
	mode := flag.Int("m", 1, "mode")
	flag.Parse()

	ctx := gousb.NewContext()
	defer ctx.Close()

	keyboard := TERRA_NOSTRA_WIRED

	dev, err := ctx.OpenDeviceWithVIDPID(keyboard.VID, keyboard.PID)
	if err != nil || dev == nil {
		log.Fatalf("Could not open a device: %s \n %v", keyboard.Name, err)
		return
	}
	defer dev.Close()

	dev.SetAutoDetach(true)

	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatal(err)
	}
	defer cfg.Close()

	intf, err := cfg.Interface(int(keyboard.InterfaceNum), int(keyboard.AltSetting))
	if err != nil {
		log.Fatalf("%s config.Interface(%d,%d): %v", dev, keyboard.InterfaceNum, keyboard.AltSetting, err)
	}
	defer intf.Close()

	endpoint, err := intf.OutEndpoint(int(keyboard.OutEndpoint))
	if err != nil {
		log.Fatalf("%s.OutEndpoint(%d): %v", intf, keyboard.OutEndpoint, err)
	}

	data, err := newPayload(*brightness, color.RGBA{
		R: 255,
		G: 255,
		B: 255,
	}, Mode(*mode), false)
	if err != nil {
		log.Fatal(err)
	}

	payload := data.GeneratePayload()

	numBytes, err := endpoint.Write(payload)
	if numBytes != len(payload) {
		log.Fatalf("%s.Write(): only %d bytes written, returned error is %v", endpoint, numBytes, err)
	}

	log.Println(dev.String())
}
