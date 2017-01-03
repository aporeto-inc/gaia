package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"
import "github.com/aporeto-inc/gaia/golang/constants"

// ServerProfileLogLevelValue represents the possible values for attribute "logLevel".
type ServerProfileLogLevelValue string

const (
	// ServerProfileLogLevelDebug represents the value debug.
	ServerProfileLogLevelDebug ServerProfileLogLevelValue = "debug"

	// ServerProfileLogLevelError represents the value error.
	ServerProfileLogLevelError ServerProfileLogLevelValue = "error"

	// ServerProfileLogLevelFatal represents the value fatal.
	ServerProfileLogLevelFatal ServerProfileLogLevelValue = "fatal"

	// ServerProfileLogLevelInfo represents the value info.
	ServerProfileLogLevelInfo ServerProfileLogLevelValue = "info"

	// ServerProfileLogLevelWarn represents the value warn.
	ServerProfileLogLevelWarn ServerProfileLogLevelValue = "warn"
)

// ServerProfileLogFormatValue represents the possible values for attribute "logFormat".
type ServerProfileLogFormatValue string

const (
	// ServerProfileLogFormatJson represents the value json.
	ServerProfileLogFormatJson ServerProfileLogFormatValue = "json"

	// ServerProfileLogFormatText represents the value text.
	ServerProfileLogFormatText ServerProfileLogFormatValue = "text"
)

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

	// AgentPort is the port the agent should use to listen for API calls
	AgentPort int `json:"agentPort" cql:"agentport,omitempty" bson:"agentport"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" cql:"annotation,omitempty" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" cql:"associatedtags,omitempty" bson:"associatedtags"`

	// CreatedAt is the time at which an entity was created
	CreatedAt time.Time `json:"createdAt" cql:"createdat,omitempty" bson:"createdat"`

	// Deleted marks if the entity has been deleted.
	Deleted bool `json:"-" cql:"deleted,omitempty" bson:"deleted"`

	// Description is the description of the object.
	Description string `json:"description" cql:"description,omitempty" bson:"description"`

	// DockerSocketAddress is the address of the docker daemon.
	DockerSocketAddress string `json:"dockerSocketAddress" cql:"dockersocketaddress,omitempty" bson:"dockersocketaddress"`

	// DockerSocketType is the type of socket to use to talk to the docker daemon.
	DockerSocketType ServerProfileDockerSocketTypeValue `json:"dockerSocketType" cql:"dockersockettype,omitempty" bson:"dockersockettype"`

	// IptablesMarkValue is the mark value to be used in an iptables implementation.
	IptablesMarkValue int `json:"iptablesMarkValue" cql:"iptablesmarkvalue,omitempty" bson:"iptablesmarkvalue"`

	// LogFormat is the format for the logs
	LogFormat ServerProfileLogFormatValue `json:"logFormat" cql:"logformat,omitempty" bson:"logformat"`

	// LogID is the ID to use to identity the logs of this server.
	LogID string `json:"logID" cql:"logid,omitempty" bson:"logid"`

	// LogLevel is the level of logging.
	LogLevel ServerProfileLogLevelValue `json:"logLevel" cql:"loglevel,omitempty" bson:"loglevel"`

	// LogServer is the address of the log server.
	LogServer string `json:"logServer" cql:"logserver,omitempty" bson:"logserver"`

	// Name is the name of the entity
	Name string `json:"name" cql:"name,omitempty" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" cql:"namespace,primarykey,omitempty" bson:"_namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" cql:"-" bson:"-"`

	// ParentID is the ID of the parent, if any,
	ParentID string `json:"parentID" cql:"parentid,omitempty" bson:"parentid"`

	// ParentType is the type of the parent, if any. It will be set to the parent's Identity.Name.
	ParentType string `json:"parentType" cql:"parenttype,omitempty" bson:"parenttype"`

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
		AgentPort:              9443,
		AssociatedTags:         []string{},
		DockerSocketAddress:    "/var/run/docker.sock",
		DockerSocketType:       "unix",
		IptablesMarkValue:      1000,
		LogFormat:              "text",
		LogID:                  "Aporeto Agent",
		LogLevel:               "info",
		NormalizedTags:         []string{},
		ReceiverNumberOfQueues: 4,
		ReceiverQueue:          0,
		ReceiverQueueSize:      500,
		RemoteEnforcer:         false,
		Status:                 constants.Active,
		TransmitterNumberOfQueues: 4,
		TransmitterQueue:          4,
		TransmitterQueueSize:      500,
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

func (o *ServerProfile) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ServerProfile) SetIdentifier(ID string) {

	o.ID = ID
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

// GetDeleted returns the deleted of the receiver
func (o *ServerProfile) GetDeleted() bool {
	return o.Deleted
}

// SetDeleted set the given deleted of the receiver
func (o *ServerProfile) SetDeleted(deleted bool) {
	o.Deleted = deleted
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

	if err := elemental.ValidateMaximumInt("agentPort", o.AgentPort, 65000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("agentPort", o.AgentPort, 1000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("dockerSocketType", string(o.DockerSocketType), []string{"tcp", "unix"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("iptablesMarkValue", o.IptablesMarkValue, 65000, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("iptablesMarkValue", o.IptablesMarkValue, 0, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("logFormat", string(o.LogFormat), []string{"json", "text"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("logLevel", string(o.LogLevel), []string{"debug", "error", "fatal", "info", "warn"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
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
	"AgentPort": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		MaxValue:       65000,
		MinValue:       1000,
		Name:           "agentPort",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"Annotation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Exposed:        true,
		Name:           "createdAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Deleted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Getter:         true,
		Name:           "deleted",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Exposed:        true,
		Filterable:     true,
		Name:           "dockerSocketType",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"IptablesMarkValue": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		MaxValue:       65000,
		Name:           "iptablesMarkValue",
		Orderable:      true,
		Stored:         true,
		Type:           "integer",
	},
	"LogFormat": elemental.AttributeSpecification{
		AllowedChoices: []string{"json", "text"},
		Exposed:        true,
		Filterable:     true,
		Name:           "logFormat",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"LogID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "logID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LogLevel": elemental.AttributeSpecification{
		AllowedChoices: []string{"debug", "error", "fatal", "info", "warn"},
		Exposed:        true,
		Filterable:     true,
		Name:           "logLevel",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"LogServer": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "logServer",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"ReceiverNumberOfQueues": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"TargetNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Exposed:        true,
		Name:           "updatedAt",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}
