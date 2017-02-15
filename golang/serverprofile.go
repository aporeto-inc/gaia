package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/golang/constants"

// ServerProfileDockerSocketTypeValue represents the possible values for attribute "dockerSocketType".
type ServerProfileDockerSocketTypeValue string

const (
	// ServerProfileDockerSocketTypeTcp represents the value tcp.
	ServerProfileDockerSocketTypeTcp ServerProfileDockerSocketTypeValue = "tcp"

	// ServerProfileDockerSocketTypeUnix represents the value unix.
	ServerProfileDockerSocketTypeUnix ServerProfileDockerSocketTypeValue = "unix"
)

// ServerProfileIdentity represents the Identity of the object
var ServerProfileIdentity = elemental.Identity{
	Name:     "serverprofile",
	Category: "serverprofiles",
}

// ServerProfilesList represents a list of ServerProfiles
type ServerProfilesList []*ServerProfile

// ServerProfile represents the model of a serverprofile
type ServerProfile struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"id,primarykey,omitempty" bson:"_id"`

	// IptablesMarkValue is the mark value to be used in an iptables implementation.
	IPTablesMarkValue int `json:"IPTablesMarkValue" cql:"iptablesmarkvalue,omitempty" bson:"iptablesmarkvalue"`

	// IgnoreExpression allows to set a tag expression that will make Aporeto to ignore docker container started with labels matching the rule.
	IgnoreExpression [][]string `json:"IgnoreExpression" cql:"ignoreexpression,omitempty" bson:"ignoreexpression"`

	// PUBookkeepingInterval configures how often the PU will be synchronized.
	PUBookkeepingInterval string `json:"PUBookkeepingInterval" cql:"pubookkeepinginterval,omitempty" bson:"pubookkeepinginterval"`

	// PUHeartbeatInterval configures the heart beat interval.
	PUHeartbeatInterval string `json:"PUHeartbeatInterval" cql:"puheartbeatinterval,omitempty" bson:"puheartbeatinterval"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty" bson:"associatedtags"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty" bson:"createdat"`

	// Description is the description of the object.
	Description string `json:"description" cql:"description,omitempty" bson:"description"`

	// DockerSocketAddress is the address of the docker daemon.
	DockerSocketAddress string `json:"dockerSocketAddress" cql:"dockersocketaddress,omitempty" bson:"dockersocketaddress"`

	// DockerSocketType is the type of socket to use to talk to the docker daemon.
	DockerSocketType ServerProfileDockerSocketTypeValue `json:"dockerSocketType" cql:"dockersockettype,omitempty" bson:"dockersockettype"`

	// kubernetesEnable enables  kubernetes mode for the agent.
	KubernetesEnable bool `json:"kubernetesEnable" cql:"kubernetesenable,omitempty" bson:"kubernetesenable"`

	// Name is the name of the entity
	Name string `json:"name" cql:"name,omitempty" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" cql:"normalizedtags,omitempty" bson:"normalizedtags"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty" bson:"parenttype"`

	// PolicySynchronizationInterval configures how often the policy will be resynchronized.
	PolicySynchronizationInterval string `json:"policySynchronizationInterval" cql:"policysynchronizationinterval,omitempty" bson:"policysynchronizationinterval"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" cql:"protected,omitempty" bson:"protected"`

	// AgentPort is the port the agent should use to listen for API calls
	ProxyListenAddress string `json:"proxyListenAddress" cql:"proxylistenaddress,omitempty" bson:"proxylistenaddress"`

	// ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network receiver starting at the ReceiverQueue
	ReceiverNumberOfQueues int `json:"receiverNumberOfQueues" cql:"receivernumberofqueues,omitempty" bson:"receivernumberofqueues"`

	// ReceiverQueue is the base queue number for traffic from the network.
	ReceiverQueue int `json:"receiverQueue" cql:"receiverqueue,omitempty" bson:"receiverqueue"`

	// ReceiverQueueSize is the queue size of the receiver
	ReceiverQueueSize int `json:"receiverQueueSize" cql:"receiverqueuesize,omitempty" bson:"receiverqueuesize"`

	// RemoteEnforcer inidicates whether a single enforcer should be used or a distributed enforcer. True means distributed.
	RemoteEnforcer bool `json:"remoteEnforcer" cql:"remoteenforcer,omitempty" bson:"remoteenforcer"`

	// Status of an entity
	Status constants.EntityStatus `json:"status" cql:"status,omitempty" bson:"status"`

	// SupportLinuxProcesses configures support for Linux processes.
	SupportLinuxProcesses bool `json:"supportLinuxProcesses" cql:"supportlinuxprocesses,omitempty" bson:"supportlinuxprocesses"`

	// TargetNetworks is the list of networks that authorization should be applied.
	TargetNetworks []string `json:"targetNetworks" cql:"targetnetworks,omitempty" bson:"targetnetworks"`

	// TransmitterNumberOfQueues is the number of queues for application traffic.
	TransmitterNumberOfQueues int `json:"transmitterNumberOfQueues" cql:"transmitternumberofqueues,omitempty" bson:"transmitternumberofqueues"`

	// TransmitterQueue is the queue number for traffic from the applications starting at the transmitterQueue
	TransmitterQueue int `json:"transmitterQueue" cql:"transmitterqueue,omitempty" bson:"transmitterqueue"`

	// TransmitterQueueSize is the size of the queue for application traffic.
	TransmitterQueueSize int `json:"transmitterQueueSize" cql:"transmitterqueuesize,omitempty" bson:"transmitterqueuesize"`

	// UpdatedAt is the time at which an entity was updated.
	UpdatedAt time.Time `json:"updatedAt" cql:"updatedat,omitempty" bson:"updatedat"`
}

// NewServerProfile returns a new *ServerProfile
func NewServerProfile() *ServerProfile {

	return &ServerProfile{
		IPTablesMarkValue:             1000,
		PUBookkeepingInterval:         "5m",
		PUHeartbeatInterval:           "5s",
		AssociatedTags:                []string{},
		DockerSocketAddress:           "/var/run/docker.sock",
		DockerSocketType:              "unix",
		KubernetesEnable:              false,
		NormalizedTags:                []string{},
		PolicySynchronizationInterval: "10m",
		ProxyListenAddress:            ":9443",
		ReceiverNumberOfQueues:        4,
		ReceiverQueue:                 0,
		ReceiverQueueSize:             500,
		RemoteEnforcer:                false,
		Status:                        constants.Active,
		SupportLinuxProcesses:         false,
		TransmitterNumberOfQueues:     4,
		TransmitterQueue:              4,
		TransmitterQueueSize:          500,
	}
}

// Identity returns the Identity of the object.
func (o *ServerProfile) Identity() elemental.Identity {

	return ServerProfileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ServerProfile) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ServerProfile) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *ServerProfile) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *ServerProfile) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *ServerProfile) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreatedAt set the given createdAt of the receiver
func (o *ServerProfile) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

// GetName returns the name of the receiver
func (o *ServerProfile) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *ServerProfile) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *ServerProfile) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *ServerProfile) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *ServerProfile) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *ServerProfile) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetParentID returns the parentID of the receiver
func (o *ServerProfile) GetParentID() string {
	return o.ParentID
}

// SetParentID set the given parentID of the receiver
func (o *ServerProfile) SetParentID(parentID string) {
	o.ParentID = parentID
}

// GetParentType returns the parentType of the receiver
func (o *ServerProfile) GetParentType() string {
	return o.ParentType
}

// SetParentType set the given parentType of the receiver
func (o *ServerProfile) SetParentType(parentType string) {
	o.ParentType = parentType
}

// GetProtected returns the protected of the receiver
func (o *ServerProfile) GetProtected() bool {
	return o.Protected
}

// GetStatus returns the status of the receiver
func (o *ServerProfile) GetStatus() constants.EntityStatus {
	return o.Status
}

// SetStatus set the given status of the receiver
func (o *ServerProfile) SetStatus(status constants.EntityStatus) {
	o.Status = status
}

// SetUpdatedAt set the given updatedAt of the receiver
func (o *ServerProfile) SetUpdatedAt(updatedAt time.Time) {
	o.UpdatedAt = updatedAt
}

// Validate valides the current information stored into the structure.
func (o *ServerProfile) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumInt("IPTablesMarkValue", o.IPTablesMarkValue, 65000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("IPTablesMarkValue", o.IPTablesMarkValue, 0, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("PUBookkeepingInterval", o.PUBookkeepingInterval, `^[0-9]+[smh]$`); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("PUHeartbeatInterval", o.PUHeartbeatInterval, `^[0-9]+[smh]$`); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("dockerSocketType", string(o.DockerSocketType), []string{"tcp", "unix"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("policySynchronizationInterval", o.PolicySynchronizationInterval, `^[0-9]+[smh]$`); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverNumberOfQueues", o.ReceiverNumberOfQueues, 16, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("receiverNumberOfQueues", o.ReceiverNumberOfQueues, 1, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverQueue", o.ReceiverQueue, 1000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("receiverQueue", o.ReceiverQueue, 0, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("receiverQueueSize", o.ReceiverQueueSize, 5000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("receiverQueueSize", o.ReceiverQueueSize, 1, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("targetNetworks", o.TargetNetworks); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("targetNetworks", o.TargetNetworks); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterNumberOfQueues", o.TransmitterNumberOfQueues, 16, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterNumberOfQueues", o.TransmitterNumberOfQueues, 1, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterQueue", o.TransmitterQueue, 1000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterQueue", o.TransmitterQueue, 1, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("transmitterQueueSize", o.TransmitterQueueSize, 1000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("transmitterQueueSize", o.TransmitterQueueSize, 1, false); err != nil {
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
func (ServerProfile) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ServerProfileAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (ServerProfile) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServerProfileAttributesMap
}

// ServerProfileAttributesMap represents the map of attribute for ServerProfile.
var ServerProfileAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
		Unique:         true,
	},
	"IPTablesMarkValue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `IptablesMarkValue is the mark value to be used in an iptables implementation.`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       65000,
		Name:           "IPTablesMarkValue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"IgnoreExpression": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `IgnoreExpression allows to set a tag expression that will make Aporeto to ignore docker container started with labels matching the rule.`,
		Exposed:        true,
		Name:           "IgnoreExpression",
		Stored:         true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"PUBookkeepingInterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
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
		Description:    `PUHeartbeatInterval configures the heart beat interval.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "PUHeartbeatInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Annotation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AssociatedTags are the list of tags attached to an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"CreatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedAt is the time at which an entity was created`,
		Exposed:        true,
		Name:           "createdAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Description:    `DockerSocketType is the type of socket to use to talk to the docker daemon.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "dockerSocketType",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"KubernetesEnable": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `kubernetesEnable enables  kubernetes mode for the agent.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "kubernetesEnable",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Unique:         true,
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentID is the ID of the parent, if any,`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"PolicySynchronizationInterval": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		Description:    `PolicySynchronizationInterval configures how often the policy will be resynchronized.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "policySynchronizationInterval",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		AllowedChoices: []string{},
		Description:    `AgentPort is the port the agent should use to listen for API calls `,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "proxyListenAddress",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ReceiverNumberOfQueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `ReceiverNumberOfQueues is the number of queues for the NFQUEUE of the network receiver starting at the ReceiverQueue`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       16,
		MinValue:       1,
		Name:           "receiverNumberOfQueues",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"ReceiverQueue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"RemoteEnforcer": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `RemoteEnforcer inidicates whether a single enforcer should be used or a distributed enforcer. True means distributed.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "remoteEnforcer",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Status of an entity`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "status",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		SubType:        "status_enum",
		Type:           "external",
	},
	"SupportLinuxProcesses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `SupportLinuxProcesses configures support for Linux processes.`,
		Exposed:        true,
		Name:           "supportLinuxProcesses",
		Stored:         true,
		Type:           "boolean",
	},
	"TargetNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `TargetNetworks is the list of networks that authorization should be applied.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetNetworks",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		SubType:        "target_networks_list",
		Type:           "external",
	},
	"TransmitterNumberOfQueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Description:    `TransmitterQueue is the queue number for traffic from the applications starting at the transmitterQueue`,
		Exposed:        true,
		Filterable:     true,
		MaxValue:       1000,
		MinValue:       1,
		Name:           "transmitterQueue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"TransmitterQueueSize": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"UpdatedAt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdatedAt is the time at which an entity was updated.`,
		Exposed:        true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
