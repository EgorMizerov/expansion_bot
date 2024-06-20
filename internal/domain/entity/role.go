package entity

type Role string

const (
	DriverRole Role = "driver"
	AdminRole  Role = "admin"
	GuestRole  Role = "guest"
)
