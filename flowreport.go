package gaia

import (
	"fmt"
	"sync"
	"time"

	"github.com/mitchellh/copystructure"
	"go.aporeto.io/elemental"
)

// FlowReportActionValue represents the possible values for attribute "action".
type FlowReportActionValue string

const (
	// FlowReportActionAccept represents the value Accept.
	FlowReportActionAccept FlowReportActionValue = "Accept"

	// FlowReportActionReject represents the value Reject.
	FlowReportActionReject FlowReportActionValue = "Reject"
)

// FlowReportDestinationTypeValue represents the possible values for attribute "destinationType".
type FlowReportDestinationTypeValue string

const (
	// FlowReportDestinationTypeClaims represents the value Claims.
	FlowReportDestinationTypeClaims FlowReportDestinationTypeValue = "Claims"

	// FlowReportDestinationTypeExternalNetwork represents the value ExternalNetwork.
	FlowReportDestinationTypeExternalNetwork FlowReportDestinationTypeValue = "ExternalNetwork"

	// FlowReportDestinationTypeProcessingUnit represents the value ProcessingUnit.
	FlowReportDestinationTypeProcessingUnit FlowReportDestinationTypeValue = "ProcessingUnit"
)

// FlowReportObservedActionValue represents the possible values for attribute "observedAction".
type FlowReportObservedActionValue string

const (
	// FlowReportObservedActionAccept represents the value Accept.
	FlowReportObservedActionAccept FlowReportObservedActionValue = "Accept"

	// FlowReportObservedActionNotApplicable represents the value NotApplicable.
	FlowReportObservedActionNotApplicable FlowReportObservedActionValue = "NotApplicable"

	// FlowReportObservedActionReject represents the value Reject.
	FlowReportObservedActionReject FlowReportObservedActionValue = "Reject"
)

// FlowReportServiceTypeValue represents the possible values for attribute "serviceType".
type FlowReportServiceTypeValue string

const (
	// FlowReportServiceTypeHTTP represents the value HTTP.
	FlowReportServiceTypeHTTP FlowReportServiceTypeValue = "HTTP"

	// FlowReportServiceTypeL3 represents the value L3.
	FlowReportServiceTypeL3 FlowReportServiceTypeValue = "L3"

	// FlowReportServiceTypeNotApplicable represents the value NotApplicable.
	FlowReportServiceTypeNotApplicable FlowReportServiceTypeValue = "NotApplicable"

	// FlowReportServiceTypeTCP represents the value TCP.
	FlowReportServiceTypeTCP FlowReportServiceTypeValue = "TCP"
)

// FlowReportSourceTypeValue represents the possible values for attribute "sourceType".
type FlowReportSourceTypeValue string

const (
	// FlowReportSourceTypeClaims represents the value Claims.
	FlowReportSourceTypeClaims FlowReportSourceTypeValue = "Claims"

	// FlowReportSourceTypeExternalNetwork represents the value ExternalNetwork.
	FlowReportSourceTypeExternalNetwork FlowReportSourceTypeValue = "ExternalNetwork"

	// FlowReportSourceTypeProcessingUnit represents the value ProcessingUnit.
	FlowReportSourceTypeProcessingUnit FlowReportSourceTypeValue = "ProcessingUnit"
)

// FlowReportIdentity represents the Identity of the object.
var FlowReportIdentity = elemental.Identity{
	Name:     "flowreport",
	Category: "flowreports",
	Package:  "zack",
	Private:  false,
}

// FlowReportsList represents a list of FlowReports
type FlowReportsList []*FlowReport

// Identity returns the identity of the objects in the list.
func (o FlowReportsList) Identity() elemental.Identity {

	return FlowReportIdentity
}

// Copy returns a pointer to a copy the FlowReportsList.
func (o FlowReportsList) Copy() elemental.Identifiables {

	copy := append(FlowReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the FlowReportsList.
func (o FlowReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(FlowReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*FlowReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o FlowReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FlowReportsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the FlowReportsList converted to SparseFlowReportsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o FlowReportsList) ToSparse(fields ...string) elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToSparse(fields...)
	}

	return out
}

// Version returns the version of the content.
func (o FlowReportsList) Version() int {

	return 1
}

// FlowReport represents the model of a flowreport
type FlowReport struct {
	// Action applied to the flow.
	Action FlowReportActionValue `json:"action" bson:"-" mapstructure:"action,omitempty"`

	// ID of the destination.
	DestinationID string `json:"destinationID" bson:"-" mapstructure:"destinationID,omitempty"`

	// Type of the destination.
	DestinationIP string `json:"destinationIP" bson:"-" mapstructure:"destinationIP,omitempty"`

	// Namespace of the receiver.
	DestinationNamespace string `json:"destinationNamespace" bson:"-" mapstructure:"destinationNamespace,omitempty"`

	// Port of the destination.
	DestinationPort int `json:"destinationPort" bson:"-" mapstructure:"destinationPort,omitempty"`

	// Type of the source.
	DestinationType FlowReportDestinationTypeValue `json:"destinationType" bson:"-" mapstructure:"destinationType,omitempty"`

	// This field is only set if 'action' is set to 'Reject' and specifies the reason
	// for the rejection.
	DropReason string `json:"dropReason" bson:"-" mapstructure:"dropReason,omitempty"`

	// Tells is the flow has been encrypted.
	Encrypted bool `json:"encrypted" bson:"-" mapstructure:"encrypted,omitempty"`

	// This is here for backward compatibility.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Tells if the flow is from design mode.
	Observed bool `json:"observed" bson:"-" mapstructure:"observed,omitempty"`

	// Action observed on the flow.
	ObservedAction FlowReportObservedActionValue `json:"observedAction" bson:"-" mapstructure:"observedAction,omitempty"`

	// This field is only set if 'observedAction' is set to 'Reject' and specifies the
	// reason for the rejection.
	ObservedDropReason string `json:"observedDropReason" bson:"-" mapstructure:"observedDropReason,omitempty"`

	// Value of the encryption of the network policy that observed the flow.
	ObservedEncrypted bool `json:"observedEncrypted" bson:"-" mapstructure:"observedEncrypted,omitempty"`

	// ID of the network policy that observed the flow.
	ObservedPolicyID string `json:"observedPolicyID" bson:"-" mapstructure:"observedPolicyID,omitempty"`

	// Namespace of the network policy that observed the flow.
	ObservedPolicyNamespace string `json:"observedPolicyNamespace" bson:"-" mapstructure:"observedPolicyNamespace,omitempty"`

	// ID of the network policy that accepted the flow.
	PolicyID string `json:"policyID" bson:"-" mapstructure:"policyID,omitempty"`

	// Namespace of the network policy that accepted the flow.
	PolicyNamespace string `json:"policyNamespace" bson:"-" mapstructure:"policyNamespace,omitempty"`

	// protocol number.
	Protocol int `json:"protocol" bson:"-" mapstructure:"protocol,omitempty"`

	// Hash of the claims used to communicate.
	ServiceClaimHash string `json:"serviceClaimHash" bson:"-" mapstructure:"serviceClaimHash,omitempty"`

	// ID of the service.
	ServiceID string `json:"serviceID" bson:"-" mapstructure:"serviceID,omitempty"`

	// Service URL accessed.
	ServiceNamespace string `json:"serviceNamespace" bson:"-" mapstructure:"serviceNamespace,omitempty"`

	// ID of the service.
	ServiceType FlowReportServiceTypeValue `json:"serviceType" bson:"-" mapstructure:"serviceType,omitempty"`

	// Service URL accessed.
	ServiceURL string `json:"serviceURL" bson:"-" mapstructure:"serviceURL,omitempty"`

	// ID of the source.
	SourceID string `json:"sourceID" bson:"-" mapstructure:"sourceID,omitempty"`

	// Type of the source.
	SourceIP string `json:"sourceIP" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Namespace of the receiver.
	SourceNamespace string `json:"sourceNamespace" bson:"-" mapstructure:"sourceNamespace,omitempty"`

	// Type of the source.
	SourceType FlowReportSourceTypeValue `json:"sourceType" bson:"-" mapstructure:"sourceType,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	// Number of flows in the report.
	Value int `json:"value" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewFlowReport returns a new *FlowReport
func NewFlowReport() *FlowReport {

	return &FlowReport{
		ModelVersion:   1,
		Mutex:          &sync.Mutex{},
		ServiceType:    FlowReportServiceTypeNotApplicable,
		ObservedAction: FlowReportObservedActionNotApplicable,
	}
}

// Identity returns the Identity of the object.
func (o *FlowReport) Identity() elemental.Identity {

	return FlowReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FlowReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FlowReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *FlowReport) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *FlowReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *FlowReport) Doc() string {
	return `Post a new flow statistics report.`
}

func (o *FlowReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *FlowReport) ToSparse(fields ...string) elemental.SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseFlowReport{
			Action:                  &o.Action,
			DestinationID:           &o.DestinationID,
			DestinationIP:           &o.DestinationIP,
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
			ServiceClaimHash:        &o.ServiceClaimHash,
			ServiceID:               &o.ServiceID,
			ServiceNamespace:        &o.ServiceNamespace,
			ServiceType:             &o.ServiceType,
			ServiceURL:              &o.ServiceURL,
			SourceID:                &o.SourceID,
			SourceIP:                &o.SourceIP,
			SourceNamespace:         &o.SourceNamespace,
			SourceType:              &o.SourceType,
			Timestamp:               &o.Timestamp,
			Value:                   &o.Value,
		}
	}

	sp := &SparseFlowReport{}
	for _, f := range fields {
		switch f {
		case "action":
			sp.Action = &(o.Action)
		case "destinationID":
			sp.DestinationID = &(o.DestinationID)
		case "destinationIP":
			sp.DestinationIP = &(o.DestinationIP)
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
		case "sourceID":
			sp.SourceID = &(o.SourceID)
		case "sourceIP":
			sp.SourceIP = &(o.SourceIP)
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

// Patch apply the non nil value of a *SparseFlowReport to the object.
func (o *FlowReport) Patch(sparse elemental.SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseFlowReport)
	if so.Action != nil {
		o.Action = *so.Action
	}
	if so.DestinationID != nil {
		o.DestinationID = *so.DestinationID
	}
	if so.DestinationIP != nil {
		o.DestinationIP = *so.DestinationIP
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
	if so.SourceID != nil {
		o.SourceID = *so.SourceID
	}
	if so.SourceIP != nil {
		o.SourceIP = *so.SourceIP
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

// DeepCopy returns a deep copy if the FlowReport.
func (o *FlowReport) DeepCopy() *FlowReport {

	if o == nil {
		return nil
	}

	out := &FlowReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *FlowReport.
func (o *FlowReport) DeepCopyInto(out *FlowReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy FlowReport: %s", err))
	}

	*out = *target.(*FlowReport)
}

// Validate valides the current information stored into the structure.
func (o *FlowReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("action", string(o.Action)); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("destinationID", o.DestinationID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("destinationType", string(o.DestinationType)); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("destinationType", string(o.DestinationType), []string{"ProcessingUnit", "ExternalNetwork", "Claims"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("observedAction", string(o.ObservedAction), []string{"Accept", "Reject", "NotApplicable"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("policyID", o.PolicyID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("policyNamespace", o.PolicyNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredInt("protocol", o.Protocol); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("serviceType", string(o.ServiceType), []string{"L3", "HTTP", "TCP", "NotApplicable"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("sourceID", o.SourceID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("sourceType", string(o.SourceType)); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("sourceType", string(o.SourceType), []string{"ProcessingUnit", "ExternalNetwork", "Claims"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredInt("value", o.Value); err != nil {
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

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*FlowReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := FlowReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return FlowReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*FlowReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FlowReportAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *FlowReport) ValueForAttribute(name string) interface{} {

	switch name {
	case "action":
		return o.Action
	case "destinationID":
		return o.DestinationID
	case "destinationIP":
		return o.DestinationIP
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
	case "sourceID":
		return o.SourceID
	case "sourceIP":
		return o.SourceIP
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

// FlowReportAttributesMap represents the map of attribute for FlowReport.
var FlowReportAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action applied to the flow.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"DestinationID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination.`,
		Exposed:        true,
		Name:           "destinationID",
		Required:       true,
		Type:           "string",
	},
	"DestinationIP": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Type of the destination.`,
		Exposed:        true,
		Name:           "destinationIP",
		Type:           "string",
	},
	"DestinationNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationNamespace",
		Description:    `Namespace of the receiver.`,
		Exposed:        true,
		Name:           "destinationNamespace",
		Type:           "string",
	},
	"DestinationPort": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `Port of the destination.`,
		Exposed:        true,
		Name:           "destinationPort",
		Type:           "integer",
	},
	"DestinationType": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "DestinationType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "destinationType",
		Required:       true,
		Type:           "enum",
	},
	"DropReason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if 'action' is set to 'Reject' and specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "dropReason",
		Type:    "string",
	},
	"Encrypted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Encrypted",
		Description:    `Tells is the flow has been encrypted.`,
		Exposed:        true,
		Name:           "encrypted",
		Type:           "boolean",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Deprecated:     true,
		Description:    `This is here for backward compatibility.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"Observed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Observed",
		Description:    `Tells if the flow is from design mode.`,
		Exposed:        true,
		Name:           "observed",
		Type:           "boolean",
	},
	"ObservedAction": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accept", "Reject", "NotApplicable"},
		ConvertedName:  "ObservedAction",
		DefaultValue:   FlowReportObservedActionNotApplicable,
		Description:    `Action observed on the flow.`,
		Exposed:        true,
		Name:           "observedAction",
		Type:           "enum",
	},
	"ObservedDropReason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedDropReason",
		Description: `This field is only set if 'observedAction' is set to 'Reject' and specifies the
reason for the rejection.`,
		Exposed: true,
		Name:    "observedDropReason",
		Type:    "string",
	},
	"ObservedEncrypted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedEncrypted",
		Description:    `Value of the encryption of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedEncrypted",
		Type:           "boolean",
	},
	"ObservedPolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyID",
		Description:    `ID of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyID",
		Type:           "string",
	},
	"ObservedPolicyNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyNamespace",
		Description:    `Namespace of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyNamespace",
		Type:           "string",
	},
	"PolicyID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyID",
		Required:       true,
		Type:           "string",
	},
	"PolicyNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `Namespace of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Required:       true,
		Type:           "string",
	},
	"Protocol": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `protocol number.`,
		Exposed:        true,
		Name:           "protocol",
		Required:       true,
		Type:           "integer",
	},
	"ServiceClaimHash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "serviceClaimHash",
		Type:           "string",
	},
	"ServiceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceID",
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceID",
		Type:           "string",
	},
	"ServiceNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceNamespace",
		Description:    `Service URL accessed.`,
		Exposed:        true,
		Name:           "serviceNamespace",
		Type:           "string",
	},
	"ServiceType": elemental.AttributeSpecification{
		AllowedChoices: []string{"L3", "HTTP", "TCP", "NotApplicable"},
		ConvertedName:  "ServiceType",
		DefaultValue:   FlowReportServiceTypeNotApplicable,
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Type:           "enum",
	},
	"ServiceURL": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceURL",
		Description:    `Service URL accessed.`,
		Exposed:        true,
		Name:           "serviceURL",
		Type:           "string",
	},
	"SourceID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
	"SourceIP": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"SourceNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Description:    `Namespace of the receiver.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Type:           "string",
	},
	"SourceType": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "SourceType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceType",
		Required:       true,
		Type:           "enum",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"Value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of flows in the report.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Type:           "integer",
	},
}

// FlowReportLowerCaseAttributesMap represents the map of attribute for FlowReport.
var FlowReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"action": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accept", "Reject"},
		ConvertedName:  "Action",
		Description:    `Action applied to the flow.`,
		Exposed:        true,
		Name:           "action",
		Required:       true,
		Type:           "enum",
	},
	"destinationid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationID",
		Description:    `ID of the destination.`,
		Exposed:        true,
		Name:           "destinationID",
		Required:       true,
		Type:           "string",
	},
	"destinationip": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationIP",
		Description:    `Type of the destination.`,
		Exposed:        true,
		Name:           "destinationIP",
		Type:           "string",
	},
	"destinationnamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationNamespace",
		Description:    `Namespace of the receiver.`,
		Exposed:        true,
		Name:           "destinationNamespace",
		Type:           "string",
	},
	"destinationport": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DestinationPort",
		Description:    `Port of the destination.`,
		Exposed:        true,
		Name:           "destinationPort",
		Type:           "integer",
	},
	"destinationtype": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "DestinationType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "destinationType",
		Required:       true,
		Type:           "enum",
	},
	"dropreason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "DropReason",
		Description: `This field is only set if 'action' is set to 'Reject' and specifies the reason
for the rejection.`,
		Exposed: true,
		Name:    "dropReason",
		Type:    "string",
	},
	"encrypted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Encrypted",
		Description:    `Tells is the flow has been encrypted.`,
		Exposed:        true,
		Name:           "encrypted",
		Type:           "boolean",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Deprecated:     true,
		Description:    `This is here for backward compatibility.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"observed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Observed",
		Description:    `Tells if the flow is from design mode.`,
		Exposed:        true,
		Name:           "observed",
		Type:           "boolean",
	},
	"observedaction": elemental.AttributeSpecification{
		AllowedChoices: []string{"Accept", "Reject", "NotApplicable"},
		ConvertedName:  "ObservedAction",
		DefaultValue:   FlowReportObservedActionNotApplicable,
		Description:    `Action observed on the flow.`,
		Exposed:        true,
		Name:           "observedAction",
		Type:           "enum",
	},
	"observeddropreason": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedDropReason",
		Description: `This field is only set if 'observedAction' is set to 'Reject' and specifies the
reason for the rejection.`,
		Exposed: true,
		Name:    "observedDropReason",
		Type:    "string",
	},
	"observedencrypted": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedEncrypted",
		Description:    `Value of the encryption of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedEncrypted",
		Type:           "boolean",
	},
	"observedpolicyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyID",
		Description:    `ID of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyID",
		Type:           "string",
	},
	"observedpolicynamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ObservedPolicyNamespace",
		Description:    `Namespace of the network policy that observed the flow.`,
		Exposed:        true,
		Name:           "observedPolicyNamespace",
		Type:           "string",
	},
	"policyid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyID",
		Description:    `ID of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyID",
		Required:       true,
		Type:           "string",
	},
	"policynamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "PolicyNamespace",
		Description:    `Namespace of the network policy that accepted the flow.`,
		Exposed:        true,
		Name:           "policyNamespace",
		Required:       true,
		Type:           "string",
	},
	"protocol": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Protocol",
		Description:    `protocol number.`,
		Exposed:        true,
		Name:           "protocol",
		Required:       true,
		Type:           "integer",
	},
	"serviceclaimhash": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceClaimHash",
		Description:    `Hash of the claims used to communicate.`,
		Exposed:        true,
		Name:           "serviceClaimHash",
		Type:           "string",
	},
	"serviceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceID",
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceID",
		Type:           "string",
	},
	"servicenamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceNamespace",
		Description:    `Service URL accessed.`,
		Exposed:        true,
		Name:           "serviceNamespace",
		Type:           "string",
	},
	"servicetype": elemental.AttributeSpecification{
		AllowedChoices: []string{"L3", "HTTP", "TCP", "NotApplicable"},
		ConvertedName:  "ServiceType",
		DefaultValue:   FlowReportServiceTypeNotApplicable,
		Description:    `ID of the service.`,
		Exposed:        true,
		Name:           "serviceType",
		Type:           "enum",
	},
	"serviceurl": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ServiceURL",
		Description:    `Service URL accessed.`,
		Exposed:        true,
		Name:           "serviceURL",
		Type:           "string",
	},
	"sourceid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceID",
		Description:    `ID of the source.`,
		Exposed:        true,
		Name:           "sourceID",
		Required:       true,
		Type:           "string",
	},
	"sourceip": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceIP",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceIP",
		Type:           "string",
	},
	"sourcenamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "SourceNamespace",
		Description:    `Namespace of the receiver.`,
		Exposed:        true,
		Name:           "sourceNamespace",
		Type:           "string",
	},
	"sourcetype": elemental.AttributeSpecification{
		AllowedChoices: []string{"ProcessingUnit", "ExternalNetwork", "Claims"},
		ConvertedName:  "SourceType",
		Description:    `Type of the source.`,
		Exposed:        true,
		Name:           "sourceType",
		Required:       true,
		Type:           "enum",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Type:           "time",
	},
	"value": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Value",
		Description:    `Number of flows in the report.`,
		Exposed:        true,
		Name:           "value",
		Required:       true,
		Type:           "integer",
	},
}

// SparseFlowReportsList represents a list of SparseFlowReports
type SparseFlowReportsList []*SparseFlowReport

// Identity returns the identity of the objects in the list.
func (o SparseFlowReportsList) Identity() elemental.Identity {

	return FlowReportIdentity
}

// Copy returns a pointer to a copy the SparseFlowReportsList.
func (o SparseFlowReportsList) Copy() elemental.Identifiables {

	copy := append(SparseFlowReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the SparseFlowReportsList.
func (o SparseFlowReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(SparseFlowReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*SparseFlowReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o SparseFlowReportsList) List() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseFlowReportsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseFlowReportsList converted to FlowReportsList.
func (o SparseFlowReportsList) ToPlain() elemental.IdentifiablesList {

	out := make(elemental.IdentifiablesList, len(o))
	for i := 0; i < len(o); i++ {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseFlowReportsList) Version() int {

	return 1
}

// SparseFlowReport represents the sparse version of a flowreport.
type SparseFlowReport struct {
	// Action applied to the flow.
	Action *FlowReportActionValue `json:"action,omitempty" bson:"-" mapstructure:"action,omitempty"`

	// ID of the destination.
	DestinationID *string `json:"destinationID,omitempty" bson:"-" mapstructure:"destinationID,omitempty"`

	// Type of the destination.
	DestinationIP *string `json:"destinationIP,omitempty" bson:"-" mapstructure:"destinationIP,omitempty"`

	// Namespace of the receiver.
	DestinationNamespace *string `json:"destinationNamespace,omitempty" bson:"-" mapstructure:"destinationNamespace,omitempty"`

	// Port of the destination.
	DestinationPort *int `json:"destinationPort,omitempty" bson:"-" mapstructure:"destinationPort,omitempty"`

	// Type of the source.
	DestinationType *FlowReportDestinationTypeValue `json:"destinationType,omitempty" bson:"-" mapstructure:"destinationType,omitempty"`

	// This field is only set if 'action' is set to 'Reject' and specifies the reason
	// for the rejection.
	DropReason *string `json:"dropReason,omitempty" bson:"-" mapstructure:"dropReason,omitempty"`

	// Tells is the flow has been encrypted.
	Encrypted *bool `json:"encrypted,omitempty" bson:"-" mapstructure:"encrypted,omitempty"`

	// This is here for backward compatibility.
	Namespace *string `json:"namespace,omitempty" bson:"-" mapstructure:"namespace,omitempty"`

	// Tells if the flow is from design mode.
	Observed *bool `json:"observed,omitempty" bson:"-" mapstructure:"observed,omitempty"`

	// Action observed on the flow.
	ObservedAction *FlowReportObservedActionValue `json:"observedAction,omitempty" bson:"-" mapstructure:"observedAction,omitempty"`

	// This field is only set if 'observedAction' is set to 'Reject' and specifies the
	// reason for the rejection.
	ObservedDropReason *string `json:"observedDropReason,omitempty" bson:"-" mapstructure:"observedDropReason,omitempty"`

	// Value of the encryption of the network policy that observed the flow.
	ObservedEncrypted *bool `json:"observedEncrypted,omitempty" bson:"-" mapstructure:"observedEncrypted,omitempty"`

	// ID of the network policy that observed the flow.
	ObservedPolicyID *string `json:"observedPolicyID,omitempty" bson:"-" mapstructure:"observedPolicyID,omitempty"`

	// Namespace of the network policy that observed the flow.
	ObservedPolicyNamespace *string `json:"observedPolicyNamespace,omitempty" bson:"-" mapstructure:"observedPolicyNamespace,omitempty"`

	// ID of the network policy that accepted the flow.
	PolicyID *string `json:"policyID,omitempty" bson:"-" mapstructure:"policyID,omitempty"`

	// Namespace of the network policy that accepted the flow.
	PolicyNamespace *string `json:"policyNamespace,omitempty" bson:"-" mapstructure:"policyNamespace,omitempty"`

	// protocol number.
	Protocol *int `json:"protocol,omitempty" bson:"-" mapstructure:"protocol,omitempty"`

	// Hash of the claims used to communicate.
	ServiceClaimHash *string `json:"serviceClaimHash,omitempty" bson:"-" mapstructure:"serviceClaimHash,omitempty"`

	// ID of the service.
	ServiceID *string `json:"serviceID,omitempty" bson:"-" mapstructure:"serviceID,omitempty"`

	// Service URL accessed.
	ServiceNamespace *string `json:"serviceNamespace,omitempty" bson:"-" mapstructure:"serviceNamespace,omitempty"`

	// ID of the service.
	ServiceType *FlowReportServiceTypeValue `json:"serviceType,omitempty" bson:"-" mapstructure:"serviceType,omitempty"`

	// Service URL accessed.
	ServiceURL *string `json:"serviceURL,omitempty" bson:"-" mapstructure:"serviceURL,omitempty"`

	// ID of the source.
	SourceID *string `json:"sourceID,omitempty" bson:"-" mapstructure:"sourceID,omitempty"`

	// Type of the source.
	SourceIP *string `json:"sourceIP,omitempty" bson:"-" mapstructure:"sourceIP,omitempty"`

	// Namespace of the receiver.
	SourceNamespace *string `json:"sourceNamespace,omitempty" bson:"-" mapstructure:"sourceNamespace,omitempty"`

	// Type of the source.
	SourceType *FlowReportSourceTypeValue `json:"sourceType,omitempty" bson:"-" mapstructure:"sourceType,omitempty"`

	// Date of the report.
	Timestamp *time.Time `json:"timestamp,omitempty" bson:"-" mapstructure:"timestamp,omitempty"`

	// Number of flows in the report.
	Value *int `json:"value,omitempty" bson:"-" mapstructure:"value,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	*sync.Mutex `json:"-" bson:"-"`
}

// NewSparseFlowReport returns a new  SparseFlowReport.
func NewSparseFlowReport() *SparseFlowReport {
	return &SparseFlowReport{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseFlowReport) Identity() elemental.Identity {

	return FlowReportIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseFlowReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseFlowReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *SparseFlowReport) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseFlowReport) ToPlain() elemental.PlainIdentifiable {

	out := NewFlowReport()
	if o.Action != nil {
		out.Action = *o.Action
	}
	if o.DestinationID != nil {
		out.DestinationID = *o.DestinationID
	}
	if o.DestinationIP != nil {
		out.DestinationIP = *o.DestinationIP
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
	if o.SourceID != nil {
		out.SourceID = *o.SourceID
	}
	if o.SourceIP != nil {
		out.SourceIP = *o.SourceIP
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

// DeepCopy returns a deep copy if the SparseFlowReport.
func (o *SparseFlowReport) DeepCopy() *SparseFlowReport {

	if o == nil {
		return nil
	}

	out := &SparseFlowReport{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseFlowReport.
func (o *SparseFlowReport) DeepCopyInto(out *SparseFlowReport) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseFlowReport: %s", err))
	}

	*out = *target.(*SparseFlowReport)
}
