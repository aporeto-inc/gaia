package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"time"
)

// K8SClusterActivationTypeValue represents the possible values for attribute "activationType".
type K8SClusterActivationTypeValue string

const (
	// K8SClusterActivationTypeKubeSquall represents the value KubeSquall.
	K8SClusterActivationTypeKubeSquall K8SClusterActivationTypeValue = "KubeSquall"

	// K8SClusterActivationTypePodAtomic represents the value PodAtomic.
	K8SClusterActivationTypePodAtomic K8SClusterActivationTypeValue = "PodAtomic"

	// K8SClusterActivationTypePodContainers represents the value PodContainers.
	K8SClusterActivationTypePodContainers K8SClusterActivationTypeValue = "PodContainers"
)

// K8SClusterNetworkPolicyTypeValue represents the possible values for attribute "networkPolicyType".
type K8SClusterNetworkPolicyTypeValue string

const (
	// K8SClusterNetworkPolicyTypeKubernetes represents the value Kubernetes.
	K8SClusterNetworkPolicyTypeKubernetes K8SClusterNetworkPolicyTypeValue = "Kubernetes"

	// K8SClusterNetworkPolicyTypeNoPolicy represents the value NoPolicy.
	K8SClusterNetworkPolicyTypeNoPolicy K8SClusterNetworkPolicyTypeValue = "NoPolicy"
)

// K8SClusterIdentity represents the Identity of the object.
var K8SClusterIdentity = elemental.Identity{
	Name:     "k8scluster",
	Category: "k8sclusters",
	Private:  false,
}

// K8SClustersList represents a list of K8SClusters
type K8SClustersList []*K8SCluster

// ContentIdentity returns the identity of the objects in the list.
func (o K8SClustersList) ContentIdentity() elemental.Identity {

	return K8SClusterIdentity
}

// Copy returns a pointer to a copy the K8SClustersList.
func (o K8SClustersList) Copy() elemental.ContentIdentifiable {

	copy := append(K8SClustersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the K8SClustersList.
func (o K8SClustersList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(K8SClustersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*K8SCluster))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o K8SClustersList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o K8SClustersList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o K8SClustersList) Version() int {

	return 1
}

// K8SCluster represents the model of a k8scluster
type K8SCluster struct {
	// Link to the API authorization policy.
	APIAuthorizationPolicyID string `json:"-" bson:"apiauthorizationpolicyid" mapstructure:"-,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Defines the mode of activation on the KubernetesCluster.
	ActivationType K8SClusterActivationTypeValue `json:"activationType" bson:"activationtype" mapstructure:"activationType,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Link to the certificate created for this cluster.
	CertificateSN string `json:"-" bson:"certificatesn" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// base64 of the .tar.gz file that contains all the .YAMLs files needed to create
	// the aporeto side on your kubernetes Cluster.
	KubernetesDefinitions string `json:"kubernetesDefinitions" bson:"-" mapstructure:"kubernetesDefinitions,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Defines what type of network policy will be applied on your cluster.
	// Kubernetes means that All the Kubernetes policies will be synced to Squall.
	// No Policies means that policies are not synced and it's up to the user to create
	// consistent policies in Squall.
	NetworkPolicyType K8SClusterNetworkPolicyTypeValue `json:"networkPolicyType" bson:"networkpolicytype" mapstructure:"networkPolicyType,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Regenerates the k8s files and certificates.
	Regenerate bool `json:"regenerate" bson:"-" mapstructure:"regenerate,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewK8SCluster returns a new *K8SCluster
func NewK8SCluster() *K8SCluster {

	return &K8SCluster{
		ModelVersion:      1,
		ActivationType:    "KubeSquall",
		Annotations:       map[string][]string{},
		AssociatedTags:    []string{},
		Metadata:          []string{},
		NetworkPolicyType: "Kubernetes",
		NormalizedTags:    []string{},
	}
}

// Identity returns the Identity of the object.
func (o *K8SCluster) Identity() elemental.Identity {

	return K8SClusterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *K8SCluster) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *K8SCluster) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *K8SCluster) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *K8SCluster) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *K8SCluster) Doc() string {
	return `Create a remote Kubernetes Cluster integration.`
}

func (o *K8SCluster) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *K8SCluster) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *K8SCluster) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *K8SCluster) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *K8SCluster) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetMetadata returns the Metadata of the receiver.
func (o *K8SCluster) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *K8SCluster) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *K8SCluster) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *K8SCluster) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *K8SCluster) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *K8SCluster) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *K8SCluster) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *K8SCluster) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *K8SCluster) GetProtected() bool {

	return o.Protected
}

// Validate valides the current information stored into the structure.
func (o *K8SCluster) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("activationType", string(o.ActivationType), []string{"KubeSquall", "PodAtomic", "PodContainers"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("networkPolicyType", string(o.NetworkPolicyType), []string{"Kubernetes", "NoPolicy"}, false); err != nil {
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
func (*K8SCluster) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := K8SClusterAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return K8SClusterLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*K8SCluster) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return K8SClusterAttributesMap
}

// K8SClusterAttributesMap represents the map of attribute for K8SCluster.
var K8SClusterAttributesMap = map[string]elemental.AttributeSpecification{
	"APIAuthorizationPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "APIAuthorizationPolicyID",
		Description:    `Link to the API authorization policy.`,
		Format:         "free",
		Name:           "APIAuthorizationPolicyID",
		Stored:         true,
		Type:           "string",
	},
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
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ActivationType": elemental.AttributeSpecification{
		AllowedChoices: []string{"KubeSquall", "PodAtomic", "PodContainers"},
		ConvertedName:  "ActivationType",
		DefaultValue:   K8SClusterActivationTypeKubeSquall,
		Description:    `Defines the mode of activation on the KubernetesCluster.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "activationType",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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
	"CertificateSN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSN",
		Description:    `Link to the certificate created for this cluster.`,
		Format:         "free",
		Name:           "certificateSN",
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Format:         "free",
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"KubernetesDefinitions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesDefinitions",
		Description: `base64 of the .tar.gz file that contains all the .YAMLs files needed to create
the aporeto side on your kubernetes Cluster.`,
		Exposed:   true,
		Format:    "free",
		Name:      "kubernetesDefinitions",
		Orderable: true,
		ReadOnly:  true,
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
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
	"NetworkPolicyType": elemental.AttributeSpecification{
		AllowedChoices: []string{"Kubernetes", "NoPolicy"},
		ConvertedName:  "NetworkPolicyType",
		DefaultValue:   K8SClusterNetworkPolicyTypeKubernetes,
		Description: `Defines what type of network policy will be applied on your cluster.
Kubernetes means that All the Kubernetes policies will be synced to Squall.
No Policies means that policies are not synced and it's up to the user to create
consistent policies in Squall.`,
		Exposed:    true,
		Filterable: true,
		Name:       "networkPolicyType",
		Orderable:  true,
		Stored:     true,
		Type:       "enum",
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
		Filterable:     true,
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
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}

// K8SClusterLowerCaseAttributesMap represents the map of attribute for K8SCluster.
var K8SClusterLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiauthorizationpolicyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "APIAuthorizationPolicyID",
		Description:    `Link to the API authorization policy.`,
		Format:         "free",
		Name:           "APIAuthorizationPolicyID",
		Stored:         true,
		Type:           "string",
	},
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
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"activationtype": elemental.AttributeSpecification{
		AllowedChoices: []string{"KubeSquall", "PodAtomic", "PodContainers"},
		ConvertedName:  "ActivationType",
		DefaultValue:   K8SClusterActivationTypeKubeSquall,
		Description:    `Defines the mode of activation on the KubernetesCluster.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "activationType",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
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
	"certificatesn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateSN",
		Description:    `Link to the certificate created for this cluster.`,
		Format:         "free",
		Name:           "certificateSN",
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Format:         "free",
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"kubernetesdefinitions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesDefinitions",
		Description: `base64 of the .tar.gz file that contains all the .YAMLs files needed to create
the aporeto side on your kubernetes Cluster.`,
		Exposed:   true,
		Format:    "free",
		Name:      "kubernetesDefinitions",
		Orderable: true,
		ReadOnly:  true,
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
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
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
	"networkpolicytype": elemental.AttributeSpecification{
		AllowedChoices: []string{"Kubernetes", "NoPolicy"},
		ConvertedName:  "NetworkPolicyType",
		DefaultValue:   K8SClusterNetworkPolicyTypeKubernetes,
		Description: `Defines what type of network policy will be applied on your cluster.
Kubernetes means that All the Kubernetes policies will be synced to Squall.
No Policies means that policies are not synced and it's up to the user to create
consistent policies in Squall.`,
		Exposed:    true,
		Filterable: true,
		Name:       "networkPolicyType",
		Orderable:  true,
		Stored:     true,
		Type:       "enum",
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
		Filterable:     true,
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
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}
