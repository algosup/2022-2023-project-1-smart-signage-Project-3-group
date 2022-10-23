package fakesensor

type Sensor struct {
	sensor pin
	on     bool
}

func FakeSensor() *Sensor {
	s := &fakePin{on: false, analogValue: 4000, voltage: 0}

	sensor := Sensor{
		sensor: s,
	}

	return &sensor
}

func (s *Sensor) String() string {
	//display the fake Sensor's state
	var onStr string
	if s.on {
		onStr = "On"
	} else {
		onStr = "Off"
	}

	return "Sensor: " + onStr
}

func (s *Sensor) On() {
	s.sensor.High()
	s.on = true
}

func (s *Sensor) GetTension() {
	s.sensor.Read()
}

type pin interface {
	Read() float32
	High()
	Low()
}

type fakePin struct {
	on          bool
	analogValue int
	voltage     float32
}

func (p *fakePin) High() {
	p.on = true
}

func (p *fakePin) Low() {
	p.on = false
}

func (p *fakePin) Read() float32 {
	p.voltage = ADC(p.analogValue)
	return p.voltage

}

func ADC(analogValue int) float32 {
	volt := float32(analogValue / 819)
	return volt
}
