package transform

import (
	"reflect"
)

type MapProvider struct {
	rv reflect.Value
	rt reflect.Type
}

// must be map[string]xxxx
func NewMapProvider(v interface{}) *MapProvider {
	rv := Indirect(v)
	if rv.Kind() != reflect.Map {
		panic(ErrTypeNotMatch)
	}
	if rv.Type().Key().Kind() != reflect.String {
		panic(ErrTypeNotMatch)
	}
	return &MapProvider{
		rv: rv,
		rt: rv.Type(),
	}
}

func (this *MapProvider) Set(f string, v interface{}) error {
	rv := reflect.ValueOf(v)
	if !rv.Type().AssignableTo(this.rt.Elem()) {
		return ErrTypeNotMatch
	}
	this.rv.SetMapIndex(reflect.ValueOf(f), rv)
	return nil
}

func (this *MapProvider) Get(f string) (interface{}, error) {
	rv := this.rv.MapIndex(reflect.ValueOf(f))
	if rv.Kind() == reflect.Invalid {
		return nil, ErrFieldNotFound
	}
	return rv.Interface(), nil
}

func (this *MapProvider) Fields() []string {
	result := make([]string, 0, this.rv.Len())
	for _, v := range this.rv.MapKeys() {
		result = append(result, v.String())
	}
	return result
}
