package refecttest

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "AA"
	var got []string

	x := struct {
		Name string
	}{expected}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("got '%s', want '%s'", got[0], expected)
	}
}

func TestWalk2(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"A"},
			[]string{"A"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
	}

	for _,test := range cases{
		t.Run(test.Name,func(t *testing.T){
			var got []string
			Walk2(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

func TestWalk3(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}

	for _,test := range cases{
		t.Run(test.Name,func(t *testing.T){
			var got []string
			Walk3(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

func TestWalk4(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}

	for _,test := range cases{
		t.Run(test.Name,func(t *testing.T){
			var got []string
			Walk4(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

func TestWalk5(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Arrays",
			[2]Profile {
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _,test := range cases{
		t.Run(test.Name,func(t *testing.T){
			var got []string
			Walk5(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

func TestWalk6(t *testing.T) {
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		Walk5(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string)  {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain '%s' but it didnt", haystack, needle)
	}
}
