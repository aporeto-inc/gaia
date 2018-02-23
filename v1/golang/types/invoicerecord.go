package types

// InvoiceRecordDetails is a service associated with an enforcer profile for a host.
type InvoiceRecordDetails struct {
	TimeStamp string `json:"timestamp" bson:"timestamp" mapstructure:"timestamp,omitempty"`
}

// InvoiceRecordDetailsList is a list of ProcessingUnitServices
type InvoiceRecordDetailsList []*InvoiceRecordDetails

// Validate will validate the types. TODO: Do we need a stub if there is nothing to validate ?
func (i InvoiceRecordDetailsList) Validate() error {
	return nil
}
