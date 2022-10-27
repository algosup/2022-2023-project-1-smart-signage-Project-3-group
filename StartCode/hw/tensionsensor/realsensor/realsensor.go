package realsensor

import "machine"

type Sensor struct {
	//Sensor struct with the chosen ADC pin and its state
	sensor machine.ADC
	on     bool
}

func RealSensor() *Sensor {
	// initialize bluepill ADC mechanism
	machine.InitADC()

	s := machine.ADC{Pin: machine.ADC5}

	s.Configure(machine.ADCConfig{})

	// OR with the values you need
	//s.Configure(machine.ADCConfig{
	//Reference:  3300, // or 5000mV, I don't know which on you need
	//Resolution: 8,    // I don't know which bit value you need here
	//Samples:    16,   // I don't know what you need here, either
	//})

	//ADCConfig{
	//	Reference uint32	// analog reference voltage (AREF) in millivolts
	//	,Resolution	uint32	// number of bits for a single conversion (e.g., 8, 10, 12)
	//	,Samples	uint32// number of samples for a single conversion (e.g., 4,})
	// }

	sensor := Sensor{
		sensor: s,
	}

	return &sensor
}

func (s *Sensor) On() {
	//set the sensor's state to High
	s.sensor.High()
	s.on = true
}

func (s *Sensor) Off() {
	//set the sensor's state to Low
	s.sensor.Low()
	s.on = false
}

func (s *Sensor) GetTension() {
	//get the tension on a given pin
	return s.Get()
}
