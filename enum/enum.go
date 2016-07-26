package enum

// EntityStatus defines the status of an entity
type EntityStatus string

const (
	// Active defines that an entity is active
	Active EntityStatus = "Active"

	// Disabled defines that an entity is disabled
	Disabled EntityStatus = "Disabled"

	// Candidate defines that an entity is a candidate and could be either activated or deleted or disabled
	Candidate EntityStatus = "Candidate"
)

// PolicyType defines the type the policy
type PolicyType string

const (
	// APIAuthorization defines that a policy is a APIAuthorization policy
	APIAuthorization PolicyType = "APIAuthorization"

	// Network defines that a policy is a Network policy
	Network PolicyType = "Network"

	// File defines that a policy is a File policy
	File PolicyType = "File"

	// Syscall defines that a policy is a Syscall policy
	Syscall PolicyType = "Syscall"

	// Statistics defines that a policy is a Statistics policy
	Statistics PolicyType = "Statistics"

	// ExtendTags defines that a policy is a ExtendTags policy
	ExtendTags PolicyType = "ExtendTags"

	// NamespaceMapping defines that a policy is a NamespaceMapping policy
	NamespaceMapping PolicyType = "NamespaceMapping"
)

// EntityStatusToString converts EntityStatusToString to its corresponding string value
func EntityStatusToString(entityStatus EntityStatus) string {

	switch entityStatus {
	case Active:
		return "Active"
	case Disabled:
		return "Disabled"
	case Candidate:
		return "Candidate"
	default:
		return ""
	}
}

// PolicyTypeToString converts PolicyTypeToString to its corresponding string value
func PolicyTypeToString(policyType PolicyType) string {

	switch policyType {
	case APIAuthorization:
		return "APIAuthorization"
	case Network:
		return "Network"
	case File:
		return "File"
	case Syscall:
		return "Syscall"
	case Statistics:
		return "Statistics"
	case ExtendTags:
		return "ExtendTags"
	case NamespaceMapping:
		return "NamespaceMapping"
	default:
		return ""
	}
}

// DependencyMapViewSelector is the type used for the selectors of a DependencyMapSubView.
type DependencyMapViewSelector []string

// DependencyMapViewSubSelector is the type used for the subselector of a DependencyMapSubView.
type DependencyMapViewSubSelector map[string]DependencyMapViewSelector
