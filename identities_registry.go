package gaia

import "go.aporeto.io/elemental"

var (
	identityNamesMap = map[string]elemental.Identity{
		"accessreport":              AccessReportIdentity,
		"account":                   AccountIdentity,
		"accountcheck":              AccountCheckIdentity,
		"activate":                  ActivateIdentity,
		"activity":                  ActivityIdentity,
		"alarm":                     AlarmIdentity,
		"apiauthorizationpolicy":    APIAuthorizationPolicyIdentity,
		"apicheck":                  APICheckIdentity,
		"app":                       AppIdentity,
		"appcredential":             AppCredentialIdentity,
		"auditprofile":              AuditProfileIdentity,
		"auditprofilemappingpolicy": AuditProfileMappingPolicyIdentity,
		"auditreport":               AuditReportIdentity,
		"authn":                     AuthnIdentity,
		"authority":                 AuthorityIdentity,
		"authz":                     AuthzIdentity,
		"automation":                AutomationIdentity,
		"automationtemplate":        AutomationTemplateIdentity,
		"awsapigateway":             AWSAPIGatewayIdentity,
		"awsregister":               AWSRegisterIdentity,
		"category":                  CategoryIdentity,

		"claims":       ClaimsIdentity,
		"clausesmatch": ClauseMatchIdentity,

		"counterreport": CounterReportIdentity,

		"customer":            CustomerIdentity,
		"datapathcertificate": DataPathCertificateIdentity,
		"dependencymap":       DependencyMapIdentity,
		"dnslookupreport":     DNSLookupReportIdentity,
		"email":               EmailIdentity,

		"enforcer":                     EnforcerIdentity,
		"enforcerlog":                  EnforcerLogIdentity,
		"enforcerprofile":              EnforcerProfileIdentity,
		"enforcerprofilemappingpolicy": EnforcerProfileMappingPolicyIdentity,
		"enforcerreport":               EnforcerReportIdentity,
		"enforcertracereport":          EnforcerTraceReportIdentity,
		"eventlog":                     EventLogIdentity,
		"export":                       ExportIdentity,
		"externalnetwork":              ExternalNetworkIdentity,
		"fileaccesspolicy":             FileAccessPolicyIdentity,
		"fileaccessreport":             FileAccessReportIdentity,
		"filepath":                     FilePathIdentity,
		"flowreport":                   FlowReportIdentity,
		"graphedge":                    GraphEdgeIdentity,

		"graphnode":                GraphNodeIdentity,
		"hit":                      HitIdentity,
		"hookpolicy":               HookPolicyIdentity,
		"hostservice":              HostServiceIdentity,
		"hostservicemappingpolicy": HostServiceMappingPolicyIdentity,
		"httpresourcespec":         HTTPResourceSpecIdentity,
		"import":                   ImportIdentity,
		"importreference":          ImportReferenceIdentity,
		"importrequest":            ImportRequestIdentity,
		"infrastructurepolicy":     InfrastructurePolicyIdentity,
		"installedapp":             InstalledAppIdentity,
		"invoice":                  InvoiceIdentity,
		"invoicerecord":            InvoiceRecordIdentity,
		"ipinfo":                   IPInfoIdentity,
		"isolationprofile":         IsolationProfileIdentity,
		"issue":                    IssueIdentity,
		"issueservicetoken":        IssueServiceTokenIdentity,

		"ldapprovider":           LDAPProviderIdentity,
		"log":                    LogIdentity,
		"message":                MessageIdentity,
		"namespace":              NamespaceIdentity,
		"namespacemappingpolicy": NamespaceMappingPolicyIdentity,
		"networkaccesspolicy":    NetworkAccessPolicyIdentity,
		"oauthinfo":              OAUTHInfoIdentity,
		"oauthkey":               OAUTHKeyIdentity,
		"oidcprovider":           OIDCProviderIdentity,
		"packetreport":           PacketReportIdentity,
		"passwordreset":          PasswordResetIdentity,

		"plan":                  PlanIdentity,
		"poke":                  PokeIdentity,
		"policy":                PolicyIdentity,
		"policygraph":           PolicyGraphIdentity,
		"policyrefresh":         PolicyRefreshIdentity,
		"policyrenderer":        PolicyRendererIdentity,
		"policyrule":            PolicyRuleIdentity,
		"policyttl":             PolicyTTLIdentity,
		"processingunit":        ProcessingUnitIdentity,
		"processingunitpolicy":  ProcessingUnitPolicyIdentity,
		"processingunitrefresh": ProcessingUnitRefreshIdentity,

		"quotacheck":  QuotaCheckIdentity,
		"quotapolicy": QuotaPolicyIdentity,
		"recipe":      RecipeIdentity,

		"remoteprocessor":        RemoteProcessorIdentity,
		"renderedpolicy":         RenderedPolicyIdentity,
		"rendertemplate":         RenderTemplateIdentity,
		"report":                 ReportIdentity,
		"revocation":             RevocationIdentity,
		"role":                   RoleIdentity,
		"root":                   RootIdentity,
		"samlprovider":           SAMLProviderIdentity,
		"sandbox":                SandboxIdentity,
		"search":                 SearchIdentity,
		"service":                ServiceIdentity,
		"servicedependency":      ServiceDependencyIdentity,
		"servicetoken":           ServiceTokenIdentity,
		"squalltag":              SquallTagIdentity,
		"sshauthority":           SSHAuthorityIdentity,
		"sshauthorizationpolicy": SSHAuthorizationPolicyIdentity,
		"sshcertificate":         SSHCertificateIdentity,
		"sshidentity":            SSHIdentityIdentity,
		"statsinfo":              StatsInfoIdentity,
		"statsquery":             StatsQueryIdentity,
		"suggestedpolicy":        SuggestedPolicyIdentity,
		"tabulation":             TabulationIdentity,
		"tag":                    TagIdentity,
		"taginject":              TagInjectIdentity,
		"tagvalue":               TagValueIdentity,
		"textindex":              TextIndexIdentity,

		"token":            TokenIdentity,
		"tokenscopepolicy": TokenScopePolicyIdentity,

		"trigger":   TriggerIdentity,
		"trustedca": TrustedCAIdentity,

		"useraccesspolicy":     UserAccessPolicyIdentity,
		"validateuiparameter":  ValidateUIParameterIdentity,
		"vulnerability":        VulnerabilityIdentity,
		"x509certificate":      X509CertificateIdentity,
		"x509certificatecheck": X509CertificateCheckIdentity,
	}

	identitycategoriesMap = map[string]elemental.Identity{
		"accessreports":               AccessReportIdentity,
		"accounts":                    AccountIdentity,
		"accountchecks":               AccountCheckIdentity,
		"activate":                    ActivateIdentity,
		"activities":                  ActivityIdentity,
		"alarms":                      AlarmIdentity,
		"apiauthorizationpolicies":    APIAuthorizationPolicyIdentity,
		"apichecks":                   APICheckIdentity,
		"apps":                        AppIdentity,
		"appcredentials":              AppCredentialIdentity,
		"auditprofiles":               AuditProfileIdentity,
		"auditprofilemappingpolicies": AuditProfileMappingPolicyIdentity,
		"auditreports":                AuditReportIdentity,
		"authn":                       AuthnIdentity,
		"authorities":                 AuthorityIdentity,
		"authz":                       AuthzIdentity,
		"automations":                 AutomationIdentity,
		"automationtemplates":         AutomationTemplateIdentity,
		"awsapigateways":              AWSAPIGatewayIdentity,
		"awsregister":                 AWSRegisterIdentity,
		"categories":                  CategoryIdentity,

		"claims":         ClaimsIdentity,
		"clausesmatches": ClauseMatchIdentity,

		"counterreports": CounterReportIdentity,

		"customers":            CustomerIdentity,
		"datapathcertificates": DataPathCertificateIdentity,
		"dependencymaps":       DependencyMapIdentity,
		"dnslookupreports":     DNSLookupReportIdentity,
		"emails":               EmailIdentity,

		"enforcers":                      EnforcerIdentity,
		"enforcerlog":                    EnforcerLogIdentity,
		"enforcerprofiles":               EnforcerProfileIdentity,
		"enforcerprofilemappingpolicies": EnforcerProfileMappingPolicyIdentity,
		"enforcerreports":                EnforcerReportIdentity,
		"enforcertracereports":           EnforcerTraceReportIdentity,
		"eventlogs":                      EventLogIdentity,
		"export":                         ExportIdentity,
		"externalnetworks":               ExternalNetworkIdentity,
		"fileaccesspolicies":             FileAccessPolicyIdentity,
		"fileaccessreports":              FileAccessReportIdentity,
		"filepaths":                      FilePathIdentity,
		"flowreports":                    FlowReportIdentity,
		"graphedges":                     GraphEdgeIdentity,

		"graphnodes":                 GraphNodeIdentity,
		"hits":                       HitIdentity,
		"hookpolicies":               HookPolicyIdentity,
		"hostservices":               HostServiceIdentity,
		"hostservicemappingpolicies": HostServiceMappingPolicyIdentity,
		"httpresourcespecs":          HTTPResourceSpecIdentity,
		"import":                     ImportIdentity,
		"importreferences":           ImportReferenceIdentity,
		"importrequests":             ImportRequestIdentity,
		"infrastructurepolicies":     InfrastructurePolicyIdentity,
		"installedapps":              InstalledAppIdentity,
		"invoices":                   InvoiceIdentity,
		"invoicerecords":             InvoiceRecordIdentity,
		"ipinfos":                    IPInfoIdentity,
		"isolationprofiles":          IsolationProfileIdentity,
		"issue":                      IssueIdentity,
		"issueservicetokens":         IssueServiceTokenIdentity,

		"ldapproviders":            LDAPProviderIdentity,
		"logs":                     LogIdentity,
		"messages":                 MessageIdentity,
		"namespaces":               NamespaceIdentity,
		"namespacemappingpolicies": NamespaceMappingPolicyIdentity,
		"networkaccesspolicies":    NetworkAccessPolicyIdentity,
		"oauthinfo":                OAUTHInfoIdentity,
		"oauthkeys":                OAUTHKeyIdentity,
		"oidcproviders":            OIDCProviderIdentity,
		"packetreports":            PacketReportIdentity,
		"passwordreset":            PasswordResetIdentity,

		"plans":                  PlanIdentity,
		"poke":                   PokeIdentity,
		"policies":               PolicyIdentity,
		"policygraphs":           PolicyGraphIdentity,
		"policyrefreshs":         PolicyRefreshIdentity,
		"policyrenderers":        PolicyRendererIdentity,
		"policyrules":            PolicyRuleIdentity,
		"policyttls":             PolicyTTLIdentity,
		"processingunits":        ProcessingUnitIdentity,
		"processingunitpolicies": ProcessingUnitPolicyIdentity,
		"processingunitrefreshs": ProcessingUnitRefreshIdentity,

		"quotacheck":    QuotaCheckIdentity,
		"quotapolicies": QuotaPolicyIdentity,
		"recipes":       RecipeIdentity,

		"remoteprocessors":         RemoteProcessorIdentity,
		"renderedpolicies":         RenderedPolicyIdentity,
		"rendertemplates":          RenderTemplateIdentity,
		"reports":                  ReportIdentity,
		"revocations":              RevocationIdentity,
		"roles":                    RoleIdentity,
		"root":                     RootIdentity,
		"samlproviders":            SAMLProviderIdentity,
		"sandboxes":                SandboxIdentity,
		"search":                   SearchIdentity,
		"services":                 ServiceIdentity,
		"servicedependencies":      ServiceDependencyIdentity,
		"servicetoken":             ServiceTokenIdentity,
		"squalltags":               SquallTagIdentity,
		"sshauthorities":           SSHAuthorityIdentity,
		"sshauthorizationpolicies": SSHAuthorizationPolicyIdentity,
		"sshcertificates":          SSHCertificateIdentity,
		"sshidentities":            SSHIdentityIdentity,
		"statsinfo":                StatsInfoIdentity,
		"statsqueries":             StatsQueryIdentity,
		"suggestedpolicies":        SuggestedPolicyIdentity,
		"tabulations":              TabulationIdentity,
		"tags":                     TagIdentity,
		"taginjects":               TagInjectIdentity,
		"tagvalues":                TagValueIdentity,
		"textindexes":              TextIndexIdentity,

		"tokens":             TokenIdentity,
		"tokenscopepolicies": TokenScopePolicyIdentity,

		"triggers":   TriggerIdentity,
		"trustedcas": TrustedCAIdentity,

		"useraccesspolicies":    UserAccessPolicyIdentity,
		"validateuiparameters":  ValidateUIParameterIdentity,
		"vulnerabilities":       VulnerabilityIdentity,
		"x509certificates":      X509CertificateIdentity,
		"x509certificatechecks": X509CertificateCheckIdentity,
	}

	aliasesMap = map[string]elemental.Identity{
		"apiauth":        APIAuthorizationPolicyIdentity,
		"apiauths":       APIAuthorizationPolicyIdentity,
		"appcred":        AppCredentialIdentity,
		"appcreds":       AppCredentialIdentity,
		"ap":             AuditProfileIdentity,
		"audpol":         AuditProfileMappingPolicyIdentity,
		"audpols":        AuditProfileMappingPolicyIdentity,
		"ca":             AuthorityIdentity,
		"autos":          AutomationIdentity,
		"auto":           AutomationIdentity,
		"autotmpl":       AutomationTemplateIdentity,
		"depmaps":        DependencyMapIdentity,
		"depmap":         DependencyMapIdentity,
		"profile":        EnforcerProfileIdentity,
		"profiles":       EnforcerProfileIdentity,
		"enfpols":        EnforcerProfileMappingPolicyIdentity,
		"enfpol":         EnforcerProfileMappingPolicyIdentity,
		"epm":            EnforcerProfileMappingPolicyIdentity,
		"extnet":         ExternalNetworkIdentity,
		"extnets":        ExternalNetworkIdentity,
		"fp":             FilePathIdentity,
		"fps":            FilePathIdentity,
		"hook":           HookPolicyIdentity,
		"hooks":          HookPolicyIdentity,
		"hookpol":        HookPolicyIdentity,
		"hookpols":       HookPolicyIdentity,
		"hostsrv":        HostServiceIdentity,
		"hostsrvs":       HostServiceIdentity,
		"hostsrvmappol":  HostServiceMappingPolicyIdentity,
		"hostsrvmappols": HostServiceMappingPolicyIdentity,
		"httpresource":   HTTPResourceSpecIdentity,
		"resource":       HTTPResourceSpecIdentity,
		"httpspec":       HTTPResourceSpecIdentity,
		"importref":      ImportReferenceIdentity,
		"impref":         ImportReferenceIdentity,
		"req":            ImportRequestIdentity,
		"reqs":           ImportRequestIdentity,
		"ireq":           ImportRequestIdentity,
		"ireqs":          ImportRequestIdentity,
		"infrapol":       InfrastructurePolicyIdentity,
		"infrapols":      InfrastructurePolicyIdentity,
		"iapps":          InstalledAppIdentity,
		"iapp":           InstalledAppIdentity,
		"ip":             IsolationProfileIdentity,
		"mess":           MessageIdentity,
		"ns":             NamespaceIdentity,
		"nspolicy":       NamespaceMappingPolicyIdentity,
		"nspolicies":     NamespaceMappingPolicyIdentity,
		"nsmap":          NamespaceMappingPolicyIdentity,
		"nsmaps":         NamespaceMappingPolicyIdentity,
		"netpol":         NetworkAccessPolicyIdentity,
		"netpols":        NetworkAccessPolicyIdentity,
		"polgraph":       PolicyGraphIdentity,
		"pu":             ProcessingUnitIdentity,
		"pus":            ProcessingUnitIdentity,
		"pup":            ProcessingUnitPolicyIdentity,
		"pups":           ProcessingUnitPolicyIdentity,
		"quota":          QuotaPolicyIdentity,
		"quotas":         QuotaPolicyIdentity,
		"quotapol":       QuotaPolicyIdentity,
		"quotapols":      QuotaPolicyIdentity,
		"rcp":            RecipeIdentity,
		"hks":            RemoteProcessorIdentity,
		"hk":             RemoteProcessorIdentity,
		"rpol":           RenderedPolicyIdentity,
		"rpols":          RenderedPolicyIdentity,
		"cook":           RenderTemplateIdentity,
		"rtpl":           RenderTemplateIdentity,
		"srv":            ServiceIdentity,
		"srvdep":         ServiceDependencyIdentity,
		"srvdeps":        ServiceDependencyIdentity,
		"sshpol":         SSHAuthorizationPolicyIdentity,
		"sshpols":        SSHAuthorizationPolicyIdentity,
		"si":             StatsInfoIdentity,
		"sq":             StatsQueryIdentity,
		"sugpol":         SuggestedPolicyIdentity,
		"sugpols":        SuggestedPolicyIdentity,
		"sugg":           SuggestedPolicyIdentity,
		"suggs":          SuggestedPolicyIdentity,
		"table":          TabulationIdentity,
		"tables":         TabulationIdentity,
		"tabs":           TabulationIdentity,
		"tab":            TabulationIdentity,
		"tsp":            TokenScopePolicyIdentity,
		"usrpol":         UserAccessPolicyIdentity,
		"usrpols":        UserAccessPolicyIdentity,
		"validparam":     ValidateUIParameterIdentity,
		"vulns":          VulnerabilityIdentity,
		"vul":            VulnerabilityIdentity,
		"vuln":           VulnerabilityIdentity,
		"vuls":           VulnerabilityIdentity,
	}

	indexesMap = map[string][][]string{
		"accessreport": nil,
		"account": [][]string{
			[]string{"resetPasswordToken"},
			[]string{"name"},
			[]string{"email"},
			[]string{"activationToken"},
			[]string{":shard", ":unique", "zone", "zHash"},
		},
		"accountcheck": nil,
		"activate":     nil,
		"activity": [][]string{
			[]string{"namespace", "date"},
			[]string{"namespace", "operation"},
			[]string{"namespace", "error.code"},
			[]string{"namespace", "targetIdentity"},
			[]string{"namespace", "data.ID"},
			[]string{"namespace", "originalData.ID"},
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
		},
		"alarm": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"status"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "kind"},
			[]string{"namespace", "kind", "status"},
			[]string{"name"},
			[]string{"kind"},
			[]string{"createIdempotencyKey"},
		},
		"apiauthorizationpolicy": nil,
		"apicheck":               nil,
		"app":                    nil,
		"appcredential": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "disabled"},
			[]string{"name"},
			[]string{"disabled"},
			[]string{"createIdempotencyKey"},
		},
		"auditprofile": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"propagate"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"auditprofilemappingpolicy": nil,
		"auditreport":               nil,
		"authn":                     nil,
		"authority": [][]string{
			[]string{"commonName"},
			[]string{":shard", ":unique", "zone", "zHash"},
		},
		"authz": nil,
		"automation": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "disabled"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"disabled"},
			[]string{"createIdempotencyKey"},
		},
		"automationtemplate": nil,
		"awsapigateway": [][]string{
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "name"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
			[]string{":shard", ":unique", "zone", "zHash"},
		},
		"awsregister": nil,
		"category":    nil,
		"claims": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"createIdempotencyKey"},
		},
		"clausesmatch":  nil,
		"counterreport": nil,
		"customer": [][]string{
			[]string{"providerCustomerID"},
		},
		"datapathcertificate": nil,
		"dependencymap":       nil,
		"dnslookupreport":     nil,
		"email":               nil,
		"enforcer": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "machineID"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "operationalStatus"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"enforcerlog": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "collectionID"},
			[]string{"namespace", "enforcerID"},
			[]string{"enforcerID"},
			[]string{"createIdempotencyKey"},
			[]string{"collectionID"},
		},
		"enforcerprofile": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"propagate"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"enforcerprofilemappingpolicy": nil,
		"enforcerreport":               nil,
		"enforcertracereport":          nil,
		"eventlog":                     nil,
		"export":                       nil,
		"externalnetwork": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"type"},
			[]string{"propagate"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "archived"},
			[]string{"namespace", "type"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
			[]string{"archived"},
		},
		"fileaccesspolicy": nil,
		"fileaccessreport": nil,
		"filepath": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"propagate"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "archived"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
			[]string{"archived"},
		},
		"flowreport": nil,
		"graphedge": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "lastSeen", "firstSeen"},
			[]string{"lastSeen", "firstSeen"},
			[]string{"lastSeen"},
			[]string{"flowID", "bucketMonth"},
			[]string{"flowID", "lastSeen"},
			[]string{"flowID", "bucketMinute"},
			[]string{"flowID", "bucketHour"},
			[]string{"flowID"},
			[]string{"flowID", "bucketDay"},
			[]string{"firstSeen"},
		},
		"graphnode":  nil,
		"hit":        nil,
		"hookpolicy": nil,
		"hostservice": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"propagate"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "archived"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
			[]string{"archived"},
		},
		"hostservicemappingpolicy": nil,
		"httpresourcespec": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"propagate"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "archived"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
			[]string{"archived"},
		},
		"import": nil,
		"importreference": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"importrequest": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"createIdempotencyKey"},
		},
		"infrastructurepolicy": nil,
		"installedapp": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"invoice":       nil,
		"invoicerecord": nil,
		"ipinfo":        nil,
		"isolationprofile": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"propagate"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"issue":             nil,
		"issueservicetoken": nil,
		"ldapprovider": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"log": nil,
		"message": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"propagate"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"namespace": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"namespacemappingpolicy": nil,
		"networkaccesspolicy":    nil,
		"oauthinfo":              nil,
		"oauthkey":               nil,
		"oidcprovider": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"packetreport":  nil,
		"passwordreset": nil,
		"plan":          nil,
		"poke":          nil,
		"policy": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"propagate"},
			[]string{"namespace", "type", "allSubjectTags", "propagate"},
			[]string{"namespace", "type", "allObjectTags"},
			[]string{"namespace", "type", "allSubjectTags"},
			[]string{"namespace", "type", "allObjectTags", "disabled"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "type"},
			[]string{"namespace", "disabled"},
			[]string{"namespace", "name"},
			[]string{"namespace", "type", "allObjectTags", "propagate"},
			[]string{"namespace"},
			[]string{"namespace", "type", "allSubjectTags", "disabled"},
			[]string{"namespace", "fallback"},
			[]string{"name"},
			[]string{"fallback"},
			[]string{"disabled"},
			[]string{"createIdempotencyKey"},
		},
		"policygraph":    nil,
		"policyrefresh":  nil,
		"policyrenderer": nil,
		"policyrule":     nil,
		"policyttl":      nil,
		"processingunit": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "archived", "createTime", "lastPokeTime"},
			[]string{"namespace", "nativeContextID"},
			[]string{"namespace", "archived"},
			[]string{"namespace", "name"},
			[]string{"namespace", "normalizedTags", "archived"},
			[]string{"namespace", "operationalStatus", "archived"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace"},
			[]string{"name"},
			[]string{"enforcerID"},
			[]string{"createIdempotencyKey"},
			[]string{"archived"},
		},
		"processingunitpolicy":  nil,
		"processingunitrefresh": nil,
		"quotacheck":            nil,
		"quotapolicy":           nil,
		"recipe": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"propagate"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"remoteprocessor": nil,
		"renderedpolicy":  nil,
		"rendertemplate":  nil,
		"report":          nil,
		"revocation": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
		},
		"role": nil,
		"root": nil,
		"samlprovider": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
		},
		"sandbox": nil,
		"search":  nil,
		"service": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"namespace", "allAPITags"},
			[]string{"namespace", "allServiceTags"},
			[]string{"namespace", "disabled"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "archived"},
			[]string{"name"},
			[]string{"disabled"},
			[]string{"createIdempotencyKey"},
			[]string{"archived"},
			[]string{"allServiceTags"},
			[]string{"allAPITags"},
		},
		"servicedependency": nil,
		"servicetoken":      nil,
		"squalltag":         nil,
		"sshauthority": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"namespace", "name"},
			[]string{"name"},
		},
		"sshauthorizationpolicy": nil,
		"sshcertificate":         nil,
		"sshidentity":            nil,
		"statsinfo":              nil,
		"statsquery":             nil,
		"suggestedpolicy":        nil,
		"tabulation":             nil,
		"tag": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
		},
		"taginject": nil,
		"tagvalue":  nil,
		"textindex": [][]string{
			[]string{"objectNamespace"},
			[]string{"objectNamespace", "objectIdentity", "objectID"},
			[]string{"objectIdentity"},
			[]string{"objectID"},
			[]string{"date"},
			[]string{":shard", ":unique", "zone", "zHash"},
		},
		"token":               nil,
		"tokenscopepolicy":    nil,
		"trigger":             nil,
		"trustedca":           nil,
		"useraccesspolicy":    nil,
		"validateuiparameter": nil,
		"vulnerability": [][]string{
			[]string{":shard", ":unique", "zone", "zHash"},
			[]string{"updateIdempotencyKey"},
			[]string{"severity"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "CVSS2Score"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "severity"},
			[]string{"name"},
			[]string{"createIdempotencyKey"},
			[]string{"CVSS2Score"},
		},
		"x509certificate":      nil,
		"x509certificatecheck": nil,
	}
)

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

type modelManager struct{}

func (f modelManager) IdentityFromName(name string) elemental.Identity {

	return identityNamesMap[name]
}

func (f modelManager) IdentityFromCategory(category string) elemental.Identity {

	return identitycategoriesMap[category]
}

func (f modelManager) IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

func (f modelManager) IdentityFromAny(any string) (i elemental.Identity) {

	if i = f.IdentityFromName(any); !i.IsEmpty() {
		return i
	}

	if i = f.IdentityFromCategory(any); !i.IsEmpty() {
		return i
	}

	return f.IdentityFromAlias(any)
}

func (f modelManager) Identifiable(identity elemental.Identity) elemental.Identifiable {

	switch identity {

	case AccessReportIdentity:
		return NewAccessReport()
	case AccountIdentity:
		return NewAccount()
	case AccountCheckIdentity:
		return NewAccountCheck()
	case ActivateIdentity:
		return NewActivate()
	case ActivityIdentity:
		return NewActivity()
	case AlarmIdentity:
		return NewAlarm()
	case APIAuthorizationPolicyIdentity:
		return NewAPIAuthorizationPolicy()
	case APICheckIdentity:
		return NewAPICheck()
	case AppIdentity:
		return NewApp()
	case AppCredentialIdentity:
		return NewAppCredential()
	case AuditProfileIdentity:
		return NewAuditProfile()
	case AuditProfileMappingPolicyIdentity:
		return NewAuditProfileMappingPolicy()
	case AuditReportIdentity:
		return NewAuditReport()
	case AuthnIdentity:
		return NewAuthn()
	case AuthorityIdentity:
		return NewAuthority()
	case AuthzIdentity:
		return NewAuthz()
	case AutomationIdentity:
		return NewAutomation()
	case AutomationTemplateIdentity:
		return NewAutomationTemplate()
	case AWSAPIGatewayIdentity:
		return NewAWSAPIGateway()
	case AWSRegisterIdentity:
		return NewAWSRegister()
	case CategoryIdentity:
		return NewCategory()
	case ClaimsIdentity:
		return NewClaims()
	case ClauseMatchIdentity:
		return NewClauseMatch()
	case CounterReportIdentity:
		return NewCounterReport()
	case CustomerIdentity:
		return NewCustomer()
	case DataPathCertificateIdentity:
		return NewDataPathCertificate()
	case DependencyMapIdentity:
		return NewDependencyMap()
	case DNSLookupReportIdentity:
		return NewDNSLookupReport()
	case EmailIdentity:
		return NewEmail()
	case EnforcerIdentity:
		return NewEnforcer()
	case EnforcerLogIdentity:
		return NewEnforcerLog()
	case EnforcerProfileIdentity:
		return NewEnforcerProfile()
	case EnforcerProfileMappingPolicyIdentity:
		return NewEnforcerProfileMappingPolicy()
	case EnforcerReportIdentity:
		return NewEnforcerReport()
	case EnforcerTraceReportIdentity:
		return NewEnforcerTraceReport()
	case EventLogIdentity:
		return NewEventLog()
	case ExportIdentity:
		return NewExport()
	case ExternalNetworkIdentity:
		return NewExternalNetwork()
	case FileAccessPolicyIdentity:
		return NewFileAccessPolicy()
	case FileAccessReportIdentity:
		return NewFileAccessReport()
	case FilePathIdentity:
		return NewFilePath()
	case FlowReportIdentity:
		return NewFlowReport()
	case GraphEdgeIdentity:
		return NewGraphEdge()
	case GraphNodeIdentity:
		return NewGraphNode()
	case HitIdentity:
		return NewHit()
	case HookPolicyIdentity:
		return NewHookPolicy()
	case HostServiceIdentity:
		return NewHostService()
	case HostServiceMappingPolicyIdentity:
		return NewHostServiceMappingPolicy()
	case HTTPResourceSpecIdentity:
		return NewHTTPResourceSpec()
	case ImportIdentity:
		return NewImport()
	case ImportReferenceIdentity:
		return NewImportReference()
	case ImportRequestIdentity:
		return NewImportRequest()
	case InfrastructurePolicyIdentity:
		return NewInfrastructurePolicy()
	case InstalledAppIdentity:
		return NewInstalledApp()
	case InvoiceIdentity:
		return NewInvoice()
	case InvoiceRecordIdentity:
		return NewInvoiceRecord()
	case IPInfoIdentity:
		return NewIPInfo()
	case IsolationProfileIdentity:
		return NewIsolationProfile()
	case IssueIdentity:
		return NewIssue()
	case IssueServiceTokenIdentity:
		return NewIssueServiceToken()
	case LDAPProviderIdentity:
		return NewLDAPProvider()
	case LogIdentity:
		return NewLog()
	case MessageIdentity:
		return NewMessage()
	case NamespaceIdentity:
		return NewNamespace()
	case NamespaceMappingPolicyIdentity:
		return NewNamespaceMappingPolicy()
	case NetworkAccessPolicyIdentity:
		return NewNetworkAccessPolicy()
	case OAUTHInfoIdentity:
		return NewOAUTHInfo()
	case OAUTHKeyIdentity:
		return NewOAUTHKey()
	case OIDCProviderIdentity:
		return NewOIDCProvider()
	case PacketReportIdentity:
		return NewPacketReport()
	case PasswordResetIdentity:
		return NewPasswordReset()
	case PlanIdentity:
		return NewPlan()
	case PokeIdentity:
		return NewPoke()
	case PolicyIdentity:
		return NewPolicy()
	case PolicyGraphIdentity:
		return NewPolicyGraph()
	case PolicyRefreshIdentity:
		return NewPolicyRefresh()
	case PolicyRendererIdentity:
		return NewPolicyRenderer()
	case PolicyRuleIdentity:
		return NewPolicyRule()
	case PolicyTTLIdentity:
		return NewPolicyTTL()
	case ProcessingUnitIdentity:
		return NewProcessingUnit()
	case ProcessingUnitPolicyIdentity:
		return NewProcessingUnitPolicy()
	case ProcessingUnitRefreshIdentity:
		return NewProcessingUnitRefresh()
	case QuotaCheckIdentity:
		return NewQuotaCheck()
	case QuotaPolicyIdentity:
		return NewQuotaPolicy()
	case RecipeIdentity:
		return NewRecipe()
	case RemoteProcessorIdentity:
		return NewRemoteProcessor()
	case RenderedPolicyIdentity:
		return NewRenderedPolicy()
	case RenderTemplateIdentity:
		return NewRenderTemplate()
	case ReportIdentity:
		return NewReport()
	case RevocationIdentity:
		return NewRevocation()
	case RoleIdentity:
		return NewRole()
	case RootIdentity:
		return NewRoot()
	case SAMLProviderIdentity:
		return NewSAMLProvider()
	case SandboxIdentity:
		return NewSandbox()
	case SearchIdentity:
		return NewSearch()
	case ServiceIdentity:
		return NewService()
	case ServiceDependencyIdentity:
		return NewServiceDependency()
	case ServiceTokenIdentity:
		return NewServiceToken()
	case SquallTagIdentity:
		return NewSquallTag()
	case SSHAuthorityIdentity:
		return NewSSHAuthority()
	case SSHAuthorizationPolicyIdentity:
		return NewSSHAuthorizationPolicy()
	case SSHCertificateIdentity:
		return NewSSHCertificate()
	case SSHIdentityIdentity:
		return NewSSHIdentity()
	case StatsInfoIdentity:
		return NewStatsInfo()
	case StatsQueryIdentity:
		return NewStatsQuery()
	case SuggestedPolicyIdentity:
		return NewSuggestedPolicy()
	case TabulationIdentity:
		return NewTabulation()
	case TagIdentity:
		return NewTag()
	case TagInjectIdentity:
		return NewTagInject()
	case TagValueIdentity:
		return NewTagValue()
	case TextIndexIdentity:
		return NewTextIndex()
	case TokenIdentity:
		return NewToken()
	case TokenScopePolicyIdentity:
		return NewTokenScopePolicy()
	case TriggerIdentity:
		return NewTrigger()
	case TrustedCAIdentity:
		return NewTrustedCA()
	case UserAccessPolicyIdentity:
		return NewUserAccessPolicy()
	case ValidateUIParameterIdentity:
		return NewValidateUIParameter()
	case VulnerabilityIdentity:
		return NewVulnerability()
	case X509CertificateIdentity:
		return NewX509Certificate()
	case X509CertificateCheckIdentity:
		return NewX509CertificateCheck()
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiable(identity elemental.Identity) elemental.SparseIdentifiable {

	switch identity {

	case AccessReportIdentity:
		return NewSparseAccessReport()
	case AccountIdentity:
		return NewSparseAccount()
	case AccountCheckIdentity:
		return NewSparseAccountCheck()
	case ActivateIdentity:
		return NewSparseActivate()
	case ActivityIdentity:
		return NewSparseActivity()
	case AlarmIdentity:
		return NewSparseAlarm()
	case APIAuthorizationPolicyIdentity:
		return NewSparseAPIAuthorizationPolicy()
	case APICheckIdentity:
		return NewSparseAPICheck()
	case AppIdentity:
		return NewSparseApp()
	case AppCredentialIdentity:
		return NewSparseAppCredential()
	case AuditProfileIdentity:
		return NewSparseAuditProfile()
	case AuditProfileMappingPolicyIdentity:
		return NewSparseAuditProfileMappingPolicy()
	case AuditReportIdentity:
		return NewSparseAuditReport()
	case AuthnIdentity:
		return NewSparseAuthn()
	case AuthorityIdentity:
		return NewSparseAuthority()
	case AuthzIdentity:
		return NewSparseAuthz()
	case AutomationIdentity:
		return NewSparseAutomation()
	case AutomationTemplateIdentity:
		return NewSparseAutomationTemplate()
	case AWSAPIGatewayIdentity:
		return NewSparseAWSAPIGateway()
	case AWSRegisterIdentity:
		return NewSparseAWSRegister()
	case CategoryIdentity:
		return NewSparseCategory()
	case ClaimsIdentity:
		return NewSparseClaims()
	case ClauseMatchIdentity:
		return NewSparseClauseMatch()
	case CounterReportIdentity:
		return NewSparseCounterReport()
	case CustomerIdentity:
		return NewSparseCustomer()
	case DataPathCertificateIdentity:
		return NewSparseDataPathCertificate()
	case DependencyMapIdentity:
		return NewSparseDependencyMap()
	case DNSLookupReportIdentity:
		return NewSparseDNSLookupReport()
	case EmailIdentity:
		return NewSparseEmail()
	case EnforcerIdentity:
		return NewSparseEnforcer()
	case EnforcerLogIdentity:
		return NewSparseEnforcerLog()
	case EnforcerProfileIdentity:
		return NewSparseEnforcerProfile()
	case EnforcerProfileMappingPolicyIdentity:
		return NewSparseEnforcerProfileMappingPolicy()
	case EnforcerReportIdentity:
		return NewSparseEnforcerReport()
	case EnforcerTraceReportIdentity:
		return NewSparseEnforcerTraceReport()
	case EventLogIdentity:
		return NewSparseEventLog()
	case ExportIdentity:
		return NewSparseExport()
	case ExternalNetworkIdentity:
		return NewSparseExternalNetwork()
	case FileAccessPolicyIdentity:
		return NewSparseFileAccessPolicy()
	case FileAccessReportIdentity:
		return NewSparseFileAccessReport()
	case FilePathIdentity:
		return NewSparseFilePath()
	case FlowReportIdentity:
		return NewSparseFlowReport()
	case GraphEdgeIdentity:
		return NewSparseGraphEdge()
	case GraphNodeIdentity:
		return NewSparseGraphNode()
	case HitIdentity:
		return NewSparseHit()
	case HookPolicyIdentity:
		return NewSparseHookPolicy()
	case HostServiceIdentity:
		return NewSparseHostService()
	case HostServiceMappingPolicyIdentity:
		return NewSparseHostServiceMappingPolicy()
	case HTTPResourceSpecIdentity:
		return NewSparseHTTPResourceSpec()
	case ImportIdentity:
		return NewSparseImport()
	case ImportReferenceIdentity:
		return NewSparseImportReference()
	case ImportRequestIdentity:
		return NewSparseImportRequest()
	case InfrastructurePolicyIdentity:
		return NewSparseInfrastructurePolicy()
	case InstalledAppIdentity:
		return NewSparseInstalledApp()
	case InvoiceIdentity:
		return NewSparseInvoice()
	case InvoiceRecordIdentity:
		return NewSparseInvoiceRecord()
	case IPInfoIdentity:
		return NewSparseIPInfo()
	case IsolationProfileIdentity:
		return NewSparseIsolationProfile()
	case IssueIdentity:
		return NewSparseIssue()
	case IssueServiceTokenIdentity:
		return NewSparseIssueServiceToken()
	case LDAPProviderIdentity:
		return NewSparseLDAPProvider()
	case LogIdentity:
		return NewSparseLog()
	case MessageIdentity:
		return NewSparseMessage()
	case NamespaceIdentity:
		return NewSparseNamespace()
	case NamespaceMappingPolicyIdentity:
		return NewSparseNamespaceMappingPolicy()
	case NetworkAccessPolicyIdentity:
		return NewSparseNetworkAccessPolicy()
	case OAUTHInfoIdentity:
		return NewSparseOAUTHInfo()
	case OAUTHKeyIdentity:
		return NewSparseOAUTHKey()
	case OIDCProviderIdentity:
		return NewSparseOIDCProvider()
	case PacketReportIdentity:
		return NewSparsePacketReport()
	case PasswordResetIdentity:
		return NewSparsePasswordReset()
	case PlanIdentity:
		return NewSparsePlan()
	case PokeIdentity:
		return NewSparsePoke()
	case PolicyIdentity:
		return NewSparsePolicy()
	case PolicyGraphIdentity:
		return NewSparsePolicyGraph()
	case PolicyRefreshIdentity:
		return NewSparsePolicyRefresh()
	case PolicyRendererIdentity:
		return NewSparsePolicyRenderer()
	case PolicyRuleIdentity:
		return NewSparsePolicyRule()
	case PolicyTTLIdentity:
		return NewSparsePolicyTTL()
	case ProcessingUnitIdentity:
		return NewSparseProcessingUnit()
	case ProcessingUnitPolicyIdentity:
		return NewSparseProcessingUnitPolicy()
	case ProcessingUnitRefreshIdentity:
		return NewSparseProcessingUnitRefresh()
	case QuotaCheckIdentity:
		return NewSparseQuotaCheck()
	case QuotaPolicyIdentity:
		return NewSparseQuotaPolicy()
	case RecipeIdentity:
		return NewSparseRecipe()
	case RemoteProcessorIdentity:
		return NewSparseRemoteProcessor()
	case RenderedPolicyIdentity:
		return NewSparseRenderedPolicy()
	case RenderTemplateIdentity:
		return NewSparseRenderTemplate()
	case ReportIdentity:
		return NewSparseReport()
	case RevocationIdentity:
		return NewSparseRevocation()
	case RoleIdentity:
		return NewSparseRole()
	case SAMLProviderIdentity:
		return NewSparseSAMLProvider()
	case SandboxIdentity:
		return NewSparseSandbox()
	case SearchIdentity:
		return NewSparseSearch()
	case ServiceIdentity:
		return NewSparseService()
	case ServiceDependencyIdentity:
		return NewSparseServiceDependency()
	case ServiceTokenIdentity:
		return NewSparseServiceToken()
	case SquallTagIdentity:
		return NewSparseSquallTag()
	case SSHAuthorityIdentity:
		return NewSparseSSHAuthority()
	case SSHAuthorizationPolicyIdentity:
		return NewSparseSSHAuthorizationPolicy()
	case SSHCertificateIdentity:
		return NewSparseSSHCertificate()
	case SSHIdentityIdentity:
		return NewSparseSSHIdentity()
	case StatsInfoIdentity:
		return NewSparseStatsInfo()
	case StatsQueryIdentity:
		return NewSparseStatsQuery()
	case SuggestedPolicyIdentity:
		return NewSparseSuggestedPolicy()
	case TabulationIdentity:
		return NewSparseTabulation()
	case TagIdentity:
		return NewSparseTag()
	case TagInjectIdentity:
		return NewSparseTagInject()
	case TagValueIdentity:
		return NewSparseTagValue()
	case TextIndexIdentity:
		return NewSparseTextIndex()
	case TokenIdentity:
		return NewSparseToken()
	case TokenScopePolicyIdentity:
		return NewSparseTokenScopePolicy()
	case TriggerIdentity:
		return NewSparseTrigger()
	case TrustedCAIdentity:
		return NewSparseTrustedCA()
	case UserAccessPolicyIdentity:
		return NewSparseUserAccessPolicy()
	case ValidateUIParameterIdentity:
		return NewSparseValidateUIParameter()
	case VulnerabilityIdentity:
		return NewSparseVulnerability()
	case X509CertificateIdentity:
		return NewSparseX509Certificate()
	case X509CertificateCheckIdentity:
		return NewSparseX509CertificateCheck()
	default:
		return nil
	}
}

func (f modelManager) Indexes(identity elemental.Identity) [][]string {

	return indexesMap[identity.Name]
}

func (f modelManager) IdentifiableFromString(any string) elemental.Identifiable {

	return f.Identifiable(f.IdentityFromAny(any))
}

func (f modelManager) Identifiables(identity elemental.Identity) elemental.Identifiables {

	switch identity {

	case AccessReportIdentity:
		return &AccessReportsList{}
	case AccountIdentity:
		return &AccountsList{}
	case AccountCheckIdentity:
		return &AccountChecksList{}
	case ActivateIdentity:
		return &ActivatesList{}
	case ActivityIdentity:
		return &ActivitiesList{}
	case AlarmIdentity:
		return &AlarmsList{}
	case APIAuthorizationPolicyIdentity:
		return &APIAuthorizationPoliciesList{}
	case APICheckIdentity:
		return &APIChecksList{}
	case AppIdentity:
		return &AppsList{}
	case AppCredentialIdentity:
		return &AppCredentialsList{}
	case AuditProfileIdentity:
		return &AuditProfilesList{}
	case AuditProfileMappingPolicyIdentity:
		return &AuditProfileMappingPoliciesList{}
	case AuditReportIdentity:
		return &AuditReportsList{}
	case AuthnIdentity:
		return &AuthnsList{}
	case AuthorityIdentity:
		return &AuthoritiesList{}
	case AuthzIdentity:
		return &AuthzsList{}
	case AutomationIdentity:
		return &AutomationsList{}
	case AutomationTemplateIdentity:
		return &AutomationTemplatesList{}
	case AWSAPIGatewayIdentity:
		return &AWSAPIGatewaysList{}
	case AWSRegisterIdentity:
		return &AWSRegistersList{}
	case CategoryIdentity:
		return &CategoriesList{}
	case ClaimsIdentity:
		return &ClaimsList{}
	case ClauseMatchIdentity:
		return &ClauseMatchesList{}
	case CounterReportIdentity:
		return &CounterReportsList{}
	case CustomerIdentity:
		return &CustomersList{}
	case DataPathCertificateIdentity:
		return &DataPathCertificatesList{}
	case DependencyMapIdentity:
		return &DependencyMapsList{}
	case DNSLookupReportIdentity:
		return &DNSLookupReportsList{}
	case EmailIdentity:
		return &EmailsList{}
	case EnforcerIdentity:
		return &EnforcersList{}
	case EnforcerLogIdentity:
		return &EnforcerLogsList{}
	case EnforcerProfileIdentity:
		return &EnforcerProfilesList{}
	case EnforcerProfileMappingPolicyIdentity:
		return &EnforcerProfileMappingPoliciesList{}
	case EnforcerReportIdentity:
		return &EnforcerReportsList{}
	case EnforcerTraceReportIdentity:
		return &EnforcerTraceReportsList{}
	case EventLogIdentity:
		return &EventLogsList{}
	case ExportIdentity:
		return &ExportsList{}
	case ExternalNetworkIdentity:
		return &ExternalNetworksList{}
	case FileAccessPolicyIdentity:
		return &FileAccessPoliciesList{}
	case FileAccessReportIdentity:
		return &FileAccessReportsList{}
	case FilePathIdentity:
		return &FilePathsList{}
	case FlowReportIdentity:
		return &FlowReportsList{}
	case GraphEdgeIdentity:
		return &GraphEdgesList{}
	case GraphNodeIdentity:
		return &GraphNodesList{}
	case HitIdentity:
		return &HitsList{}
	case HookPolicyIdentity:
		return &HookPoliciesList{}
	case HostServiceIdentity:
		return &HostServicesList{}
	case HostServiceMappingPolicyIdentity:
		return &HostServiceMappingPoliciesList{}
	case HTTPResourceSpecIdentity:
		return &HTTPResourceSpecsList{}
	case ImportIdentity:
		return &ImportsList{}
	case ImportReferenceIdentity:
		return &ImportReferencesList{}
	case ImportRequestIdentity:
		return &ImportRequestsList{}
	case InfrastructurePolicyIdentity:
		return &InfrastructurePoliciesList{}
	case InstalledAppIdentity:
		return &InstalledAppsList{}
	case InvoiceIdentity:
		return &InvoicesList{}
	case InvoiceRecordIdentity:
		return &InvoiceRecordsList{}
	case IPInfoIdentity:
		return &IPInfosList{}
	case IsolationProfileIdentity:
		return &IsolationProfilesList{}
	case IssueIdentity:
		return &IssuesList{}
	case IssueServiceTokenIdentity:
		return &IssueServiceTokensList{}
	case LDAPProviderIdentity:
		return &LDAPProvidersList{}
	case LogIdentity:
		return &LogsList{}
	case MessageIdentity:
		return &MessagesList{}
	case NamespaceIdentity:
		return &NamespacesList{}
	case NamespaceMappingPolicyIdentity:
		return &NamespaceMappingPoliciesList{}
	case NetworkAccessPolicyIdentity:
		return &NetworkAccessPoliciesList{}
	case OAUTHInfoIdentity:
		return &OAUTHInfosList{}
	case OAUTHKeyIdentity:
		return &OAUTHKeysList{}
	case OIDCProviderIdentity:
		return &OIDCProvidersList{}
	case PacketReportIdentity:
		return &PacketReportsList{}
	case PasswordResetIdentity:
		return &PasswordResetsList{}
	case PlanIdentity:
		return &PlansList{}
	case PokeIdentity:
		return &PokesList{}
	case PolicyIdentity:
		return &PoliciesList{}
	case PolicyGraphIdentity:
		return &PolicyGraphsList{}
	case PolicyRefreshIdentity:
		return &PolicyRefreshsList{}
	case PolicyRendererIdentity:
		return &PolicyRenderersList{}
	case PolicyRuleIdentity:
		return &PolicyRulesList{}
	case PolicyTTLIdentity:
		return &PolicyTTLsList{}
	case ProcessingUnitIdentity:
		return &ProcessingUnitsList{}
	case ProcessingUnitPolicyIdentity:
		return &ProcessingUnitPoliciesList{}
	case ProcessingUnitRefreshIdentity:
		return &ProcessingUnitRefreshsList{}
	case QuotaCheckIdentity:
		return &QuotaChecksList{}
	case QuotaPolicyIdentity:
		return &QuotaPoliciesList{}
	case RecipeIdentity:
		return &RecipesList{}
	case RemoteProcessorIdentity:
		return &RemoteProcessorsList{}
	case RenderedPolicyIdentity:
		return &RenderedPoliciesList{}
	case RenderTemplateIdentity:
		return &RenderTemplatesList{}
	case ReportIdentity:
		return &ReportsList{}
	case RevocationIdentity:
		return &RevocationsList{}
	case RoleIdentity:
		return &RolesList{}
	case SAMLProviderIdentity:
		return &SAMLProvidersList{}
	case SandboxIdentity:
		return &SandboxsList{}
	case SearchIdentity:
		return &SearchesList{}
	case ServiceIdentity:
		return &ServicesList{}
	case ServiceDependencyIdentity:
		return &ServiceDependenciesList{}
	case ServiceTokenIdentity:
		return &ServiceTokensList{}
	case SquallTagIdentity:
		return &SquallTagsList{}
	case SSHAuthorityIdentity:
		return &SSHAuthoritiesList{}
	case SSHAuthorizationPolicyIdentity:
		return &SSHAuthorizationPoliciesList{}
	case SSHCertificateIdentity:
		return &SSHCertificatesList{}
	case SSHIdentityIdentity:
		return &SSHIdentitiesList{}
	case StatsInfoIdentity:
		return &StatsInfosList{}
	case StatsQueryIdentity:
		return &StatsQueriesList{}
	case SuggestedPolicyIdentity:
		return &SuggestedPoliciesList{}
	case TabulationIdentity:
		return &TabulationsList{}
	case TagIdentity:
		return &TagsList{}
	case TagInjectIdentity:
		return &TagInjectsList{}
	case TagValueIdentity:
		return &TagValuesList{}
	case TextIndexIdentity:
		return &TextIndexsList{}
	case TokenIdentity:
		return &TokensList{}
	case TokenScopePolicyIdentity:
		return &TokenScopePoliciesList{}
	case TriggerIdentity:
		return &TriggersList{}
	case TrustedCAIdentity:
		return &TrustedCAsList{}
	case UserAccessPolicyIdentity:
		return &UserAccessPoliciesList{}
	case ValidateUIParameterIdentity:
		return &ValidateUIParametersList{}
	case VulnerabilityIdentity:
		return &VulnerabilitiesList{}
	case X509CertificateIdentity:
		return &X509CertificatesList{}
	case X509CertificateCheckIdentity:
		return &X509CertificateChecksList{}
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiables(identity elemental.Identity) elemental.SparseIdentifiables {

	switch identity {

	case AccessReportIdentity:
		return &SparseAccessReportsList{}
	case AccountIdentity:
		return &SparseAccountsList{}
	case AccountCheckIdentity:
		return &SparseAccountChecksList{}
	case ActivateIdentity:
		return &SparseActivatesList{}
	case ActivityIdentity:
		return &SparseActivitiesList{}
	case AlarmIdentity:
		return &SparseAlarmsList{}
	case APIAuthorizationPolicyIdentity:
		return &SparseAPIAuthorizationPoliciesList{}
	case APICheckIdentity:
		return &SparseAPIChecksList{}
	case AppIdentity:
		return &SparseAppsList{}
	case AppCredentialIdentity:
		return &SparseAppCredentialsList{}
	case AuditProfileIdentity:
		return &SparseAuditProfilesList{}
	case AuditProfileMappingPolicyIdentity:
		return &SparseAuditProfileMappingPoliciesList{}
	case AuditReportIdentity:
		return &SparseAuditReportsList{}
	case AuthnIdentity:
		return &SparseAuthnsList{}
	case AuthorityIdentity:
		return &SparseAuthoritiesList{}
	case AuthzIdentity:
		return &SparseAuthzsList{}
	case AutomationIdentity:
		return &SparseAutomationsList{}
	case AutomationTemplateIdentity:
		return &SparseAutomationTemplatesList{}
	case AWSAPIGatewayIdentity:
		return &SparseAWSAPIGatewaysList{}
	case AWSRegisterIdentity:
		return &SparseAWSRegistersList{}
	case CategoryIdentity:
		return &SparseCategoriesList{}
	case ClaimsIdentity:
		return &SparseClaimsList{}
	case ClauseMatchIdentity:
		return &SparseClauseMatchesList{}
	case CounterReportIdentity:
		return &SparseCounterReportsList{}
	case CustomerIdentity:
		return &SparseCustomersList{}
	case DataPathCertificateIdentity:
		return &SparseDataPathCertificatesList{}
	case DependencyMapIdentity:
		return &SparseDependencyMapsList{}
	case DNSLookupReportIdentity:
		return &SparseDNSLookupReportsList{}
	case EmailIdentity:
		return &SparseEmailsList{}
	case EnforcerIdentity:
		return &SparseEnforcersList{}
	case EnforcerLogIdentity:
		return &SparseEnforcerLogsList{}
	case EnforcerProfileIdentity:
		return &SparseEnforcerProfilesList{}
	case EnforcerProfileMappingPolicyIdentity:
		return &SparseEnforcerProfileMappingPoliciesList{}
	case EnforcerReportIdentity:
		return &SparseEnforcerReportsList{}
	case EnforcerTraceReportIdentity:
		return &SparseEnforcerTraceReportsList{}
	case EventLogIdentity:
		return &SparseEventLogsList{}
	case ExportIdentity:
		return &SparseExportsList{}
	case ExternalNetworkIdentity:
		return &SparseExternalNetworksList{}
	case FileAccessPolicyIdentity:
		return &SparseFileAccessPoliciesList{}
	case FileAccessReportIdentity:
		return &SparseFileAccessReportsList{}
	case FilePathIdentity:
		return &SparseFilePathsList{}
	case FlowReportIdentity:
		return &SparseFlowReportsList{}
	case GraphEdgeIdentity:
		return &SparseGraphEdgesList{}
	case GraphNodeIdentity:
		return &SparseGraphNodesList{}
	case HitIdentity:
		return &SparseHitsList{}
	case HookPolicyIdentity:
		return &SparseHookPoliciesList{}
	case HostServiceIdentity:
		return &SparseHostServicesList{}
	case HostServiceMappingPolicyIdentity:
		return &SparseHostServiceMappingPoliciesList{}
	case HTTPResourceSpecIdentity:
		return &SparseHTTPResourceSpecsList{}
	case ImportIdentity:
		return &SparseImportsList{}
	case ImportReferenceIdentity:
		return &SparseImportReferencesList{}
	case ImportRequestIdentity:
		return &SparseImportRequestsList{}
	case InfrastructurePolicyIdentity:
		return &SparseInfrastructurePoliciesList{}
	case InstalledAppIdentity:
		return &SparseInstalledAppsList{}
	case InvoiceIdentity:
		return &SparseInvoicesList{}
	case InvoiceRecordIdentity:
		return &SparseInvoiceRecordsList{}
	case IPInfoIdentity:
		return &SparseIPInfosList{}
	case IsolationProfileIdentity:
		return &SparseIsolationProfilesList{}
	case IssueIdentity:
		return &SparseIssuesList{}
	case IssueServiceTokenIdentity:
		return &SparseIssueServiceTokensList{}
	case LDAPProviderIdentity:
		return &SparseLDAPProvidersList{}
	case LogIdentity:
		return &SparseLogsList{}
	case MessageIdentity:
		return &SparseMessagesList{}
	case NamespaceIdentity:
		return &SparseNamespacesList{}
	case NamespaceMappingPolicyIdentity:
		return &SparseNamespaceMappingPoliciesList{}
	case NetworkAccessPolicyIdentity:
		return &SparseNetworkAccessPoliciesList{}
	case OAUTHInfoIdentity:
		return &SparseOAUTHInfosList{}
	case OAUTHKeyIdentity:
		return &SparseOAUTHKeysList{}
	case OIDCProviderIdentity:
		return &SparseOIDCProvidersList{}
	case PacketReportIdentity:
		return &SparsePacketReportsList{}
	case PasswordResetIdentity:
		return &SparsePasswordResetsList{}
	case PlanIdentity:
		return &SparsePlansList{}
	case PokeIdentity:
		return &SparsePokesList{}
	case PolicyIdentity:
		return &SparsePoliciesList{}
	case PolicyGraphIdentity:
		return &SparsePolicyGraphsList{}
	case PolicyRefreshIdentity:
		return &SparsePolicyRefreshsList{}
	case PolicyRendererIdentity:
		return &SparsePolicyRenderersList{}
	case PolicyRuleIdentity:
		return &SparsePolicyRulesList{}
	case PolicyTTLIdentity:
		return &SparsePolicyTTLsList{}
	case ProcessingUnitIdentity:
		return &SparseProcessingUnitsList{}
	case ProcessingUnitPolicyIdentity:
		return &SparseProcessingUnitPoliciesList{}
	case ProcessingUnitRefreshIdentity:
		return &SparseProcessingUnitRefreshsList{}
	case QuotaCheckIdentity:
		return &SparseQuotaChecksList{}
	case QuotaPolicyIdentity:
		return &SparseQuotaPoliciesList{}
	case RecipeIdentity:
		return &SparseRecipesList{}
	case RemoteProcessorIdentity:
		return &SparseRemoteProcessorsList{}
	case RenderedPolicyIdentity:
		return &SparseRenderedPoliciesList{}
	case RenderTemplateIdentity:
		return &SparseRenderTemplatesList{}
	case ReportIdentity:
		return &SparseReportsList{}
	case RevocationIdentity:
		return &SparseRevocationsList{}
	case RoleIdentity:
		return &SparseRolesList{}
	case SAMLProviderIdentity:
		return &SparseSAMLProvidersList{}
	case SandboxIdentity:
		return &SparseSandboxsList{}
	case SearchIdentity:
		return &SparseSearchesList{}
	case ServiceIdentity:
		return &SparseServicesList{}
	case ServiceDependencyIdentity:
		return &SparseServiceDependenciesList{}
	case ServiceTokenIdentity:
		return &SparseServiceTokensList{}
	case SquallTagIdentity:
		return &SparseSquallTagsList{}
	case SSHAuthorityIdentity:
		return &SparseSSHAuthoritiesList{}
	case SSHAuthorizationPolicyIdentity:
		return &SparseSSHAuthorizationPoliciesList{}
	case SSHCertificateIdentity:
		return &SparseSSHCertificatesList{}
	case SSHIdentityIdentity:
		return &SparseSSHIdentitiesList{}
	case StatsInfoIdentity:
		return &SparseStatsInfosList{}
	case StatsQueryIdentity:
		return &SparseStatsQueriesList{}
	case SuggestedPolicyIdentity:
		return &SparseSuggestedPoliciesList{}
	case TabulationIdentity:
		return &SparseTabulationsList{}
	case TagIdentity:
		return &SparseTagsList{}
	case TagInjectIdentity:
		return &SparseTagInjectsList{}
	case TagValueIdentity:
		return &SparseTagValuesList{}
	case TextIndexIdentity:
		return &SparseTextIndexsList{}
	case TokenIdentity:
		return &SparseTokensList{}
	case TokenScopePolicyIdentity:
		return &SparseTokenScopePoliciesList{}
	case TriggerIdentity:
		return &SparseTriggersList{}
	case TrustedCAIdentity:
		return &SparseTrustedCAsList{}
	case UserAccessPolicyIdentity:
		return &SparseUserAccessPoliciesList{}
	case ValidateUIParameterIdentity:
		return &SparseValidateUIParametersList{}
	case VulnerabilityIdentity:
		return &SparseVulnerabilitiesList{}
	case X509CertificateIdentity:
		return &SparseX509CertificatesList{}
	case X509CertificateCheckIdentity:
		return &SparseX509CertificateChecksList{}
	default:
		return nil
	}
}

func (f modelManager) IdentifiablesFromString(any string) elemental.Identifiables {

	return f.Identifiables(f.IdentityFromAny(any))
}

func (f modelManager) Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

var manager = modelManager{}

// Manager returns the model elemental.ModelManager.
func Manager() elemental.ModelManager { return manager }

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		AccessReportIdentity,
		AccountIdentity,
		AccountCheckIdentity,
		ActivateIdentity,
		ActivityIdentity,
		AlarmIdentity,
		APIAuthorizationPolicyIdentity,
		APICheckIdentity,
		AppIdentity,
		AppCredentialIdentity,
		AuditProfileIdentity,
		AuditProfileMappingPolicyIdentity,
		AuditReportIdentity,
		AuthnIdentity,
		AuthorityIdentity,
		AuthzIdentity,
		AutomationIdentity,
		AutomationTemplateIdentity,
		AWSAPIGatewayIdentity,
		AWSRegisterIdentity,
		CategoryIdentity,
		ClaimsIdentity,
		ClauseMatchIdentity,
		CounterReportIdentity,
		CustomerIdentity,
		DataPathCertificateIdentity,
		DependencyMapIdentity,
		DNSLookupReportIdentity,
		EmailIdentity,
		EnforcerIdentity,
		EnforcerLogIdentity,
		EnforcerProfileIdentity,
		EnforcerProfileMappingPolicyIdentity,
		EnforcerReportIdentity,
		EnforcerTraceReportIdentity,
		EventLogIdentity,
		ExportIdentity,
		ExternalNetworkIdentity,
		FileAccessPolicyIdentity,
		FileAccessReportIdentity,
		FilePathIdentity,
		FlowReportIdentity,
		GraphEdgeIdentity,
		GraphNodeIdentity,
		HitIdentity,
		HookPolicyIdentity,
		HostServiceIdentity,
		HostServiceMappingPolicyIdentity,
		HTTPResourceSpecIdentity,
		ImportIdentity,
		ImportReferenceIdentity,
		ImportRequestIdentity,
		InfrastructurePolicyIdentity,
		InstalledAppIdentity,
		InvoiceIdentity,
		InvoiceRecordIdentity,
		IPInfoIdentity,
		IsolationProfileIdentity,
		IssueIdentity,
		IssueServiceTokenIdentity,
		LDAPProviderIdentity,
		LogIdentity,
		MessageIdentity,
		NamespaceIdentity,
		NamespaceMappingPolicyIdentity,
		NetworkAccessPolicyIdentity,
		OAUTHInfoIdentity,
		OAUTHKeyIdentity,
		OIDCProviderIdentity,
		PacketReportIdentity,
		PasswordResetIdentity,
		PlanIdentity,
		PokeIdentity,
		PolicyIdentity,
		PolicyGraphIdentity,
		PolicyRefreshIdentity,
		PolicyRendererIdentity,
		PolicyRuleIdentity,
		PolicyTTLIdentity,
		ProcessingUnitIdentity,
		ProcessingUnitPolicyIdentity,
		ProcessingUnitRefreshIdentity,
		QuotaCheckIdentity,
		QuotaPolicyIdentity,
		RecipeIdentity,
		RemoteProcessorIdentity,
		RenderedPolicyIdentity,
		RenderTemplateIdentity,
		ReportIdentity,
		RevocationIdentity,
		RoleIdentity,
		RootIdentity,
		SAMLProviderIdentity,
		SandboxIdentity,
		SearchIdentity,
		ServiceIdentity,
		ServiceDependencyIdentity,
		ServiceTokenIdentity,
		SquallTagIdentity,
		SSHAuthorityIdentity,
		SSHAuthorizationPolicyIdentity,
		SSHCertificateIdentity,
		SSHIdentityIdentity,
		StatsInfoIdentity,
		StatsQueryIdentity,
		SuggestedPolicyIdentity,
		TabulationIdentity,
		TagIdentity,
		TagInjectIdentity,
		TagValueIdentity,
		TextIndexIdentity,
		TokenIdentity,
		TokenScopePolicyIdentity,
		TriggerIdentity,
		TrustedCAIdentity,
		UserAccessPolicyIdentity,
		ValidateUIParameterIdentity,
		VulnerabilityIdentity,
		X509CertificateIdentity,
		X509CertificateCheckIdentity,
	}
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case AccessReportIdentity:
		return []string{}
	case AccountIdentity:
		return []string{}
	case AccountCheckIdentity:
		return []string{}
	case ActivateIdentity:
		return []string{}
	case ActivityIdentity:
		return []string{}
	case AlarmIdentity:
		return []string{}
	case APIAuthorizationPolicyIdentity:
		return []string{
			"apiauth",
			"apiauths",
		}
	case APICheckIdentity:
		return []string{}
	case AppIdentity:
		return []string{}
	case AppCredentialIdentity:
		return []string{
			"appcred",
			"appcreds",
		}
	case AuditProfileIdentity:
		return []string{
			"ap",
		}
	case AuditProfileMappingPolicyIdentity:
		return []string{
			"audpol",
			"audpols",
		}
	case AuditReportIdentity:
		return []string{}
	case AuthnIdentity:
		return []string{}
	case AuthorityIdentity:
		return []string{
			"ca",
		}
	case AuthzIdentity:
		return []string{}
	case AutomationIdentity:
		return []string{
			"autos",
			"auto",
		}
	case AutomationTemplateIdentity:
		return []string{
			"autotmpl",
		}
	case AWSAPIGatewayIdentity:
		return []string{}
	case AWSRegisterIdentity:
		return []string{}
	case CategoryIdentity:
		return []string{}
	case ClaimsIdentity:
		return []string{}
	case ClauseMatchIdentity:
		return []string{}
	case CounterReportIdentity:
		return []string{}
	case CustomerIdentity:
		return []string{}
	case DataPathCertificateIdentity:
		return []string{}
	case DependencyMapIdentity:
		return []string{
			"depmaps",
			"depmap",
		}
	case DNSLookupReportIdentity:
		return []string{}
	case EmailIdentity:
		return []string{}
	case EnforcerIdentity:
		return []string{}
	case EnforcerLogIdentity:
		return []string{}
	case EnforcerProfileIdentity:
		return []string{
			"profile",
			"profiles",
		}
	case EnforcerProfileMappingPolicyIdentity:
		return []string{
			"enfpols",
			"enfpol",
			"epm",
		}
	case EnforcerReportIdentity:
		return []string{}
	case EnforcerTraceReportIdentity:
		return []string{}
	case EventLogIdentity:
		return []string{}
	case ExportIdentity:
		return []string{}
	case ExternalNetworkIdentity:
		return []string{
			"extnet",
			"extnets",
		}
	case FileAccessPolicyIdentity:
		return []string{}
	case FileAccessReportIdentity:
		return []string{}
	case FilePathIdentity:
		return []string{
			"fp",
			"fps",
		}
	case FlowReportIdentity:
		return []string{}
	case GraphEdgeIdentity:
		return []string{}
	case GraphNodeIdentity:
		return []string{}
	case HitIdentity:
		return []string{}
	case HookPolicyIdentity:
		return []string{
			"hook",
			"hooks",
			"hookpol",
			"hookpols",
		}
	case HostServiceIdentity:
		return []string{
			"hostsrv",
			"hostsrvs",
		}
	case HostServiceMappingPolicyIdentity:
		return []string{
			"hostsrvmappol",
			"hostsrvmappols",
		}
	case HTTPResourceSpecIdentity:
		return []string{
			"httpresource",
			"resource",
			"httpspec",
		}
	case ImportIdentity:
		return []string{}
	case ImportReferenceIdentity:
		return []string{
			"importref",
			"impref",
		}
	case ImportRequestIdentity:
		return []string{
			"req",
			"reqs",
			"ireq",
			"ireqs",
		}
	case InfrastructurePolicyIdentity:
		return []string{
			"infrapol",
			"infrapols",
		}
	case InstalledAppIdentity:
		return []string{
			"iapps",
			"iapp",
		}
	case InvoiceIdentity:
		return []string{}
	case InvoiceRecordIdentity:
		return []string{}
	case IPInfoIdentity:
		return []string{}
	case IsolationProfileIdentity:
		return []string{
			"ip",
		}
	case IssueIdentity:
		return []string{}
	case IssueServiceTokenIdentity:
		return []string{}
	case LDAPProviderIdentity:
		return []string{}
	case LogIdentity:
		return []string{}
	case MessageIdentity:
		return []string{
			"mess",
		}
	case NamespaceIdentity:
		return []string{
			"ns",
		}
	case NamespaceMappingPolicyIdentity:
		return []string{
			"nspolicy",
			"nspolicies",
			"nsmap",
			"nsmaps",
		}
	case NetworkAccessPolicyIdentity:
		return []string{
			"netpol",
			"netpols",
		}
	case OAUTHInfoIdentity:
		return []string{}
	case OAUTHKeyIdentity:
		return []string{}
	case OIDCProviderIdentity:
		return []string{}
	case PacketReportIdentity:
		return []string{}
	case PasswordResetIdentity:
		return []string{}
	case PlanIdentity:
		return []string{}
	case PokeIdentity:
		return []string{}
	case PolicyIdentity:
		return []string{}
	case PolicyGraphIdentity:
		return []string{
			"polgraph",
		}
	case PolicyRefreshIdentity:
		return []string{}
	case PolicyRendererIdentity:
		return []string{}
	case PolicyRuleIdentity:
		return []string{}
	case PolicyTTLIdentity:
		return []string{}
	case ProcessingUnitIdentity:
		return []string{
			"pu",
			"pus",
		}
	case ProcessingUnitPolicyIdentity:
		return []string{
			"pup",
			"pups",
		}
	case ProcessingUnitRefreshIdentity:
		return []string{}
	case QuotaCheckIdentity:
		return []string{}
	case QuotaPolicyIdentity:
		return []string{
			"quota",
			"quotas",
			"quotapol",
			"quotapols",
		}
	case RecipeIdentity:
		return []string{
			"rcp",
		}
	case RemoteProcessorIdentity:
		return []string{
			"hks",
			"hk",
		}
	case RenderedPolicyIdentity:
		return []string{
			"rpol",
			"rpols",
		}
	case RenderTemplateIdentity:
		return []string{
			"cook",
			"rtpl",
		}
	case ReportIdentity:
		return []string{}
	case RevocationIdentity:
		return []string{}
	case RoleIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case SAMLProviderIdentity:
		return []string{}
	case SandboxIdentity:
		return []string{}
	case SearchIdentity:
		return []string{}
	case ServiceIdentity:
		return []string{
			"srv",
		}
	case ServiceDependencyIdentity:
		return []string{
			"srvdep",
			"srvdeps",
		}
	case ServiceTokenIdentity:
		return []string{}
	case SquallTagIdentity:
		return []string{}
	case SSHAuthorityIdentity:
		return []string{}
	case SSHAuthorizationPolicyIdentity:
		return []string{
			"sshpol",
			"sshpols",
		}
	case SSHCertificateIdentity:
		return []string{}
	case SSHIdentityIdentity:
		return []string{}
	case StatsInfoIdentity:
		return []string{
			"si",
		}
	case StatsQueryIdentity:
		return []string{
			"sq",
		}
	case SuggestedPolicyIdentity:
		return []string{
			"sugpol",
			"sugpols",
			"sugg",
			"suggs",
		}
	case TabulationIdentity:
		return []string{
			"table",
			"tables",
			"tabs",
			"tab",
		}
	case TagIdentity:
		return []string{}
	case TagInjectIdentity:
		return []string{}
	case TagValueIdentity:
		return []string{}
	case TextIndexIdentity:
		return []string{}
	case TokenIdentity:
		return []string{}
	case TokenScopePolicyIdentity:
		return []string{
			"tsp",
		}
	case TriggerIdentity:
		return []string{}
	case TrustedCAIdentity:
		return []string{}
	case UserAccessPolicyIdentity:
		return []string{
			"usrpol",
			"usrpols",
		}
	case ValidateUIParameterIdentity:
		return []string{
			"validparam",
		}
	case VulnerabilityIdentity:
		return []string{
			"vulns",
			"vul",
			"vuln",
			"vuls",
		}
	case X509CertificateIdentity:
		return []string{}
	case X509CertificateCheckIdentity:
		return []string{}
	}

	return nil
}
