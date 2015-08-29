package main

import (
	"testing"
)

func TestBold(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello", "\033[1mhello\033[0m"},
	}
	for _, c := range cases {
		got := bold(c.in)
		if got != c.want {
			t.Errorf("bold(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
