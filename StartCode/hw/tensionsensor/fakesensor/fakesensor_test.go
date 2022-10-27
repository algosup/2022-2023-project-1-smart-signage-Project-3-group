package fakesensor

import (
	"testing"
)

func TestFakesensor(t *testing.T) {
	//We test functions related to the fakesensor package

	t.Run("creating a new tensionSensor", func(t *testing.T) {
		//We create a new LED struct and read the state of its "on" bool
		sensor := FakeSensor()
		got := sensor.String()
		want := "Sensor: Off"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}

func TestADC(t *testing.T) {
	//We test the ADC functions

	t.Run("Analogic value: 819", func(t *testing.T) {
		analog := 819
		got := ADC(analog)
		want := float32(1.0)

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("Analogic value: 4095", func(t *testing.T) {
		analog := 4095
		got := ADC(analog)
		want := float32(5.0)

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("Analogic value: 0", func(t *testing.T) {
		analog := 0
		got := ADC(analog)
		want := float32(0.0)

		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}
