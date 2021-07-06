package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PrivateKeyIdentity represents the Identity of the object.
var PrivateKeyIdentity = elemental.Identity{
	Name:     "privatekey",
	Category: "privatekey",
	Package:  "squall",
	Private:  false,
}

// PrivateKeysList represents a list of PrivateKeys
type PrivateKeysList []*PrivateKey

// Identity returns the identity of the objects in the list.
func (o PrivateKeysList) Identity() elemental.Identity {

	return PrivateKeyIdentity
}

// Copy returns a pointer to a copy the PrivateKeysList.
func (o PrivateKeysList) Copy() elemental.Identifiables {

	copy := append(PrivateKeysList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PrivateKeysList.
func (o PrivateKeysList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PrivateKeysList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PrivateKey))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PrivateKeysList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PrivateKeysList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToSparse returns the PrivateKeysList converted to SparsePrivateKeysList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PrivateKeysList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePrivateKeysList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePrivateKey)
	}

	return out
}

// Version returns the version of the content.
func (o PrivateKeysList) Version() int {

	return 1
}

// PrivateKey represents the model of a privatekey
type PrivateKey struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations map[string][]string `json:"annotations" msgpack:"annotations" bson:"annotations" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags []string `json:"associatedTags" msgpack:"associatedTags" bson:"associatedtags" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey string `json:"-" msgpack:"-" bson:"createidempotencykey" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// The fingerprint of the private key.
	Fingerprint string `json:"fingerprint" msgpack:"fingerprint" bson:"fingerprint" mapstructure:"fingerprint,omitempty"`

	// Private key in PEM format.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	// Name of the entity.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags []string `json:"normalizedTags" msgpack:"normalizedTags" bson:"normalizedtags" mapstructure:"normalizedTags,omitempty"`

	// Propagates the policy to all of its children.
	Propagate bool `json:"propagate" msgpack:"propagate" bson:"propagate" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected bool `json:"protected" msgpack:"protected" bson:"protected" mapstructure:"protected,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey string `json:"-" msgpack:"-" bson:"updateidempotencykey" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPrivateKey returns a new *PrivateKey
func NewPrivateKey() *PrivateKey {

	return &PrivateKey{
		ModelVersion:   1,
		Annotations:    map[string][]string{},
		AssociatedTags: []string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *PrivateKey) Identity() elemental.Identity {

	return PrivateKeyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PrivateKey) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PrivateKey) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PrivateKey) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPrivateKey{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Annotations = o.Annotations
	s.AssociatedTags = o.AssociatedTags
	s.CreateIdempotencyKey = o.CreateIdempotencyKey
	s.CreateTime = o.CreateTime
	s.Description = o.Description
	s.Fingerprint = o.Fingerprint
	s.Key = o.Key
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.NormalizedTags = o.NormalizedTags
	s.Propagate = o.Propagate
	s.Protected = o.Protected
	s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PrivateKey) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPrivateKey{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Annotations = s.Annotations
	o.AssociatedTags = s.AssociatedTags
	o.CreateIdempotencyKey = s.CreateIdempotencyKey
	o.CreateTime = s.CreateTime
	o.Description = s.Description
	o.Fingerprint = s.Fingerprint
	o.Key = s.Key
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.NormalizedTags = s.NormalizedTags
	o.Propagate = s.Propagate
	o.Protected = s.Protected
	o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PrivateKey) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PrivateKey) BleveType() string {

	return "privatekey"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PrivateKey) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// Doc returns the documentation for the object
func (o *PrivateKey) Doc() string {

	return `Manages private keys.`
}

func (o *PrivateKey) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAnnotations returns the Annotations of the receiver.
func (o *PrivateKey) GetAnnotations() map[string][]string {

	return o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the given value.
func (o *PrivateKey) SetAnnotations(annotations map[string][]string) {

	o.Annotations = annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *PrivateKey) GetAssociatedTags() []string {

	return o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the given value.
func (o *PrivateKey) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *PrivateKey) GetCreateIdempotencyKey() string {

	return o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the given value.
func (o *PrivateKey) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *PrivateKey) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *PrivateKey) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetDescription returns the Description of the receiver.
func (o *PrivateKey) GetDescription() string {

	return o.Description
}

// SetDescription sets the property Description of the receiver using the given value.
func (o *PrivateKey) SetDescription(description string) {

	o.Description = description
}

// GetName returns the Name of the receiver.
func (o *PrivateKey) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *PrivateKey) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *PrivateKey) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *PrivateKey) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *PrivateKey) GetNormalizedTags() []string {

	return o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the given value.
func (o *PrivateKey) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *PrivateKey) GetPropagate() bool {

	return o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the given value.
func (o *PrivateKey) SetPropagate(propagate bool) {

	o.Propagate = propagate
}

// GetProtected returns the Protected of the receiver.
func (o *PrivateKey) GetProtected() bool {

	return o.Protected
}

// SetProtected sets the property Protected of the receiver using the given value.
func (o *PrivateKey) SetProtected(protected bool) {

	o.Protected = protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *PrivateKey) GetUpdateIdempotencyKey() string {

	return o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the given value.
func (o *PrivateKey) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *PrivateKey) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *PrivateKey) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *PrivateKey) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *PrivateKey) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *PrivateKey) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *PrivateKey) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PrivateKey) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePrivateKey{
			ID:                   &o.ID,
			Annotations:          &o.Annotations,
			AssociatedTags:       &o.AssociatedTags,
			CreateIdempotencyKey: &o.CreateIdempotencyKey,
			CreateTime:           &o.CreateTime,
			Description:          &o.Description,
			Fingerprint:          &o.Fingerprint,
			Key:                  &o.Key,
			Name:                 &o.Name,
			Namespace:            &o.Namespace,
			NormalizedTags:       &o.NormalizedTags,
			Propagate:            &o.Propagate,
			Protected:            &o.Protected,
			UpdateIdempotencyKey: &o.UpdateIdempotencyKey,
			UpdateTime:           &o.UpdateTime,
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparsePrivateKey{}
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
		case "fingerprint":
			sp.Fingerprint = &(o.Fingerprint)
		case "key":
			sp.Key = &(o.Key)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "normalizedTags":
			sp.NormalizedTags = &(o.NormalizedTags)
		case "propagate":
			sp.Propagate = &(o.Propagate)
		case "protected":
			sp.Protected = &(o.Protected)
		case "updateIdempotencyKey":
			sp.UpdateIdempotencyKey = &(o.UpdateIdempotencyKey)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *PrivateKey) EncryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if o.Key, err = encrypter.EncryptString(o.Key); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'Key' for 'PrivateKey' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *PrivateKey) DecryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if o.Key, err = encrypter.DecryptString(o.Key); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'Key' for 'PrivateKey' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// Patch apply the non nil value of a *SparsePrivateKey to the object.
func (o *PrivateKey) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePrivateKey)
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
	if so.Fingerprint != nil {
		o.Fingerprint = *so.Fingerprint
	}
	if so.Key != nil {
		o.Key = *so.Key
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
	if so.Propagate != nil {
		o.Propagate = *so.Propagate
	}
	if so.Protected != nil {
		o.Protected = *so.Protected
	}
	if so.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = *so.UpdateIdempotencyKey
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// DeepCopy returns a deep copy if the PrivateKey.
func (o *PrivateKey) DeepCopy() *PrivateKey {

	if o == nil {
		return nil
	}

	out := &PrivateKey{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PrivateKey.
func (o *PrivateKey) DeepCopyInto(out *PrivateKey) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PrivateKey: %s", err))
	}

	*out = *target.(*PrivateKey)
}

// Validate valides the current information stored into the structure.
func (o *PrivateKey) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidateTagsWithoutReservedPrefixes("associatedTags", o.AssociatedTags); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("description", o.Description, 1024, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumLength("name", o.Name, 256, false); err != nil {
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
func (*PrivateKey) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PrivateKeyAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PrivateKeyLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PrivateKey) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PrivateKeyAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PrivateKey) ValueForAttribute(name string) interface{} {

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
	case "fingerprint":
		return o.Fingerprint
	case "key":
		return o.Key
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "normalizedTags":
		return o.NormalizedTags
	case "propagate":
		return o.Propagate
	case "protected":
		return o.Protected
	case "updateIdempotencyKey":
		return o.UpdateIdempotencyKey
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// PrivateKeyAttributesMap represents the map of attribute for PrivateKey.
var PrivateKeyAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"AssociatedTags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"CreateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
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
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"Fingerprint": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "fingerprint",
		ConvertedName:  "Fingerprint",
		Description:    `The fingerprint of the private key.`,
		Exposed:        true,
		Name:           "fingerprint",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		CreationOnly:   true,
		Description:    `Private key in PEM format.`,
		Encrypted:      true,
		Exposed:        true,
		Name:           "key",
		Required:       true,
		Secret:         true,
		Stored:         true,
		Transient:      true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
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
	"NormalizedTags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
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
	"Propagate": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagate",
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"Protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"UpdateIdempotencyKey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"UpdateTime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
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
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
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
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// PrivateKeyLowerCaseAttributesMap represents the map of attribute for PrivateKey.
var PrivateKeyLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"annotations": {
		AllowedChoices: []string{},
		BSONFieldName:  "annotations",
		ConvertedName:  "Annotations",
		Description:    `Stores additional information about an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "annotations",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string][]string",
		Type:           "external",
	},
	"associatedtags": {
		AllowedChoices: []string{},
		BSONFieldName:  "associatedtags",
		ConvertedName:  "AssociatedTags",
		Description:    `List of tags attached to an entity.`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"createidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createidempotencykey",
		ConvertedName:  "CreateIdempotencyKey",
		Description:    `internal idempotency key for a create operation.`,
		Getter:         true,
		Name:           "createIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "createtime",
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
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `Description of the object.`,
		Exposed:        true,
		Getter:         true,
		MaxLength:      1024,
		Name:           "description",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"fingerprint": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "fingerprint",
		ConvertedName:  "Fingerprint",
		Description:    `The fingerprint of the private key.`,
		Exposed:        true,
		Name:           "fingerprint",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		CreationOnly:   true,
		Description:    `Private key in PEM format.`,
		Encrypted:      true,
		Exposed:        true,
		Name:           "key",
		Required:       true,
		Secret:         true,
		Stored:         true,
		Transient:      true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `Name of the entity.`,
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
	"normalizedtags": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "normalizedtags",
		ConvertedName:  "NormalizedTags",
		Description:    `Contains the list of normalized tags of the entities.`,
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
	"propagate": {
		AllowedChoices: []string{},
		BSONFieldName:  "propagate",
		ConvertedName:  "Propagate",
		Description:    `Propagates the policy to all of its children.`,
		Exposed:        true,
		Getter:         true,
		Name:           "propagate",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"protected": {
		AllowedChoices: []string{},
		BSONFieldName:  "protected",
		ConvertedName:  "Protected",
		Description:    `Defines if the object is protected.`,
		Exposed:        true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"updateidempotencykey": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updateidempotencykey",
		ConvertedName:  "UpdateIdempotencyKey",
		Description:    `internal idempotency key for a update operation.`,
		Getter:         true,
		Name:           "updateIdempotencyKey",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"updatetime": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "updatetime",
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
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
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
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Logical storage zone. Used for sharding.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparsePrivateKeysList represents a list of SparsePrivateKeys
type SparsePrivateKeysList []*SparsePrivateKey

// Identity returns the identity of the objects in the list.
func (o SparsePrivateKeysList) Identity() elemental.Identity {

	return PrivateKeyIdentity
}

// Copy returns a pointer to a copy the SparsePrivateKeysList.
func (o SparsePrivateKeysList) Copy() elemental.Identifiables {

	copy := append(SparsePrivateKeysList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePrivateKeysList.
func (o SparsePrivateKeysList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePrivateKeysList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePrivateKey))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePrivateKeysList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePrivateKeysList) DefaultOrder() []string {

	return []string{
		"name",
	}
}

// ToPlain returns the SparsePrivateKeysList converted to PrivateKeysList.
func (o SparsePrivateKeysList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePrivateKeysList) Version() int {

	return 1
}

// SparsePrivateKey represents the sparse version of a privatekey.
type SparsePrivateKey struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Stores additional information about an entity.
	Annotations *map[string][]string `json:"annotations,omitempty" msgpack:"annotations,omitempty" bson:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// List of tags attached to an entity.
	AssociatedTags *[]string `json:"associatedTags,omitempty" msgpack:"associatedTags,omitempty" bson:"associatedtags,omitempty" mapstructure:"associatedTags,omitempty"`

	// internal idempotency key for a create operation.
	CreateIdempotencyKey *string `json:"-" msgpack:"-" bson:"createidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Description of the object.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// The fingerprint of the private key.
	Fingerprint *string `json:"fingerprint,omitempty" msgpack:"fingerprint,omitempty" bson:"fingerprint,omitempty" mapstructure:"fingerprint,omitempty"`

	// Private key in PEM format.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"key,omitempty" mapstructure:"key,omitempty"`

	// Name of the entity.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Contains the list of normalized tags of the entities.
	NormalizedTags *[]string `json:"normalizedTags,omitempty" msgpack:"normalizedTags,omitempty" bson:"normalizedtags,omitempty" mapstructure:"normalizedTags,omitempty"`

	// Propagates the policy to all of its children.
	Propagate *bool `json:"propagate,omitempty" msgpack:"propagate,omitempty" bson:"propagate,omitempty" mapstructure:"propagate,omitempty"`

	// Defines if the object is protected.
	Protected *bool `json:"protected,omitempty" msgpack:"protected,omitempty" bson:"protected,omitempty" mapstructure:"protected,omitempty"`

	// internal idempotency key for a update operation.
	UpdateIdempotencyKey *string `json:"-" msgpack:"-" bson:"updateidempotencykey,omitempty" mapstructure:"-,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePrivateKey returns a new  SparsePrivateKey.
func NewSparsePrivateKey() *SparsePrivateKey {
	return &SparsePrivateKey{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePrivateKey) Identity() elemental.Identity {

	return PrivateKeyIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePrivateKey) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePrivateKey) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePrivateKey) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePrivateKey{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Annotations != nil {
		s.Annotations = o.Annotations
	}
	if o.AssociatedTags != nil {
		s.AssociatedTags = o.AssociatedTags
	}
	if o.CreateIdempotencyKey != nil {
		s.CreateIdempotencyKey = o.CreateIdempotencyKey
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Fingerprint != nil {
		s.Fingerprint = o.Fingerprint
	}
	if o.Key != nil {
		s.Key = o.Key
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.NormalizedTags != nil {
		s.NormalizedTags = o.NormalizedTags
	}
	if o.Propagate != nil {
		s.Propagate = o.Propagate
	}
	if o.Protected != nil {
		s.Protected = o.Protected
	}
	if o.UpdateIdempotencyKey != nil {
		s.UpdateIdempotencyKey = o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePrivateKey) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePrivateKey{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Annotations != nil {
		o.Annotations = s.Annotations
	}
	if s.AssociatedTags != nil {
		o.AssociatedTags = s.AssociatedTags
	}
	if s.CreateIdempotencyKey != nil {
		o.CreateIdempotencyKey = s.CreateIdempotencyKey
	}
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Fingerprint != nil {
		o.Fingerprint = s.Fingerprint
	}
	if s.Key != nil {
		o.Key = s.Key
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.NormalizedTags != nil {
		o.NormalizedTags = s.NormalizedTags
	}
	if s.Propagate != nil {
		o.Propagate = s.Propagate
	}
	if s.Protected != nil {
		o.Protected = s.Protected
	}
	if s.UpdateIdempotencyKey != nil {
		o.UpdateIdempotencyKey = s.UpdateIdempotencyKey
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePrivateKey) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePrivateKey) ToPlain() elemental.PlainIdentifiable {

	out := NewPrivateKey()
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
	if o.Fingerprint != nil {
		out.Fingerprint = *o.Fingerprint
	}
	if o.Key != nil {
		out.Key = *o.Key
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
	if o.Propagate != nil {
		out.Propagate = *o.Propagate
	}
	if o.Protected != nil {
		out.Protected = *o.Protected
	}
	if o.UpdateIdempotencyKey != nil {
		out.UpdateIdempotencyKey = *o.UpdateIdempotencyKey
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *SparsePrivateKey) EncryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if *o.Key, err = encrypter.EncryptString(*o.Key); err != nil {
		return fmt.Errorf("unable to encrypt attribute 'Key' for 'SparsePrivateKey' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *SparsePrivateKey) DecryptAttributes(encrypter elemental.AttributeEncrypter) (err error) {

	if *o.Key, err = encrypter.DecryptString(*o.Key); err != nil {
		return fmt.Errorf("unable to decrypt attribute 'Key' for 'SparsePrivateKey' (%s): %s", o.Identifier(), err)
	}

	return nil
}

// GetAnnotations returns the Annotations of the receiver.
func (o *SparsePrivateKey) GetAnnotations() (out map[string][]string) {

	if o.Annotations == nil {
		return
	}

	return *o.Annotations
}

// SetAnnotations sets the property Annotations of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetAnnotations(annotations map[string][]string) {

	o.Annotations = &annotations
}

// GetAssociatedTags returns the AssociatedTags of the receiver.
func (o *SparsePrivateKey) GetAssociatedTags() (out []string) {

	if o.AssociatedTags == nil {
		return
	}

	return *o.AssociatedTags
}

// SetAssociatedTags sets the property AssociatedTags of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetAssociatedTags(associatedTags []string) {

	o.AssociatedTags = &associatedTags
}

// GetCreateIdempotencyKey returns the CreateIdempotencyKey of the receiver.
func (o *SparsePrivateKey) GetCreateIdempotencyKey() (out string) {

	if o.CreateIdempotencyKey == nil {
		return
	}

	return *o.CreateIdempotencyKey
}

// SetCreateIdempotencyKey sets the property CreateIdempotencyKey of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetCreateIdempotencyKey(createIdempotencyKey string) {

	o.CreateIdempotencyKey = &createIdempotencyKey
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparsePrivateKey) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetDescription returns the Description of the receiver.
func (o *SparsePrivateKey) GetDescription() (out string) {

	if o.Description == nil {
		return
	}

	return *o.Description
}

// SetDescription sets the property Description of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetDescription(description string) {

	o.Description = &description
}

// GetName returns the Name of the receiver.
func (o *SparsePrivateKey) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparsePrivateKey) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetNormalizedTags returns the NormalizedTags of the receiver.
func (o *SparsePrivateKey) GetNormalizedTags() (out []string) {

	if o.NormalizedTags == nil {
		return
	}

	return *o.NormalizedTags
}

// SetNormalizedTags sets the property NormalizedTags of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetNormalizedTags(normalizedTags []string) {

	o.NormalizedTags = &normalizedTags
}

// GetPropagate returns the Propagate of the receiver.
func (o *SparsePrivateKey) GetPropagate() (out bool) {

	if o.Propagate == nil {
		return
	}

	return *o.Propagate
}

// SetPropagate sets the property Propagate of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetPropagate(propagate bool) {

	o.Propagate = &propagate
}

// GetProtected returns the Protected of the receiver.
func (o *SparsePrivateKey) GetProtected() (out bool) {

	if o.Protected == nil {
		return
	}

	return *o.Protected
}

// SetProtected sets the property Protected of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetProtected(protected bool) {

	o.Protected = &protected
}

// GetUpdateIdempotencyKey returns the UpdateIdempotencyKey of the receiver.
func (o *SparsePrivateKey) GetUpdateIdempotencyKey() (out string) {

	if o.UpdateIdempotencyKey == nil {
		return
	}

	return *o.UpdateIdempotencyKey
}

// SetUpdateIdempotencyKey sets the property UpdateIdempotencyKey of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetUpdateIdempotencyKey(updateIdempotencyKey string) {

	o.UpdateIdempotencyKey = &updateIdempotencyKey
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparsePrivateKey) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparsePrivateKey) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparsePrivateKey) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparsePrivateKey) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparsePrivateKey.
func (o *SparsePrivateKey) DeepCopy() *SparsePrivateKey {

	if o == nil {
		return nil
	}

	out := &SparsePrivateKey{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePrivateKey.
func (o *SparsePrivateKey) DeepCopyInto(out *SparsePrivateKey) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePrivateKey: %s", err))
	}

	*out = *target.(*SparsePrivateKey)
}

type mongoAttributesPrivateKey struct {
	ID                   bson.ObjectId       `bson:"_id,omitempty"`
	Annotations          map[string][]string `bson:"annotations"`
	AssociatedTags       []string            `bson:"associatedtags"`
	CreateIdempotencyKey string              `bson:"createidempotencykey"`
	CreateTime           time.Time           `bson:"createtime"`
	Description          string              `bson:"description"`
	Fingerprint          string              `bson:"fingerprint"`
	Key                  string              `bson:"key"`
	Name                 string              `bson:"name"`
	Namespace            string              `bson:"namespace"`
	NormalizedTags       []string            `bson:"normalizedtags"`
	Propagate            bool                `bson:"propagate"`
	Protected            bool                `bson:"protected"`
	UpdateIdempotencyKey string              `bson:"updateidempotencykey"`
	UpdateTime           time.Time           `bson:"updatetime"`
	ZHash                int                 `bson:"zhash"`
	Zone                 int                 `bson:"zone"`
}
type mongoAttributesSparsePrivateKey struct {
	ID                   bson.ObjectId        `bson:"_id,omitempty"`
	Annotations          *map[string][]string `bson:"annotations,omitempty"`
	AssociatedTags       *[]string            `bson:"associatedtags,omitempty"`
	CreateIdempotencyKey *string              `bson:"createidempotencykey,omitempty"`
	CreateTime           *time.Time           `bson:"createtime,omitempty"`
	Description          *string              `bson:"description,omitempty"`
	Fingerprint          *string              `bson:"fingerprint,omitempty"`
	Key                  *string              `bson:"key,omitempty"`
	Name                 *string              `bson:"name,omitempty"`
	Namespace            *string              `bson:"namespace,omitempty"`
	NormalizedTags       *[]string            `bson:"normalizedtags,omitempty"`
	Propagate            *bool                `bson:"propagate,omitempty"`
	Protected            *bool                `bson:"protected,omitempty"`
	UpdateIdempotencyKey *string              `bson:"updateidempotencykey,omitempty"`
	UpdateTime           *time.Time           `bson:"updatetime,omitempty"`
	ZHash                *int                 `bson:"zhash,omitempty"`
	Zone                 *int                 `bson:"zone,omitempty"`
}
