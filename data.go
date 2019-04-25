package transform

type Data interface {
	Set(string, interface{}) error
	Get(string) (interface{}, error)
	KVs() []*KV
	SetKeyConvertor(keyIn, keyOut func(string) string)
}

type DataBase struct {
	KeyIn  func(string) string
	KeyOut func(string) string
}

func NewDataBase() *DataBase {
	return &DataBase{
		KeyIn:  NoChange,
		KeyOut: NoChange,
	}
}

func (this *DataBase) SetKeyConvertor(keyIn, keyOut func(string) string) {
	this.KeyIn = keyIn
	this.KeyOut = keyOut
}

func (this *DataBase) Set(key string, value interface{}) error {
	return nil
}

func (this *DataBase) Get(key string) (interface{}, error) {
	return nil, nil
}

func (this *GroupData) KVs() []*KV {
	return []*KV{}
}

type DataOption func(Data)

func KeyConvertor(keyIn, keyOut func(string) string) DataOption {
	return func(d Data) {
		d.SetKeyConvertor(keyIn, keyOut)
	}
}
