package gaia

import (
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// IPRecord represents the model of a iprecord
type IPRecord struct {
	// Number of time the port was accessed.
	Hits int `json:"hits" bson:"-" mapstructure:"hits,omitempty"`

	// Protocol used.
	Protocol int `json:"protocol" bson:"-" mapstructure:"protocol,omitempty"`

	// Date of the last access on that port.
	Timestamp time.Time `json:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewIPRecord returns a new *IPRecord
func NewIPRecord() *IPRecord {

	return &IPRecord{
		ModelVersion: 1,
	}
}

// Validate valides the current information stored into the structure.
func (o *IPRecord) Validate() error {

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
