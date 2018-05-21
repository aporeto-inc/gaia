package gaia

import (
	"fmt"
	"sync"

	"github.com/aporeto-inc/elemental"
	"github.com/aporeto-inc/gaia/v1/golang/types"
)

// AppIdentity represents the Identity of the object.
var AppIdentity = elemental.Identity{
	Name:     "app",
	Category: "apps",
	Private:  false,
}

// AppsList represents a list of Apps
type AppsList []*App

// ContentIdentity returns the identity of the objects in the list.
func (o AppsList) ContentIdentity() elemental.Identity {

	return AppIdentity
}

// Copy returns a pointer to a copy the AppsList.
func (o AppsList) Copy() elemental.ContentIdentifiable {

	copy := append(AppsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AppsList.
func (o AppsList) Append(objects ...elemental.Identifiable) elemental.ContentIdentifiable {

	out := append(AppsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*App))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AppsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AppsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Version returns the version of the content.
func (o AppsList) Version() int {

	return 1
}

// App represents the model of a app
type App struct {
	// Beta indicates if the app is in a beta version.
	Beta bool `json:"beta" bson:"-" mapstructure:"beta,omitempty"`

	// CategoryID of the app.
	CategoryID string `json:"categoryID" bson:"-" mapstructure:"categoryID,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// Icon contains a base64 image for the app.
	Icon string `json:"icon" bson:"-" mapstructure:"icon,omitempty"`

	// LongDescription contains a more detailed description of the app.
	LongDescription string `json:"longDescription" bson:"-" mapstructure:"longDescription,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// Parameters of the app the user can or has to specify.
	Parameters []*types.AppParameter `json:"parameters" bson:"-" mapstructure:"parameters,omitempty"`

	// Title represents the title of the app.
	Title string `json:"title" bson:"-" mapstructure:"title,omitempty"`

	// VersionParameters contains parameters for each available version.
	VersionParameters map[string][]*types.AppParameter `json:"versionParameters" bson:"-" mapstructure:"versionParameters,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewApp returns a new *App
func NewApp() *App {

	return &App{
		ModelVersion:      1,
		Parameters:        []*types.AppParameter{},
		VersionParameters: map[string][]*types.AppParameter{},
	}
}

// Identity returns the Identity of the object.
func (o *App) Identity() elemental.Identity {

	return AppIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *App) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *App) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *App) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *App) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *App) Doc() string {
	return `App represents an application that can be installed.`
}

func (o *App) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the Name of the receiver.
func (o *App) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *App) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *App) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*App) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AppAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AppLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*App) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AppAttributesMap
}

// AppAttributesMap represents the map of attribute for App.
var AppAttributesMap = map[string]elemental.AttributeSpecification{
	"Beta": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Beta",
		Description:    `Beta indicates if the app is in a beta version.`,
		Exposed:        true,
		Name:           "beta",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"CategoryID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `CategoryID of the app.`,
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
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Format:         "free",
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Icon": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Icon",
		Description:    `Icon contains a base64 image for the app.`,
		Exposed:        true,
		Format:         "free",
		Name:           "icon",
		ReadOnly:       true,
		Type:           "string",
	},
	"LongDescription": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LongDescription",
		Description:    `LongDescription contains a more detailed description of the app.`,
		Exposed:        true,
		Format:         "free",
		Name:           "longDescription",
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters of the app the user can or has to specify.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "app_parameters",
		Type:           "external",
	},
	"Title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		Description:    `Title represents the title of the app.`,
		Exposed:        true,
		Format:         "free",
		Name:           "title",
		Type:           "string",
	},
	"VersionParameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "VersionParameters",
		Description:    `VersionParameters contains parameters for each available version.`,
		Exposed:        true,
		Name:           "versionParameters",
		SubType:        "app_versionparameters",
		Type:           "external",
	},
}

// AppLowerCaseAttributesMap represents the map of attribute for App.
var AppLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"beta": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Beta",
		Description:    `Beta indicates if the app is in a beta version.`,
		Exposed:        true,
		Name:           "beta",
		ReadOnly:       true,
		Type:           "boolean",
	},
	"categoryid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CategoryID",
		Description:    `CategoryID of the app.`,
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
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Format:         "free",
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"icon": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Icon",
		Description:    `Icon contains a base64 image for the app.`,
		Exposed:        true,
		Format:         "free",
		Name:           "icon",
		ReadOnly:       true,
		Type:           "string",
	},
	"longdescription": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LongDescription",
		Description:    `LongDescription contains a more detailed description of the app.`,
		Exposed:        true,
		Format:         "free",
		Name:           "longDescription",
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		DefaultOrder:   true,
		Description:    `Name is the name of the entity.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"parameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Parameters",
		Description:    `Parameters of the app the user can or has to specify.`,
		Exposed:        true,
		Name:           "parameters",
		SubType:        "app_parameters",
		Type:           "external",
	},
	"title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		Description:    `Title represents the title of the app.`,
		Exposed:        true,
		Format:         "free",
		Name:           "title",
		Type:           "string",
	},
	"versionparameters": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "VersionParameters",
		Description:    `VersionParameters contains parameters for each available version.`,
		Exposed:        true,
		Name:           "versionParameters",
		SubType:        "app_versionparameters",
		Type:           "external",
	},
}
