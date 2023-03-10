package structsmethodsinterfaces

import "math"

// @Title
// @Description
// @Author
// @Update
type Shape interface {
	Area() float64
}
type Rectangle struct {
	width  float64
	height float64
}
type Triangle struct {
	Base   float64
	Height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.width + r.height) * 2
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func (c Circle) Area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
