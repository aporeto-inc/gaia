package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// GraphPolicyInfo represents the model of a graphpolicyinfo
type GraphPolicyInfo struct {
	// Number of time the policy has been hit.
	Count int `json:"count" bson:"-" mapstructure:"count,omitempty"`

	// Namespace of the policy.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewGraphPolicyInfo returns a new *GraphPolicyInfo
func NewGraphPolicyInfo() *GraphPolicyInfo {

	return &GraphPolicyInfo{
		ModelVersion: 1,
		Mutex:        &sync.Mutex{},
	}
}

// DeepCopy returns a deep copy if the GraphPolicyInfo.
func (o *GraphPolicyInfo) DeepCopy() *GraphPolicyInfo {

	if o == nil {
		return nil
	}

	out := &GraphPolicyInfo{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *GraphPolicyInfo.
func (o *GraphPolicyInfo) DeepCopyInto(out *GraphPolicyInfo) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy GraphPolicyInfo: %s", err))
	}

	*out = *target.(*GraphPolicyInfo)
}

// Validate valides the current information stored into the structure.
func (o *GraphPolicyInfo) Validate() error {

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
