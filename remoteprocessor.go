package gaia

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RemoteProcessorModeValue represents the possible values for attribute "mode".
type RemoteProcessorModeValue string

const (
	// RemoteProcessorModePost represents the value Post.
	RemoteProcessorModePost RemoteProcessorModeValue = "Post"

	// RemoteProcessorModePre represents the value Pre.
	RemoteProcessorModePre RemoteProcessorModeValue = "Pre"
)

// RemoteProcessorIdentity represents the Identity of the object.
var RemoteProcessorIdentity = elemental.Identity{
	Name:     "remoteprocessor",
	Category: "remoteprocessors",
	Package:  "rufus",
	Private:  false,
}

// RemoteProcessorsList represents a list of RemoteProcessors
type RemoteProcessorsList []*RemoteProcessor

// Identity returns the identity of the objects in the list.
func (o RemoteProcessorsList) Identity() elemental.Identity {

	return RemoteProcessorIdentity
}

// Copy returns a pointer to a copy the RemoteProcessorsList.
func (o RemoteProcessorsList) Copy() elemental.Identifiables {

	copy := append(RemoteProcessorsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RemoteProcessorsList.
func (o RemoteProcessorsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RemoteProcessorsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*RemoteProcessor))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RemoteProcessorsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RemoteProcessorsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the RemoteProcessorsList converted to SparseRemoteProcessorsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o RemoteProcessorsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o RemoteProcessorsList) Version() int {

	return 1
}

// RemoteProcessor represents the model of a remoteprocessor
type RemoteProcessor struct {
	// Represents the claims of the currently managed object.
	Claims []string `json:"claims" bson:"-" mapstructure:"claims,omitempty"`

	// Represents data received from the service.
	Input json.RawMessage `json:"input" bson:"-" mapstructure:"input,omitempty"`

	// Node defines the type of the hook.
	Mode RemoteProcessorModeValue `json:"mode" bson:"-" mapstructure:"mode,omitempty"`

	// Represents the current namespace.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Define the operation that is currently handled by the service.
	Operation elemental.Operation `json:"operation" bson:"-" mapstructure:"operation,omitempty"`

	// Returns the OutputData filled with the processor information.
	Output elemental.Identifiable `json:"output" bson:"-" mapstructure:"output,omitempty"`

	// RequestID gives the id of the request coming from the main server.
	RequestID string `json:"requestID" bson:"requestid" mapstructure:"requestID,omitempty"`

	// Represents the Identity name of the managed object.
	TargetIdentity string `json:"targetIdentity" bson:"-" mapstructure:"targetIdentity,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewRemoteProcessor returns a new *RemoteProcessor
func NewRemoteProcessor() *RemoteProcessor {

	return &RemoteProcessor{
		ModelVersion: 1,
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
func (o *RemoteProcessor) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *RemoteProcessor) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *RemoteProcessor) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *RemoteProcessor) Doc() string {
	return `Hook to integrate an Aporeto service.`
}

func (o *RemoteProcessor) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *RemoteProcessor) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseRemoteProcessor{
			Claims:         &o.Claims,
			Input:          &o.Input,
			Mode:           &o.Mode,
			Namespace:      &o.Namespace,
			Operation:      &o.Operation,
			Output:         &o.Output,
			RequestID:      &o.RequestID,
			TargetIdentity: &o.TargetIdentity,
		}
	}

	sp := &SparseRemoteProcessor{}
	for _, f := range fields {
		switch f {
		case "claims":
			sp.Claims = &(o.Claims)
		case "input":
			sp.Input = &(o.Input)
		case "mode":
			sp.Mode = &(o.Mode)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "operation":
			sp.Operation = &(o.Operation)
		case "output":
			sp.Output = &(o.Output)
		case "requestID":
			sp.RequestID = &(o.RequestID)
		case "targetIdentity":
			sp.TargetIdentity = &(o.TargetIdentity)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseRemoteProcessor to the object.
func (o *RemoteProcessor) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseRemoteProcessor)
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.Input != nil {
		o.Input = *so.Input
	}
	if so.Mode != nil {
		o.Mode = *so.Mode
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Operation != nil {
		o.Operation = *so.Operation
	}
	if so.Output != nil {
		o.Output = *so.Output
	}
	if so.RequestID != nil {
		o.RequestID = *so.RequestID
	}
	if so.TargetIdentity != nil {
		o.TargetIdentity = *so.TargetIdentity
	}
}

// DeepCopy returns a deep copy if the RemoteProcessor.
func (o *RemoteProcessor) DeepCopy() *RemoteProcessor {

	if o == nil {
		return nil
	}

	out := &RemoteProcessor{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *RemoteProcessor.
func (o *RemoteProcessor) DeepCopyInto(out *RemoteProcessor) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy RemoteProcessor: %s", err))
	}

	*out = *target.(*RemoteProcessor)
}

// Validate valides the current information stored into the structure.
func (o *RemoteProcessor) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("input", o.Input); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("mode", string(o.Mode), []string{"Post", "Pre"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("operation", o.Operation); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("targetIdentity", o.TargetIdentity); err != nil {
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
func (*RemoteProcessor) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RemoteProcessorAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RemoteProcessorLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*RemoteProcessor) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RemoteProcessorAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *RemoteProcessor) ValueForAttribute(name string) interface{} {

	switch name {
	case "claims":
		return o.Claims
	case "input":
		return o.Input
	case "mode":
		return o.Mode
	case "namespace":
		return o.Namespace
	case "operation":
		return o.Operation
	case "output":
		return o.Output
	case "requestID":
		return o.RequestID
	case "targetIdentity":
		return o.TargetIdentity
	}

	return nil
}

// RemoteProcessorAttributesMap represents the map of attribute for RemoteProcessor.
var RemoteProcessorAttributesMap = map[string]elemental.AttributeSpecification{
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Represents the claims of the currently managed object.`,
		Exposed:        true,
		Name:           "claims",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Input": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Input",
		Description:    `Represents data received from the service.`,
		Exposed:        true,
		Name:           "input",
		Required:       true,
		SubType:        "raw_json",
		Type:           "external",
	},
	"Mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Post", "Pre"},
		ConvertedName:  "Mode",
		Description:    `Node defines the type of the hook.`,
		Exposed:        true,
		Name:           "mode",
		Type:           "enum",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Represents the current namespace.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"Operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Operation",
		Description:    `Define the operation that is currently handled by the service.`,
		Exposed:        true,
		Name:           "operation",
		Required:       true,
		SubType:        "elemental_operation",
		Type:           "external",
	},
	"Output": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Output",
		Description:    `Returns the OutputData filled with the processor information.`,
		Exposed:        true,
		Name:           "output",
		ReadOnly:       true,
		SubType:        "elemental_identitifable",
		Type:           "external",
	},
	"RequestID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RequestID",
		Description:    `RequestID gives the id of the request coming from the main server.`,
		Exposed:        true,
		Name:           "requestID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `Represents the Identity name of the managed object.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Type:           "string",
	},
}

// RemoteProcessorLowerCaseAttributesMap represents the map of attribute for RemoteProcessor.
var RemoteProcessorLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Represents the claims of the currently managed object.`,
		Exposed:        true,
		Name:           "claims",
		Required:       true,
		SubType:        "string",
		Type:           "list",
	},
	"input": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Input",
		Description:    `Represents data received from the service.`,
		Exposed:        true,
		Name:           "input",
		Required:       true,
		SubType:        "raw_json",
		Type:           "external",
	},
	"mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Post", "Pre"},
		ConvertedName:  "Mode",
		Description:    `Node defines the type of the hook.`,
		Exposed:        true,
		Name:           "mode",
		Type:           "enum",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Represents the current namespace.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"operation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Operation",
		Description:    `Define the operation that is currently handled by the service.`,
		Exposed:        true,
		Name:           "operation",
		Required:       true,
		SubType:        "elemental_operation",
		Type:           "external",
	},
	"output": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Output",
		Description:    `Returns the OutputData filled with the processor information.`,
		Exposed:        true,
		Name:           "output",
		ReadOnly:       true,
		SubType:        "elemental_identitifable",
		Type:           "external",
	},
	"requestid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "RequestID",
		Description:    `RequestID gives the id of the request coming from the main server.`,
		Exposed:        true,
		Name:           "requestID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"targetidentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TargetIdentity",
		Description:    `Represents the Identity name of the managed object.`,
		Exposed:        true,
		Name:           "targetIdentity",
		Required:       true,
		Type:           "string",
	},
}

// SparseRemoteProcessorsList represents a list of SparseRemoteProcessors
type SparseRemoteProcessorsList []*SparseRemoteProcessor

// Identity returns the identity of the objects in the list.
func (o SparseRemoteProcessorsList) Identity() elemental.Identity {

	return RemoteProcessorIdentity
}

// Copy returns a pointer to a copy the SparseRemoteProcessorsList.
func (o SparseRemoteProcessorsList) Copy() elemental.Identifiables {

	copy := append(SparseRemoteProcessorsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseRemoteProcessorsList.
func (o SparseRemoteProcessorsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseRemoteProcessorsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseRemoteProcessor))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseRemoteProcessorsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseRemoteProcessorsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseRemoteProcessorsList converted to RemoteProcessorsList.
func (o SparseRemoteProcessorsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseRemoteProcessorsList) Version() int {

	return 1
}

// SparseRemoteProcessor represents the sparse version of a remoteprocessor.
type SparseRemoteProcessor struct {
	// Represents the claims of the currently managed object.
	Claims *[]string `json:"claims,omitempty" bson:"-" mapstructure:"claims,omitempty"`

	// Represents data received from the service.
	Input *json.RawMessage `json:"input,omitempty" bson:"-" mapstructure:"input,omitempty"`

	// Node defines the type of the hook.
	Mode *RemoteProcessorModeValue `json:"mode,omitempty" bson:"-" mapstructure:"mode,omitempty"`

	// Represents the current namespace.
	Namespace *string `json:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// Define the operation that is currently handled by the service.
	Operation *elemental.Operation `json:"operation,omitempty" bson:"-" mapstructure:"operation,omitempty"`

	// Returns the OutputData filled with the processor information.
	Output *elemental.Identifiable `json:"output,omitempty" bson:"-" mapstructure:"output,omitempty"`

	// RequestID gives the id of the request coming from the main server.
	RequestID *string `json:"requestID,omitempty" bson:"requestid" mapstructure:"requestID,omitempty"`

	// Represents the Identity name of the managed object.
	TargetIdentity *string `json:"targetIdentity,omitempty" bson:"-" mapstructure:"targetIdentity,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseRemoteProcessor returns a new  SparseRemoteProcessor.
func NewSparseRemoteProcessor() *SparseRemoteProcessor {
	return &SparseRemoteProcessor{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseRemoteProcessor) Identity() elemental.Identity {

	return RemoteProcessorIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseRemoteProcessor) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseRemoteProcessor) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseRemoteProcessor) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseRemoteProcessor) ToPlain() elemental.PlainIdentifiable {

	out := NewRemoteProcessor()
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.Input != nil {
		out.Input = *o.Input
	}
	if o.Mode != nil {
		out.Mode = *o.Mode
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Operation != nil {
		out.Operation = *o.Operation
	}
	if o.Output != nil {
		out.Output = *o.Output
	}
	if o.RequestID != nil {
		out.RequestID = *o.RequestID
	}
	if o.TargetIdentity != nil {
		out.TargetIdentity = *o.TargetIdentity
	}

	return out
}

// DeepCopy returns a deep copy if the SparseRemoteProcessor.
func (o *SparseRemoteProcessor) DeepCopy() *SparseRemoteProcessor {

	if o == nil {
		return nil
	}

	out := &SparseRemoteProcessor{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseRemoteProcessor.
func (o *SparseRemoteProcessor) DeepCopyInto(out *SparseRemoteProcessor) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseRemoteProcessor: %s", err))
	}

	*out = *target.(*SparseRemoteProcessor)
}
