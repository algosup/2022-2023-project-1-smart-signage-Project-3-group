package realled

import "machine"

type LED struct {
	led machine.Pin
	on  bool
}

func NewReal() *LED {
	l := machine.LED
	l.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led := LED{
		led: l,
	}

	return &led
}

func (l *LED) On() {
	l.led.High()
	l.on = true
}

func (l *LED) Off() {
	l.led.Low()
	l.on = false
}

func (l *LED) Toggle() {
	if !l.on {
		l.On()
	} else {
		l.Off()
	}
}


