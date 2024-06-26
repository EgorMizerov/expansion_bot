package template

import (
	"testing"
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/common"
	"github.com/stretchr/testify/suite"
)

type AdminTemplateTests struct {
	suite.Suite
}

func (self *AdminTemplateTests) TestRegistrationApplicationsTemplate() {
	result := ParseTemplate(RegistrationApplicationsTemplate, RegistrationApplicationsData{
		Items: []RegistrationApplication{
			{
				Date:     time.Now(),
				Fullname: "Mark Robert",
				Link:     "https://my.jump.taxi/autoregistration/list/10",
			},
			{
				Date:     time.Now(),
				Fullname: "Elon Mask",
				Link:     "https://my.jump.taxi/autoregistration/list/22",
			}},
	})

	expected := "Заявки на регистрацию\n\n2024-06-26 [Mark Robert](https://my.jump.taxi/autoregistration/list/10)\n" +
		"2024-06-26 [Elon Mask](https://my.jump.taxi/autoregistration/list/22)\n"
	self.Equal(expected, result)
}

func (self *AdminTemplateTests) TestDriversListTemplate() {
	result := ParseTemplate(DriversListTemplate, DriversListData{
		Items: []DriversListItem{
			{
				PhoneNumber: "+79956908966",
				Fullname:    "Test Name",
			},
		},
	})

	expected := "Для получения полной информации о водителе\nвведите его номер телефона\n\nTest Name +79956908966\n"
	self.Equal(expected, result)
}

func (self *AdminTemplateTests) TestRegistrationApplicationsTemplateIfNoItems() {
	result := ParseTemplate(RegistrationApplicationsTemplate, RegistrationApplicationsData{})

	expected := "Список заявок пуст!"
	self.Equal(expected, result)
}

func (self *AdminTemplateTests) TestDriverInfoTemplate() {
	result := ParseTemplate(DriverInfoTemplate, DriverInfoData{
		ID:             "test_id",
		Fullname:       "test_fullname",
		PhoneNumber:    "test_phone_number",
		City:           "test_city",
		IsSelfEmployed: false,
		WorkRule:       common.Point("test_work_rule"),
	})

	expected := `Информация о водителе

ID: test_id
ФИО: test_fullname
Номер телефона: test_phone_number
Город: test_city
Статус самозанятасти: Не активен
Выбранный тариф: test_work_rule`
	self.Equal(expected, result)
}

func (self *AdminTemplateTests) TestDriverInfoTemplateIfSelfEmployed() {
	result := ParseTemplate(DriverInfoTemplate, DriverInfoData{
		ID:             "test_id",
		Fullname:       "test_fullname",
		PhoneNumber:    "test_phone_number",
		City:           "test_city",
		IsSelfEmployed: true,
		WorkRule:       common.Point("test_work_rule"),
	})

	expected := `Информация о водителе

ID: test_id
ФИО: test_fullname
Номер телефона: test_phone_number
Город: test_city
Статус самозанятасти: Активен
Выбранный тариф: test_work_rule`
	self.Equal(expected, result)
}

func (self *AdminTemplateTests) TestDriverInfoTemplateIfWorkRuleNil() {
	result := ParseTemplate(DriverInfoTemplate, DriverInfoData{
		ID:             "test_id",
		Fullname:       "test_fullname",
		PhoneNumber:    "test_phone_number",
		City:           "test_city",
		IsSelfEmployed: false,
		WorkRule:       nil,
	})

	expected := `Информация о водителе

ID: test_id
ФИО: test_fullname
Номер телефона: test_phone_number
Город: test_city
Статус самозанятасти: Не активен
Выбранный тариф: -`
	self.Equal(expected, result)
}

func TestAdminTemplateTests(t *testing.T) {
	suite.Run(t, new(AdminTemplateTests))
}
