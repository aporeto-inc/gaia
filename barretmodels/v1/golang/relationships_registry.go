package barretmodels

import "github.com/aporeto-inc/elemental"

const nodocString = "[nodoc]" // nolint: varcheck

var relationshipsRegistry elemental.RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

func init() {
	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[AuthorityIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsDelete: map[string]bool{
			"root": true,
		},
	}
	relationshipsRegistry[CertificateIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
	}
	relationshipsRegistry[CheckIdentity] = &elemental.Relationship{
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"root": true,
		},
		AllowsInfo: map[string]bool{
			"root": true,
		},
	}
	relationshipsRegistry[PrivateKeyIdentity] = &elemental.Relationship{
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"root": true,
		},
		AllowsInfo: map[string]bool{
			"root": true,
		},
	}
	relationshipsRegistry[RevocationIdentity] = &elemental.Relationship{
		AllowsUpdate: map[string]bool{
			"root": true,
		},
	}
	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}
	relationshipsRegistry[TokenIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
	}
}
