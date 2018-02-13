package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// NamespaceIdentity represents the Identity of the object.
var NamespaceIdentity = elemental.Identity{
	Name:     "namespace",
	Category: "namespaces",
	Private:  false,
}

// NamespacesList represents a list of Namespaces
type NamespacesList []*Namespace

// ContentIdentity returns the identity of the objects in the list.
func (o NamespacesList) ContentIdentity() elemental.Identity {

	return NamespaceIdentity
}

// Copy returns a pointer to a copy the NamespacesList.
func (o NamespacesList) Copy() elemental.ContentIdentifiable {

	copy := append(NamespacesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NamespacesList.
func (o NamespacesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(NamespacesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Namespace))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NamespacesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NamespacesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o NamespacesList) Version() int {

	return 1
}

// Namespace represents the model of a namespace
type Namespace struct {
	// AssociatedLocalCAID holds the remote ID of the certificate authority to use.
	AssociatedLocalCAID string `json:"-" bson:"associatedlocalcaid" mapstructure:"-,omitempty"`

	// LocalCA holds the eventual certificate authority used by this namespace.
	LocalCA string `json:"localCA" bson:"localca" mapstructure:"localCA,omitempty"`

	// LocalCAEnabled defines if the namespace should use a local Certificate Authority. Switching it off and on again will regenerate a new CA.
	LocalCAEnabled bool `json:"localCAEnabled" bson:"localcaenabled" mapstructure:"localCAEnabled,omitempty"`

	// Name is the name of the namespace.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Annotation stores additional information about an entity
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewNamespace returns a new *Namespace
func NewNamespace() *Namespace {

	return &Namespace{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Metadata:       []string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *Namespace) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Namespace) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Namespace) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Namespace) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Namespace) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Namespace) Doc() string {
	return `A Namespace represents the core organizational unit of the system. All objects always exists in a single namespace. A Namespace can also have child namespaces. They can be used to split the system into organizations, business units, applications, services or any combination you like.`
}

func (o *Namespace) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the Name of the receiver.
func (o *Namespace) GetName() string {

	return o.Name
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Namespace) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *Namespace) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Namespace) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *Namespace) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Namespace) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *Namespace) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *Namespace) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *Namespace) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Namespace) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *Namespace) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *Namespace) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Namespace) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *Namespace) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *Namespace) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *Namespace) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// Validate valides the current information stored into the structure.
func (o *Namespace) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidatePattern("name", o.Name, `^[a-zA-Z0-9-_/]+$`, true); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
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
func (*Namespace) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := NamespaceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return NamespaceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Namespace) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespaceAttributesMap
}

// NamespaceAttributesMap represents the map of attribute for Namespace.
var NamespaceAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"AssociatedLocalCAID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedLocalCAID",
		Description:    `AssociatedLocalCAID holds the remote ID of the certificate authority to use.`,
		Format:         "free",
		Name:           "associatedLocalCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity`,
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
		Description:    `CreatedTime is the time at which the object was created`,
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
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"LocalCA": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description:    `LocalCA holds the eventual certificate authority used by this namespace.`,
		Exposed:        true,
		Format:         "free",
		Name:           "localCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"LocalCAEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LocalCAEnabled",
		Description:    `LocalCAEnabled defines if the namespace should use a local Certificate Authority. Switching it off and on again will regenerate a new CA.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "localCAEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description:    `Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "metadata",
		Setter:         true,
		Stored:         true,
		SubType:        "metadata_list",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChars:   `^[a-zA-Z0-9-_/]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `Name is the name of the namespace.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Index:          true,
		Name:           "name",
		Orderable:      true,
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Index:          true,
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
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
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
		Filterable:     true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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

// NamespaceLowerCaseAttributesMap represents the map of attribute for Namespace.
var NamespaceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"associatedlocalcaid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedLocalCAID",
		Description:    `AssociatedLocalCAID holds the remote ID of the certificate authority to use.`,
		Format:         "free",
		Name:           "associatedLocalCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity`,
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
		Description:    `CreatedTime is the time at which the object was created`,
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
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"localca": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description:    `LocalCA holds the eventual certificate authority used by this namespace.`,
		Exposed:        true,
		Format:         "free",
		Name:           "localCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"localcaenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LocalCAEnabled",
		Description:    `LocalCAEnabled defines if the namespace should use a local Certificate Authority. Switching it off and on again will regenerate a new CA.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "localCAEnabled",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description:    `Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "metadata",
		Setter:         true,
		Stored:         true,
		SubType:        "metadata_list",
		Type:           "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChars:   `^[a-zA-Z0-9-_/]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		Description:    `Name is the name of the namespace.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Index:          true,
		Name:           "name",
		Orderable:      true,
		PrimaryKey:     true,
		Required:       true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Index:          true,
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
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
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
		Filterable:     true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
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
