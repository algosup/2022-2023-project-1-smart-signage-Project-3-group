package fakesensor

type Sensor struct {
	sensor pin
	on     bool
}

func FakeSensor() *Sensor {
	s := &fakePin{on: false}

	sensor := Sensor{
		sensor: s,
	}

	return &sensor
}

func (s *Sensor) On() {
	s.sensor.High()
	s.on = true
}

func (s *Sensor) GetInfos() {
	s.sensor.Read()
}

type pin interface {
	Read()
	High()
	Low()
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

func (p *fakePin) Read() {

}
