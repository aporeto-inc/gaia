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
	// TTL is the TTL value of the packet.
	TTL int `json:"TTL" bson:"ttl" mapstructure:"TTL,omitempty"`

	// Chain is the chain that the trace was collected from.
	Chain string `json:"chain" bson:"chain" mapstructure:"chain,omitempty"`

	// DestIP is the destination IP.
	DestinationIP string `json:"destinationIP" bson:"destinationip" mapstructure:"destinationIP,omitempty"`

	// DestInterface is the destination interface of the packet.
	DestinationInterface string `json:"destinationInterface" bson:"destinationinterface" mapstructure:"destinationInterface,omitempty"`

	// DestPort is the destination Port.
	DestinationPort int `json:"destinationPort" bson:"destinationport" mapstructure:"destinationPort,omitempty"`

	// Length of the observed packet.
	Length int `json:"length" bson:"length" mapstructure:"length,omitempty"`

	// ID is the packet ID.
	PacketID int `json:"packetID" bson:"packetid" mapstructure:"packetID,omitempty"`

	// Protocol is the protocol of the packets.
	Protocol int `json:"protocol" bson:"protocol" mapstructure:"protocol,omitempty"`

	// ruleID is the index of the iptables entry that was hit.
	RuleID int `json:"ruleID" bson:"ruleid" mapstructure:"ruleID,omitempty"`

	// SrcIP is the source IP.
	SourceIP string `json:"sourceIP" bson:"sourceip" mapstructure:"sourceIP,omitempty"`

	// SrcInterface is the source interface of the packet.
	SourceInterface string `json:"sourceInterface" bson:"sourceinterface" mapstructure:"sourceInterface,omitempty"`

	// SrcPort is the source Port.
	SourcePort int `json:"sourcePort" bson:"sourceport" mapstructure:"sourcePort,omitempty"`

	// TableName is the iptable that the trace was collected.
	TableName string `json:"tableName" bson:"tablename" mapstructure:"tableName,omitempty"`

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

	if err := elemental.ValidateRequiredInt("TTL", o.TTL); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("TTL", o.TTL, int(255), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("chain", o.Chain); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("destinationIP", o.DestinationIP); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredInt("destinationPort", o.DestinationPort); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("destinationPort", o.DestinationPort, int(65536), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("destinationPort", o.DestinationPort, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredInt("length", o.Length); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("length", o.Length, int(65536), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredInt("packetID", o.PacketID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredInt("protocol", o.Protocol); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("protocol", o.Protocol, int(65536), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredInt("ruleID", o.RuleID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("sourceIP", o.SourceIP); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredInt("sourcePort", o.SourcePort); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateMaximumInt("sourcePort", o.SourcePort, int(65536), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateMinimumInt("sourcePort", o.SourcePort, int(1), false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("tableName", o.TableName); err != nil {
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
