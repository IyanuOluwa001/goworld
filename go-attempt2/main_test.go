package main

import (
	"strings"
	"testing"
)

// Test applySingleModifier

func TestApplySingleModifier(t *testing.T) {

	tests := []struct {
		word     string
		modifier string
		expected string
	}{
		{"1E", "hex", "30"},
		{"1010", "bin", "10"},
		{"hello", "up", "HELLO"},
		{"HELLO", "low", "hello"},
		{"bridge", "cap", "Bridge"},
	}

	for _, test := range tests {
		result := applySingleModifier(test.word, test.modifier)
		if result != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, result)
		}
	}
}

// Test capitalize
func TestCapitalize(t *testing.T) {
	result := capitalize("gOLang")
	expected := "Golang"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}


// Test applyModifiers (including count)
func TestApplyModifiers(t *testing.T) {

	input := "one two three four (up, 2)"
	expected := "one two THREE FOUR"

	result := applyModifiers(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

//Done
func TestHexMultiple(t *testing.T) {

	input := "FF 1A 2B 3C (hex, 4)"
	expected := "255 26 43 60"

	result := applyModifiers(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestBinMultiple(t *testing.T) {

	input := "10 11 100 101 (bin, 3)"
	expected := "10 3 4 5"

	result := applyModifiers(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}


// Test fixPunctuation
func TestFixPunctuation(t *testing.T) {

	input := "Hello , world !"
	expected := "Hello, world!"

	result := fixPunctuation(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// Test fixQuotes
func TestFixQuotes(t *testing.T) {

	input := "He said ' hello world ' loudly."
	expected := "He said 'hello world' loudly."

	result := fixQuotes(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// Test fixAAn
func TestFixAAn(t *testing.T) {

	input := "a apple and A orange but a banana"
	expected := "an apple and An orange but a banana"

	result := fixAAn(input)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

// Integration Test (Full Pipeline)
func TestFullPipeline(t *testing.T) {

	input := "1E (hex) is a apple and hello (up) world , wow !"

	result := applyModifiers(input)
	result = fixPunctuation(result)
	result = fixAAn(result)

	expectedContains := []string{
		"30",
		"an apple",
		"HELLO",
		"world, wow!",
	}

	for _, part := range expectedContains {
		if !strings.Contains(result, part) {
			t.Errorf("Expected result to contain %s, got %s", part, result)
		}
	}
}