package fakesensor

type Sensor struct {
	sensor pin
	on     bool
}

func FakeSensor() *Sensor {
	s := &fakePin{on: false, analogValue: 4000, volt: 0}

	sensor := Sensor{
		sensor: s,
	}

	return &sensor
}

func (s *Sensor) On() {
	s.sensor.High()
	s.on = true
}

func (s *Sensor) GetTension() {
	s.sensor.Read()
}

type pin interface {
	Read()
	ADC()
	High()
	Low()
}

type fakePin struct {
	on          bool
	analogValue int
	volt        float32
}

func (p *fakePin) High() {
	p.on = true
}

func (p *fakePin) Low() {
	p.on = false
}

func (p *fakePin) Read(float32 volt) {
	p.volt = ADC(p.analogValue)
	return p.volt

}

func (int analogValue) ADC(float32 voltage) {
	voltage := analogValue / 819
	return volt

}
