package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// GraphGroup represents the model of a graphgroup
type GraphGroup struct {
	// Identifier of the group.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Color to use for the group.
	Color string `json:"color" msgpack:"color" bson:"-" mapstructure:"color,omitempty"`

	// List of tag that was used to create this group.
	Match [][]string `json:"match" msgpack:"match" bson:"-" mapstructure:"match,omitempty"`

	// Name of the group.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	// ID of the parent group if any.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"-" mapstructure:"parentID,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewGraphGroup returns a new *GraphGroup
func NewGraphGroup() *GraphGroup {

	return &GraphGroup{
		ModelVersion: 1,
		Match:        [][]string{},
	}
}

// DeepCopy returns a deep copy if the GraphGroup.
func (o *GraphGroup) DeepCopy() *GraphGroup {

	if o == nil {
		return nil
	}

	out := &GraphGroup{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *GraphGroup.
func (o *GraphGroup) DeepCopyInto(out *GraphGroup) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy GraphGroup: %s", err))
	}

	*out = *target.(*GraphGroup)
}

// Validate valides the current information stored into the structure.
func (o *GraphGroup) Validate() error {

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
