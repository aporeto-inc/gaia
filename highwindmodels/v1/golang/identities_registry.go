package highwindmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(CategoryIdentity)
	elemental.RegisterIdentity(InstallationIdentity)
	elemental.RegisterIdentity(LogIdentity)
	elemental.RegisterIdentity(ServiceIdentity)
	elemental.RegisterIdentity(AvailableServiceIdentity)
	elemental.RegisterIdentity(RootIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case CategoryIdentity.Name:
		return NewCategory()
	case InstallationIdentity.Name:
		return NewInstallation()
	case LogIdentity.Name:
		return NewLog()
	case ServiceIdentity.Name:
		return NewService()
	case AvailableServiceIdentity.Name:
		return NewAvailableService()
	case RootIdentity.Name:
		return NewRoot()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {
	case CategoryIdentity.Category:
		return NewCategory()
	case InstallationIdentity.Category:
		return NewInstallation()
	case LogIdentity.Category:
		return NewLog()
	case ServiceIdentity.Category:
		return NewService()
	case AvailableServiceIdentity.Category:
		return NewAvailableService()
	case RootIdentity.Category:
		return NewRoot()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case CategoryIdentity.Name:
		return &CategoriesList{}
	case InstallationIdentity.Name:
		return &InstallationsList{}
	case LogIdentity.Name:
		return &LogsList{}
	case ServiceIdentity.Name:
		return &ServicesList{}
	case AvailableServiceIdentity.Name:
		return &AvailableServicesList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {
	case CategoryIdentity.Category:
		return &CategoriesList{}
	case InstallationIdentity.Category:
		return &InstallationsList{}
	case LogIdentity.Category:
		return &LogsList{}
	case ServiceIdentity.Category:
		return &ServicesList{}
	case AvailableServiceIdentity.Category:
		return &AvailableServicesList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		CategoryIdentity,
		InstallationIdentity,
		LogIdentity,
		ServiceIdentity,
		AvailableServiceIdentity,
		RootIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"srv":  ServiceIdentity,
	"asrv": AvailableServiceIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case CategoryIdentity:
		return []string{}
	case InstallationIdentity:
		return []string{}
	case LogIdentity:
		return []string{}
	case ServiceIdentity:
		return []string{
			"srv",
		}
	case AvailableServiceIdentity:
		return []string{
			"asrv",
		}
	case RootIdentity:
		return []string{}
	}

	return nil
}
