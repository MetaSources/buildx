package buildxjson

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/metasources/buildx/buildx/formats/buildxjson/model"
)

func validator(reader io.Reader) error {
	type Document struct {
		Schema model.Schema `json:"schema"`
	}

	dec := json.NewDecoder(reader)

	var doc Document
	err := dec.Decode(&doc)
	if err != nil {
		return fmt.Errorf("unable to decode: %w", err)
	}

	// note: we accept all schema versions
	// TODO: add per-schema version parsing
	if strings.Contains(doc.Schema.URL, "metasources/buildx") {
		return nil
	}
	return fmt.Errorf("could not extract buildx schema")
}
