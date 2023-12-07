package main

import "testing"

func TestPartOne(t *testing.T) {
	want := 57075758
	sum := PartOne()
	if want != sum {
		t.Fail()
	}
}

func TestPartTwo(t *testing.T) {
	want := 8570000
	sum := PartTwo()
	if want != sum {
		t.Fail()
	}
}
