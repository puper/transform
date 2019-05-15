package transform

type Transform struct {
	only           map[string]bool
	except         map[string]bool
	ignoreErrors   map[error]bool
	fieldIn        FieldMapper
	fieldOut       FieldMapper
	customMapper   map[string]string
	globalHandlers []FieldHandler // 只能处理某类数据
	fieldHandlers  map[string][]FieldHandler
}

func New() *Transform {
	return &Transform{
		only:          make(map[string]bool),
		except:        make(map[string]bool),
		ignoreErrors:  make(map[error]bool),
		fieldHandlers: make(map[string][]FieldHandler),
		fieldIn: func(s string) string {
			return s
		},
		fieldOut: func(s string) string {
			return s
		},
		customMapper: map[string]string{},
	}
}

func (this *Transform) Process(s, t interface{}) error {
	sp := DetectProvider(s)
	tp := DetectProvider(t)
	tfs := []string{}
	if len(this.only) > 0 {
		for f, _ := range this.only {
			tfs = append(tfs, f)
		}
	} else {
		tfs = make([]string, 0)
		if len(tp.Fields()) == 0 {
			for _, f := range sp.Fields() {
				f = this.fieldIn(f)
				if !this.except[f] {
					tfs = append(tfs, f)
				}
			}
		} else {
			for _, f := range tp.Fields() {
				if !this.except[f] {
					tfs = append(tfs, f)
				}
			}
		}
	}
	if len(tfs) == 0 {
		return nil
	}
	for _, tf := range tfs {
		sf := ""
		if _, ok := this.customMapper[tf]; ok {
			sf = this.customMapper[tf]
		} else {
			sf = this.fieldOut(tf)
		}
		// 根据source field 获取 source field value
		sv, err := sp.Get(sf)
		if !this.isIgnoreOrNilError(err) {
			return NewProcessError(sf, err)
		}
		for _, handler := range this.globalHandlers {
			tmp, err := handler(sv)
			if err == nil {
				sv = tmp
			}
			if err == ErrConvertFailed || err == ErrTypeNotMatch {
				continue
			}
			if !this.isIgnoreOrNilError(err) {
				return NewProcessError(sf, err)
			}
		}
		for _, handler := range this.fieldHandlers[tf] {
			tmp, err := handler(sv)
			if err == nil {
				sv = tmp
			}
			if !this.isIgnoreOrNilError(err) {
				return NewProcessError(sf, err)
			}
		}
		if err := tp.Set(tf, sv); !this.isIgnoreOrNilError(err) {
			return NewProcessError(sf, err)
		}
	}
	return nil
}

func (this *Transform) Only(fields ...string) *Transform {
	for _, f := range fields {
		this.only[f] = true
	}
	return this
}

func (this *Transform) Except(fields ...string) *Transform {
	for _, f := range fields {
		this.except[f] = true
	}
	return this
}

func (this *Transform) FieldMapper(in, out FieldMapper) *Transform {
	this.fieldIn = in
	this.fieldOut = out
	return this
}

func (this *Transform) CustomMapper(mapper map[string]string) *Transform {
	this.customMapper = mapper
	return this
}

func (this *Transform) IgnoreErrors(errs ...error) {
	for _, err := range errs {
		this.ignoreErrors[err] = true
	}
}

func (this *Transform) isIgnoreOrNilError(err error) bool {
	if err == nil {
		return true
	}
	return this.ignoreErrors[err]
}

func (this *Transform) FieldHandler(h FieldHandler, fields ...string) *Transform {
	if len(fields) == 0 {
		this.globalHandlers = append(this.globalHandlers, h)
	} else {
		for _, field := range fields {
			if _, ok := this.fieldHandlers[field]; ok {
				this.fieldHandlers[field] = append(this.fieldHandlers[field], h)
			} else {
				this.fieldHandlers[field] = []FieldHandler{h}
			}
		}
	}
	return this
}
