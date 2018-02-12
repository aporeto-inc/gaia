package models

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
)

// X509CertificateCheckIdentity represents the Identity of the object.
var X509CertificateCheckIdentity = elemental.Identity{
	Name:     "x509certificatecheck",
	Category: "x509certificatechecks",
}

// X509CertificateChecksList represents a list of X509CertificateChecks
type X509CertificateChecksList []*X509CertificateCheck

// ContentIdentity returns the identity of the objects in the list.
func (o X509CertificateChecksList) ContentIdentity() elemental.Identity {

	return X509CertificateCheckIdentity
}

// Copy returns a pointer to a copy the X509CertificateChecksList.
func (o X509CertificateChecksList) Copy() elemental.ContentIdentifiable {

	copy := append(X509CertificateChecksList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the X509CertificateChecksList.
func (o X509CertificateChecksList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(X509CertificateChecksList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*X509CertificateCheck))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o X509CertificateChecksList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o X509CertificateChecksList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o X509CertificateChecksList) Version() int {

	return 1
}

// X509CertificateCheck represents the model of a x509certificatecheck
type X509CertificateCheck struct {
	// ID contains the certificate serialNumber
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewX509CertificateCheck returns a new *X509CertificateCheck
func NewX509CertificateCheck() *X509CertificateCheck {

	return &X509CertificateCheck{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *X509CertificateCheck) Identity() elemental.Identity {

	return X509CertificateCheckIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *X509CertificateCheck) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *X509CertificateCheck) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *X509CertificateCheck) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *X509CertificateCheck) DefaultOrder() []string {

	return []string{}
}

func (o *X509CertificateCheck) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *X509CertificateCheck) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*X509CertificateCheck) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := X509CertificateCheckAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return X509CertificateCheckLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*X509CertificateCheck) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return X509CertificateCheckAttributesMap
}

// X509CertificateCheckAttributesMap represents the map of attribute for X509CertificateCheck.
var X509CertificateCheckAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID contains the certificate serialNumber`,
		Exposed:        true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
}

// X509CertificateCheckLowerCaseAttributesMap represents the map of attribute for X509CertificateCheck.
var X509CertificateCheckLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID contains the certificate serialNumber`,
		Exposed:        true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
}
