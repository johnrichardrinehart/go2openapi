package go2openapi

//RequestBody defines the Request Body Object
type RequestBody struct {
	Content     map[string]MediaType `json:"content"` //required
	Description string               `json:"description,omitempty"`
	Required    bool                 `json:"required,omitempty"`
}
