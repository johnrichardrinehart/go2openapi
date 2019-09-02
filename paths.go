package go2openapi

type Paths map[string]PathItem

func (p Paths) AddPath(name string, path PathItem) {
	p[name] = path
}
