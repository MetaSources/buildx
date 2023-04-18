package table

import (
	"github.com/metasources/buildx/buildx/sbom"
)

const ID sbom.FormatID = "buildx-table"

func Format() sbom.Format {
	return sbom.NewFormat(
		sbom.AnyVersion,
		encoder,
		nil,
		nil,
		ID, "table",
	)
}
