package main

import "fmt"

type car struct {
	v int
	s int
}

type moto struct {
	v int
	s int
}

func (a moto) rasx() int {
	return a.v * a.s
}

func (a car) rasx() int {
	return a.v * a.s
}

type rasxer interface {
	rasx() int
}

func main() {
	var q, w rasxer
	q = car{1, 1}
	w = moto{10, 10}

	fmt.Println(q.rasx(), w.rasx())

}
