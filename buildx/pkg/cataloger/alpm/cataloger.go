package alpm

import (
	"github.com/metasources/buildx/buildx/pkg"
	"github.com/metasources/buildx/buildx/pkg/cataloger/generic"
)

const catalogerName = "alpmdb-cataloger"

func NewAlpmdbCataloger() *generic.Cataloger {
	return generic.NewCataloger(catalogerName).
		WithParserByGlobs(parseAlpmDB, pkg.AlpmDBGlob)
}
