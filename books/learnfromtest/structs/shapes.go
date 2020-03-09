package structs

import (
	"fmt"
	"math"
	"strconv"
)

const PI = 3.14

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Length) * 2
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * PI
}

func (c Circle) Perimeter() float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", c.Radius*2*PI), 64)
	return value
}

type Triangle struct {
	Length float64
	Width  float64
	Height float64
}

func (t Triangle) Perimeter() float64 {
	return t.Length + t.Width + t.Height
}

func (t Triangle) Area() float64 {
	sum := t.Perimeter() / 2
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", math.Sqrt(sum*(sum-t.Length)*(sum-t.Width)*(sum-t.Height))), 64)
	return value
}
