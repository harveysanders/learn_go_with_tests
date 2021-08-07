package main

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rect Rectangle) float64 {
	return (rect.Width * 2) + (rect.Height * 2)
}

func Area(rect Rectangle) float64 {
	return rect.Width * rect.Height
}
