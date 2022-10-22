package fakeled

type LED struct { //LED struct with its pin and its state
	led pin
	on  bool
}

func NewFake() *LED { //Create a fakeLED with its fakepin associated

	l := &fakePin{on: false}

	led := LED{
		led: l,
		on:  false,
	}
	return &led
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
	//change the fakepin state
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
