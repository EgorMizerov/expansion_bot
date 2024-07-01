package dto

import "github.com/EgorMizerov/expansion_bot/internal/domain/entity"

type DriverWithLicenseDTO struct {
	DriverDTO
	DriverLicenseDTO
}

func (self DriverWithLicenseDTO) ToEntity() *entity.Driver {
	return self.DriverDTO.ToDriver(self.DriverLicenseDTO.ToDriverLicense())
}
