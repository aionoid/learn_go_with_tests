// @Title
// @Description
// @Author
// @Update
package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, and want %d given, %v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum of 2 slices", func(t *testing.T) {
		got := SumAll([]int{1, 3}, []int{0, 4})
		want := []int{4, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, and want %v", got, want)
		}
	})
}

func TestSumAll2(t *testing.T) {
	t.Run("Sum of 2 slices", func(t *testing.T) {
		got := SumAll2([]int{1, 3}, []int{0, 4})
		want := []int{4, 4}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, and want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, and want %v", got, want)
		}

	}
	t.Run("Sum of tails for 2 slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 3, 4}, []int{1, 4, 2})
		want := []int{7, 6}
		checkSums(t, got, want)
		// if !reflect.DeepEqual(got, want) {
		//   t.Errorf("got %v, and want %v", got, want)
		// }
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 4, 2})
		want := []int{0, 6}

		checkSums(t, got, want)
		// if !reflect.DeepEqual(got, want) {
		//   t.Errorf("got %v, and want %v", got, want)
		// }
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 3}, []int{0, 4})
	}
}
func BenchmarkSumAll2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll2([]int{1, 3}, []int{0, 4})
	}
}
