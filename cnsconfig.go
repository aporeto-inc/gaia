package gaia

import (
	"fmt"
	"time"

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

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// If `true` net effective permissions feature is enabled.
	EnableNetEffectivePermissions bool `json:"enableNetEffectivePermissions" msgpack:"enableNetEffectivePermissions" bson:"enableneteffectivepermissions" mapstructure:"enableNetEffectivePermissions,omitempty"`

	// If `true` network security feature is enabled.
	EnableNetworkSecurity bool `json:"enableNetworkSecurity" msgpack:"enableNetworkSecurity" bson:"enablenetworksecurity" mapstructure:"enableNetworkSecurity,omitempty"`

	// The unique key of the configuration.
	Key string `json:"key" msgpack:"key" bson:"key" mapstructure:"key,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCNSConfig returns a new *CNSConfig
func NewCNSConfig() *CNSConfig {

	return &CNSConfig{
		ModelVersion:  1,
		MigrationsLog: map[string]string{},
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
	s.CreateTime = o.CreateTime
	s.EnableNetEffectivePermissions = o.EnableNetEffectivePermissions
	s.EnableNetworkSecurity = o.EnableNetworkSecurity
	s.Key = o.Key
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.UpdateTime = o.UpdateTime
	s.ZHash = o.ZHash
	s.Zone = o.Zone

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
	o.CreateTime = s.CreateTime
	o.EnableNetEffectivePermissions = s.EnableNetEffectivePermissions
	o.EnableNetworkSecurity = s.EnableNetworkSecurity
	o.Key = s.Key
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.UpdateTime = s.UpdateTime
	o.ZHash = s.ZHash
	o.Zone = s.Zone

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

// GetCreateTime returns the CreateTime of the receiver.
func (o *CNSConfig) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *CNSConfig) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *CNSConfig) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *CNSConfig) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *CNSConfig) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *CNSConfig) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *CNSConfig) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *CNSConfig) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *CNSConfig) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *CNSConfig) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *CNSConfig) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *CNSConfig) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CNSConfig) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCNSConfig{
			ID:                            &o.ID,
			CreateTime:                    &o.CreateTime,
			EnableNetEffectivePermissions: &o.EnableNetEffectivePermissions,
			EnableNetworkSecurity:         &o.EnableNetworkSecurity,
			Key:                           &o.Key,
			MigrationsLog:                 &o.MigrationsLog,
			Namespace:                     &o.Namespace,
			UpdateTime:                    &o.UpdateTime,
			ZHash:                         &o.ZHash,
			Zone:                          &o.Zone,
		}
	}

	sp := &SparseCNSConfig{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "enableNetEffectivePermissions":
			sp.EnableNetEffectivePermissions = &(o.EnableNetEffectivePermissions)
		case "enableNetworkSecurity":
			sp.EnableNetworkSecurity = &(o.EnableNetworkSecurity)
		case "key":
			sp.Key = &(o.Key)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
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

// Patch apply the non nil value of a *SparseCNSConfig to the object.
func (o *CNSConfig) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCNSConfig)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
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
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
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
	case "createTime":
		return o.CreateTime
	case "enableNetEffectivePermissions":
		return o.EnableNetEffectivePermissions
	case "enableNetworkSecurity":
		return o.EnableNetworkSecurity
	case "key":
		return o.Key
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
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
	"MigrationsLog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
	"migrationslog": {
		AllowedChoices: []string{},
		BSONFieldName:  "migrationslog",
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// If `true` net effective permissions feature is enabled.
	EnableNetEffectivePermissions *bool `json:"enableNetEffectivePermissions,omitempty" msgpack:"enableNetEffectivePermissions,omitempty" bson:"enableneteffectivepermissions,omitempty" mapstructure:"enableNetEffectivePermissions,omitempty"`

	// If `true` network security feature is enabled.
	EnableNetworkSecurity *bool `json:"enableNetworkSecurity,omitempty" msgpack:"enableNetworkSecurity,omitempty" bson:"enablenetworksecurity,omitempty" mapstructure:"enableNetworkSecurity,omitempty"`

	// The unique key of the configuration.
	Key *string `json:"key,omitempty" msgpack:"key,omitempty" bson:"key,omitempty" mapstructure:"key,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

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
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
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
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
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
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.EnableNetEffectivePermissions != nil {
		o.EnableNetEffectivePermissions = s.EnableNetEffectivePermissions
	}
	if s.EnableNetworkSecurity != nil {
		o.EnableNetworkSecurity = s.EnableNetworkSecurity
	}
	if s.Key != nil {
		o.Key = s.Key
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
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
func (o *SparseCNSConfig) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCNSConfig) ToPlain() elemental.PlainIdentifiable {

	out := NewCNSConfig()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
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
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
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

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseCNSConfig) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseCNSConfig) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseCNSConfig) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseCNSConfig) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseCNSConfig) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseCNSConfig) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseCNSConfig) SetZone(zone int) {

	o.Zone = &zone
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
	ID                            bson.ObjectId     `bson:"_id,omitempty"`
	CreateTime                    time.Time         `bson:"createtime"`
	EnableNetEffectivePermissions bool              `bson:"enableneteffectivepermissions"`
	EnableNetworkSecurity         bool              `bson:"enablenetworksecurity"`
	Key                           string            `bson:"key"`
	MigrationsLog                 map[string]string `bson:"migrationslog,omitempty"`
	Namespace                     string            `bson:"namespace"`
	UpdateTime                    time.Time         `bson:"updatetime"`
	ZHash                         int               `bson:"zhash"`
	Zone                          int               `bson:"zone"`
}
type mongoAttributesSparseCNSConfig struct {
	ID                            bson.ObjectId      `bson:"_id,omitempty"`
	CreateTime                    *time.Time         `bson:"createtime,omitempty"`
	EnableNetEffectivePermissions *bool              `bson:"enableneteffectivepermissions,omitempty"`
	EnableNetworkSecurity         *bool              `bson:"enablenetworksecurity,omitempty"`
	Key                           *string            `bson:"key,omitempty"`
	MigrationsLog                 *map[string]string `bson:"migrationslog,omitempty"`
	Namespace                     *string            `bson:"namespace,omitempty"`
	UpdateTime                    *time.Time         `bson:"updatetime,omitempty"`
	ZHash                         *int               `bson:"zhash,omitempty"`
	Zone                          *int               `bson:"zone,omitempty"`
}
