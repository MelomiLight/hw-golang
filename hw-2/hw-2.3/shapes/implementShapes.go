package shapes

import "math"



/*
Полиморфизм достигается через интерфейс Shape. Все структуры, которые реализуют этот интерфейс, 
предоставляют свои собственные реализации методов Area() и Perimeter().
*/

type Circle struct {
    shape // включают shape как поле, это КОМПОЗИЦИЯ 
    Radius float64
}

func NewCircle(radius float64) *Circle {
    return &Circle{
        shape:  shape{name: "Circle"},
        Radius: radius,
    }
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}


type Rectangle struct {
    shape // включают shape как поле, это КОМПОЗИЦИЯ 
    Width  float64
    Height float64
}

func NewRectangle(width, height float64) *Rectangle {
    return &Rectangle{
        shape:  shape{name: "Rectangle"},
        Width:  width,
        Height: height,
    }
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2*r.Width + 2*r.Height
}
