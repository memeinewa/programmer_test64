package main

import "testing"

func TestCommission(t *testing.T) {

	t.Run("find commission of 30,000", func(t *testing.T) {
		var n float32 = 30000

		got := Commission(n)
		var want float32 = 3000

		if got != want {
			t.Errorf("got %f want %f given, %v", got, want, n)
		}
	})

	t.Run("find commission of 60,000", func(t *testing.T) {
		var n float32 = 60000

		got := Commission(n)
		var want float32 = 6000

		if got != want {
			t.Errorf("got %f want %f given, %v", got, want, n)
		}
	})

	t.Run("find commission of 17,000", func(t *testing.T) {
		var n float32 = 17000

		got := Commission(n)
		var want float32 = 850

		if got != want {
			t.Errorf("got %f want %f given, %v", got, want, n)
		}
	})

	t.Run("find commission of 20,000", func(t *testing.T) {
		var n float32 = 20000

		got := Commission(n)
		var want float32 = 1000

		if got != want {
			t.Errorf("got %f want %f given, %v", got, want, n)
		}
	})

	t.Run("find commission of 0", func(t *testing.T) {
		var n float32 = 0

		got := Commission(n)
		var want float32 = 0

		if got != want {
			t.Errorf("got %f want %f given, %v", got, want, n)
		}
	})

	t.Run("find commission of 10,000", func(t *testing.T) {
		var n float32 = 10000

		got := Commission(n)
		var want float32 = 0

		if got != want {
			t.Errorf("got %f want %f given, %v", got, want, n)
		}
	})

}

func TestCompareCommission(t *testing.T) {

	t.Run("compare commission when a has commission = 30,000 and b has commission = 17,000", func(t *testing.T) {
		var commissionA float32 = 30000
		var commissionB float32 = 17000

		gotA := Commission(commissionA)
		var wantA float32 = 3000

		if gotA != wantA {
			t.Errorf("got %f want %f given, %v", gotA, wantA, commissionA)
		}

		gotB := Commission(commissionB)
		var wantB float32 = 850

		if gotB != wantB {
			t.Errorf("got %f want %f given, %v", gotB, wantB, commissionB)
		}

		gotAB := CompareCommission(gotA, gotB)
		var wantAB float32 = wantA - wantB

		if gotAB != wantAB {
			t.Errorf("got %f want %f given, %v and %v", gotAB, wantAB, gotA, gotB)
		}
	})
}
