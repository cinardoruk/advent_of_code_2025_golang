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

/*
Now, an ID is invalid if it is made only of some sequence of digits repeated at least twice. So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times), and 1111111 (1 seven times) are all invalid IDs.

From the same example as before:

11-22 still has two invalid IDs, 11 and 22.
95-115 now has two invalid IDs, 99 and 111.
998-1012 now has two invalid IDs, 999 and 1010.
1188511880-1188511890 still has one invalid ID, 1188511885.
222220-222224 still has one invalid ID, 222222.
1698522-1698528 still contains no invalid IDs.
446443-446449 still has one invalid ID, 446446.
38593856-38593862 still has one invalid ID, 38593859.
565653-565659 now has one invalid ID, 565656.
824824821-824824827 now has one invalid ID, 824824824.
2121212118-2121212124 now has one invalid ID, 2121212121.
Adding up all the invalid IDs in this example produces 4174379265.

*/

// Now, an ID is invalid if it is made only of some sequence of digits repeated at least twice. So, 12341234 (1234 two times), 123123123 (123 three times), 1212121212 (12 five times), and 1111111 (1 seven times) are all invalid IDs.
func TestIdValidator2(t *testing.T) {
	t.Run("detect 998 as valid", func(t *testing.T) {
		got := validateId2("998")
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("detect 999 as invalid", func(t *testing.T) {
		got := validateId2("999")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("detect 12341234 as invalid", func(t *testing.T) {
		got := validateId2("12341234")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("detect 123123123 as invalid", func(t *testing.T) {
		got := validateId2("123123123")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("detect 1212121212 as invalid", func(t *testing.T) {
		got := validateId2("1212121212")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
	t.Run("detect 1111111 as invalid", func(t *testing.T) {
		got := validateId2("1111111")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
