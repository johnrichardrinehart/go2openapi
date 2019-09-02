package yaml

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

func Marshal(i interface{}) ([]byte, error) {
	json, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	var y interface{}
	err = yaml.Unmarshal(json, &y)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(y)
}
