package transform

import (
	"log"
	"testing"
)

type A struct {
	A       string
	C       string
	TestKey string
}

type B struct {
	B int
	C []string
	D string
}

func TestTransformStructStruct(t *testing.T) {
	a := &A{
		A: "1",
		C: "a,b,c,d",
	}
	b := &B{}
	err := Transform(
		Struct(a),
		Struct(b),
		StringToInt("A", "B"),
		StringToStringSlice("C", "C"),
		KeyMapping("A", "D"),
	)
	if err != nil {
		t.Error(err)
		return
	}
	if b.B != 1 {
		t.Errorf("b.B not eq 1")
	}
}

func TestTransformStructMap(t *testing.T) {
	a := &A{
		A:       "1",
		TestKey: "12345",
	}
	b := map[string]interface{}{}
	err := Transform(
		Struct(a),
		Map(
			b,
			Keys(
				"b",
				"testKey",
			),
			KeyConvertor(
				SnakeString,
				CamelString,
			),
		),
		StringToInt("A", "B"),
		StringToStringSlice("C", "C"),
	)
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(b["test_key"])
	if b["b"] != 1 {
		t.Errorf("b.a not eq 1")
	}
}
