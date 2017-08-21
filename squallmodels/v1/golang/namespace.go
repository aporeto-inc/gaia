package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "time"

// NamespaceIdentity represents the Identity of the object
var NamespaceIdentity = elemental.Identity{
	Name:     "namespace",
	Category: "namespaces",
}

// NamespacesList represents a list of Namespaces
type NamespacesList []*Namespace

// ContentIdentity returns the identity of the objects in the list.
func (o NamespacesList) ContentIdentity() elemental.Identity {

	return NamespaceIdentity
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
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id"`

	// Annotation stores additional information about an entity
	Annotations map[string][]string `json:"annotations" bson:"annotations"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// Metadata contains tags that can only be set during creation. They must all start with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata"`

	// Name is the name of the namespace.
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

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
func (o *Namespace) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
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

// GetAssociatedTags returns the associatedTags of the receiver
func (o *Namespace) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *Namespace) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the createTime of the receiver
func (o *Namespace) GetCreateTime() time.Time {
	return o.CreateTime
}

// SetCreateTime set the given createTime of the receiver
func (o *Namespace) SetCreateTime(createTime time.Time) {
	o.CreateTime = createTime
}

// GetMetadata returns the metadata of the receiver
func (o *Namespace) GetMetadata() []string {
	return o.Metadata
}

// SetMetadata set the given metadata of the receiver
func (o *Namespace) SetMetadata(metadata []string) {
	o.Metadata = metadata
}

// GetName returns the name of the receiver
func (o *Namespace) GetName() string {
	return o.Name
}

// GetNamespace returns the namespace of the receiver
func (o *Namespace) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *Namespace) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *Namespace) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *Namespace) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetProtected returns the protected of the receiver
func (o *Namespace) GetProtected() bool {
	return o.Protected
}

// GetUpdateTime returns the updateTime of the receiver
func (o *Namespace) GetUpdateTime() time.Time {
	return o.UpdateTime
}

// SetUpdateTime set the given updateTime of the receiver
func (o *Namespace) SetUpdateTime(updateTime time.Time) {
	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *Namespace) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidatePattern("name", o.Name, `^[^\*\=]*$`, true); err != nil {
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
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotations",
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		AllowedChars:   `^[^\*\=]*$`,
		AllowedChoices: []string{},
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
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotations",
		Stored:         true,
		SubType:        "annotations",
		Type:           "external",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
		AllowedChars:   `^[^\*\=]*$`,
		AllowedChoices: []string{},
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
