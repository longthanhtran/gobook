// http://www.hydrogen18.com/blog/golang-embedding.html
package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Android struct {
	Person // embedded
	Model  string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	p := Person{"Long"}
	a := new(Android)
	a.Person.Talk()
	p.Talk()
}
