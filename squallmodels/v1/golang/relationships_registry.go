package squallmodels

import "github.com/aporeto-inc/elemental"

const nodocString = "[nodoc]" // nolint: varcheck,deadcode

var relationshipsRegistry elemental.RelationshipsRegistry

// Relationships returns the model relationships.
func Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

func init() {

	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[APIAuthorizationPolicyIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[APICheckIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
	}

	relationshipsRegistry[ActivityIdentity] = &elemental.Relationship{
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

	relationshipsRegistry[AuditProfileIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[DependencyMapIdentity] = &elemental.Relationship{
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

	relationshipsRegistry[EnforcerIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[EnforcerProfileIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
			"root": true,
		},
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"enforcer": true,
			"root":     true,
		},
		AllowsRetrieveMany: map[string]bool{
			"enforcer": true,
			"root":     true,
		},
		AllowsInfo: map[string]bool{
			"enforcer": true,
			"root":     true,
		},
	}

	relationshipsRegistry[EnforcerProfileMappingPolicyIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[ExternalAccessIdentity] = &elemental.Relationship{
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

	relationshipsRegistry[ExternalServiceIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
			"root": true,
		},
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"networkaccesspolicy": true,
			"root":                true,
		},
		AllowsRetrieveMany: map[string]bool{
			"networkaccesspolicy": true,
			"root":                true,
		},
		AllowsInfo: map[string]bool{
			"networkaccesspolicy": true,
			"root":                true,
		},
	}

	relationshipsRegistry[FileAccessIdentity] = &elemental.Relationship{
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

	relationshipsRegistry[FileAccessPolicyIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[FilePathIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[FlowStatisticIdentity] = &elemental.Relationship{
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

	relationshipsRegistry[HookPolicyIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[MessageIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[NamespaceIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[NamespaceMappingPolicyIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[NetworkAccessPolicyIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[PokeIdentity] = &elemental.Relationship{
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

	relationshipsRegistry[PolicyIdentity] = &elemental.Relationship{
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

	relationshipsRegistry[PolicyRuleIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ProcessingUnitIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
			"root": true,
		},
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"networkaccesspolicy": true,
			"root":                true,
			"vulnerability":       true,
		},
		AllowsRetrieveMany: map[string]bool{
			"networkaccesspolicy": true,
			"root":                true,
			"vulnerability":       true,
		},
		AllowsInfo: map[string]bool{
			"networkaccesspolicy": true,
			"root":                true,
			"vulnerability":       true,
		},
	}

	relationshipsRegistry[QuotaPolicyIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[RenderedPolicyIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
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

	relationshipsRegistry[RoleIdentity] = &elemental.Relationship{
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

	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}

	relationshipsRegistry[StatsQueryIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
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

	relationshipsRegistry[SuggestedPolicyIdentity] = &elemental.Relationship{
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[SystemCallIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
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

	relationshipsRegistry[TabulationIdentity] = &elemental.Relationship{
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

	relationshipsRegistry[TagIdentity] = &elemental.Relationship{
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

	relationshipsRegistry[VulnerabilityIdentity] = &elemental.Relationship{
		AllowsCreate: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"processingunit": true,
			"root":           true,
		},
		AllowsRetrieveMany: map[string]bool{
			"processingunit": true,
			"root":           true,
		},
		AllowsInfo: map[string]bool{
			"processingunit": true,
			"root":           true,
		},
	}

}
