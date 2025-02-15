package reflection

import (
	"reflect"
	"testing"
)

// test driven development is the practice of writing tests first, and then writing code to satisfy the tests. This ensures our code satisfies
// behaviours we want it to have, so we make an efficient function for instance.
func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Sam"},
			[]string{"Sam"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Sam", "London"},
			[]string{"Sam", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
