package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PolicyRuleIdentity represents the Identity of the object.
var PolicyRuleIdentity = elemental.Identity{
	Name:     "policyrule",
	Category: "policyrules",
	Package:  "squall",
	Private:  false,
}

// PolicyRulesList represents a list of PolicyRules
type PolicyRulesList []*PolicyRule

// Identity returns the identity of the objects in the list.
func (o PolicyRulesList) Identity() elemental.Identity {

	return PolicyRuleIdentity
}

// Copy returns a pointer to a copy the PolicyRulesList.
func (o PolicyRulesList) Copy() elemental.Identifiables {

	copy := append(PolicyRulesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PolicyRulesList.
func (o PolicyRulesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PolicyRulesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PolicyRule))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PolicyRulesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PolicyRulesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the PolicyRulesList converted to SparsePolicyRulesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PolicyRulesList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o PolicyRulesList) Version() int {

	return 1
}

// PolicyRule represents the model of a policyrule
type PolicyRule struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Action defines set of actions that must be enforced when a dependency is met.
	Action map[string]map[string]interface{} `json:"action" bson:"-" mapstructure:"action,omitempty"`

	// AuditProfiles provides the audit profiles that must be applied.
	AuditProfiles AuditProfilesList `json:"auditProfiles" bson:"-" mapstructure:"auditProfiles,omitempty"`

	// EnforcerProfiles provides the information about the server profile.
	EnforcerProfiles EnforcerProfilesList `json:"enforcerProfiles" bson:"-" mapstructure:"enforcerProfiles,omitempty"`

	// Policy target networks.
	ExternalNetworks ExternalNetworksList `json:"externalNetworks" bson:"-" mapstructure:"externalNetworks,omitempty"`

	// Policy target file paths.
	FilePaths FilePathsList `json:"filePaths" bson:"-" mapstructure:"filePaths,omitempty"`

	// HostServices provides the list of host services that must be instantiated.
	HostServices HostServicesList `json:"hostServices" bson:"-" mapstructure:"hostServices,omitempty"`

	// IsolationProfiles are the isolation profiles of the rule.
	IsolationProfiles IsolationProfilesList `json:"isolationProfiles" bson:"-" mapstructure:"isolationProfiles,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Policy target namespaces.
	Namespaces NamespacesList `json:"namespaces" bson:"-" mapstructure:"namespaces,omitempty"`

	// PolicyNamespace is the namespace of the policy that created this rule.
	PolicyNamespace string `json:"policyNamespace" bson:"-" mapstructure:"policyNamespace,omitempty"`

	// Last time the policy was updated.
	PolicyUpdateTime time.Time `json:"policyUpdateTime" bson:"-" mapstructure:"policyUpdateTime,omitempty"`

	// Propagated indicates if the policy is propagated.
	Propagated bool `json:"propagated" bson:"-" mapstructure:"propagated,omitempty"`

	// Relation describes the required operation to be performed between subjects and
	// objects.
	Relation []string `json:"relation" bson:"-" mapstructure:"relation,omitempty"`

	// Services provides the services of this policy rule.
	Services ServicesList `json:"services" bson:"-" mapstructure:"services,omitempty"`

	// Policy target tags.
	TagClauses [][]string `json:"tagClauses" bson:"-" mapstructure:"tagClauses,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewPolicyRule returns a new *PolicyRule
func NewPolicyRule() *PolicyRule {

	return &PolicyRule{
		ModelVersion:  1,
		Mutex:         &sync.Mutex{},
		HostServices:  HostServicesList{},
		AuditProfiles: AuditProfilesList{},
		Relation:      []string{},
		TagClauses:    [][]string{},
	}
}

// Identity returns the Identity of the object.
func (o *PolicyRule) Identity() elemental.Identity {

	return PolicyRuleIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyRule) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyRule) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *PolicyRule) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *PolicyRule) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *PolicyRule) Doc() string {
	return `PolicyRule is an internal policy resolution API. Services can use this API to
retrieve a policy resolution.`
}

func (o *PolicyRule) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the Name of the receiver.
func (o *PolicyRule) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *PolicyRule) SetName(name string) {

	o.Name = name
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PolicyRule) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePolicyRule{
			ID:                &o.ID,
			Action:            &o.Action,
			AuditProfiles:     &o.AuditProfiles,
			EnforcerProfiles:  &o.EnforcerProfiles,
			ExternalNetworks:  &o.ExternalNetworks,
			FilePaths:         &o.FilePaths,
			HostServices:      &o.HostServices,
			IsolationProfiles: &o.IsolationProfiles,
			Name:              &o.Name,
			Namespaces:        &o.Namespaces,
			PolicyNamespace:   &o.PolicyNamespace,
			PolicyUpdateTime:  &o.PolicyUpdateTime,
			Propagated:        &o.Propagated,
			Relation:          &o.Relation,
			Services:          &o.Services,
			TagClauses:        &o.TagClauses,
		}
	}

	sp := &SparsePolicyRule{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "action":
			sp.Action = &(o.Action)
		case "auditProfiles":
			sp.AuditProfiles = &(o.AuditProfiles)
		case "enforcerProfiles":
			sp.EnforcerProfiles = &(o.EnforcerProfiles)
		case "externalNetworks":
			sp.ExternalNetworks = &(o.ExternalNetworks)
		case "filePaths":
			sp.FilePaths = &(o.FilePaths)
		case "hostServices":
			sp.HostServices = &(o.HostServices)
		case "isolationProfiles":
			sp.IsolationProfiles = &(o.IsolationProfiles)
		case "name":
			sp.Name = &(o.Name)
		case "namespaces":
			sp.Namespaces = &(o.Namespaces)
		case "policyNamespace":
			sp.PolicyNamespace = &(o.PolicyNamespace)
		case "policyUpdateTime":
			sp.PolicyUpdateTime = &(o.PolicyUpdateTime)
		case "propagated":
			sp.Propagated = &(o.Propagated)
		case "relation":
			sp.Relation = &(o.Relation)
		case "services":
			sp.Services = &(o.Services)
		case "tagClauses":
			sp.TagClauses = &(o.TagClauses)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePolicyRule to the object.
func (o *PolicyRule) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePolicyRule)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.AuditProfiles != nil {
		o.AuditProfiles = *so.AuditProfiles
	}
	if so.EnforcerProfiles != nil {
		o.EnforcerProfiles = *so.EnforcerProfiles
	}
	if so.ExternalNetworks != nil {
		o.ExternalNetworks = *so.ExternalNetworks
	}
	if so.FilePaths != nil {
		o.FilePaths = *so.FilePaths
	}
	if so.HostServices != nil {
		o.HostServices = *so.HostServices
	}
	if so.IsolationProfiles != nil {
		o.IsolationProfiles = *so.IsolationProfiles
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespaces != nil {
		o.Namespaces = *so.Namespaces
	}
	if so.PolicyNamespace != nil {
		o.PolicyNamespace = *so.PolicyNamespace
	}
	if so.PolicyUpdateTime != nil {
		o.PolicyUpdateTime = *so.PolicyUpdateTime
	}
	if so.Propagated != nil {
		o.Propagated = *so.Propagated
	}
	if so.Relation != nil {
		o.Relation = *so.Relation
	}
	if so.Services != nil {
		o.Services = *so.Services
	}
	if so.TagClauses != nil {
		o.TagClauses = *so.TagClauses
	}
}

// DeepCopy returns a deep copy if the PolicyRule.
func (o *PolicyRule) DeepCopy() *PolicyRule {

	if o == nil {
		return nil
	}

	out := &PolicyRule{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PolicyRule.
func (o *PolicyRule) DeepCopyInto(out *PolicyRule) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PolicyRule: %s", err))
	}

	*out = *target.(*PolicyRule)
}

// Validate valides the current information stored into the structure.
func (o *PolicyRule) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.AuditProfiles {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	for _, sub := range o.EnforcerProfiles {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	for _, sub := range o.ExternalNetworks {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	for _, sub := range o.FilePaths {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	for _, sub := range o.HostServices {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	for _, sub := range o.IsolationProfiles {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = append(errors, err)
	}

	for _, sub := range o.Namespaces {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	for _, sub := range o.Services {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
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
func (*PolicyRule) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PolicyRuleAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PolicyRuleLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PolicyRule) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PolicyRuleAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PolicyRule) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "action":
		return o.Action
	case "auditProfiles":
		return o.AuditProfiles
	case "enforcerProfiles":
		return o.EnforcerProfiles
	case "externalNetworks":
		return o.ExternalNetworks
	case "filePaths":
		return o.FilePaths
	case "hostServices":
		return o.HostServices
	case "isolationProfiles":
		return o.IsolationProfiles
	case "name":
		return o.Name
	case "namespaces":
		return o.Namespaces
	case "policyNamespace":
		return o.PolicyNamespace
	case "policyUpdateTime":
		return o.PolicyUpdateTime
	case "propagated":
		return o.Propagated
	case "relation":
		return o.Relation
	case "services":
		return o.Services
	case "tagClauses":
		return o.TagClauses
	}

	return nil
}

// PolicyRuleAttributesMap represents the map of attribute for PolicyRule.
var PolicyRuleAttributesMap = map[string]elemental.AttributeSpecification{
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
		Type:           "string",
	},
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Action",
		Description:    `Action defines set of actions that must be enforced when a dependency is met.`,
		Exposed:        true,
		Name:           "action",
		SubType:        "map[string]map[string]interface{}",
		Type:           "external",
	},
	"AuditProfiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditProfiles",
		Description:    `AuditProfiles provides the audit profiles that must be applied.`,
		Exposed:        true,
		Name:           "auditProfiles",
		SubType:        "auditprofile",
		Type:           "refList",
	},
	"EnforcerProfiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerProfiles",
		Description:    `EnforcerProfiles provides the information about the server profile.`,
		Exposed:        true,
		Name:           "enforcerProfiles",
		SubType:        "enforcerprofile",
		Type:           "refList",
	},
	"ExternalNetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExternalNetworks",
		Description:    `Policy target networks.`,
		Exposed:        true,
		Name:           "externalNetworks",
		SubType:        "externalnetwork",
		Type:           "refList",
	},
	"FilePaths": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FilePaths",
		Description:    `Policy target file paths.`,
		Exposed:        true,
		Name:           "filePaths",
		SubType:        "filepath",
		Type:           "refList",
	},
	"HostServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HostServices",
		Description:    `HostServices provides the list of host services that must be instantiated.`,
		Exposed:        true,
		Name:           "hostServices",
		SubType:        "hostservice",
		Type:           "refList",
	},
	"IsolationProfiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IsolationProfiles",
		Description:    `IsolationProfiles are the isolation profiles of the rule.`,
		Exposed:        true,
		Name:           "isolationProfiles",
		SubType:        "isolationprofile",
		Type:           "refList",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
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
	"Namespaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespaces",
		Description:    `Policy target namespaces.`,
		Exposed:        true,
		Name:           "namespaces",
		SubType:        "namespace",
		Type:           "refList",
	},
	"PolicyNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `PolicyNamespace is the namespace of the policy that created this rule.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Type:           "string",
	},
	"PolicyUpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyUpdateTime",
		Description:    `Last time the policy was updated.`,
		Exposed:        true,
		Name:           "policyUpdateTime",
		Type:           "time",
	},
	"Propagated": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagated",
		Description:    `Propagated indicates if the policy is propagated.`,
		Exposed:        true,
		Name:           "propagated",
		Type:           "boolean",
	},
	"Relation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Relation",
		Description: `Relation describes the required operation to be performed between subjects and
objects.`,
		Exposed: true,
		Name:    "relation",
		SubType: "string",
		Type:    "list",
	},
	"Services": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Services",
		Description:    `Services provides the services of this policy rule.`,
		Exposed:        true,
		Name:           "services",
		SubType:        "service",
		Type:           "refList",
	},
	"TagClauses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TagClauses",
		Description:    `Policy target tags.`,
		Exposed:        true,
		Name:           "tagClauses",
		SubType:        "[][]string",
		Type:           "external",
	},
}

// PolicyRuleLowerCaseAttributesMap represents the map of attribute for PolicyRule.
var PolicyRuleLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
		Type:           "string",
	},
	"action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Action",
		Description:    `Action defines set of actions that must be enforced when a dependency is met.`,
		Exposed:        true,
		Name:           "action",
		SubType:        "map[string]map[string]interface{}",
		Type:           "external",
	},
	"auditprofiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AuditProfiles",
		Description:    `AuditProfiles provides the audit profiles that must be applied.`,
		Exposed:        true,
		Name:           "auditProfiles",
		SubType:        "auditprofile",
		Type:           "refList",
	},
	"enforcerprofiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerProfiles",
		Description:    `EnforcerProfiles provides the information about the server profile.`,
		Exposed:        true,
		Name:           "enforcerProfiles",
		SubType:        "enforcerprofile",
		Type:           "refList",
	},
	"externalnetworks": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExternalNetworks",
		Description:    `Policy target networks.`,
		Exposed:        true,
		Name:           "externalNetworks",
		SubType:        "externalnetwork",
		Type:           "refList",
	},
	"filepaths": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "FilePaths",
		Description:    `Policy target file paths.`,
		Exposed:        true,
		Name:           "filePaths",
		SubType:        "filepath",
		Type:           "refList",
	},
	"hostservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "HostServices",
		Description:    `HostServices provides the list of host services that must be instantiated.`,
		Exposed:        true,
		Name:           "hostServices",
		SubType:        "hostservice",
		Type:           "refList",
	},
	"isolationprofiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "IsolationProfiles",
		Description:    `IsolationProfiles are the isolation profiles of the rule.`,
		Exposed:        true,
		Name:           "isolationProfiles",
		SubType:        "isolationprofile",
		Type:           "refList",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
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
	"namespaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespaces",
		Description:    `Policy target namespaces.`,
		Exposed:        true,
		Name:           "namespaces",
		SubType:        "namespace",
		Type:           "refList",
	},
	"policynamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `PolicyNamespace is the namespace of the policy that created this rule.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Type:           "string",
	},
	"policyupdatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyUpdateTime",
		Description:    `Last time the policy was updated.`,
		Exposed:        true,
		Name:           "policyUpdateTime",
		Type:           "time",
	},
	"propagated": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagated",
		Description:    `Propagated indicates if the policy is propagated.`,
		Exposed:        true,
		Name:           "propagated",
		Type:           "boolean",
	},
	"relation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Relation",
		Description: `Relation describes the required operation to be performed between subjects and
objects.`,
		Exposed: true,
		Name:    "relation",
		SubType: "string",
		Type:    "list",
	},
	"services": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Services",
		Description:    `Services provides the services of this policy rule.`,
		Exposed:        true,
		Name:           "services",
		SubType:        "service",
		Type:           "refList",
	},
	"tagclauses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TagClauses",
		Description:    `Policy target tags.`,
		Exposed:        true,
		Name:           "tagClauses",
		SubType:        "[][]string",
		Type:           "external",
	},
}

// SparsePolicyRulesList represents a list of SparsePolicyRules
type SparsePolicyRulesList []*SparsePolicyRule

// Identity returns the identity of the objects in the list.
func (o SparsePolicyRulesList) Identity() elemental.Identity {

	return PolicyRuleIdentity
}

// Copy returns a pointer to a copy the SparsePolicyRulesList.
func (o SparsePolicyRulesList) Copy() elemental.Identifiables {

	copy := append(SparsePolicyRulesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePolicyRulesList.
func (o SparsePolicyRulesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePolicyRulesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePolicyRule))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePolicyRulesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePolicyRulesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparsePolicyRulesList converted to PolicyRulesList.
func (o SparsePolicyRulesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePolicyRulesList) Version() int {

	return 1
}

// SparsePolicyRule represents the sparse version of a policyrule.
type SparsePolicyRule struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Action defines set of actions that must be enforced when a dependency is met.
	Action *map[string]map[string]interface{} `json:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// AuditProfiles provides the audit profiles that must be applied.
	AuditProfiles *AuditProfilesList `json:"auditProfiles,omitempty" bson:"-" mapstructure:"auditProfiles,omitempty"`

	// EnforcerProfiles provides the information about the server profile.
	EnforcerProfiles *EnforcerProfilesList `json:"enforcerProfiles,omitempty" bson:"-" mapstructure:"enforcerProfiles,omitempty"`

	// Policy target networks.
	ExternalNetworks *ExternalNetworksList `json:"externalNetworks,omitempty" bson:"-" mapstructure:"externalNetworks,omitempty"`

	// Policy target file paths.
	FilePaths *FilePathsList `json:"filePaths,omitempty" bson:"-" mapstructure:"filePaths,omitempty"`

	// HostServices provides the list of host services that must be instantiated.
	HostServices *HostServicesList `json:"hostServices,omitempty" bson:"-" mapstructure:"hostServices,omitempty"`

	// IsolationProfiles are the isolation profiles of the rule.
	IsolationProfiles *IsolationProfilesList `json:"isolationProfiles,omitempty" bson:"-" mapstructure:"isolationProfiles,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Policy target namespaces.
	Namespaces *NamespacesList `json:"namespaces,omitempty" bson:"-" mapstructure:"namespaces,omitempty"`

	// PolicyNamespace is the namespace of the policy that created this rule.
	PolicyNamespace *string `json:"policyNamespace,omitempty" bson:"-" mapstructure:"policyNamespace,omitempty"`

	// Last time the policy was updated.
	PolicyUpdateTime *time.Time `json:"policyUpdateTime,omitempty" bson:"-" mapstructure:"policyUpdateTime,omitempty"`

	// Propagated indicates if the policy is propagated.
	Propagated *bool `json:"propagated,omitempty" bson:"-" mapstructure:"propagated,omitempty"`

	// Relation describes the required operation to be performed between subjects and
	// objects.
	Relation *[]string `json:"relation,omitempty" bson:"-" mapstructure:"relation,omitempty"`

	// Services provides the services of this policy rule.
	Services *ServicesList `json:"services,omitempty" bson:"-" mapstructure:"services,omitempty"`

	// Policy target tags.
	TagClauses *[][]string `json:"tagClauses,omitempty" bson:"-" mapstructure:"tagClauses,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparsePolicyRule returns a new  SparsePolicyRule.
func NewSparsePolicyRule() *SparsePolicyRule {
	return &SparsePolicyRule{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePolicyRule) Identity() elemental.Identity {

	return PolicyRuleIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePolicyRule) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePolicyRule) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparsePolicyRule) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePolicyRule) ToPlain() elemental.PlainIdentifiable {

	out := NewPolicyRule()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.AuditProfiles != nil {
		out.AuditProfiles = *o.AuditProfiles
	}
	if o.EnforcerProfiles != nil {
		out.EnforcerProfiles = *o.EnforcerProfiles
	}
	if o.ExternalNetworks != nil {
		out.ExternalNetworks = *o.ExternalNetworks
	}
	if o.FilePaths != nil {
		out.FilePaths = *o.FilePaths
	}
	if o.HostServices != nil {
		out.HostServices = *o.HostServices
	}
	if o.IsolationProfiles != nil {
		out.IsolationProfiles = *o.IsolationProfiles
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespaces != nil {
		out.Namespaces = *o.Namespaces
	}
	if o.PolicyNamespace != nil {
		out.PolicyNamespace = *o.PolicyNamespace
	}
	if o.PolicyUpdateTime != nil {
		out.PolicyUpdateTime = *o.PolicyUpdateTime
	}
	if o.Propagated != nil {
		out.Propagated = *o.Propagated
	}
	if o.Relation != nil {
		out.Relation = *o.Relation
	}
	if o.Services != nil {
		out.Services = *o.Services
	}
	if o.TagClauses != nil {
		out.TagClauses = *o.TagClauses
	}

	return out
}

// GetName returns the Name of the receiver.
func (o *SparsePolicyRule) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparsePolicyRule) SetName(name string) {

	o.Name = &name
}

// DeepCopy returns a deep copy if the SparsePolicyRule.
func (o *SparsePolicyRule) DeepCopy() *SparsePolicyRule {

	if o == nil {
		return nil
	}

	out := &SparsePolicyRule{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePolicyRule.
func (o *SparsePolicyRule) DeepCopyInto(out *SparsePolicyRule) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePolicyRule: %s", err))
	}

	*out = *target.(*SparsePolicyRule)
}
