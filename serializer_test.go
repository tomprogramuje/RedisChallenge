package main

import "testing"

func TestSerialize(t *testing.T) {
	t.Run("sends 'OK'", func(t *testing.T) {
		got := Serialize("OK")
		want := `+OK\r\n`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("sends 'hello world'", func(t *testing.T) {
		got := Serialize("hello world")
		want := `+hello world\r\n`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
