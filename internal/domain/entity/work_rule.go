package entity

type WorkRule struct {
	ID        string
	IsEnabled bool
	Name      string
}

func (self *WorkRule) StringPointer() *string {
	if self == nil {
		return nil
	}
	return &self.Name
}

var (
	FixSelfEmployedWorkRule = WorkRule{
		ID:        "3b669e4ffb4a4803a42adaf2fe1c777e",
		IsEnabled: true,
		Name:      "Фиксированный(СМЗ)",
	}
	PercentSelfEmployedWorkRule = WorkRule{
		ID:        "7916d5a62e144072908900bb228b41dc",
		IsEnabled: true,
		Name:      "Процент(СМЗ)",
	}
	PercentWorkRule = WorkRule{
		ID:        "8f0ae1a24d3a413b91d118ed702b4ff5",
		IsEnabled: true,
		Name:      "Процент",
	}
	PerDayWorkRule = WorkRule{
		ID:        "b56444570d314ff6bc028b36af54f2fa",
		IsEnabled: true,
		Name:      "Суточный",
	}
	FixWorkRule = WorkRule{
		ID:        "e26a3cf21acfe01198d50030487e046b",
		IsEnabled: true,
		Name:      "Фиксированный",
	}
)

func WorkRuleFromID(id string) WorkRule {
	switch id {
	case "3b669e4ffb4a4803a42adaf2fe1c777e":
		return FixSelfEmployedWorkRule
	case "7916d5a62e144072908900bb228b41dc":
		return PercentSelfEmployedWorkRule
	case "8f0ae1a24d3a413b91d118ed702b4ff5":
		return PercentWorkRule
	case "b56444570d314ff6bc028b36af54f2fa":
		return PerDayWorkRule
	case "e26a3cf21acfe01198d50030487e046b":
		return FixWorkRule
	}
	return WorkRule{}
}
