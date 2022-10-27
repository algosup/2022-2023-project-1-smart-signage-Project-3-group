package tensionsensor

type Sensor interface {
	//Create a Sensor interface for using a method in both fakesensor and realsensor
	GetTension()
}
