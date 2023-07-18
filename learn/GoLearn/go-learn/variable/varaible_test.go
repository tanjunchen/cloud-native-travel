package variable

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := 0
	actual := variable()

	if actual != expected {
		t.Errorf("Expect %d, but got %d!", expected, actual)
	}
}
