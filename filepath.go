package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// FilePathIdentity represents the Identity of the object.
var FilePathIdentity = elemental.Identity{
	Name:     "filepath",
	Category: "filepaths",
	Package:  "squall",
	Private:  false,
}

// FilePathsList represents a list of FilePaths
type FilePathsList []*FilePath

// Identity returns the identity of the objects in the list.
func (o FilePathsList) Identity() elemental.Identity {

	return FilePathIdentity
}

// Copy returns a pointer to a copy the FilePathsList.
func (o FilePathsList) Copy() elemental.Identifiables {

	copy := append(FilePathsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FilePathsList.
func (o FilePathsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FilePathsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FilePath))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FilePathsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FilePathsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the FilePathsList converted to SparseFilePathsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FilePathsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o FilePathsList) Version() int {

	return 1
}

// FilePath represents the model of a filepath
type FilePath struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// FilePath refer to the file mount path.
	Filepath string `json:"filepath" bson:"filepath" mapstructure:"filepath,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// server is the server name/ID/IP associated with the file path.
	Server string `json:"server" bson:"server" mapstructure:"server,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewFilePath returns a new *FilePath
func NewFilePath() *FilePath {

	return &FilePath{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Metadata:       []string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *FilePath) Identity() elemental.Identity {

	return FilePathIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FilePath) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FilePath) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *FilePath) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *FilePath) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *FilePath) Doc() string {
	return `A File Path represents a random path to a file or a folder. They can be used in
aFile Access Policiesin order to allow Processing Units to access them, using
various modes (read, write, execute). You will need to use the File Paths tags
to set some policies. A good example would bevolume=web or file=/etc/passwd.`
}

func (o *FilePath) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *FilePath) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *FilePath) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *FilePath) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *FilePath) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *FilePath) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *FilePath) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *FilePath) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *FilePath) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *FilePath) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *FilePath) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *FilePath) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *FilePath) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *FilePath) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *FilePath) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *FilePath) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *FilePath) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *FilePath) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FilePath) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFilePath{
			ID:             &o.ID,
			Annotations:    &o.Annotations,
			AssociatedTags: &o.AssociatedTags,
			CreateTime:     &o.CreateTime,
			Description:    &o.Description,
			Filepath:       &o.Filepath,
			Metadata:       &o.Metadata,
			Name:           &o.Name,
			Namespace:      &o.Namespace,
			NormalizedTags: &o.NormalizedTags,
			Protected:      &o.Protected,
			Server:         &o.Server,
			UpdateTime:     &o.UpdateTime,
		}
	}

	sp := &SparseFilePath{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "filepath":
			sp.Filepath = &(o.Filepath)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "server":
			sp.Server = &(o.Server)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseFilePath to the object.
func (o *FilePath) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFilePath)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Filepath != nil {
		o.Filepath = *so.Filepath
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Server != nil {
		o.Server = *so.Server
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// Validate valides the current information stored into the structure.
func (o *FilePath) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("filepath", o.Filepath); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
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
func (*FilePath) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FilePathAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FilePathLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FilePath) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FilePathAttributesMap
}

// FilePathAttributesMap represents the map of attribute for FilePath.
var FilePathAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Filepath": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Filepath",
		Description:    `FilePath refer to the file mount path.`,
		Exposed:        true,
		Name:           "filepath",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "metadata_list",
		Type:       "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Server": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Server",
		CreationOnly:   true,
		Description:    `server is the server name/ID/IP associated with the file path.`,
		Exposed:        true,
		Name:           "server",
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// FilePathLowerCaseAttributesMap represents the map of attribute for FilePath.
var FilePathLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `CreatedTime is the time at which the object was created.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"filepath": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Filepath",
		Description:    `FilePath refer to the file mount path.`,
		Exposed:        true,
		Name:           "filepath",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "metadata_list",
		Type:       "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"server": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Server",
		CreationOnly:   true,
		Description:    `server is the server name/ID/IP associated with the file path.`,
		Exposed:        true,
		Name:           "server",
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// SparseFilePathsList represents a list of SparseFilePaths
type SparseFilePathsList []*SparseFilePath

// Identity returns the identity of the objects in the list.
func (o SparseFilePathsList) Identity() elemental.Identity {

	return FilePathIdentity
}

// Copy returns a pointer to a copy the SparseFilePathsList.
func (o SparseFilePathsList) Copy() elemental.Identifiables {

	copy := append(SparseFilePathsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFilePathsList.
func (o SparseFilePathsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFilePathsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFilePath))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFilePathsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFilePathsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToFull returns the SparseFilePathsList converted to FilePathsList.
func (o SparseFilePathsList) ToFull() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToFull()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFilePathsList) Version() int {

	return 1
}

// SparseFilePath represents the sparse version of a filepath.
type SparseFilePath struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// FilePath refer to the file mount path.
	Filepath *string `json:"filepath,omitempty" bson:"filepath" mapstructure:"filepath,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// server is the server name/ID/IP associated with the file path.
	Server *string `json:"server,omitempty" bson:"server" mapstructure:"server,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseFilePath returns a new  SparseFilePath.
func NewSparseFilePath() *SparseFilePath {
	return &SparseFilePath{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFilePath) Identity() elemental.Identity {

	return FilePathIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFilePath) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFilePath) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseFilePath) Version() int {

	return 1
}

// ToFull returns a full version of the sparse model.
func (o *SparseFilePath) ToFull() elemental.FullIdentifiable {

	out := NewFilePath()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Filepath != nil {
		out.Filepath = *o.Filepath
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Server != nil {
		out.Server = *o.Server
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}
