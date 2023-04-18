package syftjson

import (
	"github.com/metasources/buildx/internal"
	"github.com/metasources/buildx/buildx/sbom"
)

const ID sbom.FormatID = "buildx-json"

func Format() sbom.Format {
	return sbom.NewFormat(
		internal.JSONSchemaVersion,
		encoder,
		decoder,
		validator,
		ID, "json", "buildx",
	)
}
