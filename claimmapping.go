package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ClaimMapping represents the model of a claimmapping
type ClaimMapping struct {
	// Claim name is the name of the claim that must be mapped to an HTTP header.
	ClaimName string `json:"claimName" bson:"claimname" mapstructure:"claimName,omitempty"`

	// The target HTTP header where this claim name must be mapped.
	TargetHTTPHeader string `json:"targetHTTPHeader" bson:"targethttpheader" mapstructure:"targetHTTPHeader,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewClaimMapping returns a new *ClaimMapping
func NewClaimMapping() *ClaimMapping {

	return &ClaimMapping{
		ModelVersion: 1,
		Mutex:        &sync.Mutex{},
	}
}

// DeepCopy returns a deep copy if the ClaimMapping.
func (o *ClaimMapping) DeepCopy() *ClaimMapping {

	if o == nil {
		return nil
	}

	out := &ClaimMapping{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ClaimMapping.
func (o *ClaimMapping) DeepCopyInto(out *ClaimMapping) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ClaimMapping: %s", err))
	}

	*out = *target.(*ClaimMapping)
}

// Validate valides the current information stored into the structure.
func (o *ClaimMapping) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("claimName", o.ClaimName); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidatePattern("claimName", o.ClaimName, `^[a-zA-Z0-9-_/*#&@\+\$~:]+$`, `must be an alpha numerical character or '-', '_', '/', '*', '#', '&', '@', '_', '$' ~ or ':'`, true); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("targetHTTPHeader", o.TargetHTTPHeader); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidatePattern("targetHTTPHeader", o.TargetHTTPHeader, `^[a-zA-Z0-9-_/*#&@\+\$~:]+$`, `must be an alpha numerical character or '-', '_', '/', '*', '#', '&', '@', '_', '$' ~ or ':'`, true); err != nil {
		errors = append(errors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
