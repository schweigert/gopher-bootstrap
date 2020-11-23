package main

import f "fmt"

type Geometrico interface {
	Area() float64
}

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

	printa(quad)
	printa(tri)
}

func printa(g Geometrico) {
	f.Println("Área do quadrado:", g.Area())
}
