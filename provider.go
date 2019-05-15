package transform

import (
	"reflect"

	"github.com/kataras/iris"
)

type Provider interface {
	Set(string, interface{}) error
	Get(string) (interface{}, error)
	Fields() []string
}

func DetectProvider(v interface{}) Provider {
	if ctx, ok := v.(iris.Context); ok {
		return NewIrisContextProvider(ctx)
	}
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
