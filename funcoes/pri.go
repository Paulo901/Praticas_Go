package main

import (
	"fmt"
	"funcoes/exercicios"
	"funcoes/testes"
)

func main() {

	c := exercicios.Circulo{VRaio: 10}
	q := exercicios.Quadrado{VLado: 20}

	exercicios.Area(&c)
	exercicios.Area(&q)

	//Guardado()
}

func Guardado() {

	// testando métodos
	a := testes.Valor{Num: 1}
	a.Soma(20)
	fmt.Println(a)
	// ---------------------------------
	// testando interfaces
	b := testes.Carro{NumRodas: 4, KmRodado: 60}
	c := testes.Barco{PotenciaMotor: 40, KmRodado: 1000}

	testes.Viajou(&c, 20)
	testes.Viajou(&b, 100)

	fmt.Println(b, c)
	//-----------------------------------
	//testando retornar funções
	fmt.Println(testes.Ola("Paulo")("noite"))
	//testando callback
	valor := testes.SoImpar(testes.Soma1, 10, 11, 12, 13, 14, 15)
	fmt.Println(valor)

}
