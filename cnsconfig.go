package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CNSConfigIdentity represents the Identity of the object.
var CNSConfigIdentity = elemental.Identity{
	Name:     "cnsconfig",
	Category: "cnsconfigs",
	Package:  "karl",
	Private:  false,
}

// CNSConfigsList represents a list of CNSConfigs
type CNSConfigsList []*CNSConfig

// Identity returns the identity of the objects in the list.
func (o CNSConfigsList) Identity() elemental.Identity {

	return CNSConfigIdentity
}

// Copy returns a pointer to a copy the CNSConfigsList.
func (o CNSConfigsList) Copy() elemental.Identifiables {

	copy := append(CNSConfigsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CNSConfigsList.
func (o CNSConfigsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CNSConfigsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CNSConfig))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CNSConfigsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CNSConfigsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CNSConfigsList converted to SparseCNSConfigsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CNSConfigsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCNSConfigsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCNSConfig)
	}

	return out
}

// Version returns the version of the content.
func (o CNSConfigsList) Version() int {

	return 1
}

// CNSConfig represents the model of a cnsconfig
type CNSConfig struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// If `true` net effective permissions feature is enabled.
	EnableNetEffectivePermissions bool `json:"enableNetEffectivePermissions" msgpack:"enableNetEffectivePermissions" bson:"enableneteffectivepermissions" mapstructure:"enableNetEffectivePermissions,omitempty"`

	// If `true` network security feature is enabled.
	EnableNetworkSecurity bool `json:"enableNetworkSecurity" msgpack:"enableNetworkSecurity" bson:"enablenetworksecurity" mapstructure:"enableNetworkSecurity,omitempty"`

	// The unique key of the configuration.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCNSConfig returns a new *CNSConfig
func NewCNSConfig() *CNSConfig {

	return &CNSConfig{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *CNSConfig) Identity() elemental.Identity {

	return CNSConfigIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CNSConfig) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CNSConfig) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSConfig) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCNSConfig{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.EnableNetEffectivePermissions = o.EnableNetEffectivePermissions
	s.EnableNetworkSecurity = o.EnableNetworkSecurity
	s.Key = o.Key

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CNSConfig) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCNSConfig{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.EnableNetEffectivePermissions = s.EnableNetEffectivePermissions
	o.EnableNetworkSecurity = s.EnableNetworkSecurity
	o.Key = s.Key

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CNSConfig) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CNSConfig) BleveType() string {

	return "cnsconfig"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CNSConfig) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CNSConfig) Doc() string {

	return `Holds the CNS configuration for a namespace.`
}

func (o *CNSConfig) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CNSConfig) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCNSConfig{
			ID:                            &o.ID,
			EnableNetEffectivePermissions: &o.EnableNetEffectivePermissions,
			EnableNetworkSecurity:         &o.EnableNetworkSecurity,
			Key:                           &o.Key,
		}
	}

	sp := &SparseCNSConfig{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "enableNetEffectivePermissions":
			sp.EnableNetEffectivePermissions = &(o.EnableNetEffectivePermissions)
		case "enableNetworkSecurity":
			sp.EnableNetworkSecurity = &(o.EnableNetworkSecurity)
		case "key":
			sp.Key = &(o.Key)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCNSConfig to the object.
func (o *CNSConfig) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCNSConfig)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.EnableNetEffectivePermissions != nil {
		o.EnableNetEffectivePermissions = *so.EnableNetEffectivePermissions
	}
	if so.EnableNetworkSecurity != nil {
		o.EnableNetworkSecurity = *so.EnableNetworkSecurity
	}
	if so.Key != nil {
		o.Key = *so.Key
	}
}

// DeepCopy returns a deep copy if the CNSConfig.
func (o *CNSConfig) DeepCopy() *CNSConfig {

	if o == nil {
		return nil
	}

	out := &CNSConfig{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CNSConfig.
func (o *CNSConfig) DeepCopyInto(out *CNSConfig) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CNSConfig: %s", err))
	}

	*out = *target.(*CNSConfig)
}

// Validate valides the current information stored into the structure.
func (o *CNSConfig) Validate() error {

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
func (*CNSConfig) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CNSConfigAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CNSConfigLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CNSConfig) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CNSConfigAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CNSConfig) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "enableNetEffectivePermissions":
		return o.EnableNetEffectivePermissions
	case "enableNetworkSecurity":
		return o.EnableNetworkSecurity
	case "key":
		return o.Key
	}

	return nil
}

// CNSConfigAttributesMap represents the map of attribute for CNSConfig.
var CNSConfigAttributesMap = map[string]elemental.AttributeSpecification{
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
	"EnableNetEffectivePermissions": {
		AllowedChoices: []string{},
		BSONFieldName:  "enableneteffectivepermissions",
		ConvertedName:  "EnableNetEffectivePermissions",
		Description:    `If ` + "`" + `true` + "`" + ` net effective permissions feature is enabled.`,
		Exposed:        true,
		Name:           "enableNetEffectivePermissions",
		Stored:         true,
		Type:           "boolean",
	},
	"EnableNetworkSecurity": {
		AllowedChoices: []string{},
		BSONFieldName:  "enablenetworksecurity",
		ConvertedName:  "EnableNetworkSecurity",
		Description:    `If ` + "`" + `true` + "`" + ` network security feature is enabled.`,
		Exposed:        true,
		Name:           "enableNetworkSecurity",
		Stored:         true,
		Type:           "boolean",
	},
	"Key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		Description:    `The unique key of the configuration.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

// CNSConfigLowerCaseAttributesMap represents the map of attribute for CNSConfig.
var CNSConfigLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"enableneteffectivepermissions": {
		AllowedChoices: []string{},
		BSONFieldName:  "enableneteffectivepermissions",
		ConvertedName:  "EnableNetEffectivePermissions",
		Description:    `If ` + "`" + `true` + "`" + ` net effective permissions feature is enabled.`,
		Exposed:        true,
		Name:           "enableNetEffectivePermissions",
		Stored:         true,
		Type:           "boolean",
	},
	"enablenetworksecurity": {
		AllowedChoices: []string{},
		BSONFieldName:  "enablenetworksecurity",
		ConvertedName:  "EnableNetworkSecurity",
		Description:    `If ` + "`" + `true` + "`" + ` network security feature is enabled.`,
		Exposed:        true,
		Name:           "enableNetworkSecurity",
		Stored:         true,
		Type:           "boolean",
	},
	"key": {
		AllowedChoices: []string{},
		BSONFieldName:  "key",
		ConvertedName:  "Key",
		Description:    `The unique key of the configuration.`,
		Exposed:        true,
		Name:           "key",
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseCNSConfigsList represents a list of SparseCNSConfigs
type SparseCNSConfigsList []*SparseCNSConfig

// Identity returns the identity of the objects in the list.
func (o SparseCNSConfigsList) Identity() elemental.Identity {

	return CNSConfigIdentity
}

// Copy returns a pointer to a copy the SparseCNSConfigsList.
func (o SparseCNSConfigsList) Copy() elemental.Identifiables {

	copy := append(SparseCNSConfigsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCNSConfigsList.
func (o SparseCNSConfigsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCNSConfigsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCNSConfig))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCNSConfigsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCNSConfigsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCNSConfigsList converted to CNSConfigsList.
func (o SparseCNSConfigsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCNSConfigsList) Version() int {

	return 1
}

// SparseCNSConfig represents the sparse version of a cnsconfig.
type SparseCNSConfig struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// If `true` net effective permissions feature is enabled.
	EnableNetEffectivePermissions *bool `json:"enableNetEffectivePermissions,omitempty" msgpack:"enableNetEffectivePermissions,omitempty" bson:"enableneteffectivepermissions,omitempty" mapstructure:"enableNetEffectivePermissions,omitempty"`

	// If `true` network security feature is enabled.
	EnableNetworkSecurity *bool `json:"enableNetworkSecurity,omitempty" msgpack:"enableNetworkSecurity,omitempty" bson:"enablenetworksecurity,omitempty" mapstructure:"enableNetworkSecurity,omitempty"`

	// The unique key of the configuration.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"key,omitempty" mapstructure:"key,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCNSConfig returns a new  SparseCNSConfig.
func NewSparseCNSConfig() *SparseCNSConfig {
	return &SparseCNSConfig{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCNSConfig) Identity() elemental.Identity {

	return CNSConfigIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCNSConfig) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCNSConfig) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSConfig) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCNSConfig{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.EnableNetEffectivePermissions != nil {
		s.EnableNetEffectivePermissions = o.EnableNetEffectivePermissions
	}
	if o.EnableNetworkSecurity != nil {
		s.EnableNetworkSecurity = o.EnableNetworkSecurity
	}
	if o.Key != nil {
		s.Key = o.Key
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCNSConfig) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCNSConfig{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.EnableNetEffectivePermissions != nil {
		o.EnableNetEffectivePermissions = s.EnableNetEffectivePermissions
	}
	if s.EnableNetworkSecurity != nil {
		o.EnableNetworkSecurity = s.EnableNetworkSecurity
	}
	if s.Key != nil {
		o.Key = s.Key
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCNSConfig) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCNSConfig) ToPlain() elemental.PlainIdentifiable {

	out := NewCNSConfig()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.EnableNetEffectivePermissions != nil {
		out.EnableNetEffectivePermissions = *o.EnableNetEffectivePermissions
	}
	if o.EnableNetworkSecurity != nil {
		out.EnableNetworkSecurity = *o.EnableNetworkSecurity
	}
	if o.Key != nil {
		out.Key = *o.Key
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCNSConfig.
func (o *SparseCNSConfig) DeepCopy() *SparseCNSConfig {

	if o == nil {
		return nil
	}

	out := &SparseCNSConfig{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCNSConfig.
func (o *SparseCNSConfig) DeepCopyInto(out *SparseCNSConfig) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCNSConfig: %s", err))
	}

	*out = *target.(*SparseCNSConfig)
}

type mongoAttributesCNSConfig struct {
	ID                            bson.ObjectId `bson:"_id,omitempty"`
	EnableNetEffectivePermissions bool          `bson:"enableneteffectivepermissions"`
	EnableNetworkSecurity         bool          `bson:"enablenetworksecurity"`
	Key                           string        `bson:"key"`
}
type mongoAttributesSparseCNSConfig struct {
	ID                            bson.ObjectId `bson:"_id,omitempty"`
	EnableNetEffectivePermissions *bool         `bson:"enableneteffectivepermissions,omitempty"`
	EnableNetworkSecurity         *bool         `bson:"enablenetworksecurity,omitempty"`
	Key                           *string       `bson:"key,omitempty"`
}
