package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// DependencyMapIdentity represents the Identity of the object
var DependencyMapIdentity = elemental.Identity{
	Name:     "dependencymap",
	Category: "dependencymaps",
}

// DependencyMapsList represents a list of DependencyMaps
type DependencyMapsList []*DependencyMap

// DependencyMap represents the model of a dependencymap
type DependencyMap struct {
	// ID is the identifier of the object.
	ID string `json:"ID" cql:"-" bson:"-"`

	// edges are the edges of the map
	Edges map[string]*MapEdge `json:"edges" cql:"-" bson:"-"`

	// Groups provide information about the group values
	Groups map[string]map[string]string `json:"groups" cql:"-" bson:"-"`

	// nodes refers to the nodes of the map
	Nodes map[string]*MapNode `json:"nodes" cql:"-" bson:"-"`
}

// NewDependencyMap returns a new *DependencyMap
func NewDependencyMap() *DependencyMap {

	return &DependencyMap{
		Edges: map[string]*MapEdge{},
		Nodes: map[string]*MapNode{},
	}
}

// Identity returns the Identity of the object.
func (o *DependencyMap) Identity() elemental.Identity {

	return DependencyMapIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DependencyMap) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DependencyMap) SetIdentifier(ID string) {

	o.ID = ID
}

func (o *DependencyMap) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *DependencyMap) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredExternal("edges", o.Edges); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("edges", o.Edges); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("groups", o.Groups); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("groups", o.Groups); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredExternal("nodes", o.Nodes); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("nodes", o.Nodes); err != nil {
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
func (DependencyMap) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return DependencyMapAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (DependencyMap) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DependencyMapAttributesMap
}

// DependencyMapAttributesMap represents the map of attribute for DependencyMap.
var DependencyMapAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Edges": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `edges are the edges of the map`,
		Exposed:        true,
		Name:           "edges",
		ReadOnly:       true,
		Required:       true,
		SubType:        "edges_map",
		Type:           "external",
	},
	"Groups": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Groups provide information about the group values`,
		Exposed:        true,
		Name:           "groups",
		ReadOnly:       true,
		Required:       true,
		SubType:        "group_description_map",
		Type:           "external",
	},
	"Nodes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `nodes refers to the nodes of the map`,
		Exposed:        true,
		Name:           "nodes",
		ReadOnly:       true,
		Required:       true,
		SubType:        "nodes_map",
		Type:           "external",
	},
}
