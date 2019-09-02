package spec

import (
	"github.com/johnrichardrinehart/go2openapi"
)

//GetRoot gets the root
var GetRoot go2openapi.PathItem = go2openapi.PathItem{
	Summary:     "Get the root endpoint",
	Description: "Target this endpoint to get stuff at the root.",
}

var license go2openapi.License = go2openapi.License{
	Name: "MIT",
}

var servers = []go2openapi.Server{
	go2openapi.Server{
		URL: "http://petstore.swagger.io/v1",
	},
}

//Spec matches the OpenAPI Petstore example specification
var Spec go2openapi.Specification = go2openapi.Specification{
	Version: "3.0.0",
	Info: go2openapi.Info{
		Title:   "Swagger Petstore",
		Version: "1.0.0",
		License: &license,
	},
	Servers:    servers,
	Paths:      paths,
	Components: &comps,
}
