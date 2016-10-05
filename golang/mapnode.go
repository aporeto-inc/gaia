package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// MapNodeTypeValue represents the possible values for attribute "type".
type MapNodeTypeValue string

const (
	// MapNodeTypeContainer represents the value Container.
	MapNodeTypeContainer MapNodeTypeValue = "Container"

	// MapNodeTypeVolume represents the value Volume.
	MapNodeTypeVolume MapNodeTypeValue = "Volume"
)

// MapNodeIdentity represents the Identity of the object
var MapNodeIdentity = elemental.Identity{
	Name:     "mapnode",
	Category: "mapnodes",
}

// MapNodesList represents a list of MapNodes
type MapNodesList []*MapNode

// MapNode represents the model of a mapnode
type MapNode struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"-"`

	// Groups for organizing resources
	Groups []string `json:"groups" cql:"-"`

	// Name is the name of the entity
	Name string `json:"name" cql:"name,omitempty"`

	// Type of the resource represented in the map
	Type MapNodeTypeValue `json:"type" cql:"-"`
}

// NewMapNode returns a new *MapNode
func NewMapNode() *MapNode {

	return &MapNode{}
}

// Identity returns the Identity of the object.
func (o *MapNode) Identity() elemental.Identity {

	return MapNodeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MapNode) Identifier() string {

	return o.ID
}

func (o *MapNode) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MapNode) SetIdentifier(ID string) {

	o.ID = ID
}

// GetName returns the name of the receiver
func (o *MapNode) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *MapNode) SetName(name string) {
	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *MapNode) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"Container", "Volume"}, true); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o MapNode) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return MapNodeAttributesMap[name]
}

// MapNodeAttributesMap represents the map of attribute for MapNode.
var MapNodeAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"Groups": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Exposed:        true,
		Name:           "groups",
		ReadOnly:       true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
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
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{"Container", "Volume"},
		Autogenerated:  true,
		Exposed:        true,
		Name:           "type",
		ReadOnly:       true,
		Type:           "enum",
	},
}
