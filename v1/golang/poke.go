package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// PokeIdentity represents the Identity of the object.
var PokeIdentity = elemental.Identity{
	Name:     "poke",
	Category: "poke",
	Private:  false,
}

// PokesList represents a list of Pokes
type PokesList []*Poke

// ContentIdentity returns the identity of the objects in the list.
func (o PokesList) ContentIdentity() elemental.Identity {

	return PokeIdentity
}

// Copy returns a pointer to a copy the PokesList.
func (o PokesList) Copy() elemental.ContentIdentifiable {

	copy := append(PokesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PokesList.
func (o PokesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(PokesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Poke))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PokesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PokesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o PokesList) Version() int {

	return 1
}

// Poke represents the model of a poke
type Poke struct {
	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewPoke returns a new *Poke
func NewPoke() *Poke {

	return &Poke{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Poke) Identity() elemental.Identity {

	return PokeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Poke) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Poke) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *Poke) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Poke) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Poke) Doc() string {
	return `When available, poke can be used to update various information about the parent. For instance, for enforcers, poke will be use as the heartbeat.`
}

func (o *Poke) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Poke) Validate() error {

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
func (*Poke) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PokeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PokeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Poke) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PokeAttributesMap
}

// PokeAttributesMap represents the map of attribute for Poke.
var PokeAttributesMap = map[string]elemental.AttributeSpecification{}

// PokeLowerCaseAttributesMap represents the map of attribute for Poke.
var PokeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{}
