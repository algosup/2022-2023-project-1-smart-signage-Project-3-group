package fakeled

type LED struct {
	led pin
	on  bool
}

func NewFake() *LED {
	l := &fakePin{on: false}

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
