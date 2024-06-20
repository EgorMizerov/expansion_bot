package request

import "time"

type GetDriversSummaryRequest struct {
	From     time.Time
	To       time.Time
	DriverId *string
	SortBy   GetDriversSummarySortBy
}

func (self GetDriversSummaryRequest) ToBody() GetDriversSummaryRequestBody {
	return GetDriversSummaryRequestBody{
		DateFrom: self.From.Format(time.DateOnly),
		DateTo:   self.To.Format(time.DateOnly),
		DriverID: self.DriverId,
		Sort: GetDriversSummarySort{
			Direction: "desc",
			Field:     self.SortBy,
		},
	}
}

type GetDriversSummaryRequestBody struct {
	DateFrom string                `json:"date_from"`
	DateTo   string                `json:"date_to"`
	DriverID *string               `json:"driver_id,omitempty"`
	Sort     GetDriversSummarySort `json:"sort"`
}

type GetDriversSummarySort struct {
	Direction string                  `json:"direction"`
	Field     GetDriversSummarySortBy `json:"field"`
}

type GetDriversSummarySortBy string

const (
	GetDriversSummarySortByCountOrdersComplete GetDriversSummarySortBy = "count_orders_completed"
)
