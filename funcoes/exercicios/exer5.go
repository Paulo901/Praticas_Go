package exercicios

import "fmt"

type Quadrado struct {
	VLado, VArea int
}

func (q *Quadrado) Area() {
	q.VArea = q.VLado * 4
	fmt.Println(q.VArea)
}

type Circulo struct {
	VRaio, VArea int
}

func (c *Circulo) Area() {
	c.VArea = 2 * 3 * c.VRaio
	fmt.Println(c.VArea)
}

type Figura interface {
	Area()
}

func Area(f Figura) {
	f.Area()
}
