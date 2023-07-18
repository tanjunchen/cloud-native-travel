package iterator

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat2("a")
	}
}

func TestRepeat3(t *testing.T) {
	var count = 10
	repeated := Repeat3("a", count)
	expected := "aaaaaaaaaa"
	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}
