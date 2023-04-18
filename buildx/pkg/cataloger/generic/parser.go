package generic

import (
	"github.com/metasources/buildx/buildx/artifact"
	"github.com/metasources/buildx/buildx/linux"
	"github.com/metasources/buildx/buildx/pkg"
	"github.com/metasources/buildx/buildx/source"
)

type Environment struct {
	LinuxRelease *linux.Release
}

type Parser func(source.FileResolver, *Environment, source.LocationReadCloser) ([]pkg.Package, []artifact.Relationship, error)
