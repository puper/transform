package transform

type MapData struct {
	*DataBase
	Data map[string]interface{}
}

func (this *MapData) Set(key string, value interface{}) error {
	key = this.KeyIn(key)
	this.Data[key] = value
	return nil
}

func (this *MapData) Get(key string) (interface{}, error) {
	key = this.KeyIn(key)
	if v, ok := this.Data[key]; ok {
		return v, nil
	}
	return nil, ErrNotFound
}

func (this *MapData) KVs() []*KV {
	kvs := make([]*KV, 0, len(this.Data))
	for k, v := range this.Data {
		kvs = append(kvs, &KV{
			Key:   this.KeyOut(k),
			Value: v,
		})
	}
	return kvs
}

func Map(
	v map[string]interface{},
	keys []string,
	options ...DataOption,
) *MapData {
	d := &MapData{
		DataBase: NewDataBase(),
		Data:     v,
	}
	for _, key := range keys {
		d.Data[key] = nil
	}
	for _, option := range options {
		option(d)
	}
	return d
}

func Keys(keys ...string) []string {
	return keys
}
