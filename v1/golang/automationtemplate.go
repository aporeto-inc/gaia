package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/types"
)

// AutomationTemplateKindValue represents the possible values for attribute "kind".
type AutomationTemplateKindValue string

const (
	// AutomationTemplateKindAction represents the value Action.
	AutomationTemplateKindAction AutomationTemplateKindValue = "Action"

	// AutomationTemplateKindCondition represents the value Condition.
	AutomationTemplateKindCondition AutomationTemplateKindValue = "Condition"
)

// AutomationTemplateIdentity represents the Identity of the object.
var AutomationTemplateIdentity = elemental.Identity{
	Name:     "automationtemplate",
	Category: "automationtemplates",
	Private:  false,
}

// AutomationTemplatesList represents a list of AutomationTemplates
type AutomationTemplatesList []*AutomationTemplate

// ContentIdentity returns the identity of the objects in the list.
func (o AutomationTemplatesList) ContentIdentity() elemental.Identity {

	return AutomationTemplateIdentity
}

// Copy returns a pointer to a copy the AutomationTemplatesList.
func (o AutomationTemplatesList) Copy() elemental.ContentIdentifiable {

	copy := append(AutomationTemplatesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AutomationTemplatesList.
func (o AutomationTemplatesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(AutomationTemplatesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AutomationTemplate))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AutomationTemplatesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AutomationTemplatesList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o AutomationTemplatesList) Version() int {

	return 1
}

// AutomationTemplate represents the model of a automationtemplate
type AutomationTemplate struct {
	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Entitlements contains the entitlements needed for executing the function.
	Entitlements map[string][]elemental.Operation `json:"entitlements" bson:"-" mapstructure:"entitlements,omitempty"`

	// Function contains the code.
	Function string `json:"function" bson:"-" mapstructure:"function,omitempty"`

	// Key contains the unique identifier key for the template.
	Key string `json:"key" bson:"-" mapstructure:"key,omitempty"`

	// Kind represents the kind of template.
	Kind AutomationTemplateKindValue `json:"kind" bson:"-" mapstructure:"kind,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Parameters contains the parameter description of the function.
	Parameters map[string]types.AutomationTemplateParameter `json:"parameters" bson:"-" mapstructure:"parameters,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAutomationTemplate returns a new *AutomationTemplate
func NewAutomationTemplate() *AutomationTemplate {

	return &AutomationTemplate{
		ModelVersion: 1,
		Entitlements: map[string][]elemental.Operation{},
		Kind:         "Condition",
		Parameters:   map[string]types.AutomationTemplateParameter{},
	}
}

// Identity returns the Identity of the object.
func (o *AutomationTemplate) Identity() elemental.Identity {

	return AutomationTemplateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AutomationTemplate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AutomationTemplate) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *AutomationTemplate) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *AutomationTemplate) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *AutomationTemplate) Doc() string {
	return `Templates that ca be used in automations.`
}

func (o *AutomationTemplate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the Name of the receiver.
func (o *AutomationTemplate) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *AutomationTemplate) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *AutomationTemplate) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("kind", string(o.Kind), []string{"Action", "Condition"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*AutomationTemplate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AutomationTemplateAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AutomationTemplateLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AutomationTemplate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AutomationTemplateAttributesMap
}

// AutomationTemplateAttributesMap represents the map of attribute for AutomationTemplate.
var AutomationTemplateAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Entitlements": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Entitlements",
		Description:    `Entitlements contains the entitlements needed for executing the function.`,
		Exposed:        true,
		Name:           "entitlements",
		SubType:        "automation_entitlements",
		Type:           "external",
	},
	"Function": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Function",
		Description:    `Function contains the code.`,
		Exposed:        true,
		Format:         "free",
		Name:           "function",
		Type:           "string",
	},
	"Key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Key",
		Description:    `Key contains the unique identifier key for the template.`,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		Type:           "string",
	},
	"Kind": elemental.AttributeSpecification{
		AllowedChoices: []string{"Action", "Condition"},
		ConvertedName:  "Kind",
		DefaultValue:   AutomationTemplateKindCondition,
		Description:    `Kind represents the kind of template.`,
		Exposed:        true,
		Name:           "kind",
		Type:           "enum",
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
	"Parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters contains the parameter description of the function.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "automation_template_parameters",
		Type:           "external",
	},
}

// AutomationTemplateLowerCaseAttributesMap represents the map of attribute for AutomationTemplate.
var AutomationTemplateLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"entitlements": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Entitlements",
		Description:    `Entitlements contains the entitlements needed for executing the function.`,
		Exposed:        true,
		Name:           "entitlements",
		SubType:        "automation_entitlements",
		Type:           "external",
	},
	"function": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Function",
		Description:    `Function contains the code.`,
		Exposed:        true,
		Format:         "free",
		Name:           "function",
		Type:           "string",
	},
	"key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Key",
		Description:    `Key contains the unique identifier key for the template.`,
		Exposed:        true,
		Format:         "free",
		Name:           "key",
		Type:           "string",
	},
	"kind": elemental.AttributeSpecification{
		AllowedChoices: []string{"Action", "Condition"},
		ConvertedName:  "Kind",
		DefaultValue:   AutomationTemplateKindCondition,
		Description:    `Kind represents the kind of template.`,
		Exposed:        true,
		Name:           "kind",
		Type:           "enum",
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
	"parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters contains the parameter description of the function.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "automation_template_parameters",
		Type:           "external",
	},
}
