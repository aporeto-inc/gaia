package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// NamespaceIdentity represents the Identity of the object.
var NamespaceIdentity = elemental.Identity{
	Name:     "namespace",
	Category: "namespaces",
	Package:  "squall",
	Private:  false,
}

// NamespacesList represents a list of Namespaces
type NamespacesList []*Namespace

// Identity returns the identity of the objects in the list.
func (o NamespacesList) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Copy returns a pointer to a copy the NamespacesList.
func (o NamespacesList) Copy() elemental.Identifiables {

	copy := append(NamespacesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the NamespacesList.
func (o NamespacesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(NamespacesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Namespace))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o NamespacesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o NamespacesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the NamespacesList converted to SparseNamespacesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o NamespacesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o NamespacesList) Version() int {

	return 1
}

// Namespace represents the model of a namespace
type Namespace struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedLocalCAID holds the remote ID of the certificate authority to use.
	AssociatedLocalCAID string `json:"-" bson:"associatedlocalcaid" mapstructure:"-,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// LocalCA holds the eventual certificate authority used by this namespace.
	LocalCA string `json:"localCA" bson:"localca" mapstructure:"localCA,omitempty"`

	// LocalCAEnabled defines if the namespace should use a local Certificate
	// Authority. Switching it off and on again will regenerate a new CA.
	LocalCAEnabled bool `json:"localCAEnabled" bson:"localcaenabled" mapstructure:"localCAEnabled,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the namespace.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// List of tags that will be added to every `+"`"+`or`+"`"+` clause of all network access
	// policies in the namespace and its children.
	NetworkAccessPolicyTags []string `json:"networkAccessPolicyTags" bson:"networkaccesspolicytags" mapstructure:"networkAccessPolicyTags,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Determines the validity time of certificates issued in this namespace. Default
	// value is 1 hour.
	ServiceCertificateValidity string `json:"serviceCertificateValidity" bson:"servicecertificatevalidity" mapstructure:"serviceCertificateValidity,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewNamespace returns a new *Namespace
func NewNamespace() *Namespace {

	return &Namespace{
		ModelVersion:               1,
		Annotations:                map[string][]string{},
		AssociatedTags:             []string{},
		Metadata:                   []string{},
		NetworkAccessPolicyTags:    []string{},
		NormalizedTags:             []string{},
		ServiceCertificateValidity: "1h",
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

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Namespace) Doc() string {
	return `A Namespace represents the core organizational unit of the system. All objects
always exists in a single namespace. A Namespace can also have child namespaces.
They can be used to split the system into organizations, business units,
applications, services or any combination you like.`
}

func (o *Namespace) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Namespace) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Namespace) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Namespace) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Namespace) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Namespace) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Namespace) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Namespace) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Namespace) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *Namespace) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *Namespace) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *Namespace) GetName() string {

	return o.Name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Namespace) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Namespace) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Namespace) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
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

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Namespace) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Namespace) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseNamespace{
			ID:                         &o.ID,
			Annotations:                &o.Annotations,
			AssociatedLocalCAID:        &o.AssociatedLocalCAID,
			AssociatedTags:             &o.AssociatedTags,
			CreateTime:                 &o.CreateTime,
			Description:                &o.Description,
			LocalCA:                    &o.LocalCA,
			LocalCAEnabled:             &o.LocalCAEnabled,
			Metadata:                   &o.Metadata,
			Name:                       &o.Name,
			Namespace:                  &o.Namespace,
			NetworkAccessPolicyTags:    &o.NetworkAccessPolicyTags,
			NormalizedTags:             &o.NormalizedTags,
			Protected:                  &o.Protected,
			ServiceCertificateValidity: &o.ServiceCertificateValidity,
			UpdateTime:                 &o.UpdateTime,
		}
	}

	sp := &SparseNamespace{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedLocalCAID":
			sp.AssociatedLocalCAID = &(o.AssociatedLocalCAID)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "localCA":
			sp.LocalCA = &(o.LocalCA)
		case "localCAEnabled":
			sp.LocalCAEnabled = &(o.LocalCAEnabled)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "networkAccessPolicyTags":
			sp.NetworkAccessPolicyTags = &(o.NetworkAccessPolicyTags)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "protected":
			sp.Protected = &(o.Protected)
		case "serviceCertificateValidity":
			sp.ServiceCertificateValidity = &(o.ServiceCertificateValidity)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseNamespace to the object.
func (o *Namespace) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseNamespace)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedLocalCAID != nil {
		o.AssociatedLocalCAID = *so.AssociatedLocalCAID
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
	if so.LocalCA != nil {
		o.LocalCA = *so.LocalCA
	}
	if so.LocalCAEnabled != nil {
		o.LocalCAEnabled = *so.LocalCAEnabled
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
	if so.NetworkAccessPolicyTags != nil {
		o.NetworkAccessPolicyTags = *so.NetworkAccessPolicyTags
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.ServiceCertificateValidity != nil {
		o.ServiceCertificateValidity = *so.ServiceCertificateValidity
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the Namespace.
func (o *Namespace) DeepCopy() *Namespace {

	if o == nil {
		return nil
	}

	out := &Namespace{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Namespace.
func (o *Namespace) DeepCopyInto(out *Namespace) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Namespace: %s", err))
	}

	*out = *target.(*Namespace)
}

// Validate valides the current information stored into the structure.
func (o *Namespace) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidatePattern("name", o.Name, `^[a-zA-Z0-9-_/]+$`, `must only contain alpha numerical characters, '-' or '_'`, true); err != nil {
		errors = append(errors, err)
	}

	if err := ValidateTimeDuration("serviceCertificateValidity", o.ServiceCertificateValidity); err != nil {
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Namespace) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedLocalCAID":
		return o.AssociatedLocalCAID
	case "associatedTags":
		return o.AssociatedTags
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "localCA":
		return o.LocalCA
	case "localCAEnabled":
		return o.LocalCAEnabled
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "networkAccessPolicyTags":
		return o.NetworkAccessPolicyTags
	case "normalizedTags":
		return o.NormalizedTags
	case "protected":
		return o.Protected
	case "serviceCertificateValidity":
		return o.ServiceCertificateValidity
	case "updateTime":
		return o.UpdateTime
	}

	return nil
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
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
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
		SubType:        "map_of_string_of_list_of_strings",
		Type:           "external",
	},
	"AssociatedLocalCAID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedLocalCAID",
		Description:    `AssociatedLocalCAID holds the remote ID of the certificate authority to use.`,
		Name:           "associatedLocalCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"LocalCA": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description:    `LocalCA holds the eventual certificate authority used by this namespace.`,
		Exposed:        true,
		Name:           "localCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"LocalCAEnabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LocalCAEnabled",
		Description: `LocalCAEnabled defines if the namespace should use a local Certificate
Authority. Switching it off and on again will regenerate a new CA.`,
		Exposed:   true,
		Name:      "localCAEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
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
		AllowedChars:   `^[a-zA-Z0-9-_/]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		DefaultOrder:   true,
		Description:    `Name is the name of the namespace.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		PrimaryKey:     true,
		Required:       true,
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
	"NetworkAccessPolicyTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAccessPolicyTags",
		Description: `List of tags that will be added to every ` + "`" + `or` + "`" + ` clause of all network access
policies in the namespace and its children.`,
		Exposed: true,
		Name:    "networkAccessPolicyTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
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
	"ServiceCertificateValidity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceCertificateValidity",
		DefaultValue:   "1h",
		Description: `Determines the validity time of certificates issued in this namespace. Default
value is 1 hour.`,
		Exposed: true,
		Name:    "serviceCertificateValidity",
		Stored:  true,
		Type:    "string",
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
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
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
		SubType:        "map_of_string_of_list_of_strings",
		Type:           "external",
	},
	"associatedlocalcaid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedLocalCAID",
		Description:    `AssociatedLocalCAID holds the remote ID of the certificate authority to use.`,
		Name:           "associatedLocalCAID",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
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
	"localca": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "LocalCA",
		Description:    `LocalCA holds the eventual certificate authority used by this namespace.`,
		Exposed:        true,
		Name:           "localCA",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"localcaenabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LocalCAEnabled",
		Description: `LocalCAEnabled defines if the namespace should use a local Certificate
Authority. Switching it off and on again will regenerate a new CA.`,
		Exposed:   true,
		Name:      "localCAEnabled",
		Orderable: true,
		Stored:    true,
		Type:      "boolean",
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
		AllowedChars:   `^[a-zA-Z0-9-_/]+$`,
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		CreationOnly:   true,
		DefaultOrder:   true,
		Description:    `Name is the name of the namespace.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		PrimaryKey:     true,
		Required:       true,
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
	"networkaccesspolicytags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NetworkAccessPolicyTags",
		Description: `List of tags that will be added to every ` + "`" + `or` + "`" + ` clause of all network access
policies in the namespace and its children.`,
		Exposed: true,
		Name:    "networkAccessPolicyTags",
		Stored:  true,
		SubType: "string",
		Type:    "list",
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
	"servicecertificatevalidity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceCertificateValidity",
		DefaultValue:   "1h",
		Description: `Determines the validity time of certificates issued in this namespace. Default
value is 1 hour.`,
		Exposed: true,
		Name:    "serviceCertificateValidity",
		Stored:  true,
		Type:    "string",
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

// SparseNamespacesList represents a list of SparseNamespaces
type SparseNamespacesList []*SparseNamespace

// Identity returns the identity of the objects in the list.
func (o SparseNamespacesList) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Copy returns a pointer to a copy the SparseNamespacesList.
func (o SparseNamespacesList) Copy() elemental.Identifiables {

	copy := append(SparseNamespacesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseNamespacesList.
func (o SparseNamespacesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseNamespacesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseNamespace))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseNamespacesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseNamespacesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseNamespacesList converted to NamespacesList.
func (o SparseNamespacesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseNamespacesList) Version() int {

	return 1
}

// SparseNamespace represents the sparse version of a namespace.
type SparseNamespace struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedLocalCAID holds the remote ID of the certificate authority to use.
	AssociatedLocalCAID *string `json:"-,omitempty" bson:"associatedlocalcaid" mapstructure:"-,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// LocalCA holds the eventual certificate authority used by this namespace.
	LocalCA *string `json:"localCA,omitempty" bson:"localca" mapstructure:"localCA,omitempty"`

	// LocalCAEnabled defines if the namespace should use a local Certificate
	// Authority. Switching it off and on again will regenerate a new CA.
	LocalCAEnabled *bool `json:"localCAEnabled,omitempty" bson:"localcaenabled" mapstructure:"localCAEnabled,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the namespace.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" bson:"namespace" mapstructure:"namespace,omitempty"`

	// List of tags that will be added to every `+"`"+`or`+"`"+` clause of all network access
	// policies in the namespace and its children.
	NetworkAccessPolicyTags *[]string `json:"networkAccessPolicyTags,omitempty" bson:"networkaccesspolicytags" mapstructure:"networkAccessPolicyTags,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" bson:"protected" mapstructure:"protected,omitempty"`

	// Determines the validity time of certificates issued in this namespace. Default
	// value is 1 hour.
	ServiceCertificateValidity *string `json:"serviceCertificateValidity,omitempty" bson:"servicecertificatevalidity" mapstructure:"serviceCertificateValidity,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseNamespace returns a new  SparseNamespace.
func NewSparseNamespace() *SparseNamespace {
	return &SparseNamespace{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseNamespace) Identity() elemental.Identity {

	return NamespaceIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseNamespace) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseNamespace) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseNamespace) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseNamespace) ToPlain() elemental.PlainIdentifiable {

	out := NewNamespace()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedLocalCAID != nil {
		out.AssociatedLocalCAID = *o.AssociatedLocalCAID
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
	if o.LocalCA != nil {
		out.LocalCA = *o.LocalCA
	}
	if o.LocalCAEnabled != nil {
		out.LocalCAEnabled = *o.LocalCAEnabled
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
	if o.NetworkAccessPolicyTags != nil {
		out.NetworkAccessPolicyTags = *o.NetworkAccessPolicyTags
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.ServiceCertificateValidity != nil {
		out.ServiceCertificateValidity = *o.ServiceCertificateValidity
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseNamespace) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseNamespace) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseNamespace) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseNamespace) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseNamespace) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseNamespace) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseNamespace) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseNamespace) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseNamespace) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseNamespace) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseNamespace) GetName() string {

	return *o.Name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseNamespace) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseNamespace) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseNamespace) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseNamespace) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *SparseNamespace) GetProtected() bool {

	return *o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseNamespace) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseNamespace) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseNamespace.
func (o *SparseNamespace) DeepCopy() *SparseNamespace {

	if o == nil {
		return nil
	}

	out := &SparseNamespace{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseNamespace.
func (o *SparseNamespace) DeepCopyInto(out *SparseNamespace) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseNamespace: %s", err))
	}

	*out = *target.(*SparseNamespace)
}
