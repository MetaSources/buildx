package integration

import (
	"testing"

	"github.com/metasources/buildx/buildx/pkg"
	"github.com/metasources/buildx/buildx/source"
)

func TestRustAudit(t *testing.T) {
	sbom, _ := catalogFixtureImage(t, "image-rust-auditable", source.SquashedScope, []string{"all"})

	expectedPkgs := 2
	actualPkgs := 0
	for range sbom.Artifacts.PackageCatalog.Enumerate(pkg.RustPkg) {
		actualPkgs += 1
	}

	if actualPkgs != expectedPkgs {
		t.Errorf("unexpected number of Rust packages: %d != %d", expectedPkgs, actualPkgs)
	}
}
