package transform

import "reflect"

func Indirect(v interface{}) reflect.Value {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Interface || rv.Kind() == reflect.Ptr {
		return rv.Elem()
	}
	return rv
}
