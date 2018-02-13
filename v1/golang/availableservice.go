package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/types"
)

// AvailableServiceIdentity represents the Identity of the object.
var AvailableServiceIdentity = elemental.Identity{
	Name:     "availableservice",
	Category: "availableservices",
	Private:  false,
}

// AvailableServicesList represents a list of AvailableServices
type AvailableServicesList []*AvailableService

// ContentIdentity returns the identity of the objects in the list.
func (o AvailableServicesList) ContentIdentity() elemental.Identity {

	return AvailableServiceIdentity
}

// Copy returns a pointer to a copy the AvailableServicesList.
func (o AvailableServicesList) Copy() elemental.ContentIdentifiable {

	copy := append(AvailableServicesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AvailableServicesList.
func (o AvailableServicesList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(AvailableServicesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AvailableService))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AvailableServicesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AvailableServicesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o AvailableServicesList) Version() int {

	return 1
}

// AvailableService represents the model of a availableservice
type AvailableService struct {
	// Beta indicates if the service is in a beta version.
	Beta bool `json:"beta" bson:"-" mapstructure:"beta,omitempty"`

	// CategoryID of the service.
	CategoryID string `json:"categoryID" bson:"-" mapstructure:"categoryID,omitempty"`

	// Description of the service
	Description string `json:"description" bson:"-" mapstructure:"description,omitempty"`

	// Icon contains a base64 image for the available service.
	Icon string `json:"icon" bson:"-" mapstructure:"icon,omitempty"`

	// LongDescription contains a more detailed description of the service.
	LongDescription string `json:"longDescription" bson:"-" mapstructure:"longDescription,omitempty"`

	// Name of the Service
	Name string `json:"name" bson:"-" mapstructure:"name,omitempty"`

	// Parameters of the service the user can or has to specify
	Parameters []*types.ServiceParameter `json:"parameters" bson:"-" mapstructure:"parameters,omitempty"`

	// Title represents the title of the service.
	Title string `json:"title" bson:"-" mapstructure:"title,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewAvailableService returns a new *AvailableService
func NewAvailableService() *AvailableService {

	return &AvailableService{
		ModelVersion: 1,
		Parameters:   []*types.ServiceParameter{},
	}
}

// Identity returns the Identity of the object.
func (o *AvailableService) Identity() elemental.Identity {

	return AvailableServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AvailableService) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AvailableService) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *AvailableService) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *AvailableService) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AvailableService) Doc() string {
	return `AvailableService represents a service that is available for launching`
}

func (o *AvailableService) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *AvailableService) Validate() error {

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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*AvailableService) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AvailableServiceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AvailableServiceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AvailableService) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AvailableServiceAttributesMap
}

// AvailableServiceAttributesMap represents the map of attribute for AvailableService.
var AvailableServiceAttributesMap = map[string]elemental.AttributeSpecification{
	"Beta": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Beta",
		Description:    `Beta indicates if the service is in a beta version.`,
		Exposed:        true,
		Name:           "beta",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"CategoryID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `CategoryID of the service.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "categoryID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the service`,
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		Type:           "string",
	},
	"Icon": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Icon",
		Description:    `Icon contains a base64 image for the available service. `,
		Exposed:        true,
		Format:         "free",
		Name:           "icon",
		ReadOnly:       true,
		Type:           "string",
	},
	"LongDescription": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LongDescription",
		Description:    `LongDescription contains a more detailed description of the service.`,
		Exposed:        true,
		Format:         "free",
		Name:           "longDescription",
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the Service`,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		Type:           "string",
	},
	"Parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters of the service the user can or has to specify`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "service_parameters",
		Type:           "external",
	},
	"Title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		Description:    `Title represents the title of the service.`,
		Exposed:        true,
		Format:         "free",
		Name:           "title",
		Type:           "string",
	},
}

// AvailableServiceLowerCaseAttributesMap represents the map of attribute for AvailableService.
var AvailableServiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"beta": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Beta",
		Description:    `Beta indicates if the service is in a beta version.`,
		Exposed:        true,
		Name:           "beta",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"categoryid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `CategoryID of the service.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "categoryID",
		ReadOnly:       true,
		Type:           "string",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description of the service`,
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		Type:           "string",
	},
	"icon": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Icon",
		Description:    `Icon contains a base64 image for the available service. `,
		Exposed:        true,
		Format:         "free",
		Name:           "icon",
		ReadOnly:       true,
		Type:           "string",
	},
	"longdescription": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LongDescription",
		Description:    `LongDescription contains a more detailed description of the service.`,
		Exposed:        true,
		Format:         "free",
		Name:           "longDescription",
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the Service`,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		Type:           "string",
	},
	"parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters of the service the user can or has to specify`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "service_parameters",
		Type:           "external",
	},
	"title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		Description:    `Title represents the title of the service.`,
		Exposed:        true,
		Format:         "free",
		Name:           "title",
		Type:           "string",
	},
}
