package main

import "testing"

func TestFactorials(t *testing.T) {

	t.Run("find factorials of 5", func(t *testing.T) {
		n := 5

		got := Factorials(n)
		want := 120

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, n)
		}
	})

}
