package transform

import "reflect"

type StructProvider struct {
	rv reflect.Value
	rt reflect.Type
}

func NewStructProvider(v interface{}) *StructProvider {
	rv := Indirect(v)
	if rv.Kind() != reflect.Struct {
		panic(ErrTypeNotMatch)
	}
	return &StructProvider{
		rv: rv,
		rt: rv.Type(),
	}
}

func (this *StructProvider) Set(f string, v interface{}) error {
	rv := reflect.ValueOf(v)
	rfv := this.rv.FieldByName(f)
	if rfv.Kind() == reflect.Invalid {
		return ErrFieldNotFound
	}
	if !rv.Type().AssignableTo(rfv.Type()) {
		return ErrTypeNotMatch
	}
	rfv.Set(rv)
	return nil
}

func (this *StructProvider) Get(f string) (interface{}, error) {
	rfv := this.rv.FieldByName(f)
	if rfv.Kind() == reflect.Invalid {
		return nil, ErrFieldNotFound
	}
	return rfv.Interface(), nil
}

func (this *StructProvider) Fields() []string {
	result := make([]string, 0, this.rt.NumField())
	for i := 0; i < this.rt.NumField(); i++ {
		result = append(result, this.rt.Field(i).Name)
	}
	return result
}
