package vincemodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(AccountIdentity)
	elemental.RegisterIdentity(ActivateIdentity)
	elemental.RegisterIdentity(CertificateIdentity)
	elemental.RegisterIdentity(PasswordResetIdentity)
	elemental.RegisterIdentity(PlanIdentity)
	elemental.RegisterIdentity(AWSAccountIdentity)
	elemental.RegisterIdentity(KubernetesClusterIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(CheckIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1.0 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case AccountIdentity.Name:
		return NewAccount()
	case ActivateIdentity.Name:
		return NewActivate()
	case CertificateIdentity.Name:
		return NewCertificate()
	case PasswordResetIdentity.Name:
		return NewPasswordReset()
	case PlanIdentity.Name:
		return NewPlan()
	case AWSAccountIdentity.Name:
		return NewAWSAccount()
	case KubernetesClusterIdentity.Name:
		return NewKubernetesCluster()
	case RootIdentity.Name:
		return NewRoot()
	case CheckIdentity.Name:
		return NewCheck()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {
	case AccountIdentity.Category:
		return NewAccount()
	case ActivateIdentity.Category:
		return NewActivate()
	case CertificateIdentity.Category:
		return NewCertificate()
	case PasswordResetIdentity.Category:
		return NewPasswordReset()
	case PlanIdentity.Category:
		return NewPlan()
	case AWSAccountIdentity.Category:
		return NewAWSAccount()
	case KubernetesClusterIdentity.Category:
		return NewKubernetesCluster()
	case RootIdentity.Category:
		return NewRoot()
	case CheckIdentity.Category:
		return NewCheck()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case AccountIdentity.Name:
		return &AccountsList{}
	case ActivateIdentity.Name:
		return &ActivatesList{}
	case CertificateIdentity.Name:
		return &CertificatesList{}
	case PasswordResetIdentity.Name:
		return &PasswordResetsList{}
	case PlanIdentity.Name:
		return &PlansList{}
	case AWSAccountIdentity.Name:
		return &AWSAccountsList{}
	case KubernetesClusterIdentity.Name:
		return &KubernetesClustersList{}
	case CheckIdentity.Name:
		return &ChecksList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {
	case AccountIdentity.Category:
		return &AccountsList{}
	case ActivateIdentity.Category:
		return &ActivatesList{}
	case CertificateIdentity.Category:
		return &CertificatesList{}
	case PasswordResetIdentity.Category:
		return &PasswordResetsList{}
	case PlanIdentity.Category:
		return &PlansList{}
	case AWSAccountIdentity.Category:
		return &AWSAccountsList{}
	case KubernetesClusterIdentity.Category:
		return &KubernetesClustersList{}
	case CheckIdentity.Category:
		return &ChecksList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		AccountIdentity,
		ActivateIdentity,
		CertificateIdentity,
		PasswordResetIdentity,
		PlanIdentity,
		AWSAccountIdentity,
		KubernetesClusterIdentity,
		RootIdentity,
		CheckIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"awsacc":  AWSAccountIdentity,
	"awsaccs": AWSAccountIdentity,
	"aws":     AWSAccountIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case AccountIdentity:
		return []string{}
	case ActivateIdentity:
		return []string{}
	case CertificateIdentity:
		return []string{}
	case PasswordResetIdentity:
		return []string{}
	case PlanIdentity:
		return []string{}
	case AWSAccountIdentity:
		return []string{
			"awsacc",
			"awsaccs",
			"aws",
		}
	case KubernetesClusterIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case CheckIdentity:
		return []string{}
	}

	return nil
}
