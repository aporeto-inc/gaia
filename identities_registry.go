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
		"auditprofile":           AuditProfileIdentity,
		"auditreport":            AuditReportIdentity,
		"auth":                   AuthIdentity,
		"authority":              AuthorityIdentity,
		"automation":             AutomationIdentity,
		"automationtemplate":     AutomationTemplateIdentity,
		"awsaccount":             AWSAccountIdentity,
		"awsapigateway":          AWSAPIGatewayIdentity,
		"awsregister":            AWSRegisterIdentity,
		"category":               CategoryIdentity,
		"certificate":            CertificateIdentity,

		"customer":                     CustomerIdentity,
		"dependencymap":                DependencyMapIdentity,
		"email":                        EmailIdentity,
		"enforcer":                     EnforcerIdentity,
		"enforcerprofile":              EnforcerProfileIdentity,
		"enforcerprofilemappingpolicy": EnforcerProfileMappingPolicyIdentity,
		"enforcerreport":               EnforcerReportIdentity,
		"eventlog":                     EventLogIdentity,
		"export":                       ExportIdentity,
		"externalnetwork":              ExternalNetworkIdentity,
		"externalservice":              ExternalServiceIdentity,
		"fileaccess":                   FileAccessIdentity,
		"fileaccesspolicy":             FileAccessPolicyIdentity,
		"fileaccessreport":             FileAccessReportIdentity,
		"filepath":                     FilePathIdentity,
		"flowreport":                   FlowReportIdentity,

		"hookpolicy":             HookPolicyIdentity,
		"import":                 ImportIdentity,
		"installation":           InstallationIdentity,
		"installedapp":           InstalledAppIdentity,
		"invoice":                InvoiceIdentity,
		"invoicerecord":          InvoiceRecordIdentity,
		"isolationprofile":       IsolationProfileIdentity,
		"issue":                  IssueIdentity,
		"jaegerbatch":            JaegerbatchIdentity,
		"k8scluster":             K8SClusterIdentity,
		"log":                    LogIdentity,
		"message":                MessageIdentity,
		"namespace":              NamespaceIdentity,
		"namespacemappingpolicy": NamespaceMappingPolicyIdentity,
		"networkaccesspolicy":    NetworkAccessPolicyIdentity,
		"passwordreset":          PasswordResetIdentity,
		"plan":                   PlanIdentity,
		"poke":                   PokeIdentity,
		"policy":                 PolicyIdentity,
		"policyrefresh":          PolicyRefreshIdentity,
		"policyrenderer":         PolicyRendererIdentity,
		"policyrule":             PolicyRuleIdentity,
		"privatekey":             PrivateKeyIdentity,
		"processingunit":         ProcessingUnitIdentity,
		"processingunitpolicy":   ProcessingUnitPolicyIdentity,
		"punode":                 PUNodeIdentity,
		"quotacheck":             QuotaCheckIdentity,
		"quotapolicy":            QuotaPolicyIdentity,
		"remoteprocessor":        RemoteProcessorIdentity,
		"renderedpolicy":         RenderedPolicyIdentity,
		"report":                 ReportIdentity,
		"restapispec":            RESTAPISpecIdentity,
		"revocation":             RevocationIdentity,
		"role":                   RoleIdentity,
		"root":                   RootIdentity,
		"service":                ServiceIdentity,
		"servicedependency":      ServiceDependencyIdentity,
		"squalltag":              SquallTagIdentity,
		"statsquery":             StatsQueryIdentity,
		"suggestedpolicy":        SuggestedPolicyIdentity,
		"tabulation":             TabulationIdentity,
		"tag":                    TagIdentity,
		"taginject":              TagInjectIdentity,
		"token":                  TokenIdentity,
		"tokenscopepolicy":       TokenScopePolicyIdentity,
		"trigger":                TriggerIdentity,
		"vulnerability":          VulnerabilityIdentity,
		"x509certificate":        X509CertificateIdentity,
		"x509certificatecheck":   X509CertificateCheckIdentity,
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
		"auditprofiles":            AuditProfileIdentity,
		"auditreports":             AuditReportIdentity,
		"auth":                     AuthIdentity,
		"authorities":              AuthorityIdentity,
		"automations":              AutomationIdentity,
		"automationtemplates":      AutomationTemplateIdentity,
		"awsaccounts":              AWSAccountIdentity,
		"awsapigateways":           AWSAPIGatewayIdentity,
		"awsregister":              AWSRegisterIdentity,
		"categories":               CategoryIdentity,
		"certificates":             CertificateIdentity,

		"customers":                      CustomerIdentity,
		"dependencymaps":                 DependencyMapIdentity,
		"emails":                         EmailIdentity,
		"enforcers":                      EnforcerIdentity,
		"enforcerprofiles":               EnforcerProfileIdentity,
		"enforcerprofilemappingpolicies": EnforcerProfileMappingPolicyIdentity,
		"enforcerreports":                EnforcerReportIdentity,
		"eventlogs":                      EventLogIdentity,
		"export":                         ExportIdentity,
		"externalnetworks":               ExternalNetworkIdentity,
		"externalservices":               ExternalServiceIdentity,
		"fileaccesses":                   FileAccessIdentity,
		"fileaccesspolicies":             FileAccessPolicyIdentity,
		"fileaccessreports":              FileAccessReportIdentity,
		"filepaths":                      FilePathIdentity,
		"flowreports":                    FlowReportIdentity,

		"hookpolicies":             HookPolicyIdentity,
		"import":                   ImportIdentity,
		"installations":            InstallationIdentity,
		"installedapps":            InstalledAppIdentity,
		"invoices":                 InvoiceIdentity,
		"invoicerecords":           InvoiceRecordIdentity,
		"isolationprofiles":        IsolationProfileIdentity,
		"issue":                    IssueIdentity,
		"jaegerbatchs":             JaegerbatchIdentity,
		"k8sclusters":              K8SClusterIdentity,
		"logs":                     LogIdentity,
		"messages":                 MessageIdentity,
		"namespaces":               NamespaceIdentity,
		"namespacemappingpolicies": NamespaceMappingPolicyIdentity,
		"networkaccesspolicies":    NetworkAccessPolicyIdentity,
		"passwordreset":            PasswordResetIdentity,
		"plans":                    PlanIdentity,
		"poke":                     PokeIdentity,
		"policies":                 PolicyIdentity,
		"policyrefreshs":           PolicyRefreshIdentity,
		"policyrenderers":          PolicyRendererIdentity,
		"policyrules":              PolicyRuleIdentity,
		"privatekeys":              PrivateKeyIdentity,
		"processingunits":          ProcessingUnitIdentity,
		"processingunitpolicies":   ProcessingUnitPolicyIdentity,
		"punodes":                  PUNodeIdentity,
		"quotacheck":               QuotaCheckIdentity,
		"quotapolicies":            QuotaPolicyIdentity,
		"remoteprocessors":         RemoteProcessorIdentity,
		"renderedpolicies":         RenderedPolicyIdentity,
		"reports":                  ReportIdentity,
		"restapispecs":             RESTAPISpecIdentity,
		"revocations":              RevocationIdentity,
		"roles":                    RoleIdentity,
		"root":                     RootIdentity,
		"services":                 ServiceIdentity,
		"servicedependencies":      ServiceDependencyIdentity,
		"squalltags":               SquallTagIdentity,
		"statsqueries":             StatsQueryIdentity,
		"suggestedpolicies":        SuggestedPolicyIdentity,
		"tabulations":              TabulationIdentity,
		"tags":                     TagIdentity,
		"taginjects":               TagInjectIdentity,
		"tokens":                   TokenIdentity,
		"tokenscopepolicies":       TokenScopePolicyIdentity,
		"triggers":                 TriggerIdentity,
		"vulnerabilities":          VulnerabilityIdentity,
		"x509certificates":         X509CertificateIdentity,
		"x509certificatechecks":    X509CertificateCheckIdentity,
	}

	aliasesMap = map[string]elemental.Identity{
		"apiauth":    APIAuthorizationPolicyIdentity,
		"apiauths":   APIAuthorizationPolicyIdentity,
		"ap":         AuditProfileIdentity,
		"ca":         AuthorityIdentity,
		"autos":      AutomationIdentity,
		"auto":       AutomationIdentity,
		"autotmpl":   AutomationTemplateIdentity,
		"aws":        AWSAccountIdentity,
		"awsaccs":    AWSAccountIdentity,
		"awsacc":     AWSAccountIdentity,
		"depmaps":    DependencyMapIdentity,
		"depmap":     DependencyMapIdentity,
		"profile":    EnforcerProfileIdentity,
		"profiles":   EnforcerProfileIdentity,
		"enfpols":    EnforcerProfileMappingPolicyIdentity,
		"enfpol":     EnforcerProfileMappingPolicyIdentity,
		"extnet":     ExternalNetworkIdentity,
		"extnets":    ExternalNetworkIdentity,
		"extsrv":     ExternalServiceIdentity,
		"extsrvs":    ExternalServiceIdentity,
		"fp":         FilePathIdentity,
		"fps":        FilePathIdentity,
		"hook":       HookPolicyIdentity,
		"hooks":      HookPolicyIdentity,
		"hookpol":    HookPolicyIdentity,
		"hookpols":   HookPolicyIdentity,
		"iapps":      InstalledAppIdentity,
		"iapp":       InstalledAppIdentity,
		"ip":         IsolationProfileIdentity,
		"sp":         JaegerbatchIdentity,
		"mess":       MessageIdentity,
		"ns":         NamespaceIdentity,
		"nspolicy":   NamespaceMappingPolicyIdentity,
		"nspolicies": NamespaceMappingPolicyIdentity,
		"nsmap":      NamespaceMappingPolicyIdentity,
		"nsmaps":     NamespaceMappingPolicyIdentity,
		"netpol":     NetworkAccessPolicyIdentity,
		"netpols":    NetworkAccessPolicyIdentity,
		"pu":         ProcessingUnitIdentity,
		"pus":        ProcessingUnitIdentity,
		"pup":        ProcessingUnitPolicyIdentity,
		"quota":      QuotaPolicyIdentity,
		"quotas":     QuotaPolicyIdentity,
		"quotapol":   QuotaPolicyIdentity,
		"quotapols":  QuotaPolicyIdentity,
		"hks":        RemoteProcessorIdentity,
		"hk":         RemoteProcessorIdentity,
		"rpol":       RenderedPolicyIdentity,
		"rpols":      RenderedPolicyIdentity,
		"srv":        ServiceIdentity,
		"srvdep":     ServiceDependencyIdentity,
		"srvdeps":    ServiceDependencyIdentity,
		"sq":         StatsQueryIdentity,
		"sugpol":     SuggestedPolicyIdentity,
		"sugpols":    SuggestedPolicyIdentity,
		"sugg":       SuggestedPolicyIdentity,
		"suggs":      SuggestedPolicyIdentity,
		"table":      TabulationIdentity,
		"tables":     TabulationIdentity,
		"tabs":       TabulationIdentity,
		"tab":        TabulationIdentity,
		"tsp":        TokenScopePolicyIdentity,
		"vulns":      VulnerabilityIdentity,
		"vul":        VulnerabilityIdentity,
		"vuln":       VulnerabilityIdentity,
		"vuls":       VulnerabilityIdentity,
	}

	indexesMap = map[string][][]string{
		"account": [][]string{
			[]string{":unique", "name"},
			[]string{":unique", "email"},
			[]string{"activationtoken"},
			[]string{"resetpasswordtoken"},
		},
		"accountcheck": nil,
		"activate":     nil,
		"activity": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "date"},
			[]string{"namespace", "operation"},
		},
		"alarm": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "kind"},
		},
		"apiauthorizationpolicy": nil,
		"apicheck":               nil,
		"app":                    nil,
		"auditprofile": [][]string{
			[]string{"namespace"},
		},
		"auditreport": nil,
		"auth":        nil,
		"authority": [][]string{
			[]string{"commonName"},
			[]string{"serialNumber"},
		},
		"automation": [][]string{
			[]string{"namespace"},
		},
		"automationtemplate": nil,
		"awsaccount":         nil,
		"awsapigateway":      nil,
		"awsregister":        nil,
		"category":           nil,
		"certificate": [][]string{
			[]string{"commonName"},
			[]string{":unique", "parentID", "accountID"},
			[]string{"parentID", "commonName"},
		},
		"customer":      nil,
		"dependencymap": nil,
		"email":         nil,
		"enforcer": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
		},
		"enforcerprofile": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
		},
		"enforcerprofilemappingpolicy": nil,
		"enforcerreport":               nil,
		"eventlog":                     nil,
		"export":                       nil,
		"externalnetwork": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"namespace", "archived"},
		},
		"externalservice":  nil,
		"fileaccess":       nil,
		"fileaccesspolicy": nil,
		"fileaccessreport": nil,
		"filepath": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
			[]string{"archived"},
		},
		"flowreport":    nil,
		"hookpolicy":    nil,
		"import":        nil,
		"installation":  nil,
		"installedapp":  nil,
		"invoice":       nil,
		"invoicerecord": nil,
		"isolationprofile": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
		},
		"issue":       nil,
		"jaegerbatch": nil,
		"k8scluster": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "normalizedTags"},
		},
		"log":     nil,
		"message": nil,
		"namespace": [][]string{
			[]string{"namespace"},
			[]string{":unique", "namespace", "name"},
			[]string{"namespace", "normalizedTags"},
		},
		"namespacemappingpolicy": nil,
		"networkaccesspolicy":    nil,
		"passwordreset":          nil,
		"plan":                   nil,
		"poke":                   nil,
		"policy": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "type"},
			[]string{"namespace", "normalizedtags"},
			[]string{"namespace", "type", "allobjecttags"},
			[]string{"namespace", "type", "allsubjecttags"},
			[]string{"namespace", "type", "allobjecttags", "disabled"},
			[]string{"namespace", "type", "allsubjecttags", "disabled"},
			[]string{"namespace", "type", "allobjecttags", "propagated"},
			[]string{"namespace", "type", "allsubjecttags", "propagated"},
			[]string{"namespace", "fallback"},
		},
		"policyrefresh":  nil,
		"policyrenderer": nil,
		"policyrule":     nil,
		"privatekey":     nil,
		"processingunit": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "archived"},
			[]string{"namespace", "status"},
		},
		"processingunitpolicy": nil,
		"punode":               nil,
		"quotacheck":           nil,
		"quotapolicy":          nil,
		"remoteprocessor":      nil,
		"renderedpolicy":       nil,
		"report":               nil,
		"restapispec": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "archived"},
			[]string{"namespace", "normalizedtags"},
		},
		"revocation": [][]string{
			[]string{":unique", "serialNumber"},
		},
		"role": nil,
		"root": nil,
		"service": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "archived"},
			[]string{"namespace", "normalizedtags"},
		},
		"servicedependency": [][]string{
			[]string{"namespace"},
			[]string{"namespace", "archived"},
			[]string{"namespace", "normalizedtags"},
		},
		"squalltag":        nil,
		"statsquery":       nil,
		"suggestedpolicy":  nil,
		"tabulation":       nil,
		"tag":              nil,
		"taginject":        nil,
		"token":            nil,
		"tokenscopepolicy": nil,
		"trigger":          nil,
		"vulnerability": [][]string{
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
	case AuditProfileIdentity:
		return NewAuditProfile()
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
	case CertificateIdentity:
		return NewCertificate()
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
	case EventLogIdentity:
		return NewEventLog()
	case ExportIdentity:
		return NewExport()
	case ExternalNetworkIdentity:
		return NewExternalNetwork()
	case ExternalServiceIdentity:
		return NewExternalService()
	case FileAccessIdentity:
		return NewFileAccess()
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
	case ImportIdentity:
		return NewImport()
	case InstallationIdentity:
		return NewInstallation()
	case InstalledAppIdentity:
		return NewInstalledApp()
	case InvoiceIdentity:
		return NewInvoice()
	case InvoiceRecordIdentity:
		return NewInvoiceRecord()
	case IsolationProfileIdentity:
		return NewIsolationProfile()
	case IssueIdentity:
		return NewIssue()
	case JaegerbatchIdentity:
		return NewJaegerbatch()
	case K8SClusterIdentity:
		return NewK8SCluster()
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
	case RESTAPISpecIdentity:
		return NewRESTAPISpec()
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
	case AuditProfileIdentity:
		return &AuditProfilesList{}
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
	case CertificateIdentity:
		return &CertificatesList{}
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
	case EventLogIdentity:
		return &EventLogsList{}
	case ExportIdentity:
		return &ExportsList{}
	case ExternalNetworkIdentity:
		return &ExternalNetworksList{}
	case ExternalServiceIdentity:
		return &ExternalServicesList{}
	case FileAccessIdentity:
		return &FileAccessList{}
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
	case ImportIdentity:
		return &ImportsList{}
	case InstallationIdentity:
		return &InstallationsList{}
	case InstalledAppIdentity:
		return &InstalledAppsList{}
	case InvoiceIdentity:
		return &InvoicesList{}
	case InvoiceRecordIdentity:
		return &InvoiceRecordsList{}
	case IsolationProfileIdentity:
		return &IsolationProfilesList{}
	case IssueIdentity:
		return &IssuesList{}
	case JaegerbatchIdentity:
		return &JaegerbatchsList{}
	case K8SClusterIdentity:
		return &K8SClustersList{}
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
	case RESTAPISpecIdentity:
		return &RESTAPISpecsList{}
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
		AuditProfileIdentity,
		AuditReportIdentity,
		AuthIdentity,
		AuthorityIdentity,
		AutomationIdentity,
		AutomationTemplateIdentity,
		AWSAccountIdentity,
		AWSAPIGatewayIdentity,
		AWSRegisterIdentity,
		CategoryIdentity,
		CertificateIdentity,
		CustomerIdentity,
		DependencyMapIdentity,
		EmailIdentity,
		EnforcerIdentity,
		EnforcerProfileIdentity,
		EnforcerProfileMappingPolicyIdentity,
		EnforcerReportIdentity,
		EventLogIdentity,
		ExportIdentity,
		ExternalNetworkIdentity,
		ExternalServiceIdentity,
		FileAccessIdentity,
		FileAccessPolicyIdentity,
		FileAccessReportIdentity,
		FilePathIdentity,
		FlowReportIdentity,
		HookPolicyIdentity,
		ImportIdentity,
		InstallationIdentity,
		InstalledAppIdentity,
		InvoiceIdentity,
		InvoiceRecordIdentity,
		IsolationProfileIdentity,
		IssueIdentity,
		JaegerbatchIdentity,
		K8SClusterIdentity,
		LogIdentity,
		MessageIdentity,
		NamespaceIdentity,
		NamespaceMappingPolicyIdentity,
		NetworkAccessPolicyIdentity,
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
		RESTAPISpecIdentity,
		RevocationIdentity,
		RoleIdentity,
		RootIdentity,
		ServiceIdentity,
		ServiceDependencyIdentity,
		SquallTagIdentity,
		StatsQueryIdentity,
		SuggestedPolicyIdentity,
		TabulationIdentity,
		TagIdentity,
		TagInjectIdentity,
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
	case AuditProfileIdentity:
		return []string{
			"ap",
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
	case CertificateIdentity:
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
		}
	case EnforcerReportIdentity:
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
	case ExternalServiceIdentity:
		return []string{
			"extsrv",
			"extsrvs",
		}
	case FileAccessIdentity:
		return []string{}
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
	case ImportIdentity:
		return []string{}
	case InstallationIdentity:
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
	case K8SClusterIdentity:
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
	case RESTAPISpecIdentity:
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
