package algs

import "testing"

func TestSearchInts(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	got := BinarySearch(slice, 3)
	want := 2
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}

	got = BinarySearch(slice, -5)
	want = -1
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestSearchNegatives(t *testing.T) {
	array := [10]int{-10, -5, -3, -1, 0, 1, 3, 5, 10}

	got := BinarySearch(array[:], -3)
	want := 2
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}

	got = BinarySearch(array[:], -100)
	want = -1
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}

func TestSearchStrings(t *testing.T) {
	slice := []string{"A", "B", "C", "D", "E", "F", "G"}

	got := BinarySearch(slice, "B")
	want := 1
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
