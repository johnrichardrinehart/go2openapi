package spec

import (
	"github.com/johnrichardrinehart/go2openapi"
)

//GetRoot gets the root
var GetRoot go2openapi.PathItem = go2openapi.PathItem{
	Summary:     "Get the root endpoint",
	Description: "Target this endpoint to get stuff at the root.",
}

//Spec is a minimum viable OpenAPI specification
var Spec go2openapi.Specification = go2openapi.Specification{
	Version: go2openapi.OPENAPIVERSION,
	Info: go2openapi.Info{
		Title:       "Minimum Viable Spec",
		Description: "A demonstration of the minimum viable OpenAPI structure.",
		Version:     go2openapi.GO2OPENAPIVERSION,
	},
}
