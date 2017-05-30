package rufusmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "encoding/json"

// RemoteProcessorModeValue represents the possible values for attribute "mode".
type RemoteProcessorModeValue string

const (
	// RemoteProcessorModePost represents the value Post.
	RemoteProcessorModePost RemoteProcessorModeValue = "Post"

	// RemoteProcessorModePre represents the value Pre.
	RemoteProcessorModePre RemoteProcessorModeValue = "Pre"
)

// RemoteProcessorIdentity represents the Identity of the object
var RemoteProcessorIdentity = elemental.Identity{
	Name:     "remoteprocessor",
	Category: "remoteprocessors",
}

// RemoteProcessorsList represents a list of RemoteProcessors
type RemoteProcessorsList []*RemoteProcessor

// ContentIdentity returns the identity of the objects in the list.
func (o RemoteProcessorsList) ContentIdentity() elemental.Identity {

	return RemoteProcessorIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o RemoteProcessorsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RemoteProcessorsList) DefaultOrder() []string {

	return []string{}
}

// RemoteProcessor represents the model of a remoteprocessor
type RemoteProcessor struct {
	// Represents the claims of the currently managed object
	Claims []string `json:"claims" bson:"-"`

	// Represents data received from the service
	Input json.RawMessage `json:"input" bson:"-"`

	// Node defines the type of the hook
	Mode RemoteProcessorModeValue `json:"mode" bson:"-"`

	// Represents the current namespace
	Namespace string `json:"namespace" bson:"-"`

	// Define the operation that is currently handled by the service
	Operation elemental.Operation `json:"operation" bson:"-"`

	// Returns the OutputData filled with the processor information
	Output elemental.Identifiable `json:"output" bson:"-"`

	// requestID identifies the current context
	RequestID string `json:"requestID" bson:"requestid"`

	// Represents the Identity name of the managed object
	TargetIdentity string `json:"targetIdentity" bson:"-"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewRemoteProcessor returns a new *RemoteProcessor
func NewRemoteProcessor() *RemoteProcessor {

	return &RemoteProcessor{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *RemoteProcessor) Identity() elemental.Identity {

	return RemoteProcessorIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *RemoteProcessor) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *RemoteProcessor) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *RemoteProcessor) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *RemoteProcessor) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *RemoteProcessor) Doc() string {
	return `Hook to integrate in an Aporeto service`
}

func (o *RemoteProcessor) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *RemoteProcessor) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("input", o.Input); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("input", o.Input); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("mode", string(o.Mode), []string{"Post", "Pre"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("operation", o.Operation); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("operation", o.Operation); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("targetIdentity", o.TargetIdentity); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("targetIdentity", o.TargetIdentity); err != nil {
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
func (*RemoteProcessor) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return RemoteProcessorAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RemoteProcessor) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RemoteProcessorAttributesMap
}

// RemoteProcessorAttributesMap represents the map of attribute for RemoteProcessor.
var RemoteProcessorAttributesMap = map[string]elemental.AttributeSpecification{
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Represents the claims of the currently managed object`,
		Exposed:        true,
		Name:           "claims",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Input": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Represents data received from the service`,
		Exposed:        true,
		Name:           "input",
		Required:       true,
		SubType:        "raw_json",
		Type:           "external",
	},
	"Mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Post", "Pre"},
		Description:    `Node defines the type of the hook`,
		Exposed:        true,
		Name:           "mode",
		Required:       true,
		Type:           "enum",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Represents the current namespace`,
		Exposed:        true,
		Format:         "free",
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"Operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Define the operation that is currently handled by the service`,
		Exposed:        true,
		Name:           "operation",
		Required:       true,
		SubType:        "elemental_operation",
		Type:           "external",
	},
	"Output": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Returns the OutputData filled with the processor information`,
		Exposed:        true,
		Name:           "output",
		ReadOnly:       true,
		SubType:        "elemental_identitifable",
		Type:           "external",
	},
	"RequestID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `requestID identifies the current context`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "requestID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Represents the Identity name of the managed object`,
		Exposed:        true,
		Format:         "free",
		Name:           "targetIdentity",
		Required:       true,
		Type:           "string",
	},
}
