package go2openapi

//Components Object
type Components struct {
	*Schemas    `json:"schemas,omitempty"`
	*Parameters `json:"parameters,omitempty"`
}
