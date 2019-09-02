package go2openapi

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
	Version     string              `json:"openapi"`
	Info        Info                `json:"info"`
	Servers     []Server            `json:"servers,omitempty"`
	Paths       map[string]PathItem `json:"paths"`
	*Components `json:"components,omitempty"`
}
