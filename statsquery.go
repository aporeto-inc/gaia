package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/types"
)

// StatsQueryIdentity represents the Identity of the object.
var StatsQueryIdentity = elemental.Identity{
	Name:     "statsquery",
	Category: "statsqueries",
	Package:  "jenova",
	Private:  false,
}

// StatsQueriesList represents a list of StatsQueries
type StatsQueriesList []*StatsQuery

// Identity returns the identity of the objects in the list.
func (o StatsQueriesList) Identity() elemental.Identity {

	return StatsQueryIdentity
}

// Copy returns a pointer to a copy the StatsQueriesList.
func (o StatsQueriesList) Copy() elemental.Identifiables {

	copy := append(StatsQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the StatsQueriesList.
func (o StatsQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(StatsQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*StatsQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o StatsQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o StatsQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the StatsQueriesList converted to SparseStatsQueriesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o StatsQueriesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
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

	sync.Mutex `json:"-" bson:"-"`
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
	return `StatsQuery is a generic API to retrieve time series data stored by the Aporeto
system. The API allows different types of queries that are all protected within
the namespace of the user.`
}

func (o *StatsQuery) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *StatsQuery) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseStatsQuery{
			Results: &o.Results,
		}
	}

	sp := &SparseStatsQuery{}
	for _, f := range fields {
		switch f {
		case "results":
			sp.Results = &(o.Results)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseStatsQuery to the object.
func (o *StatsQuery) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseStatsQuery)
	if so.Results != nil {
		o.Results = *so.Results
	}
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

// SparseStatsQueriesList represents a list of SparseStatsQueries
type SparseStatsQueriesList []*SparseStatsQuery

// Identity returns the identity of the objects in the list.
func (o SparseStatsQueriesList) Identity() elemental.Identity {

	return StatsQueryIdentity
}

// Copy returns a pointer to a copy the SparseStatsQueriesList.
func (o SparseStatsQueriesList) Copy() elemental.Identifiables {

	copy := append(SparseStatsQueriesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseStatsQueriesList.
func (o SparseStatsQueriesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseStatsQueriesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseStatsQuery))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseStatsQueriesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseStatsQueriesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseStatsQueriesList converted to StatsQueriesList.
func (o SparseStatsQueriesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseStatsQueriesList) Version() int {

	return 1
}

// SparseStatsQuery represents the sparse version of a statsquery.
type SparseStatsQuery struct {
	// Results contains the result of the query.
	Results *[]*types.TimeSeriesQueryResults `json:"results,omitempty" bson:"-" mapstructure:"results,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseStatsQuery returns a new  SparseStatsQuery.
func NewSparseStatsQuery() *SparseStatsQuery {
	return &SparseStatsQuery{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseStatsQuery) Identity() elemental.Identity {

	return StatsQueryIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseStatsQuery) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseStatsQuery) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseStatsQuery) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseStatsQuery) ToPlain() elemental.PlainIdentifiable {

	out := NewStatsQuery()
	if o.Results != nil {
		out.Results = *o.Results
	}

	return out
}
