package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FileIdentity represents the Identity of the object.
var FileIdentity = elemental.Identity{
	Name:     "file",
	Category: "files",
	Package:  "guy",
	Private:  false,
}

// FilesList represents a list of Files
type FilesList []*File

// Identity returns the identity of the objects in the list.
func (o FilesList) Identity() elemental.Identity {

	return FileIdentity
}

// Copy returns a pointer to a copy the FilesList.
func (o FilesList) Copy() elemental.Identifiables {

	copy := append(FilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FilesList.
func (o FilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*File))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FilesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the FilesList converted to SparseFilesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FilesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseFilesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseFile)
	}

	return out
}

// Version returns the version of the content.
func (o FilesList) Version() int {

	return 1
}

// File represents the model of a file
type File struct {
	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewFile returns a new *File
func NewFile() *File {

	return &File{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *File) Identity() elemental.Identity {

	return FileIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *File) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *File) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *File) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesFile{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *File) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesFile{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *File) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *File) BleveType() string {

	return "file"
}

// DefaultOrder returns the list of default ordering fields.
func (o *File) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *File) Doc() string {

	return `Represents a file that can be uploaded.`
}

func (o *File) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *File) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFile{}
	}

	sp := &SparseFile{}
	for _, f := range fields {
		switch f {
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFile to the object.
func (o *File) Patch(sparse elemental.SparseIdentifiable) {
}

// DeepCopy returns a deep copy if the File.
func (o *File) DeepCopy() *File {

	if o == nil {
		return nil
	}

	out := &File{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *File.
func (o *File) DeepCopyInto(out *File) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy File: %s", err))
	}

	*out = *target.(*File)
}

// Validate valides the current information stored into the structure.
func (o *File) Validate() error {

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
func (*File) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FileAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FileLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*File) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FileAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *File) ValueForAttribute(name string) interface{} {

	switch name {
	}

	return nil
}

// FileAttributesMap represents the map of attribute for File.
var FileAttributesMap = map[string]elemental.AttributeSpecification{}

// FileLowerCaseAttributesMap represents the map of attribute for File.
var FileLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{}

// SparseFilesList represents a list of SparseFiles
type SparseFilesList []*SparseFile

// Identity returns the identity of the objects in the list.
func (o SparseFilesList) Identity() elemental.Identity {

	return FileIdentity
}

// Copy returns a pointer to a copy the SparseFilesList.
func (o SparseFilesList) Copy() elemental.Identifiables {

	copy := append(SparseFilesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFilesList.
func (o SparseFilesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFilesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFile))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFilesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFilesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseFilesList converted to FilesList.
func (o SparseFilesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFilesList) Version() int {

	return 1
}

// SparseFile represents the sparse version of a file.
type SparseFile struct {
	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseFile returns a new  SparseFile.
func NewSparseFile() *SparseFile {
	return &SparseFile{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFile) Identity() elemental.Identity {

	return FileIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFile) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFile) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFile) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseFile{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseFile) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseFile{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseFile) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFile) ToPlain() elemental.PlainIdentifiable {

	out := NewFile()

	return out
}

// DeepCopy returns a deep copy if the SparseFile.
func (o *SparseFile) DeepCopy() *SparseFile {

	if o == nil {
		return nil
	}

	out := &SparseFile{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseFile.
func (o *SparseFile) DeepCopyInto(out *SparseFile) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseFile: %s", err))
	}

	*out = *target.(*SparseFile)
}

type mongoAttributesFile struct {
}
type mongoAttributesSparseFile struct {
}
