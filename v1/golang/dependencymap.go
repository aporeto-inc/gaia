package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/types"
)

// DependencyMapIdentity represents the Identity of the object.
var DependencyMapIdentity = elemental.Identity{
	Name:     "dependencymap",
	Category: "dependencymaps",
	Private:  false,
}

// DependencyMapsList represents a list of DependencyMaps
type DependencyMapsList []*DependencyMap

// ContentIdentity returns the identity of the objects in the list.
func (o DependencyMapsList) ContentIdentity() elemental.Identity {

	return DependencyMapIdentity
}

// Copy returns a pointer to a copy the DependencyMapsList.
func (o DependencyMapsList) Copy() elemental.ContentIdentifiable {

	copy := append(DependencyMapsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DependencyMapsList.
func (o DependencyMapsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(DependencyMapsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*DependencyMap))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o DependencyMapsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DependencyMapsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o DependencyMapsList) Version() int {

	return 1
}

// DependencyMap represents the model of a dependencymap
type DependencyMap struct {
	// edges are the edges of the map
	Edges types.GraphEdgeMap `json:"edges" bson:"-" mapstructure:"edges,omitempty"`

	// Groups provide information about the group values
	Groups types.GraphGroupMap `json:"groups" bson:"-" mapstructure:"groups,omitempty"`

	// nodes refers to the nodes of the map
	Nodes types.GraphNodeMap `json:"nodes" bson:"-" mapstructure:"nodes,omitempty"`

	// viewSuggestions provides suggestion of views based on relevant tags.
	ViewSuggestions []string `json:"viewSuggestions" bson:"-" mapstructure:"viewSuggestions,omitempty"`

	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewDependencyMap returns a new *DependencyMap
func NewDependencyMap() *DependencyMap {

	return &DependencyMap{
		ModelVersion:    1,
		Edges:           types.GraphEdgeMap{},
		Groups:          types.GraphGroupMap{},
		Nodes:           types.GraphNodeMap{},
		ViewSuggestions: []string{},
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
func (o *DependencyMap) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *DependencyMap) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *DependencyMap) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *DependencyMap) Doc() string {
	return `This api returns a data structure representing the graph of all processing units
and their connections in a particular namespace, in a given time window. To pass
the time window you can use the query parameters 'startAbsolute', 'endAbsolute',
'startRelative', 'endRelative'.

For example
  "/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000"`
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

	if err := elemental.ValidateRequiredExternal("viewSuggestions", o.ViewSuggestions); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredExternal("viewSuggestions", o.ViewSuggestions); err != nil {
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
func (*DependencyMap) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DependencyMapAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DependencyMapLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DependencyMap) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DependencyMapAttributesMap
}

// DependencyMapAttributesMap represents the map of attribute for DependencyMap.
var DependencyMapAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"Edges": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Edges",
		Description:    `edges are the edges of the map`,
		Exposed:        true,
		Name:           "edges",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphedges_map",
		Type:           "external",
	},
	"Groups": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Groups",
		Description:    `Groups provide information about the group values`,
		Exposed:        true,
		Name:           "groups",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphgroups_map",
		Type:           "external",
	},
	"Nodes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Nodes",
		Description:    `nodes refers to the nodes of the map`,
		Exposed:        true,
		Name:           "nodes",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphnodes_map",
		Type:           "external",
	},
	"ViewSuggestions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ViewSuggestions",
		Description:    `viewSuggestions provides suggestion of views based on relevant tags.`,
		Exposed:        true,
		Name:           "viewSuggestions",
		ReadOnly:       true,
		Required:       true,
		SubType:        "view_suggestions",
		Type:           "external",
	},
}

// DependencyMapLowerCaseAttributesMap represents the map of attribute for DependencyMap.
var DependencyMapLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"edges": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Edges",
		Description:    `edges are the edges of the map`,
		Exposed:        true,
		Name:           "edges",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphedges_map",
		Type:           "external",
	},
	"groups": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Groups",
		Description:    `Groups provide information about the group values`,
		Exposed:        true,
		Name:           "groups",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphgroups_map",
		Type:           "external",
	},
	"nodes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Nodes",
		Description:    `nodes refers to the nodes of the map`,
		Exposed:        true,
		Name:           "nodes",
		ReadOnly:       true,
		Required:       true,
		SubType:        "graphnodes_map",
		Type:           "external",
	},
	"viewsuggestions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ViewSuggestions",
		Description:    `viewSuggestions provides suggestion of views based on relevant tags.`,
		Exposed:        true,
		Name:           "viewSuggestions",
		ReadOnly:       true,
		Required:       true,
		SubType:        "view_suggestions",
		Type:           "external",
	},
}
