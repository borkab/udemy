package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 { //16 because we have 4 elemnts in suits and 4 in value
		t.Errorf("Expected deck length of 20, but got", len(d))
	}
}

/*
go test
# learngo/udemy/StephenGrider/cards
./deck_test.go:9:11: Errorf call has arguments but no formatting directives
FAIL    learngo/udemy/StephenGrider/cards [build failed]
*/
