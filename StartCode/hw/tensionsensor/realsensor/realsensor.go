package realsensor

import "machine"

type Sensor struct {
	sensor machine.ADC
	on     bool
}

func RealSensor() *Sensor {
	s := machine.ADC
	s.Configure(
		//machine.PinConfig{Mode: machine.PinOutput}
		machine.ADCConfig{Reference: AREF, Resolution: 8, Samples: 4})

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
	s.sensor.High()
	s.on = true
}

func (s *Sensor) Off() {
	s.sensor.Low()
	s.on = false
}

func (s *Sensor) GetInfos() {
	return s.Get()
}
