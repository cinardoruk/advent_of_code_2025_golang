package main

import (
	"slices"
	"testing"
)

/*
	"some sequence of digits repeated twice"
>Your job is to find all of the invalid IDs that appear in the given ranges. In the above example:

11-22 has two invalid IDs, 11 and 22.
95-115 has one invalid ID, 99.
998-1012 has one invalid ID, 1010.
1188511880-1188511890 has one invalid ID, 1188511885.
222220-222224 has one invalid ID, 222222.
1698522-1698528 contains no invalid IDs.
446443-446449 has one invalid ID, 446446.
38593856-38593862 has one invalid ID, 38593859.
The rest of the ranges contain no invalid IDs.
*/

func TestIdValidator(t *testing.T) {
	t.Run("detect 55 as invalid", func(t *testing.T) {
		got := validateId("55")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("detect 56 as invalid", func(t *testing.T) {
		got := validateId("56")
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("detect 1010 as invalid", func(t *testing.T) {
		got := validateId("1010")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("detect 1188511885 as invalid", func(t *testing.T) {
		got := validateId("1188511885")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("detect 222222 as invalid", func(t *testing.T) {
		got := validateId("222222")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("detect 446446 as invalid", func(t *testing.T) {
		got := validateId("446446")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("detect 38593859 as invalid", func(t *testing.T) {
		got := validateId("38593859")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}

func TestInputParser(t *testing.T) {
	got := parseInput("test_input.txt")
	want := []Interval{
		{1, 17},
		{27, 86},
	}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIntervalValidator(t *testing.T) {
	testInterval := parseInput("test_input.txt")[0]
	got := validateInterval(testInterval)
	want := []int{11}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
