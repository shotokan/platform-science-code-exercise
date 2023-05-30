package commonfactor

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCounter_HaveCommonFactor(t *testing.T) {
	t.Run("HaveCommonFactor is true when a common factor gre", func(t *testing.T) {
		t.Parallel()
		a, b := 15, 10

		want := true

		got := HaveCommonFactors(a, b)

		if want != got {
			t.Error("the numbers have no common factor",
				cmp.Diff(want, got))
		}

	})

	t.Run("HaveCommonFactor have a common factor other than one", func(t *testing.T) {
		t.Parallel()
		a, b := 11, 10

		want := false

		got := HaveCommonFactors(a, b)

		if want != got {
			t.Errorf("the numbers %d, %d have no common factor other than one", a, b)
		}
	})
}
