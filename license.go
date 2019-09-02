package go2openapi

//License Object
type License struct {
	Name string `json:"name"` //required
	URL  string `json:"url,omitempty"`
}
