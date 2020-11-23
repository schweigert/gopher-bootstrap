package main

import f "fmt"

type Quadrado struct {
	A float64
}

func (obj *Quadrado) Area() float64 {
	return obj.A * obj.A
}

type Triangulo struct {
	A float64
}

func (obj *Triangulo) Area() float64 {
	return (obj.A * obj.A) / 2
}

func main() {
	quad := &Quadrado{A: 5.0}
	tri := &Triangulo{A: 5.0}

	f.Println("Área do quadrado:", quad.Area())
	f.Println("Área do triangulo:", tri.Area())
}
