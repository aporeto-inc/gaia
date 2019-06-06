package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RecipeIdentity represents the Identity of the object.
var RecipeIdentity = elemental.Identity{
	Name:     "recipe",
	Category: "recipes",
	Package:  "ignis",
	Private:  false,
}

// RecipesList represents a list of Recipes
type RecipesList []*Recipe

// Identity returns the identity of the objects in the list.
func (o RecipesList) Identity() elemental.Identity {

	return RecipeIdentity
}

// Copy returns a pointer to a copy the RecipesList.
func (o RecipesList) Copy() elemental.Identifiables {

	copy := append(RecipesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the RecipesList.
func (o RecipesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(RecipesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Recipe))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o RecipesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o RecipesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToSparse returns the RecipesList converted to SparseRecipesList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o RecipesList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseRecipesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseRecipe)
	}

	return out
}

// Version returns the version of the content.
func (o RecipesList) Version() int {

	return 1
}

// Recipe represents the model of a recipe
type Recipe struct {
	// ID is the identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// Icon contains a base64 image for the recipe.
	Icon string `json:"icon" msgpack:"icon" bson:"icon" mapstructure:"icon,omitempty"`

	// Key is the unique key of the recipe.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	// Label defines the recipe.
	Label string `json:"label" msgpack:"label" bson:"label" mapstructure:"label,omitempty"`

	// LongDescription provides a long description of the recipe.
	LongDescription string `json:"longDescription" msgpack:"longDescription" bson:"longdescription" mapstructure:"longDescription,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata []string `json:"metadata" msgpack:"metadata" bson:"metadata" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Options of the recipe.
	Options *RecipeOptions `json:"options" msgpack:"options" bson:"options" mapstructure:"options,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// Steps contains all the steps with parameters to follow for the recipe.
	Steps []*UIStep `json:"steps" msgpack:"steps" bson:"steps" mapstructure:"steps,omitempty"`

	// successfullMessage is presented if present and success.
	SuccessfullMessage string `json:"successfullMessage" msgpack:"successfullMessage" bson:"successfullmessage" mapstructure:"successfullMessage,omitempty"`

	// Template of the recipe to import.
	Template string `json:"template" msgpack:"template" bson:"template" mapstructure:"template,omitempty"`

	// templateHash is a hash of the template.
	TemplateHash string `json:"templateHash" msgpack:"templateHash" bson:"templatehash" mapstructure:"templateHash,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRecipe returns a new *Recipe
func NewRecipe() *Recipe {

	return &Recipe{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		Options:        NewRecipeOptions(),
		NormalizedTags: []string{},
		Label:          "magicpanda",
		Metadata:       []string{},
		Steps:          []*UIStep{},
	}
}

// Identity returns the Identity of the object.
func (o *Recipe) Identity() elemental.Identity {

	return RecipeIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Recipe) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Recipe) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *Recipe) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *Recipe) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// Doc returns the documentation for the object
func (o *Recipe) Doc() string {

	return `A Recipe defines a list of steps to define a workflow.`
}

func (o *Recipe) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *Recipe) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *Recipe) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *Recipe) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *Recipe) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *Recipe) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *Recipe) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *Recipe) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *Recipe) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *Recipe) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *Recipe) SetDescription(description string) {

	o.Description = description
}

// GetMetadata returns the Metadata of the receiver.
func (o *Recipe) GetMetadata() []string {

	return o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the given value.
func (o *Recipe) SetMetadata(metadata []string) {

	o.Metadata = metadata
}

// GetName returns the Name of the receiver.
func (o *Recipe) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Recipe) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Recipe) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Recipe) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *Recipe) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *Recipe) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *Recipe) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *Recipe) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *Recipe) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *Recipe) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *Recipe) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *Recipe) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *Recipe) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *Recipe) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Recipe) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseRecipe{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Icon:                 &o.Icon,
			Key:                  &o.Key,
			Label:                &o.Label,
			LongDescription:      &o.LongDescription,
			Metadata:             &o.Metadata,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Options:              &o.Options,
			Propagate:            &o.Propagate,
			Protected:            &o.Protected,
			Steps:                &o.Steps,
			SuccessfullMessage:   &o.SuccessfullMessage,
			Template:             &o.Template,
			TemplateHash:         &o.TemplateHash,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
		}
	}

	sp := &SparseRecipe{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "annotations":
			sp.Annotations = &(o.Annotations)
		case "associatedTags":
			sp.AssociatedTags = &(o.AssociatedTags)
		case "createIdempotencyKey":
			sp.CreateIdempotencyKey = &(o.CreateIdempotencyKey)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "description":
			sp.Description = &(o.Description)
		case "icon":
			sp.Icon = &(o.Icon)
		case "key":
			sp.Key = &(o.Key)
		case "label":
			sp.Label = &(o.Label)
		case "longDescription":
			sp.LongDescription = &(o.LongDescription)
		case "metadata":
			sp.Metadata = &(o.Metadata)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "options":
			sp.Options = &(o.Options)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "protected":
			sp.Protected = &(o.Protected)
		case "steps":
			sp.Steps = &(o.Steps)
		case "successfullMessage":
			sp.SuccessfullMessage = &(o.SuccessfullMessage)
		case "template":
			sp.Template = &(o.Template)
		case "templateHash":
			sp.TemplateHash = &(o.TemplateHash)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseRecipe to the object.
func (o *Recipe) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseRecipe)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Annotations != nil {
		o.Annotations = *so.Annotations
	}
	if so.AssociatedTags != nil {
		o.AssociatedTags = *so.AssociatedTags
	}
	if so.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = *so.CreateIdempotencyKey
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Icon != nil {
		o.Icon = *so.Icon
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
	if so.Label != nil {
		o.Label = *so.Label
	}
	if so.LongDescription != nil {
		o.LongDescription = *so.LongDescription
	}
	if so.Metadata != nil {
		o.Metadata = *so.Metadata
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.NormalizedTags != nil {
		o.NormalizedTags = *so.NormalizedTags
	}
	if so.Options != nil {
		o.Options = *so.Options
	}
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.Steps != nil {
		o.Steps = *so.Steps
	}
	if so.SuccessfullMessage != nil {
		o.SuccessfullMessage = *so.SuccessfullMessage
	}
	if so.Template != nil {
		o.Template = *so.Template
	}
	if so.TemplateHash != nil {
		o.TemplateHash = *so.TemplateHash
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the Recipe.
func (o *Recipe) DeepCopy() *Recipe {

	if o == nil {
		return nil
	}

	out := &Recipe{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Recipe.
func (o *Recipe) DeepCopyInto(out *Recipe) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Recipe: %s", err))
	}

	*out = *target.(*Recipe)
}

// Validate valides the current information stored into the structure.
func (o *Recipe) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("label", o.Label); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateMetadata("metadata", o.Metadata); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
		errors = errors.Append(err)
	}

	if o.Options != nil {
		if err := o.Options.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.Steps {
		if sub == nil {
			continue
		}
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
func (*Recipe) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := RecipeAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RecipeLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Recipe) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return RecipeAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Recipe) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "annotations":
		return o.Annotations
	case "associatedTags":
		return o.AssociatedTags
	case "createIdempotencyKey":
		return o.CreateIdempotencyKey
	case "createTime":
		return o.CreateTime
	case "description":
		return o.Description
	case "icon":
		return o.Icon
	case "key":
		return o.Key
	case "label":
		return o.Label
	case "longDescription":
		return o.LongDescription
	case "metadata":
		return o.Metadata
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "options":
		return o.Options
	case "propagate":
		return o.Propagate
	case "protected":
		return o.Protected
	case "steps":
		return o.Steps
	case "successfullMessage":
		return o.SuccessfullMessage
	case "template":
		return o.Template
	case "templateHash":
		return o.TemplateHash
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	}

	return nil
}

// RecipeAttributesMap represents the map of attribute for Recipe.
var RecipeAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CreateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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
		Description:    `Icon contains a base64 image for the recipe.`,
		Exposed:        true,
		Name:           "icon",
		Stored:         true,
		Type:           "string",
	},
	"Key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Key",
		Description:    `Key is the unique key of the recipe.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Label": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Label",
		CreationOnly:   true,
		DefaultValue:   "magicpanda",
		Description:    `Label defines the recipe.`,
		Exposed:        true,
		Name:           "label",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"LongDescription": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LongDescription",
		Description:    `LongDescription provides a long description of the recipe.`,
		Exposed:        true,
		Name:           "longDescription",
		Stored:         true,
		Type:           "string",
	},
	"Metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
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
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"Options": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Options",
		Description:    `Options of the recipe.`,
		Exposed:        true,
		Name:           "options",
		Stored:         true,
		SubType:        "recipeoptions",
		Type:           "ref",
	},
	"Propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Steps": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Steps",
		Description:    `Steps contains all the steps with parameters to follow for the recipe.`,
		Exposed:        true,
		Name:           "steps",
		Stored:         true,
		SubType:        "uistep",
		Type:           "refList",
	},
	"SuccessfullMessage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SuccessfullMessage",
		Description:    `successfullMessage is presented if present and success.`,
		Exposed:        true,
		Name:           "successfullMessage",
		Stored:         true,
		Type:           "string",
	},
	"Template": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Template",
		Description:    `Template of the recipe to import.`,
		Exposed:        true,
		Name:           "template",
		Stored:         true,
		Type:           "string",
	},
	"TemplateHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TemplateHash",
		Description:    `templateHash is a hash of the template.`,
		Exposed:        true,
		Name:           "templateHash",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateIdempotencyKey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// RecipeLowerCaseAttributesMap represents the map of attribute for Recipe.
var RecipeLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"annotations": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Annotations",
		Description:    `Annotation stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "AssociatedTags",
		Description:    `AssociatedTags are the list of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"createidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
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
		Description:    `Icon contains a base64 image for the recipe.`,
		Exposed:        true,
		Name:           "icon",
		Stored:         true,
		Type:           "string",
	},
	"key": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Key",
		Description:    `Key is the unique key of the recipe.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"label": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Label",
		CreationOnly:   true,
		DefaultValue:   "magicpanda",
		Description:    `Label defines the recipe.`,
		Exposed:        true,
		Name:           "label",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"longdescription": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "LongDescription",
		Description:    `LongDescription provides a long description of the recipe.`,
		Exposed:        true,
		Name:           "longDescription",
		Stored:         true,
		Type:           "string",
	},
	"metadata": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Metadata",
		CreationOnly:   true,
		Description: `Metadata contains tags that can only be set during creation. They must all start
with the '@' prefix, and should only be used by external systems.`,
		Exposed:    true,
		Filterable: true,
		Getter:     true,
		Name:       "metadata",
		Setter:     true,
		Stored:     true,
		SubType:    "string",
		Type:       "list",
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
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "Namespace",
		DefaultOrder:   true,
		Description:    `Namespace tag attached to an entity.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"normalizedtags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "NormalizedTags",
		Description:    `NormalizedTags contains the list of normalized tags of the entities.`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Transient:      true,
		Type:           "list",
	},
	"options": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Options",
		Description:    `Options of the recipe.`,
		Exposed:        true,
		Name:           "options",
		Stored:         true,
		SubType:        "recipeoptions",
		Type:           "ref",
	},
	"propagate": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Propagate",
		Description:    `Propagate will propagate the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protected",
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"steps": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Steps",
		Description:    `Steps contains all the steps with parameters to follow for the recipe.`,
		Exposed:        true,
		Name:           "steps",
		Stored:         true,
		SubType:        "uistep",
		Type:           "refList",
	},
	"successfullmessage": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SuccessfullMessage",
		Description:    `successfullMessage is presented if present and success.`,
		Exposed:        true,
		Name:           "successfullMessage",
		Stored:         true,
		Type:           "string",
	},
	"template": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Template",
		Description:    `Template of the recipe to import.`,
		Exposed:        true,
		Name:           "template",
		Stored:         true,
		Type:           "string",
	},
	"templatehash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TemplateHash",
		Description:    `templateHash is a hash of the template.`,
		Exposed:        true,
		Name:           "templateHash",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"updateidempotencykey": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}

// SparseRecipesList represents a list of SparseRecipes
type SparseRecipesList []*SparseRecipe

// Identity returns the identity of the objects in the list.
func (o SparseRecipesList) Identity() elemental.Identity {

	return RecipeIdentity
}

// Copy returns a pointer to a copy the SparseRecipesList.
func (o SparseRecipesList) Copy() elemental.Identifiables {

	copy := append(SparseRecipesList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseRecipesList.
func (o SparseRecipesList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseRecipesList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseRecipe))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseRecipesList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseRecipesList) DefaultOrder() []string {

	return []string{
		"namespace",
		"name",
	}
}

// ToPlain returns the SparseRecipesList converted to RecipesList.
func (o SparseRecipesList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseRecipesList) Version() int {

	return 1
}

// SparseRecipe represents the sparse version of a recipe.
type SparseRecipe struct {
	// ID is the identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// Annotation stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// AssociatedTags are the list of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description is the description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// Icon contains a base64 image for the recipe.
	Icon *string `json:"icon,omitempty" msgpack:"icon,omitempty" bson:"icon,omitempty" mapstructure:"icon,omitempty"`

	// Key is the unique key of the recipe.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"key,omitempty" mapstructure:"key,omitempty"`

	// Label defines the recipe.
	Label *string `json:"label,omitempty" msgpack:"label,omitempty" bson:"label,omitempty" mapstructure:"label,omitempty"`

	// LongDescription provides a long description of the recipe.
	LongDescription *string `json:"longDescription,omitempty" msgpack:"longDescription,omitempty" bson:"longdescription,omitempty" mapstructure:"longDescription,omitempty"`

	// Metadata contains tags that can only be set during creation. They must all start
	// with the '@' prefix, and should only be used by external systems.
	Metadata *[]string `json:"metadata,omitempty" msgpack:"metadata,omitempty" bson:"metadata,omitempty" mapstructure:"metadata,omitempty"`

	// Name is the name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// NormalizedTags contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Options of the recipe.
	Options **RecipeOptions `json:"options,omitempty" msgpack:"options,omitempty" bson:"options,omitempty" mapstructure:"options,omitempty"`

	// Propagate will propagate the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Protected defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// Steps contains all the steps with parameters to follow for the recipe.
	Steps *[]*UIStep `json:"steps,omitempty" msgpack:"steps,omitempty" bson:"steps,omitempty" mapstructure:"steps,omitempty"`

	// successfullMessage is presented if present and success.
	SuccessfullMessage *string `json:"successfullMessage,omitempty" msgpack:"successfullMessage,omitempty" bson:"successfullmessage,omitempty" mapstructure:"successfullMessage,omitempty"`

	// Template of the recipe to import.
	Template *string `json:"template,omitempty" msgpack:"template,omitempty" bson:"template,omitempty" mapstructure:"template,omitempty"`

	// templateHash is a hash of the template.
	TemplateHash *string `json:"templateHash,omitempty" msgpack:"templateHash,omitempty" bson:"templatehash,omitempty" mapstructure:"templateHash,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseRecipe returns a new  SparseRecipe.
func NewSparseRecipe() *SparseRecipe {
	return &SparseRecipe{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseRecipe) Identity() elemental.Identity {

	return RecipeIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseRecipe) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseRecipe) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseRecipe) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseRecipe) ToPlain() elemental.PlainIdentifiable {

	out := NewRecipe()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Annotations != nil {
		out.Annotations = *o.Annotations
	}
	if o.AssociatedTags != nil {
		out.AssociatedTags = *o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		out.CreateIdempotencyKey = *o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Icon != nil {
		out.Icon = *o.Icon
	}
	if o.Key != nil {
		out.Key = *o.Key
	}
	if o.Label != nil {
		out.Label = *o.Label
	}
	if o.LongDescription != nil {
		out.LongDescription = *o.LongDescription
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.NormalizedTags != nil {
		out.NormalizedTags = *o.NormalizedTags
	}
	if o.Options != nil {
		out.Options = *o.Options
	}
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.Steps != nil {
		out.Steps = *o.Steps
	}
	if o.SuccessfullMessage != nil {
		out.SuccessfullMessage = *o.SuccessfullMessage
	}
	if o.Template != nil {
		out.Template = *o.Template
	}
	if o.TemplateHash != nil {
		out.TemplateHash = *o.TemplateHash
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparseRecipe) GetAnnotations() map[string][]string {

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparseRecipe) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparseRecipe) GetAssociatedTags() []string {

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparseRecipe) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparseRecipe) GetCreateIdempotencyKey() string {

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseRecipe) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseRecipe) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseRecipe) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparseRecipe) GetDescription() string {

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparseRecipe) SetDescription(description string) {

	o.Description = &description
}

// GetMetadata returns the Metadata of the receiver.
func (o *SparseRecipe) GetMetadata() []string {

	return *o.Metadata
}

// SetMetadata sets the property Metadata of the receiver using the address of the given value.
func (o *SparseRecipe) SetMetadata(metadata []string) {

	o.Metadata = &metadata
}

// GetName returns the Name of the receiver.
func (o *SparseRecipe) GetName() string {

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseRecipe) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseRecipe) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseRecipe) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparseRecipe) GetNormalizedTags() []string {

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparseRecipe) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparseRecipe) GetPropagate() bool {

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparseRecipe) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparseRecipe) GetProtected() bool {

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparseRecipe) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparseRecipe) GetUpdateIdempotencyKey() string {

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparseRecipe) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseRecipe) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseRecipe) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseRecipe.
func (o *SparseRecipe) DeepCopy() *SparseRecipe {

	if o == nil {
		return nil
	}

	out := &SparseRecipe{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseRecipe.
func (o *SparseRecipe) DeepCopyInto(out *SparseRecipe) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseRecipe: %s", err))
	}

	*out = *target.(*SparseRecipe)
}
