package transform

import (
	"reflect"
)

type Provider interface {
	Set(string, interface{}) error
	Get(string) (interface{}, error)
	Fields() []string
}

func DetectProvider(v interface{}) Provider {
	rv := Indirect(v)
	if rv.Type().Implements(reflect.TypeOf((*Provider)(nil)).Elem()) {
		return rv.Interface().(Provider)
	}
	if rv.Kind() == reflect.Map {
		return NewMapProvider(v)
	}
	if rv.Kind() == reflect.Struct {
		return NewStructProvider(v)
	}
	panic(ErrTypeNotMatch)
}
