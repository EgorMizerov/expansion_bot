package types

import "time"

type Payment struct {
	ID              int64       `json:"id"`
	CreatedAt       string      `json:"created_at"`
	Driver          Driver      `json:"driver"`
	Status          Status      `json:"status"`
	Amount          float64     `json:"amount"`
	AmountPaid      float64     `json:"amount_paid"`
	PaidAt          string      `json:"paid_at"`
	PayType         Aggregator  `json:"pay_type"`
	Integration     Aggregator  `json:"integration"`
	RequestType     Aggregator  `json:"request_type"`
	Agent           Agent       `json:"agent"`
	Aggregator      Aggregator  `json:"aggregator"`
	Commission      float64     `json:"commission"`
	CommissionBank  float64     `json:"commission_bank"`
	BankAccount     BankAccount `json:"bank_account"`
	Requisites      Aggregator  `json:"requisites"`
	WriteOffAccount Aggregator  `json:"write_off_account"`
	History         []History   `json:"history"`
}

type Agent struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Aggregator struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type BankAccount struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	Balance      float64       `json:"balance"`
	IsDefault    bool          `json:"is_default"`
	PaymentTypes []PaymentType `json:"payment_types"`
	Agent        Agent         `json:"agent"`
}

type PaymentType struct {
	ID    int64  `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type Driver struct {
	ID         int64     `json:"id"`
	FullName   string    `json:"full_name"`
	FirstName  string    `json:"first_name"`
	SecondName string    `json:"second_name"`
	LastName   string    `json:"last_name"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
}

type History struct {
	DateTime string `json:"date_time"`
	Action   string `json:"action"`
}

type Status struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Theme string `json:"theme"`
}
