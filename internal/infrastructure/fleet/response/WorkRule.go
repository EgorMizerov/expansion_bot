package response

import (
	"time"
)

type WorkRulesResponse struct {
	LightWorkRule []LightWorkRule `json:"light_work_rules"`
}

type LightWorkRule struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	IsArchived bool      `json:"is_archived"`
	IsDefault  bool      `json:"is_default"`
}

type ContractorProfileID struct {
	ID string `json:"contractor_profile_id"`
}
