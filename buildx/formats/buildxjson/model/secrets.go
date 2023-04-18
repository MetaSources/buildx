package model

import (
	"github.com/metasources/buildx/buildx/file"
	"github.com/metasources/buildx/buildx/source"
)

type Secrets struct {
	Location source.Coordinates  `json:"location"`
	Secrets  []file.SearchResult `json:"secrets"`
}
