package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/constants"
)

// RenderedPolicyIdentity represents the Identity of the object.
var RenderedPolicyIdentity = elemental.Identity{
	Name:     "renderedpolicy",
	Category: "renderedpolicies",
	Package:  "squall",
	Private:  false,
}

// RenderedPoliciesList represents a list of RenderedPolicies
type RenderedPoliciesList []*RenderedPolicy

// Identity returns the identity of the objects in the list.
func (o RenderedPoliciesList) Identity() elemental.Identity {

	return RenderedPolicyIdentity
}

// Copy returns a pointer to a copy the RenderedPoliciesList.
func (o RenderedPoliciesList) Copy() elemental.Identifiables {

	copy := append(RenderedPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RenderedPoliciesList.
func (o RenderedPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RenderedPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*RenderedPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RenderedPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RenderedPoliciesList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the RenderedPoliciesList converted to SparseRenderedPoliciesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o RenderedPoliciesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o RenderedPoliciesList) Version() int {

	return 1
}

// RenderedPolicy represents the model of a renderedpolicy
type RenderedPolicy struct {
	// Certificate is the certificate associated with this PU. It will identify the PU
	// to any internal or external services.
	Certificate string `json:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// DependendServices is the list of services that this processing unit depends on.
	DependendServices ServicesList `json:"dependendServices" bson:"-" mapstructure:"dependendServices,omitempty"`

	// EgressPolicies lists all the egress policies attached to processing unit.
	EgressPolicies map[string]PolicyRulesList `json:"egressPolicies" bson:"-" mapstructure:"egressPolicies,omitempty"`

	// ExposedServices is the list of services that this processing unit is
	// implementing.
	ExposedServices ServicesList `json:"exposedServices" bson:"-" mapstructure:"exposedServices,omitempty"`

	// hashedTags contains the list of tags that matched the policies and their hashes.
	HashedTags map[string]string `json:"hashedTags" bson:"-" mapstructure:"hashedTags,omitempty"`

	// IngressPolicies lists all the ingress policies attached to processing unit.
	IngressPolicies map[string]PolicyRulesList `json:"ingressPolicies" bson:"-" mapstructure:"ingressPolicies,omitempty"`

	// MatchingTags contains the list of tags that matched the policies.
	MatchingTags []string `json:"matchingTags" bson:"-" mapstructure:"matchingTags,omitempty"`

	// Can be set during a POST operation to render a policy on a Processing Unit that
	// has not been created yet.
	ProcessingUnit *ProcessingUnit `json:"processingUnit" bson:"-" mapstructure:"processingUnit,omitempty"`

	// Identifier of the processing unit.
	ProcessingUnitID string `json:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Profile is the trust profile of the processing unit that should be used during
	// all communications.
	Profile map[string]string `json:"profile" bson:"-" mapstructure:"profile,omitempty"`

	// Scopes is the set of scopes granted to this processing unit that it has to
	// present in HTTP requests.
	Scopes []string `json:"scopes" bson:"scopes" mapstructure:"scopes,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewRenderedPolicy returns a new *RenderedPolicy
func NewRenderedPolicy() *RenderedPolicy {

	return &RenderedPolicy{
		ModelVersion:      1,
		DependendServices: ServicesList{},
		EgressPolicies: map[string]PolicyRulesList{
			string(constants.RenderedPolicyTypeNetwork):   PolicyRulesList{},
			string(constants.RenderedPolicyTypeFile):      PolicyRulesList{},
			string(constants.RenderedPolicyTypeIsolation): PolicyRulesList{},
		},
		ExposedServices: ServicesList{},
		HashedTags:      map[string]string{},
		IngressPolicies: map[string]PolicyRulesList{
			string(constants.RenderedPolicyTypeNetwork):   PolicyRulesList{},
			string(constants.RenderedPolicyTypeFile):      PolicyRulesList{},
			string(constants.RenderedPolicyTypeIsolation): PolicyRulesList{},
		},
		Scopes: []string{},
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

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *RenderedPolicy) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseRenderedPolicy{
			Certificate:       &o.Certificate,
			DependendServices: &o.DependendServices,
			EgressPolicies:    &o.EgressPolicies,
			ExposedServices:   &o.ExposedServices,
			HashedTags:        &o.HashedTags,
			IngressPolicies:   &o.IngressPolicies,
			MatchingTags:      &o.MatchingTags,
			ProcessingUnit:    &o.ProcessingUnit,
			ProcessingUnitID:  &o.ProcessingUnitID,
			Profile:           &o.Profile,
			Scopes:            &o.Scopes,
		}
	}

	sp := &SparseRenderedPolicy{}
	for _, f := range fields {
		switch f {
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "dependendServices":
			sp.DependendServices = &(o.DependendServices)
		case "egressPolicies":
			sp.EgressPolicies = &(o.EgressPolicies)
		case "exposedServices":
			sp.ExposedServices = &(o.ExposedServices)
		case "hashedTags":
			sp.HashedTags = &(o.HashedTags)
		case "ingressPolicies":
			sp.IngressPolicies = &(o.IngressPolicies)
		case "matchingTags":
			sp.MatchingTags = &(o.MatchingTags)
		case "processingUnit":
			sp.ProcessingUnit = &(o.ProcessingUnit)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "profile":
			sp.Profile = &(o.Profile)
		case "scopes":
			sp.Scopes = &(o.Scopes)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseRenderedPolicy to the object.
func (o *RenderedPolicy) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseRenderedPolicy)
	if so.Certificate != nil {
		o.Certificate = *so.Certificate
	}
	if so.DependendServices != nil {
		o.DependendServices = *so.DependendServices
	}
	if so.EgressPolicies != nil {
		o.EgressPolicies = *so.EgressPolicies
	}
	if so.ExposedServices != nil {
		o.ExposedServices = *so.ExposedServices
	}
	if so.HashedTags != nil {
		o.HashedTags = *so.HashedTags
	}
	if so.IngressPolicies != nil {
		o.IngressPolicies = *so.IngressPolicies
	}
	if so.MatchingTags != nil {
		o.MatchingTags = *so.MatchingTags
	}
	if so.ProcessingUnit != nil {
		o.ProcessingUnit = *so.ProcessingUnit
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.Profile != nil {
		o.Profile = *so.Profile
	}
	if so.Scopes != nil {
		o.Scopes = *so.Scopes
	}
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
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description: `Certificate is the certificate associated with this PU. It will identify the PU
to any internal or external services.`,
		Exposed:  true,
		Name:     "certificate",
		ReadOnly: true,
		Type:     "string",
	},
	"DependendServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DependendServices",
		Description:    `DependendServices is the list of services that this processing unit depends on.`,
		Exposed:        true,
		Name:           "dependendServices",
		SubType:        "api_services_entities",
		Type:           "external",
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
	"ExposedServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedServices",
		Description: `ExposedServices is the list of services that this processing unit is
implementing.`,
		Exposed: true,
		Name:    "exposedServices",
		SubType: "api_services_entities",
		Type:    "external",
	},
	"HashedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "HashedTags",
		Description:    `hashedTags contains the list of tags that matched the policies and their hashes.`,
		Exposed:        true,
		Name:           "hashedTags",
		ReadOnly:       true,
		SubType:        "hashed_tags",
		Type:           "external",
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
}

// RenderedPolicyLowerCaseAttributesMap represents the map of attribute for RenderedPolicy.
var RenderedPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description: `Certificate is the certificate associated with this PU. It will identify the PU
to any internal or external services.`,
		Exposed:  true,
		Name:     "certificate",
		ReadOnly: true,
		Type:     "string",
	},
	"dependendservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DependendServices",
		Description:    `DependendServices is the list of services that this processing unit depends on.`,
		Exposed:        true,
		Name:           "dependendServices",
		SubType:        "api_services_entities",
		Type:           "external",
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
	"exposedservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedServices",
		Description: `ExposedServices is the list of services that this processing unit is
implementing.`,
		Exposed: true,
		Name:    "exposedServices",
		SubType: "api_services_entities",
		Type:    "external",
	},
	"hashedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "HashedTags",
		Description:    `hashedTags contains the list of tags that matched the policies and their hashes.`,
		Exposed:        true,
		Name:           "hashedTags",
		ReadOnly:       true,
		SubType:        "hashed_tags",
		Type:           "external",
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
}

// SparseRenderedPoliciesList represents a list of SparseRenderedPolicies
type SparseRenderedPoliciesList []*SparseRenderedPolicy

// Identity returns the identity of the objects in the list.
func (o SparseRenderedPoliciesList) Identity() elemental.Identity {

	return RenderedPolicyIdentity
}

// Copy returns a pointer to a copy the SparseRenderedPoliciesList.
func (o SparseRenderedPoliciesList) Copy() elemental.Identifiables {

	copy := append(SparseRenderedPoliciesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseRenderedPoliciesList.
func (o SparseRenderedPoliciesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseRenderedPoliciesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseRenderedPolicy))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseRenderedPoliciesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseRenderedPoliciesList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseRenderedPoliciesList converted to RenderedPoliciesList.
func (o SparseRenderedPoliciesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseRenderedPoliciesList) Version() int {

	return 1
}

// SparseRenderedPolicy represents the sparse version of a renderedpolicy.
type SparseRenderedPolicy struct {
	// Certificate is the certificate associated with this PU. It will identify the PU
	// to any internal or external services.
	Certificate *string `json:"certificate,omitempty" bson:"-" mapstructure:"certificate,omitempty"`

	// DependendServices is the list of services that this processing unit depends on.
	DependendServices *ServicesList `json:"dependendServices,omitempty" bson:"-" mapstructure:"dependendServices,omitempty"`

	// EgressPolicies lists all the egress policies attached to processing unit.
	EgressPolicies *map[string]PolicyRulesList `json:"egressPolicies,omitempty" bson:"-" mapstructure:"egressPolicies,omitempty"`

	// ExposedServices is the list of services that this processing unit is
	// implementing.
	ExposedServices *ServicesList `json:"exposedServices,omitempty" bson:"-" mapstructure:"exposedServices,omitempty"`

	// hashedTags contains the list of tags that matched the policies and their hashes.
	HashedTags *map[string]string `json:"hashedTags,omitempty" bson:"-" mapstructure:"hashedTags,omitempty"`

	// IngressPolicies lists all the ingress policies attached to processing unit.
	IngressPolicies *map[string]PolicyRulesList `json:"ingressPolicies,omitempty" bson:"-" mapstructure:"ingressPolicies,omitempty"`

	// MatchingTags contains the list of tags that matched the policies.
	MatchingTags *[]string `json:"matchingTags,omitempty" bson:"-" mapstructure:"matchingTags,omitempty"`

	// Can be set during a POST operation to render a policy on a Processing Unit that
	// has not been created yet.
	ProcessingUnit **ProcessingUnit `json:"processingUnit,omitempty" bson:"-" mapstructure:"processingUnit,omitempty"`

	// Identifier of the processing unit.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Profile is the trust profile of the processing unit that should be used during
	// all communications.
	Profile *map[string]string `json:"profile,omitempty" bson:"-" mapstructure:"profile,omitempty"`

	// Scopes is the set of scopes granted to this processing unit that it has to
	// present in HTTP requests.
	Scopes *[]string `json:"scopes,omitempty" bson:"scopes" mapstructure:"scopes,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseRenderedPolicy returns a new  SparseRenderedPolicy.
func NewSparseRenderedPolicy() *SparseRenderedPolicy {
	return &SparseRenderedPolicy{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseRenderedPolicy) Identity() elemental.Identity {

	return RenderedPolicyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseRenderedPolicy) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseRenderedPolicy) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseRenderedPolicy) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseRenderedPolicy) ToPlain() elemental.PlainIdentifiable {

	out := NewRenderedPolicy()
	if o.Certificate != nil {
		out.Certificate = *o.Certificate
	}
	if o.DependendServices != nil {
		out.DependendServices = *o.DependendServices
	}
	if o.EgressPolicies != nil {
		out.EgressPolicies = *o.EgressPolicies
	}
	if o.ExposedServices != nil {
		out.ExposedServices = *o.ExposedServices
	}
	if o.HashedTags != nil {
		out.HashedTags = *o.HashedTags
	}
	if o.IngressPolicies != nil {
		out.IngressPolicies = *o.IngressPolicies
	}
	if o.MatchingTags != nil {
		out.MatchingTags = *o.MatchingTags
	}
	if o.ProcessingUnit != nil {
		out.ProcessingUnit = *o.ProcessingUnit
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.Profile != nil {
		out.Profile = *o.Profile
	}
	if o.Scopes != nil {
		out.Scopes = *o.Scopes
	}

	return out
}
