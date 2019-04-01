package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TraceMode represents the model of a tracemode
type TraceMode struct {
	// IPTables instructs the enforcers to provide an iptables trace for a PU.
	IPTables bool `json:"IPTables" bson:"iptables" mapstructure:"IPTables,omitempty"`

	// Instructs the enforcer to send records for all
	// application initiated connections.
	ApplicationConnections bool `json:"applicationConnections" bson:"applicationconnections" mapstructure:"applicationConnections,omitempty"`

	// Determines the length of the time interval that the trace must be
	// enabled.
	Interval string `json:"interval" bson:"interval" mapstructure:"interval,omitempty"`

	// Instructs the enforcer to send records for all network
	// initiated connections.
	NetworkConnections bool `json:"networkConnections" bson:"networkconnections" mapstructure:"networkConnections,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewTraceMode returns a new *TraceMode
func NewTraceMode() *TraceMode {

	return &TraceMode{
		ModelVersion: 1,
		Interval:     "10s",
	}
}

// DeepCopy returns a deep copy if the TraceMode.
func (o *TraceMode) DeepCopy() *TraceMode {

	if o == nil {
		return nil
	}

	out := &TraceMode{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TraceMode.
func (o *TraceMode) DeepCopyInto(out *TraceMode) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TraceMode: %s", err))
	}

	*out = *target.(*TraceMode)
}

// Validate valides the current information stored into the structure.
func (o *TraceMode) Validate() error {

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
