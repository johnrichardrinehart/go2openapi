package spec

import "github.com/johnrichardrinehart/go2openapi"

var petSchema go2openapi.Schema = go2openapi.Schema{
	Type:     "object",
	Required: []string{"id", "name"},
	Properties: map[string]go2openapi.Schema{
		"id": go2openapi.Schema{
			Type:   "integer",
			Format: "int64",
		},
		"name": go2openapi.Schema{
			Type: "string",
		},
		"tag": go2openapi.Schema{
			Type: "string",
		},
	},
}

var petsSchema go2openapi.Schema = go2openapi.Schema{
	Type: "array",
	Items: &go2openapi.Schema{
		Ref: "#/components/schemas/Pet",
	},
}

var errorSchema go2openapi.Schema = go2openapi.Schema{
	Type:     "object",
	Required: []string{"code", "message"},
	Properties: map[string]go2openapi.Schema{
		"code": go2openapi.Schema{
			Type:   "integer",
			Format: "int32",
		},
		"message": go2openapi.Schema{
			Type: "string",
		},
	},
}

var comps = go2openapi.Components{
	Schemas: map[string]go2openapi.Schema{
		"Pet":   petSchema,
		"Pets":  petsSchema,
		"Error": errorSchema,
	},
}
