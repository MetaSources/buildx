package integration

import (
	"testing"

	"github.com/metasources/buildx/buildx/source"
)

func TestRegressionJavaNoMainPackage(t *testing.T) { // Regression: https://github.com/metasources/buildx/issues/252
	catalogFixtureImage(t, "image-java-no-main-package", source.SquashedScope, nil)
}
