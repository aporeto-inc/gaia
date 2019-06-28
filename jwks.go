package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// JWKSAlgValue represents the possible values for attribute "alg".
type JWKSAlgValue string

const (
	// JWKSAlgES256 represents the value ES256.
	JWKSAlgES256 JWKSAlgValue = "ES256"

	// JWKSAlgRS256 represents the value RS256.
	JWKSAlgRS256 JWKSAlgValue = "RS256"
)

// JWKSCurveValue represents the possible values for attribute "curve".
type JWKSCurveValue string

const (
	// JWKSCurveP256 represents the value P256.
	JWKSCurveP256 JWKSCurveValue = "P256"
)

// JWKSKtyValue represents the possible values for attribute "kty".
type JWKSKtyValue string

const (
	// JWKSKtyEC represents the value EC.
	JWKSKtyEC JWKSKtyValue = "EC"

	// JWKSKtyRSA represents the value RSA.
	JWKSKtyRSA JWKSKtyValue = "RSA"
)

// JWKS represents the model of a jwks
type JWKS struct {
	// Alg defines the algorithm used for signing as per the JWKS specification.
	Alg JWKSAlgValue `json:"alg" msgpack:"alg" bson:"-" mapstructure:"alg,omitempty"`

	// Curve is the curve type used for signing. Valid only when signing method is EC.
	Curve JWKSCurveValue `json:"curve" msgpack:"curve" bson:"-" mapstructure:"curve,omitempty"`

	// kid is the key ID as per the JWKS specification.
	Kid string `json:"kid" msgpack:"kid" bson:"-" mapstructure:"kid,omitempty"`

	// kty defines the key type as per the JWKS specification.
	Kty JWKSKtyValue `json:"kty" msgpack:"kty" bson:"-" mapstructure:"kty,omitempty"`

	// Modulo is the modulo value of an RSA public key. Valid only when the signing
	// method is RSA.
	Modulo string `json:"modulo" msgpack:"modulo" bson:"-" mapstructure:"modulo,omitempty"`

	// Use defines the use of the signing key as per the JWKS specification.
	Use string `json:"use" msgpack:"use" bson:"-" mapstructure:"use,omitempty"`

	// X is the X value of the public key. Valid only when signing method is EC.
	X string `json:"x" msgpack:"x" bson:"-" mapstructure:"x,omitempty"`

	// Y is the Y value of the public key. Valid only when signing method is EC.
	Y string `json:"y" msgpack:"y" bson:"-" mapstructure:"y,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewJWKS returns a new *JWKS
func NewJWKS() *JWKS {

	return &JWKS{
		ModelVersion: 1,
	}
}

// BleveType implements the bleve.Classifier Interface.
func (o *JWKS) BleveType() string {

	return "jwks"
}

// DeepCopy returns a deep copy if the JWKS.
func (o *JWKS) DeepCopy() *JWKS {

	if o == nil {
		return nil
	}

	out := &JWKS{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *JWKS.
func (o *JWKS) DeepCopyInto(out *JWKS) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy JWKS: %s", err))
	}

	*out = *target.(*JWKS)
}

// Validate valides the current information stored into the structure.
func (o *JWKS) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("alg", string(o.Alg), []string{"RS256", "ES256"}, true); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("curve", string(o.Curve), []string{"P256"}, true); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("kty", string(o.Kty), []string{"RSA", "EC"}, true); err != nil {
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
