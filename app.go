package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
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
func (o AppsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAppsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseApp)
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
	Beta bool `json:"beta" msgpack:"beta" bson:"-" mapstructure:"beta,omitempty"`

	// CategoryID of the app.
	CategoryID string `json:"categoryID" msgpack:"categoryID" bson:"-" mapstructure:"categoryID,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Icon contains a base64 image for the app.
	Icon string `json:"icon" msgpack:"icon" bson:"-" mapstructure:"icon,omitempty"`

	// LatestVersion represents the latest version available of the app.
	LatestVersion string `json:"latestVersion" msgpack:"latestVersion" bson:"-" mapstructure:"latestVersion,omitempty"`

	// LongDescription contains a more detailed description of the app.
	LongDescription string `json:"longDescription" msgpack:"longDescription" bson:"-" mapstructure:"longDescription,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// List of steps that contains parameters.
	Steps []*UIStep `json:"steps" msgpack:"steps" bson:"steps" mapstructure:"steps,omitempty"`

	// Title represents the title of the app.
	Title string `json:"title" msgpack:"title" bson:"-" mapstructure:"title,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewApp returns a new *App
func NewApp() *App {

	return &App{
		ModelVersion: 1,
		Steps:        []*UIStep{},
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

// GetDescription returns the Description of the receiver.
func (o *App) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *App) SetDescription(description string) {

	o.Description = description
}

// GetName returns the Name of the receiver.
func (o *App) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *App) SetName(name string) {

	o.Name = name
}

// GetZHash returns the ZHash of the receiver.
func (o *App) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *App) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *App) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *App) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *App) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseApp{
			Beta:            &o.Beta,
			CategoryID:      &o.CategoryID,
			Description:     &o.Description,
			Icon:            &o.Icon,
			LatestVersion:   &o.LatestVersion,
			LongDescription: &o.LongDescription,
			Name:            &o.Name,
			Steps:           &o.Steps,
			Title:           &o.Title,
			ZHash:           &o.ZHash,
			Zone:            &o.Zone,
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
		case "latestVersion":
			sp.LatestVersion = &(o.LatestVersion)
		case "longDescription":
			sp.LongDescription = &(o.LongDescription)
		case "name":
			sp.Name = &(o.Name)
		case "steps":
			sp.Steps = &(o.Steps)
		case "title":
			sp.Title = &(o.Title)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
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
	if so.LatestVersion != nil {
		o.LatestVersion = *so.LatestVersion
	}
	if so.LongDescription != nil {
		o.LongDescription = *so.LongDescription
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Steps != nil {
		o.Steps = *so.Steps
	}
	if so.Title != nil {
		o.Title = *so.Title
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
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
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	for _, sub := range o.Steps {
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
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
	case "latestVersion":
		return o.LatestVersion
	case "longDescription":
		return o.LongDescription
	case "name":
		return o.Name
	case "steps":
		return o.Steps
	case "title":
		return o.Title
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
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
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
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
	"LatestVersion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LatestVersion",
		Description:    `LatestVersion represents the latest version available of the app.`,
		Exposed:        true,
		Name:           "latestVersion",
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
	"Steps": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Steps",
		Description:    `List of steps that contains parameters.`,
		Exposed:        true,
		Name:           "steps",
		Stored:         true,
		SubType:        "uistep",
		Type:           "refList",
	},
	"Title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		Description:    `Title represents the title of the app.`,
		Exposed:        true,
		Name:           "title",
		Type:           "string",
	},
	"ZHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"Zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
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
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
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
	"latestversion": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LatestVersion",
		Description:    `LatestVersion represents the latest version available of the app.`,
		Exposed:        true,
		Name:           "latestVersion",
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
	"steps": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Steps",
		Description:    `List of steps that contains parameters.`,
		Exposed:        true,
		Name:           "steps",
		Stored:         true,
		SubType:        "uistep",
		Type:           "refList",
	},
	"title": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Title",
		Description:    `Title represents the title of the app.`,
		Exposed:        true,
		Name:           "title",
		Type:           "string",
	},
	"zhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ZHash",
		Description: `geographical hash of the data. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zHash",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
	},
	"zone": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Zone",
		Description: `geographical zone. This is used for sharding and
georedundancy.`,
		Getter:   true,
		Name:     "zone",
		ReadOnly: true,
		Setter:   true,
		Stored:   true,
		Type:     "integer",
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
	Beta *bool `json:"beta,omitempty" msgpack:"beta,omitempty" bson:"-" mapstructure:"beta,omitempty"`

	// CategoryID of the app.
	CategoryID *string `json:"categoryID,omitempty" msgpack:"categoryID,omitempty" bson:"-" mapstructure:"categoryID,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Icon contains a base64 image for the app.
	Icon *string `json:"icon,omitempty" msgpack:"icon,omitempty" bson:"-" mapstructure:"icon,omitempty"`

	// LatestVersion represents the latest version available of the app.
	LatestVersion *string `json:"latestVersion,omitempty" msgpack:"latestVersion,omitempty" bson:"-" mapstructure:"latestVersion,omitempty"`

	// LongDescription contains a more detailed description of the app.
	LongDescription *string `json:"longDescription,omitempty" msgpack:"longDescription,omitempty" bson:"-" mapstructure:"longDescription,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// List of steps that contains parameters.
	Steps *[]*UIStep `json:"steps,omitempty" msgpack:"steps,omitempty" bson:"steps,omitempty" mapstructure:"steps,omitempty"`

	// Title represents the title of the app.
	Title *string `json:"title,omitempty" msgpack:"title,omitempty" bson:"-" mapstructure:"title,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// geographical zone. This is used for sharding and
	// georedundancy.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
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
	if o.LatestVersion != nil {
		out.LatestVersion = *o.LatestVersion
	}
	if o.LongDescription != nil {
		out.LongDescription = *o.LongDescription
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Steps != nil {
		out.Steps = *o.Steps
	}
	if o.Title != nil {
		out.Title = *o.Title
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetDescription returns the Description of the receiver.
func (o *SparseApp) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseApp) SetDescription(description string) {

	o.Description = &description
}

// GetName returns the Name of the receiver.
func (o *SparseApp) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseApp) SetName(name string) {

	o.Name = &name
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseApp) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseApp) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseApp) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseApp) SetZone(zone int) {

	o.Zone = &zone
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
