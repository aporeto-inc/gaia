package gaia

import "go.aporeto.io/elemental"

var relationshipsRegistry elemental.RelationshipsRegistry

func init() {

	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[APIAuthorizationPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[APICheckIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[APIProxyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AWSAPIGatewayIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AWSRegisterIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[AccessReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[AccountIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "associatedBillingID",
						Type: "string",
					},
					{
						Name: "name",
						Type: "string",
					},
					{
						Name: "status",
						Type: "string",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "associatedBillingID",
						Type: "string",
					},
					{
						Name: "name",
						Type: "string",
					},
					{
						Name: "status",
						Type: "string",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AccountCheckIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[ActivateIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"token",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "noRedirect",
						Type: "boolean",
					},
					{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"token",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "noRedirect",
						Type: "boolean",
					},
					{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[ActivityIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AlarmIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AppIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "name",
						Type: "string",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "name",
						Type: "string",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AppCredentialIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AuditProfileIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy": {},
			"enforcer":                  {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy": {},
			"enforcer":                  {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[AuditProfileMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[AuditReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[AuthnIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[AuthorityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AuthzIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[AutomationIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AutomationTemplateIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[CallIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"apiproxy": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"apiproxy": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"apiproxy": {},
		},
	}

	relationshipsRegistry[CategoryIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ClaimsIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ClauseMatchIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[CounterReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[CustomerIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[DNSLookupReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[DataPathCertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[DebugBundleIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"enforcer": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": {},
		},
	}

	relationshipsRegistry[DependencyMapIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "view",
						Type: "string",
					},
					{
						Name: "viewSuggestions",
						Type: "boolean",
					},
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "flowOffset",
						Type: "duration",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "view",
						Type: "string",
					},
					{
						Name: "viewSuggestions",
						Type: "boolean",
					},
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "flowOffset",
						Type: "duration",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[EmailIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[EnforcerIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy":    {},
			"enforcerprofilemappingpolicy": {},
			"hostservicemappingpolicy":     {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy":    {},
			"enforcerprofilemappingpolicy": {},
			"hostservicemappingpolicy":     {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[EnforcerLogIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[EnforcerProfileIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer":                     {},
			"enforcerprofilemappingpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer":                     {},
			"enforcerprofilemappingpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[EnforcerProfileMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[EnforcerRefreshIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"enforcer": {},
		},
	}

	relationshipsRegistry[EnforcerReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[EnforcerTraceReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[EventLogIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[ExportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ExternalNetworkIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[FileAccessPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[FileAccessReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[FilePathIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[FlowReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[GraphEdgeIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[GraphNodeIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[HTTPResourceSpecIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service": {},
		},
	}

	relationshipsRegistry[HitIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "reset",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"name",
								"targetID",
								"targetIdentity",
							},
							{
								"targetID",
								"targetIdentity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "name",
						Type:         "string",
						DefaultValue: "counter",
					},
					{
						Name: "targetID",
						Type: "string",
					},
					{
						Name: "targetIdentity",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"name",
								"targetID",
								"targetIdentity",
							},
							{
								"targetID",
								"targetIdentity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "name",
						Type:         "string",
						DefaultValue: "counter",
					},
					{
						Name: "targetID",
						Type: "string",
					},
					{
						Name: "targetIdentity",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[HookPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[HostServiceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "appliedServices",
						Type: "boolean",
					},
					{
						Name: "setServices",
						Type: "boolean",
					},
				},
			},
			"hostservicemappingpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "appliedServices",
						Type: "boolean",
					},
					{
						Name: "setServices",
						Type: "boolean",
					},
				},
			},
			"hostservicemappingpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[HostServiceMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[IPInfoIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"ip",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "ip",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"ip",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "ip",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ImageIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[ImageVulnerabilityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"image",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "image",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"image",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "image",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ImportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[ImportReferenceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"recipe": {},
			"root":   {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"recipe": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"recipe": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ImportRequestIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[InfrastructurePolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[InstalledAppIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[InvoiceIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[InvoiceRecordIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[IsolationProfileIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunitpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunitpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[IssueIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "asCookie",
						Type: "boolean",
					},
					{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[IssueServiceTokenIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[LDAPProviderIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[LogIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"installedapp": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"installedapp": {},
		},
	}

	relationshipsRegistry[LogoutIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[MessageIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[NamespaceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "authorized",
						Type: "boolean",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "authorized",
						Type: "boolean",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[NamespaceMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[NetworkAccessPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[OAUTHInfoIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "mode",
						Type: "enum",
						AllowedChoices: []string{
							"oidc",
						},
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "mode",
						Type: "enum",
						AllowedChoices: []string{
							"oidc",
						},
					},
				},
			},
		},
	}

	relationshipsRegistry[OAUTHKeyIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "mode",
						Type: "enum",
						AllowedChoices: []string{
							"oidc",
						},
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "mode",
						Type: "enum",
						AllowedChoices: []string{
							"oidc",
						},
					},
				},
			},
		},
	}

	relationshipsRegistry[OIDCProviderIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[PCCProviderIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[PacketReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PasswordResetIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"email",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "email",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"email",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "email",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[PingReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PlanIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PokeIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "cpuload",
						Type: "float",
					},
					{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					{
						Name: "memory",
						Type: "integer",
					},
					{
						Name: "processes",
						Type: "integer",
					},
					{
						Name: "sessionClose",
						Type: "boolean",
					},
					{
						Name: "sessionID",
						Type: "string",
					},
					{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Registered",
							"Connected",
							"Disconnected",
						},
					},
					{
						Name: "ts",
						Type: "time",
					},
					{
						Name: "version",
						Type: "string",
					},
					{
						Name: "zhash",
						Type: "integer",
					},
				},
			},
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					{
						Name: "notify",
						Type: "boolean",
					},
					{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Initialized",
							"Paused",
							"Running",
							"Stopped",
						},
					},
					{
						Name: "ts",
						Type: "time",
					},
					{
						Name: "zhash",
						Type: "integer",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "cpuload",
						Type: "float",
					},
					{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					{
						Name: "memory",
						Type: "integer",
					},
					{
						Name: "processes",
						Type: "integer",
					},
					{
						Name: "sessionClose",
						Type: "boolean",
					},
					{
						Name: "sessionID",
						Type: "string",
					},
					{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Registered",
							"Connected",
							"Disconnected",
						},
					},
					{
						Name: "ts",
						Type: "time",
					},
					{
						Name: "version",
						Type: "string",
					},
					{
						Name: "zhash",
						Type: "integer",
					},
				},
			},
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					{
						Name: "notify",
						Type: "boolean",
					},
					{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Initialized",
							"Paused",
							"Running",
							"Stopped",
						},
					},
					{
						Name: "ts",
						Type: "time",
					},
					{
						Name: "zhash",
						Type: "integer",
					},
				},
			},
		},
	}

	relationshipsRegistry[PolicyIdentity] = &elemental.Relationship{
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[PolicyGraphIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "view",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[PolicyRefreshIdentity] = &elemental.Relationship{}

	relationshipsRegistry[PolicyRendererIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PolicyRuleIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[PolicyTTLIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ProcessingUnitIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": {},
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"processingunitpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service":           {},
			"servicedependency": {},
			"vulnerability":     {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": {},
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"processingunitpolicy": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service":           {},
			"servicedependency": {},
			"vulnerability":     {},
		},
	}

	relationshipsRegistry[ProcessingUnitPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[ProcessingUnitRefreshIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"processingunit": {},
		},
	}

	relationshipsRegistry[QuotaCheckIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "remaining",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[QuotaPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[RecipeIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[RemoteProcessorIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[RenderTemplateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[RenderedPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "csr",
						Type: "string",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "csr",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "csr",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[ReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[RevocationIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[RoleIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}

	relationshipsRegistry[SAMLProviderIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[SSHAuthorityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[SSHAuthorizationPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[SSHCertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[SSHIdentityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[SandboxIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[SearchIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"q",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "q",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"q",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "q",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[ServiceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"processingunit": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"servicedependency": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"infrastructurepolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"networkaccesspolicy": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "object",
						AllowedChoices: []string{
							"subject",
							"object",
						},
					},
				},
			},
			"processingunit": {},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"servicedependency": {},
		},
	}

	relationshipsRegistry[ServiceDependencyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[ServiceTokenIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[SquallTagIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"identity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "identity",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"identity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "identity",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[StatsInfoIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[StatsQueryIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[SuggestedPolicyIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "filterAction",
						Type: "enum",
						AllowedChoices: []string{
							"include",
							"exclude",
						},
					},
					{
						Name:     "filterTags",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "flowOffset",
						Type: "duration",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"endRelative",
							},
							{
								"startRelative",
							},
							{
								"startRelative",
								"endRelative",
							},
							{
								"startRelative",
								"endAbsolute",
							},
							{
								"startAbsolute",
								"endRelative",
							},
							{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "filterAction",
						Type: "enum",
						AllowedChoices: []string{
							"include",
							"exclude",
						},
					},
					{
						Name:     "filterTags",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "endAbsolute",
						Type: "time",
					},
					{
						Name: "endRelative",
						Type: "duration",
					},
					{
						Name: "startAbsolute",
						Type: "time",
					},
					{
						Name: "startRelative",
						Type: "duration",
					},
					{
						Name: "flowOffset",
						Type: "duration",
					},
				},
			},
		},
	}

	relationshipsRegistry[TagIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[TagInjectIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[TagValueIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"key",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "key",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						{
							{
								"key",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "key",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[TextIndexIdentity] = &elemental.Relationship{}

	relationshipsRegistry[TokenIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[TokenScopePolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[TriggerIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"automation": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"automation": {},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"automation": {},
		},
	}

	relationshipsRegistry[TrustedCAIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
						},
					},
				},
			},
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
							"JWT",
						},
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
							"JWT",
						},
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
						},
					},
				},
			},
			"namespace": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
							"JWT",
						},
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:         "type",
						Type:         "enum",
						DefaultValue: "Any",
						AllowedChoices: []string{
							"Any",
							"X509",
							"SSH",
							"JWT",
						},
					},
				},
			},
		},
	}

	relationshipsRegistry[UserAccessPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[ValidateUIParameterIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[VulnerabilityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[X509CertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[X509CertificateCheckIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": {
				Parameters: []elemental.ParameterDefinition{
					{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

}
