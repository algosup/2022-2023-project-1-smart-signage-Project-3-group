package realled

import (
	"testing"
)

func TestFakeled(t *testing.T) {
	//We test functions related to the realled package

	t.Run("connecting with a LED", func(t *testing.T) {
		//We create a new LED struct and read the state of its "on" bool
		light := NewReal()
		got := light.String()
		want := "LED: Off"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
