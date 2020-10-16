package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// MetricsIdentity represents the Identity of the object.
var MetricsIdentity = elemental.Identity{
	Name:     "metrics",
	Category: "metrics",
	Package:  "jenova",
	Private:  false,
}

// MetricsList represents a list of Metrics
type MetricsList []*Metrics

// Identity returns the identity of the objects in the list.
func (o MetricsList) Identity() elemental.Identity {

	return MetricsIdentity
}

// Copy returns a pointer to a copy the MetricsList.
func (o MetricsList) Copy() elemental.Identifiables {

	copy := append(MetricsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the MetricsList.
func (o MetricsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(MetricsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*Metrics))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o MetricsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o MetricsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the MetricsList converted to SparseMetricsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o MetricsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseMetricsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseMetrics)
	}

	return out
}

// Version returns the version of the content.
func (o MetricsList) Version() int {

	return 1
}

// Metrics represents the model of a metrics
type Metrics struct {
	// End timestamp <rfc3339 | unix_timestamp>.
	End string `json:"end" msgpack:"end" bson:"-" mapstructure:"end,omitempty"`

	// Contains the remote `POST` payload.
	Query string `json:"query" msgpack:"query" bson:"-" mapstructure:"query,omitempty"`

	// Start timestamp <rfc3339 | unix_timestamp>.
	Start string `json:"start" msgpack:"start" bson:"-" mapstructure:"start,omitempty"`

	// Query resolution step width in duration format or float number of seconds.
	Step string `json:"step" msgpack:"step" bson:"-" mapstructure:"step,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewMetrics returns a new *Metrics
func NewMetrics() *Metrics {

	return &Metrics{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Metrics) Identity() elemental.Identity {

	return MetricsIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Metrics) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Metrics) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Metrics) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesMetrics{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Metrics) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesMetrics{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Metrics) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Metrics) BleveType() string {

	return "metrics"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Metrics) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Metrics) Doc() string {

	return `Prometheus compatible endpoint to evaluate an expression query over a range of
time. This can be used to retrieve back Aporeto specific metrics for a given
namespace. All queries are protected within the namespace of the caller.`
}

func (o *Metrics) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Metrics) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseMetrics{
			End:   &o.End,
			Query: &o.Query,
			Start: &o.Start,
			Step:  &o.Step,
		}
	}

	sp := &SparseMetrics{}
	for _, f := range fields {
		switch f {
		case "end":
			sp.End = &(o.End)
		case "query":
			sp.Query = &(o.Query)
		case "start":
			sp.Start = &(o.Start)
		case "step":
			sp.Step = &(o.Step)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseMetrics to the object.
func (o *Metrics) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseMetrics)
	if so.End != nil {
		o.End = *so.End
	}
	if so.Query != nil {
		o.Query = *so.Query
	}
	if so.Start != nil {
		o.Start = *so.Start
	}
	if so.Step != nil {
		o.Step = *so.Step
	}
}

// DeepCopy returns a deep copy if the Metrics.
func (o *Metrics) DeepCopy() *Metrics {

	if o == nil {
		return nil
	}

	out := &Metrics{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Metrics.
func (o *Metrics) DeepCopyInto(out *Metrics) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Metrics: %s", err))
	}

	*out = *target.(*Metrics)
}

// Validate valides the current information stored into the structure.
func (o *Metrics) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("query", o.Query); err != nil {
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
func (*Metrics) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := MetricsAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return MetricsLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Metrics) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return MetricsAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Metrics) ValueForAttribute(name string) interface{} {

	switch name {
	case "end":
		return o.End
	case "query":
		return o.Query
	case "start":
		return o.Start
	case "step":
		return o.Step
	}

	return nil
}

// MetricsAttributesMap represents the map of attribute for Metrics.
var MetricsAttributesMap = map[string]elemental.AttributeSpecification{
	"End": {
		AllowedChoices: []string{},
		ConvertedName:  "End",
		Description:    `End timestamp <rfc3339 | unix_timestamp>.`,
		Exposed:        true,
		Name:           "end",
		Type:           "string",
	},
	"Query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `Contains the remote ` + "`" + `POST` + "`" + ` payload.`,
		Exposed:        true,
		Name:           "query",
		Required:       true,
		Type:           "string",
	},
	"Start": {
		AllowedChoices: []string{},
		ConvertedName:  "Start",
		Description:    `Start timestamp <rfc3339 | unix_timestamp>.`,
		Exposed:        true,
		Name:           "start",
		Type:           "string",
	},
	"Step": {
		AllowedChoices: []string{},
		ConvertedName:  "Step",
		Description:    `Query resolution step width in duration format or float number of seconds.`,
		Exposed:        true,
		Name:           "step",
		Type:           "string",
	},
}

// MetricsLowerCaseAttributesMap represents the map of attribute for Metrics.
var MetricsLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"end": {
		AllowedChoices: []string{},
		ConvertedName:  "End",
		Description:    `End timestamp <rfc3339 | unix_timestamp>.`,
		Exposed:        true,
		Name:           "end",
		Type:           "string",
	},
	"query": {
		AllowedChoices: []string{},
		ConvertedName:  "Query",
		Description:    `Contains the remote ` + "`" + `POST` + "`" + ` payload.`,
		Exposed:        true,
		Name:           "query",
		Required:       true,
		Type:           "string",
	},
	"start": {
		AllowedChoices: []string{},
		ConvertedName:  "Start",
		Description:    `Start timestamp <rfc3339 | unix_timestamp>.`,
		Exposed:        true,
		Name:           "start",
		Type:           "string",
	},
	"step": {
		AllowedChoices: []string{},
		ConvertedName:  "Step",
		Description:    `Query resolution step width in duration format or float number of seconds.`,
		Exposed:        true,
		Name:           "step",
		Type:           "string",
	},
}

// SparseMetricsList represents a list of SparseMetrics
type SparseMetricsList []*SparseMetrics

// Identity returns the identity of the objects in the list.
func (o SparseMetricsList) Identity() elemental.Identity {

	return MetricsIdentity
}

// Copy returns a pointer to a copy the SparseMetricsList.
func (o SparseMetricsList) Copy() elemental.Identifiables {

	copy := append(SparseMetricsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseMetricsList.
func (o SparseMetricsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseMetricsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseMetrics))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseMetricsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseMetricsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseMetricsList converted to MetricsList.
func (o SparseMetricsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseMetricsList) Version() int {

	return 1
}

// SparseMetrics represents the sparse version of a metrics.
type SparseMetrics struct {
	// End timestamp <rfc3339 | unix_timestamp>.
	End *string `json:"end,omitempty" msgpack:"end,omitempty" bson:"-" mapstructure:"end,omitempty"`

	// Contains the remote `POST` payload.
	Query *string `json:"query,omitempty" msgpack:"query,omitempty" bson:"-" mapstructure:"query,omitempty"`

	// Start timestamp <rfc3339 | unix_timestamp>.
	Start *string `json:"start,omitempty" msgpack:"start,omitempty" bson:"-" mapstructure:"start,omitempty"`

	// Query resolution step width in duration format or float number of seconds.
	Step *string `json:"step,omitempty" msgpack:"step,omitempty" bson:"-" mapstructure:"step,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseMetrics returns a new  SparseMetrics.
func NewSparseMetrics() *SparseMetrics {
	return &SparseMetrics{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseMetrics) Identity() elemental.Identity {

	return MetricsIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseMetrics) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseMetrics) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseMetrics) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseMetrics{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseMetrics) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseMetrics{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseMetrics) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseMetrics) ToPlain() elemental.PlainIdentifiable {

	out := NewMetrics()
	if o.End != nil {
		out.End = *o.End
	}
	if o.Query != nil {
		out.Query = *o.Query
	}
	if o.Start != nil {
		out.Start = *o.Start
	}
	if o.Step != nil {
		out.Step = *o.Step
	}

	return out
}

// DeepCopy returns a deep copy if the SparseMetrics.
func (o *SparseMetrics) DeepCopy() *SparseMetrics {

	if o == nil {
		return nil
	}

	out := &SparseMetrics{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseMetrics.
func (o *SparseMetrics) DeepCopyInto(out *SparseMetrics) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseMetrics: %s", err))
	}

	*out = *target.(*SparseMetrics)
}

type mongoAttributesMetrics struct {
}
type mongoAttributesSparseMetrics struct {
}
