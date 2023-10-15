package shapes


type Shape interface {
    Area() float64
    Perimeter() float64
	GetName() string
}

type shape struct {
    name string
}
//shape имплементирует методы интерфейса Shape
func (s shape) Area() float64 {
    return 0
}

func (s shape) Perimeter() float64 {
    return 0
}

func (s shape) GetName() string {
    return s.name
}
