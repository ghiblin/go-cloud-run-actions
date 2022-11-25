package main

import "testing"

func TestHelloWo(t *testing.T) {
	var expected Phrase

	expected.Text = "Hello, world"
	result := HelloWorld()

	if expected.Text != result.Text {
		t.Errorf("Phrase was incorrect. Got: %s, want: %s.", result.Text, expected.Text)
	}
}
