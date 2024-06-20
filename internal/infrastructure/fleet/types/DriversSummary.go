package types

type DriversSummary struct {
	Items []DriversSummaryItem `json:"items"`
}

type DriversSummaryItem struct {
	Driver              Driver `json:"driver"`
	CountOrdersComplete int    `json:"count_orders_completed"`
}
