package gaia

import (
	"fmt"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TimeSeriesQueryResults represents the model of a timeseriesqueryresults
type TimeSeriesQueryResults struct {
	// List of rows.
	Rows []*TimeSeriesRow `json:"rows" msgpack:"rows" bson:"-" mapstructure:"rows,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTimeSeriesQueryResults returns a new *TimeSeriesQueryResults
func NewTimeSeriesQueryResults() *TimeSeriesQueryResults {

	return &TimeSeriesQueryResults{
		ModelVersion: 1,
		Rows:         []*TimeSeriesRow{},
	}
}

// DeepCopy returns a deep copy if the TimeSeriesQueryResults.
func (o *TimeSeriesQueryResults) DeepCopy() *TimeSeriesQueryResults {

	if o == nil {
		return nil
	}

	out := &TimeSeriesQueryResults{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TimeSeriesQueryResults.
func (o *TimeSeriesQueryResults) DeepCopyInto(out *TimeSeriesQueryResults) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TimeSeriesQueryResults: %s", err))
	}

	*out = *target.(*TimeSeriesQueryResults)
}

// Validate valides the current information stored into the structure.
func (o *TimeSeriesQueryResults) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Rows {
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
