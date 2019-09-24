package go2openapi

type Responses map[string]*Response

func (r Responses) AddResponse(name string, response *Response) {
	r[name] = response
}
