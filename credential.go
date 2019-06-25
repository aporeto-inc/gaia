package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// Credential represents the model of a credential
type Credential struct {
	// The Aporeto API URL.
	APIURL string `json:"APIURL" msgpack:"APIURL" bson:"-" mapstructure:"APIURL,omitempty"`

	// The ID of app credential.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The certificate data encoded in base64.
	Certificate string `json:"certificate" msgpack:"certificate" bson:"-" mapstructure:"certificate,omitempty"`

	// The certificate authority data encoded in base64.
	CertificateAuthority string `json:"certificateAuthority" msgpack:"certificateAuthority" bson:"-" mapstructure:"certificateAuthority,omitempty"`

	// The certificate key data encoded in base64.
	CertificateKey string `json:"certificateKey" msgpack:"certificateKey" bson:"-" mapstructure:"certificateKey,omitempty"`

	// The name of app credential.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	// The namespace of app credential.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCredential returns a new *Credential
func NewCredential() *Credential {

	return &Credential{
		ModelVersion: 1,
	}
}

// BleveType implements the bleve.Classifier Interface.
func (o *Credential) BleveType() string {

	return "credential"
}

// DeepCopy returns a deep copy if the Credential.
func (o *Credential) DeepCopy() *Credential {

	if o == nil {
		return nil
	}

	out := &Credential{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Credential.
func (o *Credential) DeepCopyInto(out *Credential) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Credential: %s", err))
	}

	*out = *target.(*Credential)
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
