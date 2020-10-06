package gaia

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// CachedFlowRecordActionValue represents the possible values for attribute "action".
type CachedFlowRecordActionValue string

const (
	// CachedFlowRecordActionAccept represents the value Accept.
	CachedFlowRecordActionAccept CachedFlowRecordActionValue = "Accept"

	// CachedFlowRecordActionReject represents the value Reject.
	CachedFlowRecordActionReject CachedFlowRecordActionValue = "Reject"
)

// CachedFlowRecordDestinationTypeValue represents the possible values for attribute "destinationType".
type CachedFlowRecordDestinationTypeValue string

const (
	// CachedFlowRecordDestinationTypeClaims represents the value Claims.
	CachedFlowRecordDestinationTypeClaims CachedFlowRecordDestinationTypeValue = "Claims"

	// CachedFlowRecordDestinationTypeExternalNetwork represents the value ExternalNetwork.
	CachedFlowRecordDestinationTypeExternalNetwork CachedFlowRecordDestinationTypeValue = "ExternalNetwork"

	// CachedFlowRecordDestinationTypeProcessingUnit represents the value ProcessingUnit.
	CachedFlowRecordDestinationTypeProcessingUnit CachedFlowRecordDestinationTypeValue = "ProcessingUnit"
)

// CachedFlowRecordObservedActionValue represents the possible values for attribute "observedAction".
type CachedFlowRecordObservedActionValue string

const (
	// CachedFlowRecordObservedActionAccept represents the value Accept.
	CachedFlowRecordObservedActionAccept CachedFlowRecordObservedActionValue = "Accept"

	// CachedFlowRecordObservedActionNotApplicable represents the value NotApplicable.
	CachedFlowRecordObservedActionNotApplicable CachedFlowRecordObservedActionValue = "NotApplicable"

	// CachedFlowRecordObservedActionReject represents the value Reject.
	CachedFlowRecordObservedActionReject CachedFlowRecordObservedActionValue = "Reject"
)

// CachedFlowRecordServiceTypeValue represents the possible values for attribute "serviceType".
type CachedFlowRecordServiceTypeValue string

const (
	// CachedFlowRecordServiceTypeHTTP represents the value HTTP.
	CachedFlowRecordServiceTypeHTTP CachedFlowRecordServiceTypeValue = "HTTP"

	// CachedFlowRecordServiceTypeL3 represents the value L3.
	CachedFlowRecordServiceTypeL3 CachedFlowRecordServiceTypeValue = "L3"

	// CachedFlowRecordServiceTypeNotApplicable represents the value NotApplicable.
	CachedFlowRecordServiceTypeNotApplicable CachedFlowRecordServiceTypeValue = "NotApplicable"

	// CachedFlowRecordServiceTypeTCP represents the value TCP.
	CachedFlowRecordServiceTypeTCP CachedFlowRecordServiceTypeValue = "TCP"
)

// CachedFlowRecordSourceTypeValue represents the possible values for attribute "sourceType".
type CachedFlowRecordSourceTypeValue string

const (
	// CachedFlowRecordSourceTypeClaims represents the value Claims.
	CachedFlowRecordSourceTypeClaims CachedFlowRecordSourceTypeValue = "Claims"

	// CachedFlowRecordSourceTypeExternalNetwork represents the value ExternalNetwork.
	CachedFlowRecordSourceTypeExternalNetwork CachedFlowRecordSourceTypeValue = "ExternalNetwork"

	// CachedFlowRecordSourceTypeProcessingUnit represents the value ProcessingUnit.
	CachedFlowRecordSourceTypeProcessingUnit CachedFlowRecordSourceTypeValue = "ProcessingUnit"
)

// CachedFlowRecordIdentity represents the Identity of the object.
var CachedFlowRecordIdentity = elemental.Identity{
	Name:     "cachedflowrecord",
	Category: "cachedflowrecords",
	Package:  "zack",
	Private:  false,
}

// CachedFlowRecordsList represents a list of CachedFlowRecords
type CachedFlowRecordsList []*CachedFlowRecord

// Identity returns the identity of the objects in the list.
func (o CachedFlowRecordsList) Identity() elemental.Identity {

	return CachedFlowRecordIdentity
}

// Copy returns a pointer to a copy the CachedFlowRecordsList.
func (o CachedFlowRecordsList) Copy() elemental.Identifiables {

	copy := append(CachedFlowRecordsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the CachedFlowRecordsList.
func (o CachedFlowRecordsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(CachedFlowRecordsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*CachedFlowRecord))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o CachedFlowRecordsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o CachedFlowRecordsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the CachedFlowRecordsList converted to SparseCachedFlowRecordsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o CachedFlowRecordsList) ToSparse(fields ...string) elemental.Identifiables {

	out := make(SparseCachedFlowRecordsList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...).(*SparseCachedFlowRecord)
	}

	return out
}

// Version returns the version of the content.
func (o CachedFlowRecordsList) Version() int {

	return 1
}

// CachedFlowRecord represents the model of a cachedflowrecord
type CachedFlowRecord struct {
	// Action applied to the flow.
	Action CachedFlowRecordActionValue `json:"action" msgpack:"action" bson:"-" mapstructure:"action,omitempty"`

	// Identifier of the destination controller.
	DestinationController string `json:"destinationController" msgpack:"destinationController" bson:"-" mapstructure:"destinationController,omitempty"`

	// ID of the destination.
	DestinationID string `json:"destinationID" msgpack:"destinationID" bson:"-" mapstructure:"destinationID,omitempty"`

	// Destination IP address.
	DestinationIP string `json:"destinationIP" msgpack:"destinationIP" bson:"-" mapstructure:"destinationIP,omitempty"`

	// Indicates if the destination endpoint is an enforcer-local processing unit.
	DestinationIsTemporary bool `json:"destinationIsTemporary" msgpack:"destinationIsTemporary" bson:"-" mapstructure:"destinationIsTemporary,omitempty"`

	// Namespace of the destination. This is deprecated. Use `remoteNamespace`. This
	// property does nothing.
	DestinationNamespace string `json:"destinationNamespace,omitempty" msgpack:"destinationNamespace,omitempty" bson:"-" mapstructure:"destinationNamespace,omitempty"`

	// Port of the destination.
	DestinationPort int `json:"destinationPort" msgpack:"destinationPort" bson:"-" mapstructure:"destinationPort,omitempty"`

	// Destination type.
	DestinationType CachedFlowRecordDestinationTypeValue `json:"destinationType" msgpack:"destinationType" bson:"-" mapstructure:"destinationType,omitempty"`

	// This field is only set if `action` is set to `Reject`. It specifies the reason
	// for the rejection.
	DropReason string `json:"dropReason" msgpack:"dropReason" bson:"-" mapstructure:"dropReason,omitempty"`

	// If `true`, the flow was encrypted.
	Encrypted bool `json:"encrypted" msgpack:"encrypted" bson:"-" mapstructure:"encrypted,omitempty"`

	// This is here for backward compatibility.
	Namespace string `json:"namespace" msgpack:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// If `true`, design mode is on.
	Observed bool `json:"observed" msgpack:"observed" bson:"-" mapstructure:"observed,omitempty"`

	// Action observed on the flow.
	ObservedAction CachedFlowRecordObservedActionValue `json:"observedAction" msgpack:"observedAction" bson:"-" mapstructure:"observedAction,omitempty"`

	// Specifies the reason for a rejection. Only set if `observedAction` is set
	// to `Reject`.
	ObservedDropReason string `json:"observedDropReason" msgpack:"observedDropReason" bson:"-" mapstructure:"observedDropReason,omitempty"`

	// Value of the encryption of the network policy that observed the flow.
	ObservedEncrypted bool `json:"observedEncrypted" msgpack:"observedEncrypted" bson:"-" mapstructure:"observedEncrypted,omitempty"`

	// ID of the network policy that observed the flow.
	ObservedPolicyID string `json:"observedPolicyID" msgpack:"observedPolicyID" bson:"-" mapstructure:"observedPolicyID,omitempty"`

	// Namespace of the network policy that observed the flow.
	ObservedPolicyNamespace string `json:"observedPolicyNamespace" msgpack:"observedPolicyNamespace" bson:"-" mapstructure:"observedPolicyNamespace,omitempty"`

	// ID of the network policy that accepted the flow.
	PolicyID string `json:"policyID" msgpack:"policyID" bson:"-" mapstructure:"policyID,omitempty"`

	// Namespace of the network policy that accepted the flow.
	PolicyNamespace string `json:"policyNamespace" msgpack:"policyNamespace" bson:"-" mapstructure:"policyNamespace,omitempty"`

	// Protocol number.
	Protocol int `json:"protocol" msgpack:"protocol" bson:"-" mapstructure:"protocol,omitempty"`

	// Namespace of the object at the other end of the flow.
	RemoteNamespace string `json:"remoteNamespace,omitempty" msgpack:"remoteNamespace,omitempty" bson:"-" mapstructure:"remoteNamespace,omitempty"`

	// Hash of the claims used to communicate.
	ServiceClaimHash string `json:"serviceClaimHash" msgpack:"serviceClaimHash" bson:"-" mapstructure:"serviceClaimHash,omitempty"`

	// ID of the service.
	ServiceID string `json:"serviceID" msgpack:"serviceID" bson:"-" mapstructure:"serviceID,omitempty"`

	// Namespace of Service accessed.
	ServiceNamespace string `json:"serviceNamespace" msgpack:"serviceNamespace" bson:"-" mapstructure:"serviceNamespace,omitempty"`

	// ID of the service.
	ServiceType CachedFlowRecordServiceTypeValue `json:"serviceType" msgpack:"serviceType" bson:"-" mapstructure:"serviceType,omitempty"`

	// Service URL accessed.
	ServiceURL string `json:"serviceURL" msgpack:"serviceURL" bson:"-" mapstructure:"serviceURL,omitempty"`

	// Identifier of the source controller.
	SourceController string `json:"sourceController" msgpack:"sourceController" bson:"-" mapstructure:"sourceController,omitempty"`

	// ID of the source.
	SourceID string `json:"sourceID" msgpack:"sourceID" bson:"-" mapstructure:"sourceID,omitempty"`

	// Type of the source.
	SourceIP string `json:"sourceIP" msgpack:"sourceIP" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Indicates if the source endpoint is an enforcer-local processing unit.
	SourceIsTemporary bool `json:"sourceIsTemporary" msgpack:"sourceIsTemporary" bson:"-" mapstructure:"sourceIsTemporary,omitempty"`

	// Namespace of the source. This is deprecated. Use `remoteNamespace`. This
	// property does nothing.
	SourceNamespace string `json:"sourceNamespace,omitempty" msgpack:"sourceNamespace,omitempty" bson:"-" mapstructure:"sourceNamespace,omitempty"`

	// Type of the source.
	SourceType CachedFlowRecordSourceTypeValue `json:"sourceType" msgpack:"sourceType" bson:"-" mapstructure:"sourceType,omitempty"`

	// Time and date of the log.
	Timestamp time.Time `json:"timestamp" msgpack:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	// Number of flows in the log.
	Value int `json:"value" msgpack:"value" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewCachedFlowRecord returns a new *CachedFlowRecord
func NewCachedFlowRecord() *CachedFlowRecord {

	return &CachedFlowRecord{
		ModelVersion:   1,
		ServiceType:    CachedFlowRecordServiceTypeNotApplicable,
		ObservedAction: CachedFlowRecordObservedActionNotApplicable,
	}
}

// Identity returns the Identity of the object.
func (o *CachedFlowRecord) Identity() elemental.Identity {

	return CachedFlowRecordIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *CachedFlowRecord) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *CachedFlowRecord) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CachedFlowRecord) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesCachedFlowRecord{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *CachedFlowRecord) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesCachedFlowRecord{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *CachedFlowRecord) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *CachedFlowRecord) BleveType() string {

	return "cachedflowrecord"
}

// DefaultOrder returns the list of default ordering fields.
func (o *CachedFlowRecord) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *CachedFlowRecord) Doc() string {

	return `Post a new cached flow record.`
}

func (o *CachedFlowRecord) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *CachedFlowRecord) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseCachedFlowRecord{
			Action:                  &o.Action,
			DestinationController:   &o.DestinationController,
			DestinationID:           &o.DestinationID,
			DestinationIP:           &o.DestinationIP,
			DestinationIsTemporary:  &o.DestinationIsTemporary,
			DestinationNamespace:    &o.DestinationNamespace,
			DestinationPort:         &o.DestinationPort,
			DestinationType:         &o.DestinationType,
			DropReason:              &o.DropReason,
			Encrypted:               &o.Encrypted,
			Namespace:               &o.Namespace,
			Observed:                &o.Observed,
			ObservedAction:          &o.ObservedAction,
			ObservedDropReason:      &o.ObservedDropReason,
			ObservedEncrypted:       &o.ObservedEncrypted,
			ObservedPolicyID:        &o.ObservedPolicyID,
			ObservedPolicyNamespace: &o.ObservedPolicyNamespace,
			PolicyID:                &o.PolicyID,
			PolicyNamespace:         &o.PolicyNamespace,
			Protocol:                &o.Protocol,
			RemoteNamespace:         &o.RemoteNamespace,
			ServiceClaimHash:        &o.ServiceClaimHash,
			ServiceID:               &o.ServiceID,
			ServiceNamespace:        &o.ServiceNamespace,
			ServiceType:             &o.ServiceType,
			ServiceURL:              &o.ServiceURL,
			SourceController:        &o.SourceController,
			SourceID:                &o.SourceID,
			SourceIP:                &o.SourceIP,
			SourceIsTemporary:       &o.SourceIsTemporary,
			SourceNamespace:         &o.SourceNamespace,
			SourceType:              &o.SourceType,
			Timestamp:               &o.Timestamp,
			Value:                   &o.Value,
		}
	}

	sp := &SparseCachedFlowRecord{}
	for _, f := range fields {
		switch f {
		case "action":
			sp.Action = &(o.Action)
		case "destinationController":
			sp.DestinationController = &(o.DestinationController)
		case "destinationID":
			sp.DestinationID = &(o.DestinationID)
		case "destinationIP":
			sp.DestinationIP = &(o.DestinationIP)
		case "destinationIsTemporary":
			sp.DestinationIsTemporary = &(o.DestinationIsTemporary)
		case "destinationNamespace":
			sp.DestinationNamespace = &(o.DestinationNamespace)
		case "destinationPort":
			sp.DestinationPort = &(o.DestinationPort)
		case "destinationType":
			sp.DestinationType = &(o.DestinationType)
		case "dropReason":
			sp.DropReason = &(o.DropReason)
		case "encrypted":
			sp.Encrypted = &(o.Encrypted)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "observed":
			sp.Observed = &(o.Observed)
		case "observedAction":
			sp.ObservedAction = &(o.ObservedAction)
		case "observedDropReason":
			sp.ObservedDropReason = &(o.ObservedDropReason)
		case "observedEncrypted":
			sp.ObservedEncrypted = &(o.ObservedEncrypted)
		case "observedPolicyID":
			sp.ObservedPolicyID = &(o.ObservedPolicyID)
		case "observedPolicyNamespace":
			sp.ObservedPolicyNamespace = &(o.ObservedPolicyNamespace)
		case "policyID":
			sp.PolicyID = &(o.PolicyID)
		case "policyNamespace":
			sp.PolicyNamespace = &(o.PolicyNamespace)
		case "protocol":
			sp.Protocol = &(o.Protocol)
		case "remoteNamespace":
			sp.RemoteNamespace = &(o.RemoteNamespace)
		case "serviceClaimHash":
			sp.ServiceClaimHash = &(o.ServiceClaimHash)
		case "serviceID":
			sp.ServiceID = &(o.ServiceID)
		case "serviceNamespace":
			sp.ServiceNamespace = &(o.ServiceNamespace)
		case "serviceType":
			sp.ServiceType = &(o.ServiceType)
		case "serviceURL":
			sp.ServiceURL = &(o.ServiceURL)
		case "sourceController":
			sp.SourceController = &(o.SourceController)
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
		case "sourceIsTemporary":
			sp.SourceIsTemporary = &(o.SourceIsTemporary)
		case "sourceNamespace":
			sp.SourceNamespace = &(o.SourceNamespace)
		case "sourceType":
			sp.SourceType = &(o.SourceType)
		case "timestamp":
			sp.Timestamp = &(o.Timestamp)
		case "value":
			sp.Value = &(o.Value)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseCachedFlowRecord to the object.
func (o *CachedFlowRecord) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseCachedFlowRecord)
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.DestinationController != nil {
		o.DestinationController = *so.DestinationController
	}
	if so.DestinationID != nil {
		o.DestinationID = *so.DestinationID
	}
	if so.DestinationIP != nil {
		o.DestinationIP = *so.DestinationIP
	}
	if so.DestinationIsTemporary != nil {
		o.DestinationIsTemporary = *so.DestinationIsTemporary
	}
	if so.DestinationNamespace != nil {
		o.DestinationNamespace = *so.DestinationNamespace
	}
	if so.DestinationPort != nil {
		o.DestinationPort = *so.DestinationPort
	}
	if so.DestinationType != nil {
		o.DestinationType = *so.DestinationType
	}
	if so.DropReason != nil {
		o.DropReason = *so.DropReason
	}
	if so.Encrypted != nil {
		o.Encrypted = *so.Encrypted
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.Observed != nil {
		o.Observed = *so.Observed
	}
	if so.ObservedAction != nil {
		o.ObservedAction = *so.ObservedAction
	}
	if so.ObservedDropReason != nil {
		o.ObservedDropReason = *so.ObservedDropReason
	}
	if so.ObservedEncrypted != nil {
		o.ObservedEncrypted = *so.ObservedEncrypted
	}
	if so.ObservedPolicyID != nil {
		o.ObservedPolicyID = *so.ObservedPolicyID
	}
	if so.ObservedPolicyNamespace != nil {
		o.ObservedPolicyNamespace = *so.ObservedPolicyNamespace
	}
	if so.PolicyID != nil {
		o.PolicyID = *so.PolicyID
	}
	if so.PolicyNamespace != nil {
		o.PolicyNamespace = *so.PolicyNamespace
	}
	if so.Protocol != nil {
		o.Protocol = *so.Protocol
	}
	if so.RemoteNamespace != nil {
		o.RemoteNamespace = *so.RemoteNamespace
	}
	if so.ServiceClaimHash != nil {
		o.ServiceClaimHash = *so.ServiceClaimHash
	}
	if so.ServiceID != nil {
		o.ServiceID = *so.ServiceID
	}
	if so.ServiceNamespace != nil {
		o.ServiceNamespace = *so.ServiceNamespace
	}
	if so.ServiceType != nil {
		o.ServiceType = *so.ServiceType
	}
	if so.ServiceURL != nil {
		o.ServiceURL = *so.ServiceURL
	}
	if so.SourceController != nil {
		o.SourceController = *so.SourceController
	}
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
	}
	if so.SourceIsTemporary != nil {
		o.SourceIsTemporary = *so.SourceIsTemporary
	}
	if so.SourceNamespace != nil {
		o.SourceNamespace = *so.SourceNamespace
	}
	if so.SourceType != nil {
		o.SourceType = *so.SourceType
	}
	if so.Timestamp != nil {
		o.Timestamp = *so.Timestamp
	}
	if so.Value != nil {
		o.Value = *so.Value
	}
}

// DeepCopy returns a deep copy if the CachedFlowRecord.
func (o *CachedFlowRecord) DeepCopy() *CachedFlowRecord {

	if o == nil {
		return nil
	}

	out := &CachedFlowRecord{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *CachedFlowRecord.
func (o *CachedFlowRecord) DeepCopyInto(out *CachedFlowRecord) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy CachedFlowRecord: %s", err))
	}

	*out = *target.(*CachedFlowRecord)
}

// Validate valides the current information stored into the structure.
func (o *CachedFlowRecord) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("destinationID", o.DestinationID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("destinationType", string(o.DestinationType)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("destinationType", string(o.DestinationType), []string{"ProcessingUnit", "ExternalNetwork", "Claims"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("observedAction", string(o.ObservedAction), []string{"Accept", "Reject", "NotApplicable"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("policyID", o.PolicyID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("policyNamespace", o.PolicyNamespace); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("protocol", o.Protocol); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("serviceType", string(o.ServiceType), []string{"L3", "HTTP", "TCP", "NotApplicable"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceID", o.SourceID); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateRequiredString("sourceType", string(o.SourceType)); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := elemental.ValidateStringInList("sourceType", string(o.SourceType), []string{"ProcessingUnit", "ExternalNetwork", "Claims"}, false); err != nil {
		errors = errors.Append(err)
	}

	if err := elemental.ValidateRequiredInt("value", o.Value); err != nil {
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
func (*CachedFlowRecord) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := CachedFlowRecordAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return CachedFlowRecordLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*CachedFlowRecord) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return CachedFlowRecordAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *CachedFlowRecord) ValueForAttribute(name string) interface{} {

	switch name {
	case "action":
		return o.Action
	case "destinationController":
		return o.DestinationController
	case "destinationID":
		return o.DestinationID
	case "destinationIP":
		return o.DestinationIP
	case "destinationIsTemporary":
		return o.DestinationIsTemporary
	case "destinationNamespace":
		return o.DestinationNamespace
	case "destinationPort":
		return o.DestinationPort
	case "destinationType":
		return o.DestinationType
	case "dropReason":
		return o.DropReason
	case "encrypted":
		return o.Encrypted
	case "namespace":
		return o.Namespace
	case "observed":
		return o.Observed
	case "observedAction":
		return o.ObservedAction
	case "observedDropReason":
		return o.ObservedDropReason
	case "observedEncrypted":
		return o.ObservedEncrypted
	case "observedPolicyID":
		return o.ObservedPolicyID
	case "observedPolicyNamespace":
		return o.ObservedPolicyNamespace
	case "policyID":
		return o.PolicyID
	case "policyNamespace":
		return o.PolicyNamespace
	case "protocol":
		return o.Protocol
	case "remoteNamespace":
		return o.RemoteNamespace
	case "serviceClaimHash":
		return o.ServiceClaimHash
	case "serviceID":
		return o.ServiceID
	case "serviceNamespace":
		return o.ServiceNamespace
	case "serviceType":
		return o.ServiceType
	case "serviceURL":
		return o.ServiceURL
	case "sourceController":
		return o.SourceController
	case "sourceID":
		return o.SourceID
	case "sourceIP":
		return o.SourceIP
	case "sourceIsTemporary":
		return o.SourceIsTemporary
	case "sourceNamespace":
		return o.SourceNamespace
	case "sourceType":
		return o.SourceType
	case "timestamp":
		return o.Timestamp
	case "value":
		return o.Value
	}

	return nil
}

// CachedFlowRecordAttributesMap represents the map of attribute for CachedFlowRecord.
var CachedFlowRecordAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": {
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action applied to the flow.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"DestinationController": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationController",
		Description:    `Identifier of the destination controller.`,
		Exposed:        true,
		Name:           "destinationController",
		Type:           "string",
	},
	"DestinationID": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination.`,
		Exposed:        true,
		Name:           "destinationID",
		Required:       true,
		Type:           "string",
	},
	"DestinationIP": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Destination IP address.`,
		Exposed:        true,
		Name:           "destinationIP",
		Type:           "string",
	},
	"DestinationIsTemporary": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIsTemporary",
		Description:    `Indicates if the destination endpoint is an enforcer-local processing unit.`,
		Exposed:        true,
		Name:           "destinationIsTemporary",
		Type:           "boolean",
	},
	"DestinationNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationNamespace",
		Deprecated:     true,
		Description: `Namespace of the destination. This is deprecated. Use ` + "`" + `remoteNamespace` + "`" + `. This
property does nothing.`,
		Exposed: true,
		Name:    "destinationNamespace",
		Type:    "string",
	},
	"DestinationPort": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `Port of the destination.`,
		Exposed:        true,
		Name:           "destinationPort",
		Type:           "integer",
	},
	"DestinationType": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "DestinationType",
		Description:    `Destination type.`,
		Exposed:        true,
		Name:           "destinationType",
		Required:       true,
		Type:           "enum",
	},
	"DropReason": {
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if ` + "`" + `action` + "`" + ` is set to ` + "`" + `Reject` + "`" + `. It specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "dropReason",
		Type:    "string",
	},
	"Encrypted": {
		AllowedChoices: []string{},
		ConvertedName:  "Encrypted",
		Description:    `If ` + "`" + `true` + "`" + `, the flow was encrypted.`,
		Exposed:        true,
		Name:           "encrypted",
		Type:           "boolean",
	},
	"Namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Deprecated:     true,
		Description:    `This is here for backward compatibility.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"Observed": {
		AllowedChoices: []string{},
		ConvertedName:  "Observed",
		Description:    `If ` + "`" + `true` + "`" + `, design mode is on.`,
		Exposed:        true,
		Name:           "observed",
		Type:           "boolean",
	},
	"ObservedAction": {
		AllowedChoices: []string{"Accept", "Reject", "NotApplicable"},
		ConvertedName:  "ObservedAction",
		DefaultValue:   CachedFlowRecordObservedActionNotApplicable,
		Description:    `Action observed on the flow.`,
		Exposed:        true,
		Name:           "observedAction",
		Type:           "enum",
	},
	"ObservedDropReason": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedDropReason",
		Description: `Specifies the reason for a rejection. Only set if ` + "`" + `observedAction` + "`" + ` is set
to ` + "`" + `Reject` + "`" + `.`,
		Exposed: true,
		Name:    "observedDropReason",
		Type:    "string",
	},
	"ObservedEncrypted": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedEncrypted",
		Description:    `Value of the encryption of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedEncrypted",
		Type:           "boolean",
	},
	"ObservedPolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyID",
		Description:    `ID of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyID",
		Type:           "string",
	},
	"ObservedPolicyNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyNamespace",
		Description:    `Namespace of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyNamespace",
		Type:           "string",
	},
	"PolicyID": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyID",
		Required:       true,
		Type:           "string",
	},
	"PolicyNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `Namespace of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Required:       true,
		Type:           "string",
	},
	"Protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol number.`,
		Exposed:        true,
		Name:           "protocol",
		Required:       true,
		Type:           "integer",
	},
	"RemoteNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteNamespace",
		Description:    `Namespace of the object at the other end of the flow.`,
		Exposed:        true,
		Name:           "remoteNamespace",
		Type:           "string",
	},
	"ServiceClaimHash": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "serviceClaimHash",
		Type:           "string",
	},
	"ServiceID": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceID",
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceID",
		Type:           "string",
	},
	"ServiceNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceNamespace",
		Description:    `Namespace of Service accessed.`,
		Exposed:        true,
		Name:           "serviceNamespace",
		Type:           "string",
	},
	"ServiceType": {
		AllowedChoices: []string{"L3", "HTTP", "TCP", "NotApplicable"},
		ConvertedName:  "ServiceType",
		DefaultValue:   CachedFlowRecordServiceTypeNotApplicable,
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Type:           "enum",
	},
	"ServiceURL": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceURL",
		Description:    `Service URL accessed.`,
		Exposed:        true,
		Name:           "serviceURL",
		Type:           "string",
	},
	"SourceController": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceController",
		Description:    `Identifier of the source controller.`,
		Exposed:        true,
		Name:           "sourceController",
		Type:           "string",
	},
	"SourceID": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
	"SourceIP": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"SourceIsTemporary": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIsTemporary",
		Description:    `Indicates if the source endpoint is an enforcer-local processing unit.`,
		Exposed:        true,
		Name:           "sourceIsTemporary",
		Type:           "boolean",
	},
	"SourceNamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Deprecated:     true,
		Description: `Namespace of the source. This is deprecated. Use ` + "`" + `remoteNamespace` + "`" + `. This
property does nothing.`,
		Exposed: true,
		Name:    "sourceNamespace",
		Type:    "string",
	},
	"SourceType": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "SourceType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceType",
		Required:       true,
		Type:           "enum",
	},
	"Timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the log.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"Value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of flows in the log.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Type:           "integer",
	},
}

// CachedFlowRecordLowerCaseAttributesMap represents the map of attribute for CachedFlowRecord.
var CachedFlowRecordLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"action": {
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action applied to the flow.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"destinationcontroller": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationController",
		Description:    `Identifier of the destination controller.`,
		Exposed:        true,
		Name:           "destinationController",
		Type:           "string",
	},
	"destinationid": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination.`,
		Exposed:        true,
		Name:           "destinationID",
		Required:       true,
		Type:           "string",
	},
	"destinationip": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Destination IP address.`,
		Exposed:        true,
		Name:           "destinationIP",
		Type:           "string",
	},
	"destinationistemporary": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIsTemporary",
		Description:    `Indicates if the destination endpoint is an enforcer-local processing unit.`,
		Exposed:        true,
		Name:           "destinationIsTemporary",
		Type:           "boolean",
	},
	"destinationnamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationNamespace",
		Deprecated:     true,
		Description: `Namespace of the destination. This is deprecated. Use ` + "`" + `remoteNamespace` + "`" + `. This
property does nothing.`,
		Exposed: true,
		Name:    "destinationNamespace",
		Type:    "string",
	},
	"destinationport": {
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `Port of the destination.`,
		Exposed:        true,
		Name:           "destinationPort",
		Type:           "integer",
	},
	"destinationtype": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "DestinationType",
		Description:    `Destination type.`,
		Exposed:        true,
		Name:           "destinationType",
		Required:       true,
		Type:           "enum",
	},
	"dropreason": {
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if ` + "`" + `action` + "`" + ` is set to ` + "`" + `Reject` + "`" + `. It specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "dropReason",
		Type:    "string",
	},
	"encrypted": {
		AllowedChoices: []string{},
		ConvertedName:  "Encrypted",
		Description:    `If ` + "`" + `true` + "`" + `, the flow was encrypted.`,
		Exposed:        true,
		Name:           "encrypted",
		Type:           "boolean",
	},
	"namespace": {
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Deprecated:     true,
		Description:    `This is here for backward compatibility.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"observed": {
		AllowedChoices: []string{},
		ConvertedName:  "Observed",
		Description:    `If ` + "`" + `true` + "`" + `, design mode is on.`,
		Exposed:        true,
		Name:           "observed",
		Type:           "boolean",
	},
	"observedaction": {
		AllowedChoices: []string{"Accept", "Reject", "NotApplicable"},
		ConvertedName:  "ObservedAction",
		DefaultValue:   CachedFlowRecordObservedActionNotApplicable,
		Description:    `Action observed on the flow.`,
		Exposed:        true,
		Name:           "observedAction",
		Type:           "enum",
	},
	"observeddropreason": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedDropReason",
		Description: `Specifies the reason for a rejection. Only set if ` + "`" + `observedAction` + "`" + ` is set
to ` + "`" + `Reject` + "`" + `.`,
		Exposed: true,
		Name:    "observedDropReason",
		Type:    "string",
	},
	"observedencrypted": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedEncrypted",
		Description:    `Value of the encryption of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedEncrypted",
		Type:           "boolean",
	},
	"observedpolicyid": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyID",
		Description:    `ID of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyID",
		Type:           "string",
	},
	"observedpolicynamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyNamespace",
		Description:    `Namespace of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyNamespace",
		Type:           "string",
	},
	"policyid": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyID",
		Required:       true,
		Type:           "string",
	},
	"policynamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `Namespace of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Required:       true,
		Type:           "string",
	},
	"protocol": {
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `Protocol number.`,
		Exposed:        true,
		Name:           "protocol",
		Required:       true,
		Type:           "integer",
	},
	"remotenamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "RemoteNamespace",
		Description:    `Namespace of the object at the other end of the flow.`,
		Exposed:        true,
		Name:           "remoteNamespace",
		Type:           "string",
	},
	"serviceclaimhash": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "serviceClaimHash",
		Type:           "string",
	},
	"serviceid": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceID",
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceID",
		Type:           "string",
	},
	"servicenamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceNamespace",
		Description:    `Namespace of Service accessed.`,
		Exposed:        true,
		Name:           "serviceNamespace",
		Type:           "string",
	},
	"servicetype": {
		AllowedChoices: []string{"L3", "HTTP", "TCP", "NotApplicable"},
		ConvertedName:  "ServiceType",
		DefaultValue:   CachedFlowRecordServiceTypeNotApplicable,
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Type:           "enum",
	},
	"serviceurl": {
		AllowedChoices: []string{},
		ConvertedName:  "ServiceURL",
		Description:    `Service URL accessed.`,
		Exposed:        true,
		Name:           "serviceURL",
		Type:           "string",
	},
	"sourcecontroller": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceController",
		Description:    `Identifier of the source controller.`,
		Exposed:        true,
		Name:           "sourceController",
		Type:           "string",
	},
	"sourceid": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
	"sourceip": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"sourceistemporary": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceIsTemporary",
		Description:    `Indicates if the source endpoint is an enforcer-local processing unit.`,
		Exposed:        true,
		Name:           "sourceIsTemporary",
		Type:           "boolean",
	},
	"sourcenamespace": {
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Deprecated:     true,
		Description: `Namespace of the source. This is deprecated. Use ` + "`" + `remoteNamespace` + "`" + `. This
property does nothing.`,
		Exposed: true,
		Name:    "sourceNamespace",
		Type:    "string",
	},
	"sourcetype": {
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "SourceType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceType",
		Required:       true,
		Type:           "enum",
	},
	"timestamp": {
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Time and date of the log.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"value": {
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of flows in the log.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Type:           "integer",
	},
}

// SparseCachedFlowRecordsList represents a list of SparseCachedFlowRecords
type SparseCachedFlowRecordsList []*SparseCachedFlowRecord

// Identity returns the identity of the objects in the list.
func (o SparseCachedFlowRecordsList) Identity() elemental.Identity {

	return CachedFlowRecordIdentity
}

// Copy returns a pointer to a copy the SparseCachedFlowRecordsList.
func (o SparseCachedFlowRecordsList) Copy() elemental.Identifiables {

	copy := append(SparseCachedFlowRecordsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseCachedFlowRecordsList.
func (o SparseCachedFlowRecordsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseCachedFlowRecordsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseCachedFlowRecord))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseCachedFlowRecordsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseCachedFlowRecordsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseCachedFlowRecordsList converted to CachedFlowRecordsList.
func (o SparseCachedFlowRecordsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseCachedFlowRecordsList) Version() int {

	return 1
}

// SparseCachedFlowRecord represents the sparse version of a cachedflowrecord.
type SparseCachedFlowRecord struct {
	// Action applied to the flow.
	Action *CachedFlowRecordActionValue `json:"action,omitempty" msgpack:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// Identifier of the destination controller.
	DestinationController *string `json:"destinationController,omitempty" msgpack:"destinationController,omitempty" bson:"-" mapstructure:"destinationController,omitempty"`

	// ID of the destination.
	DestinationID *string `json:"destinationID,omitempty" msgpack:"destinationID,omitempty" bson:"-" mapstructure:"destinationID,omitempty"`

	// Destination IP address.
	DestinationIP *string `json:"destinationIP,omitempty" msgpack:"destinationIP,omitempty" bson:"-" mapstructure:"destinationIP,omitempty"`

	// Indicates if the destination endpoint is an enforcer-local processing unit.
	DestinationIsTemporary *bool `json:"destinationIsTemporary,omitempty" msgpack:"destinationIsTemporary,omitempty" bson:"-" mapstructure:"destinationIsTemporary,omitempty"`

	// Namespace of the destination. This is deprecated. Use `remoteNamespace`. This
	// property does nothing.
	DestinationNamespace *string `json:"destinationNamespace,omitempty" msgpack:"destinationNamespace,omitempty" bson:"-" mapstructure:"destinationNamespace,omitempty"`

	// Port of the destination.
	DestinationPort *int `json:"destinationPort,omitempty" msgpack:"destinationPort,omitempty" bson:"-" mapstructure:"destinationPort,omitempty"`

	// Destination type.
	DestinationType *CachedFlowRecordDestinationTypeValue `json:"destinationType,omitempty" msgpack:"destinationType,omitempty" bson:"-" mapstructure:"destinationType,omitempty"`

	// This field is only set if `action` is set to `Reject`. It specifies the reason
	// for the rejection.
	DropReason *string `json:"dropReason,omitempty" msgpack:"dropReason,omitempty" bson:"-" mapstructure:"dropReason,omitempty"`

	// If `true`, the flow was encrypted.
	Encrypted *bool `json:"encrypted,omitempty" msgpack:"encrypted,omitempty" bson:"-" mapstructure:"encrypted,omitempty"`

	// This is here for backward compatibility.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// If `true`, design mode is on.
	Observed *bool `json:"observed,omitempty" msgpack:"observed,omitempty" bson:"-" mapstructure:"observed,omitempty"`

	// Action observed on the flow.
	ObservedAction *CachedFlowRecordObservedActionValue `json:"observedAction,omitempty" msgpack:"observedAction,omitempty" bson:"-" mapstructure:"observedAction,omitempty"`

	// Specifies the reason for a rejection. Only set if `observedAction` is set
	// to `Reject`.
	ObservedDropReason *string `json:"observedDropReason,omitempty" msgpack:"observedDropReason,omitempty" bson:"-" mapstructure:"observedDropReason,omitempty"`

	// Value of the encryption of the network policy that observed the flow.
	ObservedEncrypted *bool `json:"observedEncrypted,omitempty" msgpack:"observedEncrypted,omitempty" bson:"-" mapstructure:"observedEncrypted,omitempty"`

	// ID of the network policy that observed the flow.
	ObservedPolicyID *string `json:"observedPolicyID,omitempty" msgpack:"observedPolicyID,omitempty" bson:"-" mapstructure:"observedPolicyID,omitempty"`

	// Namespace of the network policy that observed the flow.
	ObservedPolicyNamespace *string `json:"observedPolicyNamespace,omitempty" msgpack:"observedPolicyNamespace,omitempty" bson:"-" mapstructure:"observedPolicyNamespace,omitempty"`

	// ID of the network policy that accepted the flow.
	PolicyID *string `json:"policyID,omitempty" msgpack:"policyID,omitempty" bson:"-" mapstructure:"policyID,omitempty"`

	// Namespace of the network policy that accepted the flow.
	PolicyNamespace *string `json:"policyNamespace,omitempty" msgpack:"policyNamespace,omitempty" bson:"-" mapstructure:"policyNamespace,omitempty"`

	// Protocol number.
	Protocol *int `json:"protocol,omitempty" msgpack:"protocol,omitempty" bson:"-" mapstructure:"protocol,omitempty"`

	// Namespace of the object at the other end of the flow.
	RemoteNamespace *string `json:"remoteNamespace,omitempty" msgpack:"remoteNamespace,omitempty" bson:"-" mapstructure:"remoteNamespace,omitempty"`

	// Hash of the claims used to communicate.
	ServiceClaimHash *string `json:"serviceClaimHash,omitempty" msgpack:"serviceClaimHash,omitempty" bson:"-" mapstructure:"serviceClaimHash,omitempty"`

	// ID of the service.
	ServiceID *string `json:"serviceID,omitempty" msgpack:"serviceID,omitempty" bson:"-" mapstructure:"serviceID,omitempty"`

	// Namespace of Service accessed.
	ServiceNamespace *string `json:"serviceNamespace,omitempty" msgpack:"serviceNamespace,omitempty" bson:"-" mapstructure:"serviceNamespace,omitempty"`

	// ID of the service.
	ServiceType *CachedFlowRecordServiceTypeValue `json:"serviceType,omitempty" msgpack:"serviceType,omitempty" bson:"-" mapstructure:"serviceType,omitempty"`

	// Service URL accessed.
	ServiceURL *string `json:"serviceURL,omitempty" msgpack:"serviceURL,omitempty" bson:"-" mapstructure:"serviceURL,omitempty"`

	// Identifier of the source controller.
	SourceController *string `json:"sourceController,omitempty" msgpack:"sourceController,omitempty" bson:"-" mapstructure:"sourceController,omitempty"`

	// ID of the source.
	SourceID *string `json:"sourceID,omitempty" msgpack:"sourceID,omitempty" bson:"-" mapstructure:"sourceID,omitempty"`

	// Type of the source.
	SourceIP *string `json:"sourceIP,omitempty" msgpack:"sourceIP,omitempty" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Indicates if the source endpoint is an enforcer-local processing unit.
	SourceIsTemporary *bool `json:"sourceIsTemporary,omitempty" msgpack:"sourceIsTemporary,omitempty" bson:"-" mapstructure:"sourceIsTemporary,omitempty"`

	// Namespace of the source. This is deprecated. Use `remoteNamespace`. This
	// property does nothing.
	SourceNamespace *string `json:"sourceNamespace,omitempty" msgpack:"sourceNamespace,omitempty" bson:"-" mapstructure:"sourceNamespace,omitempty"`

	// Type of the source.
	SourceType *CachedFlowRecordSourceTypeValue `json:"sourceType,omitempty" msgpack:"sourceType,omitempty" bson:"-" mapstructure:"sourceType,omitempty"`

	// Time and date of the log.
	Timestamp *time.Time `json:"timestamp,omitempty" msgpack:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	// Number of flows in the log.
	Value *int `json:"value,omitempty" msgpack:"value,omitempty" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseCachedFlowRecord returns a new  SparseCachedFlowRecord.
func NewSparseCachedFlowRecord() *SparseCachedFlowRecord {
	return &SparseCachedFlowRecord{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseCachedFlowRecord) Identity() elemental.Identity {

	return CachedFlowRecordIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseCachedFlowRecord) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseCachedFlowRecord) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCachedFlowRecord) GetBSON() (interface{}, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseCachedFlowRecord{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseCachedFlowRecord) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseCachedFlowRecord{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseCachedFlowRecord) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseCachedFlowRecord) ToPlain() elemental.PlainIdentifiable {

	out := NewCachedFlowRecord()
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.DestinationController != nil {
		out.DestinationController = *o.DestinationController
	}
	if o.DestinationID != nil {
		out.DestinationID = *o.DestinationID
	}
	if o.DestinationIP != nil {
		out.DestinationIP = *o.DestinationIP
	}
	if o.DestinationIsTemporary != nil {
		out.DestinationIsTemporary = *o.DestinationIsTemporary
	}
	if o.DestinationNamespace != nil {
		out.DestinationNamespace = *o.DestinationNamespace
	}
	if o.DestinationPort != nil {
		out.DestinationPort = *o.DestinationPort
	}
	if o.DestinationType != nil {
		out.DestinationType = *o.DestinationType
	}
	if o.DropReason != nil {
		out.DropReason = *o.DropReason
	}
	if o.Encrypted != nil {
		out.Encrypted = *o.Encrypted
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.Observed != nil {
		out.Observed = *o.Observed
	}
	if o.ObservedAction != nil {
		out.ObservedAction = *o.ObservedAction
	}
	if o.ObservedDropReason != nil {
		out.ObservedDropReason = *o.ObservedDropReason
	}
	if o.ObservedEncrypted != nil {
		out.ObservedEncrypted = *o.ObservedEncrypted
	}
	if o.ObservedPolicyID != nil {
		out.ObservedPolicyID = *o.ObservedPolicyID
	}
	if o.ObservedPolicyNamespace != nil {
		out.ObservedPolicyNamespace = *o.ObservedPolicyNamespace
	}
	if o.PolicyID != nil {
		out.PolicyID = *o.PolicyID
	}
	if o.PolicyNamespace != nil {
		out.PolicyNamespace = *o.PolicyNamespace
	}
	if o.Protocol != nil {
		out.Protocol = *o.Protocol
	}
	if o.RemoteNamespace != nil {
		out.RemoteNamespace = *o.RemoteNamespace
	}
	if o.ServiceClaimHash != nil {
		out.ServiceClaimHash = *o.ServiceClaimHash
	}
	if o.ServiceID != nil {
		out.ServiceID = *o.ServiceID
	}
	if o.ServiceNamespace != nil {
		out.ServiceNamespace = *o.ServiceNamespace
	}
	if o.ServiceType != nil {
		out.ServiceType = *o.ServiceType
	}
	if o.ServiceURL != nil {
		out.ServiceURL = *o.ServiceURL
	}
	if o.SourceController != nil {
		out.SourceController = *o.SourceController
	}
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
	}
	if o.SourceIsTemporary != nil {
		out.SourceIsTemporary = *o.SourceIsTemporary
	}
	if o.SourceNamespace != nil {
		out.SourceNamespace = *o.SourceNamespace
	}
	if o.SourceType != nil {
		out.SourceType = *o.SourceType
	}
	if o.Timestamp != nil {
		out.Timestamp = *o.Timestamp
	}
	if o.Value != nil {
		out.Value = *o.Value
	}

	return out
}

// DeepCopy returns a deep copy if the SparseCachedFlowRecord.
func (o *SparseCachedFlowRecord) DeepCopy() *SparseCachedFlowRecord {

	if o == nil {
		return nil
	}

	out := &SparseCachedFlowRecord{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseCachedFlowRecord.
func (o *SparseCachedFlowRecord) DeepCopyInto(out *SparseCachedFlowRecord) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseCachedFlowRecord: %s", err))
	}

	*out = *target.(*SparseCachedFlowRecord)
}

type mongoAttributesCachedFlowRecord struct {
}
type mongoAttributesSparseCachedFlowRecord struct {
}
