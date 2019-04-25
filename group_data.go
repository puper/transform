package transform

import "strings"

// GroupData Only for source.
type GroupData struct {
	*DataBase
	Data map[string]Data
}

func (this *GroupData) Get(key string) (interface{}, error) {
	keys := strings.Split(key, ".")
	if len(keys) == 2 {
		if d, ok := this.Data[keys[0]]; ok {
			return d.Get(keys[1])
		}
		return nil, ErrNotFound
	}
	return nil, ErrTypeNotMatched
}

func Group(data map[string]Data) *GroupData {
	return &GroupData{
		DataBase: NewDataBase(),
		Data:     data,
	}
}
