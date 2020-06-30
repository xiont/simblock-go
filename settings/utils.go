package settings

import "simblock-go/utils"

var Rand = utils.NewMyRand()

type funcMap map[string]interface{}

func NewFuncMap(Name ...string) *funcMap {
	m := make(funcMap)
	for _, name := range Name {
		m[name] = nil
	}
	return &m
}

var FUNCS *funcMap = NewFuncMap(TABLE, ALGO)
