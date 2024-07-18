package main

import "testing"

/* TestNumEvenOdd tests the NumEvenOdd function */
func TestNumEvenOdd(t *testing.T) {
	t.Run("even", func(t *testing.T) {
		got := NumEvenOdd(2)
		want := "even"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("odd", func(t *testing.T) {
		got := NumEvenOdd(3)
		want := "odd"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
