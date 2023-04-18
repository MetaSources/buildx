package cyclonedxxml

import (
	"flag"
	"regexp"
	"testing"

	"github.com/metasources/buildx/buildx/formats/internal/testutils"
)

var updateCycloneDx = flag.Bool("update-cyclonedx", false, "update the *.golden files for cyclone-dx encoders")

func TestCycloneDxDirectoryEncoder(t *testing.T) {
	testutils.AssertEncoderAgainstGoldenSnapshot(t,
		Format(),
		testutils.DirectoryInput(t),
		*updateCycloneDx,
		false,
		cycloneDxRedactor,
	)
}

func TestCycloneDxImageEncoder(t *testing.T) {
	testImage := "image-simple"
	testutils.AssertEncoderAgainstGoldenImageSnapshot(t,
		Format(),
		testutils.ImageInput(t, testImage),
		testImage,
		*updateCycloneDx,
		false,
		cycloneDxRedactor,
	)
}

func cycloneDxRedactor(s []byte) []byte {
	serialPattern := regexp.MustCompile(`serialNumber="[a-zA-Z0-9\-:]+"`)
	rfc3339Pattern := regexp.MustCompile(`([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])[Tt]([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9]|60)(\.[0-9]+)?(([Zz])|([\+|\-]([01][0-9]|2[0-3]):[0-5][0-9]))`)
	sha256Pattern := regexp.MustCompile(`sha256:[A-Fa-f0-9]{64}`)

	for _, pattern := range []*regexp.Regexp{serialPattern, rfc3339Pattern, sha256Pattern} {
		s = pattern.ReplaceAll(s, []byte("redacted"))
	}

	// the bom-ref will be autogenerated every time, the value here should not be directly tested in snapshot tests
	bomRefPattern := regexp.MustCompile(` bom-ref="[a-zA-Z0-9\-:]+"`)
	bomRef3339Pattern := regexp.MustCompile(`([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])[Tt]([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9]|60)(\.[0-9]+)?(([Zz])|([\+|\-]([01][0-9]|2[0-3]):[0-5][0-9]))`)
	for _, pattern := range []*regexp.Regexp{bomRefPattern, bomRef3339Pattern} {
		s = pattern.ReplaceAll(s, []byte(""))
	}

	return s
}
