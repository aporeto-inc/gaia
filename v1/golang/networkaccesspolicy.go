package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// NetworkAccessPolicyObservedTrafficActionValue represents the possible values for attribute "observedTrafficAction".
type NetworkAccessPolicyObservedTrafficActionValue string

const (
	// NetworkAccessPolicyObservedTrafficActionApply represents the value Apply.
	NetworkAccessPolicyObservedTrafficActionApply NetworkAccessPolicyObservedTrafficActionValue = "Apply"

	// NetworkAccessPolicyObservedTrafficActionContinue represents the value Continue.
	NetworkAccessPolicyObservedTrafficActionContinue NetworkAccessPolicyObservedTrafficActionValue = "Continue"
)

// NetworkAccessPolicyIdentity represents the Identity of the object.
var NetworkAccessPolicyIdentity = elemental.Identity{
	Name:     "networkaccesspolicy",
	Category: "networkaccesspolicies",
	Private:  false,
}

// NetworkAccessPoliciesList represents a list of NetworkAccessPolicies
type NetworkAccessPoliciesList []*NetworkAccessPolicy

// ContentIdentity returns the identity of the objects in the list.
func (o NetworkAccessPoliciesList) ContentIdentity() elemental.Identity {

	return NetworkAccessPolicyIdentity
}

// Copy returns a pointer to a copy the NetworkAccessPoliciesList.
func (o NetworkAccessPoliciesList) Copy() elemental.ContentIdentifiable {

	copy := append(NetworkAccessPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NetworkAccessPoliciesList.
func (o NetworkAccessPoliciesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(NetworkAccessPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*NetworkAccessPolicy))
	}

	return out
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
	// AllowsTraffic if true, the flow will be accepted. Otherwise other actions like
	// "logs" can still be done, but the traffic will be rejected.
	AllowsTraffic bool `json:"allowsTraffic" bson:"-" mapstructure:"allowsTraffic,omitempty"`

	// DestinationPorts contains the list of allowed ports and ranges.
	DestinationPorts []string `json:"destinationPorts" bson:"-" mapstructure:"destinationPorts,omitempty"`

	// EncryptionEnabled defines if the flow has to be encrypted.
	EncryptionEnabled bool `json:"encryptionEnabled" bson:"-" mapstructure:"encryptionEnabled,omitempty"`

	// LogsEnabled defines if the flow has to be logged.
	LogsEnabled bool `json:"logsEnabled" bson:"-" mapstructure:"logsEnabled,omitempty"`

	// Object of the policy.
	Object [][]string `json:"object" bson:"-" mapstructure:"object,omitempty"`

	// If set to true, the flow will be in observation mode.
	ObservationEnabled bool `json:"observationEnabled" bson:"-" mapstructure:"observationEnabled,omitempty"`

	// If observationEnabled is set to true, this will be the final action taken on the
	// packets.
	ObservedTrafficAction NetworkAccessPolicyObservedTrafficActionValue `json:"observedTrafficAction" bson:"-" mapstructure:"observedTrafficAction,omitempty"`

	// List of tags expressions to match the list of entity to pass the flow through.
	Passthrough [][]string `json:"passthrough" bson:"-" mapstructure:"passthrough,omitempty"`

	// Subject of the policy.
	Subject [][]string `json:"subject" bson:"-" mapstructure:"subject,omitempty"`

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

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// If set to true while the policy is propagating, it won't be visible to children
	// namespace, but still used for policy resolution.
	PropagationHidden bool `json:"propagationHidden" bson:"propagationhidden" mapstructure:"propagationHidden,omitempty"`

	// ActiveDuration defines for how long the policy will be active according to the
	// activeSchedule.
	ActiveDuration string `json:"activeDuration" bson:"activeduration" mapstructure:"activeDuration,omitempty"`

	// ActiveSchedule defines when the policy should be active using the cron notation.
	// The policy will be active for the given activeDuration.
	ActiveSchedule string `json:"activeSchedule" bson:"activeschedule" mapstructure:"activeSchedule,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewNetworkAccessPolicy returns a new *NetworkAccessPolicy
func NewNetworkAccessPolicy() *NetworkAccessPolicy {

	return &NetworkAccessPolicy{
		ModelVersion:          1,
		Annotations:           map[string][]string{},
		AssociatedTags:        []string{},
		DestinationPorts:      []string{},
		Metadata:              []string{},
		NormalizedTags:        []string{},
		ObservedTrafficAction: "Continue",
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
func (o *NetworkAccessPolicy) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
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
	return `Allows to define networking policies to allow or prevent processing units
identitied by their tags to talk to other processing units or external services
(also identified by their tags).`
}

func (o *NetworkAccessPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *NetworkAccessPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *NetworkAccessPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *NetworkAccessPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *NetworkAccessPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *NetworkAccessPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *NetworkAccessPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *NetworkAccessPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *NetworkAccessPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *NetworkAccessPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *NetworkAccessPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *NetworkAccessPolicy) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *NetworkAccessPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *NetworkAccessPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetDisabled returns the Disabled of the receiver.
func (o *NetworkAccessPolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the given Disabled of the receiver.
func (o *NetworkAccessPolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetMetadata returns the Metadata of the receiver.
func (o *NetworkAccessPolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *NetworkAccessPolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *NetworkAccessPolicy) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *NetworkAccessPolicy) SetName(name string) {

	o.Name = name
}

// GetPropagate returns the Propagate of the receiver.
func (o *NetworkAccessPolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the given Propagate of the receiver.
func (o *NetworkAccessPolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetPropagationHidden returns the PropagationHidden of the receiver.
func (o *NetworkAccessPolicy) GetPropagationHidden() bool {

	return o.PropagationHidden
}

// SetPropagationHidden sets the given PropagationHidden of the receiver.
func (o *NetworkAccessPolicy) SetPropagationHidden(propagationHidden bool) {

	o.PropagationHidden = propagationHidden
}

// GetActiveDuration returns the ActiveDuration of the receiver.
func (o *NetworkAccessPolicy) GetActiveDuration() string {

	return o.ActiveDuration
}

// SetActiveDuration sets the given ActiveDuration of the receiver.
func (o *NetworkAccessPolicy) SetActiveDuration(activeDuration string) {

	o.ActiveDuration = activeDuration
}

// GetActiveSchedule returns the ActiveSchedule of the receiver.
func (o *NetworkAccessPolicy) GetActiveSchedule() string {

	return o.ActiveSchedule
}

// SetActiveSchedule sets the given ActiveSchedule of the receiver.
func (o *NetworkAccessPolicy) SetActiveSchedule(activeSchedule string) {

	o.ActiveSchedule = activeSchedule
}

// Validate valides the current information stored into the structure.
func (o *NetworkAccessPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("observedTrafficAction", string(o.ObservedTrafficAction), []string{"Apply", "Continue"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("activeDuration", o.ActiveDuration, `^[0-9]+[smh]$`, false); err != nil {
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
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"ActiveDuration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "ActiveDuration",
		Description: `ActiveDuration defines for how long the policy will be active according to the
activeSchedule.`,
		Exposed: true,
		Format:  "free",
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"ActiveSchedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ActiveSchedule",
		Description: `ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		SubType: "cron_expression",
		Type:    "external",
	},
	"AllowsTraffic": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsTraffic",
		Description: `AllowsTraffic if true, the flow will be accepted. Otherwise other actions like
"logs" can still be done, but the traffic will be rejected.`,
		Exposed:    true,
		Filterable: true,
		Name:       "allowsTraffic",
		Orderable:  true,
		Type:       "boolean",
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
	"DestinationPorts": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPorts",
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
		ConvertedName:  "Disabled",
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
		ConvertedName:  "EncryptionEnabled",
		Description:    `EncryptionEnabled defines if the flow has to be encrypted.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"LogsEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogsEnabled",
		Description:    `LogsEnabled defines if the flow has to be logged.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "logsEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "metadata_list",
		Type:       "external",
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
	"Object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description:    `Object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"ObservationEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservationEnabled",
		Description:    `If set to true, the flow will be in observation mode.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "observationEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"ObservedTrafficAction": elemental.AttributeSpecification{
		AllowedChoices: []string{"Apply", "Continue"},
		ConvertedName:  "ObservedTrafficAction",
		DefaultValue:   NetworkAccessPolicyObservedTrafficActionContinue,
		Description: `If observationEnabled is set to true, this will be the final action taken on the
packets.`,
		Exposed:    true,
		Filterable: true,
		Name:       "observedTrafficAction",
		Orderable:  true,
		Type:       "enum",
	},
	"Passthrough": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Passthrough",
		Description:    `List of tags expressions to match the list of entity to pass the flow through.`,
		Exposed:        true,
		Name:           "passthrough",
		SubType:        "policies_list",
		Type:           "external",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
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
		ConvertedName:  "PropagationHidden",
		Description: `If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "propagationHidden",
		Orderable:  true,
		Setter:     true,
		Stored:     true,
		Type:       "boolean",
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
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
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

// NetworkAccessPolicyLowerCaseAttributesMap represents the map of attribute for NetworkAccessPolicy.
var NetworkAccessPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		ReadOnly:       true,
		Type:           "string",
	},
	"activeduration": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "ActiveDuration",
		Description: `ActiveDuration defines for how long the policy will be active according to the
activeSchedule.`,
		Exposed: true,
		Format:  "free",
		Getter:  true,
		Name:    "activeDuration",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"activeschedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ActiveSchedule",
		Description: `ActiveSchedule defines when the policy should be active using the cron notation.
The policy will be active for the given activeDuration.`,
		Exposed: true,
		Getter:  true,
		Name:    "activeSchedule",
		Setter:  true,
		Stored:  true,
		SubType: "cron_expression",
		Type:    "external",
	},
	"allowstraffic": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AllowsTraffic",
		Description: `AllowsTraffic if true, the flow will be accepted. Otherwise other actions like
"logs" can still be done, but the traffic will be rejected.`,
		Exposed:    true,
		Filterable: true,
		Name:       "allowsTraffic",
		Orderable:  true,
		Type:       "boolean",
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
	"destinationports": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPorts",
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
		ConvertedName:  "Disabled",
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
		ConvertedName:  "EncryptionEnabled",
		Description:    `EncryptionEnabled defines if the flow has to be encrypted.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "encryptionEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"logsenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LogsEnabled",
		Description:    `LogsEnabled defines if the flow has to be logged.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "logsEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "metadata_list",
		Type:       "external",
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
	"object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description:    `Object of the policy.`,
		Exposed:        true,
		Name:           "object",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"observationenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservationEnabled",
		Description:    `If set to true, the flow will be in observation mode.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "observationEnabled",
		Orderable:      true,
		Type:           "boolean",
	},
	"observedtrafficaction": elemental.AttributeSpecification{
		AllowedChoices: []string{"Apply", "Continue"},
		ConvertedName:  "ObservedTrafficAction",
		DefaultValue:   NetworkAccessPolicyObservedTrafficActionContinue,
		Description: `If observationEnabled is set to true, this will be the final action taken on the
packets.`,
		Exposed:    true,
		Filterable: true,
		Name:       "observedTrafficAction",
		Orderable:  true,
		Type:       "enum",
	},
	"passthrough": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Passthrough",
		Description:    `List of tags expressions to match the list of entity to pass the flow through.`,
		Exposed:        true,
		Name:           "passthrough",
		SubType:        "policies_list",
		Type:           "external",
	},
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
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
		ConvertedName:  "PropagationHidden",
		Description: `If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "propagationHidden",
		Orderable:  true,
		Setter:     true,
		Stored:     true,
		Type:       "boolean",
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
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
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
