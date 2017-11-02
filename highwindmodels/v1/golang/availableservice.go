package highwindmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

import "github.com/aporeto-inc/gaia/highwindmodels/v1/golang/types"

// AvailableServiceIdentity represents the Identity of the object
var AvailableServiceIdentity = elemental.Identity{
	Name:     "availableservice",
	Category: "availableservices",
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
	// Category of the service.
	Category string `json:"category" bson:"-"`

	// Description of the service
	Description string `json:"description" bson:"-"`

	// LongDescription contains a more detailed description of the service.
	LongDescription string `json:"longDescription" bson:"-"`

	// Name of the Service
	Name string `json:"name" bson:"-"`

	// Parameters of the service the user can or has to specify
	Parameters []*types.ServiceParameter `json:"parameters" bson:"-"`

	// Title represents the title of the service.
	Title string `json:"title" bson:"-"`

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
func (o *AvailableService) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
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
	"Category": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Category of the service.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "category",
		ReadOnly:       true,
		Type:           "string",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Description of the service`,
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		Type:           "string",
	},
	"LongDescription": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LongDescription contains a more detailed description of the service.`,
		Exposed:        true,
		Format:         "free",
		Name:           "longDescription",
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name of the Service`,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		Type:           "string",
	},
	"Parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Parameters of the service the user can or has to specify`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "service_parameters",
		Type:           "external",
	},
	"Title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Title represents the title of the service.`,
		Exposed:        true,
		Format:         "free",
		Name:           "title",
		Type:           "string",
	},
}

// AvailableServiceLowerCaseAttributesMap represents the map of attribute for AvailableService.
var AvailableServiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"category": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Category of the service.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "category",
		ReadOnly:       true,
		Type:           "string",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Description of the service`,
		Exposed:        true,
		Format:         "free",
		Name:           "description",
		Type:           "string",
	},
	"longdescription": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `LongDescription contains a more detailed description of the service.`,
		Exposed:        true,
		Format:         "free",
		Name:           "longDescription",
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name of the Service`,
		Exposed:        true,
		Format:         "free",
		Name:           "name",
		Type:           "string",
	},
	"parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Parameters of the service the user can or has to specify`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "service_parameters",
		Type:           "external",
	},
	"title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Title represents the title of the service.`,
		Exposed:        true,
		Format:         "free",
		Name:           "title",
		Type:           "string",
	},
}
