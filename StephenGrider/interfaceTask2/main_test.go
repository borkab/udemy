package main

import (
	"testing"
)

func TestPrintFileName(t *testing.T) {

	got := printFileName(file{"mondo.txt"})
	want := "mondo.txt"

	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}
