package servicelaunchermodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// ServiceIdentity represents the Identity of the object
var ServiceIdentity = elemental.Identity{
	Name:     "service",
	Category: "services",
}

// ServicesList represents a list of Services
type ServicesList []*Service

// ContentIdentity returns the identity of the objects in the list.
func (o ServicesList) ContentIdentity() elemental.Identity {

	return ServiceIdentity
}

// Copy returns a pointer to a copy the ServicesList.
func (o ServicesList) Copy() elemental.ContentIdentifiable {

	copy := append(ServicesList{}, o...)
	return &copy
}

// List converts the object to an elemental.IdentifiablesList.
func (o ServicesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ServicesList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o ServicesList) Version() int {

	return 1
}

// Service represents the model of a service
type Service struct {
	// ID of the service
	ID string `json:"ID" bson:"id"`

	// Name of the service
	Name string `json:"name" bson:"name"`

	// Type of the service
	Type string `json:"type" bson:"type"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewService returns a new *Service
func NewService() *Service {

	return &Service{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Service) Identity() elemental.Identity {

	return ServiceIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Service) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Service) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Service) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Service) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Service) Doc() string {
	return `Service represents a service that can be launched`
}

func (o *Service) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Service) Validate() error {

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
func (*Service) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ServiceAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ServiceLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Service) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ServiceAttributesMap
}

// ServiceAttributesMap represents the map of attribute for Service.
var ServiceAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `ID of the service`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name of the service`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Type": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Type of the service`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "type",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// ServiceLowerCaseAttributesMap represents the map of attribute for Service.
var ServiceLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `ID of the service`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name of the service`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"type": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Type of the service`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "type",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
