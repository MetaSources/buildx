package integration

import (
	"testing"

	"github.com/metasources/buildx/buildx/pkg"
	"github.com/metasources/buildx/buildx/source"
)

func TestMarinerDistroless(t *testing.T) {
	sbom, _ := catalogFixtureImage(t, "image-mariner-distroless", source.SquashedScope, nil)

	expectedPkgs := 12
	actualPkgs := 0
	for range sbom.Artifacts.PackageCatalog.Enumerate(pkg.RpmPkg) {
		actualPkgs += 1
	}

	if actualPkgs != expectedPkgs {
		t.Errorf("unexpected number of RPM packages: %d != %d", expectedPkgs, actualPkgs)
	}
}
