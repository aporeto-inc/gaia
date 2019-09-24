package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// EnforcerReportIdentity represents the Identity of the object.
var EnforcerReportIdentity = elemental.Identity{
	Name:     "enforcerreport",
	Category: "enforcerreports",
	Package:  "zack",
	Private:  false,
}

// EnforcerReportsList represents a list of EnforcerReports
type EnforcerReportsList []*EnforcerReport

// Identity returns the identity of the objects in the list.
func (o EnforcerReportsList) Identity() elemental.Identity {

	return EnforcerReportIdentity
}

// Copy returns a pointer to a copy the EnforcerReportsList.
func (o EnforcerReportsList) Copy() elemental.Identifiables {

	copy := append(EnforcerReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerReportsList.
func (o EnforcerReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the EnforcerReportsList converted to SparseEnforcerReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o EnforcerReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseEnforcerReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseEnforcerReport)
	}

	return out
}

// Version returns the version of the content.
func (o EnforcerReportsList) Version() int {

	return 1
}

// EnforcerReport represents the model of a enforcerreport
type EnforcerReport struct {
	// Total CPU utilization of the enforcer as a percentage of vCPUs.
	CPULoad float64 `json:"CPULoad" msgpack:"CPULoad" bson:"-" mapstructure:"CPULoad,omitempty"`

	// ID of the enforcer.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Total resident memory used by the enforcer in bytes.
	Memory int `json:"memory" msgpack:"memory" bson:"-" mapstructure:"memory,omitempty"`

	// Name of the enforcer.
	Name string `json:"name" msgpack:"name" bson:"-" mapstructure:"name,omitempty"`

	// Namespace of the enforcer.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Number of active processes of the enforcer.
	Processes int `json:"processes" msgpack:"processes" bson:"-" mapstructure:"processes,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewEnforcerReport returns a new *EnforcerReport
func NewEnforcerReport() *EnforcerReport {

	return &EnforcerReport{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerReport) Identity() elemental.Identity {

	return EnforcerReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesEnforcerReport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *EnforcerReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesEnforcerReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *EnforcerReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *EnforcerReport) BleveType() string {

	return "enforcerreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *EnforcerReport) Doc() string {

	return `Post a new enforcer statistics report.`
}

func (o *EnforcerReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *EnforcerReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseEnforcerReport{
			CPULoad:   &o.CPULoad,
			ID:        &o.ID,
			Memory:    &o.Memory,
			Name:      &o.Name,
			Namespace: &o.Namespace,
			Processes: &o.Processes,
			Timestamp: &o.Timestamp,
		}
	}

	sp := &SparseEnforcerReport{}
	for _, f := range fields {
		switch f {
		case "CPULoad":
			sp.CPULoad = &(o.CPULoad)
		case "ID":
			sp.ID = &(o.ID)
		case "memory":
			sp.Memory = &(o.Memory)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "processes":
			sp.Processes = &(o.Processes)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseEnforcerReport to the object.
func (o *EnforcerReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseEnforcerReport)
	if so.CPULoad != nil {
		o.CPULoad = *so.CPULoad
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Memory != nil {
		o.Memory = *so.Memory
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Processes != nil {
		o.Processes = *so.Processes
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
}

// DeepCopy returns a deep copy if the EnforcerReport.
func (o *EnforcerReport) DeepCopy() *EnforcerReport {

	if o == nil {
		return nil
	}

	out := &EnforcerReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *EnforcerReport.
func (o *EnforcerReport) DeepCopyInto(out *EnforcerReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy EnforcerReport: %s", err))
	}

	*out = *target.(*EnforcerReport)
}

// Validate valides the current information stored into the structure.
func (o *EnforcerReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredTime("timestamp", o.Timestamp); err != nil {
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
func (*EnforcerReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *EnforcerReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "CPULoad":
		return o.CPULoad
	case "ID":
		return o.ID
	case "memory":
		return o.Memory
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "processes":
		return o.Processes
	case "timestamp":
		return o.Timestamp
	}

	return nil
}

// EnforcerReportAttributesMap represents the map of attribute for EnforcerReport.
var EnforcerReportAttributesMap = map[string]elemental.AttributeSpecification{
	"CPULoad": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CPULoad",
		Description:    `Total CPU utilization of the enforcer as a percentage of vCPUs.`,
		Exposed:        true,
		Name:           "CPULoad",
		Type:           "float",
	},
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
	"Memory": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Memory",
		Description:    `Total resident memory used by the enforcer in bytes.`,
		Exposed:        true,
		Name:           "memory",
		Type:           "integer",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the enforcer.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"Processes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Processes",
		Description:    `Number of active processes of the enforcer.`,
		Exposed:        true,
		Name:           "processes",
		Type:           "integer",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
}

// EnforcerReportLowerCaseAttributesMap represents the map of attribute for EnforcerReport.
var EnforcerReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"cpuload": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CPULoad",
		Description:    `Total CPU utilization of the enforcer as a percentage of vCPUs.`,
		Exposed:        true,
		Name:           "CPULoad",
		Type:           "float",
	},
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
	"memory": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Memory",
		Description:    `Total resident memory used by the enforcer in bytes.`,
		Exposed:        true,
		Name:           "memory",
		Type:           "integer",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the enforcer.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"processes": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Processes",
		Description:    `Number of active processes of the enforcer.`,
		Exposed:        true,
		Name:           "processes",
		Type:           "integer",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
}

// SparseEnforcerReportsList represents a list of SparseEnforcerReports
type SparseEnforcerReportsList []*SparseEnforcerReport

// Identity returns the identity of the objects in the list.
func (o SparseEnforcerReportsList) Identity() elemental.Identity {

	return EnforcerReportIdentity
}

// Copy returns a pointer to a copy the SparseEnforcerReportsList.
func (o SparseEnforcerReportsList) Copy() elemental.Identifiables {

	copy := append(SparseEnforcerReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseEnforcerReportsList.
func (o SparseEnforcerReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseEnforcerReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseEnforcerReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseEnforcerReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseEnforcerReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseEnforcerReportsList converted to EnforcerReportsList.
func (o SparseEnforcerReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseEnforcerReportsList) Version() int {

	return 1
}

// SparseEnforcerReport represents the sparse version of a enforcerreport.
type SparseEnforcerReport struct {
	// Total CPU utilization of the enforcer as a percentage of vCPUs.
	CPULoad *float64 `json:"CPULoad,omitempty" msgpack:"CPULoad,omitempty" bson:"-" mapstructure:"CPULoad,omitempty"`

	// ID of the enforcer.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Total resident memory used by the enforcer in bytes.
	Memory *int `json:"memory,omitempty" msgpack:"memory,omitempty" bson:"-" mapstructure:"memory,omitempty"`

	// Name of the enforcer.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"-" mapstructure:"name,omitempty"`

	// Namespace of the enforcer.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// Number of active processes of the enforcer.
	Processes *int `json:"processes,omitempty" msgpack:"processes,omitempty" bson:"-" mapstructure:"processes,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseEnforcerReport returns a new  SparseEnforcerReport.
func NewSparseEnforcerReport() *SparseEnforcerReport {
	return &SparseEnforcerReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseEnforcerReport) Identity() elemental.Identity {

	return EnforcerReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseEnforcerReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseEnforcerReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcerReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseEnforcerReport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseEnforcerReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseEnforcerReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseEnforcerReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseEnforcerReport) ToPlain() elemental.PlainIdentifiable {

	out := NewEnforcerReport()
	if o.CPULoad != nil {
		out.CPULoad = *o.CPULoad
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Memory != nil {
		out.Memory = *o.Memory
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Processes != nil {
		out.Processes = *o.Processes
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}

	return out
}

// DeepCopy returns a deep copy if the SparseEnforcerReport.
func (o *SparseEnforcerReport) DeepCopy() *SparseEnforcerReport {

	if o == nil {
		return nil
	}

	out := &SparseEnforcerReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseEnforcerReport.
func (o *SparseEnforcerReport) DeepCopyInto(out *SparseEnforcerReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseEnforcerReport: %s", err))
	}

	*out = *target.(*SparseEnforcerReport)
}

type mongoAttributesEnforcerReport struct {
}
type mongoAttributesSparseEnforcerReport struct {
}
