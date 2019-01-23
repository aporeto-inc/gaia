package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// HostService represents the model of a hostservice
type HostService struct {
	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// Name of the service.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// networkonly indicate the host service is of type network only.
	Networkonly bool `json:"networkonly" bson:"networkonly" mapstructure:"networkonly,omitempty"`

	// Services lists all protocols and ports a service is running.
	Services []*ProcessingUnitService `json:"services" bson:"services" mapstructure:"services,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewHostService returns a new *HostService
func NewHostService() *HostService {

	return &HostService{
		ModelVersion:   1,
		AssociatedTags: []string{},
		Networkonly:    true,
		Services:       []*ProcessingUnitService{},
	}
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *HostService) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *HostService) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// DeepCopy returns a deep copy if the HostService.
func (o *HostService) DeepCopy() *HostService {

	if o == nil {
		return nil
	}

	out := &HostService{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *HostService.
func (o *HostService) DeepCopyInto(out *HostService) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy HostService: %s", err))
	}

	*out = *target.(*HostService)
}

// Validate valides the current information stored into the structure.
func (o *HostService) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Services {
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
