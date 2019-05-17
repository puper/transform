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
		if ai, err := strconv.Atoi(av); err == nil {
			return ai, nil
		} else {
			return 0, ErrConvertFailed
		}
	}
	return 0, ErrTypeNotMatch
}

func StringToStringSlice(a interface{}) (interface{}, error) {
	if av, ok := a.(string); ok {
		return strings.Split(av, ","), nil
	}
	return []string{}, ErrTypeNotMatch
}

func StringToIntSlice(a interface{}) (interface{}, error) {
	bv := []int{}
	a, err := StringToStringSlice(a)
	if err != nil {
		return bv, err
	}
	for _, row := range a.([]string) {
		rowInt, err := strconv.Atoi(row)
		if err != nil {
			return bv, ErrConvertFailed
		}
		bv = append(bv, rowInt)
	}
	return bv, nil
}

func Int64ToInt(a interface{}) (interface{}, error) {
	if aInt64, ok := a.(int64); ok {
		return int(aInt64), nil
	}
	return 0, ErrTypeNotMatch
}
