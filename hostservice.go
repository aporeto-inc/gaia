package gaia

import (
	"sync"

	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/types"
)

// HostService represents the model of a hostservice
type HostService struct {
	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Name of the service.
	Name string `json:"name" bson:"-" mapstructure:"name,omitempty"`

	// Services lists all protocols and ports a service is running.
	Services types.ProcessingUnitServicesList `json:"services" bson:"services" mapstructure:"services,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewHostService returns a new *HostService
func NewHostService() *HostService {

	return &HostService{
		ModelVersion:   1,
		AssociatedTags: []string{},
		Services:       types.ProcessingUnitServicesList{},
	}
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *HostService) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the given AssociatedTags of the receiver.
func (o *HostService) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// Validate valides the current information stored into the structure.
func (o *HostService) Validate() error {

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
