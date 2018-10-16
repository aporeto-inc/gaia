package gaia

import (
	"sync"

	"go.aporeto.io/elemental"
)

// Credential represents the model of a credential
type Credential struct {
	// The ID of app credential.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The certificate data encoded in base64.
	Certificate string `json:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// The certificate authority data encoded in base64.
	CertificateAuthority string `json:"certificateAuthority" bson:"-" mapstructure:"certificateAuthority,omitempty"`

	// The certificate key data encoded in base64.
	CertificateKey string `json:"certificateKey" bson:"-" mapstructure:"certificateKey,omitempty"`

	// The name of app credential.
	Name string `json:"name" bson:"-" mapstructure:"name,omitempty"`

	// The namespace of app credential.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewCredential returns a new *Credential
func NewCredential() *Credential {

	return &Credential{
		ModelVersion: 1,
	}
}

// Validate valides the current information stored into the structure.
func (o *Credential) Validate() error {

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
