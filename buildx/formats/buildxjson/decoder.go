package buildxjson

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/metasources/buildx/buildx/formats/buildxjson/model"
	"github.com/metasources/buildx/buildx/sbom"
)

func decoder(reader io.Reader) (*sbom.SBOM, error) {
	dec := json.NewDecoder(reader)

	var doc model.Document
	err := dec.Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("unable to decode buildx-json: %w", err)
	}

	return toBuildxModel(doc)
}
