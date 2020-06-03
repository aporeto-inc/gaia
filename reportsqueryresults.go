package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// ReportsQueryResults represents the model of a reportsqueryresults
type ReportsQueryResults struct {
	// List of projected fields.
	Fields []string `json:"fields" msgpack:"fields" bson:"fields" mapstructure:"fields,omitempty"`

	// List of projected fields.
	Groups map[string]interface{} `json:"groups" msgpack:"groups" bson:"groups" mapstructure:"groups,omitempty"`

	// List of values associated with the projected fields.
	Values [][]interface{} `json:"values" msgpack:"values" bson:"values" mapstructure:"values,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewReportsQueryResults returns a new *ReportsQueryResults
func NewReportsQueryResults() *ReportsQueryResults {

	return &ReportsQueryResults{
		ModelVersion: 1,
		Fields:       []string{},
		Groups:       map[string]interface{}{},
		Values:       [][]interface{}{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *ReportsQueryResults) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesReportsQueryResults{}

	s.Fields = o.Fields
	s.Groups = o.Groups
	s.Values = o.Values

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

	o.Fields = s.Fields
	o.Groups = s.Groups
	o.Values = s.Values

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

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type mongoAttributesReportsQueryResults struct {
	Fields []string               `bson:"fields"`
	Groups map[string]interface{} `bson:"groups"`
	Values [][]interface{}        `bson:"values"`
}
