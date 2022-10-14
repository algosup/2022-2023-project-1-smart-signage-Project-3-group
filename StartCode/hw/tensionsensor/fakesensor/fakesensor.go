package fakesensor 


type Sensor struct {
	sensor pin
	on  bool
}

func RealSensor() *Sensor {
	s := machine.Sensor
	s.Configure(machine.PinConfig{Mode: machine.PinOutput})

	sensor := Sensor{
		sensor: s,
	}

	return &sensor
}

func (s *Sensor) On() {
	s.sensor.High()
	s.on = true
}

func (s *Sensor) GetInfos(){
	s.Sensor.Read() 
}