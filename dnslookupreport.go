package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// DNSLookupReportActionValue represents the possible values for attribute "action".
type DNSLookupReportActionValue string

const (
	// DNSLookupReportActionAccept represents the value Accept.
	DNSLookupReportActionAccept DNSLookupReportActionValue = "Accept"

	// DNSLookupReportActionReject represents the value Reject.
	DNSLookupReportActionReject DNSLookupReportActionValue = "Reject"
)

// DNSLookupReportIdentity represents the Identity of the object.
var DNSLookupReportIdentity = elemental.Identity{
	Name:     "dnslookupreport",
	Category: "dnslookupreports",
	Package:  "zack",
	Private:  false,
}

// DNSLookupReportsList represents a list of DNSLookupReports
type DNSLookupReportsList []*DNSLookupReport

// Identity returns the identity of the objects in the list.
func (o DNSLookupReportsList) Identity() elemental.Identity {

	return DNSLookupReportIdentity
}

// Copy returns a pointer to a copy the DNSLookupReportsList.
func (o DNSLookupReportsList) Copy() elemental.Identifiables {

	copy := append(DNSLookupReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the DNSLookupReportsList.
func (o DNSLookupReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(DNSLookupReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*DNSLookupReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o DNSLookupReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o DNSLookupReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the DNSLookupReportsList converted to SparseDNSLookupReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o DNSLookupReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseDNSLookupReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseDNSLookupReport)
	}

	return out
}

// Version returns the version of the content.
func (o DNSLookupReportsList) Version() int {

	return 1
}

// DNSLookupReport represents the model of a dnslookupreport
type DNSLookupReport struct {
	// Action of the DNS request.
	Action DNSLookupReportActionValue `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// ID of the enforcer.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"-" mapstructure:"enforcerNamespace,omitempty"`

	// ID of the PU.
	ProcessingUnitID string `json:"processingUnitID" msgpack:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the PU.
	ProcessingUnitNamespace string `json:"processingUnitNamespace" msgpack:"processingUnitNamespace" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// This field is only set when the lookup fails. It specifies the reason for the
	// failure.
	Reason string `json:"reason" msgpack:"reason" bson:"-" mapstructure:"reason,omitempty"`

	// name used for DNS resolution.
	ResolvedName string `json:"resolvedName" msgpack:"resolvedName" bson:"-" mapstructure:"resolvedName,omitempty"`

	// Type of the source.
	SourceIP string `json:"sourceIP" msgpack:"sourceIP" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Time and date of the log.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	// Number of times the client saw this activity.
	Value int `json:"value" msgpack:"value" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewDNSLookupReport returns a new *DNSLookupReport
func NewDNSLookupReport() *DNSLookupReport {

	return &DNSLookupReport{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *DNSLookupReport) Identity() elemental.Identity {

	return DNSLookupReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *DNSLookupReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *DNSLookupReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DNSLookupReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesDNSLookupReport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *DNSLookupReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesDNSLookupReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *DNSLookupReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *DNSLookupReport) BleveType() string {

	return "dnslookupreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *DNSLookupReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *DNSLookupReport) Doc() string {

	return `A DNS Lookup report is used to report a DNS lookup that is happening on
behalf of a processing unit. If the DNS server is on the standard udp port 53
then enforcer is able to proxy the DNS traffic and make a report. The report
indicate whether or not the lookup was successful.`
}

func (o *DNSLookupReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *DNSLookupReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseDNSLookupReport{
			Action:                  &o.Action,
			EnforcerID:              &o.EnforcerID,
			EnforcerNamespace:       &o.EnforcerNamespace,
			ProcessingUnitID:        &o.ProcessingUnitID,
			ProcessingUnitNamespace: &o.ProcessingUnitNamespace,
			Reason:                  &o.Reason,
			ResolvedName:            &o.ResolvedName,
			SourceIP:                &o.SourceIP,
			Timestamp:               &o.Timestamp,
			Value:                   &o.Value,
		}
	}

	sp := &SparseDNSLookupReport{}
	for _, f := range fields {
		switch f {
		case "action":
			sp.Action = &(o.Action)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "reason":
			sp.Reason = &(o.Reason)
		case "resolvedName":
			sp.ResolvedName = &(o.ResolvedName)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "value":
			sp.Value = &(o.Value)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseDNSLookupReport to the object.
func (o *DNSLookupReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseDNSLookupReport)
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.Reason != nil {
		o.Reason = *so.Reason
	}
	if so.ResolvedName != nil {
		o.ResolvedName = *so.ResolvedName
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.Value != nil {
		o.Value = *so.Value
	}
}

// DeepCopy returns a deep copy if the DNSLookupReport.
func (o *DNSLookupReport) DeepCopy() *DNSLookupReport {

	if o == nil {
		return nil
	}

	out := &DNSLookupReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *DNSLookupReport.
func (o *DNSLookupReport) DeepCopyInto(out *DNSLookupReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy DNSLookupReport: %s", err))
	}

	*out = *target.(*DNSLookupReport)
}

// Validate valides the current information stored into the structure.
func (o *DNSLookupReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitID", o.ProcessingUnitID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("processingUnitNamespace", o.ProcessingUnitNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("resolvedName", o.ResolvedName); err != nil {
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
func (*DNSLookupReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := DNSLookupReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return DNSLookupReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*DNSLookupReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return DNSLookupReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *DNSLookupReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "action":
		return o.Action
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "reason":
		return o.Reason
	case "resolvedName":
		return o.ResolvedName
	case "sourceIP":
		return o.SourceIP
	case "timestamp":
		return o.Timestamp
	case "value":
		return o.Value
	}

	return nil
}

// DNSLookupReportAttributesMap represents the map of attribute for DNSLookupReport.
var DNSLookupReportAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action of the DNS request.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"EnforcerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Type:           "string",
	},
	"EnforcerNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
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
	"Reason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Reason",
		Description: `This field is only set when the lookup fails. It specifies the reason for the
failure.`,
		Exposed: true,
		Name:    "reason",
		Type:    "string",
	},
	"ResolvedName": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ResolvedName",
		Description:    `name used for DNS resolution.`,
		Exposed:        true,
		Name:           "resolvedName",
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

// DNSLookupReportLowerCaseAttributesMap represents the map of attribute for DNSLookupReport.
var DNSLookupReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action of the DNS request.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"enforcerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Type:           "string",
	},
	"enforcernamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
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
	"reason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Reason",
		Description: `This field is only set when the lookup fails. It specifies the reason for the
failure.`,
		Exposed: true,
		Name:    "reason",
		Type:    "string",
	},
	"resolvedname": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ResolvedName",
		Description:    `name used for DNS resolution.`,
		Exposed:        true,
		Name:           "resolvedName",
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

// SparseDNSLookupReportsList represents a list of SparseDNSLookupReports
type SparseDNSLookupReportsList []*SparseDNSLookupReport

// Identity returns the identity of the objects in the list.
func (o SparseDNSLookupReportsList) Identity() elemental.Identity {

	return DNSLookupReportIdentity
}

// Copy returns a pointer to a copy the SparseDNSLookupReportsList.
func (o SparseDNSLookupReportsList) Copy() elemental.Identifiables {

	copy := append(SparseDNSLookupReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseDNSLookupReportsList.
func (o SparseDNSLookupReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseDNSLookupReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseDNSLookupReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseDNSLookupReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseDNSLookupReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseDNSLookupReportsList converted to DNSLookupReportsList.
func (o SparseDNSLookupReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseDNSLookupReportsList) Version() int {

	return 1
}

// SparseDNSLookupReport represents the sparse version of a dnslookupreport.
type SparseDNSLookupReport struct {
	// Action of the DNS request.
	Action *DNSLookupReportActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"-" mapstructure:"enforcerNamespace,omitempty"`

	// ID of the PU.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the PU.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// This field is only set when the lookup fails. It specifies the reason for the
	// failure.
	Reason *string `json:"reason,omitempty" msgpack:"reason,omitempty" bson:"-" mapstructure:"reason,omitempty"`

	// name used for DNS resolution.
	ResolvedName *string `json:"resolvedName,omitempty" msgpack:"resolvedName,omitempty" bson:"-" mapstructure:"resolvedName,omitempty"`

	// Type of the source.
	SourceIP *string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Time and date of the log.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	// Number of times the client saw this activity.
	Value *int `json:"value,omitempty" msgpack:"value,omitempty" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseDNSLookupReport returns a new  SparseDNSLookupReport.
func NewSparseDNSLookupReport() *SparseDNSLookupReport {
	return &SparseDNSLookupReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseDNSLookupReport) Identity() elemental.Identity {

	return DNSLookupReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseDNSLookupReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseDNSLookupReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDNSLookupReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseDNSLookupReport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseDNSLookupReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseDNSLookupReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseDNSLookupReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseDNSLookupReport) ToPlain() elemental.PlainIdentifiable {

	out := NewDNSLookupReport()
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.Reason != nil {
		out.Reason = *o.Reason
	}
	if o.ResolvedName != nil {
		out.ResolvedName = *o.ResolvedName
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.Value != nil {
		out.Value = *o.Value
	}

	return out
}

// DeepCopy returns a deep copy if the SparseDNSLookupReport.
func (o *SparseDNSLookupReport) DeepCopy() *SparseDNSLookupReport {

	if o == nil {
		return nil
	}

	out := &SparseDNSLookupReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseDNSLookupReport.
func (o *SparseDNSLookupReport) DeepCopyInto(out *SparseDNSLookupReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseDNSLookupReport: %s", err))
	}

	*out = *target.(*SparseDNSLookupReport)
}

type mongoAttributesDNSLookupReport struct {
}
type mongoAttributesSparseDNSLookupReport struct {
}
