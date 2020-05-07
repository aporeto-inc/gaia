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
	// Time taken for a single request-response to complete.
	RTT string `json:"RTT" msgpack:"RTT" bson:"-" mapstructure:"RTT,omitempty"`

	// Controller of the transmitter.
	TXController string `json:"TXController" msgpack:"TXController" bson:"-" mapstructure:"TXController,omitempty"`

	// Type of the transmitter.
	TXType string `json:"TXType" msgpack:"TXType" bson:"-" mapstructure:"TXType,omitempty"`

	// If true, application responded to the request.
	ApplicationListening bool `json:"applicationListening" msgpack:"applicationListening" bson:"-" mapstructure:"applicationListening,omitempty"`

	// ID of the destination processing unit.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"-" mapstructure:"destinationID,omitempty"`

	// ID of the enforcer.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"-" mapstructure:"enforcerNamespace,omitempty"`

	// Semantic version of the enforcer.
	EnforcerVersion string `json:"enforcerVersion" msgpack:"enforcerVersion" bson:"-" mapstructure:"enforcerVersion,omitempty"`

	// Four tuple in the format <sip:dip:spt:dpt>.
	FourTuple string `json:"fourTuple" msgpack:"fourTuple" bson:"-" mapstructure:"fourTuple,omitempty"`

	// IterationID unique to a single ping request-response.
	IterationID string `json:"iterationID" msgpack:"iterationID" bson:"-" mapstructure:"iterationID,omitempty"`

	// Namespace of the reporting processing unit.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Size of the payload attached to the packet.
	PayloadSize int `json:"payloadSize" msgpack:"payloadSize" bson:"-" mapstructure:"payloadSize,omitempty"`

	// PingID unique to a single ping control.
	PingID string `json:"pingID" msgpack:"pingID" bson:"-" mapstructure:"pingID,omitempty"`

	// Action of the policy.
	PolicyAction string `json:"policyAction" msgpack:"policyAction" bson:"-" mapstructure:"policyAction,omitempty"`

	// ID of the policy.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"-" mapstructure:"policyID,omitempty"`

	// Protocol used for the communication.
	Protocol int `json:"protocol" msgpack:"protocol" bson:"-" mapstructure:"protocol,omitempty"`

	// Sequence number of the TCP packet. number.
	SeqNum int `json:"seqNum" msgpack:"seqNum" bson:"-" mapstructure:"seqNum,omitempty"`

	// Type of the service.
	ServiceType string `json:"serviceType" msgpack:"serviceType" bson:"-" mapstructure:"serviceType,omitempty"`

	// ID of the source PU.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"-" mapstructure:"sourceID,omitempty"`

	// Current stage when this report has been generated.
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
			RTT:                  &o.RTT,
			TXController:         &o.TXController,
			TXType:               &o.TXType,
			ApplicationListening: &o.ApplicationListening,
			DestinationID:        &o.DestinationID,
			EnforcerID:           &o.EnforcerID,
			EnforcerNamespace:    &o.EnforcerNamespace,
			EnforcerVersion:      &o.EnforcerVersion,
			FourTuple:            &o.FourTuple,
			IterationID:          &o.IterationID,
			Namespace:            &o.Namespace,
			PayloadSize:          &o.PayloadSize,
			PingID:               &o.PingID,
			PolicyAction:         &o.PolicyAction,
			PolicyID:             &o.PolicyID,
			Protocol:             &o.Protocol,
			SeqNum:               &o.SeqNum,
			ServiceType:          &o.ServiceType,
			SourceID:             &o.SourceID,
			Stage:                &o.Stage,
			Timestamp:            &o.Timestamp,
		}
	}

	sp := &SparsePingReport{}
	for _, f := range fields {
		switch f {
		case "RTT":
			sp.RTT = &(o.RTT)
		case "TXController":
			sp.TXController = &(o.TXController)
		case "TXType":
			sp.TXType = &(o.TXType)
		case "applicationListening":
			sp.ApplicationListening = &(o.ApplicationListening)
		case "destinationID":
			sp.DestinationID = &(o.DestinationID)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "enforcerVersion":
			sp.EnforcerVersion = &(o.EnforcerVersion)
		case "fourTuple":
			sp.FourTuple = &(o.FourTuple)
		case "iterationID":
			sp.IterationID = &(o.IterationID)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "payloadSize":
			sp.PayloadSize = &(o.PayloadSize)
		case "pingID":
			sp.PingID = &(o.PingID)
		case "policyAction":
			sp.PolicyAction = &(o.PolicyAction)
		case "policyID":
			sp.PolicyID = &(o.PolicyID)
		case "protocol":
			sp.Protocol = &(o.Protocol)
		case "seqNum":
			sp.SeqNum = &(o.SeqNum)
		case "serviceType":
			sp.ServiceType = &(o.ServiceType)
		case "sourceID":
			sp.SourceID = &(o.SourceID)
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
	if so.RTT != nil {
		o.RTT = *so.RTT
	}
	if so.TXController != nil {
		o.TXController = *so.TXController
	}
	if so.TXType != nil {
		o.TXType = *so.TXType
	}
	if so.ApplicationListening != nil {
		o.ApplicationListening = *so.ApplicationListening
	}
	if so.DestinationID != nil {
		o.DestinationID = *so.DestinationID
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
	if so.FourTuple != nil {
		o.FourTuple = *so.FourTuple
	}
	if so.IterationID != nil {
		o.IterationID = *so.IterationID
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.PayloadSize != nil {
		o.PayloadSize = *so.PayloadSize
	}
	if so.PingID != nil {
		o.PingID = *so.PingID
	}
	if so.PolicyAction != nil {
		o.PolicyAction = *so.PolicyAction
	}
	if so.PolicyID != nil {
		o.PolicyID = *so.PolicyID
	}
	if so.Protocol != nil {
		o.Protocol = *so.Protocol
	}
	if so.SeqNum != nil {
		o.SeqNum = *so.SeqNum
	}
	if so.ServiceType != nil {
		o.ServiceType = *so.ServiceType
	}
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
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

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("pingID", o.PingID); err != nil {
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
	case "RTT":
		return o.RTT
	case "TXController":
		return o.TXController
	case "TXType":
		return o.TXType
	case "applicationListening":
		return o.ApplicationListening
	case "destinationID":
		return o.DestinationID
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "enforcerVersion":
		return o.EnforcerVersion
	case "fourTuple":
		return o.FourTuple
	case "iterationID":
		return o.IterationID
	case "namespace":
		return o.Namespace
	case "payloadSize":
		return o.PayloadSize
	case "pingID":
		return o.PingID
	case "policyAction":
		return o.PolicyAction
	case "policyID":
		return o.PolicyID
	case "protocol":
		return o.Protocol
	case "seqNum":
		return o.SeqNum
	case "serviceType":
		return o.ServiceType
	case "sourceID":
		return o.SourceID
	case "stage":
		return o.Stage
	case "timestamp":
		return o.Timestamp
	}

	return nil
}

// PingReportAttributesMap represents the map of attribute for PingReport.
var PingReportAttributesMap = map[string]elemental.AttributeSpecification{
	"RTT": {
		AllowedChoices: []string{},
		ConvertedName:  "RTT",
		Description:    `Time taken for a single request-response to complete.`,
		Exposed:        true,
		Name:           "RTT",
		Type:           "string",
	},
	"TXController": {
		AllowedChoices: []string{},
		ConvertedName:  "TXController",
		Description:    `Controller of the transmitter.`,
		Exposed:        true,
		Name:           "TXController",
		Type:           "string",
	},
	"TXType": {
		AllowedChoices: []string{},
		ConvertedName:  "TXType",
		Description:    `Type of the transmitter.`,
		Exposed:        true,
		Name:           "TXType",
		Type:           "string",
	},
	"ApplicationListening": {
		AllowedChoices: []string{},
		ConvertedName:  "ApplicationListening",
		Description:    `If true, application responded to the request.`,
		Exposed:        true,
		Name:           "applicationListening",
		Type:           "boolean",
	},
	"DestinationID": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination processing unit.`,
		Exposed:        true,
		Name:           "destinationID",
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
	"FourTuple": {
		AllowedChoices: []string{},
		ConvertedName:  "FourTuple",
		Description:    `Four tuple in the format <sip:dip:spt:dpt>.`,
		Exposed:        true,
		Name:           "fourTuple",
		Type:           "string",
	},
	"IterationID": {
		AllowedChoices: []string{},
		ConvertedName:  "IterationID",
		Description:    `IterationID unique to a single ping request-response.`,
		Exposed:        true,
		Name:           "iterationID",
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the reporting processing unit.`,
		Exposed:        true,
		Name:           "namespace",
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
	"PingID": {
		AllowedChoices: []string{},
		ConvertedName:  "PingID",
		Description:    `PingID unique to a single ping control.`,
		Exposed:        true,
		Name:           "pingID",
		Required:       true,
		Type:           "string",
	},
	"PolicyAction": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyAction",
		Description:    `Action of the policy.`,
		Exposed:        true,
		Name:           "policyAction",
		Type:           "string",
	},
	"PolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the policy.`,
		Exposed:        true,
		Name:           "policyID",
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
	"SeqNum": {
		AllowedChoices: []string{},
		ConvertedName:  "SeqNum",
		Description:    `Sequence number of the TCP packet. number.`,
		Exposed:        true,
		Name:           "seqNum",
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
		Type:           "string",
	},
	"Stage": {
		AllowedChoices: []string{},
		ConvertedName:  "Stage",
		Description:    `Current stage when this report has been generated.`,
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
	"rtt": {
		AllowedChoices: []string{},
		ConvertedName:  "RTT",
		Description:    `Time taken for a single request-response to complete.`,
		Exposed:        true,
		Name:           "RTT",
		Type:           "string",
	},
	"txcontroller": {
		AllowedChoices: []string{},
		ConvertedName:  "TXController",
		Description:    `Controller of the transmitter.`,
		Exposed:        true,
		Name:           "TXController",
		Type:           "string",
	},
	"txtype": {
		AllowedChoices: []string{},
		ConvertedName:  "TXType",
		Description:    `Type of the transmitter.`,
		Exposed:        true,
		Name:           "TXType",
		Type:           "string",
	},
	"applicationlistening": {
		AllowedChoices: []string{},
		ConvertedName:  "ApplicationListening",
		Description:    `If true, application responded to the request.`,
		Exposed:        true,
		Name:           "applicationListening",
		Type:           "boolean",
	},
	"destinationid": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination processing unit.`,
		Exposed:        true,
		Name:           "destinationID",
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
	"fourtuple": {
		AllowedChoices: []string{},
		ConvertedName:  "FourTuple",
		Description:    `Four tuple in the format <sip:dip:spt:dpt>.`,
		Exposed:        true,
		Name:           "fourTuple",
		Type:           "string",
	},
	"iterationid": {
		AllowedChoices: []string{},
		ConvertedName:  "IterationID",
		Description:    `IterationID unique to a single ping request-response.`,
		Exposed:        true,
		Name:           "iterationID",
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the reporting processing unit.`,
		Exposed:        true,
		Name:           "namespace",
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
	"pingid": {
		AllowedChoices: []string{},
		ConvertedName:  "PingID",
		Description:    `PingID unique to a single ping control.`,
		Exposed:        true,
		Name:           "pingID",
		Required:       true,
		Type:           "string",
	},
	"policyaction": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyAction",
		Description:    `Action of the policy.`,
		Exposed:        true,
		Name:           "policyAction",
		Type:           "string",
	},
	"policyid": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the policy.`,
		Exposed:        true,
		Name:           "policyID",
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
	"seqnum": {
		AllowedChoices: []string{},
		ConvertedName:  "SeqNum",
		Description:    `Sequence number of the TCP packet. number.`,
		Exposed:        true,
		Name:           "seqNum",
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
		Type:           "string",
	},
	"stage": {
		AllowedChoices: []string{},
		ConvertedName:  "Stage",
		Description:    `Current stage when this report has been generated.`,
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
	// Time taken for a single request-response to complete.
	RTT *string `json:"RTT,omitempty" msgpack:"RTT,omitempty" bson:"-" mapstructure:"RTT,omitempty"`

	// Controller of the transmitter.
	TXController *string `json:"TXController,omitempty" msgpack:"TXController,omitempty" bson:"-" mapstructure:"TXController,omitempty"`

	// Type of the transmitter.
	TXType *string `json:"TXType,omitempty" msgpack:"TXType,omitempty" bson:"-" mapstructure:"TXType,omitempty"`

	// If true, application responded to the request.
	ApplicationListening *bool `json:"applicationListening,omitempty" msgpack:"applicationListening,omitempty" bson:"-" mapstructure:"applicationListening,omitempty"`

	// ID of the destination processing unit.
	DestinationID *string `json:"destinationID,omitempty" msgpack:"destinationID,omitempty" bson:"-" mapstructure:"destinationID,omitempty"`

	// ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"-" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"-" mapstructure:"enforcerNamespace,omitempty"`

	// Semantic version of the enforcer.
	EnforcerVersion *string `json:"enforcerVersion,omitempty" msgpack:"enforcerVersion,omitempty" bson:"-" mapstructure:"enforcerVersion,omitempty"`

	// Four tuple in the format <sip:dip:spt:dpt>.
	FourTuple *string `json:"fourTuple,omitempty" msgpack:"fourTuple,omitempty" bson:"-" mapstructure:"fourTuple,omitempty"`

	// IterationID unique to a single ping request-response.
	IterationID *string `json:"iterationID,omitempty" msgpack:"iterationID,omitempty" bson:"-" mapstructure:"iterationID,omitempty"`

	// Namespace of the reporting processing unit.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// Size of the payload attached to the packet.
	PayloadSize *int `json:"payloadSize,omitempty" msgpack:"payloadSize,omitempty" bson:"-" mapstructure:"payloadSize,omitempty"`

	// PingID unique to a single ping control.
	PingID *string `json:"pingID,omitempty" msgpack:"pingID,omitempty" bson:"-" mapstructure:"pingID,omitempty"`

	// Action of the policy.
	PolicyAction *string `json:"policyAction,omitempty" msgpack:"policyAction,omitempty" bson:"-" mapstructure:"policyAction,omitempty"`

	// ID of the policy.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"-" mapstructure:"policyID,omitempty"`

	// Protocol used for the communication.
	Protocol *int `json:"protocol,omitempty" msgpack:"protocol,omitempty" bson:"-" mapstructure:"protocol,omitempty"`

	// Sequence number of the TCP packet. number.
	SeqNum *int `json:"seqNum,omitempty" msgpack:"seqNum,omitempty" bson:"-" mapstructure:"seqNum,omitempty"`

	// Type of the service.
	ServiceType *string `json:"serviceType,omitempty" msgpack:"serviceType,omitempty" bson:"-" mapstructure:"serviceType,omitempty"`

	// ID of the source PU.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"-" mapstructure:"sourceID,omitempty"`

	// Current stage when this report has been generated.
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
	if o.RTT != nil {
		out.RTT = *o.RTT
	}
	if o.TXController != nil {
		out.TXController = *o.TXController
	}
	if o.TXType != nil {
		out.TXType = *o.TXType
	}
	if o.ApplicationListening != nil {
		out.ApplicationListening = *o.ApplicationListening
	}
	if o.DestinationID != nil {
		out.DestinationID = *o.DestinationID
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
	if o.FourTuple != nil {
		out.FourTuple = *o.FourTuple
	}
	if o.IterationID != nil {
		out.IterationID = *o.IterationID
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.PayloadSize != nil {
		out.PayloadSize = *o.PayloadSize
	}
	if o.PingID != nil {
		out.PingID = *o.PingID
	}
	if o.PolicyAction != nil {
		out.PolicyAction = *o.PolicyAction
	}
	if o.PolicyID != nil {
		out.PolicyID = *o.PolicyID
	}
	if o.Protocol != nil {
		out.Protocol = *o.Protocol
	}
	if o.SeqNum != nil {
		out.SeqNum = *o.SeqNum
	}
	if o.ServiceType != nil {
		out.ServiceType = *o.ServiceType
	}
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
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
