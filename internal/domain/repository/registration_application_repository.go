package repository

import (
	"context"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type RegistrationApplicationRepository interface {
	GetRegistrationApplication(ctx context.Context, applicationID entity.RegistrationApplicationID) (*entity.RegistrationApplication, error)
	SaveRegistrationApplication(ctx context.Context, application *entity.RegistrationApplication) error
}
