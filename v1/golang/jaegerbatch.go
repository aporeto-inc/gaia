package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/uber/jaeger-client-go/thrift-gen/jaeger"
)

// JaegerbatchIdentity represents the Identity of the object.
var JaegerbatchIdentity = elemental.Identity{
	Name:     "jaegerbatch",
	Category: "jaegerbatchs",
	Private:  false,
}

// JaegerbatchsList represents a list of Jaegerbatchs
type JaegerbatchsList []*Jaegerbatch

// ContentIdentity returns the identity of the objects in the list.
func (o JaegerbatchsList) ContentIdentity() elemental.Identity {

	return JaegerbatchIdentity
}

// Copy returns a pointer to a copy the JaegerbatchsList.
func (o JaegerbatchsList) Copy() elemental.ContentIdentifiable {

	copy := append(JaegerbatchsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the JaegerbatchsList.
func (o JaegerbatchsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(JaegerbatchsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Jaegerbatch))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o JaegerbatchsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o JaegerbatchsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o JaegerbatchsList) Version() int {

	return 1
}

// Jaegerbatch represents the model of a jaegerbatch
type Jaegerbatch struct {
	// Represent an jaeger batch
	Batch *jaeger.Batch `json:"batch" bson:"batch" mapstructure:"batch,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewJaegerbatch returns a new *Jaegerbatch
func NewJaegerbatch() *Jaegerbatch {

	return &Jaegerbatch{
		ModelVersion: 1,
		Batch:        &jaeger.Batch{},
	}
}

// Identity returns the Identity of the object.
func (o *Jaegerbatch) Identity() elemental.Identity {

	return JaegerbatchIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Jaegerbatch) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Jaegerbatch) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Jaegerbatch) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Jaegerbatch) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Jaegerbatch) Doc() string {
	return `A jaegerbatch is a batch of jaeger spans. This is used by external service to post jaeger span in our private jaeger services`
}

func (o *Jaegerbatch) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Jaegerbatch) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Jaegerbatch) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := JaegerbatchAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return JaegerbatchLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Jaegerbatch) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return JaegerbatchAttributesMap
}

// JaegerbatchAttributesMap represents the map of attribute for Jaegerbatch.
var JaegerbatchAttributesMap = map[string]elemental.AttributeSpecification{
	"Batch": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Batch",
		CreationOnly:   true,
		Description:    `Represent an jaeger batch`,
		Exposed:        true,
		Name:           "batch",
		Stored:         true,
		SubType:        "jaeger_batch",
		Type:           "external",
	},
}

// JaegerbatchLowerCaseAttributesMap represents the map of attribute for Jaegerbatch.
var JaegerbatchLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"batch": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Batch",
		CreationOnly:   true,
		Description:    `Represent an jaeger batch`,
		Exposed:        true,
		Name:           "batch",
		Stored:         true,
		SubType:        "jaeger_batch",
		Type:           "external",
	},
}
