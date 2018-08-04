package gaia

import (
	"sync"

	"go.aporeto.io/elemental"
)

// GraphPolicyInfo represents the model of a graphpolicyinfo
type GraphPolicyInfo struct {
	// Number of time the policy has been hit.
	Count int `json:"count" bson:"-" mapstructure:"count,omitempty"`

	// Namespace of the policy.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewGraphPolicyInfo returns a new *GraphPolicyInfo
func NewGraphPolicyInfo() *GraphPolicyInfo {

	return &GraphPolicyInfo{
		ModelVersion: 1,
	}
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
