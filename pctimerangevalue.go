package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PCTimeRangeValueIdentity represents the Identity of the object.
var PCTimeRangeValueIdentity = elemental.Identity{
	Name:     "pctimerangevalue",
	Category: "pctimerangevalues",
	Package:  "karl",
	Private:  false,
}

// PCTimeRangeValuesList represents a list of PCTimeRangeValues
type PCTimeRangeValuesList []*PCTimeRangeValue

// Identity returns the identity of the objects in the list.
func (o PCTimeRangeValuesList) Identity() elemental.Identity {

	return PCTimeRangeValueIdentity
}

// Copy returns a pointer to a copy the PCTimeRangeValuesList.
func (o PCTimeRangeValuesList) Copy() elemental.Identifiables {

	copy := append(PCTimeRangeValuesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PCTimeRangeValuesList.
func (o PCTimeRangeValuesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PCTimeRangeValuesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PCTimeRangeValue))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PCTimeRangeValuesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PCTimeRangeValuesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PCTimeRangeValuesList converted to SparsePCTimeRangeValuesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PCTimeRangeValuesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePCTimeRangeValuesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePCTimeRangeValue)
	}

	return out
}

// Version returns the version of the content.
func (o PCTimeRangeValuesList) Version() int {

	return 1
}

// PCTimeRangeValue represents the model of a pctimerangevalue
type PCTimeRangeValue struct {
	// The count of time durations.
	Amount int `json:"amount" msgpack:"amount" bson:"-" mapstructure:"amount,omitempty"`

	// the end time of the search, in Unix time format.
	EndTime int `json:"endTime" msgpack:"endTime" bson:"-" mapstructure:"endTime,omitempty"`

	// The start time of the search, in Unix time format.
	StartTime int `json:"startTime" msgpack:"startTime" bson:"-" mapstructure:"startTime,omitempty"`

	// The unit of the time durations.
	Unit string `json:"unit" msgpack:"unit" bson:"-" mapstructure:"unit,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPCTimeRangeValue returns a new *PCTimeRangeValue
func NewPCTimeRangeValue() *PCTimeRangeValue {

	return &PCTimeRangeValue{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PCTimeRangeValue) Identity() elemental.Identity {

	return PCTimeRangeValueIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PCTimeRangeValue) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PCTimeRangeValue) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PCTimeRangeValue) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPCTimeRangeValue{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PCTimeRangeValue) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPCTimeRangeValue{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PCTimeRangeValue) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PCTimeRangeValue) BleveType() string {

	return "pctimerangevalue"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PCTimeRangeValue) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PCTimeRangeValue) Doc() string {

	return `Represents the time range value of rql search request.`
}

func (o *PCTimeRangeValue) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PCTimeRangeValue) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePCTimeRangeValue{
			Amount:    &o.Amount,
			EndTime:   &o.EndTime,
			StartTime: &o.StartTime,
			Unit:      &o.Unit,
		}
	}

	sp := &SparsePCTimeRangeValue{}
	for _, f := range fields {
		switch f {
		case "amount":
			sp.Amount = &(o.Amount)
		case "endTime":
			sp.EndTime = &(o.EndTime)
		case "startTime":
			sp.StartTime = &(o.StartTime)
		case "unit":
			sp.Unit = &(o.Unit)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePCTimeRangeValue to the object.
func (o *PCTimeRangeValue) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePCTimeRangeValue)
	if so.Amount != nil {
		o.Amount = *so.Amount
	}
	if so.EndTime != nil {
		o.EndTime = *so.EndTime
	}
	if so.StartTime != nil {
		o.StartTime = *so.StartTime
	}
	if so.Unit != nil {
		o.Unit = *so.Unit
	}
}

// DeepCopy returns a deep copy if the PCTimeRangeValue.
func (o *PCTimeRangeValue) DeepCopy() *PCTimeRangeValue {

	if o == nil {
		return nil
	}

	out := &PCTimeRangeValue{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PCTimeRangeValue.
func (o *PCTimeRangeValue) DeepCopyInto(out *PCTimeRangeValue) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PCTimeRangeValue: %s", err))
	}

	*out = *target.(*PCTimeRangeValue)
}

// Validate valides the current information stored into the structure.
func (o *PCTimeRangeValue) Validate() error {

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
func (*PCTimeRangeValue) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PCTimeRangeValueAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PCTimeRangeValueLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PCTimeRangeValue) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PCTimeRangeValueAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PCTimeRangeValue) ValueForAttribute(name string) interface{} {

	switch name {
	case "amount":
		return o.Amount
	case "endTime":
		return o.EndTime
	case "startTime":
		return o.StartTime
	case "unit":
		return o.Unit
	}

	return nil
}

// PCTimeRangeValueAttributesMap represents the map of attribute for PCTimeRangeValue.
var PCTimeRangeValueAttributesMap = map[string]elemental.AttributeSpecification{
	"Amount": {
		AllowedChoices: []string{},
		ConvertedName:  "Amount",
		Description:    `The count of time durations.`,
		Exposed:        true,
		Name:           "amount",
		Type:           "integer",
	},
	"EndTime": {
		AllowedChoices: []string{},
		ConvertedName:  "EndTime",
		Description:    `the end time of the search, in Unix time format.`,
		Exposed:        true,
		Name:           "endTime",
		Type:           "integer",
	},
	"StartTime": {
		AllowedChoices: []string{},
		ConvertedName:  "StartTime",
		Description:    `The start time of the search, in Unix time format.`,
		Exposed:        true,
		Name:           "startTime",
		Type:           "integer",
	},
	"Unit": {
		AllowedChoices: []string{},
		ConvertedName:  "Unit",
		Description:    `The unit of the time durations.`,
		Exposed:        true,
		Name:           "unit",
		Type:           "string",
	},
}

// PCTimeRangeValueLowerCaseAttributesMap represents the map of attribute for PCTimeRangeValue.
var PCTimeRangeValueLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"amount": {
		AllowedChoices: []string{},
		ConvertedName:  "Amount",
		Description:    `The count of time durations.`,
		Exposed:        true,
		Name:           "amount",
		Type:           "integer",
	},
	"endtime": {
		AllowedChoices: []string{},
		ConvertedName:  "EndTime",
		Description:    `the end time of the search, in Unix time format.`,
		Exposed:        true,
		Name:           "endTime",
		Type:           "integer",
	},
	"starttime": {
		AllowedChoices: []string{},
		ConvertedName:  "StartTime",
		Description:    `The start time of the search, in Unix time format.`,
		Exposed:        true,
		Name:           "startTime",
		Type:           "integer",
	},
	"unit": {
		AllowedChoices: []string{},
		ConvertedName:  "Unit",
		Description:    `The unit of the time durations.`,
		Exposed:        true,
		Name:           "unit",
		Type:           "string",
	},
}

// SparsePCTimeRangeValuesList represents a list of SparsePCTimeRangeValues
type SparsePCTimeRangeValuesList []*SparsePCTimeRangeValue

// Identity returns the identity of the objects in the list.
func (o SparsePCTimeRangeValuesList) Identity() elemental.Identity {

	return PCTimeRangeValueIdentity
}

// Copy returns a pointer to a copy the SparsePCTimeRangeValuesList.
func (o SparsePCTimeRangeValuesList) Copy() elemental.Identifiables {

	copy := append(SparsePCTimeRangeValuesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePCTimeRangeValuesList.
func (o SparsePCTimeRangeValuesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePCTimeRangeValuesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePCTimeRangeValue))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePCTimeRangeValuesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePCTimeRangeValuesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePCTimeRangeValuesList converted to PCTimeRangeValuesList.
func (o SparsePCTimeRangeValuesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePCTimeRangeValuesList) Version() int {

	return 1
}

// SparsePCTimeRangeValue represents the sparse version of a pctimerangevalue.
type SparsePCTimeRangeValue struct {
	// The count of time durations.
	Amount *int `json:"amount,omitempty" msgpack:"amount,omitempty" bson:"-" mapstructure:"amount,omitempty"`

	// the end time of the search, in Unix time format.
	EndTime *int `json:"endTime,omitempty" msgpack:"endTime,omitempty" bson:"-" mapstructure:"endTime,omitempty"`

	// The start time of the search, in Unix time format.
	StartTime *int `json:"startTime,omitempty" msgpack:"startTime,omitempty" bson:"-" mapstructure:"startTime,omitempty"`

	// The unit of the time durations.
	Unit *string `json:"unit,omitempty" msgpack:"unit,omitempty" bson:"-" mapstructure:"unit,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePCTimeRangeValue returns a new  SparsePCTimeRangeValue.
func NewSparsePCTimeRangeValue() *SparsePCTimeRangeValue {
	return &SparsePCTimeRangeValue{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePCTimeRangeValue) Identity() elemental.Identity {

	return PCTimeRangeValueIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePCTimeRangeValue) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePCTimeRangeValue) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePCTimeRangeValue) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePCTimeRangeValue{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePCTimeRangeValue) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePCTimeRangeValue{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePCTimeRangeValue) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePCTimeRangeValue) ToPlain() elemental.PlainIdentifiable {

	out := NewPCTimeRangeValue()
	if o.Amount != nil {
		out.Amount = *o.Amount
	}
	if o.EndTime != nil {
		out.EndTime = *o.EndTime
	}
	if o.StartTime != nil {
		out.StartTime = *o.StartTime
	}
	if o.Unit != nil {
		out.Unit = *o.Unit
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePCTimeRangeValue.
func (o *SparsePCTimeRangeValue) DeepCopy() *SparsePCTimeRangeValue {

	if o == nil {
		return nil
	}

	out := &SparsePCTimeRangeValue{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePCTimeRangeValue.
func (o *SparsePCTimeRangeValue) DeepCopyInto(out *SparsePCTimeRangeValue) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePCTimeRangeValue: %s", err))
	}

	*out = *target.(*SparsePCTimeRangeValue)
}

type mongoAttributesPCTimeRangeValue struct {
}
type mongoAttributesSparsePCTimeRangeValue struct {
}
