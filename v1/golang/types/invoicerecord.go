package types

// InvoiceRecordDetail is a service associated with an enforcer profile for a host.
type InvoiceRecordDetail struct {
	TimeStamp string `json:"timestamp" bson:"timestamp" mapstructure:"timestamp,omitempty"`
}

// InvoiceRecordDetailsList is a list of ProcessingUnitServices
type InvoiceRecordDetailsList []*InvoiceRecordDetail

// Validate will validate the types. TODO: Do we need a stub if there is nothing to validate ?
func (i InvoiceRecordDetailsList) Validate() error {
	return nil
}
