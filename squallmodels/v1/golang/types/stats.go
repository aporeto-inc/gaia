package types

// TimeSeriesFunctionType is the type of TimerSeries Functions
type TimeSeriesFunctionType string

// Values of TimerSeriesFunctionsType
const (
	TimeSeriesFunctionTypeCount    TimeSeriesFunctionType = "COUNT"
	TimeSeriesFunctionTypeMean     TimeSeriesFunctionType = "MEAN"
	TimeSeriesFunctionTypeMedian   TimeSeriesFunctionType = "MEDIAN"
	TimeSeriesFunctionTypeDistinct TimeSeriesFunctionType = "DISTINCT"
	TimeSeriesFunctionTypeMax      TimeSeriesFunctionType = "MAX"
	TimeSeriesFunctionTypeMin      TimeSeriesFunctionType = "MIN"
)

// StatsFunctionType is the stats db function and arguments
type StatsFunctionType struct {
	Function  TimeSeriesFunctionType
	FieldName string
}
