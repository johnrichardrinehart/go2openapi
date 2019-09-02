package go2openapi

//Server Object
type Server struct {
	URL         string                    `json:"url"` // required
	Description string                    `json:"description,omitempty"`
	Variables   map[string]ServerVariable `json:"variables,omitempty"`
}

//Server Variable Object
type ServerVariable struct {
	Enum        []string `json:"enum,omitempty"`
	Default     string   `json:"default"` //required
	Description string   `json:"description,omitempty"`
}

type Servers []Server

func (ss *Servers) AddServers(servers ...Server) {
	*ss = append(*ss, servers...)
}
