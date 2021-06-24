package main

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	peri := Perimeter(rect)
	want := 40.0

	if peri != want {
		t.Errorf("got %.2f want %.2f", peri, want)
	}
}

func TestArea(t *testing.T) {

	// The only new syntax here is creating an "anonymous struct", areaTests.
	// We are declaring a slice of structs by using []struct with two fields,
	// the shape and the want. Then we fill the slice with cases.

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}

}
