package transform

import (
	"github.com/puper/orderedmap"
)

type OrderedMapProvider struct {
	data *orderedmap.OrderedMap
}

func NewOrderedMapProvider(data *orderedmap.OrderedMap) *OrderedMapProvider {
	return &OrderedMapProvider{
		data: data,
	}
}

func (this *OrderedMapProvider) Set(f string, v interface{}) error {
	this.data.Set(f, v)
	return nil
}

func (this *OrderedMapProvider) Get(f string) (interface{}, error) {
	if v, ok := this.data.Get(f); ok {
		return v, nil
	}
	return nil, ErrFieldNotFound
}

func (this *OrderedMapProvider) Fields() []string {
	return this.data.Keys()
}
