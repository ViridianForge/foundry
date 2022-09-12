package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("expected %q, observed %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("test", 5)
	}
}

// Prints "bbbb"
func ExampleRepeat() {
	repeated := Repeat("b", 4)
	fmt.Println(repeated)
}
