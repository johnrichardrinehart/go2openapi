package main

import (
	"fmt"
	"log"

	"github.com/johnrichardrinehart/go2openapi/examples/petstore/spec"
)

func main() {
	spec := spec.Spec
	yml, err := spec.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", yml)
}
