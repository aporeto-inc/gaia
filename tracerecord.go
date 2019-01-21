package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TraceRecord represents the model of a tracerecord
type TraceRecord struct {
	// ID is the packet ID.
	ID int `json:"ID" bson:"id" mapstructure:"ID,omitempty"`

	// Length of the observed packet.
	Length int `json:"Length" bson:"length" mapstructure:"Length,omitempty"`

	// TTL is the TTL value of the packet.
	TTL int `json:"TTL" bson:"ttl" mapstructure:"TTL,omitempty"`

	// Chain is the chain that the trace was collected from.
	Chain string `json:"chain" bson:"chain" mapstructure:"chain,omitempty"`

	// DestIP is the destination IP.
	DestIP string `json:"destIP" bson:"destip" mapstructure:"destIP,omitempty"`

	// DestInterface is the destination interface of the packet.
	DestInterface string `json:"destInterface" bson:"destinterface" mapstructure:"destInterface,omitempty"`

	// DstPort is the destination Port.
	DstPort int `json:"dstPort" bson:"dstport" mapstructure:"dstPort,omitempty"`

	// Policy is the index of the iptables entry that was hit.
	Policy int `json:"policy" bson:"policy" mapstructure:"policy,omitempty"`

	// Protocol is the protocol of the packets.
	Protocol int `json:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	// SrcIP is the source IP.
	SrcIP string `json:"srcIP" bson:"srcip" mapstructure:"srcIP,omitempty"`

	// SrcInterface is the source interface of the packet.
	SrcInterface string `json:"srcInterface" bson:"srcinterface" mapstructure:"srcInterface,omitempty"`

	// SrcPort is the source Port.
	SrcPort int `json:"srcPort" bson:"srcport" mapstructure:"srcPort,omitempty"`

	// Table is the iptable that the trace was collected.
	Table string `json:"table" bson:"table" mapstructure:"table,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewTraceRecord returns a new *TraceRecord
func NewTraceRecord() *TraceRecord {

	return &TraceRecord{
		ModelVersion: 1,
	}
}

// DeepCopy returns a deep copy if the TraceRecord.
func (o *TraceRecord) DeepCopy() *TraceRecord {

	if o == nil {
		return nil
	}

	out := &TraceRecord{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TraceRecord.
func (o *TraceRecord) DeepCopyInto(out *TraceRecord) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TraceRecord: %s", err))
	}

	*out = *target.(*TraceRecord)
}

// Validate valides the current information stored into the structure.
func (o *TraceRecord) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredInt("ID", o.ID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredInt("Length", o.Length); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("Length", o.Length, int(65536), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredInt("TTL", o.TTL); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("TTL", o.TTL, int(255), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("chain", o.Chain); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("destIP", o.DestIP); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredInt("dstPort", o.DstPort); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("dstPort", o.DstPort, int(65536), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredInt("policy", o.Policy); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredInt("protocol", o.Protocol); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("protocol", o.Protocol, int(65536), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("srcIP", o.SrcIP); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredInt("srcPort", o.SrcPort); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("srcPort", o.SrcPort, int(65536), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("table", o.Table); err != nil {
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
