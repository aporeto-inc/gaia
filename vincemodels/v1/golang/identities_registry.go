package vincemodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(AccountIdentity)
	elemental.RegisterIdentity(ActivateIdentity)
	elemental.RegisterIdentity(AWSAccountIdentity)
	elemental.RegisterIdentity(CertificateIdentity)
	elemental.RegisterIdentity(CheckIdentity)
	elemental.RegisterIdentity(KubernetesClusterIdentity)
	elemental.RegisterIdentity(PasswordResetIdentity)
	elemental.RegisterIdentity(PlanIdentity)
	elemental.RegisterIdentity(RootIdentity)
}

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {

	case AccountIdentity.Name:
		return NewAccount()
	case ActivateIdentity.Name:
		return NewActivate()
	case AWSAccountIdentity.Name:
		return NewAWSAccount()
	case CertificateIdentity.Name:
		return NewCertificate()
	case CheckIdentity.Name:
		return NewCheck()
	case KubernetesClusterIdentity.Name:
		return NewKubernetesCluster()
	case PasswordResetIdentity.Name:
		return NewPasswordReset()
	case PlanIdentity.Name:
		return NewPlan()
	case RootIdentity.Name:
		return NewRoot()
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
	case AWSAccountIdentity.Category:
		return NewAWSAccount()
	case CertificateIdentity.Category:
		return NewCertificate()
	case CheckIdentity.Category:
		return NewCheck()
	case KubernetesClusterIdentity.Category:
		return NewKubernetesCluster()
	case PasswordResetIdentity.Category:
		return NewPasswordReset()
	case PlanIdentity.Category:
		return NewPlan()
	case RootIdentity.Category:
		return NewRoot()
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
	case AWSAccountIdentity.Name:
		return &AWSAccountsList{}
	case CertificateIdentity.Name:
		return &CertificatesList{}
	case CheckIdentity.Name:
		return &ChecksList{}
	case KubernetesClusterIdentity.Name:
		return &KubernetesClustersList{}
	case PasswordResetIdentity.Name:
		return &PasswordResetsList{}
	case PlanIdentity.Name:
		return &PlansList{}

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
	case AWSAccountIdentity.Category:
		return &AWSAccountsList{}
	case CertificateIdentity.Category:
		return &CertificatesList{}
	case CheckIdentity.Category:
		return &ChecksList{}
	case KubernetesClusterIdentity.Category:
		return &KubernetesClustersList{}
	case PasswordResetIdentity.Category:
		return &PasswordResetsList{}
	case PlanIdentity.Category:
		return &PlansList{}

	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		AccountIdentity,
		ActivateIdentity,
		AWSAccountIdentity,
		CertificateIdentity,
		CheckIdentity,
		KubernetesClusterIdentity,
		PasswordResetIdentity,
		PlanIdentity,
		RootIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"aws":     AWSAccountIdentity,
	"awsaccs": AWSAccountIdentity,
	"awsacc":  AWSAccountIdentity,
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
	case AWSAccountIdentity:
		return []string{
			"aws",
			"awsaccs",
			"awsacc",
		}
	case CertificateIdentity:
		return []string{}
	case CheckIdentity:
		return []string{}
	case KubernetesClusterIdentity:
		return []string{}
	case PasswordResetIdentity:
		return []string{}
	case PlanIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	}

	return nil
}
