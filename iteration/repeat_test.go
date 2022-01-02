package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}
	t.Run("std example a*5", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("repeat(a, b)", func(t *testing.T) {
		repeated := Repeat("a")
		expected := strings.Repeat("a", 5)
		assertCorrectMessage(t, repeated, expected)
	})
}

func ExampleRepeat() {
	repeats := Repeat("b")
	fmt.Println(repeats)
	// Output: bbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
