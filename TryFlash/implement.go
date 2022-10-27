package main

import (
	"time"

	//"github.com/tinygo-org/tinygo/src/machine"

	"machine"
)

func main() {
	led := machine.PC13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		time.Sleep(time.Second)
		led.Low()
		time.Sleep(time.Second)
	}
}
