package main

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{12.0, 6.0}, hasArea: 72},
		{shape: Circle{10}, hasArea: 314.1592653589793},
		{shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if got != tt.hasArea {
			t.Errorf("%#v \ngot %g want %g", tt, got, tt.hasArea)
		}
	}
}
