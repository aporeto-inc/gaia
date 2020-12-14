package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CNSSearchIdentity represents the Identity of the object.
var CNSSearchIdentity = elemental.Identity{
	Name:     "cnssearch",
	Category: "cnssearches",
	Package:  "karl",
	Private:  false,
}

// CNSSearchesList represents a list of CNSSearches
type CNSSearchesList []*CNSSearch

// Identity returns the identity of the objects in the list.
func (o CNSSearchesList) Identity() elemental.Identity {

	return CNSSearchIdentity
}

// Copy returns a pointer to a copy the CNSSearchesList.
func (o CNSSearchesList) Copy() elemental.Identifiables {

	copy := append(CNSSearchesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CNSSearchesList.
func (o CNSSearchesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CNSSearchesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CNSSearch))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CNSSearchesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CNSSearchesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CNSSearchesList converted to SparseCNSSearchesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CNSSearchesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCNSSearchesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCNSSearch)
	}

	return out
}

// Version returns the version of the content.
func (o CNSSearchesList) Version() int {

	return 1
}

// CNSSearch represents the model of a cnssearch
type CNSSearch struct {
	// The payload of the search results.
	Data *PCSearchResults `json:"data" msgpack:"data" bson:"-" mapstructure:"data,omitempty"`

	// Description of the search.
	Description string `json:"description,omitempty" msgpack:"description,omitempty" bson:"-" mapstructure:"description,omitempty"`

	// ID of the search request.
	Id string `json:"id,omitempty" msgpack:"id,omitempty" bson:"-" mapstructure:"id,omitempty"`

	// The number of items to fetch.
	Limit int `json:"limit,omitempty" msgpack:"limit,omitempty" bson:"-" mapstructure:"limit,omitempty"`

	// Name of the rql search request. Should set to be empty.
	Name string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// Represents the token to fetch next page.
	PageToken string `json:"pageToken,omitempty" msgpack:"pageToken,omitempty" bson:"-" mapstructure:"pageToken,omitempty"`

	// The rql query.
	Query string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// Indicates if the search has been saved.
	Saved bool `json:"saved,omitempty" msgpack:"saved,omitempty" bson:"-" mapstructure:"saved,omitempty"`

	// Type of search request. Should set to be network.
	SearchType string `json:"searchType,omitempty" msgpack:"searchType,omitempty" bson:"-" mapstructure:"searchType,omitempty"`

	// Time range of the search.
	TimeRange *PCTimeRange `json:"timeRange,omitempty" msgpack:"timeRange,omitempty" bson:"-" mapstructure:"timeRange,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCNSSearch returns a new *CNSSearch
func NewCNSSearch() *CNSSearch {

	return &CNSSearch{
		ModelVersion: 1,
		Data:         NewPCSearchResults(),
	}
}

// Identity returns the Identity of the object.
func (o *CNSSearch) Identity() elemental.Identity {

	return CNSSearchIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CNSSearch) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CNSSearch) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSSearch) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCNSSearch{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSSearch) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCNSSearch{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CNSSearch) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CNSSearch) BleveType() string {

	return "cnssearch"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CNSSearch) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CNSSearch) Doc() string {

	return `Provide search results for Primsa Cloud's investigate page.`
}

func (o *CNSSearch) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CNSSearch) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCNSSearch{
			Data:        o.Data,
			Description: &o.Description,
			Id:          &o.Id,
			Limit:       &o.Limit,
			Name:        &o.Name,
			PageToken:   &o.PageToken,
			Query:       &o.Query,
			Saved:       &o.Saved,
			SearchType:  &o.SearchType,
			TimeRange:   o.TimeRange,
		}
	}

	sp := &SparseCNSSearch{}
	for _, f := range fields {
		switch f {
		case "data":
			sp.Data = o.Data
		case "description":
			sp.Description = &(o.Description)
		case "id":
			sp.Id = &(o.Id)
		case "limit":
			sp.Limit = &(o.Limit)
		case "name":
			sp.Name = &(o.Name)
		case "pageToken":
			sp.PageToken = &(o.PageToken)
		case "query":
			sp.Query = &(o.Query)
		case "saved":
			sp.Saved = &(o.Saved)
		case "searchType":
			sp.SearchType = &(o.SearchType)
		case "timeRange":
			sp.TimeRange = o.TimeRange
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCNSSearch to the object.
func (o *CNSSearch) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCNSSearch)
	if so.Data != nil {
		o.Data = so.Data
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Id != nil {
		o.Id = *so.Id
	}
	if so.Limit != nil {
		o.Limit = *so.Limit
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.PageToken != nil {
		o.PageToken = *so.PageToken
	}
	if so.Query != nil {
		o.Query = *so.Query
	}
	if so.Saved != nil {
		o.Saved = *so.Saved
	}
	if so.SearchType != nil {
		o.SearchType = *so.SearchType
	}
	if so.TimeRange != nil {
		o.TimeRange = so.TimeRange
	}
}

// DeepCopy returns a deep copy if the CNSSearch.
func (o *CNSSearch) DeepCopy() *CNSSearch {

	if o == nil {
		return nil
	}

	out := &CNSSearch{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CNSSearch.
func (o *CNSSearch) DeepCopyInto(out *CNSSearch) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CNSSearch: %s", err))
	}

	*out = *target.(*CNSSearch)
}

// Validate valides the current information stored into the structure.
func (o *CNSSearch) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if o.Data != nil {
		elemental.ResetDefaultForZeroValues(o.Data)
		if err := o.Data.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if o.TimeRange != nil {
		elemental.ResetDefaultForZeroValues(o.TimeRange)
		if err := o.TimeRange.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
func (*CNSSearch) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CNSSearchAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CNSSearchLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CNSSearch) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CNSSearchAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CNSSearch) ValueForAttribute(name string) interface{} {

	switch name {
	case "data":
		return o.Data
	case "description":
		return o.Description
	case "id":
		return o.Id
	case "limit":
		return o.Limit
	case "name":
		return o.Name
	case "pageToken":
		return o.PageToken
	case "query":
		return o.Query
	case "saved":
		return o.Saved
	case "searchType":
		return o.SearchType
	case "timeRange":
		return o.TimeRange
	}

	return nil
}

// CNSSearchAttributesMap represents the map of attribute for CNSSearch.
var CNSSearchAttributesMap = map[string]elemental.AttributeSpecification{
	"Data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `The payload of the search results.`,
		Exposed:        true,
		Name:           "data",
		SubType:        "pcsearchresult",
		Type:           "ref",
	},
	"Description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the search.`,
		Exposed:        true,
		Name:           "description",
		Type:           "string",
	},
	"Id": {
		AllowedChoices: []string{},
		ConvertedName:  "Id",
		Description:    `ID of the search request.`,
		Exposed:        true,
		Name:           "id",
		Type:           "string",
	},
	"Limit": {
		AllowedChoices: []string{},
		ConvertedName:  "Limit",
		Description:    `The number of items to fetch.`,
		Exposed:        true,
		Name:           "limit",
		Type:           "integer",
	},
	"Name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the rql search request. Should set to be empty.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"PageToken": {
		AllowedChoices: []string{},
		ConvertedName:  "PageToken",
		Description:    `Represents the token to fetch next page.`,
		Exposed:        true,
		Name:           "pageToken",
		Type:           "string",
	},
	"Query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `The rql query.`,
		Exposed:        true,
		Name:           "query",
		Type:           "string",
	},
	"Saved": {
		AllowedChoices: []string{},
		ConvertedName:  "Saved",
		Description:    `Indicates if the search has been saved.`,
		Exposed:        true,
		Name:           "saved",
		Type:           "boolean",
	},
	"SearchType": {
		AllowedChoices: []string{},
		ConvertedName:  "SearchType",
		Description:    `Type of search request. Should set to be network.`,
		Exposed:        true,
		Name:           "searchType",
		Type:           "string",
	},
	"TimeRange": {
		AllowedChoices: []string{},
		ConvertedName:  "TimeRange",
		Description:    `Time range of the search.`,
		Exposed:        true,
		Name:           "timeRange",
		SubType:        "pctimerange",
		Type:           "ref",
	},
}

// CNSSearchLowerCaseAttributesMap represents the map of attribute for CNSSearch.
var CNSSearchLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"data": {
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `The payload of the search results.`,
		Exposed:        true,
		Name:           "data",
		SubType:        "pcsearchresult",
		Type:           "ref",
	},
	"description": {
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the search.`,
		Exposed:        true,
		Name:           "description",
		Type:           "string",
	},
	"id": {
		AllowedChoices: []string{},
		ConvertedName:  "Id",
		Description:    `ID of the search request.`,
		Exposed:        true,
		Name:           "id",
		Type:           "string",
	},
	"limit": {
		AllowedChoices: []string{},
		ConvertedName:  "Limit",
		Description:    `The number of items to fetch.`,
		Exposed:        true,
		Name:           "limit",
		Type:           "integer",
	},
	"name": {
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the rql search request. Should set to be empty.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"pagetoken": {
		AllowedChoices: []string{},
		ConvertedName:  "PageToken",
		Description:    `Represents the token to fetch next page.`,
		Exposed:        true,
		Name:           "pageToken",
		Type:           "string",
	},
	"query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `The rql query.`,
		Exposed:        true,
		Name:           "query",
		Type:           "string",
	},
	"saved": {
		AllowedChoices: []string{},
		ConvertedName:  "Saved",
		Description:    `Indicates if the search has been saved.`,
		Exposed:        true,
		Name:           "saved",
		Type:           "boolean",
	},
	"searchtype": {
		AllowedChoices: []string{},
		ConvertedName:  "SearchType",
		Description:    `Type of search request. Should set to be network.`,
		Exposed:        true,
		Name:           "searchType",
		Type:           "string",
	},
	"timerange": {
		AllowedChoices: []string{},
		ConvertedName:  "TimeRange",
		Description:    `Time range of the search.`,
		Exposed:        true,
		Name:           "timeRange",
		SubType:        "pctimerange",
		Type:           "ref",
	},
}

// SparseCNSSearchesList represents a list of SparseCNSSearches
type SparseCNSSearchesList []*SparseCNSSearch

// Identity returns the identity of the objects in the list.
func (o SparseCNSSearchesList) Identity() elemental.Identity {

	return CNSSearchIdentity
}

// Copy returns a pointer to a copy the SparseCNSSearchesList.
func (o SparseCNSSearchesList) Copy() elemental.Identifiables {

	copy := append(SparseCNSSearchesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCNSSearchesList.
func (o SparseCNSSearchesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCNSSearchesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCNSSearch))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCNSSearchesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCNSSearchesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCNSSearchesList converted to CNSSearchesList.
func (o SparseCNSSearchesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCNSSearchesList) Version() int {

	return 1
}

// SparseCNSSearch represents the sparse version of a cnssearch.
type SparseCNSSearch struct {
	// The payload of the search results.
	Data *PCSearchResults `json:"data,omitempty" msgpack:"data,omitempty" bson:"-" mapstructure:"data,omitempty"`

	// Description of the search.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"-" mapstructure:"description,omitempty"`

	// ID of the search request.
	Id *string `json:"id,omitempty" msgpack:"id,omitempty" bson:"-" mapstructure:"id,omitempty"`

	// The number of items to fetch.
	Limit *int `json:"limit,omitempty" msgpack:"limit,omitempty" bson:"-" mapstructure:"limit,omitempty"`

	// Name of the rql search request. Should set to be empty.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// Represents the token to fetch next page.
	PageToken *string `json:"pageToken,omitempty" msgpack:"pageToken,omitempty" bson:"-" mapstructure:"pageToken,omitempty"`

	// The rql query.
	Query *string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// Indicates if the search has been saved.
	Saved *bool `json:"saved,omitempty" msgpack:"saved,omitempty" bson:"-" mapstructure:"saved,omitempty"`

	// Type of search request. Should set to be network.
	SearchType *string `json:"searchType,omitempty" msgpack:"searchType,omitempty" bson:"-" mapstructure:"searchType,omitempty"`

	// Time range of the search.
	TimeRange *PCTimeRange `json:"timeRange,omitempty" msgpack:"timeRange,omitempty" bson:"-" mapstructure:"timeRange,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCNSSearch returns a new  SparseCNSSearch.
func NewSparseCNSSearch() *SparseCNSSearch {
	return &SparseCNSSearch{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCNSSearch) Identity() elemental.Identity {

	return CNSSearchIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCNSSearch) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCNSSearch) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSSearch) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCNSSearch{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSSearch) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCNSSearch{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCNSSearch) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCNSSearch) ToPlain() elemental.PlainIdentifiable {

	out := NewCNSSearch()
	if o.Data != nil {
		out.Data = o.Data
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Id != nil {
		out.Id = *o.Id
	}
	if o.Limit != nil {
		out.Limit = *o.Limit
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.PageToken != nil {
		out.PageToken = *o.PageToken
	}
	if o.Query != nil {
		out.Query = *o.Query
	}
	if o.Saved != nil {
		out.Saved = *o.Saved
	}
	if o.SearchType != nil {
		out.SearchType = *o.SearchType
	}
	if o.TimeRange != nil {
		out.TimeRange = o.TimeRange
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCNSSearch.
func (o *SparseCNSSearch) DeepCopy() *SparseCNSSearch {

	if o == nil {
		return nil
	}

	out := &SparseCNSSearch{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCNSSearch.
func (o *SparseCNSSearch) DeepCopyInto(out *SparseCNSSearch) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCNSSearch: %s", err))
	}

	*out = *target.(*SparseCNSSearch)
}

type mongoAttributesCNSSearch struct {
}
type mongoAttributesSparseCNSSearch struct {
}
