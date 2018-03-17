package gaia

import (
	"fmt"
	"sync"

	"time"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/types"
)

// EnforcerProfileDockerSocketTypeValue represents the possible values for attribute "dockerSocketType".
type EnforcerProfileDockerSocketTypeValue string

const (
	// EnforcerProfileDockerSocketTypeTcp represents the value tcp.
	EnforcerProfileDockerSocketTypeTcp EnforcerProfileDockerSocketTypeValue = "tcp"

	// EnforcerProfileDockerSocketTypeUnix represents the value unix.
	EnforcerProfileDockerSocketTypeUnix EnforcerProfileDockerSocketTypeValue = "unix"
)

// EnforcerProfileMetadataExtractorValue represents the possible values for attribute "metadataExtractor".
type EnforcerProfileMetadataExtractorValue string

const (
	// EnforcerProfileMetadataExtractorDocker represents the value Docker.
	EnforcerProfileMetadataExtractorDocker EnforcerProfileMetadataExtractorValue = "Docker"

	// EnforcerProfileMetadataExtractorEcs represents the value ECS.
	EnforcerProfileMetadataExtractorEcs EnforcerProfileMetadataExtractorValue = "ECS"

	// EnforcerProfileMetadataExtractorKubernetes represents the value Kubernetes.
	EnforcerProfileMetadataExtractorKubernetes EnforcerProfileMetadataExtractorValue = "Kubernetes"
)

// EnforcerProfileIdentity represents the Identity of the object.
var EnforcerProfileIdentity = elemental.Identity{
	Name:     "enforcerprofile",
	Category: "enforcerprofiles",
	Private:  false,
}

// EnforcerProfilesList represents a list of EnforcerProfiles
type EnforcerProfilesList []*EnforcerProfile

// ContentIdentity returns the identity of the objects in the list.
func (o EnforcerProfilesList) ContentIdentity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Copy returns a pointer to a copy the EnforcerProfilesList.
func (o EnforcerProfilesList) Copy() elemental.ContentIdentifiable {

	copy := append(EnforcerProfilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerProfilesList.
func (o EnforcerProfilesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(EnforcerProfilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerProfile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerProfilesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerProfilesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o EnforcerProfilesList) Version() int {

	return 1
}

// EnforcerProfile represents the model of a enforcerprofile
type EnforcerProfile struct {
	// IptablesMarkValue is the mark value to be used in an iptables implementation.
	IPTablesMarkValue int `json:"IPTablesMarkValue" bson:"iptablesmarkvalue" mapstructure:"IPTablesMarkValue,omitempty"`

	// PUBookkeepingInterval configures how often the PU will be synchronized.
	PUBookkeepingInterval string `json:"PUBookkeepingInterval" bson:"pubookkeepinginterval" mapstructure:"PUBookkeepingInterval,omitempty"`

	// PUHeartbeatInterval configures the heart beat interval.
	PUHeartbeatInterval string `json:"PUHeartbeatInterval" bson:"puheartbeatinterval" mapstructure:"PUHeartbeatInterval,omitempty"`

	// AuditProfileSelectors is the list of tags (key/value pairs) that define the
	// audit policies that must be implemented by this enforcer. The enforcer will
	// implement all policies that match any of these tags.
	AuditProfileSelectors []string `json:"auditProfileSelectors" bson:"auditprofileselectors" mapstructure:"auditProfileSelectors,omitempty"`

	// AuditProfiles returns the audit rules associated with the enforcer profile. This
	// is a read only attribute when an enforcer profile is resolved for an enforcer.
	AuditProfiles AuditProfilesList `json:"auditProfiles" bson:"-" mapstructure:"auditProfiles,omitempty"`

	// AuditSocketBufferSize is the size of the audit socket buffer. Default 16384.
	AuditSocketBufferSize int `json:"auditSocketBufferSize" bson:"auditsocketbuffersize" mapstructure:"auditSocketBufferSize,omitempty"`

	// DockerSocketAddress is the address of the docker daemon.
	DockerSocketAddress string `json:"dockerSocketAddress" bson:"dockersocketaddress" mapstructure:"dockerSocketAddress,omitempty"`

	// DockerSocketType is the type of socket to use to talk to the docker daemon.
	DockerSocketType EnforcerProfileDockerSocketTypeValue `json:"dockerSocketType" bson:"dockersockettype" mapstructure:"dockerSocketType,omitempty"`

	// ExcludedInterfaces is a list of interfaces that must be excluded.
	ExcludedInterfaces []string `json:"excludedInterfaces" bson:"excludedinterfaces" mapstructure:"excludedInterfaces,omitempty"`

	// ExcludedNetworks is the list of networks that must be excluded for this
	// enforcer.
	ExcludedNetworks []string `json:"excludedNetworks" bson:"excludednetworks" mapstructure:"excludedNetworks,omitempty"`

	// HostServices is a list of services that must be activated by default to all
	// enforcers matching this profile.
	HostServices types.HostServicesList `json:"hostServices" bson:"hostservices" mapstructure:"hostServices,omitempty"`

	// IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
	// docker container started with labels matching the rule.
	IgnoreExpression [][]string `json:"ignoreExpression" bson:"ignoreexpression" mapstructure:"ignoreExpression,omitempty"`

	// KubernetesSupportEnabled enables kubernetes mode for the enforcer.
	KubernetesSupportEnabled bool `json:"kubernetesSupportEnabled" bson:"kubernetessupportenabled" mapstructure:"kubernetesSupportEnabled,omitempty"`

	// LinuxProcessesSupportEnabled configures support for Linux processes.
	LinuxProcessesSupportEnabled bool `json:"linuxProcessesSupportEnabled" bson:"linuxprocessessupportenabled" mapstructure:"linuxProcessesSupportEnabled,omitempty"`

	// Select which metadata extractor to use to process new processing units.
	MetadataExtractor EnforcerProfileMetadataExtractorValue `json:"metadataExtractor" bson:"metadataextractor" mapstructure:"metadataExtractor,omitempty"`

	// PolicySynchronizationInterval configures how often the policy will be
	// resynchronized.
	PolicySynchronizationInterval string `json:"policySynchronizationInterval" bson:"policysynchronizationinterval" mapstructure:"policySynchronizationInterval,omitempty"`

	// ProxyListenAddress is the address the enforcer should use to listen for API
	// calls. It can be a port (example :9443) or socket path
	// (example: unix:/var/run/aporeto.sock)
	ProxyListenAddress string `json:"proxyListenAddress" bson:"proxylistenaddress" mapstructure:"proxyListenAddress,omitempty"`

	// ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network
	// receiver starting at the ReceiverQueue
	ReceiverNumberOfQueues int `json:"receiverNumberOfQueues" bson:"receivernumberofqueues" mapstructure:"receiverNumberOfQueues,omitempty"`

	// ReceiverQueue is the base queue number for traffic from the network.
	ReceiverQueue int `json:"receiverQueue" bson:"receiverqueue" mapstructure:"receiverQueue,omitempty"`

	// ReceiverQueueSize is the queue size of the receiver
	ReceiverQueueSize int `json:"receiverQueueSize" bson:"receiverqueuesize" mapstructure:"receiverQueueSize,omitempty"`

	// RemoteEnforcerEnabled inidicates whether a single enforcer should be used or a
	// distributed enforcer. True means distributed.
	RemoteEnforcerEnabled bool `json:"remoteEnforcerEnabled" bson:"remoteenforcerenabled" mapstructure:"remoteEnforcerEnabled,omitempty"`

	// TargetNetworks is the list of networks that authorization should be applied.
	TargetNetworks []string `json:"targetNetworks" bson:"targetnetworks" mapstructure:"targetNetworks,omitempty"`

	// TransmitterNumberOfQueues is the number of queues for application traffic.
	TransmitterNumberOfQueues int `json:"transmitterNumberOfQueues" bson:"transmitternumberofqueues" mapstructure:"transmitterNumberOfQueues,omitempty"`

	// TransmitterQueue is the queue number for traffic from the applications starting
	// at the transmitterQueue
	TransmitterQueue int `json:"transmitterQueue" bson:"transmitterqueue" mapstructure:"transmitterQueue,omitempty"`

	// TransmitterQueueSize is the size of the queue for application traffic.
	TransmitterQueueSize int `json:"transmitterQueueSize" bson:"transmitterqueuesize" mapstructure:"transmitterQueueSize,omitempty"`

	// List of trusted CA. If empty the main chain of trust will be used.
	TrustedCAs []string `json:"trustedCAs" bson:"trustedcas" mapstructure:"trustedCAs,omitempty"`

	// Annotation stores additional information about an entity
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewEnforcerProfile returns a new *EnforcerProfile
func NewEnforcerProfile() *EnforcerProfile {

	return &EnforcerProfile{
		ModelVersion:                  1,
		Annotations:                   map[string][]string{},
		AssociatedTags:                []string{},
		AuditSocketBufferSize:         16384,
		DockerSocketAddress:           "/var/run/docker.sock",
		DockerSocketType:              "unix",
		HostServices:                  types.HostServicesList{},
		IPTablesMarkValue:             1000,
		KubernetesSupportEnabled:      false,
		LinuxProcessesSupportEnabled:  true,
		MetadataExtractor:             "Docker",
		NormalizedTags:                []string{},
		PUBookkeepingInterval:         "15m",
		PUHeartbeatInterval:           "5s",
		PolicySynchronizationInterval: "10m",
		ProxyListenAddress:            "unix:/var/run/aporeto.sock",
		ReceiverNumberOfQueues:        4,
		ReceiverQueue:                 0,
		ReceiverQueueSize:             500,
		RemoteEnforcerEnabled:         true,
		TransmitterNumberOfQueues:     4,
		TransmitterQueue:              4,
		TransmitterQueueSize:          500,
		TrustedCAs:                    []string{},
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerProfile) Identity() elemental.Identity {

	return EnforcerProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerProfile) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *EnforcerProfile) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerProfile) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *EnforcerProfile) Doc() string {
	return `Allows to create reusable configuration profile for your enforcers. Enforcer
Profiles contains various startup information that can (for some) be updated
live. Enforcer Profiles are assigned to some Enforcer using a Enforcer Profile
Mapping Policy.`
}

func (o *EnforcerProfile) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *EnforcerProfile) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *EnforcerProfile) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *EnforcerProfile) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *EnforcerProfile) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *EnforcerProfile) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *EnforcerProfile) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *EnforcerProfile) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *EnforcerProfile) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *EnforcerProfile) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *EnforcerProfile) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *EnforcerProfile) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *EnforcerProfile) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *EnforcerProfile) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetName returns the Name of the receiver.
func (o *EnforcerProfile) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *EnforcerProfile) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *EnforcerProfile) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumInt("IPTablesMarkValue", o.IPTablesMarkValue, int(65000), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("PUBookkeepingInterval", o.PUBookkeepingInterval, `^[0-9]+[smh]$`, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("PUHeartbeatInterval", o.PUHeartbeatInterval, `^[0-9]+[smh]$`, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("auditSocketBufferSize", o.AuditSocketBufferSize, int(262144), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("dockerSocketType", string(o.DockerSocketType), []string{"tcp", "unix"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("metadataExtractor", string(o.MetadataExtractor), []string{"Docker", "ECS", "Kubernetes"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("policySynchronizationInterval", o.PolicySynchronizationInterval, `^[0-9]+[smh]$`, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("proxyListenAddress", o.ProxyListenAddress, `^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix:(/[^/]{1,16}){1,5}/?)$`, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverNumberOfQueues", o.ReceiverNumberOfQueues, int(16), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("receiverNumberOfQueues", o.ReceiverNumberOfQueues, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverQueue", o.ReceiverQueue, int(1000), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverQueueSize", o.ReceiverQueueSize, int(5000), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("receiverQueueSize", o.ReceiverQueueSize, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterNumberOfQueues", o.TransmitterNumberOfQueues, int(16), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterNumberOfQueues", o.TransmitterNumberOfQueues, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterQueue", o.TransmitterQueue, int(1000), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterQueue", o.TransmitterQueue, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterQueueSize", o.TransmitterQueueSize, int(1000), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterQueueSize", o.TransmitterQueueSize, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*EnforcerProfile) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerProfileAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerProfileLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerProfile) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerProfileAttributesMap
}

// EnforcerProfileAttributesMap represents the map of attribute for EnforcerProfile.
var EnforcerProfileAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"IPTablesMarkValue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPTablesMarkValue",
		DefaultValue:   1000,
		Description:    `IptablesMarkValue is the mark value to be used in an iptables implementation.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       65000,
		Name:           "IPTablesMarkValue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"PUBookkeepingInterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PUBookkeepingInterval",
		DefaultValue:   "15m",
		Description:    `PUBookkeepingInterval configures how often the PU will be synchronized.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "PUBookkeepingInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"PUHeartbeatInterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PUHeartbeatInterval",
		DefaultValue:   "5s",
		Description:    `PUHeartbeatInterval configures the heart beat interval.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "PUHeartbeatInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"AuditProfileSelectors": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditProfileSelectors",
		Description: `AuditProfileSelectors is the list of tags (key/value pairs) that define the
audit policies that must be implemented by this enforcer. The enforcer will
implement all policies that match any of these tags.`,
		Exposed:    true,
		Filterable: true,
		Name:       "auditProfileSelectors",
		Stored:     true,
		SubType:    "audit_profile_selector",
		Type:       "external",
	},
	"AuditProfiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AuditProfiles",
		Description: `AuditProfiles returns the audit rules associated with the enforcer profile. This
is a read only attribute when an enforcer profile is resolved for an enforcer.`,
		Exposed:  true,
		Name:     "auditProfiles",
		ReadOnly: true,
		SubType:  "audit_profiles",
		Type:     "external",
	},
	"AuditSocketBufferSize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditSocketBufferSize",
		DefaultValue:   16384,
		Description:    `AuditSocketBufferSize is the size of the audit socket buffer. Default 16384.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       262144,
		Name:           "auditSocketBufferSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"DockerSocketAddress": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DockerSocketAddress",
		DefaultValue:   "/var/run/docker.sock",
		Description:    `DockerSocketAddress is the address of the docker daemon.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "dockerSocketAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"DockerSocketType": elemental.AttributeSpecification{
		AllowedChoices: []string{"tcp", "unix"},
		ConvertedName:  "DockerSocketType",
		DefaultValue:   EnforcerProfileDockerSocketTypeUnix,
		Description:    `DockerSocketType is the type of socket to use to talk to the docker daemon.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "dockerSocketType",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"ExcludedInterfaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedInterfaces",
		Description:    `ExcludedInterfaces is a list of interfaces that must be excluded.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "excludedInterfaces",
		Orderable:      true,
		Stored:         true,
		SubType:        "excluded_interfaces_list",
		Type:           "external",
	},
	"ExcludedNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedNetworks",
		Description: `ExcludedNetworks is the list of networks that must be excluded for this
enforcer.`,
		Exposed:    true,
		Filterable: true,
		Name:       "excludedNetworks",
		Orderable:  true,
		Stored:     true,
		SubType:    "excluded_networks_list",
		Type:       "external",
	},
	"HostServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HostServices",
		Description: `HostServices is a list of services that must be activated by default to all
enforcers matching this profile.`,
		Exposed:    true,
		Filterable: true,
		Name:       "hostServices",
		Orderable:  true,
		Stored:     true,
		SubType:    "host_services_list",
		Type:       "external",
	},
	"IgnoreExpression": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IgnoreExpression",
		Description: `IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
docker container started with labels matching the rule.`,
		Exposed: true,
		Name:    "ignoreExpression",
		Stored:  true,
		SubType: "policies_list",
		Type:    "external",
	},
	"KubernetesSupportEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesSupportEnabled",
		DefaultValue:   false,
		Description:    `KubernetesSupportEnabled enables kubernetes mode for the enforcer.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "kubernetesSupportEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"LinuxProcessesSupportEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LinuxProcessesSupportEnabled",
		DefaultValue:   true,
		Description:    `LinuxProcessesSupportEnabled configures support for Linux processes.`,
		Exposed:        true,
		Name:           "linuxProcessesSupportEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"MetadataExtractor": elemental.AttributeSpecification{
		AllowedChoices: []string{"Docker", "ECS", "Kubernetes"},
		ConvertedName:  "MetadataExtractor",
		DefaultValue:   EnforcerProfileMetadataExtractorDocker,
		Description:    `Select which metadata extractor to use to process new processing units.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "metadataExtractor",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Index:          true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"PolicySynchronizationInterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PolicySynchronizationInterval",
		DefaultValue:   "10m",
		Description: `PolicySynchronizationInterval configures how often the policy will be
resynchronized.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "policySynchronizationInterval",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"ProxyListenAddress": elemental.AttributeSpecification{
		AllowedChars:   `^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix:(/[^/]{1,16}){1,5}/?)$`,
		AllowedChoices: []string{},
		ConvertedName:  "ProxyListenAddress",
		DefaultValue:   "unix:/var/run/aporeto.sock",
		Description: `ProxyListenAddress is the address the enforcer should use to listen for API
calls. It can be a port (example :9443) or socket path
(example: unix:/var/run/aporeto.sock)`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "proxyListenAddress",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"ReceiverNumberOfQueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverNumberOfQueues",
		DefaultValue:   4,
		Description: `ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network
receiver starting at the ReceiverQueue`,
		Exposed:    true,
		Filterable: true,
		MaxValue:   16,
		MinValue:   1,
		Name:       "receiverNumberOfQueues",
		Orderable:  true,
		Stored:     true,
		Type:       "integer",
	},
	"ReceiverQueue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverQueue",
		DefaultValue:   0,
		Description:    `ReceiverQueue is the base queue number for traffic from the network.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       1000,
		Name:           "receiverQueue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"ReceiverQueueSize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverQueueSize",
		DefaultValue:   500,
		Description:    `ReceiverQueueSize is the queue size of the receiver`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       5000,
		MinValue:       1,
		Name:           "receiverQueueSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"RemoteEnforcerEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RemoteEnforcerEnabled",
		DefaultValue:   true,
		Description: `RemoteEnforcerEnabled inidicates whether a single enforcer should be used or a
distributed enforcer. True means distributed.`,
		Exposed:    true,
		Filterable: true,
		Name:       "remoteEnforcerEnabled",
		Orderable:  true,
		Stored:     true,
		Type:       "boolean",
	},
	"TargetNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNetworks",
		Description:    `TargetNetworks is the list of networks that authorization should be applied.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetNetworks",
		Orderable:      true,
		Stored:         true,
		SubType:        "target_networks_list",
		Type:           "external",
	},
	"TransmitterNumberOfQueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterNumberOfQueues",
		DefaultValue:   4,
		Description:    `TransmitterNumberOfQueues is the number of queues for application traffic.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       16,
		MinValue:       1,
		Name:           "transmitterNumberOfQueues",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"TransmitterQueue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterQueue",
		DefaultValue:   4,
		Description: `TransmitterQueue is the queue number for traffic from the applications starting
at the transmitterQueue`,
		Exposed:    true,
		Filterable: true,
		MaxValue:   1000,
		MinValue:   1,
		Name:       "transmitterQueue",
		Orderable:  true,
		Stored:     true,
		Type:       "integer",
	},
	"TransmitterQueueSize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterQueueSize",
		DefaultValue:   500,
		Description:    `TransmitterQueueSize is the size of the queue for application traffic.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       1000,
		MinValue:       1,
		Name:           "transmitterQueueSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"TrustedCAs": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TrustedCAs",
		Description:    `List of trusted CA. If empty the main chain of trust will be used.`,
		Exposed:        true,
		Name:           "trustedCAs",
		Stored:         true,
		SubType:        "trusted_cas_list",
		Type:           "external",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// EnforcerProfileLowerCaseAttributesMap represents the map of attribute for EnforcerProfile.
var EnforcerProfileLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"iptablesmarkvalue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPTablesMarkValue",
		DefaultValue:   1000,
		Description:    `IptablesMarkValue is the mark value to be used in an iptables implementation.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       65000,
		Name:           "IPTablesMarkValue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"pubookkeepinginterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PUBookkeepingInterval",
		DefaultValue:   "15m",
		Description:    `PUBookkeepingInterval configures how often the PU will be synchronized.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "PUBookkeepingInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"puheartbeatinterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PUHeartbeatInterval",
		DefaultValue:   "5s",
		Description:    `PUHeartbeatInterval configures the heart beat interval.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "PUHeartbeatInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"auditprofileselectors": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditProfileSelectors",
		Description: `AuditProfileSelectors is the list of tags (key/value pairs) that define the
audit policies that must be implemented by this enforcer. The enforcer will
implement all policies that match any of these tags.`,
		Exposed:    true,
		Filterable: true,
		Name:       "auditProfileSelectors",
		Stored:     true,
		SubType:    "audit_profile_selector",
		Type:       "external",
	},
	"auditprofiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "AuditProfiles",
		Description: `AuditProfiles returns the audit rules associated with the enforcer profile. This
is a read only attribute when an enforcer profile is resolved for an enforcer.`,
		Exposed:  true,
		Name:     "auditProfiles",
		ReadOnly: true,
		SubType:  "audit_profiles",
		Type:     "external",
	},
	"auditsocketbuffersize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditSocketBufferSize",
		DefaultValue:   16384,
		Description:    `AuditSocketBufferSize is the size of the audit socket buffer. Default 16384.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       262144,
		Name:           "auditSocketBufferSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"dockersocketaddress": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DockerSocketAddress",
		DefaultValue:   "/var/run/docker.sock",
		Description:    `DockerSocketAddress is the address of the docker daemon.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "dockerSocketAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"dockersockettype": elemental.AttributeSpecification{
		AllowedChoices: []string{"tcp", "unix"},
		ConvertedName:  "DockerSocketType",
		DefaultValue:   EnforcerProfileDockerSocketTypeUnix,
		Description:    `DockerSocketType is the type of socket to use to talk to the docker daemon.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "dockerSocketType",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"excludedinterfaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedInterfaces",
		Description:    `ExcludedInterfaces is a list of interfaces that must be excluded.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "excludedInterfaces",
		Orderable:      true,
		Stored:         true,
		SubType:        "excluded_interfaces_list",
		Type:           "external",
	},
	"excludednetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExcludedNetworks",
		Description: `ExcludedNetworks is the list of networks that must be excluded for this
enforcer.`,
		Exposed:    true,
		Filterable: true,
		Name:       "excludedNetworks",
		Orderable:  true,
		Stored:     true,
		SubType:    "excluded_networks_list",
		Type:       "external",
	},
	"hostservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HostServices",
		Description: `HostServices is a list of services that must be activated by default to all
enforcers matching this profile.`,
		Exposed:    true,
		Filterable: true,
		Name:       "hostServices",
		Orderable:  true,
		Stored:     true,
		SubType:    "host_services_list",
		Type:       "external",
	},
	"ignoreexpression": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IgnoreExpression",
		Description: `IgnoreExpression allows to set a tag expression that will make Aporeto to ignore
docker container started with labels matching the rule.`,
		Exposed: true,
		Name:    "ignoreExpression",
		Stored:  true,
		SubType: "policies_list",
		Type:    "external",
	},
	"kubernetessupportenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesSupportEnabled",
		DefaultValue:   false,
		Description:    `KubernetesSupportEnabled enables kubernetes mode for the enforcer.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "kubernetesSupportEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"linuxprocessessupportenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LinuxProcessesSupportEnabled",
		DefaultValue:   true,
		Description:    `LinuxProcessesSupportEnabled configures support for Linux processes.`,
		Exposed:        true,
		Name:           "linuxProcessesSupportEnabled",
		Stored:         true,
		Type:           "boolean",
	},
	"metadataextractor": elemental.AttributeSpecification{
		AllowedChoices: []string{"Docker", "ECS", "Kubernetes"},
		ConvertedName:  "MetadataExtractor",
		DefaultValue:   EnforcerProfileMetadataExtractorDocker,
		Description:    `Select which metadata extractor to use to process new processing units.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "metadataExtractor",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Index:          true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"policysynchronizationinterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "PolicySynchronizationInterval",
		DefaultValue:   "10m",
		Description: `PolicySynchronizationInterval configures how often the policy will be
resynchronized.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "policySynchronizationInterval",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"proxylistenaddress": elemental.AttributeSpecification{
		AllowedChars:   `^(:([1-9]|[1-9][0-9]|[1-9][0-9]{1,3}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|65535))$|(unix:(/[^/]{1,16}){1,5}/?)$`,
		AllowedChoices: []string{},
		ConvertedName:  "ProxyListenAddress",
		DefaultValue:   "unix:/var/run/aporeto.sock",
		Description: `ProxyListenAddress is the address the enforcer should use to listen for API
calls. It can be a port (example :9443) or socket path
(example: unix:/var/run/aporeto.sock)`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "proxyListenAddress",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"receivernumberofqueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverNumberOfQueues",
		DefaultValue:   4,
		Description: `ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network
receiver starting at the ReceiverQueue`,
		Exposed:    true,
		Filterable: true,
		MaxValue:   16,
		MinValue:   1,
		Name:       "receiverNumberOfQueues",
		Orderable:  true,
		Stored:     true,
		Type:       "integer",
	},
	"receiverqueue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverQueue",
		DefaultValue:   0,
		Description:    `ReceiverQueue is the base queue number for traffic from the network.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       1000,
		Name:           "receiverQueue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"receiverqueuesize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReceiverQueueSize",
		DefaultValue:   500,
		Description:    `ReceiverQueueSize is the queue size of the receiver`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       5000,
		MinValue:       1,
		Name:           "receiverQueueSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"remoteenforcerenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RemoteEnforcerEnabled",
		DefaultValue:   true,
		Description: `RemoteEnforcerEnabled inidicates whether a single enforcer should be used or a
distributed enforcer. True means distributed.`,
		Exposed:    true,
		Filterable: true,
		Name:       "remoteEnforcerEnabled",
		Orderable:  true,
		Stored:     true,
		Type:       "boolean",
	},
	"targetnetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNetworks",
		Description:    `TargetNetworks is the list of networks that authorization should be applied.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetNetworks",
		Orderable:      true,
		Stored:         true,
		SubType:        "target_networks_list",
		Type:           "external",
	},
	"transmitternumberofqueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterNumberOfQueues",
		DefaultValue:   4,
		Description:    `TransmitterNumberOfQueues is the number of queues for application traffic.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       16,
		MinValue:       1,
		Name:           "transmitterNumberOfQueues",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"transmitterqueue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterQueue",
		DefaultValue:   4,
		Description: `TransmitterQueue is the queue number for traffic from the applications starting
at the transmitterQueue`,
		Exposed:    true,
		Filterable: true,
		MaxValue:   1000,
		MinValue:   1,
		Name:       "transmitterQueue",
		Orderable:  true,
		Stored:     true,
		Type:       "integer",
	},
	"transmitterqueuesize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TransmitterQueueSize",
		DefaultValue:   500,
		Description:    `TransmitterQueueSize is the size of the queue for application traffic.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       1000,
		MinValue:       1,
		Name:           "transmitterQueueSize",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"trustedcas": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TrustedCAs",
		Description:    `List of trusted CA. If empty the main chain of trust will be used.`,
		Exposed:        true,
		Name:           "trustedCAs",
		Stored:         true,
		SubType:        "trusted_cas_list",
		Type:           "external",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
