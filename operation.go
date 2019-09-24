package go2openapi

//Operation Object
type Operation struct {
	Tags         []string               `json:"tags,omitempty"`
	Summary      string                 `json:"summary,omitempty"`
	Description  string                 `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
	OperationID  string                 `json:"operationId,omitempty"`
	Parameters   `json:"parameters,omitempty"`
	*RequestBody `json:"requestBody,omitempty"`
	Responses    `json:"responses"` //required
	// Callbacks    map[string]Callback    `json:"callbacks,omitempty"`
	Deprecated bool `json:"deprecated,omitempty"`
	// Security   *SecurityRequirement `json:"security,omitempty"`
	Servers `json:"servers,omitempty"`
}

//External Documentation Object
type ExternalDocumentation struct {
	URL         string `json:"url"` //required
	Description string `json:"description,omitempty"`
}
