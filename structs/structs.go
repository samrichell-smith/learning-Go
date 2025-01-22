package structs

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(height, width float64) float64 {
	return 2*height + 2*width
}

func Area(height, width float64) float64 {
	return height * width
}