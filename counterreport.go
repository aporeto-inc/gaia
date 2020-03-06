package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CounterReportIdentity represents the Identity of the object.
var CounterReportIdentity = elemental.Identity{
	Name:     "counterreport",
	Category: "counterreports",
	Package:  "zack",
	Private:  false,
}

// CounterReportsList represents a list of CounterReports
type CounterReportsList []*CounterReport

// Identity returns the identity of the objects in the list.
func (o CounterReportsList) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Copy returns a pointer to a copy the CounterReportsList.
func (o CounterReportsList) Copy() elemental.Identifiables {

	copy := append(CounterReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CounterReportsList.
func (o CounterReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CounterReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CounterReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CounterReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CounterReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CounterReportsList converted to SparseCounterReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CounterReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCounterReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCounterReport)
	}

	return out
}

// Version returns the version of the content.
func (o CounterReportsList) Version() int {

	return 1
}

// CounterReport represents the model of a counterreport
type CounterReport struct {
	// Counter for connections dropped.
	ConnectionsDropped int `json:"connectionsDropped" msgpack:"connectionsDropped" bson:"-" mapstructure:"connectionsDropped,omitempty"`

	// Counter for connections expired.
	ConnectionsExpired int `json:"connectionsExpired" msgpack:"connectionsExpired" bson:"-" mapstructure:"connectionsExpired,omitempty"`

	// Counter for connections processed.
	ConnectionsProcessed int `json:"connectionsProcessed" msgpack:"connectionsProcessed" bson:"-" mapstructure:"connectionsProcessed,omitempty"`

	// Counter for dropped packets.
	DroppedPackets int `json:"droppedPackets" msgpack:"droppedPackets" bson:"-" mapstructure:"droppedPackets,omitempty"`

	// Counter for encryption failures.
	EncryptionFailures int `json:"encryptionFailures" msgpack:"encryptionFailures" bson:"-" mapstructure:"encryptionFailures,omitempty"`

	// Identifier of the enforcer sending the report.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer sending the report.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"enforcernamespace" mapstructure:"enforcerNamespace,omitempty"`

	// Counter for external network connections.
	ExternalNetworkConnections int `json:"externalNetworkConnections" msgpack:"externalNetworkConnections" bson:"-" mapstructure:"externalNetworkConnections,omitempty"`

	// Counter for policy drops.
	PolicyDrops int `json:"policyDrops" msgpack:"policyDrops" bson:"-" mapstructure:"policyDrops,omitempty"`

	// PUID is the ID of the PU reporting the counter.
	ProcessingUnitID string `json:"processingUnitID" msgpack:"processingUnitID" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the PU reporting the counter.
	ProcessingUnitNamespace string `json:"processingUnitNamespace" msgpack:"processingUnitNamespace" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Timestamp is the date of the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	// Counter for token drops.
	TokenDrops int `json:"tokenDrops" msgpack:"tokenDrops" bson:"-" mapstructure:"tokenDrops,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCounterReport returns a new *CounterReport
func NewCounterReport() *CounterReport {

	return &CounterReport{
		ModelVersion:               1,
		ConnectionsExpired:         0,
		ConnectionsProcessed:       0,
		DroppedPackets:             0,
		EncryptionFailures:         0,
		ConnectionsDropped:         0,
		ExternalNetworkConnections: 0,
		PolicyDrops:                0,
		TokenDrops:                 0,
	}
}

// Identity returns the Identity of the object.
func (o *CounterReport) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CounterReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CounterReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CounterReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCounterReport{}

	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CounterReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCounterReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CounterReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CounterReport) BleveType() string {

	return "counterreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CounterReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CounterReport) Doc() string {

	return `Post a new counter tracing report.`
}

func (o *CounterReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CounterReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCounterReport{
			ConnectionsDropped:         &o.ConnectionsDropped,
			ConnectionsExpired:         &o.ConnectionsExpired,
			ConnectionsProcessed:       &o.ConnectionsProcessed,
			DroppedPackets:             &o.DroppedPackets,
			EncryptionFailures:         &o.EncryptionFailures,
			EnforcerID:                 &o.EnforcerID,
			EnforcerNamespace:          &o.EnforcerNamespace,
			ExternalNetworkConnections: &o.ExternalNetworkConnections,
			PolicyDrops:                &o.PolicyDrops,
			ProcessingUnitID:           &o.ProcessingUnitID,
			ProcessingUnitNamespace:    &o.ProcessingUnitNamespace,
			Timestamp:                  &o.Timestamp,
			TokenDrops:                 &o.TokenDrops,
		}
	}

	sp := &SparseCounterReport{}
	for _, f := range fields {
		switch f {
		case "connectionsDropped":
			sp.ConnectionsDropped = &(o.ConnectionsDropped)
		case "connectionsExpired":
			sp.ConnectionsExpired = &(o.ConnectionsExpired)
		case "connectionsProcessed":
			sp.ConnectionsProcessed = &(o.ConnectionsProcessed)
		case "droppedPackets":
			sp.DroppedPackets = &(o.DroppedPackets)
		case "encryptionFailures":
			sp.EncryptionFailures = &(o.EncryptionFailures)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "externalNetworkConnections":
			sp.ExternalNetworkConnections = &(o.ExternalNetworkConnections)
		case "policyDrops":
			sp.PolicyDrops = &(o.PolicyDrops)
		case "processingUnitID":
			sp.ProcessingUnitID = &(o.ProcessingUnitID)
		case "processingUnitNamespace":
			sp.ProcessingUnitNamespace = &(o.ProcessingUnitNamespace)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "tokenDrops":
			sp.TokenDrops = &(o.TokenDrops)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCounterReport to the object.
func (o *CounterReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCounterReport)
	if so.ConnectionsDropped != nil {
		o.ConnectionsDropped = *so.ConnectionsDropped
	}
	if so.ConnectionsExpired != nil {
		o.ConnectionsExpired = *so.ConnectionsExpired
	}
	if so.ConnectionsProcessed != nil {
		o.ConnectionsProcessed = *so.ConnectionsProcessed
	}
	if so.DroppedPackets != nil {
		o.DroppedPackets = *so.DroppedPackets
	}
	if so.EncryptionFailures != nil {
		o.EncryptionFailures = *so.EncryptionFailures
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.ExternalNetworkConnections != nil {
		o.ExternalNetworkConnections = *so.ExternalNetworkConnections
	}
	if so.PolicyDrops != nil {
		o.PolicyDrops = *so.PolicyDrops
	}
	if so.ProcessingUnitID != nil {
		o.ProcessingUnitID = *so.ProcessingUnitID
	}
	if so.ProcessingUnitNamespace != nil {
		o.ProcessingUnitNamespace = *so.ProcessingUnitNamespace
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.TokenDrops != nil {
		o.TokenDrops = *so.TokenDrops
	}
}

// DeepCopy returns a deep copy if the CounterReport.
func (o *CounterReport) DeepCopy() *CounterReport {

	if o == nil {
		return nil
	}

	out := &CounterReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CounterReport.
func (o *CounterReport) DeepCopyInto(out *CounterReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CounterReport: %s", err))
	}

	*out = *target.(*CounterReport)
}

// Validate valides the current information stored into the structure.
func (o *CounterReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
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
func (*CounterReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CounterReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CounterReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CounterReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CounterReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CounterReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "connectionsDropped":
		return o.ConnectionsDropped
	case "connectionsExpired":
		return o.ConnectionsExpired
	case "connectionsProcessed":
		return o.ConnectionsProcessed
	case "droppedPackets":
		return o.DroppedPackets
	case "encryptionFailures":
		return o.EncryptionFailures
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "externalNetworkConnections":
		return o.ExternalNetworkConnections
	case "policyDrops":
		return o.PolicyDrops
	case "processingUnitID":
		return o.ProcessingUnitID
	case "processingUnitNamespace":
		return o.ProcessingUnitNamespace
	case "timestamp":
		return o.Timestamp
	case "tokenDrops":
		return o.TokenDrops
	}

	return nil
}

// CounterReportAttributesMap represents the map of attribute for CounterReport.
var CounterReportAttributesMap = map[string]elemental.AttributeSpecification{
	"ConnectionsDropped": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsDropped",
		Description:    `Counter for connections dropped.`,
		Exposed:        true,
		Name:           "connectionsDropped",
		Type:           "integer",
	},
	"ConnectionsExpired": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsExpired",
		Description:    `Counter for connections expired.`,
		Exposed:        true,
		Name:           "connectionsExpired",
		Type:           "integer",
	},
	"ConnectionsProcessed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsProcessed",
		Description:    `Counter for connections processed.`,
		Exposed:        true,
		Name:           "connectionsProcessed",
		Type:           "integer",
	},
	"DroppedPackets": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DroppedPackets",
		Description:    `Counter for dropped packets.`,
		Exposed:        true,
		Name:           "droppedPackets",
		Type:           "integer",
	},
	"EncryptionFailures": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionFailures",
		Description:    `Counter for encryption failures.`,
		Exposed:        true,
		Name:           "encryptionFailures",
		Type:           "integer",
	},
	"EnforcerID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"ExternalNetworkConnections": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExternalNetworkConnections",
		Description:    `Counter for external network connections.`,
		Exposed:        true,
		Name:           "externalNetworkConnections",
		Type:           "integer",
	},
	"PolicyDrops": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyDrops",
		Description:    `Counter for policy drops.`,
		Exposed:        true,
		Name:           "policyDrops",
		Type:           "integer",
	},
	"ProcessingUnitID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `PUID is the ID of the PU reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitID",
		Type:           "string",
	},
	"ProcessingUnitNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the PU reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitNamespace",
		Type:           "string",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp is the date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"TokenDrops": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TokenDrops",
		Description:    `Counter for token drops.`,
		Exposed:        true,
		Name:           "tokenDrops",
		Type:           "integer",
	},
}

// CounterReportLowerCaseAttributesMap represents the map of attribute for CounterReport.
var CounterReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"connectionsdropped": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsDropped",
		Description:    `Counter for connections dropped.`,
		Exposed:        true,
		Name:           "connectionsDropped",
		Type:           "integer",
	},
	"connectionsexpired": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsExpired",
		Description:    `Counter for connections expired.`,
		Exposed:        true,
		Name:           "connectionsExpired",
		Type:           "integer",
	},
	"connectionsprocessed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ConnectionsProcessed",
		Description:    `Counter for connections processed.`,
		Exposed:        true,
		Name:           "connectionsProcessed",
		Type:           "integer",
	},
	"droppedpackets": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DroppedPackets",
		Description:    `Counter for dropped packets.`,
		Exposed:        true,
		Name:           "droppedPackets",
		Type:           "integer",
	},
	"encryptionfailures": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EncryptionFailures",
		Description:    `Counter for encryption failures.`,
		Exposed:        true,
		Name:           "encryptionFailures",
		Type:           "integer",
	},
	"enforcerid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the enforcer sending the report.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"externalnetworkconnections": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ExternalNetworkConnections",
		Description:    `Counter for external network connections.`,
		Exposed:        true,
		Name:           "externalNetworkConnections",
		Type:           "integer",
	},
	"policydrops": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyDrops",
		Description:    `Counter for policy drops.`,
		Exposed:        true,
		Name:           "policyDrops",
		Type:           "integer",
	},
	"processingunitid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitID",
		Description:    `PUID is the ID of the PU reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitID",
		Type:           "string",
	},
	"processingunitnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ProcessingUnitNamespace",
		Description:    `Namespace of the PU reporting the counter.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "processingUnitNamespace",
		Type:           "string",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp is the date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"tokendrops": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TokenDrops",
		Description:    `Counter for token drops.`,
		Exposed:        true,
		Name:           "tokenDrops",
		Type:           "integer",
	},
}

// SparseCounterReportsList represents a list of SparseCounterReports
type SparseCounterReportsList []*SparseCounterReport

// Identity returns the identity of the objects in the list.
func (o SparseCounterReportsList) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Copy returns a pointer to a copy the SparseCounterReportsList.
func (o SparseCounterReportsList) Copy() elemental.Identifiables {

	copy := append(SparseCounterReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCounterReportsList.
func (o SparseCounterReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCounterReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCounterReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCounterReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCounterReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCounterReportsList converted to CounterReportsList.
func (o SparseCounterReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCounterReportsList) Version() int {

	return 1
}

// SparseCounterReport represents the sparse version of a counterreport.
type SparseCounterReport struct {
	// Counter for connections dropped.
	ConnectionsDropped *int `json:"connectionsDropped,omitempty" msgpack:"connectionsDropped,omitempty" bson:"-" mapstructure:"connectionsDropped,omitempty"`

	// Counter for connections expired.
	ConnectionsExpired *int `json:"connectionsExpired,omitempty" msgpack:"connectionsExpired,omitempty" bson:"-" mapstructure:"connectionsExpired,omitempty"`

	// Counter for connections processed.
	ConnectionsProcessed *int `json:"connectionsProcessed,omitempty" msgpack:"connectionsProcessed,omitempty" bson:"-" mapstructure:"connectionsProcessed,omitempty"`

	// Counter for dropped packets.
	DroppedPackets *int `json:"droppedPackets,omitempty" msgpack:"droppedPackets,omitempty" bson:"-" mapstructure:"droppedPackets,omitempty"`

	// Counter for encryption failures.
	EncryptionFailures *int `json:"encryptionFailures,omitempty" msgpack:"encryptionFailures,omitempty" bson:"-" mapstructure:"encryptionFailures,omitempty"`

	// Identifier of the enforcer sending the report.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the enforcer sending the report.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// Counter for external network connections.
	ExternalNetworkConnections *int `json:"externalNetworkConnections,omitempty" msgpack:"externalNetworkConnections,omitempty" bson:"-" mapstructure:"externalNetworkConnections,omitempty"`

	// Counter for policy drops.
	PolicyDrops *int `json:"policyDrops,omitempty" msgpack:"policyDrops,omitempty" bson:"-" mapstructure:"policyDrops,omitempty"`

	// PUID is the ID of the PU reporting the counter.
	ProcessingUnitID *string `json:"processingUnitID,omitempty" msgpack:"processingUnitID,omitempty" bson:"-" mapstructure:"processingUnitID,omitempty"`

	// Namespace of the PU reporting the counter.
	ProcessingUnitNamespace *string `json:"processingUnitNamespace,omitempty" msgpack:"processingUnitNamespace,omitempty" bson:"-" mapstructure:"processingUnitNamespace,omitempty"`

	// Timestamp is the date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	// Counter for token drops.
	TokenDrops *int `json:"tokenDrops,omitempty" msgpack:"tokenDrops,omitempty" bson:"-" mapstructure:"tokenDrops,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCounterReport returns a new  SparseCounterReport.
func NewSparseCounterReport() *SparseCounterReport {
	return &SparseCounterReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCounterReport) Identity() elemental.Identity {

	return CounterReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCounterReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCounterReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCounterReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCounterReport{}

	if o.EnforcerID != nil {
		s.EnforcerID = o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		s.EnforcerNamespace = o.EnforcerNamespace
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCounterReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCounterReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.EnforcerID != nil {
		o.EnforcerID = s.EnforcerID
	}
	if s.EnforcerNamespace != nil {
		o.EnforcerNamespace = s.EnforcerNamespace
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCounterReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCounterReport) ToPlain() elemental.PlainIdentifiable {

	out := NewCounterReport()
	if o.ConnectionsDropped != nil {
		out.ConnectionsDropped = *o.ConnectionsDropped
	}
	if o.ConnectionsExpired != nil {
		out.ConnectionsExpired = *o.ConnectionsExpired
	}
	if o.ConnectionsProcessed != nil {
		out.ConnectionsProcessed = *o.ConnectionsProcessed
	}
	if o.DroppedPackets != nil {
		out.DroppedPackets = *o.DroppedPackets
	}
	if o.EncryptionFailures != nil {
		out.EncryptionFailures = *o.EncryptionFailures
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.ExternalNetworkConnections != nil {
		out.ExternalNetworkConnections = *o.ExternalNetworkConnections
	}
	if o.PolicyDrops != nil {
		out.PolicyDrops = *o.PolicyDrops
	}
	if o.ProcessingUnitID != nil {
		out.ProcessingUnitID = *o.ProcessingUnitID
	}
	if o.ProcessingUnitNamespace != nil {
		out.ProcessingUnitNamespace = *o.ProcessingUnitNamespace
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.TokenDrops != nil {
		out.TokenDrops = *o.TokenDrops
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCounterReport.
func (o *SparseCounterReport) DeepCopy() *SparseCounterReport {

	if o == nil {
		return nil
	}

	out := &SparseCounterReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCounterReport.
func (o *SparseCounterReport) DeepCopyInto(out *SparseCounterReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCounterReport: %s", err))
	}

	*out = *target.(*SparseCounterReport)
}

type mongoAttributesCounterReport struct {
	EnforcerID        string `bson:"enforcerid"`
	EnforcerNamespace string `bson:"enforcernamespace"`
}
type mongoAttributesSparseCounterReport struct {
	EnforcerID        *string `bson:"enforcerid,omitempty"`
	EnforcerNamespace *string `bson:"enforcernamespace,omitempty"`
}
