package gaia

import "github.com/aporeto-inc/elemental"

const (
	DependencyMapAttributeNameID    elemental.AttributeSpecificationNameKey = "dependencymap/ID"
	DependencyMapAttributeNameEdges elemental.AttributeSpecificationNameKey = "dependencymap/edges"
	DependencyMapAttributeNameNodes elemental.AttributeSpecificationNameKey = "dependencymap/nodes"
)

// DependencyMapIdentity represents the Identity of the object
var DependencyMapIdentity = elemental.Identity{
	Name:     "dependencymap",
	Category: "dependencymaps",
}

// DependencyMapsList represents a list of DependencyMaps
type DependencyMapsList []*DependencyMap

// DependencyMap represents the model of a dependencymap
type DependencyMap struct {
	ID    string    `json:"ID,omitempty" cql:"-"`
	Edges []MapEdge `json:"edges,omitempty" cql:"-"`
	Nodes []MapNode `json:"nodes,omitempty" cql:"-"`
}

// NewDependencyMap returns a new *DependencyMap
func NewDependencyMap() *DependencyMap {

	return &DependencyMap{
		Edges: []MapEdge{},

		Nodes: []MapNode{},
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

// Validate valides the current information stored into the structure.
func (o *DependencyMap) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o DependencyMap) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return DependencyMapAttributesMap[name]
}

var DependencyMapAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	DependencyMapAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
	DependencyMapAttributeNameEdges: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "edges",
		Required:       true,
		SubType:        "edge_list",
		Type:           "external",
	},
	DependencyMapAttributeNameNodes: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "nodes",
		Required:       true,
		SubType:        "node_list",
		Type:           "external",
	},
}
