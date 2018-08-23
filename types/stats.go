package types

// TimeSeriesQueryResults represents the a query result.
type TimeSeriesQueryResults struct {
	Rows []*TimeSeriesRow `json:"rows"`
}

// NewTimeSeriesQueryResults returns a new TimeSeriesQueryResults.
func NewTimeSeriesQueryResults() *TimeSeriesQueryResults {
	return &TimeSeriesQueryResults{
		Rows: []*TimeSeriesRow{},
	}
}

// TimeSeriesRow represents a single Time Series row.
type TimeSeriesRow struct {
	Name    string            `json:"name"`
	Tags    map[string]string `json:"tags"`
	Columns []string          `json:"columns"`
	Values  [][]interface{}   `json:"values"`
}

// NewTimeSeriesRow returns a new TimeSeriesRow.
func NewTimeSeriesRow(name string, tags map[string]string, columns []string, values [][]interface{}) *TimeSeriesRow {
	return &TimeSeriesRow{
		Name:    name,
		Tags:    tags,
		Columns: columns,
		Values:  values,
	}
}
