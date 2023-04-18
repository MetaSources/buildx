package text

import (
	"github.com/metasources/buildx/buildx/sbom"
)

const ID sbom.FormatID = "buildx-text"

func Format() sbom.Format {
	return sbom.NewFormat(
		sbom.AnyVersion,
		encoder,
		nil,
		nil,
		ID, "text",
	)
}
