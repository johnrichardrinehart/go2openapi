package go2openapi

//Parameter Object
type Parameter struct {
	Name            string                 `json:"name"` //required
	In              string                 `json:"in"`   //required
	Description     string                 `json:"description,omitempty"`
	Required        *bool                  `json:"required,omitempty"`
	Deprecated      *bool                  `json:"deprecated,omitempty"`
	AllowEmptyValue *bool                  `json:"allowEmptyValue,omitempty"`
	Style           string                 `json:"style,omitempty"`
	Explode         *bool                  `json:"explode,omitempty"`
	AllowReserved   *bool                  `json:"allowReserved,omitempty"`
	Schema          Schema                 `json:"schema,omitempty"`
	Example         interface{}            `json:"example,omitempty"`
	Examples        map[string]interface{} `json:"examples,omitempty"`
}

type Parameters []Parameter

func (ps *Parameters) AddParameters(parameters ...Parameter) {
	*ps = append(*ps, parameters...)
}
