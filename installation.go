package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
)

// InstallationIdentity represents the Identity of the object.
var InstallationIdentity = elemental.Identity{
	Name:     "installation",
	Category: "installations",
	Package:  "highwind",
	Private:  false,
}

// InstallationsList represents a list of Installations
type InstallationsList []*Installation

// Identity returns the identity of the objects in the list.
func (o InstallationsList) Identity() elemental.Identity {

	return InstallationIdentity
}

// Copy returns a pointer to a copy the InstallationsList.
func (o InstallationsList) Copy() elemental.Identifiables {

	copy := append(InstallationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the InstallationsList.
func (o InstallationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(InstallationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Installation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o InstallationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o InstallationsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the InstallationsList converted to SparseInstallationsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o InstallationsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o InstallationsList) Version() int {

	return 1
}

// Installation represents the model of a installation
type Installation struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// AccountName that should be installed.
	AccountName string `json:"accountName" bson:"accountname" mapstructure:"accountName,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewInstallation returns a new *Installation
func NewInstallation() *Installation {

	return &Installation{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Installation) Identity() elemental.Identity {

	return InstallationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Installation) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Installation) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Installation) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Installation) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Installation) Doc() string {
	return `Installation represents an installation for a given account.`
}

func (o *Installation) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Installation) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseInstallation{
			ID:          &o.ID,
			AccountName: &o.AccountName,
		}
	}

	sp := &SparseInstallation{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "accountName":
			sp.AccountName = &(o.AccountName)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseInstallation to the object.
func (o *Installation) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseInstallation)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.AccountName != nil {
		o.AccountName = *so.AccountName
	}
}

// Validate valides the current information stored into the structure.
func (o *Installation) Validate() error {

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
func (*Installation) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := InstallationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return InstallationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Installation) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return InstallationAttributesMap
}

// InstallationAttributesMap represents the map of attribute for Installation.
var InstallationAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	"AccountName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccountName",
		Description:    `AccountName that should be installed.`,
		Exposed:        true,
		Name:           "accountName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// InstallationLowerCaseAttributesMap represents the map of attribute for Installation.
var InstallationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		Stored:         true,
		Type:           "string",
	},
	"accountname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AccountName",
		Description:    `AccountName that should be installed.`,
		Exposed:        true,
		Name:           "accountName",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseInstallationsList represents a list of SparseInstallations
type SparseInstallationsList []*SparseInstallation

// Identity returns the identity of the objects in the list.
func (o SparseInstallationsList) Identity() elemental.Identity {

	return InstallationIdentity
}

// Copy returns a pointer to a copy the SparseInstallationsList.
func (o SparseInstallationsList) Copy() elemental.Identifiables {

	copy := append(SparseInstallationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseInstallationsList.
func (o SparseInstallationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseInstallationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseInstallation))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseInstallationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseInstallationsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseInstallationsList converted to InstallationsList.
func (o SparseInstallationsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseInstallationsList) Version() int {

	return 1
}

// SparseInstallation represents the sparse version of a installation.
type SparseInstallation struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// AccountName that should be installed.
	AccountName *string `json:"accountName,omitempty" bson:"accountname" mapstructure:"accountName,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseInstallation returns a new  SparseInstallation.
func NewSparseInstallation() *SparseInstallation {
	return &SparseInstallation{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseInstallation) Identity() elemental.Identity {

	return InstallationIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseInstallation) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseInstallation) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseInstallation) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseInstallation) ToPlain() elemental.PlainIdentifiable {

	out := NewInstallation()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.AccountName != nil {
		out.AccountName = *o.AccountName
	}

	return out
}
