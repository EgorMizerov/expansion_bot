package command

import "github.com/EgorMizerov/expansion_bot/internal/domain/entity"

type ConfirmRegistrationApplicationCommand struct {
	RegistrationApplication *entity.RegistrationApplication
	WorkRule                entity.WorkRule
	IsSelfEmployed          bool
}
