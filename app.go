package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia/types"
)

// AppIdentity represents the Identity of the object.
var AppIdentity = elemental.Identity{
	Name:     "app",
	Category: "apps",
	Package:  "highwind",
	Private:  false,
}

// AppsList represents a list of Apps
type AppsList []*App

// Identity returns the identity of the objects in the list.
func (o AppsList) Identity() elemental.Identity {

	return AppIdentity
}

// Copy returns a pointer to a copy the AppsList.
func (o AppsList) Copy() elemental.Identifiables {

	copy := append(AppsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AppsList.
func (o AppsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AppsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*App))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AppsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AppsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the AppsList converted to SparseAppsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AppsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
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

	// Title represents the title of the app.
	Title string `json:"title" bson:"-" mapstructure:"title,omitempty"`

	// VersionParameters contains parameters for each available version.
	VersionParameters map[string][]*types.AppParameter `json:"versionParameters" bson:"-" mapstructure:"versionParameters,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewApp returns a new *App
func NewApp() *App {

	return &App{
		ModelVersion:      1,
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

// SetName sets the property Name of the receiver using the given value.
func (o *App) SetName(name string) {

	o.Name = name
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *App) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseApp{
			Beta:              &o.Beta,
			CategoryID:        &o.CategoryID,
			Description:       &o.Description,
			Icon:              &o.Icon,
			LongDescription:   &o.LongDescription,
			Name:              &o.Name,
			Title:             &o.Title,
			VersionParameters: &o.VersionParameters,
		}
	}

	sp := &SparseApp{}
	for _, f := range fields {
		switch f {
		case "beta":
			sp.Beta = &(o.Beta)
		case "categoryID":
			sp.CategoryID = &(o.CategoryID)
		case "description":
			sp.Description = &(o.Description)
		case "icon":
			sp.Icon = &(o.Icon)
		case "longDescription":
			sp.LongDescription = &(o.LongDescription)
		case "name":
			sp.Name = &(o.Name)
		case "title":
			sp.Title = &(o.Title)
		case "versionParameters":
			sp.VersionParameters = &(o.VersionParameters)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseApp to the object.
func (o *App) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseApp)
	if so.Beta != nil {
		o.Beta = *so.Beta
	}
	if so.CategoryID != nil {
		o.CategoryID = *so.CategoryID
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Icon != nil {
		o.Icon = *so.Icon
	}
	if so.LongDescription != nil {
		o.LongDescription = *so.LongDescription
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Title != nil {
		o.Title = *so.Title
	}
	if so.VersionParameters != nil {
		o.VersionParameters = *so.VersionParameters
	}
}

// DeepCopy returns a deep copy if the App.
func (o *App) DeepCopy() *App {

	if o == nil {
		return nil
	}

	out := &App{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *App.
func (o *App) DeepCopyInto(out *App) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy App: %s", err))
	}

	*out = *target.(*App)
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

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *App) ValueForAttribute(name string) interface{} {

	switch name {
	case "beta":
		return o.Beta
	case "categoryID":
		return o.CategoryID
	case "description":
		return o.Description
	case "icon":
		return o.Icon
	case "longDescription":
		return o.LongDescription
	case "name":
		return o.Name
	case "title":
		return o.Title
	case "versionParameters":
		return o.VersionParameters
	}

	return nil
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
		Name:           "categoryID",
		ReadOnly:       true,
		Type:           "string",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
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
		Name:           "icon",
		ReadOnly:       true,
		Type:           "string",
	},
	"LongDescription": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LongDescription",
		Description:    `LongDescription contains a more detailed description of the app.`,
		Exposed:        true,
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
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		Description:    `Title represents the title of the app.`,
		Exposed:        true,
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
		Name:           "categoryID",
		ReadOnly:       true,
		Type:           "string",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `Description is the description of the object.`,
		Exposed:        true,
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
		Name:           "icon",
		ReadOnly:       true,
		Type:           "string",
	},
	"longdescription": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LongDescription",
		Description:    `LongDescription contains a more detailed description of the app.`,
		Exposed:        true,
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
		Getter:         true,
		MaxLength:      256,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		Description:    `Title represents the title of the app.`,
		Exposed:        true,
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

// SparseAppsList represents a list of SparseApps
type SparseAppsList []*SparseApp

// Identity returns the identity of the objects in the list.
func (o SparseAppsList) Identity() elemental.Identity {

	return AppIdentity
}

// Copy returns a pointer to a copy the SparseAppsList.
func (o SparseAppsList) Copy() elemental.Identifiables {

	copy := append(SparseAppsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAppsList.
func (o SparseAppsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAppsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseApp))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAppsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAppsList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparseAppsList converted to AppsList.
func (o SparseAppsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAppsList) Version() int {

	return 1
}

// SparseApp represents the sparse version of a app.
type SparseApp struct {
	// Beta indicates if the app is in a beta version.
	Beta *bool `json:"beta,omitempty" bson:"-" mapstructure:"beta,omitempty"`

	// CategoryID of the app.
	CategoryID *string `json:"categoryID,omitempty" bson:"-" mapstructure:"categoryID,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" bson:"description" mapstructure:"description,omitempty"`

	// Icon contains a base64 image for the app.
	Icon *string `json:"icon,omitempty" bson:"-" mapstructure:"icon,omitempty"`

	// LongDescription contains a more detailed description of the app.
	LongDescription *string `json:"longDescription,omitempty" bson:"-" mapstructure:"longDescription,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" bson:"name" mapstructure:"name,omitempty"`

	// Title represents the title of the app.
	Title *string `json:"title,omitempty" bson:"-" mapstructure:"title,omitempty"`

	// VersionParameters contains parameters for each available version.
	VersionParameters *map[string][]*types.AppParameter `json:"versionParameters,omitempty" bson:"-" mapstructure:"versionParameters,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseApp returns a new  SparseApp.
func NewSparseApp() *SparseApp {
	return &SparseApp{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseApp) Identity() elemental.Identity {

	return AppIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseApp) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseApp) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseApp) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseApp) ToPlain() elemental.PlainIdentifiable {

	out := NewApp()
	if o.Beta != nil {
		out.Beta = *o.Beta
	}
	if o.CategoryID != nil {
		out.CategoryID = *o.CategoryID
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Icon != nil {
		out.Icon = *o.Icon
	}
	if o.LongDescription != nil {
		out.LongDescription = *o.LongDescription
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Title != nil {
		out.Title = *o.Title
	}
	if o.VersionParameters != nil {
		out.VersionParameters = *o.VersionParameters
	}

	return out
}

// GetName returns the Name of the receiver.
func (o *SparseApp) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseApp) SetName(name string) {

	o.Name = &name
}

// DeepCopy returns a deep copy if the SparseApp.
func (o *SparseApp) DeepCopy() *SparseApp {

	if o == nil {
		return nil
	}

	out := &SparseApp{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseApp.
func (o *SparseApp) DeepCopyInto(out *SparseApp) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseApp: %s", err))
	}

	*out = *target.(*SparseApp)
}
