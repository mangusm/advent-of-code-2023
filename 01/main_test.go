package main

import "testing"

func TestPartOne(t *testing.T) {
	want := 54916
	sum := PartOne()
	if want != sum {
		t.Fail()
	}
}

func TestPartTwo(t *testing.T) {
	want := 54728
	sum := PartTwo()
	if want != sum {
		t.Fail()
	}
}
