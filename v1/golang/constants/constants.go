package constants

// Vulnerability defines the security vulnerability of an image or a layer
type Vulnerability int

const (
	// VulnerabilityNone defines that security vulnerability is None
	VulnerabilityNone Vulnerability = iota

	// VulnerabilityUnknown defines that security vulnerability is Unknown
	VulnerabilityUnknown

	// VulnerabilityNegligible defines that security vulnerability is Negligible
	VulnerabilityNegligible

	// VulnerabilityLow defines that security vulnerability is Low
	VulnerabilityLow

	// VulnerabilityMedium defines that security vulnerability is Medium
	VulnerabilityMedium

	// VulnerabilityHigh defines that security vulnerability is High
	VulnerabilityHigh

	// VulnerabilityCritical defines that security vulnerability is Critical
	VulnerabilityCritical

	// VulnerabilityDefcon1 defines that security vulnerability is Defcon1
	VulnerabilityDefcon1
)

// RenderedPolicyType defines the type of the RenderedPolicyType
type RenderedPolicyType string

const (
	// RenderedPolicyTypeNetwork that the RenderedPolicyType is network
	RenderedPolicyTypeNetwork RenderedPolicyType = "networkPolicyRules"
	// RenderedPolicyTypeFile that the RenderedPolicyType is file
	RenderedPolicyTypeFile RenderedPolicyType = "filePolicyRules"
	// RenderedPolicyTypeIsolation that the RenderedPolicyType for isolation policies
	RenderedPolicyTypeIsolation RenderedPolicyType = "isolationPolicyRules"
)

const (

	// StatsTagKeyDestinationID represents the destination ID tag.
	StatsTagKeyDestinationID = "@destid"

	// StatsTagKeyDestinationIP represents the destination IP tag.
	StatsTagKeyDestinationIP = "@destip"

	// StatsTagKeyDestinationType represents the destination type tag.
	StatsTagKeyDestinationType = "@desttype"

	// StatsTagKeySourceID represents the source ID tag.
	StatsTagKeySourceID = "@srcid"

	// StatsTagKeySourceIP represents the source IP tag.
	StatsTagKeySourceIP = "@srcip"

	// StatsTagKeyURI represents the URI tag.
	StatsTagKeyURI = "@uri"

	// StatsTagKeyObservationMode represents the that the flow is oberserved.
	StatsTagKeyObservationMode = "@observed"

	// StatsTagKeyObservedAction represents the action taken on an oberserved flow.
	StatsTagKeyObservedAction = "@oaction"

	// StatsTagKeyObservedPolicyID represents the observation policy ID.
	StatsTagKeyObservedPolicyID = "@opolicyid"

	// StatsTagKeySourceType represents the source type tag.
	StatsTagKeySourceType = "@srctype"

	// StatsTagKeyDestinationPort represents the destination port tag.
	StatsTagKeyDestinationPort = "@destport"

	// StatsTagKeyEncrypted represents the encrypted tag.
	StatsTagKeyEncrypted = "@encrypted"

	// StatsTagKeyPolicyID represents the policy ID tag.
	StatsTagKeyPolicyID = "@policyid"

	// StatsTagKeyProcessingUnitID represents the processing unit ID tag.
	StatsTagKeyProcessingUnitID = "$id"

	// StatsTagKeyEnforcerID represents the processing unit ID tag.
	StatsTagKeyEnforcerID = "@eid"

	// StatsTagKeyIP represents the IP tag.
	StatsTagKeyIP = "@ip"

	// StatsTagKeyAction represents the action tag.
	StatsTagKeyAction = "@action"

	// StatsTagKeyReason represents the reason tag.
	StatsTagKeyReason = "@reason"

	// StatsTagKeyEvent represents the event tag.
	StatsTagKeyEvent = "@event"

	// StatsTagKeyImage represents the image tag.
	StatsTagKeyImage = "@sys:image"

	// StatsTagKeyPid represents the parent Pid of the container
	StatsTagKeyPid = "@sys:pid"

	// StatsTagKeyUserID is the user id.
	StatsTagKeyUserID = "@sys:uid"

	// StatsTagKeyServerID represents the server ID tag.
	StatsTagKeyServerID = "@serverid"

	// StatsTagKeyNamespace represents the namespace tag.
	StatsTagKeyNamespace = "$namespace"

	// StatsTagKeyAporetoContextID represents the aporeto context ID.
	StatsTagKeyAporetoContextID = "@squallid"

	// StatsTagKeyFileAccessPath represents the file path.
	StatsTagKeyFileAccessPath = "@path"

	// StatsTagKeyFileAccessHost represents the host where the file is stored.
	StatsTagKeyFileAccessHost = "@host"

	// StatsTagKeyFileAccessProtocol represents the protocol used to access the file.
	StatsTagKeyFileAccessProtocol = "@protocol"

	// StatsTagKeyFileAccessMode represents the mode of file.
	StatsTagKeyFileAccessMode = "@mode"

	// StatsTagKeySyscallAccessName represents the number of the syscall.
	StatsTagKeySyscallAccessName = "@syscall"

	// StatsTagKeySyscallSuccess represents the success or not  of the syscall.
	StatsTagKeySyscallSuccess = "@success"

	// StatsTagKeyExecutable represents the executable of the syscall
	StatsTagKeyExecutable = "@exec"

	// StatsTagKeyEventType represents the key of the event type
	StatsTagKeyEventType = "@type"

	// StatsTagKeySyscallAccessPID represents the pid of the related process.
	StatsTagKeySyscallAccessPID = "@pid"

	// StatsTagKeySyscallAccessProcessName represents the name of the related process.
	StatsTagKeySyscallAccessProcessName = "@processname"

	// StatsTagKeyFlowTypeProcessingUnit represents the pu type value.
	StatsTagKeyFlowTypeProcessingUnit = "pu"

	// StatsTagKeyFlowTypeExternalService represents the ext value.
	StatsTagKeyFlowTypeExternalService = "ext"
)

// Shared Policy constants
const (
	// RelationConnectTagString is the system created relation=connect tag.
	RelationConnectTagString = "relation=connect"

	// RelationStartPUTagString is the system created relation=start tag.
	RelationCreatePUTagString = "relation=start"

	// RelationRunPUTagString is the system create relation=run tag.
	RelationRunPUTagString = "relation=run"

	// ActionAllowTagString is the system created action=allow tag.
	ActionAllowTagString = "action=allow"

	// ActionRejectTagString is the system created "action=reject" tag
	ActionRejectTagString = "action=reject"

	// ActionContinueTagString is the system created "action=continue" tag
	ActionContinueTagString = "action=continue"

	// ActionEncryptTagString is the system created action=encrypt tag.
	ActionEncryptTagString = "action=encrypt"

	// ActionPassthroughTagString is the system created action=passthrough tag.
	ActionPassthroughTagString = "action=passthrough"

	// ActionLogTagString is the system created action=log tag.
	ActionLogTagString = "action=log"

	// ActionStopPUTagString is the system created action-stop tag
	ActionPUDeleteTagString = "action=delete"

	// ActionStopPUTagString is the system created action-stop tag
	ActionPUSnapshotTagString = "action=snapshot"

	// ActionStopPUTagString is the system created action-stop tag
	ActionPURejectTagString = "action=reject"

	// ActionStopPUTagString is the system created action-stop tag
	ActionPULogComplianceTagString = "action=logcomply"

	// ActionStopPUTagString is the system created action-stop tag
	ActionPUEnforceTagString = "action=enforce"

	// ActionStopPUTagString is the system created action-stop tag
	ActionPUStopTagString = "action=stop"

	// ActionObserveTagString is the system created action=observe tag.
	ActionObserveTagString = "action=observe"

	// ActionOApplyTagString is the system created observedaction=allow tag.
	ActionOApplyTagString = "oaction=apply"

	// ActionOContinueTagString is the system created "observedaction=continue" tag
	ActionOContinueTagString = "oaction=continue"

	// ActionTokenScope is the system created action=scope tag.
	ActionTokenScope = "action=scope"
)

const (
	// UserLabelPrefix is the label prefix for all user defined labels
	UserLabelPrefix = "@usr:"

	// SystemLabelPrefix is the label prefix for all system defined labels
	SystemLabelPrefix = "@sys:"

	// AttributeLabelPrefix is the prefix for all attributes
	AttributeLabelPrefix = "$"
)

// File mode tags
const (
	// RelationReadTagString is the system created relation=read tag.
	RelationReadTagString = "relation=read"

	// RelationWriteTagString is the system created relation=write tag.
	RelationWriteTagString = "relation=write"

	// RelationExecuteTagString is the system created relation=execute tag.
	RelationExecuteTagString = "relation=execute"
)

// Quota Policies Constants
const (
	// ActionLimitTagString is the system created action=limit tag.
	ActionLimitTagString = "action=limit"

	// RelationQuotaTagString is the system created quota=n tag.
	ActionLimitQuotaKey = "quota"

	// RelationQuotaNamespaceTagString is the system created quota=n tag.
	ActionLimitTargetNamespaceKey = "targetNamespace"
)

// Dynamic tags constants.
const (

	// DynamicTagOperationalStatusKey is the key of the operationalstatus tag.
	DynamicTagOperationalStatusKey = "$operationalstatus"

	// DynamicTagNameKey is the key for the name tag.
	DynamicTagNameKey = "$name"

	// DynamicTagIDKey is the key of the id tag.
	DynamicTagIDKey = "$id"

	//DynamicTagTypeKey is the key of the type tag.
	DynamicTagTypeKey = "$type"

	// DynamicTagNamespaceKey is the key of the namespace tag.
	DynamicTagNamespaceKey = "$namespace"

	// DynamicTagIdentityKey is the key of the identity tag.
	DynamicTagIdentityKey = "$identity"

	// DynamicTagEnforcerIDKey is the key of the enforcerid tag.
	DynamicTagEnforcerIDKey = "$enforcerid"
)

// Metadata constants
const (
	// MetadataSysImageKey is the key of the image tag.
	MetadataSysImageKey = "@sys:image"

	// MetatataSysNameKey is the key of the name tag.
	MetatataSysNameKey = "@sys:name"

	// MetadataUsrDockerCompose  is the key of the docker compose tag.
	MetadataUsrDockerCompose = "@usr:com.docker.compose"

	// MetadataEvent is the key of the event tag
	MetadataEventKey = "@event"

	// MetadataUsrPortKey is the key of the user port tag
	MetadataUsrPortKey = "@usr:port"
)

// Auth tags constants.
const (

	// AuthKey is the key of a auth.
	AuthKey = "@auth:"

	// AuthRoleKey is the key of a role.
	AuthRoleKey = "@auth:role"

	// AuthOrganizationKey is the key of a organization.
	AuthOrganizationKey = "@auth:organization"
)

// HTTP Method tags
const (
	// HTTPMethodAnyTagString is the system created http-method=any tag.
	HTTPMethodAnyTagString = "http-method=any"

	// HTTPMethodPostTagString is the system created http-method=POST tag.
	HTTPMethodPostTagString = "http-method=post"

	// HTTPMethodPutTagString is the system created http-method=PUT tag.
	HTTPMethodPutTagString = "http-method=put"

	// HTTPMethodGetTagString is the system created http-method=GET tag.
	HTTPMethodGetTagString = "http-method=get"

	// HTTPMethodDeleteTagString is the system created http-method=DELETE tag.
	HTTPMethodDeleteTagString = "http-method=delete"

	// HTTPMethodPatchTagString is the system created http-method=PATCH tag.
	HTTPMethodPatchTagString = "http-method=patch"

	// HTTPMethodHeadTagString is the system created http-method=PATCH tag.
	HTTPMethodHeadTagString = "http-method=head"
)

// Vulnerability Constants
const (
	// VulnerabilityTagKeyLevel is the key of the vulnerability tag.
	VulnerabilityTagKeyLevel = "vulnerability:level"

	// VulnerabilityTagKeyCVE is the key of the cve tag.
	VulnerabilityTagKeyCVE = "vulnerability:cve"
)
