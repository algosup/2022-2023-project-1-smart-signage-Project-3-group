package fakeled

import (
	"testing"
)

func TestFakeled(t *testing.T) {
	//We test functions related to the fakeled package

	t.Run("creating a new LED", func(t *testing.T) {
		//We create a new LED struct and read the state of its "on" bool
		light := NewFakeLED()
		got := light.String()
		want := "LED: Off"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("creating a new LED and changing its state", func(t *testing.T) {
		//we create a new LED struct and change its state from Off to On
		newLed := NewFakeLED()
		newLed.Toggle()

		got := newLed.String()
		want := "LED: On"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}

func TestFakeSign(t *testing.T) {
	//We create a new LED struct and read the state of its "on" bool
	sign := NewFakeSign()
	got := sign.String()
	want := "SIGN: Off"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
