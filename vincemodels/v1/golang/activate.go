package vincemodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// ActivateIdentity represents the Identity of the object
var ActivateIdentity = elemental.Identity{
	Name:     "activate",
	Category: "activate",
}

// ActivatesList represents a list of Activates
type ActivatesList []*Activate

// ContentIdentity returns the identity of the objects in the list.
func (o ActivatesList) ContentIdentity() elemental.Identity {

	return ActivateIdentity
}

// List converts the object to an elemental.IdentifiablesList.
func (o ActivatesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ActivatesList) DefaultOrder() []string {

	return []string{}
}

// Activate represents the model of a activate
type Activate struct {
	// Token contains the activation token
	Token string `json:"token" bson:"-"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewActivate returns a new *Activate
func NewActivate() *Activate {

	return &Activate{
		ModelVersion: 1.0,
	}
}

// Identity returns the Identity of the object.
func (o *Activate) Identity() elemental.Identity {

	return ActivateIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Activate) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Activate) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Activate) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *Activate) DefaultOrder() []string {

	return []string{}
}

func (o *Activate) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Activate) Validate() error {

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
func (*Activate) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return ActivateAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Activate) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ActivateAttributesMap
}

// ActivateAttributesMap represents the map of attribute for Activate.
var ActivateAttributesMap = map[string]elemental.AttributeSpecification{
	"Token": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Description:    `Token contains the activation token`,
		Exposed:        true,
		Format:         "free",
		Name:           "token",
		Type:           "string",
	},
}
