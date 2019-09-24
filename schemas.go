package go2openapi

//Schema Object
type Schema struct {
	Title      string            `json:"title,omitempty"`
	Type       string            `json:"type,omitempty"`
	Format     string            `json:"format,omitempty"`
	Required   []string          `json:"required,omitempty"`
	Properties map[string]Schema `json:"properties,omitempty"`
	Items      *Schema           `json:"items,omitempty"`
	Example    interface{}       `json:"example,omitempty"`
	Ref        string            `json:"$ref,omitempty"`
}

type Schemas map[string]Schema

func (s Schemas) AddSchema(name string, schema Schema) {
	s[name] = schema
}
