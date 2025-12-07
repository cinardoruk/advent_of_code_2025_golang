package main

import "testing"

func TestGetDialPosition(t *testing.T) {
	t.Run("turning dial left within 0-99", func(t *testing.T) {
		numZeroes := 0
		got := getDialPosition(50, 5, &numZeroes)
		want := 55

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("turning dial right within 0-99", func(t *testing.T) {
		numZeroes := 0
		got := getDialPosition(50, -5, &numZeroes)
		want := 45

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("turning dial right from 0", func(t *testing.T) {
		numZeroes := 0
		got := getDialPosition(0, -2, &numZeroes)
		want := 98

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("turning dial left from 99", func(t *testing.T) {
		numZeroes := 0
		got := getDialPosition(99, 2, &numZeroes)
		want := 1

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("accumulate one 0 hit going right from 2", func(t *testing.T) {
		numZeroes := 0
		got := getDialPosition(2, -3, &numZeroes)
		want := 99
		want0s := 1

		if got != want || numZeroes != 1 {
			t.Errorf("got to %d, accumulated %d. wanted %d and %d", got, want0s, want, want0s)
		}
	})

	t.Run("accumulate one 0 hit going left from 98", func(t *testing.T) {
		numZeroes := 0
		got := getDialPosition(98, 3, &numZeroes)
		want := 1
		want0s := 1

		if got != want || numZeroes != want0s {
			t.Errorf("got to %d, accumulated %d. wanted %d and %d", got, want0s, want, want0s)
		}
	})

	t.Run("accumulate no 0 hits going left from 0", func(t *testing.T) {
		numZeroes := 0
		got := getDialPosition(0, 1, &numZeroes)
		want := 1
		want0s := 0

		if got != want || numZeroes != want0s {
			t.Errorf("got to %d, accumulated %d. wanted %d and %d", got, numZeroes, want, want0s)
		}
	})

	t.Run("accumulate no 0 hits going right from 0", func(t *testing.T) {
		numZeroes := 0
		got := getDialPosition(0, -1, &numZeroes)
		want := 99
		want0s := 0

		if got != want || numZeroes != want0s {
			t.Errorf("got to %d, accumulated %d. wanted %d and %d", got, numZeroes, want, want0s)
		}
	})
}

// run the numbers from day1 part2 example
func TestSampleCase(t *testing.T) {
	numZeroes := 0

	rotationsAndDestinations := []struct{ rotation, destination int }{
		{-68, 82},
		{-30, 52},
		{+48, 0},
		{-5, 95},
		{+60, 55},
		{-55, 0},
		{-1, 99},
		{-99, 0},
		{+14, 14},
		{-82, 32},
	}

	dialPosition := 50

	for _, rad := range rotationsAndDestinations {
		got := getDialPosition(dialPosition, rad.rotation, &numZeroes)
		want := rad.destination

		if got != want {
			t.Errorf("turned %d from %d got %d want %d", rad.rotation, dialPosition, got, want)
		}

		dialPosition = got

	}

	if numZeroes != 6 {
		t.Errorf("accumulated %d zeroes, should have been %d", numZeroes, 6)
	}
}
