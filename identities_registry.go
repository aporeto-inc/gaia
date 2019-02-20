package gaia

import "go.aporeto.io/elemental"

var (
	identityNamesMap = map[string]elemental.Identity{
		"account":                AccountIdentity,
		"accountcheck":           AccountCheckIdentity,
		"activate":               ActivateIdentity,
		"activity":               ActivityIdentity,
		"alarm":                  AlarmIdentity,
		"apiauthorizationpolicy": APIAuthorizationPolicyIdentity,
		"apicheck":               APICheckIdentity,
		"app":                    AppIdentity,
		"appcredential":          AppCredentialIdentity,

		"auditprofile":              AuditProfileIdentity,
		"auditprofilemappingpolicy": AuditProfileMappingPolicyIdentity,
		"auditreport":               AuditReportIdentity,
		"auth":                      AuthIdentity,
		"authority":                 AuthorityIdentity,
		"automation":                AutomationIdentity,
		"automationtemplate":        AutomationTemplateIdentity,

		"awsaccount":    AWSAccountIdentity,
		"awsapigateway": AWSAPIGatewayIdentity,
		"awsregister":   AWSRegisterIdentity,
		"category":      CategoryIdentity,

		"customer":      CustomerIdentity,
		"dependencymap": DependencyMapIdentity,
		"email":         EmailIdentity,

		"enforcer":                     EnforcerIdentity,
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

		"hookpolicy":               HookPolicyIdentity,
		"hostservice":              HostServiceIdentity,
		"hostservicemappingpolicy": HostServiceMappingPolicyIdentity,
		"httpresourcespec":         HTTPResourceSpecIdentity,
		"import":                   ImportIdentity,
		"installedapp":             InstalledAppIdentity,
		"invoice":                  InvoiceIdentity,
		"invoicerecord":            InvoiceRecordIdentity,
		"ipinfo":                   IPInfoIdentity,
		"isolationprofile":         IsolationProfileIdentity,
		"issue":                    IssueIdentity,
		"jaegerbatch":              JaegerbatchIdentity,
		"log":                      LogIdentity,
		"message":                  MessageIdentity,
		"namespace":                NamespaceIdentity,
		"namespacemappingpolicy":   NamespaceMappingPolicyIdentity,
		"networkaccesspolicy":      NetworkAccessPolicyIdentity,
		"oidcprovider":             OIDCProviderIdentity,
		"packetreport":             PacketReportIdentity,
		"passwordreset":            PasswordResetIdentity,
		"plan":                     PlanIdentity,
		"poke":                     PokeIdentity,
		"policy":                   PolicyIdentity,
		"policyrefresh":            PolicyRefreshIdentity,
		"policyrenderer":           PolicyRendererIdentity,
		"policyrule":               PolicyRuleIdentity,
		"privatekey":               PrivateKeyIdentity,
		"processingunit":           ProcessingUnitIdentity,
		"processingunitpolicy":     ProcessingUnitPolicyIdentity,

		"punode":                 PUNodeIdentity,
		"quotacheck":             QuotaCheckIdentity,
		"quotapolicy":            QuotaPolicyIdentity,
		"remoteprocessor":        RemoteProcessorIdentity,
		"renderedpolicy":         RenderedPolicyIdentity,
		"report":                 ReportIdentity,
		"revocation":             RevocationIdentity,
		"role":                   RoleIdentity,
		"root":                   RootIdentity,
		"service":                ServiceIdentity,
		"servicedependency":      ServiceDependencyIdentity,
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

		"token":            TokenIdentity,
		"tokenscopepolicy": TokenScopePolicyIdentity,

		"trigger":              TriggerIdentity,
		"vulnerability":        VulnerabilityIdentity,
		"x509certificate":      X509CertificateIdentity,
		"x509certificatecheck": X509CertificateCheckIdentity,
	}

	identitycategoriesMap = map[string]elemental.Identity{
		"accounts":                 AccountIdentity,
		"accountchecks":            AccountCheckIdentity,
		"activate":                 ActivateIdentity,
		"activities":               ActivityIdentity,
		"alarms":                   AlarmIdentity,
		"apiauthorizationpolicies": APIAuthorizationPolicyIdentity,
		"apichecks":                APICheckIdentity,
		"apps":                     AppIdentity,
		"appcredentials":           AppCredentialIdentity,

		"auditprofiles":               AuditProfileIdentity,
		"auditprofilemappingpolicies": AuditProfileMappingPolicyIdentity,
		"auditreports":                AuditReportIdentity,
		"auth":                        AuthIdentity,
		"authorities":                 AuthorityIdentity,
		"automations":                 AutomationIdentity,
		"automationtemplates":         AutomationTemplateIdentity,

		"awsaccounts":    AWSAccountIdentity,
		"awsapigateways": AWSAPIGatewayIdentity,
		"awsregister":    AWSRegisterIdentity,
		"categories":     CategoryIdentity,

		"customers":      CustomerIdentity,
		"dependencymaps": DependencyMapIdentity,
		"emails":         EmailIdentity,

		"enforcers":                      EnforcerIdentity,
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

		"hookpolicies":               HookPolicyIdentity,
		"hostservices":               HostServiceIdentity,
		"hostservicemappingpolicies": HostServiceMappingPolicyIdentity,
		"httpresourcespecs":          HTTPResourceSpecIdentity,
		"import":                     ImportIdentity,
		"installedapps":              InstalledAppIdentity,
		"invoices":                   InvoiceIdentity,
		"invoicerecords":             InvoiceRecordIdentity,
		"ipinfos":                    IPInfoIdentity,
		"isolationprofiles":          IsolationProfileIdentity,
		"issue":                      IssueIdentity,
		"jaegerbatchs":               JaegerbatchIdentity,
		"logs":                       LogIdentity,
		"messages":                   MessageIdentity,
		"namespaces":                 NamespaceIdentity,
		"namespacemappingpolicies":   NamespaceMappingPolicyIdentity,
		"networkaccesspolicies":      NetworkAccessPolicyIdentity,
		"oidcproviders":              OIDCProviderIdentity,
		"packetreports":              PacketReportIdentity,
		"passwordreset":              PasswordResetIdentity,
		"plans":                      PlanIdentity,
		"poke":                       PokeIdentity,
		"policies":                   PolicyIdentity,
		"policyrefreshs":             PolicyRefreshIdentity,
		"policyrenderers":            PolicyRendererIdentity,
		"policyrules":                PolicyRuleIdentity,
		"privatekeys":                PrivateKeyIdentity,
		"processingunits":            ProcessingUnitIdentity,
		"processingunitpolicies":     ProcessingUnitPolicyIdentity,

		"punodes":                  PUNodeIdentity,
		"quotacheck":               QuotaCheckIdentity,
		"quotapolicies":            QuotaPolicyIdentity,
		"remoteprocessors":         RemoteProcessorIdentity,
		"renderedpolicies":         RenderedPolicyIdentity,
		"reports":                  ReportIdentity,
		"revocations":              RevocationIdentity,
		"roles":                    RoleIdentity,
		"root":                     RootIdentity,
		"services":                 ServiceIdentity,
		"servicedependencies":      ServiceDependencyIdentity,
		"squalltags":               SquallTagIdentity,
		"sshauthorities":           SSHAuthorityIdentity,
		"sshauthorizationpolicies": SSHAuthorizationPolicyIdentity,
		"sshcertificates":          SSHCertificateIdentity,
		"sshidentitiess":           SSHIdentityIdentity,
		"statsinfo":                StatsInfoIdentity,
		"statsqueries":             StatsQueryIdentity,
		"suggestedpolicies":        SuggestedPolicyIdentity,
		"tabulations":              TabulationIdentity,
		"tags":                     TagIdentity,
		"taginjects":               TagInjectIdentity,
		"tagvalues":                TagValueIdentity,

		"tokens":             TokenIdentity,
		"tokenscopepolicies": TokenScopePolicyIdentity,

		"triggers":              TriggerIdentity,
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
		"aws":            AWSAccountIdentity,
		"awsaccs":        AWSAccountIdentity,
		"awsacc":         AWSAccountIdentity,
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
		"iapps":          InstalledAppIdentity,
		"iapp":           InstalledAppIdentity,
		"ip":             IsolationProfileIdentity,
		"sp":             JaegerbatchIdentity,
		"mess":           MessageIdentity,
		"ns":             NamespaceIdentity,
		"nspolicy":       NamespaceMappingPolicyIdentity,
		"nspolicies":     NamespaceMappingPolicyIdentity,
		"nsmap":          NamespaceMappingPolicyIdentity,
		"nsmaps":         NamespaceMappingPolicyIdentity,
		"netpol":         NetworkAccessPolicyIdentity,
		"netpols":        NetworkAccessPolicyIdentity,
		"pu":             ProcessingUnitIdentity,
		"pus":            ProcessingUnitIdentity,
		"pup":            ProcessingUnitPolicyIdentity,
		"quota":          QuotaPolicyIdentity,
		"quotas":         QuotaPolicyIdentity,
		"quotapol":       QuotaPolicyIdentity,
		"quotapols":      QuotaPolicyIdentity,
		"hks":            RemoteProcessorIdentity,
		"hk":             RemoteProcessorIdentity,
		"rpol":           RenderedPolicyIdentity,
		"rpols":          RenderedPolicyIdentity,
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
		"vulns":          VulnerabilityIdentity,
		"vul":            VulnerabilityIdentity,
		"vuln":           VulnerabilityIdentity,
		"vuls":           VulnerabilityIdentity,
	}

	indexesMap = map[string][][]string{
		"account": [][]string{
			[]string{":unique", "name"},
			[]string{":unique", "email"},
			[]string{"activationToken"},
			[]string{"resetPasswordToken"},
		},
		"accountcheck": nil,
		"activate":     nil,
		"activity": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "date"},
			[]string{"namespace", "operation"},
		},
		"alarm": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "kind"},
		},
		"apiauthorizationpolicy": nil,
		"apicheck":               nil,
		"app":                    nil,
		"appcredential": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "normalizedTags"},
		},
		"auditprofile": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
		},
		"auditprofilemappingpolicy": nil,
		"auditreport":               nil,
		"auth":                      nil,
		"authority": [][]string{
			[]string{":shard", "$hashed:serialNumber"},
			[]string{":unique", "serialNumber"},
			[]string{"commonName"},
		},
		"automation": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
		},
		"automationtemplate": nil,
		"awsaccount":         nil,
		"awsapigateway":      nil,
		"awsregister":        nil,
		"category":           nil,
		"customer": [][]string{
			[]string{"providerCustomerID"},
		},
		"dependencymap": nil,
		"email":         nil,
		"enforcer": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "normalizedTags"},
		},
		"enforcerprofile": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "normalizedTags"},
		},
		"enforcerprofilemappingpolicy": nil,
		"enforcerreport":               nil,
		"enforcertracereport":          nil,
		"eventlog":                     nil,
		"export":                       nil,
		"externalnetwork": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "archived"},
		},
		"fileaccesspolicy": nil,
		"fileaccessreport": nil,
		"filepath": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "normalizedTags"},
			[]string{"archived"},
		},
		"flowreport": nil,
		"hookpolicy": nil,
		"hostservice": [][]string{
			[]string{":shard", "zone", "zhash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "archived"},
		},
		"hostservicemappingpolicy": nil,
		"httpresourcespec": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "archived"},
			[]string{"namespace", "normalizedTags"},
		},
		"import": nil,
		"installedapp": [][]string{
			[]string{"accountName", "name"},
		},
		"invoice":       nil,
		"invoicerecord": nil,
		"ipinfo":        nil,
		"isolationprofile": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "normalizedTags"},
		},
		"issue":       nil,
		"jaegerbatch": nil,
		"log":         nil,
		"message": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
		},
		"namespace": [][]string{
			[]string{":shard", "$hashed:name"},
			[]string{":unique", "name"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
		},
		"namespacemappingpolicy": nil,
		"networkaccesspolicy":    nil,
		"oidcprovider": [][]string{
			[]string{":unique", "parentID", "name"},
		},
		"packetreport":  nil,
		"passwordreset": nil,
		"plan":          nil,
		"poke":          nil,
		"policy": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace", "name"},
			[]string{"namespace"},
			[]string{"namespace", "type"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "type", "allObjectTags"},
			[]string{"namespace", "type", "allSubjectTags"},
			[]string{"namespace", "type", "allObjectTags", "disabled"},
			[]string{"namespace", "type", "allSubjectTags", "disabled"},
			[]string{"namespace", "type", "allObjectTags", "propagate"},
			[]string{"namespace", "type", "allSubjectTags", "propagate"},
			[]string{"namespace", "fallback"},
		},
		"policyrefresh":  nil,
		"policyrenderer": nil,
		"policyrule":     nil,
		"privatekey":     nil,
		"processingunit": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "archived"},
			[]string{"namespace", "operationalStatus", "archived"},
			[]string{"namespace", "normalizedTags", "archived"},
		},
		"processingunitpolicy": nil,
		"punode":               nil,
		"quotacheck":           nil,
		"quotapolicy":          nil,
		"remoteprocessor":      nil,
		"renderedpolicy":       nil,
		"report":               nil,
		"revocation": [][]string{
			[]string{":shard", "$hashed:serialNumber"},
			[]string{":unique", "serialNumber"},
		},
		"role": nil,
		"root": nil,
		"service": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
			[]string{"namespace", "archived"},
			[]string{"namespace", "normalizedTags"},
			[]string{"allAPITags"},
			[]string{"namespace", "allAPITags"},
			[]string{"allServiceTags"},
			[]string{"namespace", "allServiceTags"},
		},
		"servicedependency":      nil,
		"squalltag":              nil,
		"sshauthority":           nil,
		"sshauthorizationpolicy": nil,
		"sshcertificate":         nil,
		"sshidentity":            nil,
		"statsinfo":              nil,
		"statsquery":             nil,
		"suggestedpolicy":        nil,
		"tabulation":             nil,
		"tag":                    nil,
		"taginject":              nil,
		"tagvalue":               nil,
		"token":                  nil,
		"tokenscopepolicy":       nil,
		"trigger":                nil,
		"vulnerability": [][]string{
			[]string{":shard", "zone", "zHash"},
			[]string{"namespace"},
			[]string{"namespace", "name"},
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
	case AuthIdentity:
		return NewAuth()
	case AuthorityIdentity:
		return NewAuthority()
	case AutomationIdentity:
		return NewAutomation()
	case AutomationTemplateIdentity:
		return NewAutomationTemplate()
	case AWSAccountIdentity:
		return NewAWSAccount()
	case AWSAPIGatewayIdentity:
		return NewAWSAPIGateway()
	case AWSRegisterIdentity:
		return NewAWSRegister()
	case CategoryIdentity:
		return NewCategory()
	case CustomerIdentity:
		return NewCustomer()
	case DependencyMapIdentity:
		return NewDependencyMap()
	case EmailIdentity:
		return NewEmail()
	case EnforcerIdentity:
		return NewEnforcer()
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
	case JaegerbatchIdentity:
		return NewJaegerbatch()
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
	case PolicyRefreshIdentity:
		return NewPolicyRefresh()
	case PolicyRendererIdentity:
		return NewPolicyRenderer()
	case PolicyRuleIdentity:
		return NewPolicyRule()
	case PrivateKeyIdentity:
		return NewPrivateKey()
	case ProcessingUnitIdentity:
		return NewProcessingUnit()
	case ProcessingUnitPolicyIdentity:
		return NewProcessingUnitPolicy()
	case PUNodeIdentity:
		return NewPUNode()
	case QuotaCheckIdentity:
		return NewQuotaCheck()
	case QuotaPolicyIdentity:
		return NewQuotaPolicy()
	case RemoteProcessorIdentity:
		return NewRemoteProcessor()
	case RenderedPolicyIdentity:
		return NewRenderedPolicy()
	case ReportIdentity:
		return NewReport()
	case RevocationIdentity:
		return NewRevocation()
	case RoleIdentity:
		return NewRole()
	case RootIdentity:
		return NewRoot()
	case ServiceIdentity:
		return NewService()
	case ServiceDependencyIdentity:
		return NewServiceDependency()
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
	case TokenIdentity:
		return NewToken()
	case TokenScopePolicyIdentity:
		return NewTokenScopePolicy()
	case TriggerIdentity:
		return NewTrigger()
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
	case AuthIdentity:
		return NewSparseAuth()
	case AuthorityIdentity:
		return NewSparseAuthority()
	case AutomationIdentity:
		return NewSparseAutomation()
	case AutomationTemplateIdentity:
		return NewSparseAutomationTemplate()
	case AWSAccountIdentity:
		return NewSparseAWSAccount()
	case AWSAPIGatewayIdentity:
		return NewSparseAWSAPIGateway()
	case AWSRegisterIdentity:
		return NewSparseAWSRegister()
	case CategoryIdentity:
		return NewSparseCategory()
	case CustomerIdentity:
		return NewSparseCustomer()
	case DependencyMapIdentity:
		return NewSparseDependencyMap()
	case EmailIdentity:
		return NewSparseEmail()
	case EnforcerIdentity:
		return NewSparseEnforcer()
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
	case JaegerbatchIdentity:
		return NewSparseJaegerbatch()
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
	case PolicyRefreshIdentity:
		return NewSparsePolicyRefresh()
	case PolicyRendererIdentity:
		return NewSparsePolicyRenderer()
	case PolicyRuleIdentity:
		return NewSparsePolicyRule()
	case PrivateKeyIdentity:
		return NewSparsePrivateKey()
	case ProcessingUnitIdentity:
		return NewSparseProcessingUnit()
	case ProcessingUnitPolicyIdentity:
		return NewSparseProcessingUnitPolicy()
	case PUNodeIdentity:
		return NewSparsePUNode()
	case QuotaCheckIdentity:
		return NewSparseQuotaCheck()
	case QuotaPolicyIdentity:
		return NewSparseQuotaPolicy()
	case RemoteProcessorIdentity:
		return NewSparseRemoteProcessor()
	case RenderedPolicyIdentity:
		return NewSparseRenderedPolicy()
	case ReportIdentity:
		return NewSparseReport()
	case RevocationIdentity:
		return NewSparseRevocation()
	case RoleIdentity:
		return NewSparseRole()
	case ServiceIdentity:
		return NewSparseService()
	case ServiceDependencyIdentity:
		return NewSparseServiceDependency()
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
	case TokenIdentity:
		return NewSparseToken()
	case TokenScopePolicyIdentity:
		return NewSparseTokenScopePolicy()
	case TriggerIdentity:
		return NewSparseTrigger()
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
	case AuthIdentity:
		return &AuthsList{}
	case AuthorityIdentity:
		return &AuthoritiesList{}
	case AutomationIdentity:
		return &AutomationsList{}
	case AutomationTemplateIdentity:
		return &AutomationTemplatesList{}
	case AWSAccountIdentity:
		return &AWSAccountsList{}
	case AWSAPIGatewayIdentity:
		return &AWSAPIGatewaysList{}
	case AWSRegisterIdentity:
		return &AWSRegistersList{}
	case CategoryIdentity:
		return &CategoriesList{}
	case CustomerIdentity:
		return &CustomersList{}
	case DependencyMapIdentity:
		return &DependencyMapsList{}
	case EmailIdentity:
		return &EmailsList{}
	case EnforcerIdentity:
		return &EnforcersList{}
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
	case JaegerbatchIdentity:
		return &JaegerbatchsList{}
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
	case PolicyRefreshIdentity:
		return &PolicyRefreshsList{}
	case PolicyRendererIdentity:
		return &PolicyRenderersList{}
	case PolicyRuleIdentity:
		return &PolicyRulesList{}
	case PrivateKeyIdentity:
		return &PrivateKeysList{}
	case ProcessingUnitIdentity:
		return &ProcessingUnitsList{}
	case ProcessingUnitPolicyIdentity:
		return &ProcessingUnitPoliciesList{}
	case PUNodeIdentity:
		return &PUNodesList{}
	case QuotaCheckIdentity:
		return &QuotaChecksList{}
	case QuotaPolicyIdentity:
		return &QuotaPoliciesList{}
	case RemoteProcessorIdentity:
		return &RemoteProcessorsList{}
	case RenderedPolicyIdentity:
		return &RenderedPoliciesList{}
	case ReportIdentity:
		return &ReportsList{}
	case RevocationIdentity:
		return &RevocationsList{}
	case RoleIdentity:
		return &RolesList{}
	case ServiceIdentity:
		return &ServicesList{}
	case ServiceDependencyIdentity:
		return &ServiceDependenciesList{}
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
	case TokenIdentity:
		return &TokensList{}
	case TokenScopePolicyIdentity:
		return &TokenScopePoliciesList{}
	case TriggerIdentity:
		return &TriggersList{}
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
	case AuthIdentity:
		return &SparseAuthsList{}
	case AuthorityIdentity:
		return &SparseAuthoritiesList{}
	case AutomationIdentity:
		return &SparseAutomationsList{}
	case AutomationTemplateIdentity:
		return &SparseAutomationTemplatesList{}
	case AWSAccountIdentity:
		return &SparseAWSAccountsList{}
	case AWSAPIGatewayIdentity:
		return &SparseAWSAPIGatewaysList{}
	case AWSRegisterIdentity:
		return &SparseAWSRegistersList{}
	case CategoryIdentity:
		return &SparseCategoriesList{}
	case CustomerIdentity:
		return &SparseCustomersList{}
	case DependencyMapIdentity:
		return &SparseDependencyMapsList{}
	case EmailIdentity:
		return &SparseEmailsList{}
	case EnforcerIdentity:
		return &SparseEnforcersList{}
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
	case JaegerbatchIdentity:
		return &SparseJaegerbatchsList{}
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
	case PolicyRefreshIdentity:
		return &SparsePolicyRefreshsList{}
	case PolicyRendererIdentity:
		return &SparsePolicyRenderersList{}
	case PolicyRuleIdentity:
		return &SparsePolicyRulesList{}
	case PrivateKeyIdentity:
		return &SparsePrivateKeysList{}
	case ProcessingUnitIdentity:
		return &SparseProcessingUnitsList{}
	case ProcessingUnitPolicyIdentity:
		return &SparseProcessingUnitPoliciesList{}
	case PUNodeIdentity:
		return &SparsePUNodesList{}
	case QuotaCheckIdentity:
		return &SparseQuotaChecksList{}
	case QuotaPolicyIdentity:
		return &SparseQuotaPoliciesList{}
	case RemoteProcessorIdentity:
		return &SparseRemoteProcessorsList{}
	case RenderedPolicyIdentity:
		return &SparseRenderedPoliciesList{}
	case ReportIdentity:
		return &SparseReportsList{}
	case RevocationIdentity:
		return &SparseRevocationsList{}
	case RoleIdentity:
		return &SparseRolesList{}
	case ServiceIdentity:
		return &SparseServicesList{}
	case ServiceDependencyIdentity:
		return &SparseServiceDependenciesList{}
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
	case TokenIdentity:
		return &SparseTokensList{}
	case TokenScopePolicyIdentity:
		return &SparseTokenScopePoliciesList{}
	case TriggerIdentity:
		return &SparseTriggersList{}
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
		AuthIdentity,
		AuthorityIdentity,
		AutomationIdentity,
		AutomationTemplateIdentity,
		AWSAccountIdentity,
		AWSAPIGatewayIdentity,
		AWSRegisterIdentity,
		CategoryIdentity,
		CustomerIdentity,
		DependencyMapIdentity,
		EmailIdentity,
		EnforcerIdentity,
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
		HookPolicyIdentity,
		HostServiceIdentity,
		HostServiceMappingPolicyIdentity,
		HTTPResourceSpecIdentity,
		ImportIdentity,
		InstalledAppIdentity,
		InvoiceIdentity,
		InvoiceRecordIdentity,
		IPInfoIdentity,
		IsolationProfileIdentity,
		IssueIdentity,
		JaegerbatchIdentity,
		LogIdentity,
		MessageIdentity,
		NamespaceIdentity,
		NamespaceMappingPolicyIdentity,
		NetworkAccessPolicyIdentity,
		OIDCProviderIdentity,
		PacketReportIdentity,
		PasswordResetIdentity,
		PlanIdentity,
		PokeIdentity,
		PolicyIdentity,
		PolicyRefreshIdentity,
		PolicyRendererIdentity,
		PolicyRuleIdentity,
		PrivateKeyIdentity,
		ProcessingUnitIdentity,
		ProcessingUnitPolicyIdentity,
		PUNodeIdentity,
		QuotaCheckIdentity,
		QuotaPolicyIdentity,
		RemoteProcessorIdentity,
		RenderedPolicyIdentity,
		ReportIdentity,
		RevocationIdentity,
		RoleIdentity,
		RootIdentity,
		ServiceIdentity,
		ServiceDependencyIdentity,
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
		TokenIdentity,
		TokenScopePolicyIdentity,
		TriggerIdentity,
		VulnerabilityIdentity,
		X509CertificateIdentity,
		X509CertificateCheckIdentity,
	}
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
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
	case AuthIdentity:
		return []string{}
	case AuthorityIdentity:
		return []string{
			"ca",
		}
	case AutomationIdentity:
		return []string{
			"autos",
			"auto",
		}
	case AutomationTemplateIdentity:
		return []string{
			"autotmpl",
		}
	case AWSAccountIdentity:
		return []string{
			"aws",
			"awsaccs",
			"awsacc",
		}
	case AWSAPIGatewayIdentity:
		return []string{}
	case AWSRegisterIdentity:
		return []string{}
	case CategoryIdentity:
		return []string{}
	case CustomerIdentity:
		return []string{}
	case DependencyMapIdentity:
		return []string{
			"depmaps",
			"depmap",
		}
	case EmailIdentity:
		return []string{}
	case EnforcerIdentity:
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
	case JaegerbatchIdentity:
		return []string{
			"sp",
		}
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
	case PolicyRefreshIdentity:
		return []string{}
	case PolicyRendererIdentity:
		return []string{}
	case PolicyRuleIdentity:
		return []string{}
	case PrivateKeyIdentity:
		return []string{}
	case ProcessingUnitIdentity:
		return []string{
			"pu",
			"pus",
		}
	case ProcessingUnitPolicyIdentity:
		return []string{
			"pup",
		}
	case PUNodeIdentity:
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
	case ReportIdentity:
		return []string{}
	case RevocationIdentity:
		return []string{}
	case RoleIdentity:
		return []string{}
	case RootIdentity:
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
	case TokenIdentity:
		return []string{}
	case TokenScopePolicyIdentity:
		return []string{
			"tsp",
		}
	case TriggerIdentity:
		return []string{}
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
