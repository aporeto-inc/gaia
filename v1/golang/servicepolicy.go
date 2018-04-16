package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/types"
	"time"
)

// ServicePolicyActionValue represents the possible values for attribute "action".
type ServicePolicyActionValue string

const (
	// ServicePolicyActionConsume represents the value Consume.
	ServicePolicyActionConsume ServicePolicyActionValue = "Consume"

	// ServicePolicyActionExpose represents the value Expose.
	ServicePolicyActionExpose ServicePolicyActionValue = "Expose"
)

// ServicePolicyIdentity represents the Identity of the object.
var ServicePolicyIdentity = elemental.Identity{
	Name:     "servicepolicy",
	Category: "servicepolicies",
	Private:  false,
}

// ServicePoliciesList represents a list of ServicePolicies
type ServicePoliciesList []*ServicePolicy

// ContentIdentity returns the identity of the objects in the list.
func (o ServicePoliciesList) ContentIdentity() elemental.Identity {

	return ServicePolicyIdentity
}

// Copy returns a pointer to a copy the ServicePoliciesList.
func (o ServicePoliciesList) Copy() elemental.ContentIdentifiable {

	copy := append(ServicePoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ServicePoliciesList.
func (o ServicePoliciesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(ServicePoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*ServicePolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ServicePoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ServicePoliciesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o ServicePoliciesList) Version() int {

	return 1
}

// ServicePolicy represents the model of a servicepolicy
type ServicePolicy struct {
	// FQDN is the fully qualified domain name of the the servers behind the services
	// matched by object. FQND must match the host part
	// of the URI that is used to call a service. It will be used for automatically
	// generating service certificates for internal services.
	FQDN string `json:"FQDN" bson:"fqdn" mapstructure:"FQDN,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// IPList is the list of IP addresses or subnets of the servers behind the services
	// matched by object. This is an optional field and it can be
	// automatically populated at runtime by the enforcers if DNS resolution is
	// available.
	IPList types.IPList `json:"IPList" bson:"iplist" mapstructure:"IPList,omitempty"`

	// Action tells if the processing units matched by the subject are
	// consuming or exposing the services matched by object.
	Action ServicePolicyActionValue `json:"action" bson:"action" mapstructure:"action,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" bson:"disabled" mapstructure:"disabled,omitempty"`

	// Fallback indicates that this is fallback policy. It will only be
	// applied if no other policies have been resolved. If the policy is also
	// propagated it will become a fallback for children namespaces.
	Fallback bool `json:"fallback" bson:"fallback" mapstructure:"fallback,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Objects contains the tag expression that an object must match in order to
	// trigger the hook.
	Object [][]string `json:"object" bson:"object" mapstructure:"object,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// If set to true while the policy is propagating, it won't be visible to children
	// namespace, but still used for policy resolution.
	PropagationHidden bool `json:"propagationHidden" bson:"propagationhidden" mapstructure:"propagationHidden,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// subject contains the tag expression that an a processing unit
	// must match in order to
	// trigger the hook.
	Subject [][]string `json:"subject" bson:"subject" mapstructure:"subject,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewServicePolicy returns a new *ServicePolicy
func NewServicePolicy() *ServicePolicy {

	return &ServicePolicy{
		ModelVersion:   1,
		Action:         "Expose",
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		IPList:         types.IPList{},
		Metadata:       []string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *ServicePolicy) Identity() elemental.Identity {

	return ServicePolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *ServicePolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *ServicePolicy) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *ServicePolicy) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *ServicePolicy) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *ServicePolicy) Doc() string {
	return `A ServicePolicy allows to declare what Services one or a group of processing
units
are exposed or consuming.`
}

func (o *ServicePolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *ServicePolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *ServicePolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *ServicePolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *ServicePolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *ServicePolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *ServicePolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDisabled returns the Disabled of the receiver.
func (o *ServicePolicy) GetDisabled() bool {

	return o.Disabled
}

// SetDisabled sets the given Disabled of the receiver.
func (o *ServicePolicy) SetDisabled(disabled bool) {

	o.Disabled = disabled
}

// GetFallback returns the Fallback of the receiver.
func (o *ServicePolicy) GetFallback() bool {

	return o.Fallback
}

// SetFallback sets the given Fallback of the receiver.
func (o *ServicePolicy) SetFallback(fallback bool) {

	o.Fallback = fallback
}

// GetMetadata returns the Metadata of the receiver.
func (o *ServicePolicy) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the given Metadata of the receiver.
func (o *ServicePolicy) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *ServicePolicy) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *ServicePolicy) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *ServicePolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *ServicePolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *ServicePolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *ServicePolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *ServicePolicy) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the given Propagate of the receiver.
func (o *ServicePolicy) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetPropagationHidden returns the PropagationHidden of the receiver.
func (o *ServicePolicy) GetPropagationHidden() bool {

	return o.PropagationHidden
}

// SetPropagationHidden sets the given PropagationHidden of the receiver.
func (o *ServicePolicy) SetPropagationHidden(propagationHidden bool) {

	o.PropagationHidden = propagationHidden
}

// GetProtected returns the Protected of the receiver.
func (o *ServicePolicy) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *ServicePolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *ServicePolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *ServicePolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Consume", "Expose"}, false); err != nil {
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

	if err := elemental.ValidateRequiredExternal("object", o.Object); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("subject", o.Subject); err != nil {
		requiredErrors = append(requiredErrors, err)
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
func (*ServicePolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ServicePolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ServicePolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*ServicePolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServicePolicyAttributesMap
}

// ServicePolicyAttributesMap represents the map of attribute for ServicePolicy.
var ServicePolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"FQDN": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FQDN",
		Description: `FQDN is the fully qualified domain name of the the servers behind the services
matched by object. FQND must match the host part
of the URI that is used to call a service. It will be used for automatically
generating service certificates for internal services.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "FQDN",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
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
		ReadOnly:       true,
		Type:           "string",
	},
	"IPList": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPList",
		Description: `IPList is the list of IP addresses or subnets of the servers behind the services
matched by object. This is an optional field and it can be
automatically populated at runtime by the enforcers if DNS resolution is
available.`,
		Exposed: true,
		Name:    "IPList",
		Stored:  true,
		SubType: "ip_list",
		Type:    "external",
	},
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Consume", "Expose"},
		ConvertedName:  "Action",
		DefaultValue:   ServicePolicyActionExpose,
		Description: `Action tells if the processing units matched by the subject are
consuming or exposing the services matched by object.`,
		Exposed:    true,
		Filterable: true,
		Name:       "action",
		Orderable:  true,
		Required:   true,
		Stored:     true,
		Type:       "enum",
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
		Format:         "free",
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "fallback",
		Orderable:  true,
		Setter:     true,
		Stored:     true,
		Type:       "boolean",
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
	"Object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description: `Objects contains the tag expression that an object must match in order to
trigger the hook.`,
		Exposed:  true,
		Name:     "object",
		Required: true,
		Stored:   true,
		SubType:  "policies_list",
		Type:     "external",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"PropagationHidden": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PropagationHidden",
		Description: `If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "propagationHidden",
		Orderable:  true,
		Setter:     true,
		Stored:     true,
		Type:       "boolean",
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
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `subject contains the tag expression that an a processing unit
must match in order to
trigger the hook.`,
		Exposed:  true,
		Name:     "subject",
		Required: true,
		Stored:   true,
		SubType:  "policies_list",
		Type:     "external",
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

// ServicePolicyLowerCaseAttributesMap represents the map of attribute for ServicePolicy.
var ServicePolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"fqdn": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FQDN",
		Description: `FQDN is the fully qualified domain name of the the servers behind the services
matched by object. FQND must match the host part
of the URI that is used to call a service. It will be used for automatically
generating service certificates for internal services.`,
		Exposed:    true,
		Filterable: true,
		Format:     "free",
		Name:       "FQDN",
		Orderable:  true,
		Stored:     true,
		Type:       "string",
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
		ReadOnly:       true,
		Type:           "string",
	},
	"iplist": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IPList",
		Description: `IPList is the list of IP addresses or subnets of the servers behind the services
matched by object. This is an optional field and it can be
automatically populated at runtime by the enforcers if DNS resolution is
available.`,
		Exposed: true,
		Name:    "IPList",
		Stored:  true,
		SubType: "ip_list",
		Type:    "external",
	},
	"action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Consume", "Expose"},
		ConvertedName:  "Action",
		DefaultValue:   ServicePolicyActionExpose,
		Description: `Action tells if the processing units matched by the subject are
consuming or exposing the services matched by object.`,
		Exposed:    true,
		Filterable: true,
		Name:       "action",
		Orderable:  true,
		Required:   true,
		Stored:     true,
		Type:       "enum",
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
		Format:         "free",
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Disabled",
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"fallback": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Fallback",
		Description: `Fallback indicates that this is fallback policy. It will only be
applied if no other policies have been resolved. If the policy is also
propagated it will become a fallback for children namespaces.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "fallback",
		Orderable:  true,
		Setter:     true,
		Stored:     true,
		Type:       "boolean",
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
	"object": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Object",
		Description: `Objects contains the tag expression that an object must match in order to
trigger the hook.`,
		Exposed:  true,
		Name:     "object",
		Required: true,
		Stored:   true,
		SubType:  "policies_list",
		Type:     "external",
	},
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"propagationhidden": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PropagationHidden",
		Description: `If set to true while the policy is propagating, it won't be visible to children
namespace, but still used for policy resolution.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "propagationHidden",
		Orderable:  true,
		Setter:     true,
		Stored:     true,
		Type:       "boolean",
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
	"subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Subject",
		Description: `subject contains the tag expression that an a processing unit
must match in order to
trigger the hook.`,
		Exposed:  true,
		Name:     "subject",
		Required: true,
		Stored:   true,
		SubType:  "policies_list",
		Type:     "external",
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
