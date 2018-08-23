package gaia

import (
	"fmt"
	"sync"

	"go.aporeto.io/elemental"
)

// DependencyMapIndexes lists the attribute compound indexes.
var DependencyMapIndexes = [][]string{}

// DependencyMapIdentity represents the Identity of the object.
var DependencyMapIdentity = elemental.Identity{
	Name:     "dependencymap",
	Category: "dependencymaps",
	Private:  false,
}

// DependencyMapsList represents a list of DependencyMaps
type DependencyMapsList []*DependencyMap

// Identity returns the identity of the objects in the list.
func (o DependencyMapsList) Identity() elemental.Identity {

	return DependencyMapIdentity
}

// Copy returns a pointer to a copy the DependencyMapsList.
func (o DependencyMapsList) Copy() elemental.Identifiables {

	copy := append(DependencyMapsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DependencyMapsList.
func (o DependencyMapsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

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
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// claims represents a user or a script that have accessed an api.
	Claims map[string][]string `json:"claims" bson:"-" mapstructure:"claims,omitempty"`

	// edges are the edges of the map.
	Edges map[string]*GraphEdge `json:"edges" bson:"-" mapstructure:"edges,omitempty"`

	// Groups provide information about the group values.
	Groups map[string]*GraphGroup `json:"groups" bson:"-" mapstructure:"groups,omitempty"`

	// nodes refers to the nodes of the map.
	Nodes map[string]*GraphNode `json:"nodes" bson:"-" mapstructure:"nodes,omitempty"`

	// viewSuggestions provides suggestion of views based on relevant tags.
	ViewSuggestions []string `json:"viewSuggestions" bson:"-" mapstructure:"viewSuggestions,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewDependencyMap returns a new *DependencyMap
func NewDependencyMap() *DependencyMap {

	return &DependencyMap{
		ModelVersion:    1,
		Claims:          map[string][]string{},
		Edges:           map[string]*GraphEdge{},
		Groups:          map[string]*GraphGroup{},
		Nodes:           map[string]*GraphNode{},
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
  "/dependencymaps?startAbsolute=1489132800000&endAbsolute=1489219200000".`
}

func (o *DependencyMap) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *DependencyMap) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Edges {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	for _, sub := range o.Groups {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
	}

	for _, sub := range o.Nodes {
		if err := sub.Validate(); err != nil {
			errors = append(errors, err)
		}
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
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `claims represents a user or a script that have accessed an api.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "graphclaims_map",
		Type:           "external",
	},
	"Edges": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Edges",
		Description:    `edges are the edges of the map.`,
		Exposed:        true,
		Name:           "edges",
		ReadOnly:       true,
		SubType:        "graphedge",
		Type:           "refMap",
	},
	"Groups": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Groups",
		Description:    `Groups provide information about the group values.`,
		Exposed:        true,
		Name:           "groups",
		ReadOnly:       true,
		SubType:        "graphgroup",
		Type:           "refMap",
	},
	"Nodes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Nodes",
		Description:    `nodes refers to the nodes of the map.`,
		Exposed:        true,
		Name:           "nodes",
		ReadOnly:       true,
		SubType:        "graphnode",
		Type:           "refMap",
	},
	"ViewSuggestions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ViewSuggestions",
		Description:    `viewSuggestions provides suggestion of views based on relevant tags.`,
		Exposed:        true,
		Name:           "viewSuggestions",
		ReadOnly:       true,
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
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
	},
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `claims represents a user or a script that have accessed an api.`,
		Exposed:        true,
		Name:           "claims",
		ReadOnly:       true,
		SubType:        "graphclaims_map",
		Type:           "external",
	},
	"edges": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Edges",
		Description:    `edges are the edges of the map.`,
		Exposed:        true,
		Name:           "edges",
		ReadOnly:       true,
		SubType:        "graphedge",
		Type:           "refMap",
	},
	"groups": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Groups",
		Description:    `Groups provide information about the group values.`,
		Exposed:        true,
		Name:           "groups",
		ReadOnly:       true,
		SubType:        "graphgroup",
		Type:           "refMap",
	},
	"nodes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Nodes",
		Description:    `nodes refers to the nodes of the map.`,
		Exposed:        true,
		Name:           "nodes",
		ReadOnly:       true,
		SubType:        "graphnode",
		Type:           "refMap",
	},
	"viewsuggestions": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ViewSuggestions",
		Description:    `viewSuggestions provides suggestion of views based on relevant tags.`,
		Exposed:        true,
		Name:           "viewSuggestions",
		ReadOnly:       true,
		SubType:        "view_suggestions",
		Type:           "external",
	},
}
