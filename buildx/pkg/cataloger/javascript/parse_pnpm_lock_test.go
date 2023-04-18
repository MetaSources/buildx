package javascript

import (
	"testing"

	"github.com/metasources/buildx/buildx/artifact"
	"github.com/metasources/buildx/buildx/pkg"
	"github.com/metasources/buildx/buildx/pkg/cataloger/internal/pkgtest"
	"github.com/metasources/buildx/buildx/source"
)

func TestParsePnpmLock(t *testing.T) {
	var expectedRelationships []artifact.Relationship
	fixture := "test-fixtures/pnpm/pnpm-lock.yaml"

	locationSet := source.NewLocationSet(source.NewLocation(fixture))

	expectedPkgs := []pkg.Package{
		{
			Name:      "nanoid",
			Version:   "3.3.4",
			PURL:      "pkg:npm/nanoid@3.3.4",
			Locations: locationSet,
			Language:  pkg.JavaScript,
			Type:      pkg.NpmPkg,
		},
		{
			Name:      "picocolors",
			Version:   "1.0.0",
			PURL:      "pkg:npm/picocolors@1.0.0",
			Locations: locationSet,
			Language:  pkg.JavaScript,
			Type:      pkg.NpmPkg,
		},
		{
			Name:      "source-map-js",
			Version:   "1.0.2",
			PURL:      "pkg:npm/source-map-js@1.0.2",
			Locations: locationSet,
			Language:  pkg.JavaScript,
			Type:      pkg.NpmPkg,
		},
	}

	pkgtest.TestFileParser(t, fixture, parsePnpmLock, expectedPkgs, expectedRelationships)
}
