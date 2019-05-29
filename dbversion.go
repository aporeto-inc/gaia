package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// DBVersionIdentity represents the Identity of the object.
var DBVersionIdentity = elemental.Identity{
	Name:     "dbversion",
	Category: "dbversions",
	Package:  "squall",
	Private:  false,
}

// DBVersionsList represents a list of DBVersions
type DBVersionsList []*DBVersion

// Identity returns the identity of the objects in the list.
func (o DBVersionsList) Identity() elemental.Identity {

	return DBVersionIdentity
}

// Copy returns a pointer to a copy the DBVersionsList.
func (o DBVersionsList) Copy() elemental.Identifiables {

	copy := append(DBVersionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DBVersionsList.
func (o DBVersionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(DBVersionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*DBVersion))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o DBVersionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DBVersionsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the DBVersionsList converted to SparseDBVersionsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o DBVersionsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseDBVersionsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseDBVersion)
	}

	return out
}

// Version returns the version of the content.
func (o DBVersionsList) Version() int {

	return 1
}

// DBVersion represents the model of a dbversion
type DBVersion struct {
	// Name of the service.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	// Version of the database.
	VersionNumber int `json:"versionNumber" msgpack:"versionNumber" bson:"-" mapstructure:"versionNumber,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewDBVersion returns a new *DBVersion
func NewDBVersion() *DBVersion {

	return &DBVersion{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *DBVersion) Identity() elemental.Identity {

	return DBVersionIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DBVersion) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DBVersion) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *DBVersion) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *DBVersion) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *DBVersion) Doc() string {

	return `A DBVersion shows the versions of each database.`
}

func (o *DBVersion) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *DBVersion) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseDBVersion{
			Name:          &o.Name,
			VersionNumber: &o.VersionNumber,
		}
	}

	sp := &SparseDBVersion{}
	for _, f := range fields {
		switch f {
		case "name":
			sp.Name = &(o.Name)
		case "versionNumber":
			sp.VersionNumber = &(o.VersionNumber)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseDBVersion to the object.
func (o *DBVersion) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseDBVersion)
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.VersionNumber != nil {
		o.VersionNumber = *so.VersionNumber
	}
}

// DeepCopy returns a deep copy if the DBVersion.
func (o *DBVersion) DeepCopy() *DBVersion {

	if o == nil {
		return nil
	}

	out := &DBVersion{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *DBVersion.
func (o *DBVersion) DeepCopyInto(out *DBVersion) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy DBVersion: %s", err))
	}

	*out = *target.(*DBVersion)
}

// Validate valides the current information stored into the structure.
func (o *DBVersion) Validate() error {

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
func (*DBVersion) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DBVersionAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DBVersionLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DBVersion) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DBVersionAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *DBVersion) ValueForAttribute(name string) interface{} {

	switch name {
	case "name":
		return o.Name
	case "versionNumber":
		return o.VersionNumber
	}

	return nil
}

// DBVersionAttributesMap represents the map of attribute for DBVersion.
var DBVersionAttributesMap = map[string]elemental.AttributeSpecification{
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the service.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"VersionNumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "VersionNumber",
		Description:    `Version of the database.`,
		Exposed:        true,
		Name:           "versionNumber",
		Type:           "integer",
	},
}

// DBVersionLowerCaseAttributesMap represents the map of attribute for DBVersion.
var DBVersionLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the service.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"versionnumber": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "VersionNumber",
		Description:    `Version of the database.`,
		Exposed:        true,
		Name:           "versionNumber",
		Type:           "integer",
	},
}

// SparseDBVersionsList represents a list of SparseDBVersions
type SparseDBVersionsList []*SparseDBVersion

// Identity returns the identity of the objects in the list.
func (o SparseDBVersionsList) Identity() elemental.Identity {

	return DBVersionIdentity
}

// Copy returns a pointer to a copy the SparseDBVersionsList.
func (o SparseDBVersionsList) Copy() elemental.Identifiables {

	copy := append(SparseDBVersionsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseDBVersionsList.
func (o SparseDBVersionsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseDBVersionsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseDBVersion))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseDBVersionsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseDBVersionsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseDBVersionsList converted to DBVersionsList.
func (o SparseDBVersionsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseDBVersionsList) Version() int {

	return 1
}

// SparseDBVersion represents the sparse version of a dbversion.
type SparseDBVersion struct {
	// Name of the service.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// Version of the database.
	VersionNumber *int `json:"versionNumber,omitempty" msgpack:"versionNumber,omitempty" bson:"-" mapstructure:"versionNumber,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseDBVersion returns a new  SparseDBVersion.
func NewSparseDBVersion() *SparseDBVersion {
	return &SparseDBVersion{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseDBVersion) Identity() elemental.Identity {

	return DBVersionIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseDBVersion) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseDBVersion) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseDBVersion) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseDBVersion) ToPlain() elemental.PlainIdentifiable {

	out := NewDBVersion()
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.VersionNumber != nil {
		out.VersionNumber = *o.VersionNumber
	}

	return out
}

// DeepCopy returns a deep copy if the SparseDBVersion.
func (o *SparseDBVersion) DeepCopy() *SparseDBVersion {

	if o == nil {
		return nil
	}

	out := &SparseDBVersion{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseDBVersion.
func (o *SparseDBVersion) DeepCopyInto(out *SparseDBVersion) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseDBVersion: %s", err))
	}

	*out = *target.(*SparseDBVersion)
}
