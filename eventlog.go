package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// EventLogIndexes lists the attribute compound indexes.
var EventLogIndexes = [][]string{}

// EventLogLevelValue represents the possible values for attribute "level".
type EventLogLevelValue string

const (
	// EventLogLevelCritical represents the value Critical.
	EventLogLevelCritical EventLogLevelValue = "Critical"

	// EventLogLevelDebug represents the value Debug.
	EventLogLevelDebug EventLogLevelValue = "Debug"

	// EventLogLevelError represents the value Error.
	EventLogLevelError EventLogLevelValue = "Error"

	// EventLogLevelInfo represents the value Info.
	EventLogLevelInfo EventLogLevelValue = "Info"

	// EventLogLevelWarning represents the value Warning.
	EventLogLevelWarning EventLogLevelValue = "Warning"
)

// EventLogIdentity represents the Identity of the object.
var EventLogIdentity = elemental.Identity{
	Name:     "eventlog",
	Category: "eventlogs",
	Private:  false,
}

// EventLogsList represents a list of EventLogs
type EventLogsList []*EventLog

// Identity returns the identity of the objects in the list.
func (o EventLogsList) Identity() elemental.Identity {

	return EventLogIdentity
}

// Copy returns a pointer to a copy the EventLogsList.
func (o EventLogsList) Copy() elemental.Identifiables {

	copy := append(EventLogsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EventLogsList.
func (o EventLogsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EventLogsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EventLog))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EventLogsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EventLogsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o EventLogsList) Version() int {

	return 1
}

// EventLog represents the model of a eventlog
type EventLog struct {
	// Category of the log.
	Category string `json:"category" bson:"category" mapstructure:"category,omitempty"`

	// Content of the log.
	Content string `json:"content" bson:"content" mapstructure:"content,omitempty"`

	// Creation date of the eventlog.
	Date time.Time `json:"date" bson:"date" mapstructure:"date,omitempty"`

	// Represent the level of the log .
	Level EventLogLevelValue `json:"level" bson:"level" mapstructure:"level,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// ID of the object this eventlog is attached to. The object must be in the same
	// namespace than the eventlog.
	TargetID string `json:"targetID" bson:"targetid" mapstructure:"targetID,omitempty"`

	// Identity of the object this eventlog is attached to.
	TargetIdentity string `json:"targetIdentity" bson:"targetidentity" mapstructure:"targetIdentity,omitempty"`

	// Title of the eventlog.
	Title string `json:"title" bson:"title" mapstructure:"title,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewEventLog returns a new *EventLog
func NewEventLog() *EventLog {

	return &EventLog{
		ModelVersion: 1,
		Level:        EventLogLevelInfo,
	}
}

// Identity returns the Identity of the object.
func (o *EventLog) Identity() elemental.Identity {

	return EventLogIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EventLog) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EventLog) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *EventLog) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *EventLog) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *EventLog) Doc() string {
	return `This api allows to report various event on any objects.`
}

func (o *EventLog) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetNamespace returns the Namespace of the receiver.
func (o *EventLog) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *EventLog) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// Validate valides the current information stored into the structure.
func (o *EventLog) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("category", o.Category); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("content", o.Content); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("level", string(o.Level), []string{"Debug", "Info", "Warning", "Error", "Critical"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("targetID", o.TargetID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("targetIdentity", o.TargetIdentity); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("title", o.Title); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*EventLog) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EventLogAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EventLogLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EventLog) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EventLogAttributesMap
}

// EventLogAttributesMap represents the map of attribute for EventLog.
var EventLogAttributesMap = map[string]elemental.AttributeSpecification{
	"Category": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Category",
		CreationOnly:   true,
		Description:    `Category of the log.`,
		Exposed:        true,
		Name:           "category",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		CreationOnly:   true,
		Description:    `Content of the log.`,
		Exposed:        true,
		Name:           "content",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Date",
		CreationOnly:   true,
		Description:    `Creation date of the eventlog.`,
		Exposed:        true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Level": elemental.AttributeSpecification{
		AllowedChoices: []string{"Debug", "Info", "Warning", "Error", "Critical"},
		ConvertedName:  "Level",
		CreationOnly:   true,
		DefaultValue:   EventLogLevelInfo,
		Description:    `Represent the level of the log .`,
		Exposed:        true,
		Name:           "level",
		Stored:         true,
		Type:           "enum",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"TargetID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetID",
		CreationOnly:   true,
		Description: `ID of the object this eventlog is attached to. The object must be in the same
namespace than the eventlog.`,
		Exposed:  true,
		Name:     "targetID",
		Required: true,
		Stored:   true,
		Type:     "string",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		CreationOnly:   true,
		Description:    `Identity of the object this eventlog is attached to.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		CreationOnly:   true,
		Description:    `Title of the eventlog.`,
		Exposed:        true,
		Name:           "title",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// EventLogLowerCaseAttributesMap represents the map of attribute for EventLog.
var EventLogLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"category": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Category",
		CreationOnly:   true,
		Description:    `Category of the log.`,
		Exposed:        true,
		Name:           "category",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"content": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Content",
		CreationOnly:   true,
		Description:    `Content of the log.`,
		Exposed:        true,
		Name:           "content",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Date",
		CreationOnly:   true,
		Description:    `Creation date of the eventlog.`,
		Exposed:        true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"level": elemental.AttributeSpecification{
		AllowedChoices: []string{"Debug", "Info", "Warning", "Error", "Critical"},
		ConvertedName:  "Level",
		CreationOnly:   true,
		DefaultValue:   EventLogLevelInfo,
		Description:    `Represent the level of the log .`,
		Exposed:        true,
		Name:           "level",
		Stored:         true,
		Type:           "enum",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"targetid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetID",
		CreationOnly:   true,
		Description: `ID of the object this eventlog is attached to. The object must be in the same
namespace than the eventlog.`,
		Exposed:  true,
		Name:     "targetID",
		Required: true,
		Stored:   true,
		Type:     "string",
	},
	"targetidentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		CreationOnly:   true,
		Description:    `Identity of the object this eventlog is attached to.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		CreationOnly:   true,
		Description:    `Title of the eventlog.`,
		Exposed:        true,
		Name:           "title",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}
