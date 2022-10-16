package main

import (
	"fmt"
	"math"
)

// * Interfaces
func main() {
	retan := retangulo{10, 15}
	escreverArea(retan)

	cir := circulo{10}
	escreverArea(cir)
}

// * Criando interface
// ? interface apenas tem assinaturas de métodos
// ? a interface é implementada de forma implícita
type forma interface {
	// Criando método que retorna float64
	area() float64
}

// ? passando a interface como parâmetro
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

// ? tem que ter o método area
type retangulo struct {
	altura  float64
	largura float64
}

// ? criação do método área vinculado ao retangulo
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}
