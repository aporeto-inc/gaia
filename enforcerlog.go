package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerLogIdentity represents the Identity of the object.
var EnforcerLogIdentity = elemental.Identity{
	Name:     "enforcerlog",
	Category: "enforcerlog",
	Package:  "squall",
	Private:  false,
}

// EnforcerLogsList represents a list of EnforcerLogs
type EnforcerLogsList []*EnforcerLog

// Identity returns the identity of the objects in the list.
func (o EnforcerLogsList) Identity() elemental.Identity {

	return EnforcerLogIdentity
}

// Copy returns a pointer to a copy the EnforcerLogsList.
func (o EnforcerLogsList) Copy() elemental.Identifiables {

	copy := append(EnforcerLogsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerLogsList.
func (o EnforcerLogsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerLogsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerLog))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerLogsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerLogsList) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// ToSparse returns the EnforcerLogsList converted to SparseEnforcerLogsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcerLogsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseEnforcerLogsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseEnforcerLog)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcerLogsList) Version() int {

	return 1
}

// EnforcerLog represents the model of a enforcerlog
type EnforcerLog struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// CollectionID is the ID of the enforcer log. CollectionID is used to
	// aggergate the multipart data.
	CollectionID string `json:"collectionID" msgpack:"collectionID" bson:"-" mapstructure:"collectionID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Represents the data collected by the enforcer.
	Data string `json:"data" msgpack:"data" bson:"data" mapstructure:"data,omitempty"`

	// ID of the enforcer.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Number assigned to each log in the increasing order.
	Page int `json:"page" msgpack:"page" bson:"-" mapstructure:"page,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone int `json:"zone" msgpack:"zone" bson:"zone" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEnforcerLog returns a new *EnforcerLog
func NewEnforcerLog() *EnforcerLog {

	return &EnforcerLog{
		ModelVersion:  1,
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerLog) Identity() elemental.Identity {

	return EnforcerLogIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerLog) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerLog) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *EnforcerLog) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *EnforcerLog) BleveType() string {

	return "enforcerlog"
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerLog) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// Doc returns the documentation for the object
func (o *EnforcerLog) Doc() string {

	return `An enforcer log represents the log collected by an enforcer. Each enforcer log
can have partial or complete data. The collectionID is used to aggregate the
multipart data into one.`
}

func (o *EnforcerLog) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *EnforcerLog) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *EnforcerLog) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *EnforcerLog) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *EnforcerLog) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *EnforcerLog) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *EnforcerLog) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *EnforcerLog) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *EnforcerLog) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *EnforcerLog) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *EnforcerLog) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *EnforcerLog) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *EnforcerLog) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EnforcerLog) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcerLog{
			ID:            &o.ID,
			CollectionID:  &o.CollectionID,
			CreateTime:    &o.CreateTime,
			Data:          &o.Data,
			EnforcerID:    &o.EnforcerID,
			MigrationsLog: &o.MigrationsLog,
			Namespace:     &o.Namespace,
			Page:          &o.Page,
			UpdateTime:    &o.UpdateTime,
			ZHash:         &o.ZHash,
			Zone:          &o.Zone,
		}
	}

	sp := &SparseEnforcerLog{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "collectionID":
			sp.CollectionID = &(o.CollectionID)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "data":
			sp.Data = &(o.Data)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "page":
			sp.Page = &(o.Page)
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

// Patch apply the non nil value of a *SparseEnforcerLog to the object.
func (o *EnforcerLog) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcerLog)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.CollectionID != nil {
		o.CollectionID = *so.CollectionID
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Data != nil {
		o.Data = *so.Data
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Page != nil {
		o.Page = *so.Page
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

// DeepCopy returns a deep copy if the EnforcerLog.
func (o *EnforcerLog) DeepCopy() *EnforcerLog {

	if o == nil {
		return nil
	}

	out := &EnforcerLog{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EnforcerLog.
func (o *EnforcerLog) DeepCopyInto(out *EnforcerLog) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EnforcerLog: %s", err))
	}

	*out = *target.(*EnforcerLog)
}

// Validate valides the current information stored into the structure.
func (o *EnforcerLog) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("collectionID", o.CollectionID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
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
func (*EnforcerLog) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerLogAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerLogLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerLog) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerLogAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EnforcerLog) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "collectionID":
		return o.CollectionID
	case "createTime":
		return o.CreateTime
	case "data":
		return o.Data
	case "enforcerID":
		return o.EnforcerID
	case "migrationsLog":
		return o.MigrationsLog
	case "namespace":
		return o.Namespace
	case "page":
		return o.Page
	case "updateTime":
		return o.UpdateTime
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// EnforcerLogAttributesMap represents the map of attribute for EnforcerLog.
var EnforcerLogAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"CollectionID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectionID",
		Description: `CollectionID is the ID of the enforcer log. CollectionID is used to
aggergate the multipart data.`,
		Exposed:  true,
		Name:     "collectionID",
		Required: true,
		Type:     "string",
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
	"Data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Represents the data collected by the enforcer.`,
		Exposed:        true,
		Name:           "data",
		Stored:         true,
		Type:           "string",
	},
	"EnforcerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Type:           "string",
	},
	"MigrationsLog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
	"Page": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Page",
		Description:    `Number assigned to each log in the increasing order.`,
		Exposed:        true,
		Name:           "page",
		Type:           "integer",
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
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// EnforcerLogLowerCaseAttributesMap represents the map of attribute for EnforcerLog.
var EnforcerLogLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	"collectionid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CollectionID",
		Description: `CollectionID is the ID of the enforcer log. CollectionID is used to
aggergate the multipart data.`,
		Exposed:  true,
		Name:     "collectionID",
		Required: true,
		Type:     "string",
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
	"data": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Data",
		Description:    `Represents the data collected by the enforcer.`,
		Exposed:        true,
		Name:           "data",
		Stored:         true,
		Type:           "string",
	},
	"enforcerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Type:           "string",
	},
	"migrationslog": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "MigrationsLog",
		Description:    `Internal property maintaining migrations information.`,
		Getter:         true,
		Name:           "migrationsLog",
		Setter:         true,
		Stored:         true,
		SubType:        "map[string]string",
		Type:           "external",
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
	"page": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Page",
		Description:    `Number assigned to each log in the increasing order.`,
		Exposed:        true,
		Name:           "page",
		Type:           "integer",
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
		Description:    `Geographical zone. Used for sharding and georedundancy.`,
		Exposed:        true,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseEnforcerLogsList represents a list of SparseEnforcerLogs
type SparseEnforcerLogsList []*SparseEnforcerLog

// Identity returns the identity of the objects in the list.
func (o SparseEnforcerLogsList) Identity() elemental.Identity {

	return EnforcerLogIdentity
}

// Copy returns a pointer to a copy the SparseEnforcerLogsList.
func (o SparseEnforcerLogsList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcerLogsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcerLogsList.
func (o SparseEnforcerLogsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcerLogsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcerLog))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcerLogsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcerLogsList) DefaultOrder() []string {

	return []string{
		"namespace",
	}
}

// ToPlain returns the SparseEnforcerLogsList converted to EnforcerLogsList.
func (o SparseEnforcerLogsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcerLogsList) Version() int {

	return 1
}

// SparseEnforcerLog represents the sparse version of a enforcerlog.
type SparseEnforcerLog struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"_id" mapstructure:"ID,omitempty"`

	// CollectionID is the ID of the enforcer log. CollectionID is used to
	// aggergate the multipart data.
	CollectionID *string `json:"collectionID,omitempty" msgpack:"collectionID,omitempty" bson:"-" mapstructure:"collectionID,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Represents the data collected by the enforcer.
	Data *string `json:"data,omitempty" msgpack:"data,omitempty" bson:"data,omitempty" mapstructure:"data,omitempty"`

	// ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace tag attached to an entity.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Number assigned to each log in the increasing order.
	Page *int `json:"page,omitempty" msgpack:"page,omitempty" bson:"-" mapstructure:"page,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Geographical zone. Used for sharding and georedundancy.
	Zone *int `json:"zone,omitempty" msgpack:"zone,omitempty" bson:"zone,omitempty" mapstructure:"zone,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseEnforcerLog returns a new  SparseEnforcerLog.
func NewSparseEnforcerLog() *SparseEnforcerLog {
	return &SparseEnforcerLog{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcerLog) Identity() elemental.Identity {

	return EnforcerLogIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcerLog) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcerLog) SetIdentifier(id string) {

	o.ID = &id
}

// Version returns the hardcoded version of the model.
func (o *SparseEnforcerLog) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcerLog) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcerLog()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.CollectionID != nil {
		out.CollectionID = *o.CollectionID
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Data != nil {
		out.Data = *o.Data
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Page != nil {
		out.Page = *o.Page
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
func (o *SparseEnforcerLog) GetCreateTime() time.Time {

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseEnforcerLog) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparseEnforcerLog) GetMigrationsLog() map[string]string {

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparseEnforcerLog) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseEnforcerLog) GetNamespace() string {

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseEnforcerLog) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseEnforcerLog) GetUpdateTime() time.Time {

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseEnforcerLog) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseEnforcerLog) GetZHash() int {

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseEnforcerLog) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseEnforcerLog) GetZone() int {

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseEnforcerLog) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseEnforcerLog.
func (o *SparseEnforcerLog) DeepCopy() *SparseEnforcerLog {

	if o == nil {
		return nil
	}

	out := &SparseEnforcerLog{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcerLog.
func (o *SparseEnforcerLog) DeepCopyInto(out *SparseEnforcerLog) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcerLog: %s", err))
	}

	*out = *target.(*SparseEnforcerLog)
}
