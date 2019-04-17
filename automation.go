package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AutomationTriggerValue represents the possible values for attribute "trigger".
type AutomationTriggerValue string

const (
	// AutomationTriggerEvent represents the value Event.
	AutomationTriggerEvent AutomationTriggerValue = "Event"

	// AutomationTriggerRemoteCall represents the value RemoteCall.
	AutomationTriggerRemoteCall AutomationTriggerValue = "RemoteCall"

	// AutomationTriggerTime represents the value Time.
	AutomationTriggerTime AutomationTriggerValue = "Time"
)

// AutomationIdentity represents the Identity of the object.
var AutomationIdentity = elemental.Identity{
	Name:     "automation",
	Category: "automations",
	Package:  "sephiroth",
	Private:  false,
}

// AutomationsList represents a list of Automations
type AutomationsList []*Automation

// Identity returns the identity of the objects in the list.
func (o AutomationsList) Identity() elemental.Identity {

	return AutomationIdentity
}

// Copy returns a pointer to a copy the AutomationsList.
func (o AutomationsList) Copy() elemental.Identifiables {

	copy := append(AutomationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AutomationsList.
func (o AutomationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AutomationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Automation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AutomationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AutomationsList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the AutomationsList converted to SparseAutomationsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AutomationsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAutomationsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAutomation)
	}

	return out
}

// Version returns the version of the content.
func (o AutomationsList) Version() int {

	return 1
}

// Automation represents the model of a automation
type Automation struct {
	// ID is the identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Action contains the code that will be executed if the condition is met.
	Actions []string `json:"actions" msgpack:"actions" bson:"actions" mapstructure:"actions,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Condition contains the code that will be executed to decide if any action should
	// be taken.
	Condition string `json:"condition" msgpack:"condition" bson:"condition" mapstructure:"condition,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" msgpack:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// Entitlements declares which operations are allowed on which identities.
	Entitlements map[string][]elemental.Operation `json:"entitlements" msgpack:"entitlements" bson:"entitlements" mapstructure:"entitlements,omitempty"`

	// Error contains the eventual error of the last run.
	Errors []string `json:"errors" msgpack:"errors" bson:"errors" mapstructure:"errors,omitempty"`

	// Events contains the identity and operation an event must have to trigger the
	// automation.
	Events map[string][]elemental.EventType `json:"events" msgpack:"events" bson:"events" mapstructure:"events,omitempty"`

	// LastExecTime holds the last successful execution tine.
	LastExecTime time.Time `json:"lastExecTime" msgpack:"lastExecTime" bson:"lastexectime" mapstructure:"lastExecTime,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Parameters are passed to the functions.
	Parameters map[string]interface{} `json:"parameters" msgpack:"parameters" bson:"parameters" mapstructure:"parameters,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Schedule tells when to run the automation. Must be a valid CRON format. This
	// only applies if the trigger is set to Time.
	Schedule string `json:"schedule" msgpack:"schedule" bson:"schedule" mapstructure:"schedule,omitempty"`

	// Stdout contains the last run standard output.
	Stdout string `json:"stdout" msgpack:"stdout" bson:"stdout" mapstructure:"stdout,omitempty"`

	// Token holds the unique access token used as a password to trigger the
	// authentication. It will be visible only after creation.
	Token string `json:"token" msgpack:"token" bson:"token" mapstructure:"token,omitempty"`

	// If set to true a new token will be issued, and the previous one invalidated.
	TokenRenew bool `json:"tokenRenew" msgpack:"tokenRenew" bson:"-" mapstructure:"tokenRenew,omitempty"`

	// Trigger controls when the automation should be triggered.
	Trigger AutomationTriggerValue `json:"trigger" msgpack:"trigger" bson:"trigger" mapstructure:"trigger,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAutomation returns a new *Automation
func NewAutomation() *Automation {

	return &Automation{
		ModelVersion:   1,
		Actions:        []string{},
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Events:         map[string][]elemental.EventType{},
		Entitlements:   map[string][]elemental.Operation{},
		Errors:         []string{},
		NormalizedTags: []string{},
		Parameters:     map[string]interface{}{},
		Trigger:        AutomationTriggerTime,
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
		"namespace",
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

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Automation) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Automation) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Automation) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *Automation) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *Automation) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Automation) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Automation) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Automation) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Automation) SetDescription(description string) {

	o.Description = description
}

// GetDisabled returns the Disabled of the receiver.
func (o *Automation) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the given value.
func (o *Automation) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetName returns the Name of the receiver.
func (o *Automation) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Automation) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Automation) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Automation) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Automation) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Automation) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Automation) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Automation) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *Automation) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *Automation) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Automation) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Automation) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Automation) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Automation) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Automation) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Automation) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Automation) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAutomation{
			ID:                   &o.ID,
			Actions:              &o.Actions,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			Condition:            &o.Condition,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Disabled:             &o.Disabled,
			Entitlements:         &o.Entitlements,
			Errors:               &o.Errors,
			Events:               &o.Events,
			LastExecTime:         &o.LastExecTime,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Parameters:           &o.Parameters,
			Protected:            &o.Protected,
			Schedule:             &o.Schedule,
			Stdout:               &o.Stdout,
			Token:                &o.Token,
			TokenRenew:           &o.TokenRenew,
			Trigger:              &o.Trigger,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseAutomation{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "actions":
			sp.Actions = &(o.Actions)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "condition":
			sp.Condition = &(o.Condition)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "disabled":
			sp.Disabled = &(o.Disabled)
		case "entitlements":
			sp.Entitlements = &(o.Entitlements)
		case "errors":
			sp.Errors = &(o.Errors)
		case "events":
			sp.Events = &(o.Events)
		case "lastExecTime":
			sp.LastExecTime = &(o.LastExecTime)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "parameters":
			sp.Parameters = &(o.Parameters)
		case "protected":
			sp.Protected = &(o.Protected)
		case "schedule":
			sp.Schedule = &(o.Schedule)
		case "stdout":
			sp.Stdout = &(o.Stdout)
		case "token":
			sp.Token = &(o.Token)
		case "tokenRenew":
			sp.TokenRenew = &(o.TokenRenew)
		case "trigger":
			sp.Trigger = &(o.Trigger)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAutomation to the object.
func (o *Automation) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAutomation)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Actions != nil {
		o.Actions = *so.Actions
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.Condition != nil {
		o.Condition = *so.Condition
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Disabled != nil {
		o.Disabled = *so.Disabled
	}
	if so.Entitlements != nil {
		o.Entitlements = *so.Entitlements
	}
	if so.Errors != nil {
		o.Errors = *so.Errors
	}
	if so.Events != nil {
		o.Events = *so.Events
	}
	if so.LastExecTime != nil {
		o.LastExecTime = *so.LastExecTime
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Parameters != nil {
		o.Parameters = *so.Parameters
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Schedule != nil {
		o.Schedule = *so.Schedule
	}
	if so.Stdout != nil {
		o.Stdout = *so.Stdout
	}
	if so.Token != nil {
		o.Token = *so.Token
	}
	if so.TokenRenew != nil {
		o.TokenRenew = *so.TokenRenew
	}
	if so.Trigger != nil {
		o.Trigger = *so.Trigger
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the Automation.
func (o *Automation) DeepCopy() *Automation {

	if o == nil {
		return nil
	}

	out := &Automation{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Automation.
func (o *Automation) DeepCopyInto(out *Automation) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Automation: %s", err))
	}

	*out = *target.(*Automation)
}

// Validate valides the current information stored into the structure.
func (o *Automation) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("condition", o.Condition); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("trigger", string(o.Trigger), []string{"Event", "RemoteCall", "Time"}, false); err != nil {
		errors = errors.Append(err)
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Automation) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "actions":
		return o.Actions
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "condition":
		return o.Condition
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "disabled":
		return o.Disabled
	case "entitlements":
		return o.Entitlements
	case "errors":
		return o.Errors
	case "events":
		return o.Events
	case "lastExecTime":
		return o.LastExecTime
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "parameters":
		return o.Parameters
	case "protected":
		return o.Protected
	case "schedule":
		return o.Schedule
	case "stdout":
		return o.Stdout
	case "token":
		return o.Token
	case "tokenRenew":
		return o.TokenRenew
	case "trigger":
		return o.Trigger
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
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
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
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
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Condition": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Condition",
		Description: `Condition contains the code that will be executed to decide if any action should
be taken.`,
		Exposed:  true,
		Name:     "condition",
		Required: true,
		Stored:   true,
		Type:     "string",
	},
	"CreateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
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
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
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
		SubType:        "_automation_entitlements",
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
automation.`,
		Exposed: true,
		Name:    "events",
		Stored:  true,
		SubType: "_automation_events",
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
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
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
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
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
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters are passed to the functions.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Schedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Schedule",
		Description: `Schedule tells when to run the automation. Must be a valid CRON format. This
only applies if the trigger is set to Time.`,
		Exposed: true,
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
		Name:           "stdout",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description: `Token holds the unique access token used as a password to trigger the
authentication. It will be visible only after creation.`,
		Exposed:   true,
		Name:      "token",
		Stored:    true,
		Transient: true,
		Type:      "string",
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
		Name:           "trigger",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"UpdateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
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
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
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
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"condition": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Condition",
		Description: `Condition contains the code that will be executed to decide if any action should
be taken.`,
		Exposed:  true,
		Name:     "condition",
		Required: true,
		Stored:   true,
		Type:     "string",
	},
	"createidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
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
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
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
		SubType:        "_automation_entitlements",
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
automation.`,
		Exposed: true,
		Name:    "events",
		Stored:  true,
		SubType: "_automation_events",
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
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
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
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
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
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters are passed to the functions.`,
		Exposed:        true,
		Name:           "parameters",
		Stored:         true,
		SubType:        "map[string]interface{}",
		Type:           "external",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"schedule": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Schedule",
		Description: `Schedule tells when to run the automation. Must be a valid CRON format. This
only applies if the trigger is set to Time.`,
		Exposed: true,
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
		Name:           "stdout",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Token",
		Description: `Token holds the unique access token used as a password to trigger the
authentication. It will be visible only after creation.`,
		Exposed:   true,
		Name:      "token",
		Stored:    true,
		Transient: true,
		Type:      "string",
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
		Name:           "trigger",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"updateidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// SparseAutomationsList represents a list of SparseAutomations
type SparseAutomationsList []*SparseAutomation

// Identity returns the identity of the objects in the list.
func (o SparseAutomationsList) Identity() elemental.Identity {

	return AutomationIdentity
}

// Copy returns a pointer to a copy the SparseAutomationsList.
func (o SparseAutomationsList) Copy() elemental.Identifiables {

	copy := append(SparseAutomationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAutomationsList.
func (o SparseAutomationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAutomationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAutomation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAutomationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAutomationsList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseAutomationsList converted to AutomationsList.
func (o SparseAutomationsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAutomationsList) Version() int {

	return 1
}

// SparseAutomation represents the sparse version of a automation.
type SparseAutomation struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Action contains the code that will be executed if the condition is met.
	Actions *[]string `json:"actions,omitempty" msgpack:"actions,omitempty" bson:"actions,omitempty" mapstructure:"actions,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Condition contains the code that will be executed to decide if any action should
	// be taken.
	Condition *string `json:"condition,omitempty" msgpack:"condition,omitempty" bson:"condition,omitempty" mapstructure:"condition,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled *bool `json:"disabled,omitempty" msgpack:"disabled,omitempty" bson:"disabled,omitempty" mapstructure:"disabled,omitempty"`

	// Entitlements declares which operations are allowed on which identities.
	Entitlements *map[string][]elemental.Operation `json:"entitlements,omitempty" msgpack:"entitlements,omitempty" bson:"entitlements,omitempty" mapstructure:"entitlements,omitempty"`

	// Error contains the eventual error of the last run.
	Errors *[]string `json:"errors,omitempty" msgpack:"errors,omitempty" bson:"errors,omitempty" mapstructure:"errors,omitempty"`

	// Events contains the identity and operation an event must have to trigger the
	// automation.
	Events *map[string][]elemental.EventType `json:"events,omitempty" msgpack:"events,omitempty" bson:"events,omitempty" mapstructure:"events,omitempty"`

	// LastExecTime holds the last successful execution tine.
	LastExecTime *time.Time `json:"lastExecTime,omitempty" msgpack:"lastExecTime,omitempty" bson:"lastexectime,omitempty" mapstructure:"lastExecTime,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Parameters are passed to the functions.
	Parameters *map[string]interface{} `json:"parameters,omitempty" msgpack:"parameters,omitempty" bson:"parameters,omitempty" mapstructure:"parameters,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Schedule tells when to run the automation. Must be a valid CRON format. This
	// only applies if the trigger is set to Time.
	Schedule *string `json:"schedule,omitempty" msgpack:"schedule,omitempty" bson:"schedule,omitempty" mapstructure:"schedule,omitempty"`

	// Stdout contains the last run standard output.
	Stdout *string `json:"stdout,omitempty" msgpack:"stdout,omitempty" bson:"stdout,omitempty" mapstructure:"stdout,omitempty"`

	// Token holds the unique access token used as a password to trigger the
	// authentication. It will be visible only after creation.
	Token *string `json:"token,omitempty" msgpack:"token,omitempty" bson:"token,omitempty" mapstructure:"token,omitempty"`

	// If set to true a new token will be issued, and the previous one invalidated.
	TokenRenew *bool `json:"tokenRenew,omitempty" msgpack:"tokenRenew,omitempty" bson:"-" mapstructure:"tokenRenew,omitempty"`

	// Trigger controls when the automation should be triggered.
	Trigger *AutomationTriggerValue `json:"trigger,omitempty" msgpack:"trigger,omitempty" bson:"trigger,omitempty" mapstructure:"trigger,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAutomation returns a new  SparseAutomation.
func NewSparseAutomation() *SparseAutomation {
	return &SparseAutomation{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAutomation) Identity() elemental.Identity {

	return AutomationIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAutomation) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAutomation) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseAutomation) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAutomation) ToPlain() elemental.PlainIdentifiable {

	out := NewAutomation()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Actions != nil {
		out.Actions = *o.Actions
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.Condition != nil {
		out.Condition = *o.Condition
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Disabled != nil {
		out.Disabled = *o.Disabled
	}
	if o.Entitlements != nil {
		out.Entitlements = *o.Entitlements
	}
	if o.Errors != nil {
		out.Errors = *o.Errors
	}
	if o.Events != nil {
		out.Events = *o.Events
	}
	if o.LastExecTime != nil {
		out.LastExecTime = *o.LastExecTime
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Parameters != nil {
		out.Parameters = *o.Parameters
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Schedule != nil {
		out.Schedule = *o.Schedule
	}
	if o.Stdout != nil {
		out.Stdout = *o.Stdout
	}
	if o.Token != nil {
		out.Token = *o.Token
	}
	if o.TokenRenew != nil {
		out.TokenRenew = *o.TokenRenew
	}
	if o.Trigger != nil {
		out.Trigger = *o.Trigger
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseAutomation) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseAutomation) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseAutomation) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseAutomation) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseAutomation) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseAutomation) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseAutomation) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAutomation) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseAutomation) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseAutomation) SetDescription(description string) {

	o.Description = &description
}

// GetDisabled returns the Disabled of the receiver.
func (o *SparseAutomation) GetDisabled() bool {

	return *o.Disabled
}

// SetDisabled sets the property Disabled of the receiver using the address of the given value.
func (o *SparseAutomation) SetDisabled(disabled bool) {

	o.Disabled = &disabled
}

// GetName returns the Name of the receiver.
func (o *SparseAutomation) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseAutomation) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseAutomation) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseAutomation) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseAutomation) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseAutomation) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseAutomation) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseAutomation) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseAutomation) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseAutomation) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAutomation) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAutomation) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAutomation) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAutomation) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAutomation) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAutomation) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAutomation.
func (o *SparseAutomation) DeepCopy() *SparseAutomation {

	if o == nil {
		return nil
	}

	out := &SparseAutomation{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAutomation.
func (o *SparseAutomation) DeepCopyInto(out *SparseAutomation) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAutomation: %s", err))
	}

	*out = *target.(*SparseAutomation)
}
