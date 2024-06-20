package fleet

type GetContractorsResponse struct {
	Contractors []Contactor `json:"contractors"`
}

type Contactor struct {
	ID string `json:"id"`
}

type GetDriverIncomeResponse struct {
	Balances GetDriverIncomeBalances `json:"balances"`
	Orders   GetDriverIncomeOrders   `json:"orders"`
	WorkTime GetDriverIncomeWorkTime `json:"work_time"`
}

type GetDriverIncomeBalances struct {
	Total         float64 `json:"total"`
	PlatformCard  float64 `json:"platform_card"`
	CashCollected float64 `json:"cash_collected"`
	PlatformTip   float64 `json:"platform_tip"`
}

type GetDriverIncomeOrders struct {
	CountCompleted float64 `json:"count_completed"`
	Mileage        float64 `json:"mileage"`
	Price          float64 `json:"price"`
}

type GetDriverIncomeWorkTime struct {
	Mph     float64 `json:"mph"`
	Seconds int64   `json:"seconds"`
}
