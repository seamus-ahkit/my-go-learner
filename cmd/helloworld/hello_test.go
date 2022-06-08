package main

import "testing"

func TestHello(t *testing.T) {
	strVal := "text"

	got := Hello(strVal)
	want := "Hello, " + strVal

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
