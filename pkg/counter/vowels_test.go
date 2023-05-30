package counter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCounter_Vowels(t *testing.T) {
	t.Run("Vowels is 0 for string without vowels", func(t *testing.T) {
		t.Parallel()

		name := "Cyndy"
		want := 0

		got := Vowels(name)

		if want != got {
			t.Error("wrong number of vowels found",
				cmp.Diff(want, got))
		}

	})

	t.Run("Vowels is 0 for an empty string", func(t *testing.T) {
		t.Parallel()

		name := ""
		want := 0

		got := Vowels(name)

		if want != got {
			t.Error("wrong number of vowels found",
				cmp.Diff(want, got))
		}

	})

	t.Run("Vowels returns the number of vowels found", func(t *testing.T) {
		t.Parallel()

		name := "Ivan"
		want := 2

		got := Vowels(name)

		if want != got {
			t.Error("wrong number of vowels found",
				cmp.Diff(want, got))
		}

	})
}
