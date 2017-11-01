package zackmodels

import "github.com/aporeto-inc/elemental"

const nodocString = "[nodoc]" // nolint: varcheck,deadcode

var relationshipsRegistry elemental.RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

func init() {
	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[ReportIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
	}
	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}
}
