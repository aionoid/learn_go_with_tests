// @Title Add func
// @Description
// @Author Belaoura Yacoub
// @Update
package integers

import (
	"fmt"
	"testing"
)

// Add takes two integers and returns the sum of theme.
func TestAdd(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("got %d, and want %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
