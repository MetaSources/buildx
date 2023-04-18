package pkg

import "github.com/metasources/buildx/buildx/artifact"

func NewRelationships(catalog *Catalog) []artifact.Relationship {
	rels := RelationshipsByFileOwnership(catalog)
	rels = append(rels, RelationshipsEvidentBy(catalog)...)
	return rels
}
