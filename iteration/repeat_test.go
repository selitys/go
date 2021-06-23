package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertCorrect := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("repeats with given count", func(t *testing.T) {
		result := Repeat("a", 5)
		expect := "aaaaa"
		assertCorrect(t, result, expect)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	rep := Repeat("a", 3)
	fmt.Println(rep)
	// Output: aaa
}
