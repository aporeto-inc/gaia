package midgardmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(IssueIdentity)
	elemental.RegisterIdentity(AuthIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case RootIdentity.Name:
		return NewRoot()
	case IssueIdentity.Name:
		return NewIssue()
	case AuthIdentity.Name:
		return NewAuth()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {
	case RootIdentity.Category:
		return NewRoot()
	case IssueIdentity.Category:
		return NewIssue()
	case AuthIdentity.Category:
		return NewAuth()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case IssueIdentity.Name:
		return &IssuesList{}
	case AuthIdentity.Name:
		return &AuthsList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {
	case IssueIdentity.Category:
		return &IssuesList{}
	case AuthIdentity.Category:
		return &AuthsList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		RootIdentity,
		IssueIdentity,
		AuthIdentity,
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
	case IssueIdentity:
		return []string{}
	case AuthIdentity:
		return []string{}
	}

	return nil
}
