package mockingpractice


import (
	"bytes"
	"testing"
)

func TestCountDown(t *testing.T){
	buffer := &bytes.Buffer{}

	spySleeper := &SpySleeper{}

	Countdown(buffer,spySleeper)

	got := buffer.String()
	want := `3
2
hello
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}