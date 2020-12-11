package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ReportsQueryResults represents the model of a reportsqueryresults
type ReportsQueryResults struct {
	// List of DNSLookupReports.
	DNSLookupReports DNSLookupReportsList `json:"DNSLookupReports,omitempty" msgpack:"DNSLookupReports,omitempty" bson:"-" mapstructure:"DNSLookupReports,omitempty"`

	// List of CounterReports.
	CounterReports CounterReportsList `json:"counterReports,omitempty" msgpack:"counterReports,omitempty" bson:"-" mapstructure:"counterReports,omitempty"`

	// List of EnforcerReports.
	EnforcerReports EnforcerReportsList `json:"enforcerReports,omitempty" msgpack:"enforcerReports,omitempty" bson:"-" mapstructure:"enforcerReports,omitempty"`

	// List of EventLogs.
	EventLogs EventLogsList `json:"eventLogs,omitempty" msgpack:"eventLogs,omitempty" bson:"-" mapstructure:"eventLogs,omitempty"`

	// List of FlowReports.
	FlowReports FlowReportsList `json:"flowReports,omitempty" msgpack:"flowReports,omitempty" bson:"-" mapstructure:"flowReports,omitempty"`

	// List of PacketReports.
	PacketReports PacketReportsList `json:"packetReports,omitempty" msgpack:"packetReports,omitempty" bson:"-" mapstructure:"packetReports,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewReportsQueryResults returns a new *ReportsQueryResults
func NewReportsQueryResults() *ReportsQueryResults {

	return &ReportsQueryResults{
		ModelVersion:     1,
		CounterReports:   CounterReportsList{},
		DNSLookupReports: DNSLookupReportsList{},
		EnforcerReports:  EnforcerReportsList{},
		EventLogs:        EventLogsList{},
		FlowReports:      FlowReportsList{},
		PacketReports:    PacketReportsList{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ReportsQueryResults) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesReportsQueryResults{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ReportsQueryResults) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesReportsQueryResults{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *ReportsQueryResults) BleveType() string {

	return "reportsqueryresults"
}

// DeepCopy returns a deep copy if the ReportsQueryResults.
func (o *ReportsQueryResults) DeepCopy() *ReportsQueryResults {

	if o == nil {
		return nil
	}

	out := &ReportsQueryResults{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *ReportsQueryResults.
func (o *ReportsQueryResults) DeepCopyInto(out *ReportsQueryResults) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy ReportsQueryResults: %s", err))
	}

	*out = *target.(*ReportsQueryResults)
}

// Validate valides the current information stored into the structure.
func (o *ReportsQueryResults) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.DNSLookupReports {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.CounterReports {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.EnforcerReports {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.EventLogs {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.FlowReports {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	for _, sub := range o.PacketReports {
		if sub == nil {
			continue
		}
		elemental.ResetDefaultForZeroValues(sub)
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
		}
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type mongoAttributesReportsQueryResults struct {
}
