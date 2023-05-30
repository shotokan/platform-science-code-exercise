package counter

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCounter_Consonants(t *testing.T) {
	t.Run("Consonants is 0 for string without consontants", func(t *testing.T) {
		t.Parallel()

		name := "Ai"
		want := 0

		got := Consonants(name)

		if want != got {
			t.Error("wrong number of consontants found",
				cmp.Diff(want, got))
		}

	})

	t.Run("Consonants is 0 for an empty string", func(t *testing.T) {
		t.Parallel()

		name := ""
		want := 0

		got := Consonants(name)

		if want != got {
			t.Error("wrong number of consontants found",
				cmp.Diff(want, got))
		}

	})

	t.Run("Consonants returns the number of consonants found", func(t *testing.T) {
		t.Parallel()

		name := "Cyndy"
		want := 5

		got := Consonants(name)

		if want != got {
			t.Error("wrong number of consontants found",
				cmp.Diff(want, got))
		}

	})
}
