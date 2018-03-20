package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// FlowStatisticMetricValue represents the possible values for attribute "metric".
type FlowStatisticMetricValue string

const (
	// FlowStatisticMetricFlows represents the value Flows.
	FlowStatisticMetricFlows FlowStatisticMetricValue = "Flows"

	// FlowStatisticMetricPorts represents the value Ports.
	FlowStatisticMetricPorts FlowStatisticMetricValue = "Ports"
)

// FlowStatisticModeValue represents the possible values for attribute "mode".
type FlowStatisticModeValue string

const (
	// FlowStatisticModeAccepted represents the value Accepted.
	FlowStatisticModeAccepted FlowStatisticModeValue = "Accepted"

	// FlowStatisticModeAny represents the value Any.
	FlowStatisticModeAny FlowStatisticModeValue = "Any"

	// FlowStatisticModeRejected represents the value Rejected.
	FlowStatisticModeRejected FlowStatisticModeValue = "Rejected"
)

// FlowStatisticTypeValue represents the possible values for attribute "type".
type FlowStatisticTypeValue string

const (
	// FlowStatisticTypeRepartition represents the value Repartition.
	FlowStatisticTypeRepartition FlowStatisticTypeValue = "Repartition"

	// FlowStatisticTypeSerie represents the value Serie.
	FlowStatisticTypeSerie FlowStatisticTypeValue = "Serie"
)

// FlowStatisticIdentity represents the Identity of the object.
var FlowStatisticIdentity = elemental.Identity{
	Name:     "flowstatistic",
	Category: "flowstatistics",
	Private:  false,
}

// FlowStatisticsList represents a list of FlowStatistics
type FlowStatisticsList []*FlowStatistic

// ContentIdentity returns the identity of the objects in the list.
func (o FlowStatisticsList) ContentIdentity() elemental.Identity {

	return FlowStatisticIdentity
}

// Copy returns a pointer to a copy the FlowStatisticsList.
func (o FlowStatisticsList) Copy() elemental.ContentIdentifiable {

	copy := append(FlowStatisticsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FlowStatisticsList.
func (o FlowStatisticsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(FlowStatisticsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FlowStatistic))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FlowStatisticsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FlowStatisticsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o FlowStatisticsList) Version() int {

	return 1
}

// FlowStatistic represents the model of a flowstatistic
type FlowStatistic struct {
	// DataPoints is a list of time/value pairs that represent the flow events over
	// time.
	DataPoints []map[string]interface{} `json:"dataPoints" bson:"-" mapstructure:"dataPoints,omitempty"`

	// DestinationIDs is the IDs of the destination.
	DestinationIDs []string `json:"destinationIDs" bson:"-" mapstructure:"destinationIDs,omitempty"`

	// DestinationTags contains the tags used to identify destination
	DestinationTags map[string]string `json:"destinationTags" bson:"-" mapstructure:"destinationTags,omitempty"`

	// Metric is the kind of metric the statistic represents.
	Metric FlowStatisticMetricValue `json:"metric" bson:"-" mapstructure:"metric,omitempty"`

	// Mode defines if the metric is for accepted or rejected flows.
	Mode FlowStatisticModeValue `json:"mode" bson:"-" mapstructure:"mode,omitempty"`

	// SourceIDs is the sources of the stats.
	SourceIDs []string `json:"sourceIDs" bson:"-" mapstructure:"sourceIDs,omitempty"`

	// SourceTags contains the tags used to identify the source.
	SourceTags map[string]string `json:"sourceTags" bson:"-" mapstructure:"sourceTags,omitempty"`

	// Type is the type of representation
	Type FlowStatisticTypeValue `json:"type" bson:"-" mapstructure:"type,omitempty"`

	// UserIdentifier can be set by the user as a query parameter. It will be returned
	// in the FlowStatistic object.
	UserIdentifier string `json:"userIdentifier" bson:"-" mapstructure:"userIdentifier,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewFlowStatistic returns a new *FlowStatistic
func NewFlowStatistic() *FlowStatistic {

	return &FlowStatistic{
		ModelVersion:   1,
		DataPoints:     []map[string]interface{}{},
		DestinationIDs: []string{},
		Metric:         "Flows",
		Mode:           "Accepted",
		SourceIDs:      []string{},
		Type:           "Serie",
	}
}

// Identity returns the Identity of the object.
func (o *FlowStatistic) Identity() elemental.Identity {

	return FlowStatisticIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FlowStatistic) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FlowStatistic) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *FlowStatistic) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *FlowStatistic) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *FlowStatistic) Doc() string {
	return `Returns network access statistics on a particular processing unit or group of
processing units based on their tags.`
}

func (o *FlowStatistic) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *FlowStatistic) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("metric", string(o.Metric), []string{"Flows", "Ports"}, true); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("mode", string(o.Mode), []string{"Accepted", "Any", "Rejected"}, true); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Repartition", "Serie"}, true); err != nil {
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
func (*FlowStatistic) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FlowStatisticAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FlowStatisticLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FlowStatistic) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FlowStatisticAttributesMap
}

// FlowStatisticAttributesMap represents the map of attribute for FlowStatistic.
var FlowStatisticAttributesMap = map[string]elemental.AttributeSpecification{
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
		ReadOnly:       true,
		Type:           "string",
	},
	"DataPoints": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "DataPoints",
		Description: `DataPoints is a list of time/value pairs that represent the flow events over
time.`,
		Exposed:  true,
		Name:     "dataPoints",
		ReadOnly: true,
		SubType:  "datapoints_list",
		Type:     "external",
	},
	"DestinationIDs": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "DestinationIDs",
		Description:    `DestinationIDs is the IDs of the destination.`,
		Exposed:        true,
		Format:         "free",
		Name:           "destinationIDs",
		ReadOnly:       true,
		SubType:        "flowstatistic_origin_list",
		Type:           "external",
	},
	"DestinationTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "DestinationTags",
		Description:    `DestinationTags contains the tags used to identify destination`,
		Exposed:        true,
		Name:           "destinationTags",
		ReadOnly:       true,
		SubType:        "selectors_list",
		Type:           "external",
	},
	"Metric": elemental.AttributeSpecification{
		AllowedChoices: []string{"Flows", "Ports"},
		Autogenerated:  true,
		ConvertedName:  "Metric",
		DefaultValue:   FlowStatisticMetricFlows,
		Description:    `Metric is the kind of metric the statistic represents.`,
		Exposed:        true,
		Name:           "metric",
		ReadOnly:       true,
		Type:           "enum",
	},
	"Mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accepted", "Any", "Rejected"},
		Autogenerated:  true,
		ConvertedName:  "Mode",
		DefaultValue:   FlowStatisticModeAccepted,
		Description:    `Mode defines if the metric is for accepted or rejected flows.`,
		Exposed:        true,
		Name:           "mode",
		ReadOnly:       true,
		Type:           "enum",
	},
	"SourceIDs": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SourceIDs",
		Description:    `SourceIDs is the sources of the stats.`,
		Exposed:        true,
		Format:         "free",
		Name:           "sourceIDs",
		ReadOnly:       true,
		SubType:        "flowstatistic_origin_list",
		Type:           "external",
	},
	"SourceTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SourceTags",
		Description:    `SourceTags contains the tags used to identify the source.`,
		Exposed:        true,
		Name:           "sourceTags",
		ReadOnly:       true,
		SubType:        "selectors_list",
		Type:           "external",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Repartition", "Serie"},
		Autogenerated:  true,
		ConvertedName:  "Type",
		DefaultValue:   FlowStatisticTypeSerie,
		Description:    `Type is the type of representation`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "enum",
	},
	"UserIdentifier": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UserIdentifier",
		Description: `UserIdentifier can be set by the user as a query parameter. It will be returned
in the FlowStatistic object.`,
		Exposed:   true,
		Format:    "free",
		Name:      "userIdentifier",
		Orderable: true,
		Type:      "string",
	},
}

// FlowStatisticLowerCaseAttributesMap represents the map of attribute for FlowStatistic.
var FlowStatisticLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		ReadOnly:       true,
		Type:           "string",
	},
	"datapoints": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "DataPoints",
		Description: `DataPoints is a list of time/value pairs that represent the flow events over
time.`,
		Exposed:  true,
		Name:     "dataPoints",
		ReadOnly: true,
		SubType:  "datapoints_list",
		Type:     "external",
	},
	"destinationids": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "DestinationIDs",
		Description:    `DestinationIDs is the IDs of the destination.`,
		Exposed:        true,
		Format:         "free",
		Name:           "destinationIDs",
		ReadOnly:       true,
		SubType:        "flowstatistic_origin_list",
		Type:           "external",
	},
	"destinationtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "DestinationTags",
		Description:    `DestinationTags contains the tags used to identify destination`,
		Exposed:        true,
		Name:           "destinationTags",
		ReadOnly:       true,
		SubType:        "selectors_list",
		Type:           "external",
	},
	"metric": elemental.AttributeSpecification{
		AllowedChoices: []string{"Flows", "Ports"},
		Autogenerated:  true,
		ConvertedName:  "Metric",
		DefaultValue:   FlowStatisticMetricFlows,
		Description:    `Metric is the kind of metric the statistic represents.`,
		Exposed:        true,
		Name:           "metric",
		ReadOnly:       true,
		Type:           "enum",
	},
	"mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accepted", "Any", "Rejected"},
		Autogenerated:  true,
		ConvertedName:  "Mode",
		DefaultValue:   FlowStatisticModeAccepted,
		Description:    `Mode defines if the metric is for accepted or rejected flows.`,
		Exposed:        true,
		Name:           "mode",
		ReadOnly:       true,
		Type:           "enum",
	},
	"sourceids": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SourceIDs",
		Description:    `SourceIDs is the sources of the stats.`,
		Exposed:        true,
		Format:         "free",
		Name:           "sourceIDs",
		ReadOnly:       true,
		SubType:        "flowstatistic_origin_list",
		Type:           "external",
	},
	"sourcetags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "SourceTags",
		Description:    `SourceTags contains the tags used to identify the source.`,
		Exposed:        true,
		Name:           "sourceTags",
		ReadOnly:       true,
		SubType:        "selectors_list",
		Type:           "external",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Repartition", "Serie"},
		Autogenerated:  true,
		ConvertedName:  "Type",
		DefaultValue:   FlowStatisticTypeSerie,
		Description:    `Type is the type of representation`,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "enum",
	},
	"useridentifier": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UserIdentifier",
		Description: `UserIdentifier can be set by the user as a query parameter. It will be returned
in the FlowStatistic object.`,
		Exposed:   true,
		Format:    "free",
		Name:      "userIdentifier",
		Orderable: true,
		Type:      "string",
	},
}
