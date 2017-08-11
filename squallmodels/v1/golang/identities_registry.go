package squallmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(APICheckIdentity)
	elemental.RegisterIdentity(HookPolicyIdentity)
	elemental.RegisterIdentity(APIAuthorizationPolicyIdentity)
	elemental.RegisterIdentity(ExternalAccessIdentity)
	elemental.RegisterIdentity(SystemCallIdentity)
	elemental.RegisterIdentity(TagIdentity)
	elemental.RegisterIdentity(EnforcerIdentity)
	elemental.RegisterIdentity(MessageIdentity)
	elemental.RegisterIdentity(PokeIdentity)
	elemental.RegisterIdentity(FilePathIdentity)
	elemental.RegisterIdentity(FileAccessIdentity)
	elemental.RegisterIdentity(NamespaceIdentity)
	elemental.RegisterIdentity(PolicyRuleIdentity)
	elemental.RegisterIdentity(ExternalServiceIdentity)
	elemental.RegisterIdentity(RoleIdentity)
	elemental.RegisterIdentity(PolicyIdentity)
	elemental.RegisterIdentity(FlowStatisticIdentity)
	elemental.RegisterIdentity(SuggestedPolicyIdentity)
	elemental.RegisterIdentity(TabulationIdentity)
	elemental.RegisterIdentity(QuotaPolicyIdentity)
	elemental.RegisterIdentity(FileAccessPolicyIdentity)
	elemental.RegisterIdentity(EnforcerProfileIdentity)
	elemental.RegisterIdentity(RenderedPolicyIdentity)
	elemental.RegisterIdentity(NamespaceMappingPolicyIdentity)
	elemental.RegisterIdentity(ProcessingUnitIdentity)
	elemental.RegisterIdentity(DependencyMapIdentity)
	elemental.RegisterIdentity(VulnerabilityIdentity)
	elemental.RegisterIdentity(EnforcerProfileMappingPolicyIdentity)
	elemental.RegisterIdentity(ActivityIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(NetworkAccessPolicyIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case APICheckIdentity.Name:
		return NewAPICheck()
	case HookPolicyIdentity.Name:
		return NewHookPolicy()
	case APIAuthorizationPolicyIdentity.Name:
		return NewAPIAuthorizationPolicy()
	case ExternalAccessIdentity.Name:
		return NewExternalAccess()
	case SystemCallIdentity.Name:
		return NewSystemCall()
	case TagIdentity.Name:
		return NewTag()
	case EnforcerIdentity.Name:
		return NewEnforcer()
	case MessageIdentity.Name:
		return NewMessage()
	case PokeIdentity.Name:
		return NewPoke()
	case FilePathIdentity.Name:
		return NewFilePath()
	case FileAccessIdentity.Name:
		return NewFileAccess()
	case NamespaceIdentity.Name:
		return NewNamespace()
	case PolicyRuleIdentity.Name:
		return NewPolicyRule()
	case ExternalServiceIdentity.Name:
		return NewExternalService()
	case RoleIdentity.Name:
		return NewRole()
	case PolicyIdentity.Name:
		return NewPolicy()
	case FlowStatisticIdentity.Name:
		return NewFlowStatistic()
	case SuggestedPolicyIdentity.Name:
		return NewSuggestedPolicy()
	case TabulationIdentity.Name:
		return NewTabulation()
	case QuotaPolicyIdentity.Name:
		return NewQuotaPolicy()
	case FileAccessPolicyIdentity.Name:
		return NewFileAccessPolicy()
	case EnforcerProfileIdentity.Name:
		return NewEnforcerProfile()
	case RenderedPolicyIdentity.Name:
		return NewRenderedPolicy()
	case NamespaceMappingPolicyIdentity.Name:
		return NewNamespaceMappingPolicy()
	case ProcessingUnitIdentity.Name:
		return NewProcessingUnit()
	case DependencyMapIdentity.Name:
		return NewDependencyMap()
	case VulnerabilityIdentity.Name:
		return NewVulnerability()
	case EnforcerProfileMappingPolicyIdentity.Name:
		return NewEnforcerProfileMappingPolicy()
	case ActivityIdentity.Name:
		return NewActivity()
	case RootIdentity.Name:
		return NewRoot()
	case NetworkAccessPolicyIdentity.Name:
		return NewNetworkAccessPolicy()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {
	case APICheckIdentity.Category:
		return NewAPICheck()
	case HookPolicyIdentity.Category:
		return NewHookPolicy()
	case APIAuthorizationPolicyIdentity.Category:
		return NewAPIAuthorizationPolicy()
	case ExternalAccessIdentity.Category:
		return NewExternalAccess()
	case SystemCallIdentity.Category:
		return NewSystemCall()
	case TagIdentity.Category:
		return NewTag()
	case EnforcerIdentity.Category:
		return NewEnforcer()
	case MessageIdentity.Category:
		return NewMessage()
	case PokeIdentity.Category:
		return NewPoke()
	case FilePathIdentity.Category:
		return NewFilePath()
	case FileAccessIdentity.Category:
		return NewFileAccess()
	case NamespaceIdentity.Category:
		return NewNamespace()
	case PolicyRuleIdentity.Category:
		return NewPolicyRule()
	case ExternalServiceIdentity.Category:
		return NewExternalService()
	case RoleIdentity.Category:
		return NewRole()
	case PolicyIdentity.Category:
		return NewPolicy()
	case FlowStatisticIdentity.Category:
		return NewFlowStatistic()
	case SuggestedPolicyIdentity.Category:
		return NewSuggestedPolicy()
	case TabulationIdentity.Category:
		return NewTabulation()
	case QuotaPolicyIdentity.Category:
		return NewQuotaPolicy()
	case FileAccessPolicyIdentity.Category:
		return NewFileAccessPolicy()
	case EnforcerProfileIdentity.Category:
		return NewEnforcerProfile()
	case RenderedPolicyIdentity.Category:
		return NewRenderedPolicy()
	case NamespaceMappingPolicyIdentity.Category:
		return NewNamespaceMappingPolicy()
	case ProcessingUnitIdentity.Category:
		return NewProcessingUnit()
	case DependencyMapIdentity.Category:
		return NewDependencyMap()
	case VulnerabilityIdentity.Category:
		return NewVulnerability()
	case EnforcerProfileMappingPolicyIdentity.Category:
		return NewEnforcerProfileMappingPolicy()
	case ActivityIdentity.Category:
		return NewActivity()
	case RootIdentity.Category:
		return NewRoot()
	case NetworkAccessPolicyIdentity.Category:
		return NewNetworkAccessPolicy()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case APICheckIdentity.Name:
		return &APIChecksList{}
	case HookPolicyIdentity.Name:
		return &HookPoliciesList{}
	case APIAuthorizationPolicyIdentity.Name:
		return &APIAuthorizationPoliciesList{}
	case ExternalAccessIdentity.Name:
		return &ExternalAccessList{}
	case SystemCallIdentity.Name:
		return &SystemCallsList{}
	case TagIdentity.Name:
		return &TagsList{}
	case EnforcerIdentity.Name:
		return &EnforcersList{}
	case MessageIdentity.Name:
		return &MessagesList{}
	case PokeIdentity.Name:
		return &PokesList{}
	case FilePathIdentity.Name:
		return &FilePathsList{}
	case FileAccessIdentity.Name:
		return &FileAccessList{}
	case NamespaceIdentity.Name:
		return &NamespacesList{}
	case PolicyRuleIdentity.Name:
		return &PolicyRulesList{}
	case ExternalServiceIdentity.Name:
		return &ExternalServicesList{}
	case RoleIdentity.Name:
		return &RolesList{}
	case PolicyIdentity.Name:
		return &PoliciesList{}
	case FlowStatisticIdentity.Name:
		return &FlowStatisticsList{}
	case SuggestedPolicyIdentity.Name:
		return &SuggestedPoliciesList{}
	case TabulationIdentity.Name:
		return &TabulationsList{}
	case QuotaPolicyIdentity.Name:
		return &QuotaPoliciesList{}
	case FileAccessPolicyIdentity.Name:
		return &FileAccessPoliciesList{}
	case EnforcerProfileIdentity.Name:
		return &EnforcerProfilesList{}
	case RenderedPolicyIdentity.Name:
		return &RenderedPoliciesList{}
	case NamespaceMappingPolicyIdentity.Name:
		return &NamespaceMappingPoliciesList{}
	case ProcessingUnitIdentity.Name:
		return &ProcessingUnitsList{}
	case DependencyMapIdentity.Name:
		return &DependencyMapsList{}
	case VulnerabilityIdentity.Name:
		return &VulnerabilitiesList{}
	case EnforcerProfileMappingPolicyIdentity.Name:
		return &EnforcerProfileMappingPoliciesList{}
	case ActivityIdentity.Name:
		return &ActivitiesList{}
	case NetworkAccessPolicyIdentity.Name:
		return &NetworkAccessPoliciesList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {
	case APICheckIdentity.Category:
		return &APIChecksList{}
	case HookPolicyIdentity.Category:
		return &HookPoliciesList{}
	case APIAuthorizationPolicyIdentity.Category:
		return &APIAuthorizationPoliciesList{}
	case ExternalAccessIdentity.Category:
		return &ExternalAccessList{}
	case SystemCallIdentity.Category:
		return &SystemCallsList{}
	case TagIdentity.Category:
		return &TagsList{}
	case EnforcerIdentity.Category:
		return &EnforcersList{}
	case MessageIdentity.Category:
		return &MessagesList{}
	case PokeIdentity.Category:
		return &PokesList{}
	case FilePathIdentity.Category:
		return &FilePathsList{}
	case FileAccessIdentity.Category:
		return &FileAccessList{}
	case NamespaceIdentity.Category:
		return &NamespacesList{}
	case PolicyRuleIdentity.Category:
		return &PolicyRulesList{}
	case ExternalServiceIdentity.Category:
		return &ExternalServicesList{}
	case RoleIdentity.Category:
		return &RolesList{}
	case PolicyIdentity.Category:
		return &PoliciesList{}
	case FlowStatisticIdentity.Category:
		return &FlowStatisticsList{}
	case SuggestedPolicyIdentity.Category:
		return &SuggestedPoliciesList{}
	case TabulationIdentity.Category:
		return &TabulationsList{}
	case QuotaPolicyIdentity.Category:
		return &QuotaPoliciesList{}
	case FileAccessPolicyIdentity.Category:
		return &FileAccessPoliciesList{}
	case EnforcerProfileIdentity.Category:
		return &EnforcerProfilesList{}
	case RenderedPolicyIdentity.Category:
		return &RenderedPoliciesList{}
	case NamespaceMappingPolicyIdentity.Category:
		return &NamespaceMappingPoliciesList{}
	case ProcessingUnitIdentity.Category:
		return &ProcessingUnitsList{}
	case DependencyMapIdentity.Category:
		return &DependencyMapsList{}
	case VulnerabilityIdentity.Category:
		return &VulnerabilitiesList{}
	case EnforcerProfileMappingPolicyIdentity.Category:
		return &EnforcerProfileMappingPoliciesList{}
	case ActivityIdentity.Category:
		return &ActivitiesList{}
	case NetworkAccessPolicyIdentity.Category:
		return &NetworkAccessPoliciesList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		APICheckIdentity,
		HookPolicyIdentity,
		APIAuthorizationPolicyIdentity,
		ExternalAccessIdentity,
		SystemCallIdentity,
		TagIdentity,
		EnforcerIdentity,
		MessageIdentity,
		PokeIdentity,
		FilePathIdentity,
		FileAccessIdentity,
		NamespaceIdentity,
		PolicyRuleIdentity,
		ExternalServiceIdentity,
		RoleIdentity,
		PolicyIdentity,
		FlowStatisticIdentity,
		SuggestedPolicyIdentity,
		TabulationIdentity,
		QuotaPolicyIdentity,
		FileAccessPolicyIdentity,
		EnforcerProfileIdentity,
		RenderedPolicyIdentity,
		NamespaceMappingPolicyIdentity,
		ProcessingUnitIdentity,
		DependencyMapIdentity,
		VulnerabilityIdentity,
		EnforcerProfileMappingPolicyIdentity,
		ActivityIdentity,
		RootIdentity,
		NetworkAccessPolicyIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"hook":       HookPolicyIdentity,
	"hooks":      HookPolicyIdentity,
	"hookpol":    HookPolicyIdentity,
	"hookpols":   HookPolicyIdentity,
	"apiauths":   APIAuthorizationPolicyIdentity,
	"apiauth":    APIAuthorizationPolicyIdentity,
	"extac":      ExternalAccessIdentity,
	"extacs":     ExternalAccessIdentity,
	"mess":       MessageIdentity,
	"fp":         FilePathIdentity,
	"fps":        FilePathIdentity,
	"ns":         NamespaceIdentity,
	"extsrv":     ExternalServiceIdentity,
	"extsrvs":    ExternalServiceIdentity,
	"sugpol":     SuggestedPolicyIdentity,
	"sugpols":    SuggestedPolicyIdentity,
	"sugg":       SuggestedPolicyIdentity,
	"suggs":      SuggestedPolicyIdentity,
	"table":      TabulationIdentity,
	"tables":     TabulationIdentity,
	"tabs":       TabulationIdentity,
	"tab":        TabulationIdentity,
	"quota":      QuotaPolicyIdentity,
	"quotas":     QuotaPolicyIdentity,
	"quotapol":   QuotaPolicyIdentity,
	"quotapols":  QuotaPolicyIdentity,
	"profile":    EnforcerProfileIdentity,
	"profiles":   EnforcerProfileIdentity,
	"rpol":       RenderedPolicyIdentity,
	"rpols":      RenderedPolicyIdentity,
	"nspolicy":   NamespaceMappingPolicyIdentity,
	"nspolicies": NamespaceMappingPolicyIdentity,
	"nsmap":      NamespaceMappingPolicyIdentity,
	"nsmaps":     NamespaceMappingPolicyIdentity,
	"pu":         ProcessingUnitIdentity,
	"pus":        ProcessingUnitIdentity,
	"depmap":     DependencyMapIdentity,
	"depmaps":    DependencyMapIdentity,
	"vulns":      VulnerabilityIdentity,
	"vul":        VulnerabilityIdentity,
	"vuln":       VulnerabilityIdentity,
	"vuls":       VulnerabilityIdentity,
	"srvpols":    EnforcerProfileMappingPolicyIdentity,
	"srvpol":     EnforcerProfileMappingPolicyIdentity,
	"netpol":     NetworkAccessPolicyIdentity,
	"netpols":    NetworkAccessPolicyIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case APICheckIdentity:
		return []string{}
	case HookPolicyIdentity:
		return []string{
			"hook",
			"hooks",
			"hookpol",
			"hookpols",
		}
	case APIAuthorizationPolicyIdentity:
		return []string{
			"apiauths",
			"apiauth",
		}
	case ExternalAccessIdentity:
		return []string{
			"extac",
			"extacs",
		}
	case SystemCallIdentity:
		return []string{}
	case TagIdentity:
		return []string{}
	case EnforcerIdentity:
		return []string{}
	case MessageIdentity:
		return []string{
			"mess",
		}
	case PokeIdentity:
		return []string{}
	case FilePathIdentity:
		return []string{
			"fp",
			"fps",
		}
	case FileAccessIdentity:
		return []string{}
	case NamespaceIdentity:
		return []string{
			"ns",
		}
	case PolicyRuleIdentity:
		return []string{}
	case ExternalServiceIdentity:
		return []string{
			"extsrv",
			"extsrvs",
		}
	case RoleIdentity:
		return []string{}
	case PolicyIdentity:
		return []string{}
	case FlowStatisticIdentity:
		return []string{}
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
	case QuotaPolicyIdentity:
		return []string{
			"quota",
			"quotas",
			"quotapol",
			"quotapols",
		}
	case FileAccessPolicyIdentity:
		return []string{}
	case EnforcerProfileIdentity:
		return []string{
			"profile",
			"profiles",
		}
	case RenderedPolicyIdentity:
		return []string{
			"rpol",
			"rpols",
		}
	case NamespaceMappingPolicyIdentity:
		return []string{
			"nspolicy",
			"nspolicies",
			"nsmap",
			"nsmaps",
		}
	case ProcessingUnitIdentity:
		return []string{
			"pu",
			"pus",
		}
	case DependencyMapIdentity:
		return []string{
			"depmap",
			"depmaps",
		}
	case VulnerabilityIdentity:
		return []string{
			"vulns",
			"vul",
			"vuln",
			"vuls",
		}
	case EnforcerProfileMappingPolicyIdentity:
		return []string{
			"srvpols",
			"srvpol",
		}
	case ActivityIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case NetworkAccessPolicyIdentity:
		return []string{
			"netpol",
			"netpols",
		}
	}

	return nil
}
