package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// FlowReportIndexes lists the attribute compound indexes.
var FlowReportIndexes = [][]string{}

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

	// FlowReportDestinationTypeExternalService represents the value ExternalService.
	FlowReportDestinationTypeExternalService FlowReportDestinationTypeValue = "ExternalService"

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

	// FlowReportServiceTypeTCP represents the value TCP.
	FlowReportServiceTypeTCP FlowReportServiceTypeValue = "TCP"
)

// FlowReportSourceTypeValue represents the possible values for attribute "sourceType".
type FlowReportSourceTypeValue string

const (
	// FlowReportSourceTypeClaims represents the value Claims.
	FlowReportSourceTypeClaims FlowReportSourceTypeValue = "Claims"

	// FlowReportSourceTypeExternalService represents the value ExternalService.
	FlowReportSourceTypeExternalService FlowReportSourceTypeValue = "ExternalService"

	// FlowReportSourceTypeProcessingUnit represents the value ProcessingUnit.
	FlowReportSourceTypeProcessingUnit FlowReportSourceTypeValue = "ProcessingUnit"
)

// FlowReportIdentity represents the Identity of the object.
var FlowReportIdentity = elemental.Identity{
	Name:     "flowreport",
	Category: "flowreports",
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

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o FlowReportsList) DefaultOrder() []string {

	return []string{}
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

	// Reason for the rejection.
	DropReason string `json:"dropReason" bson:"-" mapstructure:"dropReason,omitempty"`

	// Tells is the flow has been encrypted.
	Encrypted bool `json:"encrypted" bson:"-" mapstructure:"encrypted,omitempty"`

	// This is here for backward compatibility.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Tells if the flow is from design mode.
	Observed bool `json:"observed" bson:"-" mapstructure:"observed,omitempty"`

	// Action observed on the flow.
	ObservedAction FlowReportObservedActionValue `json:"observedAction" bson:"-" mapstructure:"observedAction,omitempty"`

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

	sync.Mutex
}

// NewFlowReport returns a new *FlowReport
func NewFlowReport() *FlowReport {

	return &FlowReport{
		ModelVersion:   1,
		ObservedAction: "NotApplicable",
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

// Validate valides the current information stored into the structure.
func (o *FlowReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateStringInList("action", string(o.Action), []string{"Accept", "Reject"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("destinationID", o.DestinationID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("destinationType", string(o.DestinationType), []string{"ProcessingUnit", "ExternalService", "Claims"}, false); err != nil {
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

	if err := elemental.ValidateStringInList("serviceType", string(o.ServiceType), []string{"L3", "HTTP", "TCP"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("sourceID", o.SourceID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateStringInList("sourceType", string(o.SourceType), []string{"ProcessingUnit", "ExternalService", "Claims"}, false); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredTime("timestamp", o.Timestamp); err != nil {
		requiredErrors = append(requiredErrors, err)
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
		AllowedChoices: []string{"ProcessingUnit", "ExternalService", "Claims"},
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
		Description:    `Reason for the rejection.`,
		Exposed:        true,
		Name:           "dropReason",
		Type:           "string",
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
		AllowedChoices: []string{"L3", "HTTP", "TCP"},
		ConvertedName:  "ServiceType",
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
		AllowedChoices: []string{"ProcessingUnit", "ExternalService", "Claims"},
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
		Required:       true,
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
		AllowedChoices: []string{"ProcessingUnit", "ExternalService", "Claims"},
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
		Description:    `Reason for the rejection.`,
		Exposed:        true,
		Name:           "dropReason",
		Type:           "string",
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
		AllowedChoices: []string{"L3", "HTTP", "TCP"},
		ConvertedName:  "ServiceType",
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
		AllowedChoices: []string{"ProcessingUnit", "ExternalService", "Claims"},
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
		Required:       true,
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
