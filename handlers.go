package transform

import (
	"strconv"
	"strings"
)

type FieldHandler func(interface{}) (interface{}, error)

func StringTrimSpace(a interface{}) (interface{}, error) {
	if av, ok := a.(string); ok {
		return strings.TrimSpace(av), nil
	}
	return a, ErrTypeNotMatch
}

func StringToInt(a interface{}) (interface{}, error) {
	if av, ok := a.(string); ok {
		return strconv.Atoi(av)
	}
	return a, ErrTypeNotMatch
}

func StringToStringSlice(a interface{}) (interface{}, error) {
	if av, ok := a.(string); ok {
		return strings.Split(av, ","), nil
	}
	return a, ErrTypeNotMatch
}

func StringToIntSlice(a interface{}) (interface{}, error) {
	a, err := StringToStringSlice(a)
	if err != nil {
		return a, err
	}
	bv := []int{}
	for _, row := range a.([]string) {
		rowInt, err := strconv.Atoi(row)
		if err != nil {
			return a, ErrConvertFailed
		}
		bv = append(bv, rowInt)
	}
	return bv, nil
}
