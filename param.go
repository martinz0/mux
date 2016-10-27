package mux

type PathParam map[string][]byte

func NewPathParam() PathParam {
	return make(PathParam)
}

type QueryParam map[string]string

func NewQueryParam() QueryParam {
	return make(QueryParam)
}
