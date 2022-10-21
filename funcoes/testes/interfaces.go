package testes

type Carro struct {
	NumRodas, KmRodado int
}

func (c *Carro) Viajou(km int) {
	c.KmRodado += km
}

type Barco struct {
	PotenciaMotor, KmRodado int
}

func (c *Barco) Viajou(km int) {
	c.KmRodado += km
}

type Transporte interface {
	Viajou(v int)
}

func Viajou(t Transporte, a int) {

	t.Viajou(a)
}
