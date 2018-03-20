package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// AutomationTriggerValue represents the possible values for attribute "trigger".
type AutomationTriggerValue string

const (
	// AutomationTriggerEvent represents the value Event.
	AutomationTriggerEvent AutomationTriggerValue = "Event"

	// AutomationTriggerRemotecall represents the value RemoteCall.
	AutomationTriggerRemotecall AutomationTriggerValue = "RemoteCall"

	// AutomationTriggerTime represents the value Time.
	AutomationTriggerTime AutomationTriggerValue = "Time"
)

// AutomationIdentity represents the Identity of the object.
var AutomationIdentity = elemental.Identity{
	Name:     "automation",
	Category: "automations",
	Private:  false,
}

// AutomationsList represents a list of Automations
type AutomationsList []*Automation

// ContentIdentity returns the identity of the objects in the list.
func (o AutomationsList) ContentIdentity() elemental.Identity {

	return AutomationIdentity
}

// Copy returns a pointer to a copy the AutomationsList.
func (o AutomationsList) Copy() elemental.ContentIdentifiable {

	copy := append(AutomationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AutomationsList.
func (o AutomationsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(AutomationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Automation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AutomationsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AutomationsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o AutomationsList) Version() int {

	return 1
}

// Automation represents the model of a automation
type Automation struct {
	// Action contains the code that will be executed if the condition is met.
	Actions []string `json:"actions" bson:"actions" mapstructure:"actions,omitempty"`

	// Condition contains the code that will be executed to decide if any action should
	// be taken.
	Condition string `json:"condition" bson:"condition" mapstructure:"condition,omitempty"`

	// Entitlements declares which operations are allowed on which identities.
	Entitlements map[string][]elemental.Operation `json:"entitlements" bson:"entitlements" mapstructure:"entitlements,omitempty"`

	// Error contains the eventual error of the last run.
	Errors []string `json:"errors" bson:"errors" mapstructure:"errors,omitempty"`

	// Events contains the identity and operation an event must have to trigger the
	// automation
	Events map[string][]elemental.EventType `json:"events" bson:"events" mapstructure:"events,omitempty"`

	// LastExecTime holds the last successful execution tine.
	LastExecTime time.Time `json:"lastExecTime" bson:"lastexectime" mapstructure:"lastExecTime,omitempty"`

	// Parameters are passed to the functions.
	Parameters map[string]interface{} `json:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	// Schedule tells when to run the automation. Must be a valid CRON format. This
	// only applies if the trigger is set to Time.
	Schedule string `json:"schedule" bson:"schedule" mapstructure:"schedule,omitempty"`

	// Stdout contains the last run standard output.
	Stdout string `json:"stdout" bson:"stdout" mapstructure:"stdout,omitempty"`

	// Token holds the unique access token used as a password to trigger the
	// authentication. It will be visible only after creation.
	Token string `json:"token" bson:"token" mapstructure:"token,omitempty"`

	// If set to true a new token will be issued, and the previous one invalidated.
	TokenRenew bool `json:"tokenRenew" bson:"-" mapstructure:"tokenRenew,omitempty"`

	// Trigger controls when the automation should be triggered.
	Trigger AutomationTriggerValue `json:"trigger" bson:"trigger" mapstructure:"trigger,omitempty"`

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
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAutomation returns a new *Automation
func NewAutomation() *Automation {

	return &Automation{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Entitlements:   map[string][]elemental.Operation{},
		Events:         map[string][]elemental.EventType{},
		NormalizedTags: []string{},
		Parameters:     map[string]interface{}{},
		Trigger:        "Time",
	}
}

// Identity returns the Identity of the object.
func (o *Automation) Identity() elemental.Identity {

	return AutomationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Automation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Automation) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Automation) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Automation) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Automation) Doc() string {
	return `An automation needs documentation.`
}

func (o *Automation) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Automation) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *Automation) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Automation) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *Automation) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Automation) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *Automation) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *Automation) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *Automation) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Automation) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *Automation) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Automation) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Automation) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *Automation) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetDisabled returns the Disabled of the receiver.
func (o *Automation) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the given Disabled of the receiver.
func (o *Automation) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetName returns the Name of the receiver.
func (o *Automation) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *Automation) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *Automation) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("condition", o.Condition); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("condition", o.Condition); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("trigger", string(o.Trigger), []string{"Event", "RemoteCall", "Time"}, false); err != nil {
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
func (*Automation) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AutomationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AutomationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Automation) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AutomationAttributesMap
}

// AutomationAttributesMap represents the map of attribute for Automation.
var AutomationAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Actions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Actions",
		Description:    `Action contains the code that will be executed if the condition is met.`,
		Exposed:        true,
		Name:           "actions",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"Condition": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Condition",
		Description: `Condition contains the code that will be executed to decide if any action should
be taken.`,
		Exposed:  true,
		Format:   "free",
		Name:     "condition",
		Required: true,
		Stored:   true,
		Type:     "string",
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
	"Entitlements": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Entitlements",
		Description:    `Entitlements declares which operations are allowed on which identities.`,
		Exposed:        true,
		Name:           "entitlements",
		Stored:         true,
		SubType:        "automation_entitlements",
		Type:           "external",
	},
	"Errors": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Errors",
		Description:    `Error contains the eventual error of the last run.`,
		Exposed:        true,
		Name:           "errors",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Events": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Events",
		Description: `Events contains the identity and operation an event must have to trigger the
automation`,
		Exposed: true,
		Name:    "events",
		Stored:  true,
		SubType: "automation_events",
		Type:    "external",
	},
	"LastExecTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastExecTime",
		Description:    `LastExecTime holds the last successful execution tine.`,
		Exposed:        true,
		Name:           "lastExecTime",
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
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
	"Parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters are passed to the functions.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "automation_parameters",
		Type:           "external",
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
	"Schedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Schedule",
		Description: `Schedule tells when to run the automation. Must be a valid CRON format. This
only applies if the trigger is set to Time.`,
		Exposed: true,
		Format:  "free",
		Name:    "schedule",
		Stored:  true,
		Type:    "string",
	},
	"Stdout": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Stdout",
		Description:    `Stdout contains the last run standard output.`,
		Exposed:        true,
		Format:         "free",
		Name:           "stdout",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		CreationOnly:   true,
		Description: `Token holds the unique access token used as a password to trigger the
authentication. It will be visible only after creation.`,
		Exposed: true,
		Format:  "free",
		Name:    "token",
		Stored:  true,
		Type:    "string",
	},
	"TokenRenew": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TokenRenew",
		Description:    `If set to true a new token will be issued, and the previous one invalidated.`,
		Exposed:        true,
		Name:           "tokenRenew",
		Type:           "boolean",
	},
	"Trigger": elemental.AttributeSpecification{
		AllowedChoices: []string{"Event", "RemoteCall", "Time"},
		ConvertedName:  "Trigger",
		DefaultValue:   AutomationTriggerTime,
		Description:    `Trigger controls when the automation should be triggered.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "trigger",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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

// AutomationLowerCaseAttributesMap represents the map of attribute for Automation.
var AutomationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"actions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Actions",
		Description:    `Action contains the code that will be executed if the condition is met.`,
		Exposed:        true,
		Name:           "actions",
		Required:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
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
	"condition": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Condition",
		Description: `Condition contains the code that will be executed to decide if any action should
be taken.`,
		Exposed:  true,
		Format:   "free",
		Name:     "condition",
		Required: true,
		Stored:   true,
		Type:     "string",
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
	"entitlements": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Entitlements",
		Description:    `Entitlements declares which operations are allowed on which identities.`,
		Exposed:        true,
		Name:           "entitlements",
		Stored:         true,
		SubType:        "automation_entitlements",
		Type:           "external",
	},
	"errors": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Errors",
		Description:    `Error contains the eventual error of the last run.`,
		Exposed:        true,
		Name:           "errors",
		ReadOnly:       true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"events": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Events",
		Description: `Events contains the identity and operation an event must have to trigger the
automation`,
		Exposed: true,
		Name:    "events",
		Stored:  true,
		SubType: "automation_events",
		Type:    "external",
	},
	"lastexectime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LastExecTime",
		Description:    `LastExecTime holds the last successful execution tine.`,
		Exposed:        true,
		Name:           "lastExecTime",
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
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
	"parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters are passed to the functions.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "automation_parameters",
		Type:           "external",
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
	"schedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Schedule",
		Description: `Schedule tells when to run the automation. Must be a valid CRON format. This
only applies if the trigger is set to Time.`,
		Exposed: true,
		Format:  "free",
		Name:    "schedule",
		Stored:  true,
		Type:    "string",
	},
	"stdout": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Stdout",
		Description:    `Stdout contains the last run standard output.`,
		Exposed:        true,
		Format:         "free",
		Name:           "stdout",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		CreationOnly:   true,
		Description: `Token holds the unique access token used as a password to trigger the
authentication. It will be visible only after creation.`,
		Exposed: true,
		Format:  "free",
		Name:    "token",
		Stored:  true,
		Type:    "string",
	},
	"tokenrenew": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TokenRenew",
		Description:    `If set to true a new token will be issued, and the previous one invalidated.`,
		Exposed:        true,
		Name:           "tokenRenew",
		Type:           "boolean",
	},
	"trigger": elemental.AttributeSpecification{
		AllowedChoices: []string{"Event", "RemoteCall", "Time"},
		ConvertedName:  "Trigger",
		DefaultValue:   AutomationTriggerTime,
		Description:    `Trigger controls when the automation should be triggered.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "trigger",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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
