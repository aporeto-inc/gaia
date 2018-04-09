package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/constants"
	"time"
)

// RenderedPolicyIdentity represents the Identity of the object.
var RenderedPolicyIdentity = elemental.Identity{
	Name:     "renderedpolicy",
	Category: "renderedpolicies",
	Private:  false,
}

// RenderedPoliciesList represents a list of RenderedPolicies
type RenderedPoliciesList []*RenderedPolicy

// ContentIdentity returns the identity of the objects in the list.
func (o RenderedPoliciesList) ContentIdentity() elemental.Identity {

	return RenderedPolicyIdentity
}

// Copy returns a pointer to a copy the RenderedPoliciesList.
func (o RenderedPoliciesList) Copy() elemental.ContentIdentifiable {

	copy := append(RenderedPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RenderedPoliciesList.
func (o RenderedPoliciesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(RenderedPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*RenderedPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RenderedPoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RenderedPoliciesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o RenderedPoliciesList) Version() int {

	return 1
}

// RenderedPolicy represents the model of a renderedpolicy
type RenderedPolicy struct {
	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Certificate is the certificate associated with this PU. It will identify the PU
	// to any internal or external services.
	Certificate string `json:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// CreatedTime is the time at which the object was created.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// EgressPolicies lists all the egress policies attached to processing unit.
	EgressPolicies map[string]PolicyRulesList `json:"egressPolicies" bson:"-" mapstructure:"egressPolicies,omitempty"`

	// ExposedAPIServices is the list of services that this processing unit is
	// implementing.
	ExposedAPIServices APIServicesList `json:"exposedAPIServices" bson:"-" mapstructure:"exposedAPIServices,omitempty"`

	// IngressPolicies lists all the ingress policies attached to processing unit.
	IngressPolicies map[string]PolicyRulesList `json:"ingressPolicies" bson:"-" mapstructure:"ingressPolicies,omitempty"`

	// MatchingTags contains the list of tags that matched the policies.
	MatchingTags []string `json:"matchingTags" bson:"-" mapstructure:"matchingTags,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Can be set during a POST operation to render a policy on a Processing Unit that
	// has not been created yet.
	ProcessingUnit *ProcessingUnit `json:"processingUnit" bson:"-" mapstructure:"processingUnit,omitempty"`

	// Identifier of the processing unit.
	ProcessingUnitID string `json:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Profile is the trust profile of the processing unit that should be used during
	// all communications.
	Profile map[string]string `json:"profile" bson:"-" mapstructure:"profile,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Scopes is the set of scopes granted to this processing unit that it has to
	// present in HTTP requests.
	Scopes []string `json:"scopes" bson:"scopes" mapstructure:"scopes,omitempty"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewRenderedPolicy returns a new *RenderedPolicy
func NewRenderedPolicy() *RenderedPolicy {

	return &RenderedPolicy{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		EgressPolicies: map[string]PolicyRulesList{
			string(constants.RenderedPolicyTypeNetwork):   PolicyRulesList{},
			string(constants.RenderedPolicyTypeFile):      PolicyRulesList{},
			string(constants.RenderedPolicyTypeIsolation): PolicyRulesList{},
		},
		ExposedAPIServices: APIServicesList{},
		IngressPolicies: map[string]PolicyRulesList{
			string(constants.RenderedPolicyTypeNetwork):   PolicyRulesList{},
			string(constants.RenderedPolicyTypeFile):      PolicyRulesList{},
			string(constants.RenderedPolicyTypeIsolation): PolicyRulesList{},
		},
		NormalizedTags: []string{},
		Scopes:         []string{},
	}
}

// Identity returns the Identity of the object.
func (o *RenderedPolicy) Identity() elemental.Identity {

	return RenderedPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RenderedPolicy) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RenderedPolicy) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *RenderedPolicy) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *RenderedPolicy) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *RenderedPolicy) Doc() string {
	return `Retrieve the aggregated policies applied to a particular processing unit.`
}

func (o *RenderedPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *RenderedPolicy) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the given Annotations of the receiver.
func (o *RenderedPolicy) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *RenderedPolicy) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *RenderedPolicy) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *RenderedPolicy) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the given CreateTime of the receiver.
func (o *RenderedPolicy) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetNamespace returns the Namespace of the receiver.
func (o *RenderedPolicy) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the given Namespace of the receiver.
func (o *RenderedPolicy) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *RenderedPolicy) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the given NormalizedTags of the receiver.
func (o *RenderedPolicy) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetProtected returns the Protected of the receiver.
func (o *RenderedPolicy) GetProtected() bool {

	return o.Protected
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *RenderedPolicy) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the given UpdateTime of the receiver.
func (o *RenderedPolicy) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *RenderedPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("processingUnit", o.ProcessingUnit); err != nil {
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
func (*RenderedPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RenderedPolicyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RenderedPolicyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RenderedPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RenderedPolicyAttributesMap
}

// RenderedPolicyAttributesMap represents the map of attribute for RenderedPolicy.
var RenderedPolicyAttributesMap = map[string]elemental.AttributeSpecification{
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
		Description: `Certificate is the certificate associated with this PU. It will identify the PU
to any internal or external services.`,
		Exposed:  true,
		Format:   "free",
		Name:     "certificate",
		ReadOnly: true,
		Type:     "string",
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
	"EgressPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "EgressPolicies",
		Description:    `EgressPolicies lists all the egress policies attached to processing unit.`,
		Exposed:        true,
		Name:           "egressPolicies",
		ReadOnly:       true,
		SubType:        "rendered_policy",
		Type:           "external",
	},
	"ExposedAPIServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedAPIServices",
		Description: `ExposedAPIServices is the list of services that this processing unit is
implementing.`,
		Exposed: true,
		Name:    "exposedAPIServices",
		SubType: "api_services_entities",
		Type:    "external",
	},
	"IngressPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IngressPolicies",
		Description:    `IngressPolicies lists all the ingress policies attached to processing unit.`,
		Exposed:        true,
		Name:           "ingressPolicies",
		ReadOnly:       true,
		SubType:        "rendered_policy",
		Type:           "external",
	},
	"MatchingTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "MatchingTags",
		Description:    `MatchingTags contains the list of tags that matched the policies.`,
		Exposed:        true,
		Name:           "matchingTags",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
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
	"ProcessingUnit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnit",
		CreationOnly:   true,
		Description: `Can be set during a POST operation to render a policy on a Processing Unit that
has not been created yet.`,
		Exposed:  true,
		Name:     "processingUnit",
		Required: true,
		SubType:  "processingunit",
		Type:     "external",
	},
	"ProcessingUnitID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ProcessingUnitID",
		Description:    `Identifier of the processing unit.`,
		Exposed:        true,
		Format:         "free",
		Name:           "processingUnitID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Profile": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Profile",
		Description: `Profile is the trust profile of the processing unit that should be used during
all communications.`,
		Exposed:  true,
		Name:     "profile",
		ReadOnly: true,
		SubType:  "trust_profile",
		Type:     "external",
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
	"Scopes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Scopes",
		Description: `Scopes is the set of scopes granted to this processing unit that it has to
present in HTTP requests.`,
		Exposed: true,
		Name:    "scopes",
		Stored:  true,
		SubType: "scopes_list",
		Type:    "external",
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

// RenderedPolicyLowerCaseAttributesMap represents the map of attribute for RenderedPolicy.
var RenderedPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Description: `Certificate is the certificate associated with this PU. It will identify the PU
to any internal or external services.`,
		Exposed:  true,
		Format:   "free",
		Name:     "certificate",
		ReadOnly: true,
		Type:     "string",
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
	"egresspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "EgressPolicies",
		Description:    `EgressPolicies lists all the egress policies attached to processing unit.`,
		Exposed:        true,
		Name:           "egressPolicies",
		ReadOnly:       true,
		SubType:        "rendered_policy",
		Type:           "external",
	},
	"exposedapiservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedAPIServices",
		Description: `ExposedAPIServices is the list of services that this processing unit is
implementing.`,
		Exposed: true,
		Name:    "exposedAPIServices",
		SubType: "api_services_entities",
		Type:    "external",
	},
	"ingresspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IngressPolicies",
		Description:    `IngressPolicies lists all the ingress policies attached to processing unit.`,
		Exposed:        true,
		Name:           "ingressPolicies",
		ReadOnly:       true,
		SubType:        "rendered_policy",
		Type:           "external",
	},
	"matchingtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "MatchingTags",
		Description:    `MatchingTags contains the list of tags that matched the policies.`,
		Exposed:        true,
		Name:           "matchingTags",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
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
	"processingunit": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnit",
		CreationOnly:   true,
		Description: `Can be set during a POST operation to render a policy on a Processing Unit that
has not been created yet.`,
		Exposed:  true,
		Name:     "processingUnit",
		Required: true,
		SubType:  "processingunit",
		Type:     "external",
	},
	"processingunitid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ProcessingUnitID",
		Description:    `Identifier of the processing unit.`,
		Exposed:        true,
		Format:         "free",
		Name:           "processingUnitID",
		ReadOnly:       true,
		Type:           "string",
	},
	"profile": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Profile",
		Description: `Profile is the trust profile of the processing unit that should be used during
all communications.`,
		Exposed:  true,
		Name:     "profile",
		ReadOnly: true,
		SubType:  "trust_profile",
		Type:     "external",
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
	"scopes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Scopes",
		Description: `Scopes is the set of scopes granted to this processing unit that it has to
present in HTTP requests.`,
		Exposed: true,
		Name:    "scopes",
		Stored:  true,
		SubType: "scopes_list",
		Type:    "external",
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
