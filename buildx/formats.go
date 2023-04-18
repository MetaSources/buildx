package buildx

import (
	"github.com/metasources/buildx/buildx/formats"
	"github.com/metasources/buildx/buildx/formats/cyclonedxjson"
	"github.com/metasources/buildx/buildx/formats/cyclonedxxml"
	"github.com/metasources/buildx/buildx/formats/github"
	"github.com/metasources/buildx/buildx/formats/spdxjson"
	"github.com/metasources/buildx/buildx/formats/spdxtagvalue"
	"github.com/metasources/buildx/buildx/formats/buildxjson"
	"github.com/metasources/buildx/buildx/formats/table"
	"github.com/metasources/buildx/buildx/formats/template"
	"github.com/metasources/buildx/buildx/formats/text"
	"github.com/metasources/buildx/buildx/sbom"
)

// these have been exported for the benefit of API users
// TODO: deprecated: now that the formats package has been moved to buildx/formats, will be removed in v1.0.0
const (
	JSONFormatID          = buildxjson.ID
	TextFormatID          = text.ID
	TableFormatID         = table.ID
	CycloneDxXMLFormatID  = cyclonedxxml.ID
	CycloneDxJSONFormatID = cyclonedxjson.ID
	GitHubFormatID        = github.ID
	SPDXTagValueFormatID  = spdxtagvalue.ID
	SPDXJSONFormatID      = spdxjson.ID
	TemplateFormatID      = template.ID
)

// TODO: deprecated, moved to buildx/formats/formats.go. will be removed in v1.0.0
func FormatIDs() (ids []sbom.FormatID) {
	return formats.AllIDs()
}

// TODO: deprecated, moved to buildx/formats/formats.go. will be removed in v1.0.0
func FormatByID(id sbom.FormatID) sbom.Format {
	return formats.ByNameAndVersion(string(id), "")
}

// TODO: deprecated, moved to buildx/formats/formats.go. will be removed in v1.0.0
func FormatByName(name string) sbom.Format {
	return formats.ByName(name)
}

// TODO: deprecated, moved to buildx/formats/formats.go. will be removed in v1.0.0
func IdentifyFormat(by []byte) sbom.Format {
	return formats.Identify(by)
}
