package constants

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

// Vulnerability defines the security vulnerability of an image or a layer
type Vulnerability int

const (
	// None defines that security vulnerability is None
	None Vulnerability = iota
	// Unknown defines that security vulnerability is Unknown
	Unknown
	// Negligible defines that security vulnerability is Negligible
	Negligible
	// Low defines that security vulnerability is Low
	Low
	// Medium defines that security vulnerability is Medium
	Medium
	// High defines that security vulnerability is High
	High
	// Critical defines that security vulnerability is Critical
	Critical
	// Defcon1 defines that security vulnerability is Defcon1
	Defcon1
)

var vulnerabilityValues = [][]string{
	{"none", "None"},
	{"unknown", "Unknown"},
	{"low", "Low"},
	{"negligible", "Negligible"},
	{"medium", "Medium"},
	{"critical", "Critical"},
	{"high", "High"},
	{"defcon1", "Defcon1"},
}

var vulnerabilityMap = map[string]Vulnerability{
	"None":       None,
	"Unknown":    Unknown,
	"Low":        Low,
	"Negligible": Negligible,
	"Medium":     Medium,
	"Critical":   Critical,
	"High":       High,
	"Defcon1":    Defcon1,
}

// TagValue returns the value of vulnerability system tag
func (sl Vulnerability) TagValue() string { return vulnerabilityValues[sl][0] }

func (sl Vulnerability) String() string { return vulnerabilityValues[sl][1] }

// VulnerabilityConst return constant value of vulnerability string
func VulnerabilityConst(vulnerability string) Vulnerability { return vulnerabilityMap[vulnerability] }

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

// IntegrationType represents the possible type of integration servers
type IntegrationType string

const (
	// Registry represents the value Registry
	Registry IntegrationType = "Registry"

	// Vulnerabilityscanner represents the value VulnerabilityScanner
	Vulnerabilityscanner IntegrationType = "VulnerabilityScanner"
)

const (

	// StatsTagKeyDestinationID represents the destination ID tag.
	StatsTagKeyDestinationID = "destID"

	// StatsTagKeySourceID represents the source ID tag.
	StatsTagKeySourceID = "srcID"

	// StatsTagKeyDestinationPort represents the destination port tag.
	StatsTagKeyDestinationPort = "destPort"

	// StatsTagKeyFlowContext represents the flow context tag.
	StatsTagKeyFlowContext = "flowContext"

	// StatsTagKeyProcessingUnitID represents the processing unit ID tag.
	StatsTagKeyProcessingUnitID = "ID"

	// StatsTagKeyIP represents the IP tag.
	StatsTagKeyIP = "IP"

	// StatsTagKeyAction represents the action tag.
	StatsTagKeyAction = "action"

	// StatsTagKeyReason represents the reason tag.
	StatsTagKeyReason = "reason"

	// StatsTagKeyEvent represents the event tag.
	StatsTagKeyEvent = "event"

	// StatsTagKeyImage represents the image tag.
	StatsTagKeyImage = "image"

	// StatsTagKeyServerID represents the server ID tag.
	StatsTagKeyServerID = "serverID"

	// StatsTagKeyNamespace represents the namespace tag.
	StatsTagKeyNamespace = "AporetoNamespace"

	// StatsTagKeyAporetoContextID represents the aporeto context ID.
	StatsTagKeyAporetoContextID = "AporetoContextID"

	// StatsTagKeyFileAccessPath represents the file path.
	StatsTagKeyFileAccessPath = "path"

	// StatsTagKeyFileAccessHost represents the host where the file is stored.
	StatsTagKeyFileAccessHost = "host"

	// StatsTagKeyFileAccessProtocol represents the protocol used to access the file.
	StatsTagKeyFileAccessProtocol = "protocol"

	// StatsTagKeyFileAccessMode represents the mode of file.
	StatsTagKeyFileAccessMode = "mode"

	// StatsTagKeyFileAccessContext represents the full context of the file access.
	StatsTagKeyFileAccessContext = "fileContext"
)
