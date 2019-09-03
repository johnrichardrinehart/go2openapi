package spec

import "github.com/johnrichardrinehart/go2openapi"

var getPetsParameters = go2openapi.Parameters{
	go2openapi.Parameter{
		Name:        "limit",
		In:          "query",
		Description: "How many items to return at one time (max 100)",
		Required:    go2openapi.Pfalse,
		Schema: go2openapi.Schema{
			Type:   "integer",
			Format: "int32",
		},
	},
}

var getPets = go2openapi.Operation{
	Summary:     "List all pets",
	OperationID: "listPets",
	Tags:        []string{"pets"},
	Parameters:  &getPetsParameters,
	Responses: map[string]go2openapi.Response{
		"200": go2openapi.Response{
			Description: "A paged array of pets",
			Headers: map[string]go2openapi.Header{
				"x-next": go2openapi.Header{
					Description: "A link to the next page of responses",
					Schema: go2openapi.Schema{
						Type: "string",
					},
				},
			},
			Content: map[string]go2openapi.MediaType{
				"application/json": go2openapi.MediaType{
					Schema: &go2openapi.Schema{
						Ref: "#/components/schemas/Pets",
					},
				},
			},
		},
		"default": go2openapi.Response{
			Description: "unexpected error",
			Content: map[string]go2openapi.MediaType{
				"application/json": go2openapi.MediaType{
					Schema: &go2openapi.Schema{
						Ref: "#/components/schemas/Error",
					},
				},
			},
		},
	},
}

var postPets = go2openapi.Operation{
	Summary:     "Create a pet",
	OperationID: "createPets",
	Tags:        []string{"pets"},
	Responses: map[string]go2openapi.Response{
		"default": go2openapi.Response{
			Description: "unexpected error",
			Content: map[string]go2openapi.MediaType{
				"application/json": go2openapi.MediaType{
					Schema: &go2openapi.Schema{
						Ref: "#/components/schemas/Error",
					},
				},
			},
		},
		"201": go2openapi.Response{
			Description: "Null response",
		},
	},
}

var getPetsPetIdParameters = go2openapi.Parameters{
	go2openapi.Parameter{
		Name:        "petId",
		In:          "path",
		Required:    go2openapi.Ptrue,
		Description: "The id of the pet to retrieve",
		Schema: go2openapi.Schema{
			Type: "string",
		},
	},
}

var getPetsPetIDResponses = make(go2openapi.Responses)

var getPetsPetID = go2openapi.Operation{
	Summary:     "Info for a specific pet",
	OperationID: "showPetById",
	Tags:        []string{"pets"},
	Parameters:  &getPetsPetIdParameters,
}

var paths = go2openapi.Paths{
	`/pets`: go2openapi.PathItem{
		Get:  &getPets,
		Post: &postPets,
	},
	`/pets/{petId}`: go2openapi.PathItem{
		Get: &getPetsPetID,
	},
}

func init() {
	getPetsPetIDResponses.AddResponse("200", go2openapi.Response{
		Description: "Expected response to a valid request",
		Content: map[string]go2openapi.MediaType{
			"application/json": go2openapi.MediaType{
				Schema: &go2openapi.Schema{
					Ref: "#/components/schemas/Pet",
				},
			},
		},
	})
	getPetsPetIDResponses.AddResponse("default", go2openapi.Response{
		Description: "unexpected error",
		Content: map[string]go2openapi.MediaType{
			"application/json": go2openapi.MediaType{
				Schema: &go2openapi.Schema{
					Ref: "#/components/schemas/Error",
				},
			},
		},
	})

	getPetsPetID.Responses = getPetsPetID.Responses
}
