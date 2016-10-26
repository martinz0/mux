package mux

type PathParam map[string][]byte

func NewPathParam() PathParam {
	return make(PathParam)
}
