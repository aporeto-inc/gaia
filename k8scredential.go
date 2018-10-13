package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// K8SCredentialIdentity represents the Identity of the object.
var K8SCredentialIdentity = elemental.Identity{
	Name:     "k8scredential",
	Category: "k8scredentials",
	Package:  "cactuar",
	Private:  false,
}

// K8SCredentialsList represents a list of K8SCredentials
type K8SCredentialsList []*K8SCredential

// Identity returns the identity of the objects in the list.
func (o K8SCredentialsList) Identity() elemental.Identity {

	return K8SCredentialIdentity
}

// Copy returns a pointer to a copy the K8SCredentialsList.
func (o K8SCredentialsList) Copy() elemental.Identifiables {

	copy := append(K8SCredentialsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the K8SCredentialsList.
func (o K8SCredentialsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(K8SCredentialsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*K8SCredential))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o K8SCredentialsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o K8SCredentialsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o K8SCredentialsList) Version() int {

	return 1
}

// K8SCredential represents the model of a k8scredential
type K8SCredential struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The string representation of the Certificate used by the Kubernetes cluster.
	Certificate string `json:"certificate" bson:"certificate" mapstructure:"certificate,omitempty"`

	// Link to the certificate created for this cluster.
	CertificateSN string `json:"-" bson:"certificatesn" mapstructure:"-,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// The email address that will receive a copy of the Kubernetes cluster YAMLs
	// definition.
	Email string `json:"email" bson:"email" mapstructure:"email,omitempty"`

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

	// Regenerates the k8s files and certificates.
	Regenerate bool `json:"regenerate" bson:"-" mapstructure:"regenerate,omitempty"`

	// The secret file to deploy on your cluster.
	SecretDefinition map[string]interface{} `json:"secretDefinition" bson:"-" mapstructure:"secretDefinition,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewK8SCredential returns a new *K8SCredential
func NewK8SCredential() *K8SCredential {

	return &K8SCredential{
		ModelVersion:     1,
		Annotations:      map[string][]string{},
		AssociatedTags:   []string{},
		Metadata:         []string{},
		NormalizedTags:   []string{},
		SecretDefinition: map[string]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *K8SCredential) Identity() elemental.Identity {

	return K8SCredentialIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *K8SCredential) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *K8SCredential) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *K8SCredential) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *K8SCredential) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *K8SCredential) Doc() string {
	return `Create a credential for a kubernetes cluster.`
}

func (o *K8SCredential) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *K8SCredential) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *K8SCredential) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *K8SCredential) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *K8SCredential) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *K8SCredential) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *K8SCredential) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMetadata returns the Metadata of the receiver.
func (o *K8SCredential) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *K8SCredential) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *K8SCredential) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *K8SCredential) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *K8SCredential) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *K8SCredential) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *K8SCredential) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *K8SCredential) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *K8SCredential) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *K8SCredential) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *K8SCredential) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *K8SCredential) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
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
func (*K8SCredential) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := K8SCredentialAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return K8SCredentialLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*K8SCredential) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return K8SCredentialAttributesMap
}

// K8SCredentialAttributesMap represents the map of attribute for K8SCredential.
var K8SCredentialAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description:    `The string representation of the Certificate used by the Kubernetes cluster.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CertificateSN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSN",
		Description:    `Link to the certificate created for this cluster.`,
		Name:           "certificateSN",
		Stored:         true,
		Type:           "string",
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
	"Email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		Description: `The email address that will receive a copy of the Kubernetes cluster YAMLs
definition.`,
		Exposed:   true,
		Name:      "email",
		Orderable: true,
		Stored:    true,
		Type:      "string",
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
		CreationOnly:   true,
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
	"Regenerate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Regenerate",
		Description:    `Regenerates the k8s files and certificates.`,
		Exposed:        true,
		Name:           "regenerate",
		Type:           "boolean",
	},
	"SecretDefinition": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SecretDefinition",
		Description:    `The secret file to deploy on your cluster.`,
		Exposed:        true,
		Name:           "secretDefinition",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "k8s_secret_definition",
		Type:           "external",
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

// K8SCredentialLowerCaseAttributesMap represents the map of attribute for K8SCredential.
var K8SCredentialLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description:    `The string representation of the Certificate used by the Kubernetes cluster.`,
		Exposed:        true,
		Name:           "certificate",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"certificatesn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSN",
		Description:    `Link to the certificate created for this cluster.`,
		Name:           "certificateSN",
		Stored:         true,
		Type:           "string",
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
	"email": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Email",
		Description: `The email address that will receive a copy of the Kubernetes cluster YAMLs
definition.`,
		Exposed:   true,
		Name:      "email",
		Orderable: true,
		Stored:    true,
		Type:      "string",
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
		CreationOnly:   true,
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
	"regenerate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Regenerate",
		Description:    `Regenerates the k8s files and certificates.`,
		Exposed:        true,
		Name:           "regenerate",
		Type:           "boolean",
	},
	"secretdefinition": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SecretDefinition",
		Description:    `The secret file to deploy on your cluster.`,
		Exposed:        true,
		Name:           "secretDefinition",
		Orderable:      true,
		ReadOnly:       true,
		SubType:        "k8s_secret_definition",
		Type:           "external",
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
