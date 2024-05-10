[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-interfaces/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-interfaces/blob/main/README.md)  
go version 1.22.1  

## Interfaces
The ```interface``` to represent "objects", for example we have 2 ```structs``` one called ```circle``` and another called ```rectangle```, if I want to calculate a area, would I need two functions?
We will now see:  
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
Here, if you want to calculate the area of ​​two structs, it would not be possible, as they have different types, and different inputs, or the opposite is true.  
Let's see how to solve this with interfaces, similar to struct, but when you use the word ```struct``` you use the keyword ```interface```, remembering that the interface only has method signatures, you You can't pass fields, just as reported in a ```struct```, you can pass method signatures that say what these methods should be like, let's see:  
```go
...
func writeArea(f from) {
	fmt.Printf("The area of ​​the shape is %0.2f", f.area())
}

type from interface {
  area() float64
}



func main() {

}
```
Now let's see how to pass the parameters to our ```area()``` function:  
```go
...
func main() {
	r := rectangle{10,15}
	writeArea(r) // <<
  // will complain that you cannot use rectangle as a form type, as it does not meet the requirements
  // to be a form. There must be a method called area that returns a float64
}
```
For the function ```writeArea``` you must have the method whose signature is equal to ```form``` interface, that is, it must be an area, it cannot have parameters otherwise it is not considered a ``` form```, let's see:  
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
