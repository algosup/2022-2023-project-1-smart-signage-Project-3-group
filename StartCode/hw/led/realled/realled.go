package realled

import "machine"

//Use the machine package to use the hardware related functions

type LED struct { //LED struct with the chosen pin and its state
	led machine.Pin
	on  bool
}

func NewReal() *LED { //detect a new LED and configure the chosen pin as its output
	l := machine.PC13
	l.Configure(machine.PinConfig{Mode: machine.PinOutput})

	led := LED{
		led: l,
	}

	return &led
}

func (l *LED) On() {
	//set the led's state to High
	l.led.High()
	l.on = true
}

func (l *LED) Off() {
	//set the led's state to Low
	l.led.Low()
	l.on = false
}

func (l *LED) Toggle() {
	//change the pin state
	if !l.on {
		// if its pin state is false, switch it to true
		l.On()
	} else {
		// if its pin state is true, switch it to false
		l.Off()
	}
}

func (l LED) String() string {
	//display the LED's state
	var onStr string
	if l.on {
		onStr = "On"
	} else {
		onStr = "Off"
	}

	return "LED: " + onStr
}
