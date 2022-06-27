package main

import "fmt"

func main(){
	c := NewCircle(4.5)
	a := NewRectangle(4, 5)
	fmt.Println(c.Type())
	fmt.Println(a.Area())
}


const PI = 3.141592653589793;

type Rectangle struct{
	figureType string
	a, b float64
}

func NewRectangle(a, b float64) *Rectangle {
    return &Rectangle{"Прямоугольник", a, b}
}

func (rect Rectangle) Area() float64 {
    return rect.a * rect.b
}

func (rect Rectangle) Type() string {
    return rect.figureType
}

type Circle struct{
	figureType string
	r float64
}

func NewCircle(r float64) *Circle {
    return &Circle{"Круг", r}
}

func (circ Circle) Area() float64 {
    return circ.r * circ.r * PI
}

func (circ Circle) Type() string {
    return circ.figureType
}