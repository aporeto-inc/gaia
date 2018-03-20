package gaia

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
	}

	relationshipsRegistry[APICheckIdentity] = &elemental.Relationship{}

	relationshipsRegistry[APIServiceIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[AWSAccountIdentity] = &elemental.Relationship{
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
	}

	relationshipsRegistry[AccountIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[AccountCheckIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ActivateIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ActivityIdentity] = &elemental.Relationship{
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
	}

	relationshipsRegistry[AlarmIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[AuditProfileIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[AuthIdentity] = &elemental.Relationship{}

	relationshipsRegistry[AuthorityIdentity] = &elemental.Relationship{
		AllowsDelete: map[string]bool{
			"root": true,
		},
	}

	relationshipsRegistry[AutomationIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[AutomationTemplateIdentity] = &elemental.Relationship{
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
	}

	relationshipsRegistry[AvailableServiceIdentity] = &elemental.Relationship{}

	relationshipsRegistry[CategoryIdentity] = &elemental.Relationship{}

	relationshipsRegistry[CertificateIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[DependencyMapIdentity] = &elemental.Relationship{}

	relationshipsRegistry[EmailIdentity] = &elemental.Relationship{}

	relationshipsRegistry[EnforcerIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[EnforcerProfileIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[EnforcerProfileMappingPolicyIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[ExportIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ExternalAccessIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ExternalServiceIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[FileAccessIdentity] = &elemental.Relationship{}

	relationshipsRegistry[FileAccessPolicyIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[FilePathIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[FlowStatisticIdentity] = &elemental.Relationship{}

	relationshipsRegistry[HookPolicyIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[ImportIdentity] = &elemental.Relationship{}

	relationshipsRegistry[InstallationIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[IsolationProfileIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[IssueIdentity] = &elemental.Relationship{}

	relationshipsRegistry[JaegerbatchIdentity] = &elemental.Relationship{}

	relationshipsRegistry[KubernetesClusterIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[LogIdentity] = &elemental.Relationship{}

	relationshipsRegistry[MessageIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[NamespaceIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[NamespaceMappingPolicyIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[NetworkAccessPolicyIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[PasswordResetIdentity] = &elemental.Relationship{}

	relationshipsRegistry[PlanIdentity] = &elemental.Relationship{
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
	}

	relationshipsRegistry[PokeIdentity] = &elemental.Relationship{}

	relationshipsRegistry[PolicyIdentity] = &elemental.Relationship{
		AllowsDelete: map[string]bool{
			"root": true,
		},
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
	}

	relationshipsRegistry[PolicyRefreshIdentity] = &elemental.Relationship{}

	relationshipsRegistry[PolicyRuleIdentity] = &elemental.Relationship{
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
	}

	relationshipsRegistry[PrivateKeyIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[ProcessingUnitIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[ProcessingUnitPolicyIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[QuotaPolicyIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[RemoteProcessorIdentity] = &elemental.Relationship{}

	relationshipsRegistry[RenderedPolicyIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ReportIdentity] = &elemental.Relationship{}

	relationshipsRegistry[RevocationIdentity] = &elemental.Relationship{
		AllowsUpdate: map[string]bool{
			"root": true,
		},
		AllowsPatch: map[string]bool{
			"root": true,
		},
	}

	relationshipsRegistry[RoleIdentity] = &elemental.Relationship{}

	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ServiceIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[StatsQueryIdentity] = &elemental.Relationship{}

	relationshipsRegistry[SuggestedPolicyIdentity] = &elemental.Relationship{}

	relationshipsRegistry[SystemCallIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[TabulationIdentity] = &elemental.Relationship{}

	relationshipsRegistry[TagIdentity] = &elemental.Relationship{}

	relationshipsRegistry[TokenIdentity] = &elemental.Relationship{}

	relationshipsRegistry[TokenScopePolicyIdentity] = &elemental.Relationship{
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
	}

	relationshipsRegistry[TriggerIdentity] = &elemental.Relationship{}

	relationshipsRegistry[VulnerabilityIdentity] = &elemental.Relationship{
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
	}

	relationshipsRegistry[X509CertificateIdentity] = &elemental.Relationship{}

	relationshipsRegistry[X509CertificateCheckIdentity] = &elemental.Relationship{
		AllowsRetrieve: map[string]bool{
			"root": true,
		},
	}

}
