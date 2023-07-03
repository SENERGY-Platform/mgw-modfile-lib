package main

import (
	"github.com/SENERGY-Platform/mgw-modfile-lib/v1/model"
	"github.com/invopop/jsonschema"
	"log"
	"os"
)

func main() {
	t := jsonschema.Reflect(&model.ModFile{})
	b, err := t.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("v1/schema.json", b, 0666); err != nil {
		log.Fatal(err)
	}
}
