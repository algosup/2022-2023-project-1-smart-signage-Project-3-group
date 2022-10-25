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
		//we create a new LED struct and change its state from Off to On
		newSign := NewFakeSign("neon")
		newSign.Toggle()

		got := newSign.String()
		want := "neon: On"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
