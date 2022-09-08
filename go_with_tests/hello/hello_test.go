package hello

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris!")
	want := "Hello, Chris!"

	if got != want {
		t.Errorf("got %q, expected %q", got, want)
	}
}
