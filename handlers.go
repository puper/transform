package transform

import (
	"strconv"
	"strings"
)

func KeyMapping(keys ...string) Option {
	var (
		aKey, bKey string
	)
	if len(keys) == 2 {
		aKey = keys[0]
		bKey = keys[1]
	} else {
		aKey = keys[0]
		bKey = keys[0]
	}
	return func(c *Config) {
		c.Handlers[bKey] = func(a Data, b Data) error {
			v, err := a.Get(aKey)
			if err != nil {
				return err
			}
			return b.Set(bKey, v)
		}
	}
}

func StringToInt(keys ...string) Option {
	var (
		aKey, bKey string
	)
	if len(keys) == 2 {
		aKey = keys[0]
		bKey = keys[1]
	} else {
		aKey = keys[0]
		bKey = keys[0]
	}
	return func(c *Config) {
		c.Handlers[bKey] = func(a Data, b Data) error {
			v, err := a.Get(aKey)
			if err != nil {
				return err
			}
			vs, ok := v.(string)
			if ok {
				vi, err := strconv.Atoi(vs)
				if err != nil {
					return err
				}
				return b.Set(bKey, vi)
			}
			return ErrTypeNotMatched
		}
	}
}

func StringToStringSlice(keys ...string) Option {
	var (
		aKey, bKey string
	)
	if len(keys) == 2 {
		aKey = keys[0]
		bKey = keys[1]
	} else {
		aKey = keys[0]
		bKey = keys[0]
	}
	return func(c *Config) {
		c.Handlers[bKey] = func(a Data, b Data) error {
			v, err := a.Get(aKey)
			if err != nil {
				return err
			}
			vs, ok := v.(string)
			if ok {
				return b.Set(bKey, strings.Split(vs, ","))
			}
			return ErrTypeNotMatched
		}
	}
}
