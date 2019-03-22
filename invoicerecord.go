package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// InvoiceRecordIdentity represents the Identity of the object.
var InvoiceRecordIdentity = elemental.Identity{
	Name:     "invoicerecord",
	Category: "invoicerecords",
	Package:  "bill",
	Private:  false,
}

// InvoiceRecordsList represents a list of InvoiceRecords
type InvoiceRecordsList []*InvoiceRecord

// Identity returns the identity of the objects in the list.
func (o InvoiceRecordsList) Identity() elemental.Identity {

	return InvoiceRecordIdentity
}

// Copy returns a pointer to a copy the InvoiceRecordsList.
func (o InvoiceRecordsList) Copy() elemental.Identifiables {

	copy := append(InvoiceRecordsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the InvoiceRecordsList.
func (o InvoiceRecordsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(InvoiceRecordsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*InvoiceRecord))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o InvoiceRecordsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o InvoiceRecordsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the InvoiceRecordsList converted to SparseInvoiceRecordsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o InvoiceRecordsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o InvoiceRecordsList) Version() int {

	return 1
}

// InvoiceRecord represents the model of a invoicerecord
type InvoiceRecord struct {
	// ID is the id of this invoice record.
	ID string `json:"ID" bson:"id" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime time.Time `json:"createTime" bson:"createtime" mapstructure:"createTime,omitempty"`

	// InvoiceID references the id of the invoice that this invoice record provides
	// details for.
	InvoiceID string `json:"invoiceID" bson:"invoiceid" mapstructure:"invoiceID,omitempty"`

	// InvoiceRecords provides details about billing units.
	InvoiceRecords []string `json:"invoiceRecords" bson:"invoicerecords" mapstructure:"invoiceRecords,omitempty"`

	// Last update date of the object.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewInvoiceRecord returns a new *InvoiceRecord
func NewInvoiceRecord() *InvoiceRecord {

	return &InvoiceRecord{
		ModelVersion:   1,
		InvoiceRecords: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *InvoiceRecord) Identity() elemental.Identity {

	return InvoiceRecordIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *InvoiceRecord) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *InvoiceRecord) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *InvoiceRecord) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *InvoiceRecord) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *InvoiceRecord) Doc() string {
	return `This api allows to view detailed records of invoices for Aporeto customers.`
}

func (o *InvoiceRecord) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *InvoiceRecord) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseInvoiceRecord{
			ID:             &o.ID,
			CreateTime:     &o.CreateTime,
			InvoiceID:      &o.InvoiceID,
			InvoiceRecords: &o.InvoiceRecords,
			UpdateTime:     &o.UpdateTime,
		}
	}

	sp := &SparseInvoiceRecord{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "createTime":
			sp.CreateTime = &(o.CreateTime)
		case "invoiceID":
			sp.InvoiceID = &(o.InvoiceID)
		case "invoiceRecords":
			sp.InvoiceRecords = &(o.InvoiceRecords)
		case "updateTime":
			sp.UpdateTime = &(o.UpdateTime)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseInvoiceRecord to the object.
func (o *InvoiceRecord) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseInvoiceRecord)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.CreateTime != nil {
		o.CreateTime = *so.CreateTime
	}
	if so.InvoiceID != nil {
		o.InvoiceID = *so.InvoiceID
	}
	if so.InvoiceRecords != nil {
		o.InvoiceRecords = *so.InvoiceRecords
	}
	if so.UpdateTime != nil {
		o.UpdateTime = *so.UpdateTime
	}
}

// DeepCopy returns a deep copy if the InvoiceRecord.
func (o *InvoiceRecord) DeepCopy() *InvoiceRecord {

	if o == nil {
		return nil
	}

	out := &InvoiceRecord{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *InvoiceRecord.
func (o *InvoiceRecord) DeepCopyInto(out *InvoiceRecord) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy InvoiceRecord: %s", err))
	}

	*out = *target.(*InvoiceRecord)
}

// Validate valides the current information stored into the structure.
func (o *InvoiceRecord) Validate() error {

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
func (*InvoiceRecord) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := InvoiceRecordAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return InvoiceRecordLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*InvoiceRecord) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return InvoiceRecordAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *InvoiceRecord) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "createTime":
		return o.CreateTime
	case "invoiceID":
		return o.InvoiceID
	case "invoiceRecords":
		return o.InvoiceRecords
	case "updateTime":
		return o.UpdateTime
	}

	return nil
}

// InvoiceRecordAttributesMap represents the map of attribute for InvoiceRecord.
var InvoiceRecordAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the id of this invoice record.`,
		Exposed:        true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"InvoiceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InvoiceID",
		Description: `InvoiceID references the id of the invoice that this invoice record provides
details for.`,
		Exposed:   true,
		Name:      "invoiceID",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"InvoiceRecords": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InvoiceRecords",
		Description:    `InvoiceRecords provides details about billing units.`,
		Exposed:        true,
		Name:           "invoiceRecords",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}

// InvoiceRecordLowerCaseAttributesMap represents the map of attribute for InvoiceRecord.
var InvoiceRecordLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the id of this invoice record.`,
		Exposed:        true,
		Name:           "ID",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"createtime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "CreateTime",
		Description:    `Creation date of the object.`,
		Exposed:        true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
	"invoiceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InvoiceID",
		Description: `InvoiceID references the id of the invoice that this invoice record provides
details for.`,
		Exposed:   true,
		Name:      "invoiceID",
		Orderable: true,
		Stored:    true,
		Type:      "string",
	},
	"invoicerecords": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "InvoiceRecords",
		Description:    `InvoiceRecords provides details about billing units.`,
		Exposed:        true,
		Name:           "invoiceRecords",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"updatetime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "UpdateTime",
		Description:    `Last update date of the object.`,
		Exposed:        true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "time",
	},
}

// SparseInvoiceRecordsList represents a list of SparseInvoiceRecords
type SparseInvoiceRecordsList []*SparseInvoiceRecord

// Identity returns the identity of the objects in the list.
func (o SparseInvoiceRecordsList) Identity() elemental.Identity {

	return InvoiceRecordIdentity
}

// Copy returns a pointer to a copy the SparseInvoiceRecordsList.
func (o SparseInvoiceRecordsList) Copy() elemental.Identifiables {

	copy := append(SparseInvoiceRecordsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseInvoiceRecordsList.
func (o SparseInvoiceRecordsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseInvoiceRecordsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseInvoiceRecord))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseInvoiceRecordsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseInvoiceRecordsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseInvoiceRecordsList converted to InvoiceRecordsList.
func (o SparseInvoiceRecordsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseInvoiceRecordsList) Version() int {

	return 1
}

// SparseInvoiceRecord represents the sparse version of a invoicerecord.
type SparseInvoiceRecord struct {
	// ID is the id of this invoice record.
	ID *string `json:"ID,omitempty" bson:"id" mapstructure:"ID,omitempty"`

	// Creation date of the object.
	CreateTime *time.Time `json:"createTime,omitempty" bson:"createtime" mapstructure:"createTime,omitempty"`

	// InvoiceID references the id of the invoice that this invoice record provides
	// details for.
	InvoiceID *string `json:"invoiceID,omitempty" bson:"invoiceid" mapstructure:"invoiceID,omitempty"`

	// InvoiceRecords provides details about billing units.
	InvoiceRecords *[]string `json:"invoiceRecords,omitempty" bson:"invoicerecords" mapstructure:"invoiceRecords,omitempty"`

	// Last update date of the object.
	UpdateTime *time.Time `json:"updateTime,omitempty" bson:"updatetime" mapstructure:"updateTime,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewSparseInvoiceRecord returns a new  SparseInvoiceRecord.
func NewSparseInvoiceRecord() *SparseInvoiceRecord {
	return &SparseInvoiceRecord{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseInvoiceRecord) Identity() elemental.Identity {

	return InvoiceRecordIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseInvoiceRecord) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseInvoiceRecord) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseInvoiceRecord) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseInvoiceRecord) ToPlain() elemental.PlainIdentifiable {

	out := NewInvoiceRecord()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.CreateTime != nil {
		out.CreateTime = *o.CreateTime
	}
	if o.InvoiceID != nil {
		out.InvoiceID = *o.InvoiceID
	}
	if o.InvoiceRecords != nil {
		out.InvoiceRecords = *o.InvoiceRecords
	}
	if o.UpdateTime != nil {
		out.UpdateTime = *o.UpdateTime
	}

	return out
}

// DeepCopy returns a deep copy if the SparseInvoiceRecord.
func (o *SparseInvoiceRecord) DeepCopy() *SparseInvoiceRecord {

	if o == nil {
		return nil
	}

	out := &SparseInvoiceRecord{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseInvoiceRecord.
func (o *SparseInvoiceRecord) DeepCopyInto(out *SparseInvoiceRecord) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseInvoiceRecord: %s", err))
	}

	*out = *target.(*SparseInvoiceRecord)
}
