package interfaces

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

//go:generate mockery --name RegistrationApplicationService
type RegistrationApplicationService interface {
	GetRegistrationApplication(ctx context.Context, applicationID entity.RegistrationApplicationID) (*entity.RegistrationApplication, error)
	SaveRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error
	ConfirmRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error
}
