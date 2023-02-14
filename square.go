package figures


type Square struct{
    Width int
    Height int
}

func (square *Square) getPerimeter() float64 {
    return float64((2 * square.Width) + (2 * square.Height))
}

func (square *Square) getArea() float64 {
    return float64(square.Width * square.Height)
}
