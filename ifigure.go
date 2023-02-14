package figures

import "fmt"


type IFigure interface{
    getPerimeter() float64
    getArea() float64
}

func GetPerimeter(geometricFigure IFigure) {
    fmt.Printf("perimeter: %.2f cm\n", geometricFigure.getPerimeter())
}

func GetArea(geometricFigure IFigure) {
    fmt.Printf("area: %.2f cm^2\n", geometricFigure.getArea())
}
