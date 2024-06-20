package request

import (
	"strconv"
	"time"
)

type CreateRegularChargesRequest struct {
	DriverID   string
	ChargingAt time.Time
	CarID      string
	DailyPrice int
}

func (self CreateRegularChargesRequest) ToBody() CreateRegularCharges {
	return CreateRegularCharges{
		DriverID:   self.DriverID,
		ChargingAt: self.ChargingAt.Format("2006-01-02T15:04:05.999+07:00"),
		Charging: Charging{
			Type:       "active_days",
			DailyPrice: strconv.Itoa(self.DailyPrice),
		},
		Asset: Asset{
			Type:  "car",
			CarID: self.CarID,
		},
		BalanceNotifyLimit: "0.00",
	}
}

type CreateRegularCharges struct {
	DriverID           string   `json:"driver_id"`
	ChargingAt         string   `json:"charging_at"`
	Asset              Asset    `json:"asset"`
	Charging           Charging `json:"charging"`
	BalanceNotifyLimit string   `json:"balance_notify_limit"`
}

type Asset struct {
	Type  string `json:"type"`
	CarID string `json:"car_id"`
}

type Charging struct {
	Type       string `json:"type"`
	DailyPrice string `json:"daily_price"`
}
