package main

import "testing"

func TestSerialize(t *testing.T) {
	got := Serialize("OK")
	want := `+OK\r\n`

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
