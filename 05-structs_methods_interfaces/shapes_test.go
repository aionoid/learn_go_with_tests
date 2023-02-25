// @Title
// @Description
// @Author
// @Update
package structsmethodsinterfaces

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g for shape: %#v", got, want, shape)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Triangle{12, 6}, 36.0},
		{Circle{10.0}, 314.1592653589793},
	}
	for _, tt := range areaTests {
		t.Run(fmt.Sprintf("%T", tt.shape), func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}

	// t.Run("rectangles", func(t *testing.T) {
	//   rect := Rectangle{12.0, 6.0}
	//   checkArea(t, rect, 72.0)
	// })
	//
	// t.Run("circles", func(t *testing.T) {
	//   circle := Circle{10.0}
	//   checkArea(t, circle, 314.1592653589793)
	// })
}
