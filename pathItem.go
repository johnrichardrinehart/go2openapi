package go2openapi

//Path Item Object
type PathItem struct {
	Summary     string     `json:"summary,omitempty"`
	Description string     `json:"description,omitempty"`
	Get         *Operation `json:"get,omitempty"`
	Patch       *Operation `json:"patch,omitempty"`
	Post        *Operation `json:"post,omitempty"`
	*Parameters `json:"parameters,omitempty"`
}

//Operation Object
type Operation struct {
	Tags         []string               `json:"tags,omitempty"`
	Summary      string                 `json:"summary,omitempty"`
	Description  string                 `json:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`
	OperationID  string                 `json:"operationId,omitempty"`
	*Parameters  `json:"parameters,omitempty"`
	// RequestBody  *RequestBody           `json:"requestBody,omitempty"`
	Responses `json:"responses"` //required
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

//Response Object
type Response struct {
	Description string               `json:"description"` //required
	Headers     map[string]Header    `json:"headers,omitempty"`
	Content     map[string]MediaType `json:"content,omitempty"`
	// Links       map[string]Link      `json:"links,omitempty"`
}

//Header Object
type Header struct {
	Description     string      `json:"description,omitempty"`
	Required        bool        `json:"required,omitempty"`
	Deprecated      bool        `json:"deprecated,omitempty"`
	AllowEmptyValue bool        `json:"allowEmptyValue,omitempty"`
	Style           string      `json:"style,omitempty"`
	Explode         bool        `json:"explode,omitempty"`
	AllowReserved   bool        `json:"allowReserved,omitempty"`
	Schema          Schema      `json:"schema,omitempty"`
	Example         interface{} `json:"example,omitempty"`
}

//MediaType Object
type MediaType struct {
	Schema  *Schema     `json:"schema,omitempty"`
	Example interface{} `json:"example,omitempty"`
	// Examples map[string]Example `json:"examples,omitempty"`
	// Encoding map[string]Encoding `json:"encoding,omitempty"`
}
