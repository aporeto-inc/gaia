package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PingReportIdentity represents the Identity of the object.
var PingReportIdentity = elemental.Identity{
	Name:     "pingreport",
	Category: "pingreports",
	Package:  "zack",
	Private:  false,
}

// PingReportsList represents a list of PingReports
type PingReportsList []*PingReport

// Identity returns the identity of the objects in the list.
func (o PingReportsList) Identity() elemental.Identity {

	return PingReportIdentity
}

// Copy returns a pointer to a copy the PingReportsList.
func (o PingReportsList) Copy() elemental.Identifiables {

	copy := append(PingReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PingReportsList.
func (o PingReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PingReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PingReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PingReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PingReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PingReportsList converted to SparsePingReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PingReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePingReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePingReport)
	}

	return out
}

// Version returns the version of the content.
func (o PingReportsList) Version() int {

	return 1
}

// PingReport represents the model of a pingreport
type PingReport struct {
	// ID unique to a single origin and reply report.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// ID of the destination PU.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"-" mapstructure:"destinationID,omitempty"`

	// Namespace of the destination PU.
	DestinationNamespace string `json:"destinationNamespace" msgpack:"destinationNamespace" bson:"-" mapstructure:"destinationNamespace,omitempty"`

	// ID of the enforcer.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"-" mapstructure:"enforcerNamespace,omitempty"`

	// Semantic version of the enforcer.
	EnforcerVersion string `json:"enforcerVersion" msgpack:"enforcerVersion" bson:"-" mapstructure:"enforcerVersion,omitempty"`

	// Flow tuple in the format <sip:dip:spt:dpt>.
	FlowTuple string `json:"flowTuple" msgpack:"flowTuple" bson:"-" mapstructure:"flowTuple,omitempty"`

	// Time taken for a single request to complete.
	Latency string `json:"latency" msgpack:"latency" bson:"-" mapstructure:"latency,omitempty"`

	// Size of the payload attached to the packet.
	PayloadSize int `json:"payloadSize" msgpack:"payloadSize" bson:"-" mapstructure:"payloadSize,omitempty"`

	// Represents the ping type used.
	PingType string `json:"pingType" msgpack:"pingType" bson:"-" mapstructure:"pingType,omitempty"`

	// Protocol used for the communication.
	Protocol int `json:"protocol" msgpack:"protocol" bson:"-" mapstructure:"protocol,omitempty"`

	// Request represents the request number.
	Request int `json:"request" msgpack:"request" bson:"-" mapstructure:"request,omitempty"`

	// Type of the service.
	ServiceType string `json:"serviceType" msgpack:"serviceType" bson:"-" mapstructure:"serviceType,omitempty"`

	// ID of the source PU.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"-" mapstructure:"sourceID,omitempty"`

	// Namespace of the source PU.
	SourceNamespace string `json:"sourceNamespace" msgpack:"sourceNamespace" bson:"-" mapstructure:"sourceNamespace,omitempty"`

	// Stage when the packet is received.
	Stage string `json:"stage" msgpack:"stage" bson:"-" mapstructure:"stage,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPingReport returns a new *PingReport
func NewPingReport() *PingReport {

	return &PingReport{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *PingReport) Identity() elemental.Identity {

	return PingReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PingReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PingReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPingReport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPingReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PingReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PingReport) BleveType() string {

	return "pingreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PingReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PingReport) Doc() string {

	return `Post a new pu diagnostics report.`
}

func (o *PingReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PingReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePingReport{
			ID:                   &o.ID,
			DestinationID:        &o.DestinationID,
			DestinationNamespace: &o.DestinationNamespace,
			EnforcerID:           &o.EnforcerID,
			EnforcerNamespace:    &o.EnforcerNamespace,
			EnforcerVersion:      &o.EnforcerVersion,
			FlowTuple:            &o.FlowTuple,
			Latency:              &o.Latency,
			PayloadSize:          &o.PayloadSize,
			PingType:             &o.PingType,
			Protocol:             &o.Protocol,
			Request:              &o.Request,
			ServiceType:          &o.ServiceType,
			SourceID:             &o.SourceID,
			SourceNamespace:      &o.SourceNamespace,
			Stage:                &o.Stage,
			Timestamp:            &o.Timestamp,
		}
	}

	sp := &SparsePingReport{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "destinationID":
			sp.DestinationID = &(o.DestinationID)
		case "destinationNamespace":
			sp.DestinationNamespace = &(o.DestinationNamespace)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "enforcerVersion":
			sp.EnforcerVersion = &(o.EnforcerVersion)
		case "flowTuple":
			sp.FlowTuple = &(o.FlowTuple)
		case "latency":
			sp.Latency = &(o.Latency)
		case "payloadSize":
			sp.PayloadSize = &(o.PayloadSize)
		case "pingType":
			sp.PingType = &(o.PingType)
		case "protocol":
			sp.Protocol = &(o.Protocol)
		case "request":
			sp.Request = &(o.Request)
		case "serviceType":
			sp.ServiceType = &(o.ServiceType)
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "sourceNamespace":
			sp.SourceNamespace = &(o.SourceNamespace)
		case "stage":
			sp.Stage = &(o.Stage)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePingReport to the object.
func (o *PingReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePingReport)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.DestinationID != nil {
		o.DestinationID = *so.DestinationID
	}
	if so.DestinationNamespace != nil {
		o.DestinationNamespace = *so.DestinationNamespace
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.EnforcerVersion != nil {
		o.EnforcerVersion = *so.EnforcerVersion
	}
	if so.FlowTuple != nil {
		o.FlowTuple = *so.FlowTuple
	}
	if so.Latency != nil {
		o.Latency = *so.Latency
	}
	if so.PayloadSize != nil {
		o.PayloadSize = *so.PayloadSize
	}
	if so.PingType != nil {
		o.PingType = *so.PingType
	}
	if so.Protocol != nil {
		o.Protocol = *so.Protocol
	}
	if so.Request != nil {
		o.Request = *so.Request
	}
	if so.ServiceType != nil {
		o.ServiceType = *so.ServiceType
	}
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
	}
	if so.SourceNamespace != nil {
		o.SourceNamespace = *so.SourceNamespace
	}
	if so.Stage != nil {
		o.Stage = *so.Stage
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
}

// DeepCopy returns a deep copy if the PingReport.
func (o *PingReport) DeepCopy() *PingReport {

	if o == nil {
		return nil
	}

	out := &PingReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PingReport.
func (o *PingReport) DeepCopyInto(out *PingReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PingReport: %s", err))
	}

	*out = *target.(*PingReport)
}

// Validate valides the current information stored into the structure.
func (o *PingReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceID", o.SourceID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceNamespace", o.SourceNamespace); err != nil {
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
func (*PingReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PingReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PingReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PingReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PingReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PingReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "ID":
		return o.ID
	case "destinationID":
		return o.DestinationID
	case "destinationNamespace":
		return o.DestinationNamespace
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "enforcerVersion":
		return o.EnforcerVersion
	case "flowTuple":
		return o.FlowTuple
	case "latency":
		return o.Latency
	case "payloadSize":
		return o.PayloadSize
	case "pingType":
		return o.PingType
	case "protocol":
		return o.Protocol
	case "request":
		return o.Request
	case "serviceType":
		return o.ServiceType
	case "sourceID":
		return o.SourceID
	case "sourceNamespace":
		return o.SourceNamespace
	case "stage":
		return o.Stage
	case "timestamp":
		return o.Timestamp
	}

	return nil
}

// PingReportAttributesMap represents the map of attribute for PingReport.
var PingReportAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID unique to a single origin and reply report.`,
		Exposed:        true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
	"DestinationID": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination PU.`,
		Exposed:        true,
		Name:           "destinationID",
		Type:           "string",
	},
	"DestinationNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationNamespace",
		Description:    `Namespace of the destination PU.`,
		Exposed:        true,
		Name:           "destinationNamespace",
		Type:           "string",
	},
	"EnforcerID": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Type:           "string",
	},
	"EnforcerNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Type:           "string",
	},
	"EnforcerVersion": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerVersion",
		Description:    `Semantic version of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerVersion",
		Type:           "string",
	},
	"FlowTuple": {
		AllowedChoices: []string{},
		ConvertedName:  "FlowTuple",
		Description:    `Flow tuple in the format <sip:dip:spt:dpt>.`,
		Exposed:        true,
		Name:           "flowTuple",
		Type:           "string",
	},
	"Latency": {
		AllowedChoices: []string{},
		ConvertedName:  "Latency",
		Description:    `Time taken for a single request to complete.`,
		Exposed:        true,
		Name:           "latency",
		Type:           "string",
	},
	"PayloadSize": {
		AllowedChoices: []string{},
		ConvertedName:  "PayloadSize",
		Description:    `Size of the payload attached to the packet.`,
		Exposed:        true,
		Name:           "payloadSize",
		Type:           "integer",
	},
	"PingType": {
		AllowedChoices: []string{},
		ConvertedName:  "PingType",
		Description:    `Represents the ping type used.`,
		Exposed:        true,
		Name:           "pingType",
		Type:           "string",
	},
	"Protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol used for the communication.`,
		Exposed:        true,
		Name:           "protocol",
		Type:           "integer",
	},
	"Request": {
		AllowedChoices: []string{},
		ConvertedName:  "Request",
		Description:    `Request represents the request number.`,
		Exposed:        true,
		Name:           "request",
		Type:           "integer",
	},
	"ServiceType": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceType",
		Description:    `Type of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Type:           "string",
	},
	"SourceID": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source PU.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
	"SourceNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Description:    `Namespace of the source PU.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Required:       true,
		Type:           "string",
	},
	"Stage": {
		AllowedChoices: []string{},
		ConvertedName:  "Stage",
		Description:    `Stage when the packet is received.`,
		Exposed:        true,
		Name:           "stage",
		Type:           "string",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
}

// PingReportLowerCaseAttributesMap represents the map of attribute for PingReport.
var PingReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID unique to a single origin and reply report.`,
		Exposed:        true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
	"destinationid": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination PU.`,
		Exposed:        true,
		Name:           "destinationID",
		Type:           "string",
	},
	"destinationnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationNamespace",
		Description:    `Namespace of the destination PU.`,
		Exposed:        true,
		Name:           "destinationNamespace",
		Type:           "string",
	},
	"enforcerid": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Type:           "string",
	},
	"enforcernamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Type:           "string",
	},
	"enforcerversion": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerVersion",
		Description:    `Semantic version of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerVersion",
		Type:           "string",
	},
	"flowtuple": {
		AllowedChoices: []string{},
		ConvertedName:  "FlowTuple",
		Description:    `Flow tuple in the format <sip:dip:spt:dpt>.`,
		Exposed:        true,
		Name:           "flowTuple",
		Type:           "string",
	},
	"latency": {
		AllowedChoices: []string{},
		ConvertedName:  "Latency",
		Description:    `Time taken for a single request to complete.`,
		Exposed:        true,
		Name:           "latency",
		Type:           "string",
	},
	"payloadsize": {
		AllowedChoices: []string{},
		ConvertedName:  "PayloadSize",
		Description:    `Size of the payload attached to the packet.`,
		Exposed:        true,
		Name:           "payloadSize",
		Type:           "integer",
	},
	"pingtype": {
		AllowedChoices: []string{},
		ConvertedName:  "PingType",
		Description:    `Represents the ping type used.`,
		Exposed:        true,
		Name:           "pingType",
		Type:           "string",
	},
	"protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol used for the communication.`,
		Exposed:        true,
		Name:           "protocol",
		Type:           "integer",
	},
	"request": {
		AllowedChoices: []string{},
		ConvertedName:  "Request",
		Description:    `Request represents the request number.`,
		Exposed:        true,
		Name:           "request",
		Type:           "integer",
	},
	"servicetype": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceType",
		Description:    `Type of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Type:           "string",
	},
	"sourceid": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source PU.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
	"sourcenamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Description:    `Namespace of the source PU.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Required:       true,
		Type:           "string",
	},
	"stage": {
		AllowedChoices: []string{},
		ConvertedName:  "Stage",
		Description:    `Stage when the packet is received.`,
		Exposed:        true,
		Name:           "stage",
		Type:           "string",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
}

// SparsePingReportsList represents a list of SparsePingReports
type SparsePingReportsList []*SparsePingReport

// Identity returns the identity of the objects in the list.
func (o SparsePingReportsList) Identity() elemental.Identity {

	return PingReportIdentity
}

// Copy returns a pointer to a copy the SparsePingReportsList.
func (o SparsePingReportsList) Copy() elemental.Identifiables {

	copy := append(SparsePingReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePingReportsList.
func (o SparsePingReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePingReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePingReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePingReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePingReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePingReportsList converted to PingReportsList.
func (o SparsePingReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePingReportsList) Version() int {

	return 1
}

// SparsePingReport represents the sparse version of a pingreport.
type SparsePingReport struct {
	// ID unique to a single origin and reply report.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// ID of the destination PU.
	DestinationID *string `json:"destinationID,omitempty" msgpack:"destinationID,omitempty" bson:"-" mapstructure:"destinationID,omitempty"`

	// Namespace of the destination PU.
	DestinationNamespace *string `json:"destinationNamespace,omitempty" msgpack:"destinationNamespace,omitempty" bson:"-" mapstructure:"destinationNamespace,omitempty"`

	// ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"-" mapstructure:"enforcerNamespace,omitempty"`

	// Semantic version of the enforcer.
	EnforcerVersion *string `json:"enforcerVersion,omitempty" msgpack:"enforcerVersion,omitempty" bson:"-" mapstructure:"enforcerVersion,omitempty"`

	// Flow tuple in the format <sip:dip:spt:dpt>.
	FlowTuple *string `json:"flowTuple,omitempty" msgpack:"flowTuple,omitempty" bson:"-" mapstructure:"flowTuple,omitempty"`

	// Time taken for a single request to complete.
	Latency *string `json:"latency,omitempty" msgpack:"latency,omitempty" bson:"-" mapstructure:"latency,omitempty"`

	// Size of the payload attached to the packet.
	PayloadSize *int `json:"payloadSize,omitempty" msgpack:"payloadSize,omitempty" bson:"-" mapstructure:"payloadSize,omitempty"`

	// Represents the ping type used.
	PingType *string `json:"pingType,omitempty" msgpack:"pingType,omitempty" bson:"-" mapstructure:"pingType,omitempty"`

	// Protocol used for the communication.
	Protocol *int `json:"protocol,omitempty" msgpack:"protocol,omitempty" bson:"-" mapstructure:"protocol,omitempty"`

	// Request represents the request number.
	Request *int `json:"request,omitempty" msgpack:"request,omitempty" bson:"-" mapstructure:"request,omitempty"`

	// Type of the service.
	ServiceType *string `json:"serviceType,omitempty" msgpack:"serviceType,omitempty" bson:"-" mapstructure:"serviceType,omitempty"`

	// ID of the source PU.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"-" mapstructure:"sourceID,omitempty"`

	// Namespace of the source PU.
	SourceNamespace *string `json:"sourceNamespace,omitempty" msgpack:"sourceNamespace,omitempty" bson:"-" mapstructure:"sourceNamespace,omitempty"`

	// Stage when the packet is received.
	Stage *string `json:"stage,omitempty" msgpack:"stage,omitempty" bson:"-" mapstructure:"stage,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePingReport returns a new  SparsePingReport.
func NewSparsePingReport() *SparsePingReport {
	return &SparsePingReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePingReport) Identity() elemental.Identity {

	return PingReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePingReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePingReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePingReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePingReport{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePingReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePingReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePingReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePingReport) ToPlain() elemental.PlainIdentifiable {

	out := NewPingReport()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.DestinationID != nil {
		out.DestinationID = *o.DestinationID
	}
	if o.DestinationNamespace != nil {
		out.DestinationNamespace = *o.DestinationNamespace
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.EnforcerVersion != nil {
		out.EnforcerVersion = *o.EnforcerVersion
	}
	if o.FlowTuple != nil {
		out.FlowTuple = *o.FlowTuple
	}
	if o.Latency != nil {
		out.Latency = *o.Latency
	}
	if o.PayloadSize != nil {
		out.PayloadSize = *o.PayloadSize
	}
	if o.PingType != nil {
		out.PingType = *o.PingType
	}
	if o.Protocol != nil {
		out.Protocol = *o.Protocol
	}
	if o.Request != nil {
		out.Request = *o.Request
	}
	if o.ServiceType != nil {
		out.ServiceType = *o.ServiceType
	}
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
	}
	if o.SourceNamespace != nil {
		out.SourceNamespace = *o.SourceNamespace
	}
	if o.Stage != nil {
		out.Stage = *o.Stage
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePingReport.
func (o *SparsePingReport) DeepCopy() *SparsePingReport {

	if o == nil {
		return nil
	}

	out := &SparsePingReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePingReport.
func (o *SparsePingReport) DeepCopyInto(out *SparsePingReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePingReport: %s", err))
	}

	*out = *target.(*SparsePingReport)
}

type mongoAttributesPingReport struct {
}
type mongoAttributesSparsePingReport struct {
}
