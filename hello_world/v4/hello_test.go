package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Sangil")
	want := "Hello, Sangil"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
