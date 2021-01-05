package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// RouteData represents the model of a routedata
type RouteData struct {
	// The gateway that this route table is associated with.
	GatewayID string `json:"gatewayID" msgpack:"gatewayID" bson:"gatewayid" mapstructure:"gatewayID,omitempty"`

	// Indicates that this is the default route table for the VPC.
	MainTable bool `json:"mainTable" msgpack:"mainTable" bson:"maintable" mapstructure:"mainTable,omitempty"`

	// Routes associated with this route table.
	Routelist CloudRoutesList `json:"routelist" msgpack:"routelist" bson:"routelist" mapstructure:"routelist,omitempty"`

	// The list of subnets that this route table is associated with.
	SubnetAssociations []string `json:"subnetAssociations" msgpack:"subnetAssociations" bson:"subnetassociations" mapstructure:"subnetAssociations,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRouteData returns a new *RouteData
func NewRouteData() *RouteData {

	return &RouteData{
		ModelVersion:       1,
		Routelist:          CloudRoutesList{},
		SubnetAssociations: []string{},
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RouteData) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRouteData{}

	s.GatewayID = o.GatewayID
	s.MainTable = o.MainTable
	s.Routelist = o.Routelist
	s.SubnetAssociations = o.SubnetAssociations

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *RouteData) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesRouteData{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.GatewayID = s.GatewayID
	o.MainTable = s.MainTable
	o.Routelist = s.Routelist
	o.SubnetAssociations = s.SubnetAssociations

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *RouteData) BleveType() string {

	return "routedata"
}

// DeepCopy returns a deep copy if the RouteData.
func (o *RouteData) DeepCopy() *RouteData {

	if o == nil {
		return nil
	}

	out := &RouteData{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *RouteData.
func (o *RouteData) DeepCopyInto(out *RouteData) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy RouteData: %s", err))
	}

	*out = *target.(*RouteData)
}

// Validate valides the current information stored into the structure.
func (o *RouteData) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	for _, sub := range o.Routelist {
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

type mongoAttributesRouteData struct {
	GatewayID          string          `bson:"gatewayid"`
	MainTable          bool            `bson:"maintable"`
	Routelist          CloudRoutesList `bson:"routelist"`
	SubnetAssociations []string        `bson:"subnetassociations"`
}
