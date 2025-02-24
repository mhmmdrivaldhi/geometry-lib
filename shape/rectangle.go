package shape

type Rectangle struct {
	Width float32
	Lenght float32
}

func(rectangle *Rectangle) Area() float32 {
	return rectangle.Width * rectangle.Lenght
}

func(rectangle *Rectangle) Perimeter() float32 {
	return 2 * rectangle.Width + rectangle.Lenght 
}