package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := 2
	actual := HelloWorld()

	if actual != expected {
		t.Errorf("Expect hello, but got %d!", actual)
	}
}

func TestSum(t *testing.T) {
	set := []int{17, 23, 100, 76, 55}
	expected := 271
	actual := Sum(set)

	if actual != expected {
		t.Errorf("Expect %d, but got %d!", expected, actual)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("tanjunchen")
		want := "Hello,tanjunchen"
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("tanjunchen")
		want := "Hello,World"
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
}

func TestHelloRefactor(t *testing.T) {
	// refactor common code
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloRefactor("tanjunchen")
		want := "Hello,tanjunchen"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := HelloRefactor("World")
		want := "Hello,World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := HelloParameter("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
