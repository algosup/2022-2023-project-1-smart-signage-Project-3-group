package fakeled

type SIGN struct {
	signName string
	leds     []LED
	output   pin
	pinState bool
}

type LED struct { //LED struct with its pin and its state
	led pin
	on  bool
}

func NewFakeSign() *SIGN { //Create a fakeSIGN with its fakepin associated

	s := &fakePin{on: false}

	sign := SIGN{
		signName: "neon",
		output:   s,
		pinState: s.on,
	}
	return &sign
}

func NewFakeLED() *LED { //Create a fakeLED

	l := &fakePin{on: false}

	led := LED{
		led: l,
		on:  false,
	}
	return &led
}

func (s *SIGN) String() string {
	//display the fake SIGN's state
	var onStr string
	if s.pinState {
		onStr = "On"
	} else {
		onStr = "Off"
	}

	return "SIGN: " + onStr
}

func (l *LED) String() string {
	//display the fake LED's state
	var onStr string
	if l.on {
		onStr = "On"
	} else {
		onStr = "Off"
	}

	return "LED: " + onStr
}

func (l *LED) On() {
	//set the fake led's state to High
	l.led.High()
	l.on = true
}

func (l *LED) Off() {
	//set the fake led's state to Low
	l.led.Low()
	l.on = false
}

func (l *LED) Toggle() {
	//change the fakeLED state
	if !l.on {
		// if its pin state is false, switch it to true
		l.On()
	} else {
		// if its pin state is true, switch it to false
		l.Off()
	}
}

type pin interface {
	//create an interface to simulate the Low and High state of a fake pin
	Low()
	High()
}

type fakePin struct {
	//create a struct which simulates a fake pin for a fake LED with a "on" bool for designing its state
	on bool
}

func (p *fakePin) High() {
	p.on = true
	//get the fakepin's memory address and set its state to true
}

func (p *fakePin) Low() {
	p.on = false
	//get the fakepin's memory address and set its state to false
}
