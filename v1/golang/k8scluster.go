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

	return []string{}
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

	// Link to the certificate created in Vince for this cluster.
	CertificateID string `json:"-" bson:"certificateid" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// base64 of the .tar.gz file that contains all the .YAMLs files needed to create
	// the aporeto side on your kubernetes Cluster.
	KubernetesDefinitions string `json:"kubernetesDefinitions" bson:"-" mapstructure:"kubernetesDefinitions,omitempty"`

	// The name of your cluster.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Link to your namespace.
	NamespaceID string `json:"-" bson:"namespaceid" mapstructure:"-,omitempty"`

	// ID of the parent account.
	ParentID string `json:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// Regenerates the k8s files and certificates.
	Regenerate bool `json:"regenerate" bson:"-" mapstructure:"regenerate,omitempty"`

	// The namespace in which the Kubernetes specific namespace will be created. By
	// default your account namespace.
	TargetNamespace string `json:"targetNamespace" bson:"targetnamespace" mapstructure:"targetNamespace,omitempty"`

	// List of target networks.
	TargetNetworks []string `json:"targetNetworks" bson:"targetnetworks" mapstructure:"targetNetworks,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewK8SCluster returns a new *K8SCluster
func NewK8SCluster() *K8SCluster {

	return &K8SCluster{
		ModelVersion:   1,
		ActivationType: "KubeSquall",
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

	return []string{}
}

// Doc returns the documentation for the object
func (o *K8SCluster) Doc() string {
	return `Create a remote Kubernetes Cluster integration.`
}

func (o *K8SCluster) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *K8SCluster) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("activationType", string(o.ActivationType), []string{"KubeSquall", "PodAtomic", "PodContainers"}, false); err != nil {
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
	"CertificateID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateID",
		Description:    `Link to the certificate created in Vince for this cluster.`,
		Format:         "free",
		Name:           "certificateID",
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
	"KubernetesDefinitions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesDefinitions",
		Description: `base64 of the .tar.gz file that contains all the .YAMLs files needed to create
the aporeto side on your kubernetes Cluster.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "kubernetesDefinitions",
		Orderable:  true,
		ReadOnly:   true,
		Type:       "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name of your cluster.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"NamespaceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		Description:    `Link to your namespace.`,
		Format:         "free",
		Name:           "namespaceID",
		Stored:         true,
		Type:           "string",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ParentID",
		Description:    `ID of the parent account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Regenerate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Regenerate",
		Description:    `Regenerates the k8s files and certificates.`,
		Exposed:        true,
		Name:           "regenerate",
		Type:           "boolean",
	},
	"TargetNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description: `The namespace in which the Kubernetes specific namespace will be created. By
default your account namespace.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "targetNamespace",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"TargetNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNetworks",
		Deprecated:     true,
		Description:    `List of target networks.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetNetworks",
		Orderable:      true,
		Stored:         true,
		SubType:        "target_networks_list",
		Type:           "external",
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
	"certificateid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CertificateID",
		Description:    `Link to the certificate created in Vince for this cluster.`,
		Format:         "free",
		Name:           "certificateID",
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
	"kubernetesdefinitions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesDefinitions",
		Description: `base64 of the .tar.gz file that contains all the .YAMLs files needed to create
the aporeto side on your kubernetes Cluster.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "kubernetesDefinitions",
		Orderable:  true,
		ReadOnly:   true,
		Type:       "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name of your cluster.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"namespaceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		Description:    `Link to your namespace.`,
		Format:         "free",
		Name:           "namespaceID",
		Stored:         true,
		Type:           "string",
	},
	"parentid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ParentID",
		Description:    `ID of the parent account.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"regenerate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Regenerate",
		Description:    `Regenerates the k8s files and certificates.`,
		Exposed:        true,
		Name:           "regenerate",
		Type:           "boolean",
	},
	"targetnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNamespace",
		Description: `The namespace in which the Kubernetes specific namespace will be created. By
default your account namespace.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "targetNamespace",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
	},
	"targetnetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetNetworks",
		Deprecated:     true,
		Description:    `List of target networks.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "targetNetworks",
		Orderable:      true,
		Stored:         true,
		SubType:        "target_networks_list",
		Type:           "external",
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
