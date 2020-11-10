package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TagIdentity represents the Identity of the object.
var TagIdentity = elemental.Identity{
	Name:     "tag",
	Category: "tags",
	Package:  "tagle",
	Private:  false,
}

// TagsList represents a list of Tags
type TagsList []*Tag

// Identity returns the identity of the objects in the list.
func (o TagsList) Identity() elemental.Identity {

	return TagIdentity
}

// Copy returns a pointer to a copy the TagsList.
func (o TagsList) Copy() elemental.Identifiables {

	copy := append(TagsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the TagsList.
func (o TagsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(TagsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Tag))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o TagsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TagsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TagsList converted to SparseTagsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TagsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseTagsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseTag)
	}

	return out
}

// Version returns the version of the content.
func (o TagsList) Version() int {

	return 1
}

// Tag represents the model of a tag
type Tag struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Represents the number of times the tag is used.
	Count int `json:"count" msgpack:"count" bson:"count" mapstructure:"count,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Represents the value of the tag.
	Value string `json:"value" msgpack:"value" bson:"value" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTag returns a new *Tag
func NewTag() *Tag {

	return &Tag{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Tag) Identity() elemental.Identity {

	return TagIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Tag) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Tag) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Tag) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTag{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Count = o.Count
	s.Namespace = o.Namespace
	s.Value = o.Value

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Tag) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesTag{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Count = s.Count
	o.Namespace = s.Namespace
	o.Value = s.Value

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Tag) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Tag) BleveType() string {

	return "tag"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Tag) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Tag) Doc() string {

	return `A tag is a key-value pair in string form that can applied to all objects in
the system. They are used for policy resolution. Tags starting with ` + "`" + `$` + "`" + ` are
derived from the property of an object. For example an object with an ID set to
` + "`" + `xxx` + "`" + ` and a name set to ` + "`" + `the name` + "`" + ` will be tagged by default with ` + "`" + `$name=the name` + "`" + `
and ` + "`" + `$id=xxx` + "`" + `. Tags starting with an ` + "`" + `@` + "`" + ` have been generated by an external
system.`
}

func (o *Tag) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetNamespace returns the Namespace of the receiver.
func (o *Tag) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Tag) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Tag) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTag{
			ID:        &o.ID,
			Count:     &o.Count,
			Namespace: &o.Namespace,
			Value:     &o.Value,
		}
	}

	sp := &SparseTag{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "count":
			sp.Count = &(o.Count)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "value":
			sp.Value = &(o.Value)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTag to the object.
func (o *Tag) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTag)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Count != nil {
		o.Count = *so.Count
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Value != nil {
		o.Value = *so.Value
	}
}

// DeepCopy returns a deep copy if the Tag.
func (o *Tag) DeepCopy() *Tag {

	if o == nil {
		return nil
	}

	out := &Tag{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Tag.
func (o *Tag) DeepCopyInto(out *Tag) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Tag: %s", err))
	}

	*out = *target.(*Tag)
}

// Validate valides the current information stored into the structure.
func (o *Tag) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("value", o.Value); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateTag("value", o.Value); err != nil {
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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Tag) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := TagAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TagLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Tag) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TagAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Tag) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "count":
		return o.Count
	case "namespace":
		return o.Namespace
	case "value":
		return o.Value
	}

	return nil
}

// TagAttributesMap represents the map of attribute for Tag.
var TagAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Count": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "count",
		ConvertedName:  "Count",
		Description:    `Represents the number of times the tag is used.`,
		Exposed:        true,
		Name:           "count",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
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
	"Value": {
		AllowedChoices: []string{},
		BSONFieldName:  "value",
		ConvertedName:  "Value",
		CreationOnly:   true,
		Description:    `Represents the value of the tag.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// TagLowerCaseAttributesMap represents the map of attribute for Tag.
var TagLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
		ConvertedName:  "ID",
		Description:    `Identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"count": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "count",
		ConvertedName:  "Count",
		Description:    `Represents the number of times the tag is used.`,
		Exposed:        true,
		Name:           "count",
		ReadOnly:       true,
		Stored:         true,
		Type:           "integer",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
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
	"value": {
		AllowedChoices: []string{},
		BSONFieldName:  "value",
		ConvertedName:  "Value",
		CreationOnly:   true,
		Description:    `Represents the value of the tag.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseTagsList represents a list of SparseTags
type SparseTagsList []*SparseTag

// Identity returns the identity of the objects in the list.
func (o SparseTagsList) Identity() elemental.Identity {

	return TagIdentity
}

// Copy returns a pointer to a copy the SparseTagsList.
func (o SparseTagsList) Copy() elemental.Identifiables {

	copy := append(SparseTagsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTagsList.
func (o SparseTagsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseTagsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTag))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseTagsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTagsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTagsList converted to TagsList.
func (o SparseTagsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTagsList) Version() int {

	return 1
}

// SparseTag represents the sparse version of a tag.
type SparseTag struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Represents the number of times the tag is used.
	Count *int `json:"count,omitempty" msgpack:"count,omitempty" bson:"count,omitempty" mapstructure:"count,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Represents the value of the tag.
	Value *string `json:"value,omitempty" msgpack:"value,omitempty" bson:"value,omitempty" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTag returns a new  SparseTag.
func NewSparseTag() *SparseTag {
	return &SparseTag{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTag) Identity() elemental.Identity {

	return TagIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTag) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTag) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTag) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTag{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Count != nil {
		s.Count = o.Count
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.Value != nil {
		s.Value = o.Value
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTag) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTag{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Count != nil {
		o.Count = s.Count
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.Value != nil {
		o.Value = s.Value
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTag) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTag) ToPlain() elemental.PlainIdentifiable {

	out := NewTag()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Value != nil {
		out.Value = *o.Value
	}

	return out
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseTag) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseTag) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// DeepCopy returns a deep copy if the SparseTag.
func (o *SparseTag) DeepCopy() *SparseTag {

	if o == nil {
		return nil
	}

	out := &SparseTag{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTag.
func (o *SparseTag) DeepCopyInto(out *SparseTag) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTag: %s", err))
	}

	*out = *target.(*SparseTag)
}

type mongoAttributesTag struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Count     int           `bson:"count"`
	Namespace string        `bson:"namespace"`
	Value     string        `bson:"value"`
}
type mongoAttributesSparseTag struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Count     *int          `bson:"count,omitempty"`
	Namespace *string       `bson:"namespace,omitempty"`
	Value     *string       `bson:"value,omitempty"`
}
