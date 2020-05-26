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
	// Identifier of the object.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Time taken for a single request-response to complete.
	RTT string `json:"RTT" msgpack:"RTT" bson:"rtt" mapstructure:"RTT,omitempty"`

	// Controller of the transmitter.
	TXController string `json:"TXController" msgpack:"TXController" bson:"txcontroller" mapstructure:"TXController,omitempty"`

	// Type of the transmitter.
	TXType string `json:"TXType" msgpack:"TXType" bson:"txtype" mapstructure:"TXType,omitempty"`

	// If true, application responded to the request.
	ApplicationListening bool `json:"applicationListening" msgpack:"applicationListening" bson:"applicationlistening" mapstructure:"applicationListening,omitempty"`

	// ID of the destination processing unit.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"destinationid" mapstructure:"destinationID,omitempty"`

	// ID of the enforcer.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"enforcernamespace" mapstructure:"enforcerNamespace,omitempty"`

	// Semantic version of the enforcer.
	EnforcerVersion string `json:"enforcerVersion" msgpack:"enforcerVersion" bson:"enforcerversion" mapstructure:"enforcerVersion,omitempty"`

	// Four tuple in the format <sip:dip:spt:dpt>.
	FourTuple string `json:"fourTuple" msgpack:"fourTuple" bson:"fourtuple" mapstructure:"fourTuple,omitempty"`

	// IterationID unique to a single ping request-response.
	IterationID string `json:"iterationID" msgpack:"iterationID" bson:"iterationid" mapstructure:"iterationID,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog map[string]string `json:"-" msgpack:"-" bson:"migrationslog" mapstructure:"-,omitempty"`

	// Namespace of the reporting processing unit.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// Size of the payload attached to the packet.
	PayloadSize int `json:"payloadSize" msgpack:"payloadSize" bson:"payloadsize" mapstructure:"payloadSize,omitempty"`

	// PingID unique to a single ping control.
	PingID string `json:"pingID" msgpack:"pingID" bson:"pingid" mapstructure:"pingID,omitempty"`

	// Action of the policy.
	PolicyAction string `json:"policyAction" msgpack:"policyAction" bson:"policyaction" mapstructure:"policyAction,omitempty"`

	// ID of the policy.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"policyid" mapstructure:"policyID,omitempty"`

	// Protocol used for the communication.
	Protocol int `json:"protocol" msgpack:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	// Sequence number of the TCP packet. number.
	SeqNum int `json:"seqNum" msgpack:"seqNum" bson:"seqnum" mapstructure:"seqNum,omitempty"`

	// Type of the service.
	ServiceType string `json:"serviceType" msgpack:"serviceType" bson:"servicetype" mapstructure:"serviceType,omitempty"`

	// ID of the source PU.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"sourceid" mapstructure:"sourceID,omitempty"`

	// Current stage when this report has been generated.
	Stage string `json:"stage" msgpack:"stage" bson:"stage" mapstructure:"stage,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"timestamp" mapstructure:"timestamp,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPingReport returns a new *PingReport
func NewPingReport() *PingReport {

	return &PingReport{
		ModelVersion:  1,
		MigrationsLog: map[string]string{},
	}
}

// Identity returns the Identity of the object.
func (o *PingReport) Identity() elemental.Identity {

	return PingReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PingReport) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PingReport) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPingReport{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.RTT = o.RTT
	s.TXController = o.TXController
	s.TXType = o.TXType
	s.ApplicationListening = o.ApplicationListening
	s.DestinationID = o.DestinationID
	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace
	s.EnforcerVersion = o.EnforcerVersion
	s.FourTuple = o.FourTuple
	s.IterationID = o.IterationID
	s.MigrationsLog = o.MigrationsLog
	s.Namespace = o.Namespace
	s.PayloadSize = o.PayloadSize
	s.PingID = o.PingID
	s.PolicyAction = o.PolicyAction
	s.PolicyID = o.PolicyID
	s.Protocol = o.Protocol
	s.SeqNum = o.SeqNum
	s.ServiceType = o.ServiceType
	s.SourceID = o.SourceID
	s.Stage = o.Stage
	s.Timestamp = o.Timestamp
	s.ZHash = o.ZHash
	s.Zone = o.Zone

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

	o.ID = s.ID.Hex()
	o.RTT = s.RTT
	o.TXController = s.TXController
	o.TXType = s.TXType
	o.ApplicationListening = s.ApplicationListening
	o.DestinationID = s.DestinationID
	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace
	o.EnforcerVersion = s.EnforcerVersion
	o.FourTuple = s.FourTuple
	o.IterationID = s.IterationID
	o.MigrationsLog = s.MigrationsLog
	o.Namespace = s.Namespace
	o.PayloadSize = s.PayloadSize
	o.PingID = s.PingID
	o.PolicyAction = s.PolicyAction
	o.PolicyID = s.PolicyID
	o.Protocol = s.Protocol
	o.SeqNum = s.SeqNum
	o.ServiceType = s.ServiceType
	o.SourceID = s.SourceID
	o.Stage = s.Stage
	o.Timestamp = s.Timestamp
	o.ZHash = s.ZHash
	o.Zone = s.Zone

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

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *PingReport) GetMigrationsLog() map[string]string {

	return o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the given value.
func (o *PingReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *PingReport) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *PingReport) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *PingReport) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *PingReport) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PingReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePingReport{
			ID:                   &o.ID,
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
			MigrationsLog:        &o.MigrationsLog,
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
			ZHash:                &o.ZHash,
			Zone:                 &o.Zone,
		}
	}

	sp := &SparsePingReport{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
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
		case "migrationsLog":
			sp.MigrationsLog = &(o.MigrationsLog)
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
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
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
	if so.MigrationsLog != nil {
		o.MigrationsLog = *so.MigrationsLog
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
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
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
	case "ID":
		return o.ID
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
	case "migrationsLog":
		return o.MigrationsLog
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
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// PingReportAttributesMap represents the map of attribute for PingReport.
var PingReportAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": {
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
	"RTT": {
		AllowedChoices: []string{},
		ConvertedName:  "RTT",
		Description:    `Time taken for a single request-response to complete.`,
		Exposed:        true,
		Name:           "RTT",
		Stored:         true,
		Type:           "string",
	},
	"TXController": {
		AllowedChoices: []string{},
		ConvertedName:  "TXController",
		Description:    `Controller of the transmitter.`,
		Exposed:        true,
		Name:           "TXController",
		Stored:         true,
		Type:           "string",
	},
	"TXType": {
		AllowedChoices: []string{},
		ConvertedName:  "TXType",
		Description:    `Type of the transmitter.`,
		Exposed:        true,
		Name:           "TXType",
		Stored:         true,
		Type:           "string",
	},
	"ApplicationListening": {
		AllowedChoices: []string{},
		ConvertedName:  "ApplicationListening",
		Description:    `If true, application responded to the request.`,
		Exposed:        true,
		Name:           "applicationListening",
		Stored:         true,
		Type:           "boolean",
	},
	"DestinationID": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination processing unit.`,
		Exposed:        true,
		Name:           "destinationID",
		Stored:         true,
		Type:           "string",
	},
	"EnforcerID": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerVersion": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerVersion",
		Description:    `Semantic version of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerVersion",
		Stored:         true,
		Type:           "string",
	},
	"FourTuple": {
		AllowedChoices: []string{},
		ConvertedName:  "FourTuple",
		Description:    `Four tuple in the format <sip:dip:spt:dpt>.`,
		Exposed:        true,
		Name:           "fourTuple",
		Stored:         true,
		Type:           "string",
	},
	"IterationID": {
		AllowedChoices: []string{},
		ConvertedName:  "IterationID",
		Description:    `IterationID unique to a single ping request-response.`,
		Exposed:        true,
		Name:           "iterationID",
		Stored:         true,
		Type:           "string",
	},
	"MigrationsLog": {
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
	"Namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the reporting processing unit.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"PayloadSize": {
		AllowedChoices: []string{},
		ConvertedName:  "PayloadSize",
		Description:    `Size of the payload attached to the packet.`,
		Exposed:        true,
		Name:           "payloadSize",
		Stored:         true,
		Type:           "integer",
	},
	"PingID": {
		AllowedChoices: []string{},
		ConvertedName:  "PingID",
		Description:    `PingID unique to a single ping control.`,
		Exposed:        true,
		Name:           "pingID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"PolicyAction": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyAction",
		Description:    `Action of the policy.`,
		Exposed:        true,
		Name:           "policyAction",
		Stored:         true,
		Type:           "string",
	},
	"PolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the policy.`,
		Exposed:        true,
		Name:           "policyID",
		Stored:         true,
		Type:           "string",
	},
	"Protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol used for the communication.`,
		Exposed:        true,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"SeqNum": {
		AllowedChoices: []string{},
		ConvertedName:  "SeqNum",
		Description:    `Sequence number of the TCP packet. number.`,
		Exposed:        true,
		Name:           "seqNum",
		Stored:         true,
		Type:           "integer",
	},
	"ServiceType": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceType",
		Description:    `Type of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Stored:         true,
		Type:           "string",
	},
	"SourceID": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source PU.`,
		Exposed:        true,
		Name:           "sourceID",
		Stored:         true,
		Type:           "string",
	},
	"Stage": {
		AllowedChoices: []string{},
		ConvertedName:  "Stage",
		Description:    `Current stage when this report has been generated.`,
		Exposed:        true,
		Name:           "stage",
		Stored:         true,
		Type:           "string",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Stored:         true,
		Type:           "time",
	},
	"ZHash": {
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
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
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

// PingReportLowerCaseAttributesMap represents the map of attribute for PingReport.
var PingReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": {
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
	"rtt": {
		AllowedChoices: []string{},
		ConvertedName:  "RTT",
		Description:    `Time taken for a single request-response to complete.`,
		Exposed:        true,
		Name:           "RTT",
		Stored:         true,
		Type:           "string",
	},
	"txcontroller": {
		AllowedChoices: []string{},
		ConvertedName:  "TXController",
		Description:    `Controller of the transmitter.`,
		Exposed:        true,
		Name:           "TXController",
		Stored:         true,
		Type:           "string",
	},
	"txtype": {
		AllowedChoices: []string{},
		ConvertedName:  "TXType",
		Description:    `Type of the transmitter.`,
		Exposed:        true,
		Name:           "TXType",
		Stored:         true,
		Type:           "string",
	},
	"applicationlistening": {
		AllowedChoices: []string{},
		ConvertedName:  "ApplicationListening",
		Description:    `If true, application responded to the request.`,
		Exposed:        true,
		Name:           "applicationListening",
		Stored:         true,
		Type:           "boolean",
	},
	"destinationid": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination processing unit.`,
		Exposed:        true,
		Name:           "destinationID",
		Stored:         true,
		Type:           "string",
	},
	"enforcerid": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `ID of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcerversion": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerVersion",
		Description:    `Semantic version of the enforcer.`,
		Exposed:        true,
		Name:           "enforcerVersion",
		Stored:         true,
		Type:           "string",
	},
	"fourtuple": {
		AllowedChoices: []string{},
		ConvertedName:  "FourTuple",
		Description:    `Four tuple in the format <sip:dip:spt:dpt>.`,
		Exposed:        true,
		Name:           "fourTuple",
		Stored:         true,
		Type:           "string",
	},
	"iterationid": {
		AllowedChoices: []string{},
		ConvertedName:  "IterationID",
		Description:    `IterationID unique to a single ping request-response.`,
		Exposed:        true,
		Name:           "iterationID",
		Stored:         true,
		Type:           "string",
	},
	"migrationslog": {
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
	"namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the reporting processing unit.`,
		Exposed:        true,
		Name:           "namespace",
		Stored:         true,
		Type:           "string",
	},
	"payloadsize": {
		AllowedChoices: []string{},
		ConvertedName:  "PayloadSize",
		Description:    `Size of the payload attached to the packet.`,
		Exposed:        true,
		Name:           "payloadSize",
		Stored:         true,
		Type:           "integer",
	},
	"pingid": {
		AllowedChoices: []string{},
		ConvertedName:  "PingID",
		Description:    `PingID unique to a single ping control.`,
		Exposed:        true,
		Name:           "pingID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"policyaction": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyAction",
		Description:    `Action of the policy.`,
		Exposed:        true,
		Name:           "policyAction",
		Stored:         true,
		Type:           "string",
	},
	"policyid": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the policy.`,
		Exposed:        true,
		Name:           "policyID",
		Stored:         true,
		Type:           "string",
	},
	"protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol used for the communication.`,
		Exposed:        true,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"seqnum": {
		AllowedChoices: []string{},
		ConvertedName:  "SeqNum",
		Description:    `Sequence number of the TCP packet. number.`,
		Exposed:        true,
		Name:           "seqNum",
		Stored:         true,
		Type:           "integer",
	},
	"servicetype": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceType",
		Description:    `Type of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Stored:         true,
		Type:           "string",
	},
	"sourceid": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source PU.`,
		Exposed:        true,
		Name:           "sourceID",
		Stored:         true,
		Type:           "string",
	},
	"stage": {
		AllowedChoices: []string{},
		ConvertedName:  "Stage",
		Description:    `Current stage when this report has been generated.`,
		Exposed:        true,
		Name:           "stage",
		Stored:         true,
		Type:           "string",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Stored:         true,
		Type:           "time",
	},
	"zhash": {
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
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
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
	// Identifier of the object.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// Time taken for a single request-response to complete.
	RTT *string `json:"RTT,omitempty" msgpack:"RTT,omitempty" bson:"rtt,omitempty" mapstructure:"RTT,omitempty"`

	// Controller of the transmitter.
	TXController *string `json:"TXController,omitempty" msgpack:"TXController,omitempty" bson:"txcontroller,omitempty" mapstructure:"TXController,omitempty"`

	// Type of the transmitter.
	TXType *string `json:"TXType,omitempty" msgpack:"TXType,omitempty" bson:"txtype,omitempty" mapstructure:"TXType,omitempty"`

	// If true, application responded to the request.
	ApplicationListening *bool `json:"applicationListening,omitempty" msgpack:"applicationListening,omitempty" bson:"applicationlistening,omitempty" mapstructure:"applicationListening,omitempty"`

	// ID of the destination processing unit.
	DestinationID *string `json:"destinationID,omitempty" msgpack:"destinationID,omitempty" bson:"destinationid,omitempty" mapstructure:"destinationID,omitempty"`

	// ID of the enforcer.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Semantic version of the enforcer.
	EnforcerVersion *string `json:"enforcerVersion,omitempty" msgpack:"enforcerVersion,omitempty" bson:"enforcerversion,omitempty" mapstructure:"enforcerVersion,omitempty"`

	// Four tuple in the format <sip:dip:spt:dpt>.
	FourTuple *string `json:"fourTuple,omitempty" msgpack:"fourTuple,omitempty" bson:"fourtuple,omitempty" mapstructure:"fourTuple,omitempty"`

	// IterationID unique to a single ping request-response.
	IterationID *string `json:"iterationID,omitempty" msgpack:"iterationID,omitempty" bson:"iterationid,omitempty" mapstructure:"iterationID,omitempty"`

	// Internal property maintaining migrations information.
	MigrationsLog *map[string]string `json:"-" msgpack:"-" bson:"migrationslog,omitempty" mapstructure:"-,omitempty"`

	// Namespace of the reporting processing unit.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// Size of the payload attached to the packet.
	PayloadSize *int `json:"payloadSize,omitempty" msgpack:"payloadSize,omitempty" bson:"payloadsize,omitempty" mapstructure:"payloadSize,omitempty"`

	// PingID unique to a single ping control.
	PingID *string `json:"pingID,omitempty" msgpack:"pingID,omitempty" bson:"pingid,omitempty" mapstructure:"pingID,omitempty"`

	// Action of the policy.
	PolicyAction *string `json:"policyAction,omitempty" msgpack:"policyAction,omitempty" bson:"policyaction,omitempty" mapstructure:"policyAction,omitempty"`

	// ID of the policy.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"policyid,omitempty" mapstructure:"policyID,omitempty"`

	// Protocol used for the communication.
	Protocol *int `json:"protocol,omitempty" msgpack:"protocol,omitempty" bson:"protocol,omitempty" mapstructure:"protocol,omitempty"`

	// Sequence number of the TCP packet. number.
	SeqNum *int `json:"seqNum,omitempty" msgpack:"seqNum,omitempty" bson:"seqnum,omitempty" mapstructure:"seqNum,omitempty"`

	// Type of the service.
	ServiceType *string `json:"serviceType,omitempty" msgpack:"serviceType,omitempty" bson:"servicetype,omitempty" mapstructure:"serviceType,omitempty"`

	// ID of the source PU.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"sourceid,omitempty" mapstructure:"sourceID,omitempty"`

	// Current stage when this report has been generated.
	Stage *string `json:"stage,omitempty" msgpack:"stage,omitempty" bson:"stage,omitempty" mapstructure:"stage,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"timestamp,omitempty" mapstructure:"timestamp,omitempty"`

	// geographical hash of the data. This is used for sharding and
	// georedundancy.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Logical storage zone. Used for sharding.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

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

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePingReport) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePingReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePingReport{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.RTT != nil {
		s.RTT = o.RTT
	}
	if o.TXController != nil {
		s.TXController = o.TXController
	}
	if o.TXType != nil {
		s.TXType = o.TXType
	}
	if o.ApplicationListening != nil {
		s.ApplicationListening = o.ApplicationListening
	}
	if o.DestinationID != nil {
		s.DestinationID = o.DestinationID
	}
	if o.EnforcerID != nil {
		s.EnforcerID = o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		s.EnforcerNamespace = o.EnforcerNamespace
	}
	if o.EnforcerVersion != nil {
		s.EnforcerVersion = o.EnforcerVersion
	}
	if o.FourTuple != nil {
		s.FourTuple = o.FourTuple
	}
	if o.IterationID != nil {
		s.IterationID = o.IterationID
	}
	if o.MigrationsLog != nil {
		s.MigrationsLog = o.MigrationsLog
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.PayloadSize != nil {
		s.PayloadSize = o.PayloadSize
	}
	if o.PingID != nil {
		s.PingID = o.PingID
	}
	if o.PolicyAction != nil {
		s.PolicyAction = o.PolicyAction
	}
	if o.PolicyID != nil {
		s.PolicyID = o.PolicyID
	}
	if o.Protocol != nil {
		s.Protocol = o.Protocol
	}
	if o.SeqNum != nil {
		s.SeqNum = o.SeqNum
	}
	if o.ServiceType != nil {
		s.ServiceType = o.ServiceType
	}
	if o.SourceID != nil {
		s.SourceID = o.SourceID
	}
	if o.Stage != nil {
		s.Stage = o.Stage
	}
	if o.Timestamp != nil {
		s.Timestamp = o.Timestamp
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
func (o *SparsePingReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePingReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.RTT != nil {
		o.RTT = s.RTT
	}
	if s.TXController != nil {
		o.TXController = s.TXController
	}
	if s.TXType != nil {
		o.TXType = s.TXType
	}
	if s.ApplicationListening != nil {
		o.ApplicationListening = s.ApplicationListening
	}
	if s.DestinationID != nil {
		o.DestinationID = s.DestinationID
	}
	if s.EnforcerID != nil {
		o.EnforcerID = s.EnforcerID
	}
	if s.EnforcerNamespace != nil {
		o.EnforcerNamespace = s.EnforcerNamespace
	}
	if s.EnforcerVersion != nil {
		o.EnforcerVersion = s.EnforcerVersion
	}
	if s.FourTuple != nil {
		o.FourTuple = s.FourTuple
	}
	if s.IterationID != nil {
		o.IterationID = s.IterationID
	}
	if s.MigrationsLog != nil {
		o.MigrationsLog = s.MigrationsLog
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.PayloadSize != nil {
		o.PayloadSize = s.PayloadSize
	}
	if s.PingID != nil {
		o.PingID = s.PingID
	}
	if s.PolicyAction != nil {
		o.PolicyAction = s.PolicyAction
	}
	if s.PolicyID != nil {
		o.PolicyID = s.PolicyID
	}
	if s.Protocol != nil {
		o.Protocol = s.Protocol
	}
	if s.SeqNum != nil {
		o.SeqNum = s.SeqNum
	}
	if s.ServiceType != nil {
		o.ServiceType = s.ServiceType
	}
	if s.SourceID != nil {
		o.SourceID = s.SourceID
	}
	if s.Stage != nil {
		o.Stage = s.Stage
	}
	if s.Timestamp != nil {
		o.Timestamp = s.Timestamp
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
func (o *SparsePingReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePingReport) ToPlain() elemental.PlainIdentifiable {

	out := NewPingReport()
	if o.ID != nil {
		out.ID = *o.ID
	}
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
	if o.MigrationsLog != nil {
		out.MigrationsLog = *o.MigrationsLog
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
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// GetMigrationsLog returns the MigrationsLog of the receiver.
func (o *SparsePingReport) GetMigrationsLog() (out map[string]string) {

	if o.MigrationsLog == nil {
		return
	}

	return *o.MigrationsLog
}

// SetMigrationsLog sets the property MigrationsLog of the receiver using the address of the given value.
func (o *SparsePingReport) SetMigrationsLog(migrationsLog map[string]string) {

	o.MigrationsLog = &migrationsLog
}

// GetZHash returns the ZHash of the receiver.
func (o *SparsePingReport) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparsePingReport) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparsePingReport) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparsePingReport) SetZone(zone int) {

	o.Zone = &zone
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
	ID                   bson.ObjectId     `bson:"_id,omitempty"`
	RTT                  string            `bson:"rtt"`
	TXController         string            `bson:"txcontroller"`
	TXType               string            `bson:"txtype"`
	ApplicationListening bool              `bson:"applicationlistening"`
	DestinationID        string            `bson:"destinationid"`
	EnforcerID           string            `bson:"enforcerid"`
	EnforcerNamespace    string            `bson:"enforcernamespace"`
	EnforcerVersion      string            `bson:"enforcerversion"`
	FourTuple            string            `bson:"fourtuple"`
	IterationID          string            `bson:"iterationid"`
	MigrationsLog        map[string]string `bson:"migrationslog,omitempty"`
	Namespace            string            `bson:"namespace"`
	PayloadSize          int               `bson:"payloadsize"`
	PingID               string            `bson:"pingid"`
	PolicyAction         string            `bson:"policyaction"`
	PolicyID             string            `bson:"policyid"`
	Protocol             int               `bson:"protocol"`
	SeqNum               int               `bson:"seqnum"`
	ServiceType          string            `bson:"servicetype"`
	SourceID             string            `bson:"sourceid"`
	Stage                string            `bson:"stage"`
	Timestamp            time.Time         `bson:"timestamp"`
	ZHash                int               `bson:"zhash"`
	Zone                 int               `bson:"zone"`
}
type mongoAttributesSparsePingReport struct {
	ID                   bson.ObjectId      `bson:"_id,omitempty"`
	RTT                  *string            `bson:"rtt,omitempty"`
	TXController         *string            `bson:"txcontroller,omitempty"`
	TXType               *string            `bson:"txtype,omitempty"`
	ApplicationListening *bool              `bson:"applicationlistening,omitempty"`
	DestinationID        *string            `bson:"destinationid,omitempty"`
	EnforcerID           *string            `bson:"enforcerid,omitempty"`
	EnforcerNamespace    *string            `bson:"enforcernamespace,omitempty"`
	EnforcerVersion      *string            `bson:"enforcerversion,omitempty"`
	FourTuple            *string            `bson:"fourtuple,omitempty"`
	IterationID          *string            `bson:"iterationid,omitempty"`
	MigrationsLog        *map[string]string `bson:"migrationslog,omitempty"`
	Namespace            *string            `bson:"namespace,omitempty"`
	PayloadSize          *int               `bson:"payloadsize,omitempty"`
	PingID               *string            `bson:"pingid,omitempty"`
	PolicyAction         *string            `bson:"policyaction,omitempty"`
	PolicyID             *string            `bson:"policyid,omitempty"`
	Protocol             *int               `bson:"protocol,omitempty"`
	SeqNum               *int               `bson:"seqnum,omitempty"`
	ServiceType          *string            `bson:"servicetype,omitempty"`
	SourceID             *string            `bson:"sourceid,omitempty"`
	Stage                *string            `bson:"stage,omitempty"`
	Timestamp            *time.Time         `bson:"timestamp,omitempty"`
	ZHash                *int               `bson:"zhash,omitempty"`
	Zone                 *int               `bson:"zone,omitempty"`
}
