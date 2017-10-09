package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// PolicyRuleIdentity represents the Identity of the object
var PolicyRuleIdentity = elemental.Identity{
	Name:     "policyrule",
	Category: "policyrules",
}

// PolicyRulesList represents a list of PolicyRules
type PolicyRulesList []*PolicyRule

// ContentIdentity returns the identity of the objects in the list.
func (o PolicyRulesList) ContentIdentity() elemental.Identity {

	return PolicyRuleIdentity
}

// Copy returns a pointer to a copy the PolicyRulesList.
func (o PolicyRulesList) Copy() elemental.ContentIdentifiable {

	copy := append(PolicyRulesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PolicyRulesList.
func (o PolicyRulesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(PolicyRulesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PolicyRule))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PolicyRulesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PolicyRulesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o PolicyRulesList) Version() int {

	return 1
}

// PolicyRule represents the model of a policyrule
type PolicyRule struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-"`

	// Action defines set of actions that must be enforced when a dependency is met.
	Action map[string]map[string]string `json:"action" bson:"-"`

	// EnforcerProfiles provides the information about the server profile.
	EnforcerProfiles EnforcerProfilesList `json:"enforcerProfiles" bson:"-"`

	// Policy target networks
	ExternalServices ExternalServicesList `json:"externalServices" bson:"-"`

	// Policy target networks
	FilePaths FilePathsList `json:"filePaths" bson:"-"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Policy target networks
	Namespaces NamespacesList `json:"namespaces" bson:"-"`

	// Propagated indicates if the policy is propagated.
	Propagated bool `json:"propagated" bson:"-"`

	// Relation describes the required operation to be performed between subjects and objects
	Relation []string `json:"relation" bson:"-"`

	// Policy target networks
	SystemCalls SystemCallsList `json:"systemCalls" bson:"-"`

	// Policy target tags
	TagClauses [][]string `json:"tagClauses" bson:"-"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewPolicyRule returns a new *PolicyRule
func NewPolicyRule() *PolicyRule {

	return &PolicyRule{
		ModelVersion:     1,
		EnforcerProfiles: EnforcerProfilesList{},
		ExternalServices: ExternalServicesList{},
		FilePaths:        FilePathsList{},
		Namespaces:       NamespacesList{},
		SystemCalls:      SystemCallsList{},
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
func (o *PolicyRule) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
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
	return nodocString
}

func (o *PolicyRule) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the name of the receiver
func (o *PolicyRule) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *PolicyRule) SetName(name string) {
	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *PolicyRule) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
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

// PolicyRuleAttributesMap represents the map of attribute for PolicyRule.
var PolicyRuleAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Action defines set of actions that must be enforced when a dependency is met.`,
		Exposed:        true,
		Name:           "action",
		SubType:        "actions_list",
		Type:           "external",
	},
	"EnforcerProfiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `EnforcerProfiles provides the information about the server profile.`,
		Exposed:        true,
		Name:           "enforcerProfiles",
		SubType:        "enforcerprofiles_list",
		Type:           "external",
	},
	"ExternalServices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "externalServices",
		SubType:        "network_entities",
		Type:           "external",
	},
	"FilePaths": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "filePaths",
		SubType:        "file_entities",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultOrder:   true,
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Namespaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "namespaces",
		SubType:        "namespace_entities",
		Type:           "external",
	},
	"Propagated": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Propagated indicates if the policy is propagated.`,
		Exposed:        true,
		Name:           "propagated",
		Type:           "boolean",
	},
	"Relation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Relation describes the required operation to be performed between subjects and objects`,
		Exposed:        true,
		Name:           "relation",
		SubType:        "relations_list",
		Type:           "external",
	},
	"SystemCalls": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "systemCalls",
		SubType:        "syscall_entities",
		Type:           "external",
	},
	"TagClauses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Policy target tags`,
		Exposed:        true,
		Name:           "tagClauses",
		SubType:        "target_tags",
		Type:           "external",
	},
}

// PolicyRuleLowerCaseAttributesMap represents the map of attribute for PolicyRule.
var PolicyRuleLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Action defines set of actions that must be enforced when a dependency is met.`,
		Exposed:        true,
		Name:           "action",
		SubType:        "actions_list",
		Type:           "external",
	},
	"enforcerprofiles": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `EnforcerProfiles provides the information about the server profile.`,
		Exposed:        true,
		Name:           "enforcerProfiles",
		SubType:        "enforcerprofiles_list",
		Type:           "external",
	},
	"externalservices": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "externalServices",
		SubType:        "network_entities",
		Type:           "external",
	},
	"filepaths": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "filePaths",
		SubType:        "file_entities",
		Type:           "external",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		DefaultOrder:   true,
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"namespaces": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "namespaces",
		SubType:        "namespace_entities",
		Type:           "external",
	},
	"propagated": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Propagated indicates if the policy is propagated.`,
		Exposed:        true,
		Name:           "propagated",
		Type:           "boolean",
	},
	"relation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Relation describes the required operation to be performed between subjects and objects`,
		Exposed:        true,
		Name:           "relation",
		SubType:        "relations_list",
		Type:           "external",
	},
	"systemcalls": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Policy target networks `,
		Exposed:        true,
		Name:           "systemCalls",
		SubType:        "syscall_entities",
		Type:           "external",
	},
	"tagclauses": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Policy target tags`,
		Exposed:        true,
		Name:           "tagClauses",
		SubType:        "target_tags",
		Type:           "external",
	},
}
