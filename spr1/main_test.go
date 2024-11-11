package main

import "testing"

func TestSpamMask(t *testing.T) {
	result := SpamMask("http://serdar.com")
	expected := "http://**********"

	if result != expected {
		t.Errorf("Expected: %s but got: %s", expected, result)
	}
}
