package types

// An AutomationTemplateParameterType represent the type of an AutomationTemplateParameter.
type AutomationTemplateParameterType string

// Various possible values for AutomationTemplateParameterType.
const (
	AutomationTemplateParameterTypeString AutomationTemplateParameterType = "string"
	AutomationTemplateParameterTypeBool   AutomationTemplateParameterType = "boolean"
	AutomationTemplateParameterTypeInt    AutomationTemplateParameterType = "int"
	AutomationTemplateParameterTypeFloat  AutomationTemplateParameterType = "float"
)

// An AutomationTemplateParameter represents an automation template parameter.
type AutomationTemplateParameter struct {
	Name         string                          `json:"name"`
	Description  string                          `json:"description"`
	Type         AutomationTemplateParameterType `json:"type"`
	DefaultValue interface{}                     `json:"defaultValue"`
	Required     bool                            `json:"required"`
}
