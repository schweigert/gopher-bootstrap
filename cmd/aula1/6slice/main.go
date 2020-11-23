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
	objs := []Geometrico{
		&Quadrado{A: 2},
		&Quadrado{A: 3},
		&Quadrado{A: 4},
		&Quadrado{A: 5},
	}

	for _, obj := range objs {
		printa(obj)
	}
}

func printa(g Geometrico) {
	f.Println("Área do quadrado:", g.Area())
}
