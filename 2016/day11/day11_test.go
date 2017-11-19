package main

import "testing"

func TestFloor_burned(t *testing.T) {
	f := floor{}
	f.add(object{otype: CHIP, element: RUTHENIUM})

	if f.burned() {
		t.Fatalf("Floor should not be burned with only one chip.")
	}

	f.add(object{otype: GEN, element: POLONIUM})
	if !f.burned() {
		t.Fatalf("Floor should be burned with one chip and non-matching generator.")
	}

	f.add(object{otype: GEN, element: RUTHENIUM})
	if f.burned() {
		t.Fatalf("Floor should not be burned with the chip connected to its generator.")
	}
}

func TestFloor_AddRemoveContainsEmpty(t *testing.T) {
	f := floor{}

	if !f.empty() {
		t.Fatalf("Floor should be empty")
	}

	f.add(object{otype: CHIP, element: RUTHENIUM})
	if f.empty() {
		t.Fatalf("Floor should not be empty")
	}
	if !f.containsChip(RUTHENIUM) {
		t.Fatalf("Floor should contain ruthenium chip")
	}
	if f.containsGenerator(RUTHENIUM) {
		t.Fatalf("Floor should not contain ruthenium generator")
	}

	f.remove(object{otype: CHIP, element: RUTHENIUM})
	if !f.empty() {
		t.Fatalf("Floor should be empty")
	}
	if f.containsChip(RUTHENIUM) {
		t.Fatalf("Floor should not contain ruthenium chip")
	}
	if f.containsGenerator(RUTHENIUM) {
		t.Fatalf("Floor should not contain ruthenium generator")
	}

	// Same tests, but with a generator
	f.add(object{otype: GEN, element: RUTHENIUM})
	if f.empty() {
		t.Fatalf("Floor should not be empty")
	}
	if f.containsChip(RUTHENIUM) {
		t.Fatalf("Floor should not contain ruthenium chip")
	}
	if !f.containsGenerator(RUTHENIUM) {
		t.Fatalf("Floor should contain ruthenium generator")
	}

	f.remove(object{otype: GEN, element: RUTHENIUM})
	if !f.empty() {
		t.Fatalf("Floor should be empty")
	}
	if f.containsChip(RUTHENIUM) {
		t.Fatalf("Floor should not contain ruthenium chip")
	}
	if f.containsGenerator(RUTHENIUM) {
		t.Fatalf("Floor should not contain ruthenium generator")
	}
}

func TestFacility_burned(t *testing.T) {
	///
}
