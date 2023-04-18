package integration

import (
	"testing"

	"github.com/metasources/buildx/buildx/pkg"
	"github.com/metasources/buildx/buildx/source"
)

func TestSqliteRpm(t *testing.T) {
	// This is a regression test for issue #469 (https://github.com/metasources/buildx/issues/469). Recent RPM
	// based distribution store package data in an sqlite database
	sbom, _ := catalogFixtureImage(t, "image-sqlite-rpmdb", source.SquashedScope, nil)

	expectedPkgs := 139
	actualPkgs := 0
	for range sbom.Artifacts.PackageCatalog.Enumerate(pkg.RpmPkg) {
		actualPkgs += 1
	}

	if actualPkgs != expectedPkgs {
		t.Errorf("unexpected number of RPM packages: %d != %d", expectedPkgs, actualPkgs)
	}
}
