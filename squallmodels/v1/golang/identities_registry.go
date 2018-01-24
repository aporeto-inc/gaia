package squallmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(ActivityIdentity)
	elemental.RegisterIdentity(APIAuthorizationPolicyIdentity)
	elemental.RegisterIdentity(APICheckIdentity)
	elemental.RegisterIdentity(AuditProfileIdentity)
	elemental.RegisterIdentity(DependencyMapIdentity)
	elemental.RegisterIdentity(EnforcerIdentity)
	elemental.RegisterIdentity(EnforcerProfileIdentity)
	elemental.RegisterIdentity(EnforcerProfileMappingPolicyIdentity)
	elemental.RegisterIdentity(ExternalAccessIdentity)
	elemental.RegisterIdentity(ExternalServiceIdentity)
	elemental.RegisterIdentity(FileAccessIdentity)
	elemental.RegisterIdentity(FileAccessPolicyIdentity)
	elemental.RegisterIdentity(FilePathIdentity)
	elemental.RegisterIdentity(FlowStatisticIdentity)
	elemental.RegisterIdentity(HookPolicyIdentity)
	elemental.RegisterIdentity(MessageIdentity)
	elemental.RegisterIdentity(NamespaceIdentity)
	elemental.RegisterIdentity(NamespaceMappingPolicyIdentity)
	elemental.RegisterIdentity(NetworkAccessPolicyIdentity)
	elemental.RegisterIdentity(PokeIdentity)
	elemental.RegisterIdentity(PolicyIdentity)
	elemental.RegisterIdentity(PolicyRuleIdentity)
	elemental.RegisterIdentity(ProcessingUnitIdentity)
	elemental.RegisterIdentity(QuotaPolicyIdentity)
	elemental.RegisterIdentity(RenderedPolicyIdentity)
	elemental.RegisterIdentity(RoleIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(StatsQueryIdentity)
	elemental.RegisterIdentity(SuggestedPolicyIdentity)
	elemental.RegisterIdentity(SystemCallIdentity)
	elemental.RegisterIdentity(TabulationIdentity)
	elemental.RegisterIdentity(TagIdentity)
	elemental.RegisterIdentity(VulnerabilityIdentity)
}

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {

	case ActivityIdentity.Name:
		return NewActivity()
	case APIAuthorizationPolicyIdentity.Name:
		return NewAPIAuthorizationPolicy()
	case APICheckIdentity.Name:
		return NewAPICheck()
	case AuditProfileIdentity.Name:
		return NewAuditProfile()
	case DependencyMapIdentity.Name:
		return NewDependencyMap()
	case EnforcerIdentity.Name:
		return NewEnforcer()
	case EnforcerProfileIdentity.Name:
		return NewEnforcerProfile()
	case EnforcerProfileMappingPolicyIdentity.Name:
		return NewEnforcerProfileMappingPolicy()
	case ExternalAccessIdentity.Name:
		return NewExternalAccess()
	case ExternalServiceIdentity.Name:
		return NewExternalService()
	case FileAccessIdentity.Name:
		return NewFileAccess()
	case FileAccessPolicyIdentity.Name:
		return NewFileAccessPolicy()
	case FilePathIdentity.Name:
		return NewFilePath()
	case FlowStatisticIdentity.Name:
		return NewFlowStatistic()
	case HookPolicyIdentity.Name:
		return NewHookPolicy()
	case MessageIdentity.Name:
		return NewMessage()
	case NamespaceIdentity.Name:
		return NewNamespace()
	case NamespaceMappingPolicyIdentity.Name:
		return NewNamespaceMappingPolicy()
	case NetworkAccessPolicyIdentity.Name:
		return NewNetworkAccessPolicy()
	case PokeIdentity.Name:
		return NewPoke()
	case PolicyIdentity.Name:
		return NewPolicy()
	case PolicyRuleIdentity.Name:
		return NewPolicyRule()
	case ProcessingUnitIdentity.Name:
		return NewProcessingUnit()
	case QuotaPolicyIdentity.Name:
		return NewQuotaPolicy()
	case RenderedPolicyIdentity.Name:
		return NewRenderedPolicy()
	case RoleIdentity.Name:
		return NewRole()
	case RootIdentity.Name:
		return NewRoot()
	case StatsQueryIdentity.Name:
		return NewStatsQuery()
	case SuggestedPolicyIdentity.Name:
		return NewSuggestedPolicy()
	case SystemCallIdentity.Name:
		return NewSystemCall()
	case TabulationIdentity.Name:
		return NewTabulation()
	case TagIdentity.Name:
		return NewTag()
	case VulnerabilityIdentity.Name:
		return NewVulnerability()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {

	case ActivityIdentity.Category:
		return NewActivity()
	case APIAuthorizationPolicyIdentity.Category:
		return NewAPIAuthorizationPolicy()
	case APICheckIdentity.Category:
		return NewAPICheck()
	case AuditProfileIdentity.Category:
		return NewAuditProfile()
	case DependencyMapIdentity.Category:
		return NewDependencyMap()
	case EnforcerIdentity.Category:
		return NewEnforcer()
	case EnforcerProfileIdentity.Category:
		return NewEnforcerProfile()
	case EnforcerProfileMappingPolicyIdentity.Category:
		return NewEnforcerProfileMappingPolicy()
	case ExternalAccessIdentity.Category:
		return NewExternalAccess()
	case ExternalServiceIdentity.Category:
		return NewExternalService()
	case FileAccessIdentity.Category:
		return NewFileAccess()
	case FileAccessPolicyIdentity.Category:
		return NewFileAccessPolicy()
	case FilePathIdentity.Category:
		return NewFilePath()
	case FlowStatisticIdentity.Category:
		return NewFlowStatistic()
	case HookPolicyIdentity.Category:
		return NewHookPolicy()
	case MessageIdentity.Category:
		return NewMessage()
	case NamespaceIdentity.Category:
		return NewNamespace()
	case NamespaceMappingPolicyIdentity.Category:
		return NewNamespaceMappingPolicy()
	case NetworkAccessPolicyIdentity.Category:
		return NewNetworkAccessPolicy()
	case PokeIdentity.Category:
		return NewPoke()
	case PolicyIdentity.Category:
		return NewPolicy()
	case PolicyRuleIdentity.Category:
		return NewPolicyRule()
	case ProcessingUnitIdentity.Category:
		return NewProcessingUnit()
	case QuotaPolicyIdentity.Category:
		return NewQuotaPolicy()
	case RenderedPolicyIdentity.Category:
		return NewRenderedPolicy()
	case RoleIdentity.Category:
		return NewRole()
	case RootIdentity.Category:
		return NewRoot()
	case StatsQueryIdentity.Category:
		return NewStatsQuery()
	case SuggestedPolicyIdentity.Category:
		return NewSuggestedPolicy()
	case SystemCallIdentity.Category:
		return NewSystemCall()
	case TabulationIdentity.Category:
		return NewTabulation()
	case TagIdentity.Category:
		return NewTag()
	case VulnerabilityIdentity.Category:
		return NewVulnerability()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {

	case ActivityIdentity.Name:
		return &ActivitiesList{}
	case APIAuthorizationPolicyIdentity.Name:
		return &APIAuthorizationPoliciesList{}
	case APICheckIdentity.Name:
		return &APIChecksList{}
	case AuditProfileIdentity.Name:
		return &AuditProfilesList{}
	case DependencyMapIdentity.Name:
		return &DependencyMapsList{}
	case EnforcerIdentity.Name:
		return &EnforcersList{}
	case EnforcerProfileIdentity.Name:
		return &EnforcerProfilesList{}
	case EnforcerProfileMappingPolicyIdentity.Name:
		return &EnforcerProfileMappingPoliciesList{}
	case ExternalAccessIdentity.Name:
		return &ExternalAccessList{}
	case ExternalServiceIdentity.Name:
		return &ExternalServicesList{}
	case FileAccessIdentity.Name:
		return &FileAccessList{}
	case FileAccessPolicyIdentity.Name:
		return &FileAccessPoliciesList{}
	case FilePathIdentity.Name:
		return &FilePathsList{}
	case FlowStatisticIdentity.Name:
		return &FlowStatisticsList{}
	case HookPolicyIdentity.Name:
		return &HookPoliciesList{}
	case MessageIdentity.Name:
		return &MessagesList{}
	case NamespaceIdentity.Name:
		return &NamespacesList{}
	case NamespaceMappingPolicyIdentity.Name:
		return &NamespaceMappingPoliciesList{}
	case NetworkAccessPolicyIdentity.Name:
		return &NetworkAccessPoliciesList{}
	case PokeIdentity.Name:
		return &PokesList{}
	case PolicyIdentity.Name:
		return &PoliciesList{}
	case PolicyRuleIdentity.Name:
		return &PolicyRulesList{}
	case ProcessingUnitIdentity.Name:
		return &ProcessingUnitsList{}
	case QuotaPolicyIdentity.Name:
		return &QuotaPoliciesList{}
	case RenderedPolicyIdentity.Name:
		return &RenderedPoliciesList{}
	case RoleIdentity.Name:
		return &RolesList{}
	case RootIdentity.Name:
		return &RootsList{}
	case StatsQueryIdentity.Name:
		return &StatsQueriesList{}
	case SuggestedPolicyIdentity.Name:
		return &SuggestedPoliciesList{}
	case SystemCallIdentity.Name:
		return &SystemCallsList{}
	case TabulationIdentity.Name:
		return &TabulationsList{}
	case TagIdentity.Name:
		return &TagsList{}
	case VulnerabilityIdentity.Name:
		return &VulnerabilitiesList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {

	case ActivityIdentity.Category:
		return &ActivitiesList{}
	case APIAuthorizationPolicyIdentity.Category:
		return &APIAuthorizationPoliciesList{}
	case APICheckIdentity.Category:
		return &APIChecksList{}
	case AuditProfileIdentity.Category:
		return &AuditProfilesList{}
	case DependencyMapIdentity.Category:
		return &DependencyMapsList{}
	case EnforcerIdentity.Category:
		return &EnforcersList{}
	case EnforcerProfileIdentity.Category:
		return &EnforcerProfilesList{}
	case EnforcerProfileMappingPolicyIdentity.Category:
		return &EnforcerProfileMappingPoliciesList{}
	case ExternalAccessIdentity.Category:
		return &ExternalAccessList{}
	case ExternalServiceIdentity.Category:
		return &ExternalServicesList{}
	case FileAccessIdentity.Category:
		return &FileAccessList{}
	case FileAccessPolicyIdentity.Category:
		return &FileAccessPoliciesList{}
	case FilePathIdentity.Category:
		return &FilePathsList{}
	case FlowStatisticIdentity.Category:
		return &FlowStatisticsList{}
	case HookPolicyIdentity.Category:
		return &HookPoliciesList{}
	case MessageIdentity.Category:
		return &MessagesList{}
	case NamespaceIdentity.Category:
		return &NamespacesList{}
	case NamespaceMappingPolicyIdentity.Category:
		return &NamespaceMappingPoliciesList{}
	case NetworkAccessPolicyIdentity.Category:
		return &NetworkAccessPoliciesList{}
	case PokeIdentity.Category:
		return &PokesList{}
	case PolicyIdentity.Category:
		return &PoliciesList{}
	case PolicyRuleIdentity.Category:
		return &PolicyRulesList{}
	case ProcessingUnitIdentity.Category:
		return &ProcessingUnitsList{}
	case QuotaPolicyIdentity.Category:
		return &QuotaPoliciesList{}
	case RenderedPolicyIdentity.Category:
		return &RenderedPoliciesList{}
	case RoleIdentity.Category:
		return &RolesList{}
	case RootIdentity.Category:
		return &RootsList{}
	case StatsQueryIdentity.Category:
		return &StatsQueriesList{}
	case SuggestedPolicyIdentity.Category:
		return &SuggestedPoliciesList{}
	case SystemCallIdentity.Category:
		return &SystemCallsList{}
	case TabulationIdentity.Category:
		return &TabulationsList{}
	case TagIdentity.Category:
		return &TagsList{}
	case VulnerabilityIdentity.Category:
		return &VulnerabilitiesList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		ActivityIdentity,
		APIAuthorizationPolicyIdentity,
		APICheckIdentity,
		AuditProfileIdentity,
		DependencyMapIdentity,
		EnforcerIdentity,
		EnforcerProfileIdentity,
		EnforcerProfileMappingPolicyIdentity,
		ExternalAccessIdentity,
		ExternalServiceIdentity,
		FileAccessIdentity,
		FileAccessPolicyIdentity,
		FilePathIdentity,
		FlowStatisticIdentity,
		HookPolicyIdentity,
		MessageIdentity,
		NamespaceIdentity,
		NamespaceMappingPolicyIdentity,
		NetworkAccessPolicyIdentity,
		PokeIdentity,
		PolicyIdentity,
		PolicyRuleIdentity,
		ProcessingUnitIdentity,
		QuotaPolicyIdentity,
		RenderedPolicyIdentity,
		RoleIdentity,
		RootIdentity,
		StatsQueryIdentity,
		SuggestedPolicyIdentity,
		SystemCallIdentity,
		TabulationIdentity,
		TagIdentity,
		VulnerabilityIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"apiauths":   APIAuthorizationPolicyIdentity,
	"apiauth":    APIAuthorizationPolicyIdentity,
	"ap":         AuditProfileIdentity,
	"depmap":     DependencyMapIdentity,
	"depmaps":    DependencyMapIdentity,
	"profile":    EnforcerProfileIdentity,
	"profiles":   EnforcerProfileIdentity,
	"srvpols":    EnforcerProfileMappingPolicyIdentity,
	"srvpol":     EnforcerProfileMappingPolicyIdentity,
	"extac":      ExternalAccessIdentity,
	"extacs":     ExternalAccessIdentity,
	"extsrv":     ExternalServiceIdentity,
	"extsrvs":    ExternalServiceIdentity,
	"fp":         FilePathIdentity,
	"fps":        FilePathIdentity,
	"flowstats":  FlowStatisticIdentity,
	"flowstat":   FlowStatisticIdentity,
	"hook":       HookPolicyIdentity,
	"hooks":      HookPolicyIdentity,
	"hookpol":    HookPolicyIdentity,
	"hookpols":   HookPolicyIdentity,
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
	"quota":      QuotaPolicyIdentity,
	"quotas":     QuotaPolicyIdentity,
	"quotapol":   QuotaPolicyIdentity,
	"quotapols":  QuotaPolicyIdentity,
	"rpol":       RenderedPolicyIdentity,
	"rpols":      RenderedPolicyIdentity,
	"sq":         StatsQueryIdentity,
	"sugpol":     SuggestedPolicyIdentity,
	"sugpols":    SuggestedPolicyIdentity,
	"sugg":       SuggestedPolicyIdentity,
	"suggs":      SuggestedPolicyIdentity,
	"table":      TabulationIdentity,
	"tables":     TabulationIdentity,
	"tabs":       TabulationIdentity,
	"tab":        TabulationIdentity,
	"vulns":      VulnerabilityIdentity,
	"vul":        VulnerabilityIdentity,
	"vuln":       VulnerabilityIdentity,
	"vuls":       VulnerabilityIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case ActivityIdentity:
		return []string{}
	case APIAuthorizationPolicyIdentity:
		return []string{
			"apiauths",
			"apiauth",
		}
	case APICheckIdentity:
		return []string{}
	case AuditProfileIdentity:
		return []string{
			"ap",
		}
	case DependencyMapIdentity:
		return []string{
			"depmap",
			"depmaps",
		}
	case EnforcerIdentity:
		return []string{}
	case EnforcerProfileIdentity:
		return []string{
			"profile",
			"profiles",
		}
	case EnforcerProfileMappingPolicyIdentity:
		return []string{
			"srvpols",
			"srvpol",
		}
	case ExternalAccessIdentity:
		return []string{
			"extac",
			"extacs",
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
	case FilePathIdentity:
		return []string{
			"fp",
			"fps",
		}
	case FlowStatisticIdentity:
		return []string{
			"flowstats",
			"flowstat",
		}
	case HookPolicyIdentity:
		return []string{
			"hook",
			"hooks",
			"hookpol",
			"hookpols",
		}
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
	case PokeIdentity:
		return []string{}
	case PolicyIdentity:
		return []string{}
	case PolicyRuleIdentity:
		return []string{}
	case ProcessingUnitIdentity:
		return []string{
			"pu",
			"pus",
		}
	case QuotaPolicyIdentity:
		return []string{
			"quota",
			"quotas",
			"quotapol",
			"quotapols",
		}
	case RenderedPolicyIdentity:
		return []string{
			"rpol",
			"rpols",
		}
	case RoleIdentity:
		return []string{}
	case RootIdentity:
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
	case SystemCallIdentity:
		return []string{}
	case TabulationIdentity:
		return []string{
			"table",
			"tables",
			"tabs",
			"tab",
		}
	case TagIdentity:
		return []string{}
	case VulnerabilityIdentity:
		return []string{
			"vulns",
			"vul",
			"vuln",
			"vuls",
		}
	}

	return nil
}
