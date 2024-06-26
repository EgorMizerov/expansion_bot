package repository

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type RegistrationApplicationRepository interface {
	SaveRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error
	GetRegistrationApplication(ctx context.Context, applicationID entity.RegistrationApplicationID) (*entity.RegistrationApplication, error)
	GetRegistrationApplications(ctx context.Context) ([]*entity.RegistrationApplication, error)
}
