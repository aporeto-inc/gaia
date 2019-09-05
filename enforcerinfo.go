package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerInfoIdentity represents the Identity of the object.
var EnforcerInfoIdentity = elemental.Identity{
	Name:     "enforcerinfo",
	Category: "enforcerinfo",
	Package:  "squall",
	Private:  false,
}

// EnforcerInfosList represents a list of EnforcerInfos
type EnforcerInfosList []*EnforcerInfo

// Identity returns the identity of the objects in the list.
func (o EnforcerInfosList) Identity() elemental.Identity {

	return EnforcerInfoIdentity
}

// Copy returns a pointer to a copy the EnforcerInfosList.
func (o EnforcerInfosList) Copy() elemental.Identifiables {

	copy := append(EnforcerInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerInfosList.
func (o EnforcerInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerInfosList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the EnforcerInfosList converted to SparseEnforcerInfosList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcerInfosList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseEnforcerInfosList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseEnforcerInfo)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcerInfosList) Version() int {

	return 1
}

// EnforcerInfo represents the model of a enforcerinfo
type EnforcerInfo struct {
	// Represents the latest information collected by the enforcer.
	CollectedInfo map[string]string `json:"collectedInfo" msgpack:"collectedInfo" bson:"collectedinfo" mapstructure:"collectedInfo,omitempty"`

	// ID of the enforcer.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"-" mapstructure:"enforcerID,omitempty"`

	// EnforcerInfoID is the ID of the enforcer information. EnforcerInfoID is used to
	// aggergate the multipart requests.
	EnforcerInfoID string `json:"enforcerInfoID" msgpack:"enforcerInfoID" bson:"-" mapstructure:"enforcerInfoID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEnforcerInfo returns a new *EnforcerInfo
func NewEnforcerInfo() *EnforcerInfo {

	return &EnforcerInfo{
		ModelVersion:  1,
		CollectedInfo: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerInfo) Identity() elemental.Identity {

	return EnforcerInfoIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerInfo) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *EnforcerInfo) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *EnforcerInfo) BleveType() string {

	return "enforcerinfo"
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerInfo) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *EnforcerInfo) Doc() string {

	return `Post a new enforcer information.`
}

func (o *EnforcerInfo) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EnforcerInfo) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcerInfo{
			CollectedInfo:  &o.CollectedInfo,
			EnforcerID:     &o.EnforcerID,
			EnforcerInfoID: &o.EnforcerInfoID,
		}
	}

	sp := &SparseEnforcerInfo{}
	for _, f := range fields {
		switch f {
		case "collectedInfo":
			sp.CollectedInfo = &(o.CollectedInfo)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerInfoID":
			sp.EnforcerInfoID = &(o.EnforcerInfoID)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEnforcerInfo to the object.
func (o *EnforcerInfo) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcerInfo)
	if so.CollectedInfo != nil {
		o.CollectedInfo = *so.CollectedInfo
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerInfoID != nil {
		o.EnforcerInfoID = *so.EnforcerInfoID
	}
}

// DeepCopy returns a deep copy if the EnforcerInfo.
func (o *EnforcerInfo) DeepCopy() *EnforcerInfo {

	if o == nil {
		return nil
	}

	out := &EnforcerInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EnforcerInfo.
func (o *EnforcerInfo) DeepCopyInto(out *EnforcerInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EnforcerInfo: %s", err))
	}

	*out = *target.(*EnforcerInfo)
}

// Validate valides the current information stored into the structure.
func (o *EnforcerInfo) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerInfoID", o.EnforcerInfoID); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*EnforcerInfo) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerInfoAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerInfoLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerInfo) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerInfoAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EnforcerInfo) ValueForAttribute(name string) interface{} {

	switch name {
	case "collectedInfo":
		return o.CollectedInfo
	case "enforcerID":
		return o.EnforcerID
	case "enforcerInfoID":
		return o.EnforcerInfoID
	}

	return nil
}

// EnforcerInfoAttributesMap represents the map of attribute for EnforcerInfo.
var EnforcerInfoAttributesMap = map[string]elemental.AttributeSpecification{
	"CollectedInfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectedInfo",
		Description:    `Represents the latest information collected by the enforcer.`,
		Exposed:        true,
		Name:           "collectedInfo",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"EnforcerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Type:           "string",
	},
	"EnforcerInfoID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerInfoID",
		Description: `EnforcerInfoID is the ID of the enforcer information. EnforcerInfoID is used to
aggergate the multipart requests.`,
		Exposed:  true,
		Name:     "enforcerInfoID",
		Required: true,
		Type:     "string",
	},
}

// EnforcerInfoLowerCaseAttributesMap represents the map of attribute for EnforcerInfo.
var EnforcerInfoLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"collectedinfo": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectedInfo",
		Description:    `Represents the latest information collected by the enforcer.`,
		Exposed:        true,
		Name:           "collectedInfo",
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"enforcerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Type:           "string",
	},
	"enforcerinfoid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerInfoID",
		Description: `EnforcerInfoID is the ID of the enforcer information. EnforcerInfoID is used to
aggergate the multipart requests.`,
		Exposed:  true,
		Name:     "enforcerInfoID",
		Required: true,
		Type:     "string",
	},
}

// SparseEnforcerInfosList represents a list of SparseEnforcerInfos
type SparseEnforcerInfosList []*SparseEnforcerInfo

// Identity returns the identity of the objects in the list.
func (o SparseEnforcerInfosList) Identity() elemental.Identity {

	return EnforcerInfoIdentity
}

// Copy returns a pointer to a copy the SparseEnforcerInfosList.
func (o SparseEnforcerInfosList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcerInfosList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcerInfosList.
func (o SparseEnforcerInfosList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcerInfosList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcerInfo))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcerInfosList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcerInfosList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseEnforcerInfosList converted to EnforcerInfosList.
func (o SparseEnforcerInfosList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcerInfosList) Version() int {

	return 1
}

// SparseEnforcerInfo represents the sparse version of a enforcerinfo.
type SparseEnforcerInfo struct {
	// Represents the latest information collected by the enforcer.
	CollectedInfo *map[string]string `json:"collectedInfo,omitempty" msgpack:"collectedInfo,omitempty" bson:"collectedinfo,omitempty" mapstructure:"collectedInfo,omitempty"`

	// ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"-" mapstructure:"enforcerID,omitempty"`

	// EnforcerInfoID is the ID of the enforcer information. EnforcerInfoID is used to
	// aggergate the multipart requests.
	EnforcerInfoID *string `json:"enforcerInfoID,omitempty" msgpack:"enforcerInfoID,omitempty" bson:"-" mapstructure:"enforcerInfoID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseEnforcerInfo returns a new  SparseEnforcerInfo.
func NewSparseEnforcerInfo() *SparseEnforcerInfo {
	return &SparseEnforcerInfo{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcerInfo) Identity() elemental.Identity {

	return EnforcerInfoIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcerInfo) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcerInfo) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseEnforcerInfo) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcerInfo) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcerInfo()
	if o.CollectedInfo != nil {
		out.CollectedInfo = *o.CollectedInfo
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerInfoID != nil {
		out.EnforcerInfoID = *o.EnforcerInfoID
	}

	return out
}

// DeepCopy returns a deep copy if the SparseEnforcerInfo.
func (o *SparseEnforcerInfo) DeepCopy() *SparseEnforcerInfo {

	if o == nil {
		return nil
	}

	out := &SparseEnforcerInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcerInfo.
func (o *SparseEnforcerInfo) DeepCopyInto(out *SparseEnforcerInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcerInfo: %s", err))
	}

	*out = *target.(*SparseEnforcerInfo)
}
