package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// MessageLevelValue represents the possible values for attribute "level".
type MessageLevelValue string

const (
	// MessageLevelDanger represents the value Danger.
	MessageLevelDanger MessageLevelValue = "Danger"

	// MessageLevelInfo represents the value Info.
	MessageLevelInfo MessageLevelValue = "Info"

	// MessageLevelWarning represents the value Warning.
	MessageLevelWarning MessageLevelValue = "Warning"
)

// MessageIdentity represents the Identity of the object.
var MessageIdentity = elemental.Identity{
	Name:     "message",
	Category: "messages",
	Private:  false,
}

// MessagesList represents a list of Messages
type MessagesList []*Message

// ContentIdentity returns the identity of the objects in the list.
func (o MessagesList) ContentIdentity() elemental.Identity {

	return MessageIdentity
}

// Copy returns a pointer to a copy the MessagesList.
func (o MessagesList) Copy() elemental.ContentIdentifiable {

	copy := append(MessagesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the MessagesList.
func (o MessagesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(MessagesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Message))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o MessagesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o MessagesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o MessagesList) Version() int {

	return 1
}

// Message represents the model of a message
type Message struct {
	// expirationTime is the time after which the message will be deleted.
	ExpirationTime time.Time `json:"expirationTime" bson:"expirationtime" mapstructure:"expirationTime,omitempty"`

	// Level defines how the message is important.
	Level MessageLevelValue `json:"level" bson:"level" mapstructure:"level,omitempty"`

	// If local is set, the message will only be visible in the current namespace.
	Local bool `json:"local" bson:"local" mapstructure:"local,omitempty"`

	// If enabled, the message will be sent to the email associated in namespaces annotations.
	NotifyByEmail bool `json:"notifyByEmail" bson:"-" mapstructure:"notifyByEmail,omitempty"`

	// Validity set using golang time duration, when the message will be automatically deleted.
	Validity string `json:"validity" bson:"validity" mapstructure:"validity,omitempty"`

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

// NewMessage returns a new *Message
func NewMessage() *Message {

	return &Message{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Level:          "Info",
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *Message) Identity() elemental.Identity {

	return MessageIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Message) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Message) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Message) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Message) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Message) Doc() string {
	return `The Message API allows to post public messages that will be visible through all children namespaces`
}

func (o *Message) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Message) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *Message) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Message) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *Message) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Message) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *Message) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *Message) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *Message) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Message) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *Message) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Message) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Message) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *Message) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetName returns the Name of the receiver.
func (o *Message) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *Message) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *Message) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("level", string(o.Level), []string{"Danger", "Info", "Warning"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidatePattern("validity", o.Validity, `^[0-9]+[smh]$`, false); err != nil {
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
func (*Message) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := MessageAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return MessageLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Message) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return MessageAttributesMap
}

// MessageAttributesMap represents the map of attribute for Message.
var MessageAttributesMap = map[string]elemental.AttributeSpecification{
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
	"ExpirationTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `expirationTime is the time after which the message will be deleted.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "expirationTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Level": elemental.AttributeSpecification{
		AllowedChoices: []string{"Danger", "Info", "Warning"},
		ConvertedName:  "Level",
		DefaultValue:   MessageLevelInfo,
		Description:    `Level defines how the message is important.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "level",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"Local": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Local",
		Description:    `If local is set, the message will only be visible in the current namespace.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "local",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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
	"NotifyByEmail": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NotifyByEmail",
		CreationOnly:   true,
		Description:    `If enabled, the message will be sent to the email associated in namespaces annotations.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "notifyByEmail",
		Type:           "boolean",
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
	"Validity": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		Description:    `Validity set using golang time duration, when the message will be automatically deleted.`,
		Exposed:        true,
		Format:         "free",
		Name:           "validity",
		Stored:         true,
		Type:           "string",
	},
}

// MessageLowerCaseAttributesMap represents the map of attribute for Message.
var MessageLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"expirationtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExpirationTime",
		Description:    `expirationTime is the time after which the message will be deleted.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "expirationTime",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"level": elemental.AttributeSpecification{
		AllowedChoices: []string{"Danger", "Info", "Warning"},
		ConvertedName:  "Level",
		DefaultValue:   MessageLevelInfo,
		Description:    `Level defines how the message is important.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "level",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"local": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Local",
		Description:    `If local is set, the message will only be visible in the current namespace.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "local",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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
	"notifybyemail": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NotifyByEmail",
		CreationOnly:   true,
		Description:    `If enabled, the message will be sent to the email associated in namespaces annotations.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "notifyByEmail",
		Type:           "boolean",
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
	"validity": elemental.AttributeSpecification{
		AllowedChars:   `^[0-9]+[smh]$`,
		AllowedChoices: []string{},
		ConvertedName:  "Validity",
		Description:    `Validity set using golang time duration, when the message will be automatically deleted.`,
		Exposed:        true,
		Format:         "free",
		Name:           "validity",
		Stored:         true,
		Type:           "string",
	},
}
