// @Title
// @Description
// @Author
// @Update
package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 6)
	want := "aaaaaa"
	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("b", 5)
	}
}

func ExampleRepeat() {
	out := Repeat("k", 6)
	fmt.Println(out)
	//Output: kkkkkk
}
