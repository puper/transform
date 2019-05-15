# struct to map，map to struct， map to map， struct to struct
# example:
```
package main

import (
	"log"

	"github.com/puper/transform"
)

type Test struct {
	A int
	B string
}

type Test2 struct {
	A []int
	B int
}

func main() {
	a := &Test{
		A: 1,
		B: "1,2,3",
	}
	b := &Test2{}
	err := transform.New().
		FieldMapper(transform.LowerFirst, transform.UpperFirst).
		CustomMapper(map[string]string{
			"B": "A",
			"A": "B",
		}).
		FieldHandler(transform.StringToIntSlice, "A").
		Process(a, b)
	log.Println(err, b.A)
}

```