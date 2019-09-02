package go2openapi

//Info Object
type Info struct {
	Title          string `json:"title"`
	Description    string `json:"description,omitempty"`
	TermsOfService string `json:"termsOfService,omitempty"`
	// Contact        *Contact `json:"contact,omitempty"`
	License *License `json:"license,omitempty"`
	Version string   `json:"version"`
}
