package gaia

import (
	"sync"

	"go.aporeto.io/elemental"
)

// GraphGroup represents the model of a graphgroup
type GraphGroup struct {
	// Identifier of the group.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Color to use for the group.
	Color string `json:"color" bson:"-" mapstructure:"color,omitempty"`

	// List of tag that was used to create this group.
	Match [][]string `json:"match" bson:"-" mapstructure:"match,omitempty"`

	// Name of the group.
	Name string `json:"name" bson:"-" mapstructure:"name,omitempty"`

	// ID of the parent group if any.
	ParentID string `json:"parentID" bson:"-" mapstructure:"parentID,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewGraphGroup returns a new *GraphGroup
func NewGraphGroup() *GraphGroup {

	return &GraphGroup{
		ModelVersion: 1,
		Match:        [][]string{},
	}
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
