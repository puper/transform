# struct to map，map to struct， map to map， struct to struct
# parse iris.Context to custom request struct
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
# support iris context
```
package main

import (
	"github.com/kataras/iris"
	"github.com/puper/transform"
)

func main() {
	app := iris.Default()
	app.Post("/", func(ctx iris.Context) {
		req := &struct {
			A int
			B string
			C string
		}{}
		err := transform.New().CustomMapper(
			map[string]string{
				"A": "post.a",
				"B": "query.b",
				"C": "query.d",
			},
		).
			FieldHandler(transform.StringToInt, "A").
			IgnoreErrors(transform.ErrTypeNotMatch).
			Process(ctx, req)
		ctx.JSON(iris.Map{
			"err":  err,
			"data": req,
		})
	})
	app.Run(iris.Addr(":8080"))
}
```