package fakesensor

type Sensor struct {
	//with its fakepin associated
	sensor pin
	on     bool
}

func FakeSensor() *Sensor {
	//create a fake Sensor with its fakepin associated, and also its analogValue
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
	//set the fake sensor's state to High
	s.sensor.High()
	s.on = true
}

func (s *Sensor) GetTension() {
	//Get the tension detected by the fake Sensor
	s.sensor.Read()
}

type pin interface {
	//create a fakepin interface to use methods for the fakepin
	Read() float32
	High()
	Low()
}

type fakePin struct {
	//create a struct which simulates a fakepin for a fake sensor
	//with a "on" bool for designing its state, its analogic value
	//and the voltage he has to detect

	on          bool
	analogValue int
	voltage     float32
}

func (p *fakePin) High() {
	//get the fakepin's memory address and set its state to true
	p.on = true
}

func (p *fakePin) Low() {
	//get the fakepin's memory address and set its state to false
	p.on = false
}

func (p *fakePin) Read() float32 {
	//read the fakepin detected analogValue, and return the voltage
	p.voltage = ADC(p.analogValue)
	return p.voltage

}

func ADC(analogValue int) float32 {
	// Convert an analogic Value to a tension in volts
	volt := float32(analogValue / 819)
	return volt
}
