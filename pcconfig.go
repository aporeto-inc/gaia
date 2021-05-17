package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PCConfigIdentity represents the Identity of the object.
var PCConfigIdentity = elemental.Identity{
	Name:     "pcconfig",
	Category: "pcconfig",
	Package:  "karl",
	Private:  false,
}

// PCConfigsList represents a list of PCConfigs
type PCConfigsList []*PCConfig

// Identity returns the identity of the objects in the list.
func (o PCConfigsList) Identity() elemental.Identity {

	return PCConfigIdentity
}

// Copy returns a pointer to a copy the PCConfigsList.
func (o PCConfigsList) Copy() elemental.Identifiables {

	copy := append(PCConfigsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PCConfigsList.
func (o PCConfigsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PCConfigsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PCConfig))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PCConfigsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PCConfigsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PCConfigsList converted to SparsePCConfigsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PCConfigsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePCConfigsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePCConfig)
	}

	return out
}

// Version returns the version of the content.
func (o PCConfigsList) Version() int {

	return 1
}

// PCConfig represents the model of a pcconfig
type PCConfig struct {
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

// NewPCConfig returns a new *PCConfig
func NewPCConfig() *PCConfig {

	return &PCConfig{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PCConfig) Identity() elemental.Identity {

	return PCConfigIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PCConfig) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PCConfig) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PCConfig) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPCConfig{}

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
func (o *PCConfig) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPCConfig{}
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
func (o *PCConfig) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PCConfig) BleveType() string {

	return "pcconfig"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PCConfig) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PCConfig) Doc() string {

	return `Holds the Prisma Cloud configuration for a namespace.`
}

func (o *PCConfig) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PCConfig) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePCConfig{
			ID:                            &o.ID,
			EnableNetEffectivePermissions: &o.EnableNetEffectivePermissions,
			EnableNetworkSecurity:         &o.EnableNetworkSecurity,
			Key:                           &o.Key,
		}
	}

	sp := &SparsePCConfig{}
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

// Patch apply the non nil value of a *SparsePCConfig to the object.
func (o *PCConfig) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePCConfig)
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

// DeepCopy returns a deep copy if the PCConfig.
func (o *PCConfig) DeepCopy() *PCConfig {

	if o == nil {
		return nil
	}

	out := &PCConfig{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PCConfig.
func (o *PCConfig) DeepCopyInto(out *PCConfig) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PCConfig: %s", err))
	}

	*out = *target.(*PCConfig)
}

// Validate valides the current information stored into the structure.
func (o *PCConfig) Validate() error {

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
func (*PCConfig) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PCConfigAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PCConfigLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PCConfig) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PCConfigAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PCConfig) ValueForAttribute(name string) interface{} {

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

// PCConfigAttributesMap represents the map of attribute for PCConfig.
var PCConfigAttributesMap = map[string]elemental.AttributeSpecification{
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

// PCConfigLowerCaseAttributesMap represents the map of attribute for PCConfig.
var PCConfigLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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

// SparsePCConfigsList represents a list of SparsePCConfigs
type SparsePCConfigsList []*SparsePCConfig

// Identity returns the identity of the objects in the list.
func (o SparsePCConfigsList) Identity() elemental.Identity {

	return PCConfigIdentity
}

// Copy returns a pointer to a copy the SparsePCConfigsList.
func (o SparsePCConfigsList) Copy() elemental.Identifiables {

	copy := append(SparsePCConfigsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePCConfigsList.
func (o SparsePCConfigsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePCConfigsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePCConfig))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePCConfigsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePCConfigsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePCConfigsList converted to PCConfigsList.
func (o SparsePCConfigsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePCConfigsList) Version() int {

	return 1
}

// SparsePCConfig represents the sparse version of a pcconfig.
type SparsePCConfig struct {
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

// NewSparsePCConfig returns a new  SparsePCConfig.
func NewSparsePCConfig() *SparsePCConfig {
	return &SparsePCConfig{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePCConfig) Identity() elemental.Identity {

	return PCConfigIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePCConfig) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePCConfig) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePCConfig) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePCConfig{}

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
func (o *SparsePCConfig) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePCConfig{}
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
func (o *SparsePCConfig) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePCConfig) ToPlain() elemental.PlainIdentifiable {

	out := NewPCConfig()
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

// DeepCopy returns a deep copy if the SparsePCConfig.
func (o *SparsePCConfig) DeepCopy() *SparsePCConfig {

	if o == nil {
		return nil
	}

	out := &SparsePCConfig{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePCConfig.
func (o *SparsePCConfig) DeepCopyInto(out *SparsePCConfig) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePCConfig: %s", err))
	}

	*out = *target.(*SparsePCConfig)
}

type mongoAttributesPCConfig struct {
	ID                            bson.ObjectId `bson:"_id,omitempty"`
	EnableNetEffectivePermissions bool          `bson:"enableneteffectivepermissions"`
	EnableNetworkSecurity         bool          `bson:"enablenetworksecurity"`
	Key                           string        `bson:"key"`
}
type mongoAttributesSparsePCConfig struct {
	ID                            bson.ObjectId `bson:"_id,omitempty"`
	EnableNetEffectivePermissions *bool         `bson:"enableneteffectivepermissions,omitempty"`
	EnableNetworkSecurity         *bool         `bson:"enablenetworksecurity,omitempty"`
	Key                           *string       `bson:"key,omitempty"`
}
