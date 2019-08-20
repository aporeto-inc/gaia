package gaia

import (
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// DNSReportIdentity represents the Identity of the object.
var DNSReportIdentity = elemental.Identity{
	Name:     "dnsreport",
	Category: "dnsreports",
	Package:  "zack",
	Private:  false,
}

// DNSReportsList represents a list of DNSReports
type DNSReportsList []*DNSReport

// Identity returns the identity of the objects in the list.
func (o DNSReportsList) Identity() elemental.Identity {

	return DNSReportIdentity
}

// Copy returns a pointer to a copy the DNSReportsList.
func (o DNSReportsList) Copy() elemental.Identifiables {

	copy := append(DNSReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DNSReportsList.
func (o DNSReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(DNSReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*DNSReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o DNSReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DNSReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the DNSReportsList converted to SparseDNSReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o DNSReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseDNSReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseDNSReport)
	}

	return out
}

// Version returns the version of the content.
func (o DNSReportsList) Version() int {

	return 1
}

// DNSReport represents the model of a dnsreport
type DNSReport struct {
	// Number of times the client saw this activity.
	Count int `json:"count" msgpack:"count" bson:"-" mapstructure:"count,omitempty"`

	// If the result is false, error reports the reason of the dns failure.
	Error string `json:"error" msgpack:"error" bson:"-" mapstructure:"error,omitempty"`

	// name looked up by PU.
	NameLookup string `json:"nameLookup" msgpack:"nameLookup" bson:"-" mapstructure:"nameLookup,omitempty"`

	// Result reports whether dns request succeeded or failed.
	Result bool `json:"result" msgpack:"result" bson:"-" mapstructure:"result,omitempty"`

	// ID of the source.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"-" mapstructure:"sourceID,omitempty"`

	// Type of the source.
	SourceIP string `json:"sourceIP" msgpack:"sourceIP" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Namespace of the source.
	SourceNamespace string `json:"sourceNamespace" msgpack:"sourceNamespace" bson:"-" mapstructure:"sourceNamespace,omitempty"`

	// Time and date of the log.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewDNSReport returns a new *DNSReport
func NewDNSReport() *DNSReport {

	return &DNSReport{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *DNSReport) Identity() elemental.Identity {

	return DNSReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DNSReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DNSReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *DNSReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *DNSReport) BleveType() string {

	return "dnsreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *DNSReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *DNSReport) Doc() string {

	return `Post a new dns request report.`
}

func (o *DNSReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *DNSReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseDNSReport{
			Count:           &o.Count,
			Error:           &o.Error,
			NameLookup:      &o.NameLookup,
			Result:          &o.Result,
			SourceID:        &o.SourceID,
			SourceIP:        &o.SourceIP,
			SourceNamespace: &o.SourceNamespace,
			Timestamp:       &o.Timestamp,
		}
	}

	sp := &SparseDNSReport{}
	for _, f := range fields {
		switch f {
		case "count":
			sp.Count = &(o.Count)
		case "error":
			sp.Error = &(o.Error)
		case "nameLookup":
			sp.NameLookup = &(o.NameLookup)
		case "result":
			sp.Result = &(o.Result)
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "sourceNamespace":
			sp.SourceNamespace = &(o.SourceNamespace)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseDNSReport to the object.
func (o *DNSReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseDNSReport)
	if so.Count != nil {
		o.Count = *so.Count
	}
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.NameLookup != nil {
		o.NameLookup = *so.NameLookup
	}
	if so.Result != nil {
		o.Result = *so.Result
	}
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
	}
	if so.SourceNamespace != nil {
		o.SourceNamespace = *so.SourceNamespace
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
}

// DeepCopy returns a deep copy if the DNSReport.
func (o *DNSReport) DeepCopy() *DNSReport {

	if o == nil {
		return nil
	}

	out := &DNSReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *DNSReport.
func (o *DNSReport) DeepCopyInto(out *DNSReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy DNSReport: %s", err))
	}

	*out = *target.(*DNSReport)
}

// Validate valides the current information stored into the structure.
func (o *DNSReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredInt("count", o.Count); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("nameLookup", o.NameLookup); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceID", o.SourceID); err != nil {
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
func (*DNSReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DNSReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DNSReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DNSReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DNSReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *DNSReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "count":
		return o.Count
	case "error":
		return o.Error
	case "nameLookup":
		return o.NameLookup
	case "result":
		return o.Result
	case "sourceID":
		return o.SourceID
	case "sourceIP":
		return o.SourceIP
	case "sourceNamespace":
		return o.SourceNamespace
	case "timestamp":
		return o.Timestamp
	}

	return nil
}

// DNSReportAttributesMap represents the map of attribute for DNSReport.
var DNSReportAttributesMap = map[string]elemental.AttributeSpecification{
	"Count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Count",
		Description:    `Number of times the client saw this activity.`,
		Exposed:        true,
		Name:           "count",
		Required:       true,
		Type:           "integer",
	},
	"Error": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `If the result is false, error reports the reason of the dns failure.`,
		Exposed:        true,
		Name:           "error",
		Type:           "string",
	},
	"NameLookup": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NameLookup",
		Description:    `name looked up by PU.`,
		Exposed:        true,
		Name:           "nameLookup",
		Required:       true,
		Type:           "string",
	},
	"Result": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Result",
		Description:    `Result reports whether dns request succeeded or failed.`,
		Exposed:        true,
		Name:           "result",
		Required:       true,
		Type:           "boolean",
	},
	"SourceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
	"SourceIP": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"SourceNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Description:    `Namespace of the source.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Type:           "string",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the log.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
}

// DNSReportLowerCaseAttributesMap represents the map of attribute for DNSReport.
var DNSReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Count",
		Description:    `Number of times the client saw this activity.`,
		Exposed:        true,
		Name:           "count",
		Required:       true,
		Type:           "integer",
	},
	"error": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Error",
		Description:    `If the result is false, error reports the reason of the dns failure.`,
		Exposed:        true,
		Name:           "error",
		Type:           "string",
	},
	"namelookup": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "NameLookup",
		Description:    `name looked up by PU.`,
		Exposed:        true,
		Name:           "nameLookup",
		Required:       true,
		Type:           "string",
	},
	"result": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Result",
		Description:    `Result reports whether dns request succeeded or failed.`,
		Exposed:        true,
		Name:           "result",
		Required:       true,
		Type:           "boolean",
	},
	"sourceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
	"sourceip": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"sourcenamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Description:    `Namespace of the source.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Type:           "string",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the log.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
}

// SparseDNSReportsList represents a list of SparseDNSReports
type SparseDNSReportsList []*SparseDNSReport

// Identity returns the identity of the objects in the list.
func (o SparseDNSReportsList) Identity() elemental.Identity {

	return DNSReportIdentity
}

// Copy returns a pointer to a copy the SparseDNSReportsList.
func (o SparseDNSReportsList) Copy() elemental.Identifiables {

	copy := append(SparseDNSReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseDNSReportsList.
func (o SparseDNSReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseDNSReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseDNSReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseDNSReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseDNSReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseDNSReportsList converted to DNSReportsList.
func (o SparseDNSReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseDNSReportsList) Version() int {

	return 1
}

// SparseDNSReport represents the sparse version of a dnsreport.
type SparseDNSReport struct {
	// Number of times the client saw this activity.
	Count *int `json:"count,omitempty" msgpack:"count,omitempty" bson:"-" mapstructure:"count,omitempty"`

	// If the result is false, error reports the reason of the dns failure.
	Error *string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// name looked up by PU.
	NameLookup *string `json:"nameLookup,omitempty" msgpack:"nameLookup,omitempty" bson:"-" mapstructure:"nameLookup,omitempty"`

	// Result reports whether dns request succeeded or failed.
	Result *bool `json:"result,omitempty" msgpack:"result,omitempty" bson:"-" mapstructure:"result,omitempty"`

	// ID of the source.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"-" mapstructure:"sourceID,omitempty"`

	// Type of the source.
	SourceIP *string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Namespace of the source.
	SourceNamespace *string `json:"sourceNamespace,omitempty" msgpack:"sourceNamespace,omitempty" bson:"-" mapstructure:"sourceNamespace,omitempty"`

	// Time and date of the log.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseDNSReport returns a new  SparseDNSReport.
func NewSparseDNSReport() *SparseDNSReport {
	return &SparseDNSReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseDNSReport) Identity() elemental.Identity {

	return DNSReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseDNSReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseDNSReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseDNSReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseDNSReport) ToPlain() elemental.PlainIdentifiable {

	out := NewDNSReport()
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.NameLookup != nil {
		out.NameLookup = *o.NameLookup
	}
	if o.Result != nil {
		out.Result = *o.Result
	}
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
	}
	if o.SourceNamespace != nil {
		out.SourceNamespace = *o.SourceNamespace
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}

	return out
}

// DeepCopy returns a deep copy if the SparseDNSReport.
func (o *SparseDNSReport) DeepCopy() *SparseDNSReport {

	if o == nil {
		return nil
	}

	out := &SparseDNSReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseDNSReport.
func (o *SparseDNSReport) DeepCopyInto(out *SparseDNSReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseDNSReport: %s", err))
	}

	*out = *target.(*SparseDNSReport)
}
