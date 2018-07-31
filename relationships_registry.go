package gaia

import "go.aporeto.io/elemental"

const nodocString = "[nodoc]" // nolint: varcheck,deadcode

var relationshipsRegistry elemental.RelationshipsRegistry

func init() {

	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[APIAuthorizationPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[APICheckIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AWSAPIGatewayIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AWSAccountIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AWSRegisterIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AccountIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AccountCheckIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[ActivateIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[ActivityIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AlarmIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AppIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AuditProfileIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcerprofile": (&elemental.RelationshipInfo{}).Build(),
			"root":            (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcerprofile": (&elemental.RelationshipInfo{}).Build(),
			"root":            (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AuthIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AuthorityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AutomationIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[AutomationTemplateIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[CategoryIdentity] = &elemental.Relationship{}

	relationshipsRegistry[CertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[CustomerIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[DependencyMapIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[EmailIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[EnforcerIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcerprofilemappingpolicy": (&elemental.RelationshipInfo{}).Build(),
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcerprofilemappingpolicy": (&elemental.RelationshipInfo{}).Build(),
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[EnforcerProfileIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer":                     (&elemental.RelationshipInfo{}).Build(),
			"enforcerprofilemappingpolicy": (&elemental.RelationshipInfo{}).Build(),
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer":                     (&elemental.RelationshipInfo{}).Build(),
			"enforcerprofilemappingpolicy": (&elemental.RelationshipInfo{}).Build(),
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[EnforcerProfileMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[EventLogIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[ExportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[ExternalAccessIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[ExternalNetworkIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[ExternalServiceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"networkaccesspolicy": (&elemental.RelationshipInfo{}).Build(),
			"root":                (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"networkaccesspolicy": (&elemental.RelationshipInfo{}).Build(),
			"root":                (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[FileAccessIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": (&elemental.RelationshipInfo{}).Build(),
			"root":           (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": (&elemental.RelationshipInfo{}).Build(),
			"root":           (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[FileAccessPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[FilePathIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[FlowReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[FlowStatisticIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[HookPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[ImportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[InstallationIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[InstalledAppIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[InvoiceIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[InvoiceRecordIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[IsolationProfileIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[IssueIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[JaegerbatchIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[K8SClusterIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[KubernetesClusterIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[LogIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"installedapp": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"installedapp": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[MessageIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[NamespaceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[NamespaceMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[NetworkAccessPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[PasswordResetIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[PlanIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[PokeIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer": (&elemental.RelationshipInfo{}).Build(),
			"processingunit": (&elemental.RelationshipInfo{
				Parameters: []elemental.Parameter{
					elemental.Parameter{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Paused",
							"Running",
							"Stopped",
						},
					},
				},
			}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": (&elemental.RelationshipInfo{}).Build(),
			"processingunit": (&elemental.RelationshipInfo{
				Parameters: []elemental.Parameter{
					elemental.Parameter{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Paused",
							"Running",
							"Stopped",
						},
					},
				},
			}).Build(),
		},
	}

	relationshipsRegistry[PolicyIdentity] = &elemental.Relationship{
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[PolicyRefreshIdentity] = &elemental.Relationship{}

	relationshipsRegistry[PolicyRendererIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[PolicyRuleIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[PrivateKeyIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[ProcessingUnitIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"networkaccesspolicy": (&elemental.RelationshipInfo{}).Build(),
			"root":                (&elemental.RelationshipInfo{}).Build(),
			"service":             (&elemental.RelationshipInfo{}).Build(),
			"servicedependency":   (&elemental.RelationshipInfo{}).Build(),
			"vulnerability":       (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"networkaccesspolicy": (&elemental.RelationshipInfo{}).Build(),
			"root":                (&elemental.RelationshipInfo{}).Build(),
			"service":             (&elemental.RelationshipInfo{}).Build(),
			"servicedependency":   (&elemental.RelationshipInfo{}).Build(),
			"vulnerability":       (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[ProcessingUnitPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[QuotaCheckIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[QuotaPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[RESTAPISpecIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root":    (&elemental.RelationshipInfo{}).Build(),
			"service": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root":    (&elemental.RelationshipInfo{}).Build(),
			"service": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[RemoteProcessorIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[RenderedPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[ReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[RevocationIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[RoleIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ServiceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit":    (&elemental.RelationshipInfo{}).Build(),
			"root":              (&elemental.RelationshipInfo{}).Build(),
			"servicedependency": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit":    (&elemental.RelationshipInfo{}).Build(),
			"root":              (&elemental.RelationshipInfo{}).Build(),
			"servicedependency": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[ServiceDependencyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[SquallTagIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[StatsQueryIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[SuggestedPolicyIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[TabulationIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[TagIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[TagInjectIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[TokenIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[TokenScopePolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[TriggerIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"automation": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"automation": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"automation": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[VulnerabilityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": (&elemental.RelationshipInfo{}).Build(),
			"root":           (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": (&elemental.RelationshipInfo{}).Build(),
			"root":           (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[X509CertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

	relationshipsRegistry[X509CertificateCheckIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": (&elemental.RelationshipInfo{}).Build(),
		},
	}

}
