package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// PolicyRendererTypeValue represents the possible values for attribute "type".
type PolicyRendererTypeValue string

const (
	// PolicyRendererTypeAPIAuthorization represents the value APIAuthorization.
	PolicyRendererTypeAPIAuthorization PolicyRendererTypeValue = "APIAuthorization"

	// PolicyRendererTypeEnforcerProfile represents the value EnforcerProfile.
	PolicyRendererTypeEnforcerProfile PolicyRendererTypeValue = "EnforcerProfile"

	// PolicyRendererTypeFile represents the value File.
	PolicyRendererTypeFile PolicyRendererTypeValue = "File"

	// PolicyRendererTypeHook represents the value Hook.
	PolicyRendererTypeHook PolicyRendererTypeValue = "Hook"

	// PolicyRendererTypeNamespaceMapping represents the value NamespaceMapping.
	PolicyRendererTypeNamespaceMapping PolicyRendererTypeValue = "NamespaceMapping"

	// PolicyRendererTypeNetwork represents the value Network.
	PolicyRendererTypeNetwork PolicyRendererTypeValue = "Network"

	// PolicyRendererTypeProcessingUnit represents the value ProcessingUnit.
	PolicyRendererTypeProcessingUnit PolicyRendererTypeValue = "ProcessingUnit"

	// PolicyRendererTypeQuota represents the value Quota.
	PolicyRendererTypeQuota PolicyRendererTypeValue = "Quota"

	// PolicyRendererTypeSyscall represents the value Syscall.
	PolicyRendererTypeSyscall PolicyRendererTypeValue = "Syscall"

	// PolicyRendererTypeTokenScope represents the value TokenScope.
	PolicyRendererTypeTokenScope PolicyRendererTypeValue = "TokenScope"
)

// PolicyRendererIdentity represents the Identity of the object.
var PolicyRendererIdentity = elemental.Identity{
	Name:     "policyrenderer",
	Category: "policyrenderers",
	Private:  true,
}

// PolicyRenderersList represents a list of PolicyRenderers
type PolicyRenderersList []*PolicyRenderer

// ContentIdentity returns the identity of the objects in the list.
func (o PolicyRenderersList) ContentIdentity() elemental.Identity {

	return PolicyRendererIdentity
}

// Copy returns a pointer to a copy the PolicyRenderersList.
func (o PolicyRenderersList) Copy() elemental.ContentIdentifiable {

	copy := append(PolicyRenderersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PolicyRenderersList.
func (o PolicyRenderersList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(PolicyRenderersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PolicyRenderer))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PolicyRenderersList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PolicyRenderersList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o PolicyRenderersList) Version() int {

	return 1
}

// PolicyRenderer represents the model of a policyrenderer
type PolicyRenderer struct {
	// List of policies rendered for the given set of tags.
	Policies PolicyRulesList `json:"policies" bson:"-" mapstructure:"policies,omitempty"`

	// List of tags of the object to render the hook policy for.
	Tags []string `json:"tags" bson:"-" mapstructure:"tags,omitempty"`

	// Type of the policy to render.
	Type PolicyRendererTypeValue `json:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewPolicyRenderer returns a new *PolicyRenderer
func NewPolicyRenderer() *PolicyRenderer {

	return &PolicyRenderer{
		ModelVersion: 1,
		Policies:     PolicyRulesList{},
	}
}

// Identity returns the Identity of the object.
func (o *PolicyRenderer) Identity() elemental.Identity {

	return PolicyRendererIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PolicyRenderer) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PolicyRenderer) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *PolicyRenderer) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *PolicyRenderer) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PolicyRenderer) Doc() string {
	return `Returns the list of hook policies that matches the given set of tags.`
}

func (o *PolicyRenderer) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *PolicyRenderer) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"APIAuthorization", "EnforcerProfile", "File", "Hook", "NamespaceMapping", "Network", "ProcessingUnit", "Quota", "Syscall", "TokenScope"}, false); err != nil {
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
func (*PolicyRenderer) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PolicyRendererAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PolicyRendererLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PolicyRenderer) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PolicyRendererAttributesMap
}

// PolicyRendererAttributesMap represents the map of attribute for PolicyRenderer.
var PolicyRendererAttributesMap = map[string]elemental.AttributeSpecification{
	"Policies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Policies",
		Description:    `List of policies rendered for the given set of tags.`,
		Exposed:        true,
		Name:           "policies",
		ReadOnly:       true,
		SubType:        "policy_rules_list",
		Type:           "external",
	},
	"Tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `List of tags of the object to render the hook policy for.`,
		Exposed:        true,
		Name:           "tags",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"APIAuthorization", "EnforcerProfile", "File", "Hook", "NamespaceMapping", "Network", "ProcessingUnit", "Quota", "Syscall", "TokenScope"},
		ConvertedName:  "Type",
		Description:    `Type of the policy to render.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Type:           "enum",
	},
}

// PolicyRendererLowerCaseAttributesMap represents the map of attribute for PolicyRenderer.
var PolicyRendererLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"policies": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Policies",
		Description:    `List of policies rendered for the given set of tags.`,
		Exposed:        true,
		Name:           "policies",
		ReadOnly:       true,
		SubType:        "policy_rules_list",
		Type:           "external",
	},
	"tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `List of tags of the object to render the hook policy for.`,
		Exposed:        true,
		Name:           "tags",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{"APIAuthorization", "EnforcerProfile", "File", "Hook", "NamespaceMapping", "Network", "ProcessingUnit", "Quota", "Syscall", "TokenScope"},
		ConvertedName:  "Type",
		Description:    `Type of the policy to render.`,
		Exposed:        true,
		Name:           "type",
		Required:       true,
		Type:           "enum",
	},
}
