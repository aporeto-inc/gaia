package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// PingConfigTypeValue represents the possible values for attribute "type".
type PingConfigTypeValue string

const (
	// PingConfigTypeAporetoIdentity represents the value AporetoIdentity.
	PingConfigTypeAporetoIdentity PingConfigTypeValue = "AporetoIdentity"

	// PingConfigTypeAporetoIdentityPassthrough represents the value AporetoIdentityPassthrough.
	PingConfigTypeAporetoIdentityPassthrough PingConfigTypeValue = "AporetoIdentityPassthrough"

	// PingConfigTypeCustomIdentity represents the value CustomIdentity.
	PingConfigTypeCustomIdentity PingConfigTypeValue = "CustomIdentity"

	// PingConfigTypeNone represents the value None.
	PingConfigTypeNone PingConfigTypeValue = "None"
)

// PingConfig represents the model of a pingconfig
type PingConfig struct {
	// Destination network to test.
	Network string `json:"network" msgpack:"network" bson:"network" mapstructure:"network,omitempty"`

	// Destination port(s) to test.
	Ports []string `json:"ports" msgpack:"ports" bson:"ports" mapstructure:"ports,omitempty"`

	// Number of requests to make on one call.
	Requests int `json:"requests" msgpack:"requests" bson:"requests" mapstructure:"requests,omitempty"`

	// Ping type.
	Type PingConfigTypeValue `json:"type" msgpack:"type" bson:"type" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewPingConfig returns a new *PingConfig
func NewPingConfig() *PingConfig {

	return &PingConfig{
		ModelVersion: 1,
		Ports:        []string{},
		Type:         PingConfigTypeNone,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingConfig) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesPingConfig{}

	s.Network = o.Network
	s.Ports = o.Ports
	s.Requests = o.Requests
	s.Type = o.Type

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *PingConfig) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesPingConfig{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.Network = s.Network
	o.Ports = s.Ports
	o.Requests = s.Requests
	o.Type = s.Type

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *PingConfig) BleveType() string {

	return "pingconfig"
}

// DeepCopy returns a deep copy if the PingConfig.
func (o *PingConfig) DeepCopy() *PingConfig {

	if o == nil {
		return nil
	}

	out := &PingConfig{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *PingConfig.
func (o *PingConfig) DeepCopyInto(out *PingConfig) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy PingConfig: %s", err))
	}

	*out = *target.(*PingConfig)
}

// Validate valides the current information stored into the structure.
func (o *PingConfig) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := ValidatePortStringList("ports", o.Ports); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateStringInList("type", string(o.Type), []string{"None", "AporetoIdentity", "CustomIdentity", "AporetoIdentityPassthrough"}, false); err != nil {
		errors = errors.Append(err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

type mongoAttributesPingConfig struct {
	Network  string              `bson:"network"`
	Ports    []string            `bson:"ports"`
	Requests int                 `bson:"requests"`
	Type     PingConfigTypeValue `bson:"type"`
}
