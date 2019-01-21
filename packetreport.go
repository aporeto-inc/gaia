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
	// PacketReportEventDrop represents the value Drop.
	PacketReportEventDrop PacketReportEventValue = "Drop"

	// PacketReportEventRcv represents the value Rcv.
	PacketReportEventRcv PacketReportEventValue = "Rcv"

	// PacketReportEventTxt represents the value Txt.
	PacketReportEventTxt PacketReportEventValue = "Txt"
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
	// DstPort is the destination port of the packet.
	DstPort int `json:"DstPort" bson:"-" mapstructure:"DstPort,omitempty"`

	// ID is the packet ID from the IP header.
	ID int `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// SrcPort is the source port of the packet.
	SrcPort int `json:"SrcPort" bson:"-" mapstructure:"SrcPort,omitempty"`

	// Claims is the list of claims detected for the packet.
	Claims map[string]string `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// Type of the destination.
	DestinationIP string `json:"destinationIP" bson:"-" mapstructure:"destinationIP,omitempty"`

	// This field is only set if 'action' is set to 'Reject' and specifies the reason
	// for the rejection.
	DropReason string `json:"dropReason" bson:"-" mapstructure:"dropReason,omitempty"`

	// Encrypt indicates that the packet was encrypted.
	Encrypt bool `json:"encrypt" bson:"-" mapstructure:"encrypt,omitempty"`

	// Event is the event that triggered the report.
	Event PacketReportEventValue `json:"event" bson:"-" mapstructure:"event,omitempty"`

	// Flags are the TCP flags of the packet.
	Flags int `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// Length is the length of the packet.
	Length int `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// Mark is the mark value of the packet.
	Mark int `json:"-" bson:"-" mapstructure:"-,omitempty"`

	// Namespace of the PU reporting the packet.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Protocol number.
	Protocol int `json:"protocol" bson:"-" mapstructure:"protocol,omitempty"`

	// ID of the PU reporting the packet.
	PuID string `json:"puID" bson:"-" mapstructure:"puID,omitempty"`

	// Type of the source.
	SourceIP string `json:"sourceIP" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Date of the report.
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
			DstPort:       &o.DstPort,
			ID:            &o.ID,
			SrcPort:       &o.SrcPort,
			Claims:        &o.Claims,
			DestinationIP: &o.DestinationIP,
			DropReason:    &o.DropReason,
			Encrypt:       &o.Encrypt,
			Event:         &o.Event,
			Flags:         &o.Flags,
			Length:        &o.Length,
			Mark:          &o.Mark,
			Namespace:     &o.Namespace,
			Protocol:      &o.Protocol,
			PuID:          &o.PuID,
			SourceIP:      &o.SourceIP,
			Timestamp:     &o.Timestamp,
			TriremePacket: &o.TriremePacket,
		}
	}

	sp := &SparsePacketReport{}
	for _, f := range fields {
		switch f {
		case "DstPort":
			sp.DstPort = &(o.DstPort)
		case "ID":
			sp.ID = &(o.ID)
		case "SrcPort":
			sp.SrcPort = &(o.SrcPort)
		case "claims":
			sp.Claims = &(o.Claims)
		case "destinationIP":
			sp.DestinationIP = &(o.DestinationIP)
		case "dropReason":
			sp.DropReason = &(o.DropReason)
		case "encrypt":
			sp.Encrypt = &(o.Encrypt)
		case "event":
			sp.Event = &(o.Event)
		case "flags":
			sp.Flags = &(o.Flags)
		case "length":
			sp.Length = &(o.Length)
		case "mark":
			sp.Mark = &(o.Mark)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "protocol":
			sp.Protocol = &(o.Protocol)
		case "puID":
			sp.PuID = &(o.PuID)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
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
	if so.DstPort != nil {
		o.DstPort = *so.DstPort
	}
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.SrcPort != nil {
		o.SrcPort = *so.SrcPort
	}
	if so.Claims != nil {
		o.Claims = *so.Claims
	}
	if so.DestinationIP != nil {
		o.DestinationIP = *so.DestinationIP
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
	if so.Flags != nil {
		o.Flags = *so.Flags
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
	if so.Protocol != nil {
		o.Protocol = *so.Protocol
	}
	if so.PuID != nil {
		o.PuID = *so.PuID
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
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

	if err := elemental.ValidateMaximumInt("DstPort", o.DstPort, int(65536), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMaximumInt("SrcPort", o.SrcPort, int(65536), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateStringInList("event", string(o.Event), []string{"Rcv", "Txt", "Drop"}, false); err != nil {
		errors = append(errors, err)
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
	case "DstPort":
		return o.DstPort
	case "ID":
		return o.ID
	case "SrcPort":
		return o.SrcPort
	case "claims":
		return o.Claims
	case "destinationIP":
		return o.DestinationIP
	case "dropReason":
		return o.DropReason
	case "encrypt":
		return o.Encrypt
	case "event":
		return o.Event
	case "flags":
		return o.Flags
	case "length":
		return o.Length
	case "mark":
		return o.Mark
	case "namespace":
		return o.Namespace
	case "protocol":
		return o.Protocol
	case "puID":
		return o.PuID
	case "sourceIP":
		return o.SourceIP
	case "timestamp":
		return o.Timestamp
	case "triremePacket":
		return o.TriremePacket
	}

	return nil
}

// PacketReportAttributesMap represents the map of attribute for PacketReport.
var PacketReportAttributesMap = map[string]elemental.AttributeSpecification{
	"DstPort": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DstPort",
		Description:    `DstPort is the destination port of the packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "DstPort",
		Type:           "integer",
	},
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the packet ID from the IP header.`,
		Name:           "ID",
		Required:       true,
		Type:           "integer",
	},
	"SrcPort": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SrcPort",
		Description:    `SrcPort is the source port of the packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "SrcPort",
		Type:           "integer",
	},
	"Claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims is the list of claims detected for the packet.`,
		Name:           "claims",
		SubType:        "enforcer_claims",
		Type:           "external",
	},
	"DestinationIP": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Type of the destination.`,
		Exposed:        true,
		Name:           "destinationIP",
		Type:           "string",
	},
	"DropReason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if 'action' is set to 'Reject' and specifies the reason
for the rejection.`,
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
		Required:       true,
		Type:           "boolean",
	},
	"Event": elemental.AttributeSpecification{
		AllowedChoices: []string{"Rcv", "Txt", "Drop"},
		ConvertedName:  "Event",
		Description:    `Event is the event that triggered the report.`,
		Exposed:        true,
		Name:           "event",
		Required:       true,
		Type:           "enum",
	},
	"Flags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Flags",
		Description:    `Flags are the TCP flags of the packet.`,
		Name:           "flags",
		Type:           "integer",
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
		Type:           "string",
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
		Description:    `ID of the PU reporting the packet.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "puID",
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
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
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
	"dstport": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DstPort",
		Description:    `DstPort is the destination port of the packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "DstPort",
		Type:           "integer",
	},
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID is the packet ID from the IP header.`,
		Name:           "ID",
		Required:       true,
		Type:           "integer",
	},
	"srcport": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SrcPort",
		Description:    `SrcPort is the source port of the packet.`,
		Exposed:        true,
		MaxValue:       65536,
		Name:           "SrcPort",
		Type:           "integer",
	},
	"claims": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Claims",
		Description:    `Claims is the list of claims detected for the packet.`,
		Name:           "claims",
		SubType:        "enforcer_claims",
		Type:           "external",
	},
	"destinationip": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Type of the destination.`,
		Exposed:        true,
		Name:           "destinationIP",
		Type:           "string",
	},
	"dropreason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if 'action' is set to 'Reject' and specifies the reason
for the rejection.`,
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
		Required:       true,
		Type:           "boolean",
	},
	"event": elemental.AttributeSpecification{
		AllowedChoices: []string{"Rcv", "Txt", "Drop"},
		ConvertedName:  "Event",
		Description:    `Event is the event that triggered the report.`,
		Exposed:        true,
		Name:           "event",
		Required:       true,
		Type:           "enum",
	},
	"flags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Flags",
		Description:    `Flags are the TCP flags of the packet.`,
		Name:           "flags",
		Type:           "integer",
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
		Type:           "string",
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
		Description:    `ID of the PU reporting the packet.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "puID",
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
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
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
	// DstPort is the destination port of the packet.
	DstPort *int `json:"DstPort,omitempty" bson:"-" mapstructure:"DstPort,omitempty"`

	// ID is the packet ID from the IP header.
	ID *int `json:"-,omitempty" bson:"-" mapstructure:"-,omitempty"`

	// SrcPort is the source port of the packet.
	SrcPort *int `json:"SrcPort,omitempty" bson:"-" mapstructure:"SrcPort,omitempty"`

	// Claims is the list of claims detected for the packet.
	Claims *map[string]string `json:"-,omitempty" bson:"-" mapstructure:"-,omitempty"`

	// Type of the destination.
	DestinationIP *string `json:"destinationIP,omitempty" bson:"-" mapstructure:"destinationIP,omitempty"`

	// This field is only set if 'action' is set to 'Reject' and specifies the reason
	// for the rejection.
	DropReason *string `json:"dropReason,omitempty" bson:"-" mapstructure:"dropReason,omitempty"`

	// Encrypt indicates that the packet was encrypted.
	Encrypt *bool `json:"encrypt,omitempty" bson:"-" mapstructure:"encrypt,omitempty"`

	// Event is the event that triggered the report.
	Event *PacketReportEventValue `json:"event,omitempty" bson:"-" mapstructure:"event,omitempty"`

	// Flags are the TCP flags of the packet.
	Flags *int `json:"-,omitempty" bson:"-" mapstructure:"-,omitempty"`

	// Length is the length of the packet.
	Length *int `json:"-,omitempty" bson:"-" mapstructure:"-,omitempty"`

	// Mark is the mark value of the packet.
	Mark *int `json:"-,omitempty" bson:"-" mapstructure:"-,omitempty"`

	// Namespace of the PU reporting the packet.
	Namespace *string `json:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// Protocol number.
	Protocol *int `json:"protocol,omitempty" bson:"-" mapstructure:"protocol,omitempty"`

	// ID of the PU reporting the packet.
	PuID *string `json:"puID,omitempty" bson:"-" mapstructure:"puID,omitempty"`

	// Type of the source.
	SourceIP *string `json:"sourceIP,omitempty" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Date of the report.
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
	if o.DstPort != nil {
		out.DstPort = *o.DstPort
	}
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.SrcPort != nil {
		out.SrcPort = *o.SrcPort
	}
	if o.Claims != nil {
		out.Claims = *o.Claims
	}
	if o.DestinationIP != nil {
		out.DestinationIP = *o.DestinationIP
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
	if o.Flags != nil {
		out.Flags = *o.Flags
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
	if o.Protocol != nil {
		out.Protocol = *o.Protocol
	}
	if o.PuID != nil {
		out.PuID = *o.PuID
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
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
