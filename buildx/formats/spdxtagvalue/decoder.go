package spdxtagvalue

import (
	"fmt"
	"io"

	"github.com/spdx/tools-golang/tagvalue"

	"github.com/metasources/buildx/buildx/formats/common/spdxhelpers"
	"github.com/metasources/buildx/buildx/sbom"
)

func decoder(reader io.Reader) (*sbom.SBOM, error) {
	doc, err := tagvalue.Read(reader)
	if err != nil {
		return nil, fmt.Errorf("unable to decode spdx-tag-value: %w", err)
	}

	return spdxhelpers.ToBuildxModel(doc)
}
