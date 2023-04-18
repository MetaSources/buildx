package cyclonedxjson

import (
	"github.com/CycloneDX/cyclonedx-go"

	"github.com/metasources/buildx/buildx/formats/common/cyclonedxhelpers"
	"github.com/metasources/buildx/buildx/sbom"
)

const ID sbom.FormatID = "cyclonedx-json"

func Format() sbom.Format {
	return sbom.NewFormat(
		sbom.AnyVersion,
		encoder,
		cyclonedxhelpers.GetDecoder(cyclonedx.BOMFileFormatJSON),
		cyclonedxhelpers.GetValidator(cyclonedx.BOMFileFormatJSON),
		ID,
	)
}
