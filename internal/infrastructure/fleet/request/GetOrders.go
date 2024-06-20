package request

import (
	"time"
)

type GetOrdersRequest struct {
	BookedAt *FromToTime
	DriverID *string
	Status   *string

	Limit  int
	Cursor *string
}

func (self GetOrdersRequest) ToBody(parkID string) GetOrdersRequestBody {
	request := GetOrdersRequestBody{
		Cursor: self.Cursor,
		Limit:  self.Limit,
		Query: GetOrdersQuery{
			Park: GetOrdersQueryPark{
				ID:    parkID,
				Order: GetOrdersQueryParkOrder{
					//Statuses: func() *[]string {
					//	if self.Status != nil {
					//		slc := []string{*self.Status}
					//		return &slc
					//	}
					//	return nil
					//}(),
				},
			},
		},
	}

	if self.BookedAt != nil {
		request.Query.Park.Order.BookedAt = self.BookedAt
	}
	if self.DriverID != nil {
		request.Query.Park.DriverProfile = &GetOrdersQueryParkDriverProfile{
			ID: *self.DriverID,
		}
	}
	if self.Status != nil {
		status := []string{*self.Status}
		request.Query.Park.Order.Statuses = &status
	}

	return request
}

type GetOrdersRequestBody struct {
	Cursor *string        `json:"cursor,omitempty"`
	Limit  int            `json:"limit,omitempty"`
	Query  GetOrdersQuery `json:"query,omitempty"`
}

type GetOrdersQuery struct {
	Park GetOrdersQueryPark `json:"park,omitempty"`
}

type GetOrdersQueryPark struct {
	ID            string                           `json:"id,omitempty"`
	Order         GetOrdersQueryParkOrder          `json:"order,omitempty"`
	DriverProfile *GetOrdersQueryParkDriverProfile `json:"driver_profile,omitempty"`
}

type GetOrdersQueryParkOrder struct {
	BookedAt *FromToTime `json:"booked_at,omitempty"`
	EndedAt  *FromToTime `json:"ended_at,omitempty"`
	Statuses *[]string   `json:"statuses,omitempty"`
}

type GetOrdersQueryParkDriverProfile struct {
	ID string `json:"id,omitempty"`
}

type FromToTime struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"to,omitempty"`
}
