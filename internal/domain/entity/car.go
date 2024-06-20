package entity

type CarID string

type Car struct {
	ID                      CarID
	Brand                   string
	Model                   string
	Year                    int
	Color                   string
	VIN                     string
	LicensePlateNumber      string
	RegistrationCertificate string
}

func NewCar(brand string, model string, year int, color string, VIN string, licensePlateNumber string, registrationCertificate string) *Car {
	return &Car{Brand: brand, Model: model, Year: year, Color: color, VIN: VIN, LicensePlateNumber: licensePlateNumber, RegistrationCertificate: registrationCertificate}
}

func (self *Car) SetID(id CarID) {
	self.ID = id
}
