package transform

import (
	"reflect"
)

type StructData struct {
	*DataBase
	Data  interface{}
	Type  reflect.Type
	Value reflect.Value
}

func (this *StructData) Set(key string, value interface{}) error {
	key = this.KeyIn(key)
	rv := reflect.ValueOf(value)
	if rv.Kind() != this.Value.FieldByName(key).Kind() {
		return ErrTypeNotMatched
	}
	this.Value.FieldByName(key).Set(rv)
	return nil
}

func (this *StructData) Get(key string) (interface{}, error) {
	key = this.KeyIn(key)
	return this.Value.FieldByName(key).Interface(), nil
}

func (this *StructData) KVs() []*KV {
	kvs := []*KV{}
	for i := 0; i < this.Value.NumField(); i++ {
		kvs = append(kvs, &KV{
			Key:   this.KeyOut(this.Type.Field(i).Name),
			Value: this.Value.Field(i).Interface(),
		})
	}
	return kvs
}

func Struct(v interface{}, options ...DataOption) *StructData {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Interface || rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	rt := rv.Type()
	d := &StructData{
		DataBase: NewDataBase(),
		Data:     v,
		Type:     rt,
		Value:    rv,
	}
	for _, option := range options {
		option(d)
	}
	return d
}
