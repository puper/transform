package transform

type KV struct {
	Key   string
	Value interface{}
}

type Option func(*Config)

type Handler func(Data, Data) error

type Config struct {
	Only     map[string]bool
	Except   map[string]bool
	Handlers map[string]Handler
}

func Transform(a Data, b Data, options ...Option) (err error) {
	c := &Config{
		Only:     map[string]bool{},
		Except:   map[string]bool{},
		Handlers: map[string]Handler{},
	}
	for _, option := range options {
		option(c)
	}
	for _, kv := range b.KVs() {
		if len(c.Only) > 0 && !c.Only[kv.Key] {
			continue
		}
		if len(c.Except) > 0 && c.Except[kv.Key] {
			continue
		}
		if handler, ok := c.Handlers[kv.Key]; ok {
			err = handler(a, b)
			if err != nil {
				return nil
			}
		} else {
			v, err := a.Get(kv.Key)
			if err != nil {
				return err
			}
			err = b.Set(kv.Key, v)
			if err != nil {
				return err
			}
		}
	}
	return
}

func Only(keys ...string) Option {
	return func(c *Config) {
		for _, key := range keys {
			c.Only[key] = true
		}
	}
}

func Except(keys ...string) Option {
	return func(c *Config) {
		for _, key := range keys {
			c.Except[key] = true
		}
	}
}
