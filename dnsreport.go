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
	// ID of the enforcer.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"-" mapstructure:"enforcerID,omitempty"`

	// If the result is false, error reports the reason of the dns failure.
	Error string `json:"error" msgpack:"error" bson:"-" mapstructure:"error,omitempty"`

	// name looked up by PU.
	NameLookup string `json:"nameLookup" msgpack:"nameLookup" bson:"-" mapstructure:"nameLookup,omitempty"`

	// Namespace of the enforcer.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// ID of the PU.
	ProcessingUnitID string `json:"processingUnitID" msgpack:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the PU.
	ProcessingUnitNamespace string `json:"processingUnitNamespace" msgpack:"processingUnitNamespace" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Type of the source.
	SourceIP string `json:"sourceIP" msgpack:"sourceIP" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Result reports whether dns request succeeded or failed.
	Success bool `json:"success" msgpack:"success" bson:"-" mapstructure:"success,omitempty"`

	// Time and date of the log.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	// Number of times the client saw this activity.
	Value int `json:"value" msgpack:"value" bson:"-" mapstructure:"value,omitempty"`

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

	return `` + "`" + `dnsreport` + "`" + ` is used to report all the dns lookups that are happening on
behalf of the PU. If the DNS server is on the standard udp port 53 then
enforcer is able to proxy the dns traffic and report the dns lookups. It
also reports whether the lookup was successful or not.`
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
			EnforcerID:              &o.EnforcerID,
			Error:                   &o.Error,
			NameLookup:              &o.NameLookup,
			Namespace:               &o.Namespace,
			ProcessingUnitID:        &o.ProcessingUnitID,
			ProcessingUnitNamespace: &o.ProcessingUnitNamespace,
			SourceIP:                &o.SourceIP,
			Success:                 &o.Success,
			Timestamp:               &o.Timestamp,
			Value:                   &o.Value,
		}
	}

	sp := &SparseDNSReport{}
	for _, f := range fields {
		switch f {
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "error":
			sp.Error = &(o.Error)
		case "nameLookup":
			sp.NameLookup = &(o.NameLookup)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "success":
			sp.Success = &(o.Success)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "value":
			sp.Value = &(o.Value)
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
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.Error != nil {
		o.Error = *so.Error
	}
	if so.NameLookup != nil {
		o.NameLookup = *so.NameLookup
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
	}
	if so.Success != nil {
		o.Success = *so.Success
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.Value != nil {
		o.Value = *so.Value
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

	if err := elemental.ValidateRequiredString("nameLookup", o.NameLookup); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitID", o.ProcessingUnitID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitNamespace", o.ProcessingUnitNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceIP", o.SourceIP); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("value", o.Value); err != nil {
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
	case "enforcerID":
		return o.EnforcerID
	case "error":
		return o.Error
	case "nameLookup":
		return o.NameLookup
	case "namespace":
		return o.Namespace
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "sourceIP":
		return o.SourceIP
	case "success":
		return o.Success
	case "timestamp":
		return o.Timestamp
	case "value":
		return o.Value
	}

	return nil
}

// DNSReportAttributesMap represents the map of attribute for DNSReport.
var DNSReportAttributesMap = map[string]elemental.AttributeSpecification{
	"EnforcerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Type:           "string",
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
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"ProcessingUnitID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the PU.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Type:           "string",
	},
	"ProcessingUnitNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the PU.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Required:       true,
		Type:           "string",
	},
	"SourceIP": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Required:       true,
		Type:           "string",
	},
	"Success": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Success",
		Description:    `Result reports whether dns request succeeded or failed.`,
		Exposed:        true,
		Name:           "success",
		Required:       true,
		Type:           "boolean",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the log.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"Value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of times the client saw this activity.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Type:           "integer",
	},
}

// DNSReportLowerCaseAttributesMap represents the map of attribute for DNSReport.
var DNSReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"enforcerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Type:           "string",
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
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"processingunitid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `ID of the PU.`,
		Exposed:        true,
		Name:           "processingUnitID",
		Required:       true,
		Type:           "string",
	},
	"processingunitnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the PU.`,
		Exposed:        true,
		Name:           "processingUnitNamespace",
		Required:       true,
		Type:           "string",
	},
	"sourceip": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Required:       true,
		Type:           "string",
	},
	"success": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Success",
		Description:    `Result reports whether dns request succeeded or failed.`,
		Exposed:        true,
		Name:           "success",
		Required:       true,
		Type:           "boolean",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the log.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of times the client saw this activity.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Type:           "integer",
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
	// ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"-" mapstructure:"enforcerID,omitempty"`

	// If the result is false, error reports the reason of the dns failure.
	Error *string `json:"error,omitempty" msgpack:"error,omitempty" bson:"-" mapstructure:"error,omitempty"`

	// name looked up by PU.
	NameLookup *string `json:"nameLookup,omitempty" msgpack:"nameLookup,omitempty" bson:"-" mapstructure:"nameLookup,omitempty"`

	// Namespace of the enforcer.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// ID of the PU.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the PU.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Type of the source.
	SourceIP *string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Result reports whether dns request succeeded or failed.
	Success *bool `json:"success,omitempty" msgpack:"success,omitempty" bson:"-" mapstructure:"success,omitempty"`

	// Time and date of the log.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	// Number of times the client saw this activity.
	Value *int `json:"value,omitempty" msgpack:"value,omitempty" bson:"-" mapstructure:"value,omitempty"`

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
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.Error != nil {
		out.Error = *o.Error
	}
	if o.NameLookup != nil {
		out.NameLookup = *o.NameLookup
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
	}
	if o.Success != nil {
		out.Success = *o.Success
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.Value != nil {
		out.Value = *o.Value
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
