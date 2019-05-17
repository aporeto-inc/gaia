package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// Endpoint represents the model of a endpoint
type Endpoint struct {
	// URI of the exposed API.
	URI string `json:"URI" msgpack:"URI" bson:"uri" mapstructure:"URI,omitempty"`

	// AllowedScopes authorized to access the API.
	AllowedScopes [][]string `json:"allowedScopes" msgpack:"allowedScopes" bson:"allowedscopes" mapstructure:"allowedScopes,omitempty"`

	// methods exposed to access the API.
	Methods []string `json:"methods" msgpack:"methods" bson:"methods" mapstructure:"methods,omitempty"`

	// public defines if the api is public or not.
	Public bool `json:"public" msgpack:"public" bson:"public" mapstructure:"public,omitempty"`

	// Scopes is deprecated.
	Scopes []string `json:"scopes" msgpack:"scopes" bson:"-" mapstructure:"scopes,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEndpoint returns a new *Endpoint
func NewEndpoint() *Endpoint {

	return &Endpoint{
		ModelVersion:  1,
		AllowedScopes: [][]string{},
		Methods:       []string{},
		Scopes:        []string{},
	}
}

// DeepCopy returns a deep copy if the Endpoint.
func (o *Endpoint) DeepCopy() *Endpoint {

	if o == nil {
		return nil
	}

	out := &Endpoint{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Endpoint.
func (o *Endpoint) DeepCopyInto(out *Endpoint) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Endpoint: %s", err))
	}

	*out = *target.(*Endpoint)
}

// Validate valides the current information stored into the structure.
func (o *Endpoint) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateHTTPMethods("methods", o.Methods); err != nil {
		errors = errors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
