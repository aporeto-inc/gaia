package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// PUNodeIdentity represents the Identity of the object.
var PUNodeIdentity = elemental.Identity{
	Name:     "punode",
	Category: "punodes",
	Package:  "squall",
	Private:  true,
}

// PUNodesList represents a list of PUNodes
type PUNodesList []*PUNode

// Identity returns the identity of the objects in the list.
func (o PUNodesList) Identity() elemental.Identity {

	return PUNodeIdentity
}

// Copy returns a pointer to a copy the PUNodesList.
func (o PUNodesList) Copy() elemental.Identifiables {

	copy := append(PUNodesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PUNodesList.
func (o PUNodesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PUNodesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PUNode))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PUNodesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PUNodesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o PUNodesList) Version() int {

	return 1
}

// PUNode represents the model of a punode
type PUNode struct {
	// Identifier of the pu.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Enforcement status of the pu.
	EnforcementStatus string `json:"enforcementStatus" bson:"-" mapstructure:"enforcementStatus,omitempty"`

	// Image of the pu.
	Image string `json:"image" bson:"-" mapstructure:"image,omitempty"`

	// Last update of the pu.
	LastUpdate time.Time `json:"lastUpdate" bson:"-" mapstructure:"lastUpdate,omitempty"`

	// Name of the pu.
	Name string `json:"name" bson:"-" mapstructure:"name,omitempty"`

	// Namespace of the pu.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Status of the pu.
	Status string `json:"status" bson:"-" mapstructure:"status,omitempty"`

	// Tags of the pu.
	Tags map[string]string `json:"tags" bson:"-" mapstructure:"tags,omitempty"`

	// Type of the pu.
	Type string `json:"type" bson:"-" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewPUNode returns a new *PUNode
func NewPUNode() *PUNode {

	return &PUNode{
		ModelVersion: 1,
		Tags:         map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *PUNode) Identity() elemental.Identity {

	return PUNodeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PUNode) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PUNode) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *PUNode) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *PUNode) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PUNode) Doc() string {
	return `Internal scaling down of a pu for dep map representation.`
}

func (o *PUNode) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *PUNode) Validate() error {

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
func (*PUNode) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PUNodeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PUNodeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PUNode) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PUNodeAttributesMap
}

// PUNodeAttributesMap represents the map of attribute for PUNode.
var PUNodeAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `Identifier of the pu.`,
		Exposed:        true,
		Name:           "ID",
		Type:           "string",
	},
	"EnforcementStatus": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcementStatus",
		Description:    `Enforcement status of the pu.`,
		Exposed:        true,
		Name:           "enforcementStatus",
		Type:           "string",
	},
	"Image": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Image",
		Description:    `Image of the pu.`,
		Exposed:        true,
		Name:           "image",
		Type:           "string",
	},
	"LastUpdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastUpdate",
		Description:    `Last update of the pu.`,
		Exposed:        true,
		Name:           "lastUpdate",
		Type:           "time",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the pu.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the pu.`,
		Exposed:        true,
		Name:           "namespace",
		Type:           "string",
	},
	"Status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `Status of the pu.`,
		Exposed:        true,
		Name:           "status",
		Type:           "string",
	},
	"Tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `Tags of the pu.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "tags_map",
		Type:           "external",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Type",
		Description:    `Type of the pu.`,
		Exposed:        true,
		Name:           "type",
		Type:           "string",
	},
}

// PUNodeLowerCaseAttributesMap represents the map of attribute for PUNode.
var PUNodeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `Identifier of the pu.`,
		Exposed:        true,
		Name:           "ID",
		Type:           "string",
	},
	"enforcementstatus": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcementStatus",
		Description:    `Enforcement status of the pu.`,
		Exposed:        true,
		Name:           "enforcementStatus",
		Type:           "string",
	},
	"image": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Image",
		Description:    `Image of the pu.`,
		Exposed:        true,
		Name:           "image",
		Type:           "string",
	},
	"lastupdate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LastUpdate",
		Description:    `Last update of the pu.`,
		Exposed:        true,
		Name:           "lastUpdate",
		Type:           "time",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the pu.`,
		Exposed:        true,
		Name:           "name",
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the pu.`,
		Exposed:        true,
		Name:           "namespace",
		Type:           "string",
	},
	"status": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Status",
		Description:    `Status of the pu.`,
		Exposed:        true,
		Name:           "status",
		Type:           "string",
	},
	"tags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Tags",
		Description:    `Tags of the pu.`,
		Exposed:        true,
		Name:           "tags",
		SubType:        "tags_map",
		Type:           "external",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Type",
		Description:    `Type of the pu.`,
		Exposed:        true,
		Name:           "type",
		Type:           "string",
	},
}
