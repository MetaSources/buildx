/*
Package portage provides a concrete Cataloger implementation for Gentoo Portage.
*/
package portage

import (
	"github.com/metasources/buildx/buildx/pkg/cataloger/generic"
)

func NewPortageCataloger() *generic.Cataloger {
	return generic.NewCataloger("portage-cataloger").
		WithParserByGlobs(parsePortageContents, "**/var/db/pkg/*/*/CONTENTS")
}
