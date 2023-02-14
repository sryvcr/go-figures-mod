package figures

import "math"


type Circle struct{
    Radio int
}

func (circle *Circle) getPerimeter() float64 {
    return 2 * math.Pi * float64(circle.Radio)
}

func (circle *Circle) getArea() float64 {
    return math.Pi * float64(circle.Radio * circle.Radio)
}
