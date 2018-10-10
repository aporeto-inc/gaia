package gaia

import (
	"sync"

	"go.aporeto.io/elemental"
)

// ProcessingUnitService represents the model of a processingunitservice
type ProcessingUnitService struct {
	// porst contains the list of allowed ports and ranges.
	Ports string `json:"ports" bson:"-" mapstructure:"ports,omitempty"`

	// Protocol used by the service.
	Protocol int `json:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewProcessingUnitService returns a new *ProcessingUnitService
func NewProcessingUnitService() *ProcessingUnitService {

	return &ProcessingUnitService{
		ModelVersion: 1,
	}
}

// Validate valides the current information stored into the structure.
func (o *ProcessingUnitService) Validate() error {

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
