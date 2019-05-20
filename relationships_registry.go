package gaia

import "go.aporeto.io/elemental"

var relationshipsRegistry elemental.RelationshipsRegistry

func init() {

	relationshipsRegistry = elemental.RelationshipsRegistry{}

	relationshipsRegistry[APIAuthorizationPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[APICheckIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[AWSAPIGatewayIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AWSAccountIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "accountid",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "accountid",
						Type: "string",
					},
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[AccountIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "name",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "status",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "name",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "status",
						Type: "string",
					},
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[ActivateIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"token",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "noRedirect",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"token",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "noRedirect",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[ActivityIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy": &elemental.RelationshipInfo{},
			"enforcer":                  &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy": &elemental.RelationshipInfo{},
			"enforcer":                  &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[AuditProfileMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[AuditReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[AuthIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[AuthorityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AutomationIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[CategoryIdentity] = &elemental.Relationship{}

	relationshipsRegistry[ClaimsIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[CustomerIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[DataPathCertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[DependencyMapIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"endRelative",
							},
							[]string{
								"startRelative",
							},
							[]string{
								"startRelative",
								"endRelative",
							},
							[]string{
								"startRelative",
								"endAbsolute",
							},
							[]string{
								"startAbsolute",
								"endRelative",
							},
							[]string{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "view",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "viewSuggestions",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "endAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "endRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "startAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "startRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "flowOffset",
						Type: "duration",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"endRelative",
							},
							[]string{
								"startRelative",
							},
							[]string{
								"startRelative",
								"endRelative",
							},
							[]string{
								"startRelative",
								"endAbsolute",
							},
							[]string{
								"startAbsolute",
								"endRelative",
							},
							[]string{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "view",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "viewSuggestions",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "endAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "endRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "startAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "startRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "flowOffset",
						Type: "duration",
					},
				},
			},
		},
	}

	relationshipsRegistry[EmailIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[EnforcerIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy":    &elemental.RelationshipInfo{},
			"enforcerprofilemappingpolicy": &elemental.RelationshipInfo{},
			"hostservicemappingpolicy":     &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"auditprofilemappingpolicy":    &elemental.RelationshipInfo{},
			"enforcerprofilemappingpolicy": &elemental.RelationshipInfo{},
			"hostservicemappingpolicy":     &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer":                     &elemental.RelationshipInfo{},
			"enforcerprofilemappingpolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer":                     &elemental.RelationshipInfo{},
			"enforcerprofilemappingpolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[EnforcerProfileMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[EnforcerReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[EnforcerTraceReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[EventLogIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[ExportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "ignoredTags",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ExternalNetworkIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"networkaccesspolicy": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"networkaccesspolicy": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[FileAccessPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[FileAccessReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[FilePathIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[FlowReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[HTTPResourceSpecIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[HookPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[HostServiceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "appliedServices",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "setServices",
						Type: "boolean",
					},
				},
			},
			"hostservicemappingpolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "appliedServices",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "setServices",
						Type: "boolean",
					},
				},
			},
			"hostservicemappingpolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[HostServiceMappingPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[IPInfoIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"ip",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "ip",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"ip",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "ip",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ImportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[ImportReferenceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "tag",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[InvoiceRecordIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[IsolationProfileIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunitpolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunitpolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[IssueIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "token",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[JaegerbatchIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[LDAPProviderIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"installedapp": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"installedapp": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[MessageIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[NamespaceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "authorized",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "authorized",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[OIDCProviderIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[PUNodeIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"endRelative",
							},
							[]string{
								"startRelative",
							},
							[]string{
								"startRelative",
								"endRelative",
							},
							[]string{
								"startRelative",
								"endAbsolute",
							},
							[]string{
								"startAbsolute",
								"endRelative",
							},
							[]string{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "endAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "endRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "startAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "startRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"endRelative",
							},
							[]string{
								"startRelative",
							},
							[]string{
								"startRelative",
								"endRelative",
							},
							[]string{
								"startRelative",
								"endAbsolute",
							},
							[]string{
								"startAbsolute",
								"endRelative",
							},
							[]string{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "endAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "endRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "startAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "startRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[PasswordResetIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"email",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "email",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"email",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "email",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[PlanIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[PokeIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "cpuload",
						Type: "float",
					},
					elemental.ParameterDefinition{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					elemental.ParameterDefinition{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "memory",
						Type: "integer",
					},
					elemental.ParameterDefinition{
						Name: "processes",
						Type: "integer",
					},
					elemental.ParameterDefinition{
						Name: "sessionClose",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "sessionID",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Registered",
							"Connected",
							"Disconnected",
						},
					},
					elemental.ParameterDefinition{
						Name: "ts",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "version",
						Type: "string",
					},
				},
			},
			"processingunit": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					elemental.ParameterDefinition{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "notify",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Initialized",
							"Paused",
							"Running",
							"Stopped",
						},
					},
					elemental.ParameterDefinition{
						Name: "ts",
						Type: "time",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "cpuload",
						Type: "float",
					},
					elemental.ParameterDefinition{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					elemental.ParameterDefinition{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "memory",
						Type: "integer",
					},
					elemental.ParameterDefinition{
						Name: "processes",
						Type: "integer",
					},
					elemental.ParameterDefinition{
						Name: "sessionClose",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "sessionID",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Registered",
							"Connected",
							"Disconnected",
						},
					},
					elemental.ParameterDefinition{
						Name: "ts",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "version",
						Type: "string",
					},
				},
			},
			"processingunit": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "enforcementStatus",
						Type: "enum",
						AllowedChoices: []string{
							"Failed",
							"Inactive",
							"Active",
						},
					},
					elemental.ParameterDefinition{
						Name: "forceFullPoke",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "notify",
						Type: "boolean",
					},
					elemental.ParameterDefinition{
						Name: "status",
						Type: "enum",
						AllowedChoices: []string{
							"Initialized",
							"Paused",
							"Running",
							"Stopped",
						},
					},
					elemental.ParameterDefinition{
						Name: "ts",
						Type: "time",
					},
				},
			},
		},
	}

	relationshipsRegistry[PolicyIdentity] = &elemental.Relationship{
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[PolicyGraphIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[PolicyRefreshIdentity] = &elemental.Relationship{}

	relationshipsRegistry[PolicyRendererIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[PolicyRuleIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[PolicyTTLIdentity] = &elemental.Relationship{}

	relationshipsRegistry[PrivateKeyIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[ProcessingUnitIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": &elemental.RelationshipInfo{},
			"networkaccesspolicy": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"processingunitpolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "objects",
						AllowedChoices: []string{
							"subjects",
							"object",
						},
					},
				},
			},
			"servicedependency": &elemental.RelationshipInfo{},
			"vulnerability":     &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"fileaccesspolicy": &elemental.RelationshipInfo{},
			"networkaccesspolicy": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"processingunitpolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"service": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "mode",
						Type:         "enum",
						DefaultValue: "objects",
						AllowedChoices: []string{
							"subjects",
							"object",
						},
					},
				},
			},
			"servicedependency": &elemental.RelationshipInfo{},
			"vulnerability":     &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[ProcessingUnitPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[ProcessingUnitRefreshIdentity] = &elemental.Relationship{}

	relationshipsRegistry[QuotaCheckIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[QuotaPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[RecipeIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[RemoteProcessorIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[RenderTemplateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[RenderedPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "csr",
						Type: "string",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "csr",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "csr",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[ReportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[RevocationIdentity] = &elemental.Relationship{
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[RootIdentity] = &elemental.Relationship{}

	relationshipsRegistry[SSHAuthorityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[SSHAuthorizationPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[SSHCertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[SSHIdentityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[ServiceIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"networkaccesspolicy": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"processingunit": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"servicedependency": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"networkaccesspolicy": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
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
			"processingunit": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "archived",
						Type: "boolean",
					},
				},
			},
			"servicedependency": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[ServiceDependencyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[SquallTagIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"identity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "identity",
						Type: "string",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"identity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "identity",
						Type: "string",
					},
				},
			},
		},
	}

	relationshipsRegistry[StatsInfoIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[StatsQueryIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"endRelative",
							},
							[]string{
								"startRelative",
							},
							[]string{
								"startRelative",
								"endRelative",
							},
							[]string{
								"startRelative",
								"endAbsolute",
							},
							[]string{
								"startAbsolute",
								"endRelative",
							},
							[]string{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "endAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "endRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "startAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "startRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
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
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"endRelative",
							},
							[]string{
								"startRelative",
							},
							[]string{
								"startRelative",
								"endRelative",
							},
							[]string{
								"startRelative",
								"endAbsolute",
							},
							[]string{
								"startAbsolute",
								"endRelative",
							},
							[]string{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "filterAction",
						Type: "enum",
						AllowedChoices: []string{
							"include",
							"exclude",
						},
					},
					elemental.ParameterDefinition{
						Name:     "filterTags",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "endAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "endRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "startAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "startRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "flowOffset",
						Type: "duration",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"endRelative",
							},
							[]string{
								"startRelative",
							},
							[]string{
								"startRelative",
								"endRelative",
							},
							[]string{
								"startRelative",
								"endAbsolute",
							},
							[]string{
								"startAbsolute",
								"endRelative",
							},
							[]string{
								"startAbsolute",
								"endAbsolute",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name: "filterAction",
						Type: "enum",
						AllowedChoices: []string{
							"include",
							"exclude",
						},
					},
					elemental.ParameterDefinition{
						Name:     "filterTags",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "endAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "endRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "startAbsolute",
						Type: "time",
					},
					elemental.ParameterDefinition{
						Name: "startRelative",
						Type: "duration",
					},
					elemental.ParameterDefinition{
						Name: "flowOffset",
						Type: "duration",
					},
				},
			},
		},
	}

	relationshipsRegistry[TabulationIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"identity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "column",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "identity",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"identity",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "column",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "identity",
						Type: "string",
					},
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[TagIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[TagInjectIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[TagValueIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"key",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "key",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"key",
							},
						},
					},
				),
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "key",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[TokenIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[TokenScopePolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Update: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Patch: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
					elemental.ParameterDefinition{
						Name: "propagated",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[TriggerIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"automation": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"automation": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"automation": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[ValidateUIParameterIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[VulnerabilityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

	relationshipsRegistry[X509CertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[X509CertificateCheckIdentity] = &elemental.Relationship{
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:     "q",
						Type:     "string",
						Multiple: true,
					},
				},
			},
		},
	}

}
