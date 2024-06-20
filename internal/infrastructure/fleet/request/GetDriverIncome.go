package request

import (
	"time"
)

type GetDriverIncomeRequest struct {
	DateFrom time.Time
	DateTo time.Time
	DriverID string
}

func (self GetDriverIncomeRequest) ToBody() GetDriverIncomeRequestBody {
	request := GetDriverIncomeRequestBody{
		DateFrom: self.DateFrom.Format("2006-01-02T15:04:05.999+07:00"),
		DateTo:   self.DateTo.Format("2006-01-02T15:04:05.999+07:00"),
		DriverID: self.DriverID,
	}

	return request
}

type GetDriverIncomeRequestBody struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
	DriverID string `json:"driver_id"`
}
