package python

import (
	"fmt"

	"github.com/pelletier/go-toml"

	"github.com/metasources/buildx/buildx/artifact"
	"github.com/metasources/buildx/buildx/pkg"
	"github.com/metasources/buildx/buildx/pkg/cataloger/generic"
	"github.com/metasources/buildx/buildx/source"
)

// integrity check
var _ generic.Parser = parsePoetryLock

type poetryMetadata struct {
	Packages []struct {
		Name        string `toml:"name"`
		Version     string `toml:"version"`
		Category    string `toml:"category"`
		Description string `toml:"description"`
		Optional    bool   `toml:"optional"`
	} `toml:"package"`
}

// parsePoetryLock is a parser function for poetry.lock contents, returning all python packages discovered.
func parsePoetryLock(_ source.FileResolver, _ *generic.Environment, reader source.LocationReadCloser) ([]pkg.Package, []artifact.Relationship, error) {
	tree, err := toml.LoadReader(reader)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to load poetry.lock for parsing: %w", err)
	}

	metadata := poetryMetadata{}
	err = tree.Unmarshal(&metadata)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to parse poetry.lock: %w", err)
	}

	var pkgs []pkg.Package
	for _, p := range metadata.Packages {
		pkgs = append(
			pkgs,
			newPackageForIndex(
				p.Name,
				p.Version,
				reader.Location.WithAnnotation(pkg.EvidenceAnnotationKey, pkg.PrimaryEvidenceAnnotation),
			),
		)
	}

	return pkgs, nil, nil
}
