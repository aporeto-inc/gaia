package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// PUReportIndexes lists the attribute compound indexes.
var PUReportIndexes = [][]string{}

// PUReportIdentity represents the Identity of the object.
var PUReportIdentity = elemental.Identity{
	Name:     "pureport",
	Category: "pureports",
	Private:  false,
}

// PUReportsList represents a list of PUReports
type PUReportsList []*PUReport

// Identity returns the identity of the objects in the list.
func (o PUReportsList) Identity() elemental.Identity {

	return PUReportIdentity
}

// Copy returns a pointer to a copy the PUReportsList.
func (o PUReportsList) Copy() elemental.Identifiables {

	copy := append(PUReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PUReportsList.
func (o PUReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PUReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PUReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PUReportsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PUReportsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o PUReportsList) Version() int {

	return 1
}

// PUReport represents the model of a pureport
type PUReport struct {
	// PU to report.
	PU *ProcessingUnit `json:"PU" bson:"-" mapstructure:"PU,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewPUReport returns a new *PUReport
func NewPUReport() *PUReport {

	return &PUReport{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PUReport) Identity() elemental.Identity {

	return PUReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PUReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PUReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *PUReport) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *PUReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PUReport) Doc() string {
	return `Post a new pu statistics report.`
}

func (o *PUReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *PUReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("PU", o.PU); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredTime("timestamp", o.Timestamp); err != nil {
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
func (*PUReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PUReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PUReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PUReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PUReportAttributesMap
}

// PUReportAttributesMap represents the map of attribute for PUReport.
var PUReportAttributesMap = map[string]elemental.AttributeSpecification{
	"PU": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PU",
		Description:    `PU to report.`,
		Exposed:        true,
		Name:           "PU",
		Required:       true,
		SubType:        "processingunit",
		Type:           "external",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
}

// PUReportLowerCaseAttributesMap represents the map of attribute for PUReport.
var PUReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"pu": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PU",
		Description:    `PU to report.`,
		Exposed:        true,
		Name:           "PU",
		Required:       true,
		SubType:        "processingunit",
		Type:           "external",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
}
