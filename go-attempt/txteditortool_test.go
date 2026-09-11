package main

import "testing"

//Edited
func TestHexMultiple(t *testing.T) {
	input := "FF 1A 2B 3C (hex, 4)"
	expected := "255 26 43 60"

	result, _ := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestUp(t *testing.T) {
	input := "hello (up)"
	expected := "HELLO"

	result, _ := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestArticle(t *testing.T) {
	input := "a apple"
	expected := "an apple"

	result, _ := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestGroupedPunctuation(t *testing.T) {
	input := "What !? Wait ..."
	expected := "What!? Wait..."

	result, _ := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestUpMultiple(t *testing.T) {
	input := "This is so exciting (up, 2)"
	expected := "This is SO EXCITING"

	result, _ := ProcessText(input)

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

