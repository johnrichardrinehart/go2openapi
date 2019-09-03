package go2openapi

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

//GO2OPENAPIVERSION is the version of this library
const GO2OPENAPIVERSION = "0.0.1"

//OPENAPIVERSION is the OpenAPI version against which this code was written
const OPENAPIVERSION = "3.0.2"

//Ptrue - pointer to true
var Ptrue = new(bool)

//Pfalse - pointer to false
var Pfalse = new(bool)

func init() {
	*Ptrue = true
	*Pfalse = false
}

//Specification Object
type Specification struct {
	OpenAPIVersion string         `json:"openapi"` // required
	Info           Info           `json:"info"`    // required
	Paths          `json:"paths"` //required
	*Servers       `json:"servers,omitempty"`
	*Components    `json:"components,omitempty"`
}

//MarshalJSON covers the case of the Paths variable being unset since, as a map, its 0 value is `nil` which marshals to JSON null, not an empty object as desired.
func (s Specification) MarshalJSON() ([]byte, error) {
	type Alias Specification
	if s.Paths == nil {
		s.Paths = make(map[string]PathItem)
	}
	return json.Marshal(Alias(s))
}

func (s Specification) MarshalYAML() ([]byte, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var y interface{}
	err = yaml.Unmarshal(b, &y)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(y)
}
