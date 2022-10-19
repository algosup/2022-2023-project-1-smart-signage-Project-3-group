package fakeled

type LED struct {
	led pin
	on  bool
}

func NewFake() *LED {
	l := &fakePin{on: true}

	led := LED{
		led: l,
		on:  true,
	}
	return &led
}

func (l *LED) String() string {
	var onStr string
	if l.on {
		onStr = "On"
	} else {
		onStr = "Off"
	}

	return "LED: " + onStr
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

type pin interface {
	Low()
	High()
}

type fakePin struct {
	on bool
}

func (p *fakePin) High() {
	p.on = true
}

func (p *fakePin) Low() {
	p.on = false
}
