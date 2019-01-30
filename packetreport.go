package gaia

import (
	"fmt"
	"sync"
	"time"

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
func (o PacketReportsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
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
	TCPFlags int `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// Claims is the list of claims detected for the packet.
	Claims map[string]string `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// DestinationIP is the IP address of the destination.
	DestinationIP string `json:"destinationIP" bson:"-" mapstructure:"destinationIP,omitempty"`

	// DestinationPort is the destination port of a TCP or UDP packet.
	DestinationPort int `json:"destinationPort" bson:"-" mapstructure:"destinationPort,omitempty"`

	// This field is only set if 'event' is set to 'Dropped' and specifies the reason
	// for the drop.
	DropReason string `json:"dropReason" bson:"-" mapstructure:"dropReason,omitempty"`

	// Encrypt indicates that the packet was encrypted.
	Encrypt bool `json:"encrypt" bson:"-" mapstructure:"encrypt,omitempty"`

	// Event is the event that triggered the report.
	Event PacketReportEventValue `json:"event" bson:"-" mapstructure:"event,omitempty"`

	// Length is the length of the packet.
	Length int `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// Mark is the mark value of the packet.
	Mark int `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// Namespace of the PU reporting the packet.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// PacketID is the ID from the IP header of the packet.
	PacketID int `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// Protocol number.
	Protocol int `json:"protocol" bson:"-" mapstructure:"protocol,omitempty"`

	// PUID is the ID of the PU reporting the packet.
	PuID string `json:"puID" bson:"-" mapstructure:"puID,omitempty"`

	// SourceIP is the source IP address of the packet.
	SourceIP string `json:"sourceIP" bson:"-" mapstructure:"sourceIP,omitempty"`

	// SourcePort is the source port of the packet.
	SourcePort int `json:"sourcePort" bson:"-" mapstructure:"sourcePort,omitempty"`

	// Timestamp is the date of the report.
	Timestamp time.Time `json:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	// TriremePacket is set if the packet arrived with the Trireme options.
	TriremePacket bool `json:"triremePacket" bson:"triremepacket" mapstructure:"triremePacket,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewPacketReport returns a new *PacketReport
func NewPacketReport() *PacketReport {

	return &PacketReport{
		ModelVersion:  1,
		Claims:        map[string]string{},
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

// Version returns the hardcoded version of the model.
func (o *PacketReport) Version() int {

	return 1
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
			TCPFlags:        &o.TCPFlags,
			Claims:          &o.Claims,
			DestinationIP:   &o.DestinationIP,
			DestinationPort: &o.DestinationPort,
			DropReason:      &o.DropReason,
			Encrypt:         &o.Encrypt,
			Event:           &o.Event,
			Length:          &o.Length,
			Mark:            &o.Mark,
			Namespace:       &o.Namespace,
			PacketID:        &o.PacketID,
			Protocol:        &o.Protocol,
			PuID:            &o.PuID,
			SourceIP:        &o.SourceIP,
			SourcePort:      &o.SourcePort,
			Timestamp:       &o.Timestamp,
			TriremePacket:   &o.TriremePacket,
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
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("event", string(o.Event), []string{"Received", "Transmitted", "Dropped"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredInt("protocol", o.Protocol); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("protocol", o.Protocol, int(255), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("puID", o.PuID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("sourcePort", o.SourcePort, int(65536), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredTime("timestamp", o.Timestamp); err != nil {
		requiredErrors = append(requiredErrors, err)
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
	"TCPFlags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TCPFlags",
		Description:    `Flags are the TCP flags of the packet.`,
		Name:           "TCPFlags",
		Type:           "integer",
	},
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims is the list of claims detected for the packet.`,
		Name:           "claims",
		SubType:        "map_of_string_of_strings",
		Type:           "external",
	},
	"DestinationIP": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `DestinationIP is the IP address of the destination.`,
		Exposed:        true,
		Name:           "destinationIP",
		Type:           "string",
	},
	"DestinationPort": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `DestinationPort is the destination port of a TCP or UDP packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "destinationPort",
		Type:           "integer",
	},
	"DropReason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if 'event' is set to 'Dropped' and specifies the reason
for the drop.`,
		Exposed: true,
		Name:    "dropReason",
		Type:    "string",
	},
	"Encrypt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Encrypt",
		Description:    `Encrypt indicates that the packet was encrypted.`,
		Exposed:        true,
		Name:           "encrypt",
		Type:           "boolean",
	},
	"Event": elemental.AttributeSpecification{
		AllowedChoices: []string{"Received", "Transmitted", "Dropped"},
		ConvertedName:  "Event",
		Description:    `Event is the event that triggered the report.`,
		Exposed:        true,
		Name:           "event",
		Required:       true,
		Type:           "enum",
	},
	"Length": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Length",
		Description:    `Length is the length of the packet.`,
		MaxValue:       65536,
		Name:           "length",
		Type:           "integer",
	},
	"Mark": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Mark",
		Description:    `Mark is the mark value of the packet.`,
		Name:           "mark",
		Required:       true,
		Type:           "integer",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the PU reporting the packet.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"PacketID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PacketID",
		Description:    `PacketID is the ID from the IP header of the packet.`,
		Name:           "packetID",
		Required:       true,
		Type:           "integer",
	},
	"Protocol": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol number.`,
		Exposed:        true,
		MaxValue:       255,
		Name:           "protocol",
		Required:       true,
		Type:           "integer",
	},
	"PuID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PuID",
		Description:    `PUID is the ID of the PU reporting the packet.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "puID",
		Required:       true,
		Type:           "string",
	},
	"SourceIP": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `SourceIP is the source IP address of the packet.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"SourcePort": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourcePort",
		Description:    `SourcePort is the source port of the packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "sourcePort",
		Type:           "integer",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp is the date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
	"TriremePacket": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TriremePacket",
		DefaultValue:   true,
		Description:    `TriremePacket is set if the packet arrived with the Trireme options.`,
		Exposed:        true,
		Name:           "triremePacket",
		Required:       true,
		Stored:         true,
		Type:           "boolean",
	},
}

// PacketReportLowerCaseAttributesMap represents the map of attribute for PacketReport.
var PacketReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"tcpflags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TCPFlags",
		Description:    `Flags are the TCP flags of the packet.`,
		Name:           "TCPFlags",
		Type:           "integer",
	},
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims is the list of claims detected for the packet.`,
		Name:           "claims",
		SubType:        "map_of_string_of_strings",
		Type:           "external",
	},
	"destinationip": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `DestinationIP is the IP address of the destination.`,
		Exposed:        true,
		Name:           "destinationIP",
		Type:           "string",
	},
	"destinationport": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `DestinationPort is the destination port of a TCP or UDP packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "destinationPort",
		Type:           "integer",
	},
	"dropreason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if 'event' is set to 'Dropped' and specifies the reason
for the drop.`,
		Exposed: true,
		Name:    "dropReason",
		Type:    "string",
	},
	"encrypt": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Encrypt",
		Description:    `Encrypt indicates that the packet was encrypted.`,
		Exposed:        true,
		Name:           "encrypt",
		Type:           "boolean",
	},
	"event": elemental.AttributeSpecification{
		AllowedChoices: []string{"Received", "Transmitted", "Dropped"},
		ConvertedName:  "Event",
		Description:    `Event is the event that triggered the report.`,
		Exposed:        true,
		Name:           "event",
		Required:       true,
		Type:           "enum",
	},
	"length": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Length",
		Description:    `Length is the length of the packet.`,
		MaxValue:       65536,
		Name:           "length",
		Type:           "integer",
	},
	"mark": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Mark",
		Description:    `Mark is the mark value of the packet.`,
		Name:           "mark",
		Required:       true,
		Type:           "integer",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the PU reporting the packet.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"packetid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PacketID",
		Description:    `PacketID is the ID from the IP header of the packet.`,
		Name:           "packetID",
		Required:       true,
		Type:           "integer",
	},
	"protocol": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol number.`,
		Exposed:        true,
		MaxValue:       255,
		Name:           "protocol",
		Required:       true,
		Type:           "integer",
	},
	"puid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PuID",
		Description:    `PUID is the ID of the PU reporting the packet.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "puID",
		Required:       true,
		Type:           "string",
	},
	"sourceip": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `SourceIP is the source IP address of the packet.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"sourceport": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourcePort",
		Description:    `SourcePort is the source port of the packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "sourcePort",
		Type:           "integer",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Timestamp is the date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
	"triremepacket": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "TriremePacket",
		DefaultValue:   true,
		Description:    `TriremePacket is set if the packet arrived with the Trireme options.`,
		Exposed:        true,
		Name:           "triremePacket",
		Required:       true,
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
	TCPFlags *int `json:"-,omitempty" bson:"-" mapstructure:"-,omitempty"`

	// Claims is the list of claims detected for the packet.
	Claims *map[string]string `json:"-,omitempty" bson:"-" mapstructure:"-,omitempty"`

	// DestinationIP is the IP address of the destination.
	DestinationIP *string `json:"destinationIP,omitempty" bson:"-" mapstructure:"destinationIP,omitempty"`

	// DestinationPort is the destination port of a TCP or UDP packet.
	DestinationPort *int `json:"destinationPort,omitempty" bson:"-" mapstructure:"destinationPort,omitempty"`

	// This field is only set if 'event' is set to 'Dropped' and specifies the reason
	// for the drop.
	DropReason *string `json:"dropReason,omitempty" bson:"-" mapstructure:"dropReason,omitempty"`

	// Encrypt indicates that the packet was encrypted.
	Encrypt *bool `json:"encrypt,omitempty" bson:"-" mapstructure:"encrypt,omitempty"`

	// Event is the event that triggered the report.
	Event *PacketReportEventValue `json:"event,omitempty" bson:"-" mapstructure:"event,omitempty"`

	// Length is the length of the packet.
	Length *int `json:"-,omitempty" bson:"-" mapstructure:"-,omitempty"`

	// Mark is the mark value of the packet.
	Mark *int `json:"-,omitempty" bson:"-" mapstructure:"-,omitempty"`

	// Namespace of the PU reporting the packet.
	Namespace *string `json:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// PacketID is the ID from the IP header of the packet.
	PacketID *int `json:"-,omitempty" bson:"-" mapstructure:"-,omitempty"`

	// Protocol number.
	Protocol *int `json:"protocol,omitempty" bson:"-" mapstructure:"protocol,omitempty"`

	// PUID is the ID of the PU reporting the packet.
	PuID *string `json:"puID,omitempty" bson:"-" mapstructure:"puID,omitempty"`

	// SourceIP is the source IP address of the packet.
	SourceIP *string `json:"sourceIP,omitempty" bson:"-" mapstructure:"sourceIP,omitempty"`

	// SourcePort is the source port of the packet.
	SourcePort *int `json:"sourcePort,omitempty" bson:"-" mapstructure:"sourcePort,omitempty"`

	// Timestamp is the date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	// TriremePacket is set if the packet arrived with the Trireme options.
	TriremePacket *bool `json:"triremePacket,omitempty" bson:"triremepacket" mapstructure:"triremePacket,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
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
