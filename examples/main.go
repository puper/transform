package main

import (
	"log"
	"time"

	"github.com/puper/transform"
)

type A struct {
	A int64
	B string
	C int
	D string
	E string
}

type B struct {
	A  int
	B1 []string
	B2 []int
	C  time.Time
	D  int
}

func main() {
	a := &A{
		A: 1234,
		B: "1,2,3,4,5",
		C: 1234,
		D: "1234",
		E: "not use",
	}
	b := &B{}
	err := transform.Transform(
		a,
		b,
		transform.Only("A", "B"),
	)
	log.Println(b.A, err)
}
