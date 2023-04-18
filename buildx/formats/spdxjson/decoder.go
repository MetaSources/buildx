package spdxjson

import (
	"fmt"
	"io"

	"github.com/spdx/tools-golang/json"

	"github.com/metasources/buildx/buildx/formats/common/spdxhelpers"
	"github.com/metasources/buildx/buildx/sbom"
)

func decoder(reader io.Reader) (s *sbom.SBOM, err error) {
	doc, err := json.Read(reader)
	if err != nil {
		return nil, fmt.Errorf("unable to decode spdx-json: %w", err)
	}

	return spdxhelpers.ToBuildxModel(doc)
}
