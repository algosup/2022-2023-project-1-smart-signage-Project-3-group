package fakeled

import "testing"

func TestLed(t *testing.T) {

	led := LED{}

	got := NewFake(led)
	want := &led

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
