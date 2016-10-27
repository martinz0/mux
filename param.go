package mux

import "bytes"

type params struct {
	params []*param
}

type param struct {
	key, value []byte
}

func NewParams() *params {
	return &params{
		params: make([]*param, 0),
	}
}

func (p *params) Get(key []byte) []byte {
	for _, param := range p.params {
		if bytes.Equal(param.key, key) {
			return param.value
		}
	}
	return nil
}

func (p *params) Set(key, value []byte) {
	p.params = append(p.params, &param{key, value})
}
