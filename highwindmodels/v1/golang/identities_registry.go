package highwindmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(InstallationIdentity)
	elemental.RegisterIdentity(ServiceIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case RootIdentity.Name:
		return NewRoot()
	case InstallationIdentity.Name:
		return NewInstallation()
	case ServiceIdentity.Name:
		return NewService()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {
	case RootIdentity.Category:
		return NewRoot()
	case InstallationIdentity.Category:
		return NewInstallation()
	case ServiceIdentity.Category:
		return NewService()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case InstallationIdentity.Name:
		return &InstallationsList{}
	case ServiceIdentity.Name:
		return &ServicesList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {
	case InstallationIdentity.Category:
		return &InstallationsList{}
	case ServiceIdentity.Category:
		return &ServicesList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		RootIdentity,
		InstallationIdentity,
		ServiceIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case RootIdentity:
		return []string{}
	case InstallationIdentity:
		return []string{}
	case ServiceIdentity:
		return []string{}
	}

	return nil
}
