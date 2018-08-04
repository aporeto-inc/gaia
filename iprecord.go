package gaia

import (
	"sync"

	"go.aporeto.io/elemental"
)

// IPRecord represents the model of a iprecord
type IPRecord struct {
	// Actual IP Address.
	IP string `json:"IP" bson:"-" mapstructure:"IP,omitempty"`

	// List of actions applied from that IP.
	Actions []string `json:"actions" bson:"-" mapstructure:"actions,omitempty"`

	// City of the IP.
	City string `json:"city" bson:"-" mapstructure:"city,omitempty"`

	// Country of the IP.
	Country string `json:"country" bson:"-" mapstructure:"country,omitempty"`

	// List of ports applied used.
	DestinationPorts []string `json:"destinationPorts" bson:"-" mapstructure:"destinationPorts,omitempty"`

	// Number of hits.
	Hits int `json:"hits" bson:"-" mapstructure:"hits,omitempty"`

	// Eventual list of hostnames associated to the IP.
	Hostnames []string `json:"hostnames" bson:"-" mapstructure:"hostnames,omitempty"`

	// Protocol used.
	L4Protocol string `json:"l4Protocol" bson:"-" mapstructure:"l4Protocol,omitempty"`

	// Latitude of the IP.
	Latitude float64 `json:"latitude" bson:"-" mapstructure:"latitude,omitempty"`

	// Longitude of the IP.
	Longitude float64 `json:"longitude" bson:"-" mapstructure:"longitude,omitempty"`

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
