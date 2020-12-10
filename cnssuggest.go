package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CNSSuggestIdentity represents the Identity of the object.
var CNSSuggestIdentity = elemental.Identity{
	Name:     "cnssuggest",
	Category: "cnssuggests",
	Package:  "karl",
	Private:  false,
}

// CNSSuggestsList represents a list of CNSSuggests
type CNSSuggestsList []*CNSSuggest

// Identity returns the identity of the objects in the list.
func (o CNSSuggestsList) Identity() elemental.Identity {

	return CNSSuggestIdentity
}

// Copy returns a pointer to a copy the CNSSuggestsList.
func (o CNSSuggestsList) Copy() elemental.Identifiables {

	copy := append(CNSSuggestsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CNSSuggestsList.
func (o CNSSuggestsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CNSSuggestsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CNSSuggest))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CNSSuggestsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CNSSuggestsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CNSSuggestsList converted to SparseCNSSuggestsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CNSSuggestsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCNSSuggestsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCNSSuggest)
	}

	return out
}

// Version returns the version of the content.
func (o CNSSuggestsList) Version() int {

	return 1
}

// CNSSuggest represents the model of a cnssuggest
type CNSSuggest struct {
	// Rquired by Prisma Cloud. Always set to true.
	NeedsOffsetUpdate bool `json:"needsOffsetUpdate" msgpack:"needsOffsetUpdate" bson:"-" mapstructure:"needsOffsetUpdate,omitempty"`

	// The length of the rql query part that is valid.
	Offset int `json:"offset" msgpack:"offset" bson:"-" mapstructure:"offset,omitempty"`

	// Prisma Cloud's rql query.
	Query string `json:"query" msgpack:"query" bson:"-" mapstructure:"query,omitempty"`

	// List of query suggestions.
	Suggestions []string `json:"suggestions" msgpack:"suggestions" bson:"-" mapstructure:"suggestions,omitempty"`

	// Rquired by Prisma Cloud. Always set to false.
	Translate bool `json:"translate" msgpack:"translate" bson:"-" mapstructure:"translate,omitempty"`

	// The validity of the rql query.
	Valid bool `json:"valid" msgpack:"valid" bson:"-" mapstructure:"valid,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCNSSuggest returns a new *CNSSuggest
func NewCNSSuggest() *CNSSuggest {

	return &CNSSuggest{
		ModelVersion:      1,
		NeedsOffsetUpdate: true,
		Offset:            0,
		Suggestions:       []string{},
		Translate:         false,
		Valid:             false,
	}
}

// Identity returns the Identity of the object.
func (o *CNSSuggest) Identity() elemental.Identity {

	return CNSSuggestIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CNSSuggest) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CNSSuggest) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSSuggest) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCNSSuggest{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSSuggest) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCNSSuggest{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CNSSuggest) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CNSSuggest) BleveType() string {

	return "cnssuggest"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CNSSuggest) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CNSSuggest) Doc() string {

	return `Provide search suggestion query for Primsa Cloud's investigate page.`
}

func (o *CNSSuggest) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CNSSuggest) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCNSSuggest{
			NeedsOffsetUpdate: &o.NeedsOffsetUpdate,
			Offset:            &o.Offset,
			Query:             &o.Query,
			Suggestions:       &o.Suggestions,
			Translate:         &o.Translate,
			Valid:             &o.Valid,
		}
	}

	sp := &SparseCNSSuggest{}
	for _, f := range fields {
		switch f {
		case "needsOffsetUpdate":
			sp.NeedsOffsetUpdate = &(o.NeedsOffsetUpdate)
		case "offset":
			sp.Offset = &(o.Offset)
		case "query":
			sp.Query = &(o.Query)
		case "suggestions":
			sp.Suggestions = &(o.Suggestions)
		case "translate":
			sp.Translate = &(o.Translate)
		case "valid":
			sp.Valid = &(o.Valid)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCNSSuggest to the object.
func (o *CNSSuggest) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCNSSuggest)
	if so.NeedsOffsetUpdate != nil {
		o.NeedsOffsetUpdate = *so.NeedsOffsetUpdate
	}
	if so.Offset != nil {
		o.Offset = *so.Offset
	}
	if so.Query != nil {
		o.Query = *so.Query
	}
	if so.Suggestions != nil {
		o.Suggestions = *so.Suggestions
	}
	if so.Translate != nil {
		o.Translate = *so.Translate
	}
	if so.Valid != nil {
		o.Valid = *so.Valid
	}
}

// DeepCopy returns a deep copy if the CNSSuggest.
func (o *CNSSuggest) DeepCopy() *CNSSuggest {

	if o == nil {
		return nil
	}

	out := &CNSSuggest{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CNSSuggest.
func (o *CNSSuggest) DeepCopyInto(out *CNSSuggest) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CNSSuggest: %s", err))
	}

	*out = *target.(*CNSSuggest)
}

// Validate valides the current information stored into the structure.
func (o *CNSSuggest) Validate() error {

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
func (*CNSSuggest) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CNSSuggestAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CNSSuggestLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CNSSuggest) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CNSSuggestAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CNSSuggest) ValueForAttribute(name string) interface{} {

	switch name {
	case "needsOffsetUpdate":
		return o.NeedsOffsetUpdate
	case "offset":
		return o.Offset
	case "query":
		return o.Query
	case "suggestions":
		return o.Suggestions
	case "translate":
		return o.Translate
	case "valid":
		return o.Valid
	}

	return nil
}

// CNSSuggestAttributesMap represents the map of attribute for CNSSuggest.
var CNSSuggestAttributesMap = map[string]elemental.AttributeSpecification{
	"NeedsOffsetUpdate": {
		AllowedChoices: []string{},
		ConvertedName:  "NeedsOffsetUpdate",
		DefaultValue:   true,
		Description:    `Rquired by Prisma Cloud. Always set to true.`,
		Exposed:        true,
		Name:           "needsOffsetUpdate",
		Type:           "boolean",
	},
	"Offset": {
		AllowedChoices: []string{},
		ConvertedName:  "Offset",
		Description:    `The length of the rql query part that is valid.`,
		Exposed:        true,
		Name:           "offset",
		Type:           "integer",
	},
	"Query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `Prisma Cloud's rql query.`,
		Exposed:        true,
		Name:           "query",
		ReadOnly:       true,
		Type:           "string",
	},
	"Suggestions": {
		AllowedChoices: []string{},
		ConvertedName:  "Suggestions",
		Description:    `List of query suggestions.`,
		Exposed:        true,
		Name:           "suggestions",
		SubType:        "string",
		Type:           "list",
	},
	"Translate": {
		AllowedChoices: []string{},
		ConvertedName:  "Translate",
		Description:    `Rquired by Prisma Cloud. Always set to false.`,
		Exposed:        true,
		Name:           "translate",
		Type:           "boolean",
	},
	"Valid": {
		AllowedChoices: []string{},
		ConvertedName:  "Valid",
		Description:    `The validity of the rql query.`,
		Exposed:        true,
		Name:           "valid",
		Type:           "boolean",
	},
}

// CNSSuggestLowerCaseAttributesMap represents the map of attribute for CNSSuggest.
var CNSSuggestLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"needsoffsetupdate": {
		AllowedChoices: []string{},
		ConvertedName:  "NeedsOffsetUpdate",
		DefaultValue:   true,
		Description:    `Rquired by Prisma Cloud. Always set to true.`,
		Exposed:        true,
		Name:           "needsOffsetUpdate",
		Type:           "boolean",
	},
	"offset": {
		AllowedChoices: []string{},
		ConvertedName:  "Offset",
		Description:    `The length of the rql query part that is valid.`,
		Exposed:        true,
		Name:           "offset",
		Type:           "integer",
	},
	"query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `Prisma Cloud's rql query.`,
		Exposed:        true,
		Name:           "query",
		ReadOnly:       true,
		Type:           "string",
	},
	"suggestions": {
		AllowedChoices: []string{},
		ConvertedName:  "Suggestions",
		Description:    `List of query suggestions.`,
		Exposed:        true,
		Name:           "suggestions",
		SubType:        "string",
		Type:           "list",
	},
	"translate": {
		AllowedChoices: []string{},
		ConvertedName:  "Translate",
		Description:    `Rquired by Prisma Cloud. Always set to false.`,
		Exposed:        true,
		Name:           "translate",
		Type:           "boolean",
	},
	"valid": {
		AllowedChoices: []string{},
		ConvertedName:  "Valid",
		Description:    `The validity of the rql query.`,
		Exposed:        true,
		Name:           "valid",
		Type:           "boolean",
	},
}

// SparseCNSSuggestsList represents a list of SparseCNSSuggests
type SparseCNSSuggestsList []*SparseCNSSuggest

// Identity returns the identity of the objects in the list.
func (o SparseCNSSuggestsList) Identity() elemental.Identity {

	return CNSSuggestIdentity
}

// Copy returns a pointer to a copy the SparseCNSSuggestsList.
func (o SparseCNSSuggestsList) Copy() elemental.Identifiables {

	copy := append(SparseCNSSuggestsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCNSSuggestsList.
func (o SparseCNSSuggestsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCNSSuggestsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCNSSuggest))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCNSSuggestsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCNSSuggestsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCNSSuggestsList converted to CNSSuggestsList.
func (o SparseCNSSuggestsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCNSSuggestsList) Version() int {

	return 1
}

// SparseCNSSuggest represents the sparse version of a cnssuggest.
type SparseCNSSuggest struct {
	// Rquired by Prisma Cloud. Always set to true.
	NeedsOffsetUpdate *bool `json:"needsOffsetUpdate,omitempty" msgpack:"needsOffsetUpdate,omitempty" bson:"-" mapstructure:"needsOffsetUpdate,omitempty"`

	// The length of the rql query part that is valid.
	Offset *int `json:"offset,omitempty" msgpack:"offset,omitempty" bson:"-" mapstructure:"offset,omitempty"`

	// Prisma Cloud's rql query.
	Query *string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// List of query suggestions.
	Suggestions *[]string `json:"suggestions,omitempty" msgpack:"suggestions,omitempty" bson:"-" mapstructure:"suggestions,omitempty"`

	// Rquired by Prisma Cloud. Always set to false.
	Translate *bool `json:"translate,omitempty" msgpack:"translate,omitempty" bson:"-" mapstructure:"translate,omitempty"`

	// The validity of the rql query.
	Valid *bool `json:"valid,omitempty" msgpack:"valid,omitempty" bson:"-" mapstructure:"valid,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCNSSuggest returns a new  SparseCNSSuggest.
func NewSparseCNSSuggest() *SparseCNSSuggest {
	return &SparseCNSSuggest{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCNSSuggest) Identity() elemental.Identity {

	return CNSSuggestIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCNSSuggest) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCNSSuggest) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSSuggest) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCNSSuggest{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSSuggest) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCNSSuggest{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCNSSuggest) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCNSSuggest) ToPlain() elemental.PlainIdentifiable {

	out := NewCNSSuggest()
	if o.NeedsOffsetUpdate != nil {
		out.NeedsOffsetUpdate = *o.NeedsOffsetUpdate
	}
	if o.Offset != nil {
		out.Offset = *o.Offset
	}
	if o.Query != nil {
		out.Query = *o.Query
	}
	if o.Suggestions != nil {
		out.Suggestions = *o.Suggestions
	}
	if o.Translate != nil {
		out.Translate = *o.Translate
	}
	if o.Valid != nil {
		out.Valid = *o.Valid
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCNSSuggest.
func (o *SparseCNSSuggest) DeepCopy() *SparseCNSSuggest {

	if o == nil {
		return nil
	}

	out := &SparseCNSSuggest{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCNSSuggest.
func (o *SparseCNSSuggest) DeepCopyInto(out *SparseCNSSuggest) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCNSSuggest: %s", err))
	}

	*out = *target.(*SparseCNSSuggest)
}

type mongoAttributesCNSSuggest struct {
}
type mongoAttributesSparseCNSSuggest struct {
}
