package main

import "testing"

// Test uppercase modifier
func TestUpModifier(t *testing.T) {

	input := "hello (up)"
	expected := "HELLO"

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

// Test uppercase modifier with count
func TestUpMultiple(t *testing.T) {

	input := "This is so exciting (up, 2)"
	expected := "This is SO EXCITING"

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

// Test hex conversion modifier
func TestHexMultiple(t *testing.T) {

	input := "FF 1A 2B 3C (hex, 4)"
	expected := "255 26 43 60"

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

// Test grammar rule for article correction
func TestArticleCorrection(t *testing.T) {

	input := "a apple"
	expected := "an apple"

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

// Test punctuation grouping
func TestGroupedPunctuation(t *testing.T) {

	input := "What !? Wait ..."
	expected := "What!? Wait..."

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

// Test binary conversion
func TestBinaryConversion(t *testing.T) {

	input := "1010 (bin)"
	expected := "10"

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

// Test capitalization modifier
func TestCapModifier(t *testing.T) {

	input := "hELLo (cap)"
	expected := "Hello"

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

// Test lowercase modifier
func TestLowModifier(t *testing.T) {

	input := "HELLO (low)"
	expected := "hello"

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}

// NEW TESTS FOR QUOTES AND PUNCTUATION

// Test for single word inside quotes
func TestQuotesSingleWord(t *testing.T) {
	input := "' hello '"
	expected := "'hello'"

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test for multiple words inside quotes
func TestQuotesMultipleWords(t *testing.T) {
	input := "' I am learning Go '"
	expected := "'I am learning Go'"

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test for punctuation next to quotes
func TestQuotesWithPunctuation(t *testing.T) {
	input := "He said : ' Hello world ' , and left ."
	expected := "He said: 'Hello world', and left."

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test for Input1 example from problem
func TestInput1Example(t *testing.T) {
	input := "As the poet once said: ' imagination is more important than knowledge ' , and I believe it ... truly !! Do you think so ? Or is it just ' a dream ' ?"
	expected := "As the poet once said: 'imagination is more important than knowledge', and I believe it... truly!! Do you think so? Or is it just 'a dream'?"

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// Test for Input2 example from problem
func TestInput2Example(t *testing.T) {
	input := "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '"
	expected := "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"

	result, err := ProcessText(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}