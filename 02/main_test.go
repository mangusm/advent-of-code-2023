package main

import "testing"

func TestPartOne(t *testing.T) {
	want := 2176
	sum := PartOne()
	if want != sum {
		t.Fail()
	}
}

func TestPartTwo(t *testing.T) {
	want := 63700
	sum := PartTwo()
	if want != sum {
		t.Fail()
	}
}
