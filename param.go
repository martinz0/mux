package mux

type Params []param

type param struct {
	key, value string
}

func (p Params) Get(key string) string {
	for _, param := range p {
		if param.key == key {
			return param.value
		}
	}
	return ""
}

func (p *Params) Set(key, value string) {
	*p = append(*p, param{key, value})
}

func (p *Params) reset() {
	*p = (*p)[:0]
}
