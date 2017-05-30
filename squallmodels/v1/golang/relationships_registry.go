package squallmodels

import "github.com/aporeto-inc/elemental"

const nodocString = "[nodoc]" // nolint: varcheck

var relationshipsRegistry elemental.RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

func init() {
	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[elemental.IdentityFromName("apicheck")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
	}
	relationshipsRegistry[elemental.IdentityFromName("hookpolicy")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
	relationshipsRegistry[elemental.IdentityFromName("apiauthorizationpolicy")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
	relationshipsRegistry[elemental.IdentityFromName("syscallaccess")] = &elemental.Relationship{
		AllowsRetrieve: map[string]bool{
			"processingunit": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"processingunit": true,
		},
		AllowsInfo: map[string]bool{
			"processingunit": true,
		},
	}
	relationshipsRegistry[elemental.IdentityFromName("systemcall")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
	relationshipsRegistry[elemental.IdentityFromName("tag")] = &elemental.Relationship{
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
	relationshipsRegistry[elemental.IdentityFromName("enforcer")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
	relationshipsRegistry[elemental.IdentityFromName("poke")] = &elemental.Relationship{
		AllowsRetrieve: map[string]bool{
			"enforcer": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"enforcer": true,
		},
		AllowsInfo: map[string]bool{
			"enforcer": true,
		},
	}
	relationshipsRegistry[elemental.IdentityFromName("filepath")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
	relationshipsRegistry[elemental.IdentityFromName("fileaccess")] = &elemental.Relationship{
		AllowsRetrieve: map[string]bool{
			"processingunit": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"processingunit": true,
		},
		AllowsInfo: map[string]bool{
			"processingunit": true,
		},
	}
	relationshipsRegistry[elemental.IdentityFromName("namespace")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
	relationshipsRegistry[elemental.IdentityFromName("policyrule")] = &elemental.Relationship{}
	relationshipsRegistry[elemental.IdentityFromName("externalservice")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
	relationshipsRegistry[elemental.IdentityFromName("role")] = &elemental.Relationship{
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
	relationshipsRegistry[elemental.IdentityFromName("policy")] = &elemental.Relationship{
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
	relationshipsRegistry[elemental.IdentityFromName("flowstatistic")] = &elemental.Relationship{
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
	relationshipsRegistry[elemental.IdentityFromName("suggestedpolicy")] = &elemental.Relationship{
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
	relationshipsRegistry[elemental.IdentityFromName("tabulation")] = &elemental.Relationship{
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
	relationshipsRegistry[elemental.IdentityFromName("quotapolicy")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
	relationshipsRegistry[elemental.IdentityFromName("fileaccesspolicy")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
	relationshipsRegistry[elemental.IdentityFromName("enforcerprofile")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"root":     true,
			"enforcer": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"root":     true,
			"enforcer": true,
		},
		AllowsInfo: map[string]bool{
			"root":     true,
			"enforcer": true,
		},
	}
	relationshipsRegistry[elemental.IdentityFromName("renderedpolicy")] = &elemental.Relationship{
		AllowsRetrieve: map[string]bool{
			"processingunit": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"processingunit": true,
		},
		AllowsInfo: map[string]bool{
			"processingunit": true,
		},
	}
	relationshipsRegistry[elemental.IdentityFromName("namespacemappingpolicy")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
	relationshipsRegistry[elemental.IdentityFromName("processingunit")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"vulnerability": true,
			"root":          true,
		},
		AllowsRetrieveMany: map[string]bool{
			"vulnerability": true,
			"root":          true,
		},
		AllowsInfo: map[string]bool{
			"vulnerability": true,
			"root":          true,
		},
	}
	relationshipsRegistry[elemental.IdentityFromName("dependencymap")] = &elemental.Relationship{
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
	relationshipsRegistry[elemental.IdentityFromName("vulnerability")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"root":           true,
			"processingunit": true,
		},
		AllowsRetrieveMany: map[string]bool{
			"root":           true,
			"processingunit": true,
		},
		AllowsInfo: map[string]bool{
			"root":           true,
			"processingunit": true,
		},
	}
	relationshipsRegistry[elemental.IdentityFromName("enforcerprofilemappingpolicy")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
	relationshipsRegistry[elemental.IdentityFromName("activity")] = &elemental.Relationship{
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
	relationshipsRegistry[elemental.IdentityFromName("root")] = &elemental.Relationship{}
	relationshipsRegistry[elemental.IdentityFromName("networkaccesspolicy")] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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
}
