package main

import (
	"fmt"
	"log"

	"github.com/johnrichardrinehart/go2openapi/examples/mvp/spec"
	"github.com/johnrichardrinehart/go2openapi/yaml"
)

func main() {
	spec := spec.Spec
	yml, err := yaml.Marshal(spec)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", yml)
}
