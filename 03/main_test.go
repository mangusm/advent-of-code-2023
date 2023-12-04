package main

import "testing"

func TestPartOne(t *testing.T) {
	want := 525181
	sum := PartOne()
	if want != sum {
		t.Fail()
	}
}

func TestPartTwo(t *testing.T) {
	want := 84289137
	sum := PartTwo()
	if want != sum {
		t.Fail()
	}
}
