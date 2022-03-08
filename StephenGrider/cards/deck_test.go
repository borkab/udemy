package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 { //16 because we have 4 elemnts in suits and 4 in value
		t.Errorf("Expected deck length of 20, but got", len(d))
	}
}
