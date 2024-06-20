package types

type Transaction struct {
	ID          int64                      `json:"id"`
	Date        string                     `json:"date"`
	Description string                     `json:"description"`
	Driver      GetTransactionsDriver      `json:"driver"`
	Type        GetTransactionsIntegration `json:"type"`
	Integration GetTransactionsIntegration `json:"integration"`
	Amount      int64                      `json:"amount"`
	PaymentID   int64                      `json:"payment_id"`
	Trip        GetTransactionsTrip        `json:"trip"`
}

type GetTransactionsDriver struct {
	ID         int64  `json:"id"`
	Phone      int64  `json:"phone"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	FullName   string `json:"full_name"`
	MiddleName string `json:"middle_name"`
}

type GetTransactionsIntegration struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GetTransactionsTrip struct {
	Date        string                     `json:"date"`
	AddressFrom string                     `json:"address_from"`
	AddressTo   string                     `json:"address_to"`
	Duration    int64                      `json:"duration"`
	Distance    int64                      `json:"distance"`
	PaymentType GetTransactionsIntegration `json:"payment_type"`
	Amount      GetTransactionsAmount      `json:"amount"`
	Comment     string                     `json:"comment"`
}

type GetTransactionsAmount struct {
	Order                int64   `json:"order"`
	Tip                  int64   `json:"tip"`
	Cash                 int64   `json:"cash"`
	Cashless             int64   `json:"cashless"`
	AggregatorCommission float64 `json:"aggregator_commission"`
	CustomerCommission   float64 `json:"customer_commission"`
	Total                float64 `json:"total"`
}
