package rufusmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(RemoteProcessorIdentity)
	elemental.RegisterIdentity(RootIdentity)
}

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {

	case RemoteProcessorIdentity.Name:
		return NewRemoteProcessor()
	case RootIdentity.Name:
		return NewRoot()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {

	case RemoteProcessorIdentity.Category:
		return NewRemoteProcessor()
	case RootIdentity.Category:
		return NewRoot()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {

	case RemoteProcessorIdentity.Name:
		return &RemoteProcessorsList{}

	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {

	case RemoteProcessorIdentity.Category:
		return &RemoteProcessorsList{}

	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		RemoteProcessorIdentity,
		RootIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"hks": RemoteProcessorIdentity,
	"hk":  RemoteProcessorIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case RemoteProcessorIdentity:
		return []string{
			"hks",
			"hk",
		}
	case RootIdentity:
		return []string{}
	}

	return nil
}
