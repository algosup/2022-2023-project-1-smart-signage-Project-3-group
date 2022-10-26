package fakeled

import (
	"testing"
)

func TestFakeled(t *testing.T) {
	//We test functions related to the fake LED struct

	t.Run("creating a new LED", func(t *testing.T) {
		//We create a new LED struct and read the state of its "on" bool
		light := NewFakeLED("light")
		got := light.String()
		want := "light: Off"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("creating a new LED and changing its state", func(t *testing.T) {
		//we create a new LED struct and change its state from Off to On
		newLed := NewFakeLED("light")
		newLed.Toggle()

		got := newLed.String()
		want := "light: On"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}

func TestFakeSign(t *testing.T) {
	//We test functions related to the fake SIGN struct

	t.Run("creating a new SIGN", func(t *testing.T) {
		//We create a new SIGN struct and read the state of its "on" bool
		sign := NewFakeSign("neon")
		got := sign.String()
		want := "neon: Off"

		if got != want {
			t.Errorf("got %s want %s", got, want)

		}
	})

	t.Run("creating a new SIGN and changing its fakepin state", func(t *testing.T) {
		//we create a new SIGN struct and change its state from Off to On
		newSign := NewFakeSign("neon")
		newSign.Toggle()

		got := newSign.String()
		want := "neon: On"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}

func TestFakeSignWithFakeLeds(t *testing.T) {
	//We test functions related to both fake SIGN and fake LED structs
	t.Run("creating a new SIGN and a new LED part of it", func(t *testing.T) {
		//We create a new SIGN struct and then a new LED struct part of it

		//sign := NewFakeSign("motherSIGN")

		led := NewFakeLED("daughterLED")

		signWithLEd := SIGN{leds: led}

		got := signWithLEd

		want := SIGN[{"daugherLED", false}]

		if got != want {
			t.Errorf("got %#v want %s", got, want)
		}
	})

	t.Run("creating a new SIGN and with a few new leds", func(t *testing.T)

	t.Run("creating a new SIGN with new leds off, and toggle the sign to turn them on", func(t *testing.T)
}
