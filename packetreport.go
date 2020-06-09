package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PacketReportEventValue represents the possible values for attribute "event".
type PacketReportEventValue string

const (
	// PacketReportEventDropped represents the value Dropped.
	PacketReportEventDropped PacketReportEventValue = "Dropped"

	// PacketReportEventReceived represents the value Received.
	PacketReportEventReceived PacketReportEventValue = "Received"

	// PacketReportEventTransmitted represents the value Transmitted.
	PacketReportEventTransmitted PacketReportEventValue = "Transmitted"
)

// PacketReportIdentity represents the Identity of the object.
var PacketReportIdentity = elemental.Identity{
	Name:     "packetreport",
	Category: "packetreports",
	Package:  "zack",
	Private:  false,
}

// PacketReportsList represents a list of PacketReports
type PacketReportsList []*PacketReport

// Identity returns the identity of the objects in the list.
func (o PacketReportsList) Identity() elemental.Identity {

	return PacketReportIdentity
}

// Copy returns a pointer to a copy the PacketReportsList.
func (o PacketReportsList) Copy() elemental.Identifiables {

	copy := append(PacketReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the PacketReportsList.
func (o PacketReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(PacketReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*PacketReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o PacketReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o PacketReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the PacketReportsList converted to SparsePacketReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o PacketReportsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparsePacketReportsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparsePacketReport)
	}

	return out
}

// Version returns the version of the content.
func (o PacketReportsList) Version() int {

	return 1
}

// PacketReport represents the model of a packetreport
type PacketReport struct {
	// Flags are the TCP flags of the packet.
	TCPFlags int `json:"TCPFlags" msgpack:"TCPFlags" bson:"tcpflags" mapstructure:"TCPFlags,omitempty"`

	// Claims is the list of claims detected for the packet.
	Claims []string `json:"claims" msgpack:"claims" bson:"claims" mapstructure:"claims,omitempty"`

	// The destination IP address of the packet.
	DestinationIP string `json:"destinationIP" msgpack:"destinationIP" bson:"destinationip" mapstructure:"destinationIP,omitempty"`

	// The destination port of a TCP or UDP packet.
	DestinationPort int `json:"destinationPort" msgpack:"destinationPort" bson:"destinationport" mapstructure:"destinationPort,omitempty"`

	// If `event` is set to `Dropped`, contains the reason that the packet was dropped.
	// Otherwise empty.
	DropReason string `json:"dropReason" msgpack:"dropReason" bson:"dropreason" mapstructure:"dropReason,omitempty"`

	// Set to `true` if the packet was encrypted.
	Encrypt bool `json:"encrypt" msgpack:"encrypt" bson:"encrypt" mapstructure:"encrypt,omitempty"`

	// Identifier of the defender sending the report.
	EnforcerID string `json:"enforcerID" msgpack:"enforcerID" bson:"enforcerid" mapstructure:"enforcerID,omitempty"`

	// Namespace of the defender sending the report.
	EnforcerNamespace string `json:"enforcerNamespace" msgpack:"enforcerNamespace" bson:"enforcernamespace" mapstructure:"enforcerNamespace,omitempty"`

	// The event that triggered the report.
	Event PacketReportEventValue `json:"event" msgpack:"event" bson:"-" mapstructure:"event,omitempty"`

	// Length is the length of the packet.
	Length int `json:"-" msgpack:"-" bson:"length" mapstructure:"-,omitempty"`

	// Mark is the mark value of the packet.
	Mark int `json:"mark" msgpack:"mark" bson:"mark" mapstructure:"mark,omitempty"`

	// Namespace of the processing unit reporting the packet.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"namespace" mapstructure:"namespace,omitempty"`

	// The ID of the IP header of the reported packet.
	PacketID int `json:"packetID" msgpack:"packetID" bson:"packetid" mapstructure:"packetID,omitempty"`

	// Protocol number.
	Protocol int `json:"protocol" msgpack:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	// The ID of the processing unit reporting the packet.
	PuID string `json:"puID" msgpack:"puID" bson:"puid" mapstructure:"puID,omitempty"`

	// The first 64 bytes of the packet.
	RawPacket string `json:"rawPacket" msgpack:"rawPacket" bson:"rawpacket" mapstructure:"rawPacket,omitempty"`

	// The source IP address of the packet.
	SourceIP string `json:"sourceIP" msgpack:"sourceIP" bson:"sourceip" mapstructure:"sourceIP,omitempty"`

	// The source port of the packet.
	SourcePort int `json:"sourcePort" msgpack:"sourcePort" bson:"sourceport" mapstructure:"sourcePort,omitempty"`

	// The time-date stamp of the report.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"timestamp" mapstructure:"timestamp,omitempty"`

	// Set to `true` if the packet arrived with the Trireme options (default).
	TriremePacket bool `json:"triremePacket" msgpack:"triremePacket" bson:"triremepacket" mapstructure:"triremePacket,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPacketReport returns a new *PacketReport
func NewPacketReport() *PacketReport {

	return &PacketReport{
		ModelVersion:  1,
		Claims:        []string{},
		RawPacket:     "abcd",
		TriremePacket: true,
	}
}

// Identity returns the Identity of the object.
func (o *PacketReport) Identity() elemental.Identity {

	return PacketReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *PacketReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *PacketReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PacketReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPacketReport{}

	s.TCPFlags = o.TCPFlags
	s.Claims = o.Claims
	s.DestinationIP = o.DestinationIP
	s.DestinationPort = o.DestinationPort
	s.DropReason = o.DropReason
	s.Encrypt = o.Encrypt
	s.EnforcerID = o.EnforcerID
	s.EnforcerNamespace = o.EnforcerNamespace
	s.Length = o.Length
	s.Mark = o.Mark
	s.Namespace = o.Namespace
	s.PacketID = o.PacketID
	s.Protocol = o.Protocol
	s.PuID = o.PuID
	s.RawPacket = o.RawPacket
	s.SourceIP = o.SourceIP
	s.SourcePort = o.SourcePort
	s.Timestamp = o.Timestamp
	s.TriremePacket = o.TriremePacket

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PacketReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPacketReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.TCPFlags = s.TCPFlags
	o.Claims = s.Claims
	o.DestinationIP = s.DestinationIP
	o.DestinationPort = s.DestinationPort
	o.DropReason = s.DropReason
	o.Encrypt = s.Encrypt
	o.EnforcerID = s.EnforcerID
	o.EnforcerNamespace = s.EnforcerNamespace
	o.Length = s.Length
	o.Mark = s.Mark
	o.Namespace = s.Namespace
	o.PacketID = s.PacketID
	o.Protocol = s.Protocol
	o.PuID = s.PuID
	o.RawPacket = s.RawPacket
	o.SourceIP = s.SourceIP
	o.SourcePort = s.SourcePort
	o.Timestamp = s.Timestamp
	o.TriremePacket = s.TriremePacket

	return nil
}

// Version returns the hardcoded version of the model.
func (o *PacketReport) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *PacketReport) BleveType() string {

	return "packetreport"
}

// DefaultOrder returns the list of default ordering fields.
func (o *PacketReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *PacketReport) Doc() string {

	return `Post a new packet tracing report.`
}

func (o *PacketReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *PacketReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparsePacketReport{
			TCPFlags:          &o.TCPFlags,
			Claims:            &o.Claims,
			DestinationIP:     &o.DestinationIP,
			DestinationPort:   &o.DestinationPort,
			DropReason:        &o.DropReason,
			Encrypt:           &o.Encrypt,
			EnforcerID:        &o.EnforcerID,
			EnforcerNamespace: &o.EnforcerNamespace,
			Event:             &o.Event,
			Length:            &o.Length,
			Mark:              &o.Mark,
			Namespace:         &o.Namespace,
			PacketID:          &o.PacketID,
			Protocol:          &o.Protocol,
			PuID:              &o.PuID,
			RawPacket:         &o.RawPacket,
			SourceIP:          &o.SourceIP,
			SourcePort:        &o.SourcePort,
			Timestamp:         &o.Timestamp,
			TriremePacket:     &o.TriremePacket,
		}
	}

	sp := &SparsePacketReport{}
	for _, f := range fields {
		switch f {
		case "TCPFlags":
			sp.TCPFlags = &(o.TCPFlags)
		case "claims":
			sp.Claims = &(o.Claims)
		case "destinationIP":
			sp.DestinationIP = &(o.DestinationIP)
		case "destinationPort":
			sp.DestinationPort = &(o.DestinationPort)
		case "dropReason":
			sp.DropReason = &(o.DropReason)
		case "encrypt":
			sp.Encrypt = &(o.Encrypt)
		case "enforcerID":
			sp.EnforcerID = &(o.EnforcerID)
		case "enforcerNamespace":
			sp.EnforcerNamespace = &(o.EnforcerNamespace)
		case "event":
			sp.Event = &(o.Event)
		case "length":
			sp.Length = &(o.Length)
		case "mark":
			sp.Mark = &(o.Mark)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "packetID":
			sp.PacketID = &(o.PacketID)
		case "protocol":
			sp.Protocol = &(o.Protocol)
		case "puID":
			sp.PuID = &(o.PuID)
		case "rawPacket":
			sp.RawPacket = &(o.RawPacket)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "sourcePort":
			sp.SourcePort = &(o.SourcePort)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "triremePacket":
			sp.TriremePacket = &(o.TriremePacket)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparsePacketReport to the object.
func (o *PacketReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparsePacketReport)
	if so.TCPFlags != nil {
		o.TCPFlags = *so.TCPFlags
	}
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.DestinationIP != nil {
		o.DestinationIP = *so.DestinationIP
	}
	if so.DestinationPort != nil {
		o.DestinationPort = *so.DestinationPort
	}
	if so.DropReason != nil {
		o.DropReason = *so.DropReason
	}
	if so.Encrypt != nil {
		o.Encrypt = *so.Encrypt
	}
	if so.EnforcerID != nil {
		o.EnforcerID = *so.EnforcerID
	}
	if so.EnforcerNamespace != nil {
		o.EnforcerNamespace = *so.EnforcerNamespace
	}
	if so.Event != nil {
		o.Event = *so.Event
	}
	if so.Length != nil {
		o.Length = *so.Length
	}
	if so.Mark != nil {
		o.Mark = *so.Mark
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.PacketID != nil {
		o.PacketID = *so.PacketID
	}
	if so.Protocol != nil {
		o.Protocol = *so.Protocol
	}
	if so.PuID != nil {
		o.PuID = *so.PuID
	}
	if so.RawPacket != nil {
		o.RawPacket = *so.RawPacket
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
	}
	if so.SourcePort != nil {
		o.SourcePort = *so.SourcePort
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.TriremePacket != nil {
		o.TriremePacket = *so.TriremePacket
	}
}

// DeepCopy returns a deep copy if the PacketReport.
func (o *PacketReport) DeepCopy() *PacketReport {

	if o == nil {
		return nil
	}

	out := &PacketReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PacketReport.
func (o *PacketReport) DeepCopyInto(out *PacketReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PacketReport: %s", err))
	}

	*out = *target.(*PacketReport)
}

// Validate valides the current information stored into the structure.
func (o *PacketReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateMaximumInt("destinationPort", o.DestinationPort, int(65536), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerID", o.EnforcerID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("enforcerNamespace", o.EnforcerNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("event", string(o.Event)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("event", string(o.Event), []string{"Received", "Transmitted", "Dropped"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("protocol", o.Protocol, int(255), false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateMaximumInt("sourcePort", o.SourcePort, int(65536), false); err != nil {
		errors = errors.Append(err)
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
func (*PacketReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := PacketReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return PacketReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*PacketReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return PacketReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *PacketReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "TCPFlags":
		return o.TCPFlags
	case "claims":
		return o.Claims
	case "destinationIP":
		return o.DestinationIP
	case "destinationPort":
		return o.DestinationPort
	case "dropReason":
		return o.DropReason
	case "encrypt":
		return o.Encrypt
	case "enforcerID":
		return o.EnforcerID
	case "enforcerNamespace":
		return o.EnforcerNamespace
	case "event":
		return o.Event
	case "length":
		return o.Length
	case "mark":
		return o.Mark
	case "namespace":
		return o.Namespace
	case "packetID":
		return o.PacketID
	case "protocol":
		return o.Protocol
	case "puID":
		return o.PuID
	case "rawPacket":
		return o.RawPacket
	case "sourceIP":
		return o.SourceIP
	case "sourcePort":
		return o.SourcePort
	case "timestamp":
		return o.Timestamp
	case "triremePacket":
		return o.TriremePacket
	}

	return nil
}

// PacketReportAttributesMap represents the map of attribute for PacketReport.
var PacketReportAttributesMap = map[string]elemental.AttributeSpecification{
	"TCPFlags": {
		AllowedChoices: []string{},
		ConvertedName:  "TCPFlags",
		Description:    `Flags are the TCP flags of the packet.`,
		Exposed:        true,
		Name:           "TCPFlags",
		Stored:         true,
		Type:           "integer",
	},
	"Claims": {
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims is the list of claims detected for the packet.`,
		Exposed:        true,
		Name:           "claims",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"DestinationIP": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `The destination IP address of the packet.`,
		Exposed:        true,
		Name:           "destinationIP",
		Stored:         true,
		Type:           "string",
	},
	"DestinationPort": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `The destination port of a TCP or UDP packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "destinationPort",
		Stored:         true,
		Type:           "integer",
	},
	"DropReason": {
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `If ` + "`" + `event` + "`" + ` is set to ` + "`" + `Dropped` + "`" + `, contains the reason that the packet was dropped.
Otherwise empty.`,
		Exposed: true,
		Name:    "dropReason",
		Stored:  true,
		Type:    "string",
	},
	"Encrypt": {
		AllowedChoices: []string{},
		ConvertedName:  "Encrypt",
		Description:    `Set to ` + "`" + `true` + "`" + ` if the packet was encrypted.`,
		Exposed:        true,
		Name:           "encrypt",
		Stored:         true,
		Type:           "boolean",
	},
	"EnforcerID": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the defender sending the report.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"EnforcerNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the defender sending the report.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Event": {
		AllowedChoices: []string{"Received", "Transmitted", "Dropped"},
		ConvertedName:  "Event",
		Description:    `The event that triggered the report.`,
		Exposed:        true,
		Name:           "event",
		Required:       true,
		Type:           "enum",
	},
	"Length": {
		AllowedChoices: []string{},
		ConvertedName:  "Length",
		Description:    `Length is the length of the packet.`,
		MaxValue:       65536,
		Name:           "length",
		Stored:         true,
		Type:           "integer",
	},
	"Mark": {
		AllowedChoices: []string{},
		ConvertedName:  "Mark",
		Description:    `Mark is the mark value of the packet.`,
		Exposed:        true,
		Name:           "mark",
		Stored:         true,
		Type:           "integer",
	},
	"Namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the processing unit reporting the packet.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "namespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"PacketID": {
		AllowedChoices: []string{},
		ConvertedName:  "PacketID",
		Description:    `The ID of the IP header of the reported packet.`,
		Exposed:        true,
		Name:           "packetID",
		Stored:         true,
		Type:           "integer",
	},
	"Protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol number.`,
		Exposed:        true,
		MaxValue:       255,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"PuID": {
		AllowedChoices: []string{},
		ConvertedName:  "PuID",
		Description:    `The ID of the processing unit reporting the packet.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "puID",
		Stored:         true,
		Type:           "string",
	},
	"RawPacket": {
		AllowedChoices: []string{},
		ConvertedName:  "RawPacket",
		DefaultValue:   "abcd",
		Description:    `The first 64 bytes of the packet.`,
		Exposed:        true,
		Name:           "rawPacket",
		Stored:         true,
		Type:           "string",
	},
	"SourceIP": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `The source IP address of the packet.`,
		Exposed:        true,
		Name:           "sourceIP",
		Stored:         true,
		Type:           "string",
	},
	"SourcePort": {
		AllowedChoices: []string{},
		ConvertedName:  "SourcePort",
		Description:    `The source port of the packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "sourcePort",
		Stored:         true,
		Type:           "integer",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `The time-date stamp of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Stored:         true,
		Type:           "time",
	},
	"TriremePacket": {
		AllowedChoices: []string{},
		ConvertedName:  "TriremePacket",
		DefaultValue:   true,
		Description:    `Set to ` + "`" + `true` + "`" + ` if the packet arrived with the Trireme options (default).`,
		Exposed:        true,
		Name:           "triremePacket",
		Stored:         true,
		Type:           "boolean",
	},
}

// PacketReportLowerCaseAttributesMap represents the map of attribute for PacketReport.
var PacketReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"tcpflags": {
		AllowedChoices: []string{},
		ConvertedName:  "TCPFlags",
		Description:    `Flags are the TCP flags of the packet.`,
		Exposed:        true,
		Name:           "TCPFlags",
		Stored:         true,
		Type:           "integer",
	},
	"claims": {
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims is the list of claims detected for the packet.`,
		Exposed:        true,
		Name:           "claims",
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"destinationip": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `The destination IP address of the packet.`,
		Exposed:        true,
		Name:           "destinationIP",
		Stored:         true,
		Type:           "string",
	},
	"destinationport": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `The destination port of a TCP or UDP packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "destinationPort",
		Stored:         true,
		Type:           "integer",
	},
	"dropreason": {
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `If ` + "`" + `event` + "`" + ` is set to ` + "`" + `Dropped` + "`" + `, contains the reason that the packet was dropped.
Otherwise empty.`,
		Exposed: true,
		Name:    "dropReason",
		Stored:  true,
		Type:    "string",
	},
	"encrypt": {
		AllowedChoices: []string{},
		ConvertedName:  "Encrypt",
		Description:    `Set to ` + "`" + `true` + "`" + ` if the packet was encrypted.`,
		Exposed:        true,
		Name:           "encrypt",
		Stored:         true,
		Type:           "boolean",
	},
	"enforcerid": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerID",
		Description:    `Identifier of the defender sending the report.`,
		Exposed:        true,
		Name:           "enforcerID",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"enforcernamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "EnforcerNamespace",
		Description:    `Namespace of the defender sending the report.`,
		Exposed:        true,
		Name:           "enforcerNamespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"event": {
		AllowedChoices: []string{"Received", "Transmitted", "Dropped"},
		ConvertedName:  "Event",
		Description:    `The event that triggered the report.`,
		Exposed:        true,
		Name:           "event",
		Required:       true,
		Type:           "enum",
	},
	"length": {
		AllowedChoices: []string{},
		ConvertedName:  "Length",
		Description:    `Length is the length of the packet.`,
		MaxValue:       65536,
		Name:           "length",
		Stored:         true,
		Type:           "integer",
	},
	"mark": {
		AllowedChoices: []string{},
		ConvertedName:  "Mark",
		Description:    `Mark is the mark value of the packet.`,
		Exposed:        true,
		Name:           "mark",
		Stored:         true,
		Type:           "integer",
	},
	"namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the processing unit reporting the packet.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "namespace",
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"packetid": {
		AllowedChoices: []string{},
		ConvertedName:  "PacketID",
		Description:    `The ID of the IP header of the reported packet.`,
		Exposed:        true,
		Name:           "packetID",
		Stored:         true,
		Type:           "integer",
	},
	"protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol number.`,
		Exposed:        true,
		MaxValue:       255,
		Name:           "protocol",
		Stored:         true,
		Type:           "integer",
	},
	"puid": {
		AllowedChoices: []string{},
		ConvertedName:  "PuID",
		Description:    `The ID of the processing unit reporting the packet.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "puID",
		Stored:         true,
		Type:           "string",
	},
	"rawpacket": {
		AllowedChoices: []string{},
		ConvertedName:  "RawPacket",
		DefaultValue:   "abcd",
		Description:    `The first 64 bytes of the packet.`,
		Exposed:        true,
		Name:           "rawPacket",
		Stored:         true,
		Type:           "string",
	},
	"sourceip": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `The source IP address of the packet.`,
		Exposed:        true,
		Name:           "sourceIP",
		Stored:         true,
		Type:           "string",
	},
	"sourceport": {
		AllowedChoices: []string{},
		ConvertedName:  "SourcePort",
		Description:    `The source port of the packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "sourcePort",
		Stored:         true,
		Type:           "integer",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `The time-date stamp of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Stored:         true,
		Type:           "time",
	},
	"triremepacket": {
		AllowedChoices: []string{},
		ConvertedName:  "TriremePacket",
		DefaultValue:   true,
		Description:    `Set to ` + "`" + `true` + "`" + ` if the packet arrived with the Trireme options (default).`,
		Exposed:        true,
		Name:           "triremePacket",
		Stored:         true,
		Type:           "boolean",
	},
}

// SparsePacketReportsList represents a list of SparsePacketReports
type SparsePacketReportsList []*SparsePacketReport

// Identity returns the identity of the objects in the list.
func (o SparsePacketReportsList) Identity() elemental.Identity {

	return PacketReportIdentity
}

// Copy returns a pointer to a copy the SparsePacketReportsList.
func (o SparsePacketReportsList) Copy() elemental.Identifiables {

	copy := append(SparsePacketReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparsePacketReportsList.
func (o SparsePacketReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparsePacketReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparsePacketReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparsePacketReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparsePacketReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparsePacketReportsList converted to PacketReportsList.
func (o SparsePacketReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparsePacketReportsList) Version() int {

	return 1
}

// SparsePacketReport represents the sparse version of a packetreport.
type SparsePacketReport struct {
	// Flags are the TCP flags of the packet.
	TCPFlags *int `json:"TCPFlags,omitempty" msgpack:"TCPFlags,omitempty" bson:"tcpflags,omitempty" mapstructure:"TCPFlags,omitempty"`

	// Claims is the list of claims detected for the packet.
	Claims *[]string `json:"claims,omitempty" msgpack:"claims,omitempty" bson:"claims,omitempty" mapstructure:"claims,omitempty"`

	// The destination IP address of the packet.
	DestinationIP *string `json:"destinationIP,omitempty" msgpack:"destinationIP,omitempty" bson:"destinationip,omitempty" mapstructure:"destinationIP,omitempty"`

	// The destination port of a TCP or UDP packet.
	DestinationPort *int `json:"destinationPort,omitempty" msgpack:"destinationPort,omitempty" bson:"destinationport,omitempty" mapstructure:"destinationPort,omitempty"`

	// If `event` is set to `Dropped`, contains the reason that the packet was dropped.
	// Otherwise empty.
	DropReason *string `json:"dropReason,omitempty" msgpack:"dropReason,omitempty" bson:"dropreason,omitempty" mapstructure:"dropReason,omitempty"`

	// Set to `true` if the packet was encrypted.
	Encrypt *bool `json:"encrypt,omitempty" msgpack:"encrypt,omitempty" bson:"encrypt,omitempty" mapstructure:"encrypt,omitempty"`

	// Identifier of the defender sending the report.
	EnforcerID *string `json:"enforcerID,omitempty" msgpack:"enforcerID,omitempty" bson:"enforcerid,omitempty" mapstructure:"enforcerID,omitempty"`

	// Namespace of the defender sending the report.
	EnforcerNamespace *string `json:"enforcerNamespace,omitempty" msgpack:"enforcerNamespace,omitempty" bson:"enforcernamespace,omitempty" mapstructure:"enforcerNamespace,omitempty"`

	// The event that triggered the report.
	Event *PacketReportEventValue `json:"event,omitempty" msgpack:"event,omitempty" bson:"-" mapstructure:"event,omitempty"`

	// Length is the length of the packet.
	Length *int `json:"-" msgpack:"-" bson:"length,omitempty" mapstructure:"-,omitempty"`

	// Mark is the mark value of the packet.
	Mark *int `json:"mark,omitempty" msgpack:"mark,omitempty" bson:"mark,omitempty" mapstructure:"mark,omitempty"`

	// Namespace of the processing unit reporting the packet.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The ID of the IP header of the reported packet.
	PacketID *int `json:"packetID,omitempty" msgpack:"packetID,omitempty" bson:"packetid,omitempty" mapstructure:"packetID,omitempty"`

	// Protocol number.
	Protocol *int `json:"protocol,omitempty" msgpack:"protocol,omitempty" bson:"protocol,omitempty" mapstructure:"protocol,omitempty"`

	// The ID of the processing unit reporting the packet.
	PuID *string `json:"puID,omitempty" msgpack:"puID,omitempty" bson:"puid,omitempty" mapstructure:"puID,omitempty"`

	// The first 64 bytes of the packet.
	RawPacket *string `json:"rawPacket,omitempty" msgpack:"rawPacket,omitempty" bson:"rawpacket,omitempty" mapstructure:"rawPacket,omitempty"`

	// The source IP address of the packet.
	SourceIP *string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"sourceip,omitempty" mapstructure:"sourceIP,omitempty"`

	// The source port of the packet.
	SourcePort *int `json:"sourcePort,omitempty" msgpack:"sourcePort,omitempty" bson:"sourceport,omitempty" mapstructure:"sourcePort,omitempty"`

	// The time-date stamp of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"timestamp,omitempty" mapstructure:"timestamp,omitempty"`

	// Set to `true` if the packet arrived with the Trireme options (default).
	TriremePacket *bool `json:"triremePacket,omitempty" msgpack:"triremePacket,omitempty" bson:"triremepacket,omitempty" mapstructure:"triremePacket,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparsePacketReport returns a new  SparsePacketReport.
func NewSparsePacketReport() *SparsePacketReport {
	return &SparsePacketReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparsePacketReport) Identity() elemental.Identity {

	return PacketReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparsePacketReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparsePacketReport) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePacketReport) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparsePacketReport{}

	if o.TCPFlags != nil {
		s.TCPFlags = o.TCPFlags
	}
	if o.Claims != nil {
		s.Claims = o.Claims
	}
	if o.DestinationIP != nil {
		s.DestinationIP = o.DestinationIP
	}
	if o.DestinationPort != nil {
		s.DestinationPort = o.DestinationPort
	}
	if o.DropReason != nil {
		s.DropReason = o.DropReason
	}
	if o.Encrypt != nil {
		s.Encrypt = o.Encrypt
	}
	if o.EnforcerID != nil {
		s.EnforcerID = o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		s.EnforcerNamespace = o.EnforcerNamespace
	}
	if o.Length != nil {
		s.Length = o.Length
	}
	if o.Mark != nil {
		s.Mark = o.Mark
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.PacketID != nil {
		s.PacketID = o.PacketID
	}
	if o.Protocol != nil {
		s.Protocol = o.Protocol
	}
	if o.PuID != nil {
		s.PuID = o.PuID
	}
	if o.RawPacket != nil {
		s.RawPacket = o.RawPacket
	}
	if o.SourceIP != nil {
		s.SourceIP = o.SourceIP
	}
	if o.SourcePort != nil {
		s.SourcePort = o.SourcePort
	}
	if o.Timestamp != nil {
		s.Timestamp = o.Timestamp
	}
	if o.TriremePacket != nil {
		s.TriremePacket = o.TriremePacket
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparsePacketReport) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparsePacketReport{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	if s.TCPFlags != nil {
		o.TCPFlags = s.TCPFlags
	}
	if s.Claims != nil {
		o.Claims = s.Claims
	}
	if s.DestinationIP != nil {
		o.DestinationIP = s.DestinationIP
	}
	if s.DestinationPort != nil {
		o.DestinationPort = s.DestinationPort
	}
	if s.DropReason != nil {
		o.DropReason = s.DropReason
	}
	if s.Encrypt != nil {
		o.Encrypt = s.Encrypt
	}
	if s.EnforcerID != nil {
		o.EnforcerID = s.EnforcerID
	}
	if s.EnforcerNamespace != nil {
		o.EnforcerNamespace = s.EnforcerNamespace
	}
	if s.Length != nil {
		o.Length = s.Length
	}
	if s.Mark != nil {
		o.Mark = s.Mark
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.PacketID != nil {
		o.PacketID = s.PacketID
	}
	if s.Protocol != nil {
		o.Protocol = s.Protocol
	}
	if s.PuID != nil {
		o.PuID = s.PuID
	}
	if s.RawPacket != nil {
		o.RawPacket = s.RawPacket
	}
	if s.SourceIP != nil {
		o.SourceIP = s.SourceIP
	}
	if s.SourcePort != nil {
		o.SourcePort = s.SourcePort
	}
	if s.Timestamp != nil {
		o.Timestamp = s.Timestamp
	}
	if s.TriremePacket != nil {
		o.TriremePacket = s.TriremePacket
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparsePacketReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparsePacketReport) ToPlain() elemental.PlainIdentifiable {

	out := NewPacketReport()
	if o.TCPFlags != nil {
		out.TCPFlags = *o.TCPFlags
	}
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.DestinationIP != nil {
		out.DestinationIP = *o.DestinationIP
	}
	if o.DestinationPort != nil {
		out.DestinationPort = *o.DestinationPort
	}
	if o.DropReason != nil {
		out.DropReason = *o.DropReason
	}
	if o.Encrypt != nil {
		out.Encrypt = *o.Encrypt
	}
	if o.EnforcerID != nil {
		out.EnforcerID = *o.EnforcerID
	}
	if o.EnforcerNamespace != nil {
		out.EnforcerNamespace = *o.EnforcerNamespace
	}
	if o.Event != nil {
		out.Event = *o.Event
	}
	if o.Length != nil {
		out.Length = *o.Length
	}
	if o.Mark != nil {
		out.Mark = *o.Mark
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.PacketID != nil {
		out.PacketID = *o.PacketID
	}
	if o.Protocol != nil {
		out.Protocol = *o.Protocol
	}
	if o.PuID != nil {
		out.PuID = *o.PuID
	}
	if o.RawPacket != nil {
		out.RawPacket = *o.RawPacket
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
	}
	if o.SourcePort != nil {
		out.SourcePort = *o.SourcePort
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.TriremePacket != nil {
		out.TriremePacket = *o.TriremePacket
	}

	return out
}

// DeepCopy returns a deep copy if the SparsePacketReport.
func (o *SparsePacketReport) DeepCopy() *SparsePacketReport {

	if o == nil {
		return nil
	}

	out := &SparsePacketReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparsePacketReport.
func (o *SparsePacketReport) DeepCopyInto(out *SparsePacketReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparsePacketReport: %s", err))
	}

	*out = *target.(*SparsePacketReport)
}

type mongoAttributesPacketReport struct {
	TCPFlags          int       `bson:"tcpflags"`
	Claims            []string  `bson:"claims"`
	DestinationIP     string    `bson:"destinationip"`
	DestinationPort   int       `bson:"destinationport"`
	DropReason        string    `bson:"dropreason"`
	Encrypt           bool      `bson:"encrypt"`
	EnforcerID        string    `bson:"enforcerid"`
	EnforcerNamespace string    `bson:"enforcernamespace"`
	Length            int       `bson:"length"`
	Mark              int       `bson:"mark"`
	Namespace         string    `bson:"namespace"`
	PacketID          int       `bson:"packetid"`
	Protocol          int       `bson:"protocol"`
	PuID              string    `bson:"puid"`
	RawPacket         string    `bson:"rawpacket"`
	SourceIP          string    `bson:"sourceip"`
	SourcePort        int       `bson:"sourceport"`
	Timestamp         time.Time `bson:"timestamp"`
	TriremePacket     bool      `bson:"triremepacket"`
}
type mongoAttributesSparsePacketReport struct {
	TCPFlags          *int       `bson:"tcpflags,omitempty"`
	Claims            *[]string  `bson:"claims,omitempty"`
	DestinationIP     *string    `bson:"destinationip,omitempty"`
	DestinationPort   *int       `bson:"destinationport,omitempty"`
	DropReason        *string    `bson:"dropreason,omitempty"`
	Encrypt           *bool      `bson:"encrypt,omitempty"`
	EnforcerID        *string    `bson:"enforcerid,omitempty"`
	EnforcerNamespace *string    `bson:"enforcernamespace,omitempty"`
	Length            *int       `bson:"length,omitempty"`
	Mark              *int       `bson:"mark,omitempty"`
	Namespace         *string    `bson:"namespace,omitempty"`
	PacketID          *int       `bson:"packetid,omitempty"`
	Protocol          *int       `bson:"protocol,omitempty"`
	PuID              *string    `bson:"puid,omitempty"`
	RawPacket         *string    `bson:"rawpacket,omitempty"`
	SourceIP          *string    `bson:"sourceip,omitempty"`
	SourcePort        *int       `bson:"sourceport,omitempty"`
	Timestamp         *time.Time `bson:"timestamp,omitempty"`
	TriremePacket     *bool      `bson:"triremepacket,omitempty"`
}
