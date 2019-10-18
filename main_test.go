package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "1"},
	}

	for _, c := range cases {
		got := FizzBuzz(c.in)

		if got != c.want {
			t.Errorf("want %s, but got %s", c.want, got)
		}
	}
}
