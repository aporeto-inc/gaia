package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// NetworkAccessPolicyIdentity represents the Identity of the object
var NetworkAccessPolicyIdentity = elemental.Identity{
	Name:     "networkaccesspolicy",
	Category: "networkaccesspolicies",
}

// NetworkAccessPoliciesList represents a list of NetworkAccessPolicies
type NetworkAccessPoliciesList []*NetworkAccessPolicy

// ContentIdentity returns the identity of the objects in the list.
func (o NetworkAccessPoliciesList) ContentIdentity() elemental.Identity {

	return NetworkAccessPolicyIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o NetworkAccessPoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NetworkAccessPoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o NetworkAccessPoliciesList) Version() int {

	return 1
}

// NetworkAccessPolicy represents the model of a networkaccesspolicy
type NetworkAccessPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-"`

	// ActiveDuration defines for how long the policy will be active according to the activeSchedule.
	ActiveDuration string `json:"activeDuration" bson:"activeduration"`

	// ActiveSchedule defines when the policy should be active using the cron notation. The policy will be active for the given activeDuration.
	ActiveSchedule string `json:"activeSchedule" bson:"activeschedule"`

	// AllowsTraffic if true, the flow will be accepted. Otherwise other actions like "logs" can still be done, but the traffic will be rejected.
	AllowsTraffic bool `json:"allowsTraffic" bson:"-"`

	// Annotation stores additional information about an entity
	Annotations map[string][]string `json:"annotations" bson:"annotations"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// DestinationPorts contains the list of allowed ports and ranges.
	DestinationPorts []string `json:"destinationPorts" bson:"-"`

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" bson:"disabled"`

	// EncryptionEnabled defines if the flow has to be encrypted.
	EncryptionEnabled bool `json:"encryptionEnabled" bson:"-"`

	// LogsEnabled defines if the flow has to be logged.
	LogsEnabled bool `json:"logsEnabled" bson:"-"`

	// Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// Object of the policy.
	Object [][]string `json:"object" bson:"-"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" bson:"propagate"`

	// If set to true while the policy is propagating, it won't be visible to children namespace, but still used for policy resolution.
	PropagationHidden bool `json:"propagationHidden" bson:"propagationhidden"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// Subject of the policy.
	Subject [][]string `json:"subject" bson:"-"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewNetworkAccessPolicy returns a new *NetworkAccessPolicy
func NewNetworkAccessPolicy() *NetworkAccessPolicy {

	return &NetworkAccessPolicy{
		ModelVersion:     1,
		Annotations:      map[string][]string{},
		AssociatedTags:   []string{},
		DestinationPorts: []string{},
		Metadata:         []string{},
		NormalizedTags:   []string{},
	}
}

// Identity returns the Identity of the object.
func (o *NetworkAccessPolicy) Identity() elemental.Identity {

	return NetworkAccessPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NetworkAccessPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NetworkAccessPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *NetworkAccessPolicy) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *NetworkAccessPolicy) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *NetworkAccessPolicy) Doc() string {
	return `Allows to define networking policies to allow or prevent processing units identitied by their tags to talk to other processing units or external services (also identified by their tags).`
}

func (o *NetworkAccessPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetActiveDuration returns the activeDuration of the receiver
func (o *NetworkAccessPolicy) GetActiveDuration() string {
	return o.ActiveDuration
}

// SetActiveDuration set the given activeDuration of the receiver
func (o *NetworkAccessPolicy) SetActiveDuration(activeDuration string) {
	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the activeSchedule of the receiver
func (o *NetworkAccessPolicy) GetActiveSchedule() string {
	return o.ActiveSchedule
}

// SetActiveSchedule set the given activeSchedule of the receiver
func (o *NetworkAccessPolicy) SetActiveSchedule(activeSchedule string) {
	o.ActiveSchedule = activeSchedule
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *NetworkAccessPolicy) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *NetworkAccessPolicy) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the createTime of the receiver
func (o *NetworkAccessPolicy) GetCreateTime() time.Time {
	return o.CreateTime
}

// SetCreateTime set the given createTime of the receiver
func (o *NetworkAccessPolicy) SetCreateTime(createTime time.Time) {
	o.CreateTime = createTime
}

// GetDisabled returns the disabled of the receiver
func (o *NetworkAccessPolicy) GetDisabled() bool {
	return o.Disabled
}

// SetDisabled set the given disabled of the receiver
func (o *NetworkAccessPolicy) SetDisabled(disabled bool) {
	o.Disabled = disabled
}

// GetMetadata returns the metadata of the receiver
func (o *NetworkAccessPolicy) GetMetadata() []string {
	return o.Metadata
}

// SetMetadata set the given metadata of the receiver
func (o *NetworkAccessPolicy) SetMetadata(metadata []string) {
	o.Metadata = metadata
}

// GetName returns the name of the receiver
func (o *NetworkAccessPolicy) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *NetworkAccessPolicy) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *NetworkAccessPolicy) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *NetworkAccessPolicy) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *NetworkAccessPolicy) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *NetworkAccessPolicy) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the propagate of the receiver
func (o *NetworkAccessPolicy) GetPropagate() bool {
	return o.Propagate
}

// SetPropagate set the given propagate of the receiver
func (o *NetworkAccessPolicy) SetPropagate(propagate bool) {
	o.Propagate = propagate
}

// GetPropagationHidden returns the propagationHidden of the receiver
func (o *NetworkAccessPolicy) GetPropagationHidden() bool {
	return o.PropagationHidden
}

// SetPropagationHidden set the given propagationHidden of the receiver
func (o *NetworkAccessPolicy) SetPropagationHidden(propagationHidden bool) {
	o.PropagationHidden = propagationHidden
}

// GetProtected returns the protected of the receiver
func (o *NetworkAccessPolicy) GetProtected() bool {
	return o.Protected
}

// GetUpdateTime returns the updateTime of the receiver
func (o *NetworkAccessPolicy) GetUpdateTime() time.Time {
	return o.UpdateTime
}

// SetUpdateTime set the given updateTime of the receiver
func (o *NetworkAccessPolicy) SetUpdateTime(updateTime time.Time) {
	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *NetworkAccessPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidatePattern("activeDuration", o.ActiveDuration, `^[0-9]+[smh]$`, false); err != nil {
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
func (*NetworkAccessPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NetworkAccessPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NetworkAccessPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*NetworkAccessPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NetworkAccessPolicyAttributesMap
}

// NetworkAccessPolicyAttributesMap represents the map of attribute for NetworkAccessPolicy.
var NetworkAccessPolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"ActiveDuration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		Description:    `ActiveDuration defines for how long the policy will be active according to the activeSchedule.`,
		Exposed:        true,
		Format:         "free",
		Getter:         true,
		Name:           "activeDuration",
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ActiveSchedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `ActiveSchedule defines when the policy should be active using the cron notation. The policy will be active for the given activeDuration.`,
		Exposed:        true,
		Getter:         true,
		Name:           "activeSchedule",
		Setter:         true,
		Stored:         true,
		SubType:        "cron_expression",
		Type:           "external",
	},
	"AllowsTraffic": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AllowsTraffic if true, the flow will be accepted. Otherwise other actions like "logs" can still be done, but the traffic will be rejected.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsTraffic",
		Orderable:      true,
		Type:           "boolean",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotations",
		Stored:         true,
		SubType:        "annotations",
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
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"DestinationPorts": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `DestinationPorts contains the list of allowed ports and ranges.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "destinationPorts",
		Orderable:      true,
		SubType:        "ports_list",
		Type:           "external",
	},
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"EncryptionEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `EncryptionEnabled defines if the flow has to be encrypted.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"LogsEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LogsEnabled defines if the flow has to be logged.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "logsEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "metadata",
		Setter:         true,
		Stored:         true,
		SubType:        "metadata_list",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"Object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"PropagationHidden": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `If set to true while the policy is propagating, it won't be visible to children namespace, but still used for policy resolution.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagationHidden",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
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
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Subject of the policy.`,
		Exposed:        true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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

// NetworkAccessPolicyLowerCaseAttributesMap represents the map of attribute for NetworkAccessPolicy.
var NetworkAccessPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"activeduration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		Description:    `ActiveDuration defines for how long the policy will be active according to the activeSchedule.`,
		Exposed:        true,
		Format:         "free",
		Getter:         true,
		Name:           "activeDuration",
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"activeschedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `ActiveSchedule defines when the policy should be active using the cron notation. The policy will be active for the given activeDuration.`,
		Exposed:        true,
		Getter:         true,
		Name:           "activeSchedule",
		Setter:         true,
		Stored:         true,
		SubType:        "cron_expression",
		Type:           "external",
	},
	"allowstraffic": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AllowsTraffic if true, the flow will be accepted. Otherwise other actions like "logs" can still be done, but the traffic will be rejected.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "allowsTraffic",
		Orderable:      true,
		Type:           "boolean",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotations",
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"associatedtags": elemental.AttributeSpecification{
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
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"destinationports": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `DestinationPorts contains the list of allowed ports and ranges.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "destinationPorts",
		Orderable:      true,
		SubType:        "ports_list",
		Type:           "external",
	},
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"encryptionenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `EncryptionEnabled defines if the flow has to be encrypted.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"logsenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LogsEnabled defines if the flow has to be logged.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "logsEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "metadata",
		Setter:         true,
		Stored:         true,
		SubType:        "metadata_list",
		Type:           "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Unique:         true,
	},
	"namespace": elemental.AttributeSpecification{
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
	"normalizedtags": elemental.AttributeSpecification{
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
	"object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"propagationhidden": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `If set to true while the policy is propagating, it won't be visible to children namespace, but still used for policy resolution.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagationHidden",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"protected": elemental.AttributeSpecification{
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
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Subject of the policy.`,
		Exposed:        true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
