[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-interfaces/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-interfaces/blob/main/README.md)  
go version 1.22.1  

## Interfaces
As ```interface``` para representar "objetos", por exemplo caso tenhamos 2 ```structs``` um chamado ```circle``` e outro chamado ```rectangle```, caso eu queira calcular uma área, precisaria de duas funções?  
Veremos agora:  
```go
package main

type circle struct {
	rad float64
}
type rectangle struct {
	height float64
	width  float64
}

func writeArea(r rectangle) {

}

func main() {

}
```
Aqui, caso você queira calcular a area dos dois structs, não seria possivel, pois possuem tipos diferentes, e entradas diferentes, o memso vale pro contrário.  
Vejamos como solucionar isso com interfaces, parecido com o struct, mas ao inves de usar a palavra ```struct``` você usa a palavra chave ```interface```, lembrando que a interface só tem assinatura de métodos, você não pode passar campos, assim como passamos num ```struct```, você pode passara ssinaturas de métodos que diz como esses métodos devem ser, vejamos:  
```go
...
func writeArea(f from) {
	fmt.Printf("A área da forma é %0.2f", f.area())
}

type from interface {
  area() float64
}



func main() {

}
```
Agora vamos ver como passar os parâmetros para nossa função ```area()```:  
```go
...
func main() {
	r := rectangle{10,15}
	writeArea(r) // <<
  // irá reclamar que você não pode usar rectangle como tipo form, pois nao atendem os requisitos
  // para ser uma form. Tem que ter obrigatoriamente ter o método chamado área que retorna um float64
}
```
Para a função ```writeArea``` você precisa obrigatoriamente ter o metodo cuja assinatura é igual a ```form``` interface, ou seja, deve ser área, não pode ter parâmeros se não não é considerado uma ```form```, vejamos:  
```go
package main

import (
	"fmt"
	"math"
)

func (r rectangle) area() float64 {
	return r.height * r.width
}
func (c circle) area() float64 {
	return math.Pi * (c.rad * c.rad)
}

type circle struct {
	rad float64
}

type rectangle struct {
	height float64
	width  float64
}

type from interface {
	area() float64
}

func writeArea(f from) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

func main() {
	r := rectangle{10, 15}
	writeArea(r) // 150
	c := circle{10}
	writeArea(c) // 314.16
}
```
