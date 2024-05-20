package structs

type Dimensions struct {
	Width  int
	Height int
}

func NewDimensions(width, height int) Dimensions {
	return Dimensions{
		Width:  width,
		Height: height,
	}
}
