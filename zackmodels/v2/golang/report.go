package zackmodels

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// ReportKindValue represents the possible values for attribute "kind".
type ReportKindValue string

const (
	// ReportKindEnforcer represents the value Enforcer.
	ReportKindEnforcer ReportKindValue = "Enforcer"

	// ReportKindFileaccess represents the value FileAccess.
	ReportKindFileaccess ReportKindValue = "FileAccess"

	// ReportKindFlow represents the value Flow.
	ReportKindFlow ReportKindValue = "Flow"

	// ReportKindProcessingunit represents the value ProcessingUnit.
	ReportKindProcessingunit ReportKindValue = "ProcessingUnit"

	// ReportKindSyscall represents the value Syscall.
	ReportKindSyscall ReportKindValue = "Syscall"
)

// ReportIdentity represents the Identity of the object.
var ReportIdentity = elemental.Identity{
	Name:     "report",
	Category: "reports",
}

// ReportsList represents a list of Reports
type ReportsList []*Report

// ContentIdentity returns the identity of the objects in the list.
func (o ReportsList) ContentIdentity() elemental.Identity {

	return ReportIdentity
}

// Copy returns a pointer to a copy the ReportsList.
func (o ReportsList) Copy() elemental.ContentIdentifiable {

	copy := append(ReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ReportsList.
func (o ReportsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(ReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Report))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ReportsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ReportsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o ReportsList) Version() int {

	return 2
}

// Report represents the model of a report
type Report struct {
	// TSDB Fields to set for the report.
	Fields map[string]interface{} `json:"fields" bson:"-"`

	// Kind contains the kind of report.
	Kind ReportKindValue `json:"kind" bson:"-"`

	// Tags contains the tags associated to the data point.
	Tags map[string]string `json:"tags" bson:"-"`

	// Timestamp contains the time for the report.
	Timestamp time.Time `json:"timestamp" bson:"-"`

	// Value contains the value for the report.
	Value float64 `json:"value" bson:"-"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewReport returns a new *Report
func NewReport() *Report {

	return &Report{
		ModelVersion: 2,
		Fields:       map[string]interface{}{},
		Tags:         map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *Report) Identity() elemental.Identity {

	return ReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Report) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Report) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Report) Version() int {

	return 2
}

// DefaultOrder returns the list of default ordering fields.
func (o *Report) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Report) Doc() string {
	return `Post a new statistics report.`
}

func (o *Report) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Report) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("kind", string(o.Kind), []string{"Enforcer", "FileAccess", "Flow", "ProcessingUnit", "Syscall"}, false); err != nil {
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
func (*Report) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Report) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ReportAttributesMap
}

// ReportAttributesMap represents the map of attribute for Report.
var ReportAttributesMap = map[string]elemental.AttributeSpecification{
	"Fields": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fields",
		Description:    `TSDB Fields to set for the report.`,
		Exposed:        true,
		Name:           "fields",
		SubType:        "tsdb_fields",
		Type:           "external",
	},
	"Kind": elemental.AttributeSpecification{
		AllowedChoices: []string{"Enforcer", "FileAccess", "Flow", "ProcessingUnit", "Syscall"},
		ConvertedName:  "Kind",
		Description:    `Kind contains the kind of report.`,
		Exposed:        true,
		Name:           "kind",
		Type:           "enum",
	},
	"Tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `Tags contains the tags associated to the data point.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "tags_map",
		Type:           "external",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp contains the time for the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"Value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Value contains the value for the report.`,
		Exposed:        true,
		Name:           "value",
		Type:           "float",
	},
}

// ReportLowerCaseAttributesMap represents the map of attribute for Report.
var ReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"fields": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fields",
		Description:    `TSDB Fields to set for the report.`,
		Exposed:        true,
		Name:           "fields",
		SubType:        "tsdb_fields",
		Type:           "external",
	},
	"kind": elemental.AttributeSpecification{
		AllowedChoices: []string{"Enforcer", "FileAccess", "Flow", "ProcessingUnit", "Syscall"},
		ConvertedName:  "Kind",
		Description:    `Kind contains the kind of report.`,
		Exposed:        true,
		Name:           "kind",
		Type:           "enum",
	},
	"tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `Tags contains the tags associated to the data point.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "tags_map",
		Type:           "external",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp contains the time for the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Value contains the value for the report.`,
		Exposed:        true,
		Name:           "value",
		Type:           "float",
	},
}
