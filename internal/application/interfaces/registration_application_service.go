package interfaces

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

//go:generate mockery --name RegistrationApplicationService
type RegistrationApplicationService interface {
	SaveRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error
	GetRegistrationApplication(ctx context.Context, applicationID entity.RegistrationApplicationID) (*entity.RegistrationApplication, error)
	GetRegistrationApplications(ctx context.Context) ([]*entity.RegistrationApplication, error)
	ConfirmRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error
}
