package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// AWSRegisterIdentity represents the Identity of the object.
var AWSRegisterIdentity = elemental.Identity{
	Name:     "awsregister",
	Category: "awsregister",
	Package:  "bill",
	Private:  false,
}

// AWSRegistersList represents a list of AWSRegisters
type AWSRegistersList []*AWSRegister

// Identity returns the identity of the objects in the list.
func (o AWSRegistersList) Identity() elemental.Identity {

	return AWSRegisterIdentity
}

// Copy returns a pointer to a copy the AWSRegistersList.
func (o AWSRegistersList) Copy() elemental.Identifiables {

	copy := append(AWSRegistersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the AWSRegistersList.
func (o AWSRegistersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(AWSRegistersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*AWSRegister))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o AWSRegistersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o AWSRegistersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the AWSRegistersList converted to SparseAWSRegistersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o AWSRegistersList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseAWSRegistersList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseAWSRegister)
	}

	return out
}

// Version returns the version of the content.
func (o AWSRegistersList) Version() int {

	return 1
}

// AWSRegister represents the model of a awsregister
type AWSRegister struct {
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" msgpack:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// Token Provided by AWS.
	Provider string `json:"provider" msgpack:"provider" bson:"-" mapstructure:"provider,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" msgpack:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewAWSRegister returns a new *AWSRegister
func NewAWSRegister() *AWSRegister {

	return &AWSRegister{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *AWSRegister) Identity() elemental.Identity {

	return AWSRegisterIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *AWSRegister) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *AWSRegister) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AWSRegister) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesAWSRegister{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.CreateTime = o.CreateTime
	s.UpdateTime = o.UpdateTime

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *AWSRegister) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesAWSRegister{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.CreateTime = s.CreateTime
	o.UpdateTime = s.UpdateTime

	return nil
}

// Version returns the hardcoded version of the model.
func (o *AWSRegister) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *AWSRegister) BleveType() string {

	return "awsregister"
}

// DefaultOrder returns the list of default ordering fields.
func (o *AWSRegister) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *AWSRegister) Doc() string {

	return `This api allows AWS customer to register with Aporeto SaaS for billing.`
}

func (o *AWSRegister) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *AWSRegister) GetCreateTime() time.Time {

	return o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the given value.
func (o *AWSRegister) SetCreateTime(createTime time.Time) {

	o.CreateTime = createTime
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *AWSRegister) GetUpdateTime() time.Time {

	return o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the given value.
func (o *AWSRegister) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = updateTime
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *AWSRegister) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseAWSRegister{
			ID:         &o.ID,
			CreateTime: &o.CreateTime,
			Provider:   &o.Provider,
			UpdateTime: &o.UpdateTime,
		}
	}

	sp := &SparseAWSRegister{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "provider":
			sp.Provider = &(o.Provider)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseAWSRegister to the object.
func (o *AWSRegister) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseAWSRegister)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.Provider != nil {
		o.Provider = *so.Provider
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the AWSRegister.
func (o *AWSRegister) DeepCopy() *AWSRegister {

	if o == nil {
		return nil
	}

	out := &AWSRegister{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *AWSRegister.
func (o *AWSRegister) DeepCopyInto(out *AWSRegister) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy AWSRegister: %s", err))
	}

	*out = *target.(*AWSRegister)
}

// Validate valides the current information stored into the structure.
func (o *AWSRegister) Validate() error {

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
func (*AWSRegister) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := AWSRegisterAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return AWSRegisterLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*AWSRegister) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return AWSRegisterAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *AWSRegister) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "createTime":
		return o.CreateTime
	case "provider":
		return o.Provider
	case "updateTime":
		return o.UpdateTime
	}

	return nil
}

// AWSRegisterAttributesMap represents the map of attribute for AWSRegister.
var AWSRegisterAttributesMap = map[string]elemental.AttributeSpecification{
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
	"Provider": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Provider",
		Description:    `Token Provided by AWS.`,
		Exposed:        true,
		Name:           "provider",
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

// AWSRegisterLowerCaseAttributesMap represents the map of attribute for AWSRegister.
var AWSRegisterLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
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
	"provider": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Provider",
		Description:    `Token Provided by AWS.`,
		Exposed:        true,
		Name:           "provider",
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

// SparseAWSRegistersList represents a list of SparseAWSRegisters
type SparseAWSRegistersList []*SparseAWSRegister

// Identity returns the identity of the objects in the list.
func (o SparseAWSRegistersList) Identity() elemental.Identity {

	return AWSRegisterIdentity
}

// Copy returns a pointer to a copy the SparseAWSRegistersList.
func (o SparseAWSRegistersList) Copy() elemental.Identifiables {

	copy := append(SparseAWSRegistersList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseAWSRegistersList.
func (o SparseAWSRegistersList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseAWSRegistersList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseAWSRegister))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseAWSRegistersList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseAWSRegistersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseAWSRegistersList converted to AWSRegistersList.
func (o SparseAWSRegistersList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseAWSRegistersList) Version() int {

	return 1
}

// SparseAWSRegister represents the sparse version of a awsregister.
type SparseAWSRegister struct {
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" msgpack:"createTime,omitempty" bson:"createtime,omitempty" mapstructure:"createTime,omitempty"`

	// Token Provided by AWS.
	Provider *string `json:"provider,omitempty" msgpack:"provider,omitempty" bson:"-" mapstructure:"provider,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" msgpack:"updateTime,omitempty" bson:"updatetime,omitempty" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseAWSRegister returns a new  SparseAWSRegister.
func NewSparseAWSRegister() *SparseAWSRegister {
	return &SparseAWSRegister{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseAWSRegister) Identity() elemental.Identity {

	return AWSRegisterIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseAWSRegister) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseAWSRegister) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAWSRegister) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseAWSRegister{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.CreateTime != nil {
		s.CreateTime = o.CreateTime
	}
	if o.UpdateTime != nil {
		s.UpdateTime = o.UpdateTime
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseAWSRegister) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseAWSRegister{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.CreateTime != nil {
		o.CreateTime = s.CreateTime
	}
	if s.UpdateTime != nil {
		o.UpdateTime = s.UpdateTime
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseAWSRegister) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseAWSRegister) ToPlain() elemental.PlainIdentifiable {

	out := NewAWSRegister()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.Provider != nil {
		out.Provider = *o.Provider
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// GetCreateTime returns the CreateTime of the receiver.
func (o *SparseAWSRegister) GetCreateTime() (out time.Time) {

	if o.CreateTime == nil {
		return
	}

	return *o.CreateTime
}

// SetCreateTime sets the property CreateTime of the receiver using the address of the given value.
func (o *SparseAWSRegister) SetCreateTime(createTime time.Time) {

	o.CreateTime = &createTime
}

// GetUpdateTime returns the UpdateTime of the receiver.
func (o *SparseAWSRegister) GetUpdateTime() (out time.Time) {

	if o.UpdateTime == nil {
		return
	}

	return *o.UpdateTime
}

// SetUpdateTime sets the property UpdateTime of the receiver using the address of the given value.
func (o *SparseAWSRegister) SetUpdateTime(updateTime time.Time) {

	o.UpdateTime = &updateTime
}

// DeepCopy returns a deep copy if the SparseAWSRegister.
func (o *SparseAWSRegister) DeepCopy() *SparseAWSRegister {

	if o == nil {
		return nil
	}

	out := &SparseAWSRegister{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseAWSRegister.
func (o *SparseAWSRegister) DeepCopyInto(out *SparseAWSRegister) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseAWSRegister: %s", err))
	}

	*out = *target.(*SparseAWSRegister)
}

type mongoAttributesAWSRegister struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	CreateTime time.Time     `bson:"createtime"`
	UpdateTime time.Time     `bson:"updatetime"`
}
type mongoAttributesSparseAWSRegister struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	CreateTime *time.Time    `bson:"createtime,omitempty"`
	UpdateTime *time.Time    `bson:"updatetime,omitempty"`
}
