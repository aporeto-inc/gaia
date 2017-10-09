package zackmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// ReportIdentity represents the Identity of the object
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

	return 1
}

// Report represents the model of a report
type Report struct {
	// Name contains the name metric of statistics data.
	Name string `json:"name" bson:"-"`

	// Tags contains the tags associated to the data point.
	Tags map[string]string `json:"tags" bson:"-"`

	// Timestamp contains the unix timestamp for the report.
	Timestamp int64 `json:"timestamp" bson:"-"`

	// TTL set the time to live for the report in seconds.
	Ttl int `json:"ttl" bson:"-"`

	// Value contains the value for the report.
	Value float64 `json:"value" bson:"-"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewReport returns a new *Report
func NewReport() *Report {

	return &Report{
		ModelVersion: 1,
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
func (o *Report) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Report) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Report) DefaultOrder() []string {

	return []string{}
}

func (o *Report) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Report) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

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
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name contains the name metric of statistics data.`,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		Type:           "string",
	},
	"Tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Tags contains the tags associated to the data point.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "tags_map",
		Type:           "external",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Timestamp contains the unix timestamp for the report.`,
		Exposed:        true,
		Name:           "timestamp",
		SubType:        "int64",
		Type:           "external",
	},
	"Ttl": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `TTL set the time to live for the report in seconds.`,
		Exposed:        true,
		Name:           "ttl",
		Type:           "integer",
	},
	"Value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Value contains the value for the report.`,
		Exposed:        true,
		Name:           "value",
		Type:           "float",
	},
}

// ReportLowerCaseAttributesMap represents the map of attribute for Report.
var ReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name contains the name metric of statistics data.`,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		Type:           "string",
	},
	"tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Tags contains the tags associated to the data point.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "tags_map",
		Type:           "external",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Timestamp contains the unix timestamp for the report.`,
		Exposed:        true,
		Name:           "timestamp",
		SubType:        "int64",
		Type:           "external",
	},
	"ttl": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `TTL set the time to live for the report in seconds.`,
		Exposed:        true,
		Name:           "ttl",
		Type:           "integer",
	},
	"value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Value contains the value for the report.`,
		Exposed:        true,
		Name:           "value",
		Type:           "float",
	},
}
