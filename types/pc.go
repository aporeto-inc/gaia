package types

const (
	// PCTimeTypeRelative is the relative time range type
	PCTimeTypeRelative PCTimeType = "relative"
	// PCTimeTypeAbsolute is the absolute time range type
	PCTimeTypeAbsolute PCTimeType = "absolute"
	// PCTimeUnitHour is the unit for hour
	PCTimeUnitHour PCTimeValueUnit = "hour"
	// PCTimeUnitDay is the unit for day
	PCTimeUnitDay PCTimeValueUnit = "day"
	// PCTimeUnitWeek is the unit for week
	PCTimeUnitWeek PCTimeValueUnit = "week"
)

// PCTimeType represetns the type of time range
type PCTimeType string

// PCTimeValueUnit represetns the type of time unit
type PCTimeValueUnit string

// PCRqlTimeRange represents Primsma Cloud's time range type
type PCRqlTimeRange struct {
	Type  PCTimeType  `json:"type,omitempty"`
	Value PCTimeValue `json:"value,omitempty"`
}

// PCTimeValue represents Prisma Cloud's time value type
type PCTimeValue struct {
	Amount    int              `json:"amount,omitempty"`
	Unit      *PCTimeValueUnit `json:"unit,omitempty"`
	StartTime int              `json:"startTime,omitempty"`
	EndTime   int              `json:"endTime,omitempty"`
}

// PCRqlTableData represents Prisma Cloud's table data type
type PCRqlTableData struct {
	TotalRows     int           `json:"totalRows,omitempty"`
	NextPageToken string        `json:"nextPageToken,omitempty"`
	Items         []interface{} `json:"items,omitempty"`
}
