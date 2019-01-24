package gaia

import (
	"fmt"
	"sync"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// TimeSeriesRow represents the model of a timeseriesrow
type TimeSeriesRow struct {
	// colums of the row.
	Columns []string `json:"columns" bson:"-" mapstructure:"columns,omitempty"`

	// the name of row.
	Name string `json:"name" bson:"-" mapstructure:"name,omitempty"`

	// List of tags.
	Tags map[string]string `json:"tags" bson:"-" mapstructure:"tags,omitempty"`

	// List of tags.
	Values [][]interface{} `json:"values" bson:"-" mapstructure:"values,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewTimeSeriesRow returns a new *TimeSeriesRow
func NewTimeSeriesRow() *TimeSeriesRow {

	return &TimeSeriesRow{
		ModelVersion: 1,
		Columns:      []string{},
		Tags:         map[string]string{},
		Values:       [][]interface{}{},
	}
}

// DeepCopy returns a deep copy if the TimeSeriesRow.
func (o *TimeSeriesRow) DeepCopy() *TimeSeriesRow {

	if o == nil {
		return nil
	}

	out := &TimeSeriesRow{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *TimeSeriesRow.
func (o *TimeSeriesRow) DeepCopyInto(out *TimeSeriesRow) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy TimeSeriesRow: %s", err))
	}

	*out = *target.(*TimeSeriesRow)
}

// Validate valides the current information stored into the structure.
func (o *TimeSeriesRow) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
