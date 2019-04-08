package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AccountCheckIdentity represents the Identity of the object.
var AccountCheckIdentity = elemental.Identity{
	Name:     "accountcheck",
	Category: "accountchecks",
	Package:  "vince",
	Private:  true,
}

// AccountChecksList represents a list of AccountChecks
type AccountChecksList []*AccountCheck

// Identity returns the identity of the objects in the list.
func (o AccountChecksList) Identity() elemental.Identity {

	return AccountCheckIdentity
}

// Copy returns a pointer to a copy the AccountChecksList.
func (o AccountChecksList) Copy() elemental.Identifiables {

	copy := append(AccountChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AccountChecksList.
func (o AccountChecksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AccountChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AccountCheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AccountChecksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AccountChecksList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AccountChecksList converted to SparseAccountChecksList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AccountChecksList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o AccountChecksList) Version() int {

	return 1
}

// AccountCheck represents the model of a accountcheck
type AccountCheck struct {
	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewAccountCheck returns a new *AccountCheck
func NewAccountCheck() *AccountCheck {

	return &AccountCheck{
		ModelVersion: 1,
		Mutex:        &sync.Mutex{},
	}
}

// Identity returns the Identity of the object.
func (o *AccountCheck) Identity() elemental.Identity {

	return AccountCheckIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AccountCheck) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AccountCheck) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *AccountCheck) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *AccountCheck) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AccountCheck) Doc() string {

	return `Validates the password for an account.`
}

func (o *AccountCheck) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AccountCheck) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAccountCheck{}
	}

	sp := &SparseAccountCheck{}
	for _, f := range fields {
		switch f {
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAccountCheck to the object.
func (o *AccountCheck) Patch(sparse elemental.SparseIdentifiable) {
}

// DeepCopy returns a deep copy if the AccountCheck.
func (o *AccountCheck) DeepCopy() *AccountCheck {

	if o == nil {
		return nil
	}

	out := &AccountCheck{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AccountCheck.
func (o *AccountCheck) DeepCopyInto(out *AccountCheck) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AccountCheck: %s", err))
	}

	*out = *target.(*AccountCheck)
}

// Validate valides the current information stored into the structure.
func (o *AccountCheck) Validate() error {

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
func (*AccountCheck) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AccountCheckAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AccountCheckLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AccountCheck) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AccountCheckAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AccountCheck) ValueForAttribute(name string) interface{} {

	switch name {
	}

	return nil
}

// AccountCheckAttributesMap represents the map of attribute for AccountCheck.
var AccountCheckAttributesMap = map[string]elemental.AttributeSpecification{}

// AccountCheckLowerCaseAttributesMap represents the map of attribute for AccountCheck.
var AccountCheckLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{}

// SparseAccountChecksList represents a list of SparseAccountChecks
type SparseAccountChecksList []*SparseAccountCheck

// Identity returns the identity of the objects in the list.
func (o SparseAccountChecksList) Identity() elemental.Identity {

	return AccountCheckIdentity
}

// Copy returns a pointer to a copy the SparseAccountChecksList.
func (o SparseAccountChecksList) Copy() elemental.Identifiables {

	copy := append(SparseAccountChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAccountChecksList.
func (o SparseAccountChecksList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAccountChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAccountCheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAccountChecksList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAccountChecksList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAccountChecksList converted to AccountChecksList.
func (o SparseAccountChecksList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAccountChecksList) Version() int {

	return 1
}

// SparseAccountCheck represents the sparse version of a accountcheck.
type SparseAccountCheck struct {
	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparseAccountCheck returns a new  SparseAccountCheck.
func NewSparseAccountCheck() *SparseAccountCheck {
	return &SparseAccountCheck{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAccountCheck) Identity() elemental.Identity {

	return AccountCheckIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAccountCheck) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAccountCheck) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseAccountCheck) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAccountCheck) ToPlain() elemental.PlainIdentifiable {

	out := NewAccountCheck()

	return out
}

// DeepCopy returns a deep copy if the SparseAccountCheck.
func (o *SparseAccountCheck) DeepCopy() *SparseAccountCheck {

	if o == nil {
		return nil
	}

	out := &SparseAccountCheck{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAccountCheck.
func (o *SparseAccountCheck) DeepCopyInto(out *SparseAccountCheck) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAccountCheck: %s", err))
	}

	*out = *target.(*SparseAccountCheck)
}
