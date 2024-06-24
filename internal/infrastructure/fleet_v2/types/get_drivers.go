package types

type GetDriversItem struct {
	DriverProfile GetDriversDriverProfile `json:"driver_profile" mapstructure:"driver_profile"`
	Car           GetDriversCar           `json:"car" mapstructure:"car"`
}

type GetDriversDriverProfile struct {
	ID     string   `json:"id" mapstructure:"id"`
	Phones []string `json:"phones" mapstructure:"phones"`
}

type GetDriversCar struct {
	ID string `json:"id"`
}
