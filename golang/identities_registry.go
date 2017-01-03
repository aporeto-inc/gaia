package gaia

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(NamespaceMappingPolicyIdentity)
	elemental.RegisterIdentity(DependencyMapSubviewIdentity)
	elemental.RegisterIdentity(APIAuthorizationPolicyIdentity)
	elemental.RegisterIdentity(NamespaceContentIdentity)
	elemental.RegisterIdentity(SystemInfoIdentity)
	elemental.RegisterIdentity(SyscallAccessIdentity)
	elemental.RegisterIdentity(ComputedPolicyIdentity)
	elemental.RegisterIdentity(TagIdentity)
	elemental.RegisterIdentity(MapEdgeIdentity)
	elemental.RegisterIdentity(CertificateIdentity)
	elemental.RegisterIdentity(FilePathIdentity)
	elemental.RegisterIdentity(FileAccessIdentity)
	elemental.RegisterIdentity(NamespaceIdentity)
	elemental.RegisterIdentity(IntegrationIdentity)
	elemental.RegisterIdentity(PolicyRuleIdentity)
	elemental.RegisterIdentity(ExternalServiceIdentity)
	elemental.RegisterIdentity(PolicyIdentity)
	elemental.RegisterIdentity(FlowStatisticIdentity)
	elemental.RegisterIdentity(ServerProfileIdentity)
	elemental.RegisterIdentity(ServerPolicyIdentity)
	elemental.RegisterIdentity(ComputedDependencyMapViewIdentity)
	elemental.RegisterIdentity(SystemCallIdentity)
	elemental.RegisterIdentity(AuthenticatorIdentity)
	elemental.RegisterIdentity(FileAccessPolicyIdentity)
	elemental.RegisterIdentity(RenderedPolicyIdentity)
	elemental.RegisterIdentity(ProcessingUnitIdentity)
	elemental.RegisterIdentity(DependencyMapViewIdentity)
	elemental.RegisterIdentity(DependencyMapIdentity)
	elemental.RegisterIdentity(VulnerabilityIdentity)
	elemental.RegisterIdentity(ServerIdentity)
	elemental.RegisterIdentity(MapNodeIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(NetworkAccessPolicyIdentity)
	elemental.RegisterIdentity(UserIdentity)
}

func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case NamespaceMappingPolicyIdentity.Name:
		return NewNamespaceMappingPolicy()
	case DependencyMapSubviewIdentity.Name:
		return NewDependencyMapSubview()
	case APIAuthorizationPolicyIdentity.Name:
		return NewAPIAuthorizationPolicy()
	case NamespaceContentIdentity.Name:
		return NewNamespaceContent()
	case SystemInfoIdentity.Name:
		return NewSystemInfo()
	case SyscallAccessIdentity.Name:
		return NewSyscallAccess()
	case ComputedPolicyIdentity.Name:
		return NewComputedPolicy()
	case TagIdentity.Name:
		return NewTag()
	case MapEdgeIdentity.Name:
		return NewMapEdge()
	case CertificateIdentity.Name:
		return NewCertificate()
	case FilePathIdentity.Name:
		return NewFilePath()
	case FileAccessIdentity.Name:
		return NewFileAccess()
	case NamespaceIdentity.Name:
		return NewNamespace()
	case IntegrationIdentity.Name:
		return NewIntegration()
	case PolicyRuleIdentity.Name:
		return NewPolicyRule()
	case ExternalServiceIdentity.Name:
		return NewExternalService()
	case PolicyIdentity.Name:
		return NewPolicy()
	case FlowStatisticIdentity.Name:
		return NewFlowStatistic()
	case ServerProfileIdentity.Name:
		return NewServerProfile()
	case ServerPolicyIdentity.Name:
		return NewServerPolicy()
	case ComputedDependencyMapViewIdentity.Name:
		return NewComputedDependencyMapView()
	case SystemCallIdentity.Name:
		return NewSystemCall()
	case AuthenticatorIdentity.Name:
		return NewAuthenticator()
	case FileAccessPolicyIdentity.Name:
		return NewFileAccessPolicy()
	case RenderedPolicyIdentity.Name:
		return NewRenderedPolicy()
	case ProcessingUnitIdentity.Name:
		return NewProcessingUnit()
	case DependencyMapViewIdentity.Name:
		return NewDependencyMapView()
	case DependencyMapIdentity.Name:
		return NewDependencyMap()
	case VulnerabilityIdentity.Name:
		return NewVulnerability()
	case ServerIdentity.Name:
		return NewServer()
	case MapNodeIdentity.Name:
		return NewMapNode()
	case RootIdentity.Name:
		return NewRoot()
	case NetworkAccessPolicyIdentity.Name:
		return NewNetworkAccessPolicy()
	case UserIdentity.Name:
		return NewUser()
	default:
		return nil
	}

	return nil
}
