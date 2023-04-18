package text

import (
	"flag"
	"testing"

	"github.com/metasources/buildx/buildx/formats/internal/testutils"
)

var updateTextEncoderGoldenFiles = flag.Bool("update-text", false, "update the *.golden files for text encoder")

func TestTextDirectoryEncoder(t *testing.T) {
	testutils.AssertEncoderAgainstGoldenSnapshot(t,
		Format(),
		testutils.DirectoryInput(t),
		*updateTextEncoderGoldenFiles,
		false,
	)
}

func TestTextImageEncoder(t *testing.T) {
	testImage := "image-simple"
	testutils.AssertEncoderAgainstGoldenImageSnapshot(t,
		Format(),
		testutils.ImageInput(t, testImage, testutils.FromSnapshot()),
		testImage,
		*updateTextEncoderGoldenFiles,
		false,
	)
}
