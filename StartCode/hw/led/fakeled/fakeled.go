package fakeled

type SIGN struct {
	//SIGN struct with its name, the leds it contains, its fake output and its state
	signName string
	leds     []LED
	output   pin
	pinState bool
}

type LED struct {
	//LED struct with its name and its state
	ledName string
	on      bool
}

func NewFakeSign(nameString string) *SIGN {
	//Create a fakeSIGN with its fakepin associated

	s := &fakePin{on: false}

	sign := SIGN{
		signName: nameString,
		output:   s,
		pinState: s.on,
	}
	return &sign
}

func NewFakeLED(nameLed string) *LED {
	//Create a fakeLED set on false

	led := LED{
		ledName: nameLed,
		on:      false,
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

	return s.signName + ": " + onStr
}

func (l *LED) String() string {
	//display the fake LED's state
	var onStr string
	if l.on {
		onStr = "On"
	} else {
		onStr = "Off"
	}

	return l.ledName + ": " + onStr
}

func (s *SIGN) On() {
	//set the fake led's state to High
	s.output.High()
	s.pinState = true
}

func (s *SIGN) Off() {
	//set the fake led's state to Low
	s.output.Low()
	s.pinState = false
}

func (l *LED) On() {
	//set the fake led's state to High
	l.on = true
}

func (l *LED) Off() {
	//set the fake led's state to Low
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

func (s *SIGN) Toggle() {
	//change the fakeLED state
	if !s.pinState {
		// if its pin state is false, switch it to true
		s.On()
	} else {
		// if its pin state is true, switch it to false
		s.Off()
	}
}

type pin interface {
	//create an interface to simulate the Low and High state of a fake pin
	Low()
	High()
}

type fakePin struct {
	//create a struct which simulates a fake pin for a fake SIGN with a "on" bool for designing its state
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
