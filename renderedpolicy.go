package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/constants"
)

// RenderedPolicyDatapathTypeValue represents the possible values for attribute "datapathType".
type RenderedPolicyDatapathTypeValue string

const (
	// RenderedPolicyDatapathTypeAporeto represents the value Aporeto.
	RenderedPolicyDatapathTypeAporeto RenderedPolicyDatapathTypeValue = "Aporeto"

	// RenderedPolicyDatapathTypeDefault represents the value Default.
	RenderedPolicyDatapathTypeDefault RenderedPolicyDatapathTypeValue = "Default"

	// RenderedPolicyDatapathTypeEnvoyAuthorizer represents the value EnvoyAuthorizer.
	RenderedPolicyDatapathTypeEnvoyAuthorizer RenderedPolicyDatapathTypeValue = "EnvoyAuthorizer"
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
func (o RenderedPoliciesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseRenderedPoliciesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseRenderedPolicy)
	}

	return out
}

// Version returns the version of the content.
func (o RenderedPoliciesList) Version() int {

	return 1
}

// RenderedPolicy represents the model of a renderedpolicy
type RenderedPolicy struct {
	// The certificate associated with this processing unit. It will identify the
	// processing unit to any internal or external services.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// The datapath type that this processing unit must implement according to
	// the rendered policy:
	// - `Default`: This policy is not making a decision for the datapath.
	// - `Aporeto`: The enforcer is managing and handling the datapath.
	// - `EnvoyAuthorizer`: The enforcer is serving envoy compatible gRPC APIs
	// that for example can be used by an envoy proxy to use the Aporeto PKI
	// and implement Aporeto network access policies. NOTE: The enforcer is not
	// owning the datapath in this case. It is merely providing an authorizer API.
	DatapathType RenderedPolicyDatapathTypeValue `json:"datapathType" msgpack:"datapathType" bson:"-" mapstructure:"datapathType,omitempty"`

	// The list of services that this processing unit depends on.
	DependendServices ServicesList `json:"dependendServices" msgpack:"dependendServices" bson:"-" mapstructure:"dependendServices,omitempty"`

	// Lists all the egress policies attached to processing unit.
	EgressPolicies map[string]PolicyRulesList `json:"egressPolicies" msgpack:"egressPolicies" bson:"-" mapstructure:"egressPolicies,omitempty"`

	// The list of services that this processing unit is implementing.
	ExposedServices ServicesList `json:"exposedServices" msgpack:"exposedServices" bson:"-" mapstructure:"exposedServices,omitempty"`

	// Contains the list of tags that matched the policies and their hashes.
	HashedTags map[string]string `json:"hashedTags" msgpack:"hashedTags" bson:"-" mapstructure:"hashedTags,omitempty"`

	// Lists all the ingress policies attached to the processing unit.
	IngressPolicies map[string]PolicyRulesList `json:"ingressPolicies" msgpack:"ingressPolicies" bson:"-" mapstructure:"ingressPolicies,omitempty"`

	// Contains the list of tags that matched the policies.
	MatchingTags []string `json:"matchingTags" msgpack:"matchingTags" bson:"-" mapstructure:"matchingTags,omitempty"`

	// Can be set during a `POST` operation to render a policy on a processing unit
	// that
	// has not been created yet.
	ProcessingUnit *ProcessingUnit `json:"processingUnit" msgpack:"processingUnit" bson:"-" mapstructure:"processingUnit,omitempty"`

	// Identifier of the processing unit.
	ProcessingUnitID string `json:"processingUnitID" msgpack:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// The set of scopes granted to this processing unit that has to be
	// present in HTTP requests.
	Scopes []string `json:"scopes" msgpack:"scopes" bson:"scopes" mapstructure:"scopes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRenderedPolicy returns a new *RenderedPolicy
func NewRenderedPolicy() *RenderedPolicy {

	return &RenderedPolicy{
		ModelVersion: 1,
		EgressPolicies: map[string]PolicyRulesList{
			string(constants.RenderedPolicyTypeNetwork):   PolicyRulesList{},
			string(constants.RenderedPolicyTypeFile):      PolicyRulesList{},
			string(constants.RenderedPolicyTypeIsolation): PolicyRulesList{},
		},
		HashedTags: map[string]string{},
		IngressPolicies: map[string]PolicyRulesList{
			string(constants.RenderedPolicyTypeNetwork):   PolicyRulesList{},
			string(constants.RenderedPolicyTypeFile):      PolicyRulesList{},
			string(constants.RenderedPolicyTypeIsolation): PolicyRulesList{},
		},
		MatchingTags: []string{},
		Scopes:       []string{},
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

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RenderedPolicy) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRenderedPolicy{}

	s.Scopes = o.Scopes

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RenderedPolicy) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesRenderedPolicy{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Scopes = s.Scopes

	return nil
}

// Version returns the hardcoded version of the model.
func (o *RenderedPolicy) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *RenderedPolicy) BleveType() string {

	return "renderedpolicy"
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
			DatapathType:      &o.DatapathType,
			DependendServices: &o.DependendServices,
			EgressPolicies:    &o.EgressPolicies,
			ExposedServices:   &o.ExposedServices,
			HashedTags:        &o.HashedTags,
			IngressPolicies:   &o.IngressPolicies,
			MatchingTags:      &o.MatchingTags,
			ProcessingUnit:    o.ProcessingUnit,
			ProcessingUnitID:  &o.ProcessingUnitID,
			Scopes:            &o.Scopes,
		}
	}

	sp := &SparseRenderedPolicy{}
	for _, f := range fields {
		switch f {
		case "certificate":
			sp.Certificate = &(o.Certificate)
		case "datapathType":
			sp.DatapathType = &(o.DatapathType)
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
			sp.ProcessingUnit = o.ProcessingUnit
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
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
	if so.DatapathType != nil {
		o.DatapathType = *so.DatapathType
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
		o.ProcessingUnit = so.ProcessingUnit
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.Scopes != nil {
		o.Scopes = *so.Scopes
	}
}

// DeepCopy returns a deep copy if the RenderedPolicy.
func (o *RenderedPolicy) DeepCopy() *RenderedPolicy {

	if o == nil {
		return nil
	}

	out := &RenderedPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *RenderedPolicy.
func (o *RenderedPolicy) DeepCopyInto(out *RenderedPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy RenderedPolicy: %s", err))
	}

	*out = *target.(*RenderedPolicy)
}

// Validate valides the current information stored into the structure.
func (o *RenderedPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("datapathType", string(o.DatapathType), []string{"Default", "Aporeto", "EnvoyAuthorizer"}, true); err != nil {
		errors = errors.Append(err)
	}

	for _, sub := range o.DependendServices {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.ExposedServices {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if o.ProcessingUnit != nil {
		elemental.ResetDefaultForZeroValues(o.ProcessingUnit)
		if err := o.ProcessingUnit.Validate(); err != nil {
			errors = errors.Append(err)
		}
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *RenderedPolicy) ValueForAttribute(name string) interface{} {

	switch name {
	case "certificate":
		return o.Certificate
	case "datapathType":
		return o.DatapathType
	case "dependendServices":
		return o.DependendServices
	case "egressPolicies":
		return o.EgressPolicies
	case "exposedServices":
		return o.ExposedServices
	case "hashedTags":
		return o.HashedTags
	case "ingressPolicies":
		return o.IngressPolicies
	case "matchingTags":
		return o.MatchingTags
	case "processingUnit":
		return o.ProcessingUnit
	case "processingUnitID":
		return o.ProcessingUnitID
	case "scopes":
		return o.Scopes
	}

	return nil
}

// RenderedPolicyAttributesMap represents the map of attribute for RenderedPolicy.
var RenderedPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"Certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description: `The certificate associated with this processing unit. It will identify the
processing unit to any internal or external services.`,
		Exposed:  true,
		Name:     "certificate",
		ReadOnly: true,
		Type:     "string",
	},
	"DatapathType": elemental.AttributeSpecification{
		AllowedChoices: []string{"Default", "Aporeto", "EnvoyAuthorizer"},
		Autogenerated:  true,
		ConvertedName:  "DatapathType",
		Description: `The datapath type that this processing unit must implement according to
the rendered policy:
- ` + "`" + `Default` + "`" + `: This policy is not making a decision for the datapath.
- ` + "`" + `Aporeto` + "`" + `: The enforcer is managing and handling the datapath.
- ` + "`" + `EnvoyAuthorizer` + "`" + `: The enforcer is serving envoy compatible gRPC APIs
that for example can be used by an envoy proxy to use the Aporeto PKI
and implement Aporeto network access policies. NOTE: The enforcer is not
owning the datapath in this case. It is merely providing an authorizer API.`,
		Exposed:  true,
		Name:     "datapathType",
		ReadOnly: true,
		Type:     "enum",
	},
	"DependendServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DependendServices",
		Description:    `The list of services that this processing unit depends on.`,
		Exposed:        true,
		Name:           "dependendServices",
		SubType:        "service",
		Type:           "refList",
	},
	"EgressPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "EgressPolicies",
		Description:    `Lists all the egress policies attached to processing unit.`,
		Exposed:        true,
		Name:           "egressPolicies",
		ReadOnly:       true,
		SubType:        "_rendered_policy",
		Type:           "external",
	},
	"ExposedServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedServices",
		Description:    `The list of services that this processing unit is implementing.`,
		Exposed:        true,
		Name:           "exposedServices",
		SubType:        "service",
		Type:           "refList",
	},
	"HashedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "HashedTags",
		Description:    `Contains the list of tags that matched the policies and their hashes.`,
		Exposed:        true,
		Name:           "hashedTags",
		ReadOnly:       true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"IngressPolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IngressPolicies",
		Description:    `Lists all the ingress policies attached to the processing unit.`,
		Exposed:        true,
		Name:           "ingressPolicies",
		ReadOnly:       true,
		SubType:        "_rendered_policy",
		Type:           "external",
	},
	"MatchingTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "MatchingTags",
		Description:    `Contains the list of tags that matched the policies.`,
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
		Description: `Can be set during a ` + "`" + `POST` + "`" + ` operation to render a policy on a processing unit
that
has not been created yet.`,
		Exposed:  true,
		Name:     "processingUnit",
		Required: true,
		SubType:  "processingunit",
		Type:     "ref",
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
	"Scopes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Scopes",
		Description: `The set of scopes granted to this processing unit that has to be
present in HTTP requests.`,
		Exposed: true,
		Name:    "scopes",
		Stored:  true,
		SubType: "string",
		Type:    "list",
	},
}

// RenderedPolicyLowerCaseAttributesMap represents the map of attribute for RenderedPolicy.
var RenderedPolicyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"certificate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Certificate",
		Description: `The certificate associated with this processing unit. It will identify the
processing unit to any internal or external services.`,
		Exposed:  true,
		Name:     "certificate",
		ReadOnly: true,
		Type:     "string",
	},
	"datapathtype": elemental.AttributeSpecification{
		AllowedChoices: []string{"Default", "Aporeto", "EnvoyAuthorizer"},
		Autogenerated:  true,
		ConvertedName:  "DatapathType",
		Description: `The datapath type that this processing unit must implement according to
the rendered policy:
- ` + "`" + `Default` + "`" + `: This policy is not making a decision for the datapath.
- ` + "`" + `Aporeto` + "`" + `: The enforcer is managing and handling the datapath.
- ` + "`" + `EnvoyAuthorizer` + "`" + `: The enforcer is serving envoy compatible gRPC APIs
that for example can be used by an envoy proxy to use the Aporeto PKI
and implement Aporeto network access policies. NOTE: The enforcer is not
owning the datapath in this case. It is merely providing an authorizer API.`,
		Exposed:  true,
		Name:     "datapathType",
		ReadOnly: true,
		Type:     "enum",
	},
	"dependendservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DependendServices",
		Description:    `The list of services that this processing unit depends on.`,
		Exposed:        true,
		Name:           "dependendServices",
		SubType:        "service",
		Type:           "refList",
	},
	"egresspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "EgressPolicies",
		Description:    `Lists all the egress policies attached to processing unit.`,
		Exposed:        true,
		Name:           "egressPolicies",
		ReadOnly:       true,
		SubType:        "_rendered_policy",
		Type:           "external",
	},
	"exposedservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExposedServices",
		Description:    `The list of services that this processing unit is implementing.`,
		Exposed:        true,
		Name:           "exposedServices",
		SubType:        "service",
		Type:           "refList",
	},
	"hashedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "HashedTags",
		Description:    `Contains the list of tags that matched the policies and their hashes.`,
		Exposed:        true,
		Name:           "hashedTags",
		ReadOnly:       true,
		SubType:        "map[string]string",
		Type:           "external",
	},
	"ingresspolicies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "IngressPolicies",
		Description:    `Lists all the ingress policies attached to the processing unit.`,
		Exposed:        true,
		Name:           "ingressPolicies",
		ReadOnly:       true,
		SubType:        "_rendered_policy",
		Type:           "external",
	},
	"matchingtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "MatchingTags",
		Description:    `Contains the list of tags that matched the policies.`,
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
		Description: `Can be set during a ` + "`" + `POST` + "`" + ` operation to render a policy on a processing unit
that
has not been created yet.`,
		Exposed:  true,
		Name:     "processingUnit",
		Required: true,
		SubType:  "processingunit",
		Type:     "ref",
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
	"scopes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Scopes",
		Description: `The set of scopes granted to this processing unit that has to be
present in HTTP requests.`,
		Exposed: true,
		Name:    "scopes",
		Stored:  true,
		SubType: "string",
		Type:    "list",
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
	// The certificate associated with this processing unit. It will identify the
	// processing unit to any internal or external services.
	Certificate *string `json:"certificate,omitempty" msgpack:"certificate,omitempty" bson:"-" mapstructure:"certificate,omitempty"`

	// The datapath type that this processing unit must implement according to
	// the rendered policy:
	// - `Default`: This policy is not making a decision for the datapath.
	// - `Aporeto`: The enforcer is managing and handling the datapath.
	// - `EnvoyAuthorizer`: The enforcer is serving envoy compatible gRPC APIs
	// that for example can be used by an envoy proxy to use the Aporeto PKI
	// and implement Aporeto network access policies. NOTE: The enforcer is not
	// owning the datapath in this case. It is merely providing an authorizer API.
	DatapathType *RenderedPolicyDatapathTypeValue `json:"datapathType,omitempty" msgpack:"datapathType,omitempty" bson:"-" mapstructure:"datapathType,omitempty"`

	// The list of services that this processing unit depends on.
	DependendServices *ServicesList `json:"dependendServices,omitempty" msgpack:"dependendServices,omitempty" bson:"-" mapstructure:"dependendServices,omitempty"`

	// Lists all the egress policies attached to processing unit.
	EgressPolicies *map[string]PolicyRulesList `json:"egressPolicies,omitempty" msgpack:"egressPolicies,omitempty" bson:"-" mapstructure:"egressPolicies,omitempty"`

	// The list of services that this processing unit is implementing.
	ExposedServices *ServicesList `json:"exposedServices,omitempty" msgpack:"exposedServices,omitempty" bson:"-" mapstructure:"exposedServices,omitempty"`

	// Contains the list of tags that matched the policies and their hashes.
	HashedTags *map[string]string `json:"hashedTags,omitempty" msgpack:"hashedTags,omitempty" bson:"-" mapstructure:"hashedTags,omitempty"`

	// Lists all the ingress policies attached to the processing unit.
	IngressPolicies *map[string]PolicyRulesList `json:"ingressPolicies,omitempty" msgpack:"ingressPolicies,omitempty" bson:"-" mapstructure:"ingressPolicies,omitempty"`

	// Contains the list of tags that matched the policies.
	MatchingTags *[]string `json:"matchingTags,omitempty" msgpack:"matchingTags,omitempty" bson:"-" mapstructure:"matchingTags,omitempty"`

	// Can be set during a `POST` operation to render a policy on a processing unit
	// that
	// has not been created yet.
	ProcessingUnit *ProcessingUnit `json:"processingUnit,omitempty" msgpack:"processingUnit,omitempty" bson:"-" mapstructure:"processingUnit,omitempty"`

	// Identifier of the processing unit.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// The set of scopes granted to this processing unit that has to be
	// present in HTTP requests.
	Scopes *[]string `json:"scopes,omitempty" msgpack:"scopes,omitempty" bson:"scopes,omitempty" mapstructure:"scopes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
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

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRenderedPolicy) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseRenderedPolicy{}

	if o.Scopes != nil {
		s.Scopes = o.Scopes
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseRenderedPolicy) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseRenderedPolicy{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.Scopes != nil {
		o.Scopes = s.Scopes
	}

	return nil
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
	if o.DatapathType != nil {
		out.DatapathType = *o.DatapathType
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
		out.ProcessingUnit = o.ProcessingUnit
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.Scopes != nil {
		out.Scopes = *o.Scopes
	}

	return out
}

// DeepCopy returns a deep copy if the SparseRenderedPolicy.
func (o *SparseRenderedPolicy) DeepCopy() *SparseRenderedPolicy {

	if o == nil {
		return nil
	}

	out := &SparseRenderedPolicy{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseRenderedPolicy.
func (o *SparseRenderedPolicy) DeepCopyInto(out *SparseRenderedPolicy) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseRenderedPolicy: %s", err))
	}

	*out = *target.(*SparseRenderedPolicy)
}

type mongoAttributesRenderedPolicy struct {
	Scopes []string `bson:"scopes"`
}
type mongoAttributesSparseRenderedPolicy struct {
	Scopes *[]string `bson:"scopes,omitempty"`
}
