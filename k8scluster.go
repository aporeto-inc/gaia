package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
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
	Package:  "cactuar",
	Private:  false,
}

// K8SClustersList represents a list of K8SClusters
type K8SClustersList []*K8SCluster

// Identity returns the identity of the objects in the list.
func (o K8SClustersList) Identity() elemental.Identity {

	return K8SClusterIdentity
}

// Copy returns a pointer to a copy the K8SClustersList.
func (o K8SClustersList) Copy() elemental.Identifiables {

	copy := append(K8SClustersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the K8SClustersList.
func (o K8SClustersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(K8SClustersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*K8SCluster))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o K8SClustersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o K8SClustersList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the K8SClustersList converted to SparseK8SClustersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o K8SClustersList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
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

	// The email address that will receive a copy of the Kubernetes cluster YAMLs
	// definition.
	AdminEmail string `json:"adminEmail" bson:"adminemail" mapstructure:"adminEmail,omitempty"`

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

	// Link to the cluster namespace.
	NamespaceID string `json:"-" bson:"namespaceid" mapstructure:"-,omitempty"`

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

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewK8SCluster returns a new *K8SCluster
func NewK8SCluster() *K8SCluster {

	return &K8SCluster{
		ModelVersion:      1,
		Metadata:          []string{},
		ActivationType:    K8SClusterActivationTypeKubeSquall,
		Annotations:       map[string][]string{},
		AssociatedTags:    []string{},
		NetworkPolicyType: K8SClusterNetworkPolicyTypeKubernetes,
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

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *K8SCluster) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *K8SCluster) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *K8SCluster) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *K8SCluster) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *K8SCluster) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *K8SCluster) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *K8SCluster) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *K8SCluster) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *K8SCluster) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *K8SCluster) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *K8SCluster) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *K8SCluster) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *K8SCluster) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *K8SCluster) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *K8SCluster) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *K8SCluster) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *K8SCluster) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *K8SCluster) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *K8SCluster) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *K8SCluster) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *K8SCluster) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *K8SCluster) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *K8SCluster) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseK8SCluster{
			APIAuthorizationPolicyID: &o.APIAuthorizationPolicyID,
			ID:                       &o.ID,
			ActivationType:           &o.ActivationType,
			AdminEmail:               &o.AdminEmail,
			Annotations:              &o.Annotations,
			AssociatedTags:           &o.AssociatedTags,
			Certificate:              &o.Certificate,
			CertificateSN:            &o.CertificateSN,
			CreateTime:               &o.CreateTime,
			Description:              &o.Description,
			KubernetesDefinitions:    &o.KubernetesDefinitions,
			Metadata:                 &o.Metadata,
			Name:                     &o.Name,
			Namespace:                &o.Namespace,
			NamespaceID:              &o.NamespaceID,
			NetworkPolicyType:        &o.NetworkPolicyType,
			NormalizedTags:           &o.NormalizedTags,
			Protected:                &o.Protected,
			Regenerate:               &o.Regenerate,
			UpdateTime:               &o.UpdateTime,
			ZHash:                    &o.ZHash,
			Zone:                     &o.Zone,
		}
	}

	sp := &SparseK8SCluster{}
	for _, f := range fields {
		switch f {
		case "APIAuthorizationPolicyID":
			sp.APIAuthorizationPolicyID = &(o.APIAuthorizationPolicyID)
		case "ID":
			sp.ID = &(o.ID)
		case "activationType":
			sp.ActivationType = &(o.ActivationType)
		case "adminEmail":
			sp.AdminEmail = &(o.AdminEmail)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "certificateSN":
			sp.CertificateSN = &(o.CertificateSN)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "kubernetesDefinitions":
			sp.KubernetesDefinitions = &(o.KubernetesDefinitions)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "namespaceID":
			sp.NamespaceID = &(o.NamespaceID)
		case "networkPolicyType":
			sp.NetworkPolicyType = &(o.NetworkPolicyType)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "regenerate":
			sp.Regenerate = &(o.Regenerate)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseK8SCluster to the object.
func (o *K8SCluster) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseK8SCluster)
	if so.APIAuthorizationPolicyID != nil {
		o.APIAuthorizationPolicyID = *so.APIAuthorizationPolicyID
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.ActivationType != nil {
		o.ActivationType = *so.ActivationType
	}
	if so.AdminEmail != nil {
		o.AdminEmail = *so.AdminEmail
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.CertificateSN != nil {
		o.CertificateSN = *so.CertificateSN
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.KubernetesDefinitions != nil {
		o.KubernetesDefinitions = *so.KubernetesDefinitions
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
	if so.NamespaceID != nil {
		o.NamespaceID = *so.NamespaceID
	}
	if so.NetworkPolicyType != nil {
		o.NetworkPolicyType = *so.NetworkPolicyType
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Regenerate != nil {
		o.Regenerate = *so.Regenerate
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the K8SCluster.
func (o *K8SCluster) DeepCopy() *K8SCluster {

	if o == nil {
		return nil
	}

	out := &K8SCluster{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *K8SCluster.
func (o *K8SCluster) DeepCopyInto(out *K8SCluster) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy K8SCluster: %s", err))
	}

	*out = *target.(*K8SCluster)
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *K8SCluster) ValueForAttribute(name string) interface{} {

	switch name {
	case "APIAuthorizationPolicyID":
		return o.APIAuthorizationPolicyID
	case "ID":
		return o.ID
	case "activationType":
		return o.ActivationType
	case "adminEmail":
		return o.AdminEmail
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "certificate":
		return o.Certificate
	case "certificateSN":
		return o.CertificateSN
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "kubernetesDefinitions":
		return o.KubernetesDefinitions
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "namespaceID":
		return o.NamespaceID
	case "networkPolicyType":
		return o.NetworkPolicyType
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "regenerate":
		return o.Regenerate
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// K8SClusterAttributesMap represents the map of attribute for K8SCluster.
var K8SClusterAttributesMap = map[string]elemental.AttributeSpecification{
	"APIAuthorizationPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "APIAuthorizationPolicyID",
		Description:    `Link to the API authorization policy.`,
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
		Name:           "activationType",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"AdminEmail": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AdminEmail",
		Description: `The email address that will receive a copy of the Kubernetes cluster YAMLs
definition.`,
		Exposed:   true,
		Name:      "adminEmail",
		Orderable: true,
		Stored:    true,
		Type:      "string",
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
		SubType:        "map_of_string_of_list_of_strings",
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
		SubType:        "string",
		Type:           "list",
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
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"KubernetesDefinitions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesDefinitions",
		Description: `base64 of the .tar.gz file that contains all the .YAMLs files needed to create
the aporeto side on your kubernetes Cluster.`,
		Exposed:   true,
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
		SubType:    "string",
		Type:       "list",
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
	"NamespaceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		Description:    `Link to the cluster namespace.`,
		Name:           "namespaceID",
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
		Exposed:   true,
		Name:      "networkPolicyType",
		Orderable: true,
		Stored:    true,
		Type:      "enum",
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
		SubType:        "string",
		Transient:      true,
		Type:           "list",
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
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// K8SClusterLowerCaseAttributesMap represents the map of attribute for K8SCluster.
var K8SClusterLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"apiauthorizationpolicyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "APIAuthorizationPolicyID",
		Description:    `Link to the API authorization policy.`,
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
		Name:           "activationType",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"adminemail": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AdminEmail",
		Description: `The email address that will receive a copy of the Kubernetes cluster YAMLs
definition.`,
		Exposed:   true,
		Name:      "adminEmail",
		Orderable: true,
		Stored:    true,
		Type:      "string",
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
		SubType:        "map_of_string_of_list_of_strings",
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
		SubType:        "string",
		Type:           "list",
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
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"kubernetesdefinitions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "KubernetesDefinitions",
		Description: `base64 of the .tar.gz file that contains all the .YAMLs files needed to create
the aporeto side on your kubernetes Cluster.`,
		Exposed:   true,
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
		SubType:    "string",
		Type:       "list",
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
	"namespaceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NamespaceID",
		Description:    `Link to the cluster namespace.`,
		Name:           "namespaceID",
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
		Exposed:   true,
		Name:      "networkPolicyType",
		Orderable: true,
		Stored:    true,
		Type:      "enum",
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
		SubType:        "string",
		Transient:      true,
		Type:           "list",
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
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
}

// SparseK8SClustersList represents a list of SparseK8SClusters
type SparseK8SClustersList []*SparseK8SCluster

// Identity returns the identity of the objects in the list.
func (o SparseK8SClustersList) Identity() elemental.Identity {

	return K8SClusterIdentity
}

// Copy returns a pointer to a copy the SparseK8SClustersList.
func (o SparseK8SClustersList) Copy() elemental.Identifiables {

	copy := append(SparseK8SClustersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseK8SClustersList.
func (o SparseK8SClustersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseK8SClustersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseK8SCluster))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseK8SClustersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseK8SClustersList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseK8SClustersList converted to K8SClustersList.
func (o SparseK8SClustersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseK8SClustersList) Version() int {

	return 1
}

// SparseK8SCluster represents the sparse version of a k8scluster.
type SparseK8SCluster struct {
	// Link to the API authorization policy.
	APIAuthorizationPolicyID *string `json:"-,omitempty" bson:"apiauthorizationpolicyid" mapstructure:"-,omitempty"`

	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Defines the mode of activation on the KubernetesCluster.
	ActivationType *K8SClusterActivationTypeValue `json:"activationType,omitempty" bson:"activationtype" mapstructure:"activationType,omitempty"`

	// The email address that will receive a copy of the Kubernetes cluster YAMLs
	// definition.
	AdminEmail *string `json:"adminEmail,omitempty" bson:"adminemail" mapstructure:"adminEmail,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// The string representation of the Certificate used by the Kubernetes cluster.
	Certificate *string `json:"certificate,omitempty" bson:"certificate" mapstructure:"certificate,omitempty"`

	// Link to the certificate created for this cluster.
	CertificateSN *string `json:"-,omitempty" bson:"certificatesn" mapstructure:"-,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// base64 of the .tar.gz file that contains all the .YAMLs files needed to create
	// the aporeto side on your kubernetes Cluster.
	KubernetesDefinitions *string `json:"kubernetesDefinitions,omitempty" bson:"-" mapstructure:"kubernetesDefinitions,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Link to the cluster namespace.
	NamespaceID *string `json:"-,omitempty" bson:"namespaceid" mapstructure:"-,omitempty"`

	// Defines what type of network policy will be applied on your cluster.
	// Kubernetes means that All the Kubernetes policies will be synced to Squall.
	// No Policies means that policies are not synced and it's up to the user to create
	// consistent policies in Squall.
	NetworkPolicyType *K8SClusterNetworkPolicyTypeValue `json:"networkPolicyType,omitempty" bson:"networkpolicytype" mapstructure:"networkPolicyType,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// Regenerates the k8s files and certificates.
	Regenerate *bool `json:"regenerate,omitempty" bson:"-" mapstructure:"regenerate,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-,omitempty" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-,omitempty" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseK8SCluster returns a new  SparseK8SCluster.
func NewSparseK8SCluster() *SparseK8SCluster {
	return &SparseK8SCluster{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseK8SCluster) Identity() elemental.Identity {

	return K8SClusterIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseK8SCluster) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseK8SCluster) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseK8SCluster) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseK8SCluster) ToPlain() elemental.PlainIdentifiable {

	out := NewK8SCluster()
	if o.APIAuthorizationPolicyID != nil {
		out.APIAuthorizationPolicyID = *o.APIAuthorizationPolicyID
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.ActivationType != nil {
		out.ActivationType = *o.ActivationType
	}
	if o.AdminEmail != nil {
		out.AdminEmail = *o.AdminEmail
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.CertificateSN != nil {
		out.CertificateSN = *o.CertificateSN
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.KubernetesDefinitions != nil {
		out.KubernetesDefinitions = *o.KubernetesDefinitions
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
	if o.NamespaceID != nil {
		out.NamespaceID = *o.NamespaceID
	}
	if o.NetworkPolicyType != nil {
		out.NetworkPolicyType = *o.NetworkPolicyType
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Regenerate != nil {
		out.Regenerate = *o.Regenerate
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseK8SCluster) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseK8SCluster) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseK8SCluster) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseK8SCluster) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseK8SCluster) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseK8SCluster) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseK8SCluster) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseK8SCluster) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseK8SCluster) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseK8SCluster) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseK8SCluster) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseK8SCluster) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseK8SCluster) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseK8SCluster) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseK8SCluster) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseK8SCluster) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseK8SCluster) GetProtected() bool {

	return *o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseK8SCluster) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseK8SCluster) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseK8SCluster) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseK8SCluster) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseK8SCluster) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseK8SCluster) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseK8SCluster.
func (o *SparseK8SCluster) DeepCopy() *SparseK8SCluster {

	if o == nil {
		return nil
	}

	out := &SparseK8SCluster{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseK8SCluster.
func (o *SparseK8SCluster) DeepCopyInto(out *SparseK8SCluster) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseK8SCluster: %s", err))
	}

	*out = *target.(*SparseK8SCluster)
}
