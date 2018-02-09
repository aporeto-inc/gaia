package barretmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(AuthorityIdentity)
	elemental.RegisterIdentity(CertificateIdentity)
	elemental.RegisterIdentity(CheckIdentity)
	elemental.RegisterIdentity(PrivateKeyIdentity)
	elemental.RegisterIdentity(RevocationIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(TokenIdentity)
}

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {

	case AuthorityIdentity.Name:
		return NewAuthority()
	case CertificateIdentity.Name:
		return NewCertificate()
	case CheckIdentity.Name:
		return NewCheck()
	case PrivateKeyIdentity.Name:
		return NewPrivateKey()
	case RevocationIdentity.Name:
		return NewRevocation()
	case RootIdentity.Name:
		return NewRoot()
	case TokenIdentity.Name:
		return NewToken()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {

	case AuthorityIdentity.Category:
		return NewAuthority()
	case CertificateIdentity.Category:
		return NewCertificate()
	case CheckIdentity.Category:
		return NewCheck()
	case PrivateKeyIdentity.Category:
		return NewPrivateKey()
	case RevocationIdentity.Category:
		return NewRevocation()
	case RootIdentity.Category:
		return NewRoot()
	case TokenIdentity.Category:
		return NewToken()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {

	case AuthorityIdentity.Name:
		return &AuthoritiesList{}
	case CertificateIdentity.Name:
		return &CertificatesList{}
	case CheckIdentity.Name:
		return &ChecksList{}
	case PrivateKeyIdentity.Name:
		return &PrivateKeysList{}
	case RevocationIdentity.Name:
		return &RevocationsList{}

	case TokenIdentity.Name:
		return &TokensList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {

	case AuthorityIdentity.Category:
		return &AuthoritiesList{}
	case CertificateIdentity.Category:
		return &CertificatesList{}
	case CheckIdentity.Category:
		return &ChecksList{}
	case PrivateKeyIdentity.Category:
		return &PrivateKeysList{}
	case RevocationIdentity.Category:
		return &RevocationsList{}

	case TokenIdentity.Category:
		return &TokensList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		AuthorityIdentity,
		CertificateIdentity,
		CheckIdentity,
		PrivateKeyIdentity,
		RevocationIdentity,
		RootIdentity,
		TokenIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"ca":       AuthorityIdentity,
	"api":      CertificateIdentity,
	"apicert":  CertificateIdentity,
	"apis":     CertificateIdentity,
	"apicerts": CertificateIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case AuthorityIdentity:
		return []string{
			"ca",
		}
	case CertificateIdentity:
		return []string{
			"api",
			"apicert",
			"apis",
			"apicerts",
		}
	case CheckIdentity:
		return []string{}
	case PrivateKeyIdentity:
		return []string{}
	case RevocationIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case TokenIdentity:
		return []string{}
	}

	return nil
}
