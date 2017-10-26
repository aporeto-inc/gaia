package barretmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(RevocationIdentity)
	elemental.RegisterIdentity(CertificateIdentity)
	elemental.RegisterIdentity(PrivateKeyIdentity)
	elemental.RegisterIdentity(AuthorityIdentity)
	elemental.RegisterIdentity(TokenIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(CheckIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case RevocationIdentity.Name:
		return NewRevocation()
	case CertificateIdentity.Name:
		return NewCertificate()
	case PrivateKeyIdentity.Name:
		return NewPrivateKey()
	case AuthorityIdentity.Name:
		return NewAuthority()
	case TokenIdentity.Name:
		return NewToken()
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
	case RevocationIdentity.Category:
		return NewRevocation()
	case CertificateIdentity.Category:
		return NewCertificate()
	case PrivateKeyIdentity.Category:
		return NewPrivateKey()
	case AuthorityIdentity.Category:
		return NewAuthority()
	case TokenIdentity.Category:
		return NewToken()
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
	case RevocationIdentity.Name:
		return &RevocationsList{}
	case CertificateIdentity.Name:
		return &CertificatesList{}
	case PrivateKeyIdentity.Name:
		return &PrivateKeysList{}
	case AuthorityIdentity.Name:
		return &AuthoritiesList{}
	case TokenIdentity.Name:
		return &TokensList{}
	case CheckIdentity.Name:
		return &ChecksList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {
	case RevocationIdentity.Category:
		return &RevocationsList{}
	case CertificateIdentity.Category:
		return &CertificatesList{}
	case PrivateKeyIdentity.Category:
		return &PrivateKeysList{}
	case AuthorityIdentity.Category:
		return &AuthoritiesList{}
	case TokenIdentity.Category:
		return &TokensList{}
	case CheckIdentity.Category:
		return &ChecksList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		RevocationIdentity,
		CertificateIdentity,
		PrivateKeyIdentity,
		AuthorityIdentity,
		TokenIdentity,
		RootIdentity,
		CheckIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"api":      CertificateIdentity,
	"apicert":  CertificateIdentity,
	"apis":     CertificateIdentity,
	"apicerts": CertificateIdentity,
	"ca":       AuthorityIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case RevocationIdentity:
		return []string{}
	case CertificateIdentity:
		return []string{
			"api",
			"apicert",
			"apis",
			"apicerts",
		}
	case PrivateKeyIdentity:
		return []string{}
	case AuthorityIdentity:
		return []string{
			"ca",
		}
	case TokenIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case CheckIdentity:
		return []string{}
	}

	return nil
}
