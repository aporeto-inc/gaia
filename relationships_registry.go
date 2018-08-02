package gaia

import "go.aporeto.io/elemental"

const nodocString = "[nodoc]" // nolint: varcheck,deadcode

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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "accountid",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "accountid",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "name",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "status",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "name",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "status",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AccountCheckIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
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
						Name:         "noRedirect",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "token",
						Type:         "string",
						DefaultValue: "<no value>",
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
						Name:         "noRedirect",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "token",
						Type:         "string",
						DefaultValue: "<no value>",
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
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcerprofile": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcerprofile": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
	}

	relationshipsRegistry[AuthIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[AuthorityIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Delete: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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

	relationshipsRegistry[CertificateIdentity] = &elemental.Relationship{
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
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
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
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "view",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "viewSuggestions",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "flowOffset",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
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
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "view",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "viewSuggestions",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "flowOffset",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
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
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcerprofilemappingpolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcerprofilemappingpolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"enforcer":                     &elemental.RelationshipInfo{},
			"enforcerprofilemappingpolicy": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
	}

	relationshipsRegistry[EventLogIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "category",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "id",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "identity",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "level",
						Type:         "string",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "category",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "id",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "identity",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "level",
						Type:         "string",
						DefaultValue: "<no value>",
					},
				},
			},
		},
	}

	relationshipsRegistry[ExportIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "ignoredTags",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
	}

	relationshipsRegistry[ExternalAccessIdentity] = &elemental.Relationship{
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
						Name:         "destinationID",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "destinationType",
						Type:         "enum",
						DefaultValue: "<no value>",
						AllowedChoices: []string{
							"ext",
							"pu",
						},
					},
					elemental.ParameterDefinition{
						Name:         "geoloc",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "resolve",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "sourceID",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "sourceType",
						Type:         "enum",
						DefaultValue: "<no value>",
						AllowedChoices: []string{
							"ext",
							"pu",
						},
					},
					elemental.ParameterDefinition{
						Name:         "endAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "flowOffset",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
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
						Name:         "destinationID",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "destinationType",
						Type:         "enum",
						DefaultValue: "<no value>",
						AllowedChoices: []string{
							"ext",
							"pu",
						},
					},
					elemental.ParameterDefinition{
						Name:         "geoloc",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "resolve",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "sourceID",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "sourceType",
						Type:         "enum",
						DefaultValue: "<no value>",
						AllowedChoices: []string{
							"ext",
							"pu",
						},
					},
					elemental.ParameterDefinition{
						Name:         "endAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "flowOffset",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
	}

	relationshipsRegistry[ExternalServiceIdentity] = &elemental.Relationship{
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
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
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
						DefaultValue: "objects",
						AllowedChoices: []string{
							"subjects",
							"object",
						},
					},
				},
			},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
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
						DefaultValue: "objects",
						AllowedChoices: []string{
							"subjects",
							"object",
						},
					},
				},
			},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
	}

	relationshipsRegistry[FileAccessIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"puID",
							},
						},
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
						Name:         "puID",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "flowOffset",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				RequiredParameters: elemental.NewParametersRequirement(
					[][][]string{
						[][]string{
							[]string{
								"puID",
							},
						},
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
						Name:         "puID",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "flowOffset",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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

	relationshipsRegistry[FlowStatisticIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Deprecated: true,
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
						Name:         "action",
						Type:         "enum",
						DefaultValue: "any",
						AllowedChoices: []string{
							"accept",
							"reject",
							"any",
						},
					},
					elemental.ParameterDefinition{
						Name:         "averageInterval",
						Type:         "duration",
						DefaultValue: "1h",
					},
					elemental.ParameterDefinition{
						Name:         "destinationID",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "flowMode",
						Type:         "enum",
						DefaultValue: "applied",
						AllowedChoices: []string{
							"observed",
							"applied",
							"all",
						},
					},
					elemental.ParameterDefinition{
						Name:         "metric",
						Type:         "enum",
						DefaultValue: "Flows",
						AllowedChoices: []string{
							"Flows",
							"Ports",
						},
					},
					elemental.ParameterDefinition{
						Name:         "sourceID",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "userIdentifier",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "flowOffset",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Deprecated: true,
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
						Name:         "action",
						Type:         "enum",
						DefaultValue: "any",
						AllowedChoices: []string{
							"accept",
							"reject",
							"any",
						},
					},
					elemental.ParameterDefinition{
						Name:         "averageInterval",
						Type:         "duration",
						DefaultValue: "1h",
					},
					elemental.ParameterDefinition{
						Name:         "destinationID",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "flowMode",
						Type:         "enum",
						DefaultValue: "applied",
						AllowedChoices: []string{
							"observed",
							"applied",
							"all",
						},
					},
					elemental.ParameterDefinition{
						Name:         "metric",
						Type:         "enum",
						DefaultValue: "Flows",
						AllowedChoices: []string{
							"Flows",
							"Ports",
						},
					},
					elemental.ParameterDefinition{
						Name:         "sourceID",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "userIdentifier",
						Type:         "string",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "flowOffset",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
				},
			},
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
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

	relationshipsRegistry[InstallationIdentity] = &elemental.Relationship{
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
			"root": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
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
			"root": &elemental.RelationshipInfo{},
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
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
						Name:         "token",
						Type:         "string",
						DefaultValue: "<no value>",
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

	relationshipsRegistry[K8SClusterIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Deprecated: true,
			},
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
				Deprecated: true,
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Deprecated: true,
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
	}

	relationshipsRegistry[KubernetesClusterIdentity] = &elemental.Relationship{
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
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
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
						Name:         "email",
						Type:         "string",
						DefaultValue: "<no value>",
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
						Name:         "email",
						Type:         "string",
						DefaultValue: "<no value>",
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
			"enforcer": &elemental.RelationshipInfo{},
			"processingunit": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "status",
						Type:         "enum",
						DefaultValue: "<no value>",
						AllowedChoices: []string{
							"Paused",
							"Running",
							"Stopped",
						},
					},
					elemental.ParameterDefinition{
						Name:         "ts",
						Type:         "time",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"enforcer": &elemental.RelationshipInfo{},
			"processingunit": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "status",
						Type:         "enum",
						DefaultValue: "<no value>",
						AllowedChoices: []string{
							"Paused",
							"Running",
							"Stopped",
						},
					},
					elemental.ParameterDefinition{
						Name:         "ts",
						Type:         "time",
						DefaultValue: "<no value>",
					},
				},
			},
		},
	}

	relationshipsRegistry[PolicyIdentity] = &elemental.Relationship{
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
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
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

	relationshipsRegistry[PrivateKeyIdentity] = &elemental.Relationship{
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
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
						DefaultValue: "objects",
						AllowedChoices: []string{
							"subjects",
							"object",
						},
					},
				},
			},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
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
			"networkaccesspolicy": &elemental.RelationshipInfo{
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
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
	}

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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
	}

	relationshipsRegistry[RESTAPISpecIdentity] = &elemental.Relationship{
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
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
			"service": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
			"service": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[RemoteProcessorIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
	}

	relationshipsRegistry[RenderedPolicyIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "csr",
						Type:         "string",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "csr",
						Type:         "string",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "csr",
						Type:         "string",
						DefaultValue: "<no value>",
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
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"processingunit": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
			"servicedependency": &elemental.RelationshipInfo{},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "archived",
						Type:         "boolean",
						DefaultValue: "<no value>",
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
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
						Name:         "identity",
						Type:         "string",
						DefaultValue: "<no value>",
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
						Name:         "identity",
						Type:         "string",
						DefaultValue: "<no value>",
					},
				},
			},
		},
	}

	relationshipsRegistry[StatsQueryIdentity] = &elemental.Relationship{
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
						Name:         "field",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "function",
						Type:         "enum",
						DefaultValue: "<no value>",
						AllowedChoices: []string{
							"count",
							"mean",
							"median",
							"sum",
							"first",
							"last",
							"max",
							"min",
						},
					},
					elemental.ParameterDefinition{
						Name:         "groupBy",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "interval",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "measurement",
						Type:         "enum",
						DefaultValue: "flows",
						AllowedChoices: []string{
							"flows",
							"audit",
							"enforcers",
							"files",
						},
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "where",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "endAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "flowOffset",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
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
						Name:         "field",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "function",
						Type:         "enum",
						DefaultValue: "<no value>",
						AllowedChoices: []string{
							"count",
							"mean",
							"median",
							"sum",
							"first",
							"last",
							"max",
							"min",
						},
					},
					elemental.ParameterDefinition{
						Name:         "groupBy",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "interval",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "measurement",
						Type:         "enum",
						DefaultValue: "flows",
						AllowedChoices: []string{
							"flows",
							"audit",
							"enforcers",
							"files",
						},
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "where",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "endAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "endRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "flowOffset",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startAbsolute",
						Type:         "time",
						DefaultValue: "<no value>",
					},
					elemental.ParameterDefinition{
						Name:         "startRelative",
						Type:         "duration",
						DefaultValue: "<no value>",
					},
				},
			},
		},
	}

	relationshipsRegistry[SuggestedPolicyIdentity] = &elemental.Relationship{
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "excludeTagPrefix",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "excludeTagPrefix",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
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
						Name:         "column",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "identity",
						Type:         "string",
						DefaultValue: "<no value>",
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
						Name:         "column",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "identity",
						Type:         "string",
						DefaultValue: "<no value>",
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
			"root": &elemental.RelationshipInfo{},
		},
		Retrieve: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "propagated",
						Type:         "boolean",
						DefaultValue: "<no value>",
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
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"processingunit": &elemental.RelationshipInfo{},
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
	}

	relationshipsRegistry[X509CertificateIdentity] = &elemental.Relationship{
		Create: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{},
		},
		RetrieveMany: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
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
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
		Info: map[string]*elemental.RelationshipInfo{
			"root": &elemental.RelationshipInfo{
				Parameters: []elemental.ParameterDefinition{
					elemental.ParameterDefinition{
						Name:         "q",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
					elemental.ParameterDefinition{
						Name:         "tag",
						Type:         "string",
						DefaultValue: "<no value>",
						Multiple:     true,
					},
				},
			},
		},
	}

}
