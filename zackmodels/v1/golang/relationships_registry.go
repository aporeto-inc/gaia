package zackmodels

import "github.com/aporeto-inc/elemental"

const nodocString = "[nodoc]" // nolint: varcheck

var relationshipsRegistry elemental.RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

func init() {
	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[elemental.IdentityFromName("report")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
	}
	relationshipsRegistry[elemental.IdentityFromName("root")] = &elemental.Relationship{}
}
