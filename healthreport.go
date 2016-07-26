// Copyright 2016 Aporeto Inc.

package gaia

import "github.com/aporeto-inc/elemental"

const (
	HealthReportAttributeNameAPIVersion       elemental.AttributeSpecificationNameKey = "healthreport/APIVersion"
	HealthReportAttributeNameID               elemental.AttributeSpecificationNameKey = "healthreport/ID"
	HealthReportAttributeNameBahamutVersion   elemental.AttributeSpecificationNameKey = "healthreport/bahamutVersion"
	HealthReportAttributeNameElementalVersion elemental.AttributeSpecificationNameKey = "healthreport/elementalVersion"
	HealthReportAttributeNameGaiaVersion      elemental.AttributeSpecificationNameKey = "healthreport/gaiaVersion"
	HealthReportAttributeNameSquallVersion    elemental.AttributeSpecificationNameKey = "healthreport/squallVersion"
	HealthReportAttributeNameStatus           elemental.AttributeSpecificationNameKey = "healthreport/status"
)

// HealthReportStatusValue represents the possible values for attribute "status".
type HealthReportStatusValue string

const (
	HealthReportStatusDegraded HealthReportStatusValue = "Degraded"
	HealthReportStatusFailure  HealthReportStatusValue = "Failure"
	HealthReportStatusOk       HealthReportStatusValue = "Ok"
)

// HealthReportIdentity represents the Identity of the object
var HealthReportIdentity = elemental.Identity{
	Name:     "healthreport",
	Category: "healthreports",
}

// HealthReportsList represents a list of HealthReports
type HealthReportsList []*HealthReport

// HealthReport represents the model of a healthreport
type HealthReport struct {
	APIVersion       string                  `json:"APIVersion,omitempty" cql:"apiversion,omitempty"`
	ID               string                  `json:"ID,omitempty" cql:"id,omitempty"`
	BahamutVersion   string                  `json:"bahamutVersion,omitempty" cql:"bahamutversion,omitempty"`
	ElementalVersion string                  `json:"elementalVersion,omitempty" cql:"elementalversion,omitempty"`
	GaiaVersion      string                  `json:"gaiaVersion,omitempty" cql:"gaiaversion,omitempty"`
	SquallVersion    string                  `json:"squallVersion,omitempty" cql:"squallversion,omitempty"`
	Status           HealthReportStatusValue `json:"status,omitempty" cql:"status,omitempty"`
}

// NewHealthReport returns a new *HealthReport
func NewHealthReport() *HealthReport {

	return &HealthReport{}
}

// Identity returns the Identity of the object.
func (o *HealthReport) Identity() elemental.Identity {

	return HealthReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *HealthReport) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *HealthReport) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *HealthReport) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("status", string(o.Status), []string{"Degraded", "Failure", "Ok"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o HealthReport) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return HealthReportAttributesMap[name]
}

var HealthReportAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	HealthReportAttributeNameAPIVersion: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "APIVersion",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	HealthReportAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	HealthReportAttributeNameBahamutVersion: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "bahamutVersion",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	HealthReportAttributeNameElementalVersion: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "elementalVersion",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	HealthReportAttributeNameGaiaVersion: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "gaiaVersion",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	HealthReportAttributeNameSquallVersion: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "squallVersion",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	HealthReportAttributeNameStatus: elemental.AttributeSpecification{
		AllowedChoices: []string{"Degraded", "Failure", "Ok"},
		Autogenerated:  true,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
}
