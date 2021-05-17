package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PrismaCloudConfigurationIdentity represents the Identity of the object.
var PrismaCloudConfigurationIdentity = elemental.Identity{
	Name:     "prismacloudconfiguration",
	Category: "prismacloudconfiguration",
	Package:  "karl",
	Private:  false,
}

// PrismaCloudConfigurationsList represents a list of PrismaCloudConfigurations
type PrismaCloudConfigurationsList []*PrismaCloudConfiguration

// Identity returns the identity of the objects in the list.
func (o PrismaCloudConfigurationsList) Identity() elemental.Identity {

	return PrismaCloudConfigurationIdentity
}

// Copy returns a pointer to a copy the PrismaCloudConfigurationsList.
func (o PrismaCloudConfigurationsList) Copy() elemental.Identifiables {

	copy := append(PrismaCloudConfigurationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PrismaCloudConfigurationsList.
func (o PrismaCloudConfigurationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PrismaCloudConfigurationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PrismaCloudConfiguration))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PrismaCloudConfigurationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PrismaCloudConfigurationsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PrismaCloudConfigurationsList converted to SparsePrismaCloudConfigurationsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PrismaCloudConfigurationsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePrismaCloudConfigurationsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePrismaCloudConfiguration)
	}

	return out
}

// Version returns the version of the content.
func (o PrismaCloudConfigurationsList) Version() int {

	return 1
}

// PrismaCloudConfiguration represents the model of a prismacloudconfiguration
type PrismaCloudConfiguration struct {
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

// NewPrismaCloudConfiguration returns a new *PrismaCloudConfiguration
func NewPrismaCloudConfiguration() *PrismaCloudConfiguration {

	return &PrismaCloudConfiguration{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PrismaCloudConfiguration) Identity() elemental.Identity {

	return PrismaCloudConfigurationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PrismaCloudConfiguration) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PrismaCloudConfiguration) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PrismaCloudConfiguration) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPrismaCloudConfiguration{}

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
func (o *PrismaCloudConfiguration) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPrismaCloudConfiguration{}
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
func (o *PrismaCloudConfiguration) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PrismaCloudConfiguration) BleveType() string {

	return "prismacloudconfiguration"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PrismaCloudConfiguration) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PrismaCloudConfiguration) Doc() string {

	return `Holds the various Prisma Cloud configuration for a namespace.`
}

func (o *PrismaCloudConfiguration) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PrismaCloudConfiguration) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePrismaCloudConfiguration{
			ID:                            &o.ID,
			EnableNetEffectivePermissions: &o.EnableNetEffectivePermissions,
			EnableNetworkSecurity:         &o.EnableNetworkSecurity,
			Key:                           &o.Key,
		}
	}

	sp := &SparsePrismaCloudConfiguration{}
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

// Patch apply the non nil value of a *SparsePrismaCloudConfiguration to the object.
func (o *PrismaCloudConfiguration) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePrismaCloudConfiguration)
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

// DeepCopy returns a deep copy if the PrismaCloudConfiguration.
func (o *PrismaCloudConfiguration) DeepCopy() *PrismaCloudConfiguration {

	if o == nil {
		return nil
	}

	out := &PrismaCloudConfiguration{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PrismaCloudConfiguration.
func (o *PrismaCloudConfiguration) DeepCopyInto(out *PrismaCloudConfiguration) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PrismaCloudConfiguration: %s", err))
	}

	*out = *target.(*PrismaCloudConfiguration)
}

// Validate valides the current information stored into the structure.
func (o *PrismaCloudConfiguration) Validate() error {

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
func (*PrismaCloudConfiguration) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PrismaCloudConfigurationAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PrismaCloudConfigurationLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PrismaCloudConfiguration) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PrismaCloudConfigurationAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PrismaCloudConfiguration) ValueForAttribute(name string) interface{} {

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

// PrismaCloudConfigurationAttributesMap represents the map of attribute for PrismaCloudConfiguration.
var PrismaCloudConfigurationAttributesMap = map[string]elemental.AttributeSpecification{
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

// PrismaCloudConfigurationLowerCaseAttributesMap represents the map of attribute for PrismaCloudConfiguration.
var PrismaCloudConfigurationLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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

// SparsePrismaCloudConfigurationsList represents a list of SparsePrismaCloudConfigurations
type SparsePrismaCloudConfigurationsList []*SparsePrismaCloudConfiguration

// Identity returns the identity of the objects in the list.
func (o SparsePrismaCloudConfigurationsList) Identity() elemental.Identity {

	return PrismaCloudConfigurationIdentity
}

// Copy returns a pointer to a copy the SparsePrismaCloudConfigurationsList.
func (o SparsePrismaCloudConfigurationsList) Copy() elemental.Identifiables {

	copy := append(SparsePrismaCloudConfigurationsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePrismaCloudConfigurationsList.
func (o SparsePrismaCloudConfigurationsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePrismaCloudConfigurationsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePrismaCloudConfiguration))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePrismaCloudConfigurationsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePrismaCloudConfigurationsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePrismaCloudConfigurationsList converted to PrismaCloudConfigurationsList.
func (o SparsePrismaCloudConfigurationsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePrismaCloudConfigurationsList) Version() int {

	return 1
}

// SparsePrismaCloudConfiguration represents the sparse version of a prismacloudconfiguration.
type SparsePrismaCloudConfiguration struct {
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

// NewSparsePrismaCloudConfiguration returns a new  SparsePrismaCloudConfiguration.
func NewSparsePrismaCloudConfiguration() *SparsePrismaCloudConfiguration {
	return &SparsePrismaCloudConfiguration{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePrismaCloudConfiguration) Identity() elemental.Identity {

	return PrismaCloudConfigurationIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePrismaCloudConfiguration) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePrismaCloudConfiguration) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePrismaCloudConfiguration) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePrismaCloudConfiguration{}

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
func (o *SparsePrismaCloudConfiguration) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePrismaCloudConfiguration{}
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
func (o *SparsePrismaCloudConfiguration) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePrismaCloudConfiguration) ToPlain() elemental.PlainIdentifiable {

	out := NewPrismaCloudConfiguration()
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

// DeepCopy returns a deep copy if the SparsePrismaCloudConfiguration.
func (o *SparsePrismaCloudConfiguration) DeepCopy() *SparsePrismaCloudConfiguration {

	if o == nil {
		return nil
	}

	out := &SparsePrismaCloudConfiguration{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePrismaCloudConfiguration.
func (o *SparsePrismaCloudConfiguration) DeepCopyInto(out *SparsePrismaCloudConfiguration) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePrismaCloudConfiguration: %s", err))
	}

	*out = *target.(*SparsePrismaCloudConfiguration)
}

type mongoAttributesPrismaCloudConfiguration struct {
	ID                            bson.ObjectId `bson:"_id,omitempty"`
	EnableNetEffectivePermissions bool          `bson:"enableneteffectivepermissions"`
	EnableNetworkSecurity         bool          `bson:"enablenetworksecurity"`
	Key                           string        `bson:"key"`
}
type mongoAttributesSparsePrismaCloudConfiguration struct {
	ID                            bson.ObjectId `bson:"_id,omitempty"`
	EnableNetEffectivePermissions *bool         `bson:"enableneteffectivepermissions,omitempty"`
	EnableNetworkSecurity         *bool         `bson:"enablenetworksecurity,omitempty"`
	Key                           *string       `bson:"key,omitempty"`
}
