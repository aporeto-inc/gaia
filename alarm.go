package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AlarmStatusValue represents the possible values for attribute "status".
type AlarmStatusValue string

const (
	// AlarmStatusAcknowledged represents the value Acknowledged.
	AlarmStatusAcknowledged AlarmStatusValue = "Acknowledged"

	// AlarmStatusOpen represents the value Open.
	AlarmStatusOpen AlarmStatusValue = "Open"

	// AlarmStatusResolved represents the value Resolved.
	AlarmStatusResolved AlarmStatusValue = "Resolved"
)

// AlarmIdentity represents the Identity of the object.
var AlarmIdentity = elemental.Identity{
	Name:     "alarm",
	Category: "alarms",
	Package:  "sephiroth",
	Private:  false,
}

// AlarmsList represents a list of Alarms
type AlarmsList []*Alarm

// Identity returns the identity of the objects in the list.
func (o AlarmsList) Identity() elemental.Identity {

	return AlarmIdentity
}

// Copy returns a pointer to a copy the AlarmsList.
func (o AlarmsList) Copy() elemental.Identifiables {

	copy := append(AlarmsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AlarmsList.
func (o AlarmsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AlarmsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Alarm))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AlarmsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AlarmsList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the AlarmsList converted to SparseAlarmsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AlarmsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAlarmsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAlarm)
	}

	return out
}

// Version returns the version of the content.
func (o AlarmsList) Version() int {

	return 1
}

// Alarm represents the model of a alarm
type Alarm struct {
	// ID is the identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Content of the alarm.
	Content string `json:"content" msgpack:"content" bson:"content" mapstructure:"content,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Data represent user data related to the alams.
	Data []map[string]string `json:"data" msgpack:"data" bson:"data" mapstructure:"data,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Emails is a list of recipients that should be emailed when this alarm is
	// created.
	Emails []string `json:"emails" msgpack:"emails" bson:"emails" mapstructure:"emails,omitempty"`

	// Kind identifies the kind of alarms. If two alarms are created with the same
	// identifier, then only the occurrence will be incremented.
	Kind string `json:"kind" msgpack:"kind" bson:"kind" mapstructure:"kind,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Number of time this alarm have been seen.
	Occurrences []time.Time `json:"occurrences" msgpack:"occurrences" bson:"occurrences" mapstructure:"occurrences,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Status of the alarm.
	Status AlarmStatusValue `json:"status" msgpack:"status" bson:"status" mapstructure:"status,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAlarm returns a new *Alarm
func NewAlarm() *Alarm {

	return &Alarm{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		Data:           []map[string]string{},
		AssociatedTags: []string{},
		Emails:         []string{},
		NormalizedTags: []string{},
		Occurrences:    []time.Time{},
		Status:         AlarmStatusOpen,
	}
}

// Identity returns the Identity of the object.
func (o *Alarm) Identity() elemental.Identity {

	return AlarmIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Alarm) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Alarm) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Alarm) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Alarm) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Alarm) Doc() string {

	return `An alarm represents an event requiring attention.`
}

func (o *Alarm) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Alarm) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Alarm) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Alarm) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Alarm) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *Alarm) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *Alarm) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Alarm) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Alarm) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Alarm) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Alarm) SetDescription(description string) {

	o.Description = description
}

// GetName returns the Name of the receiver.
func (o *Alarm) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Alarm) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Alarm) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Alarm) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Alarm) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Alarm) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Alarm) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Alarm) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *Alarm) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *Alarm) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Alarm) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Alarm) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *Alarm) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *Alarm) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *Alarm) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *Alarm) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Alarm) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAlarm{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			Content:              &o.Content,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Data:                 &o.Data,
			Description:          &o.Description,
			Emails:               &o.Emails,
			Kind:                 &o.Kind,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Occurrences:          &o.Occurrences,
			Protected:            &o.Protected,
			Status:               &o.Status,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparseAlarm{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "content":
			sp.Content = &(o.Content)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "data":
			sp.Data = &(o.Data)
		case "description":
			sp.Description = &(o.Description)
		case "emails":
			sp.Emails = &(o.Emails)
		case "kind":
			sp.Kind = &(o.Kind)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "occurrences":
			sp.Occurrences = &(o.Occurrences)
		case "protected":
			sp.Protected = &(o.Protected)
		case "status":
			sp.Status = &(o.Status)
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

// Patch apply the non nil value of a *SparseAlarm to the object.
func (o *Alarm) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAlarm)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.Content != nil {
		o.Content = *so.Content
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Emails != nil {
		o.Emails = *so.Emails
	}
	if so.Kind != nil {
		o.Kind = *so.Kind
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
	if so.Occurrences != nil {
		o.Occurrences = *so.Occurrences
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Status != nil {
		o.Status = *so.Status
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

// DeepCopy returns a deep copy if the Alarm.
func (o *Alarm) DeepCopy() *Alarm {

	if o == nil {
		return nil
	}

	out := &Alarm{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Alarm.
func (o *Alarm) DeepCopyInto(out *Alarm) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Alarm: %s", err))
	}

	*out = *target.(*Alarm)
}

// Validate valides the current information stored into the structure.
func (o *Alarm) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("content", o.Content); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("kind", o.Kind); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Acknowledged", "Open", "Resolved"}, false); err != nil {
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
func (*Alarm) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AlarmAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AlarmLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Alarm) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AlarmAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Alarm) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "content":
		return o.Content
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "data":
		return o.Data
	case "description":
		return o.Description
	case "emails":
		return o.Emails
	case "kind":
		return o.Kind
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "occurrences":
		return o.Occurrences
	case "protected":
		return o.Protected
	case "status":
		return o.Status
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

// AlarmAttributesMap represents the map of attribute for Alarm.
var AlarmAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		CreationOnly:   true,
		Description:    `Content of the alarm.`,
		Exposed:        true,
		Name:           "content",
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data represent user data related to the alams.`,
		Exposed:        true,
		Name:           "data",
		Stored:         true,
		SubType:        "[]map[string]string",
		Type:           "external",
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
	"Emails": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Emails",
		Description: `Emails is a list of recipients that should be emailed when this alarm is
created.`,
		Exposed: true,
		Name:    "emails",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"Kind": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Kind",
		CreationOnly:   true,
		Description: `Kind identifies the kind of alarms. If two alarms are created with the same
identifier, then only the occurrence will be incremented.`,
		Exposed:   true,
		Name:      "kind",
		Orderable: true,
		Required:  true,
		Stored:    true,
		Type:      "string",
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
	"Occurrences": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Occurrences",
		CreationOnly:   true,
		Description:    `Number of time this alarm have been seen.`,
		Exposed:        true,
		Name:           "occurrences",
		Stored:         true,
		SubType:        "[]time.Time",
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
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Acknowledged", "Open", "Resolved"},
		ConvertedName:  "Status",
		DefaultValue:   AlarmStatusOpen,
		Description:    `Status of the alarm.`,
		Exposed:        true,
		Name:           "status",
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
		Exposed:   true,
		Getter:    true,
		Name:      "zone",
		ReadOnly:  true,
		Setter:    true,
		Stored:    true,
		Transient: true,
		Type:      "integer",
	},
}

// AlarmLowerCaseAttributesMap represents the map of attribute for Alarm.
var AlarmLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		CreationOnly:   true,
		Description:    `Content of the alarm.`,
		Exposed:        true,
		Name:           "content",
		Required:       true,
		Stored:         true,
		Type:           "string",
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
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Data represent user data related to the alams.`,
		Exposed:        true,
		Name:           "data",
		Stored:         true,
		SubType:        "[]map[string]string",
		Type:           "external",
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
	"emails": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Emails",
		Description: `Emails is a list of recipients that should be emailed when this alarm is
created.`,
		Exposed: true,
		Name:    "emails",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
	"kind": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Kind",
		CreationOnly:   true,
		Description: `Kind identifies the kind of alarms. If two alarms are created with the same
identifier, then only the occurrence will be incremented.`,
		Exposed:   true,
		Name:      "kind",
		Orderable: true,
		Required:  true,
		Stored:    true,
		Type:      "string",
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
	"occurrences": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Occurrences",
		CreationOnly:   true,
		Description:    `Number of time this alarm have been seen.`,
		Exposed:        true,
		Name:           "occurrences",
		Stored:         true,
		SubType:        "[]time.Time",
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
	"status": elemental.AttributeSpecification{
		AllowedChoices: []string{"Acknowledged", "Open", "Resolved"},
		ConvertedName:  "Status",
		DefaultValue:   AlarmStatusOpen,
		Description:    `Status of the alarm.`,
		Exposed:        true,
		Name:           "status",
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
		Exposed:   true,
		Getter:    true,
		Name:      "zone",
		ReadOnly:  true,
		Setter:    true,
		Stored:    true,
		Transient: true,
		Type:      "integer",
	},
}

// SparseAlarmsList represents a list of SparseAlarms
type SparseAlarmsList []*SparseAlarm

// Identity returns the identity of the objects in the list.
func (o SparseAlarmsList) Identity() elemental.Identity {

	return AlarmIdentity
}

// Copy returns a pointer to a copy the SparseAlarmsList.
func (o SparseAlarmsList) Copy() elemental.Identifiables {

	copy := append(SparseAlarmsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAlarmsList.
func (o SparseAlarmsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAlarmsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAlarm))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAlarmsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAlarmsList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseAlarmsList converted to AlarmsList.
func (o SparseAlarmsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAlarmsList) Version() int {

	return 1
}

// SparseAlarm represents the sparse version of a alarm.
type SparseAlarm struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// Content of the alarm.
	Content *string `json:"content,omitempty" msgpack:"content,omitempty" bson:"content,omitempty" mapstructure:"content,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Data represent user data related to the alams.
	Data *[]map[string]string `json:"data,omitempty" msgpack:"data,omitempty" bson:"data,omitempty" mapstructure:"data,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Emails is a list of recipients that should be emailed when this alarm is
	// created.
	Emails *[]string `json:"emails,omitempty" msgpack:"emails,omitempty" bson:"emails,omitempty" mapstructure:"emails,omitempty"`

	// Kind identifies the kind of alarms. If two alarms are created with the same
	// identifier, then only the occurrence will be incremented.
	Kind *string `json:"kind,omitempty" msgpack:"kind,omitempty" bson:"kind,omitempty" mapstructure:"kind,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Number of time this alarm have been seen.
	Occurrences *[]time.Time `json:"occurrences,omitempty" msgpack:"occurrences,omitempty" bson:"occurrences,omitempty" mapstructure:"occurrences,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Status of the alarm.
	Status *AlarmStatusValue `json:"status,omitempty" msgpack:"status,omitempty" bson:"status,omitempty" mapstructure:"status,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAlarm returns a new  SparseAlarm.
func NewSparseAlarm() *SparseAlarm {
	return &SparseAlarm{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAlarm) Identity() elemental.Identity {

	return AlarmIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAlarm) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAlarm) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseAlarm) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAlarm) ToPlain() elemental.PlainIdentifiable {

	out := NewAlarm()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.Content != nil {
		out.Content = *o.Content
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Emails != nil {
		out.Emails = *o.Emails
	}
	if o.Kind != nil {
		out.Kind = *o.Kind
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
	if o.Occurrences != nil {
		out.Occurrences = *o.Occurrences
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Status != nil {
		out.Status = *o.Status
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
func (o *SparseAlarm) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseAlarm) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseAlarm) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseAlarm) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseAlarm) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseAlarm) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseAlarm) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAlarm) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseAlarm) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseAlarm) SetDescription(description string) {

	o.Description = &description
}

// GetName returns the Name of the receiver.
func (o *SparseAlarm) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseAlarm) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseAlarm) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseAlarm) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseAlarm) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseAlarm) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseAlarm) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseAlarm) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseAlarm) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseAlarm) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAlarm) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAlarm) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseAlarm) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseAlarm) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseAlarm) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseAlarm) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseAlarm.
func (o *SparseAlarm) DeepCopy() *SparseAlarm {

	if o == nil {
		return nil
	}

	out := &SparseAlarm{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAlarm.
func (o *SparseAlarm) DeepCopyInto(out *SparseAlarm) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAlarm: %s", err))
	}

	*out = *target.(*SparseAlarm)
}
