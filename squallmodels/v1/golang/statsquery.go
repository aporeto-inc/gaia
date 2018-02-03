package squallmodels

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/squallmodels/v1/golang/types"
)

// StatsQueryIdentity represents the Identity of the object.
var StatsQueryIdentity = elemental.Identity{
	Name:     "statsquery",
	Category: "statsqueries",
}

// StatsQueriesList represents a list of StatsQueries
type StatsQueriesList []*StatsQuery

// ContentIdentity returns the identity of the objects in the list.
func (o StatsQueriesList) ContentIdentity() elemental.Identity {

	return StatsQueryIdentity
}

// Copy returns a pointer to a copy the StatsQueriesList.
func (o StatsQueriesList) Copy() elemental.ContentIdentifiable {

	copy := append(StatsQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the StatsQueriesList.
func (o StatsQueriesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(StatsQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*StatsQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o StatsQueriesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o StatsQueriesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o StatsQueriesList) Version() int {

	return 1
}

// StatsQuery represents the model of a statsquery
type StatsQuery struct {
	// Results contains the result of the query.
	Results []*types.TimeSeriesQueryResults `json:"results" bson:"-" mapstructure:"results,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewStatsQuery returns a new *StatsQuery
func NewStatsQuery() *StatsQuery {

	return &StatsQuery{
		ModelVersion: 1,
		Results:      []*types.TimeSeriesQueryResults{},
	}
}

// Identity returns the Identity of the object.
func (o *StatsQuery) Identity() elemental.Identity {

	return StatsQueryIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *StatsQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *StatsQuery) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *StatsQuery) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *StatsQuery) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *StatsQuery) Doc() string {
	return `TimeSeriesData is a generic API to retrieve time series data stored by the Aporeto system. The API allows different types of queries that are all protected within the namespace of the user.`
}

func (o *StatsQuery) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *StatsQuery) Validate() error {

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
func (*StatsQuery) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := StatsQueryAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return StatsQueryLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*StatsQuery) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return StatsQueryAttributesMap
}

// StatsQueryAttributesMap represents the map of attribute for StatsQuery.
var StatsQueryAttributesMap = map[string]elemental.AttributeSpecification{
	"Results": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Results",
		Description:    `Results contains the result of the query.`,
		Exposed:        true,
		Name:           "results",
		ReadOnly:       true,
		SubType:        "time_series_results",
		Type:           "external",
	},
}

// StatsQueryLowerCaseAttributesMap represents the map of attribute for StatsQuery.
var StatsQueryLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"results": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Results",
		Description:    `Results contains the result of the query.`,
		Exposed:        true,
		Name:           "results",
		ReadOnly:       true,
		SubType:        "time_series_results",
		Type:           "external",
	},
}
