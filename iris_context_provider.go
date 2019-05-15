package transform

import (
	"strings"

	"github.com/kataras/iris"
)

// !!!NOT COMPLETED!!!
type IrisContextProvider struct {
	context iris.Context
}

func (this *IrisContextProvider) Set(f string, v interface{}) error {
	return nil
}

func (this *IrisContextProvider) Get(f string) (interface{}, error) {
	fs := strings.SplitN(f, ".", 1)
	if len(fs) != 2 {
		panic(ErrFieldNotFound)
	}
	if fs[0] == "query" {
		return this.context.URLParam(fs[1]), nil
	} else if fs[0] == "post" {
		return this.context.PostValue(fs[1]), nil
	} else if fs[0] == "params" {
		return this.context.Params().Get(fs[1]), nil
	} else if fs[0] == "cookie" {
		return this.context.GetCookie(fs[1]), nil
	}
	return "", nil
}

func (this *IrisContextProvider) Fields() []string {
	return []string{}
}
