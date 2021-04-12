package gaia

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CloudEndpointDataProduct represents the model of a cloudendpointdataproduct
type CloudEndpointDataProduct struct {
	// The ID of the corresponding product.
	ProductID string `json:"productID,omitempty" msgpack:"productID,omitempty" bson:"productid,omitempty" mapstructure:"productID,omitempty"`

	// The type of the product.
	Type string `json:"type,omitempty" msgpack:"type,omitempty" bson:"type,omitempty" mapstructure:"type,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCloudEndpointDataProduct returns a new *CloudEndpointDataProduct
func NewCloudEndpointDataProduct() *CloudEndpointDataProduct {

	return &CloudEndpointDataProduct{
		ModelVersion: 1,
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpointDataProduct) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCloudEndpointDataProduct{}

	s.ProductID = o.ProductID
	s.Type = o.Type

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CloudEndpointDataProduct) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCloudEndpointDataProduct{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ProductID = s.ProductID
	o.Type = s.Type

	return nil
}

// BleveType implements the bleve.Classifier Interface.
func (o *CloudEndpointDataProduct) BleveType() string {

	return "cloudendpointdataproduct"
}

// DeepCopy returns a deep copy if the CloudEndpointDataProduct.
func (o *CloudEndpointDataProduct) DeepCopy() *CloudEndpointDataProduct {

	if o == nil {
		return nil
	}

	out := &CloudEndpointDataProduct{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CloudEndpointDataProduct.
func (o *CloudEndpointDataProduct) DeepCopyInto(out *CloudEndpointDataProduct) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CloudEndpointDataProduct: %s", err))
	}

	*out = *target.(*CloudEndpointDataProduct)
}

// Validate valides the current information stored into the structure.
func (o *CloudEndpointDataProduct) Validate() error {

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

type mongoAttributesCloudEndpointDataProduct struct {
	ProductID string `bson:"productid,omitempty"`
	Type      string `bson:"type,omitempty"`
}
