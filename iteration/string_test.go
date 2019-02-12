package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestTrimPrefix(t *testing.T) {
	got := strings.TrimPrefix("incomplete", "in")
	want := "complete"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func ExampleTrimPrefix() {
	expected := strings.TrimPrefix("impossible", "im")
	fmt.Println(expected)
	// Output: possible
}

func BenchmarkTrimPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.TrimPrefix("not good", "not ")
	}
}
