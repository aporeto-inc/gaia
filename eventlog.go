package gaia

import (
	"fmt"
	"sync"
	"time"

	"go.aporeto.io/elemental"
)

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
	Package:  "leon",
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

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EventLogsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the EventLogsList converted to SparseEventLogsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EventLogsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
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

	sync.Mutex `json:"-" bson:"-"`
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

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EventLog) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEventLog{
			Category:       &o.Category,
			Content:        &o.Content,
			Date:           &o.Date,
			Level:          &o.Level,
			Namespace:      &o.Namespace,
			TargetID:       &o.TargetID,
			TargetIdentity: &o.TargetIdentity,
			Title:          &o.Title,
		}
	}

	sp := &SparseEventLog{}
	for _, f := range fields {
		switch f {
		case "category":
			sp.Category = &(o.Category)
		case "content":
			sp.Content = &(o.Content)
		case "date":
			sp.Date = &(o.Date)
		case "level":
			sp.Level = &(o.Level)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "targetID":
			sp.TargetID = &(o.TargetID)
		case "targetIdentity":
			sp.TargetIdentity = &(o.TargetIdentity)
		case "title":
			sp.Title = &(o.Title)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEventLog to the object.
func (o *EventLog) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEventLog)
	if so.Category != nil {
		o.Category = *so.Category
	}
	if so.Content != nil {
		o.Content = *so.Content
	}
	if so.Date != nil {
		o.Date = *so.Date
	}
	if so.Level != nil {
		o.Level = *so.Level
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.TargetID != nil {
		o.TargetID = *so.TargetID
	}
	if so.TargetIdentity != nil {
		o.TargetIdentity = *so.TargetIdentity
	}
	if so.Title != nil {
		o.Title = *so.Title
	}
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

// SparseEventLogsList represents a list of SparseEventLogs
type SparseEventLogsList []*SparseEventLog

// Identity returns the identity of the objects in the list.
func (o SparseEventLogsList) Identity() elemental.Identity {

	return EventLogIdentity
}

// Copy returns a pointer to a copy the SparseEventLogsList.
func (o SparseEventLogsList) Copy() elemental.Identifiables {

	copy := append(SparseEventLogsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEventLogsList.
func (o SparseEventLogsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEventLogsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEventLog))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEventLogsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEventLogsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseEventLogsList converted to EventLogsList.
func (o SparseEventLogsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEventLogsList) Version() int {

	return 1
}

// SparseEventLog represents the sparse version of a eventlog.
type SparseEventLog struct {
	// Category of the log.
	Category *string `json:"category,omitempty" bson:"category" mapstructure:"category,omitempty"`

	// Content of the log.
	Content *string `json:"content,omitempty" bson:"content" mapstructure:"content,omitempty"`

	// Creation date of the eventlog.
	Date *time.Time `json:"date,omitempty" bson:"date" mapstructure:"date,omitempty"`

	// Represent the level of the log .
	Level *EventLogLevelValue `json:"level,omitempty" bson:"level" mapstructure:"level,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// ID of the object this eventlog is attached to. The object must be in the same
	// namespace than the eventlog.
	TargetID *string `json:"targetID,omitempty" bson:"targetid" mapstructure:"targetID,omitempty"`

	// Identity of the object this eventlog is attached to.
	TargetIdentity *string `json:"targetIdentity,omitempty" bson:"targetidentity" mapstructure:"targetIdentity,omitempty"`

	// Title of the eventlog.
	Title *string `json:"title,omitempty" bson:"title" mapstructure:"title,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseEventLog returns a new  SparseEventLog.
func NewSparseEventLog() *SparseEventLog {
	return &SparseEventLog{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEventLog) Identity() elemental.Identity {

	return EventLogIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEventLog) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEventLog) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseEventLog) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEventLog) ToPlain() elemental.PlainIdentifiable {

	out := NewEventLog()
	if o.Category != nil {
		out.Category = *o.Category
	}
	if o.Content != nil {
		out.Content = *o.Content
	}
	if o.Date != nil {
		out.Date = *o.Date
	}
	if o.Level != nil {
		out.Level = *o.Level
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.TargetID != nil {
		out.TargetID = *o.TargetID
	}
	if o.TargetIdentity != nil {
		out.TargetIdentity = *o.TargetIdentity
	}
	if o.Title != nil {
		out.Title = *o.Title
	}

	return out
}
