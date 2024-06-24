package perimeter

type Rectangle struct {
	Width  float32
	Height float32
}

func Perimeter(rectangle Rectangle) float32 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float32 {
	return rectangle.Width * rectangle.Height
}
