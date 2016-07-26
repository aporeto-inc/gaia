package gaia

import "github.com/aporeto-inc/elemental"

const (
	MapEdgeAttributeNameID            elemental.AttributeSpecificationNameKey = "mapedge/ID"
	MapEdgeAttributeNameDestinationID elemental.AttributeSpecificationNameKey = "mapedge/destinationID"
	MapEdgeAttributeNameLabels        elemental.AttributeSpecificationNameKey = "mapedge/labels"
	MapEdgeAttributeNameName          elemental.AttributeSpecificationNameKey = "mapedge/name"
	MapEdgeAttributeNameSourceID      elemental.AttributeSpecificationNameKey = "mapedge/sourceID"
)

// MapEdgeIdentity represents the Identity of the object
var MapEdgeIdentity = elemental.Identity{
	Name:     "mapedge",
	Category: "mapedges",
}

// MapEdgesList represents a list of MapEdges
type MapEdgesList []*MapEdge

// MapEdge represents the model of a mapedge
type MapEdge struct {
	ID            string   `json:"ID,omitempty" cql:"-"`
	DestinationID string   `json:"destinationID,omitempty" cql:"-"`
	Labels        []string `json:"labels,omitempty" cql:"-"`
	Name          string   `json:"name,omitempty" cql:"-"`
	SourceID      string   `json:"sourceID,omitempty" cql:"-"`
}

// NewMapEdge returns a new *MapEdge
func NewMapEdge() *MapEdge {

	return &MapEdge{}
}

// Identity returns the Identity of the object.
func (o *MapEdge) Identity() elemental.Identity {

	return MapEdgeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *MapEdge) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *MapEdge) SetIdentifier(ID string) {

	o.ID = ID
}

// Validate valides the current information stored into the structure.
func (o *MapEdge) Validate() elemental.Errors {

	errors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("destinationID", o.DestinationID); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("sourceID", o.SourceID); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (o MapEdge) SpecificationForAttribute(name elemental.AttributeSpecificationNameKey) elemental.AttributeSpecification {

	return MapEdgeAttributesMap[name]
}

var MapEdgeAttributesMap = map[elemental.AttributeSpecificationNameKey]elemental.AttributeSpecification{
	MapEdgeAttributeNameID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
	MapEdgeAttributeNameDestinationID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "destinationID",
		Required:       true,
		Type:           "string",
	},
	MapEdgeAttributeNameLabels: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "labels",
		Required:       true,
		SubType:        "tag_list",
		Type:           "external",
	},
	MapEdgeAttributeNameName: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		Required:       true,
		Type:           "string",
	},
	MapEdgeAttributeNameSourceID: elemental.AttributeSpecification{
		AllowedChoices: []string{},
		CreationOnly:   true,
		Exposed:        true,
		Format:         "free",
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
}
