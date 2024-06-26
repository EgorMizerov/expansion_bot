package rest

import (
	"bytes"
	"context"
	"net/http"
	"testing"
	"text/template"
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/application/interfaces/mocks"
	"github.com/EgorMizerov/expansion_bot/internal/common"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/interface/rest/testdata"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
)

type JumpWebhookTests struct {
	suite.Suite

	handler                        http.Handler
	registrationApplicationService *mocks.RegistrationApplicationService
}

func (self *JumpWebhookTests) SetupTest() {
	self.registrationApplicationService = mocks.NewRegistrationApplicationService(self.T())
	self.handler = NewJumpWebhook(self.registrationApplicationService, nil)
}

func (self *JumpWebhookTests) TestWebhookNewApplication() {
	registrationApplication := entity.NewRegistrationApplication(222, time.Unix(1, 0))
	registrationApplication.SetPhone("+79956782366")

	self.registrationApplicationService.On("SaveRegistrationApplication", context.Background(), registrationApplication).
		Return(nil)

	body := self.parseBody(testdata.NewApplicationBody, map[string]interface{}{
		"id":    registrationApplication.ID,
		"date":  registrationApplication.Date.Format("2006-01-02T15:04:05-07:00"),
		"phone": registrationApplication.PhoneNumber,
	})

	apitest.New("TestWebhookNewApplication").
		HandlerFunc(self.handler.ServeHTTP).
		Post("/webhook").
		Body(body).
		Expect(self.T()).
		Status(http.StatusOK).
		End()
}

func (self *JumpWebhookTests) TestWebhookChangeApplicationProfession() {
	registrationApplication := entity.NewRegistrationApplication(222, time.Unix(1, 0))

	self.registrationApplicationService.On("GetRegistrationApplication", context.Background(), registrationApplication.ID).
		Return(registrationApplication, nil)
	self.registrationApplicationService.On("SaveRegistrationApplication", context.Background(), registrationApplication).
		Return(nil)

	body := self.parseBody(testdata.ChangeApplicationProfessionBody, map[string]interface{}{
		"id": registrationApplication.ID,
	})

	apitest.New("TestWebhookChangeApplicationProfession").
		HandlerFunc(self.handler.ServeHTTP).
		Post("/webhook").
		Body(body).
		Expect(self.T()).
		Status(http.StatusOK).
		End()
}

func (self *JumpWebhookTests) TestWebhookChangeApplicationPersonInfo() {
	registrationApplication := entity.NewRegistrationApplication(222, time.Unix(1, 0))
	savedRegistrationApplication := *registrationApplication
	savedRegistrationApplication.SetPersonInfo("Иванов", "Иван", common.Point("Иванович"), "Пермь")

	self.registrationApplicationService.On("GetRegistrationApplication", context.Background(), registrationApplication.ID).
		Return(registrationApplication, nil)
	self.registrationApplicationService.On("SaveRegistrationApplication", context.Background(), &savedRegistrationApplication).
		Return(nil)

	body := self.parseBody(testdata.ChangeApplicationPersonInfoBody, map[string]interface{}{
		"id":          savedRegistrationApplication.ID,
		"last_name":   savedRegistrationApplication.LastName,
		"first_name":  savedRegistrationApplication.FirstName,
		"middle_name": savedRegistrationApplication.MiddleName,
		"city":        savedRegistrationApplication.City,
	})

	apitest.New("TestWebhookChangeApplicationPersonInfo").
		HandlerFunc(self.handler.ServeHTTP).
		Post("/webhook").
		Body(body).
		Expect(self.T()).
		Status(http.StatusOK).
		End()
}

func (self *JumpWebhookTests) TestWebhookChangeApplicationDriverLicense() {
	nowDate := time.Date(2024, 06, 23, 0, 0, 0, 0, time.UTC)
	registrationApplication := entity.NewRegistrationApplication(222, time.Unix(1, 0))
	savedRegistrationApplication := *registrationApplication
	savedRegistrationApplication.SetDriverLicense("rus", "112214212", nowDate, nowDate, nowDate)

	self.registrationApplicationService.On("GetRegistrationApplication", context.Background(), registrationApplication.ID).
		Return(registrationApplication, nil)
	self.registrationApplicationService.On("SaveRegistrationApplication", context.Background(), &savedRegistrationApplication).
		Return(nil)

	body := self.parseBody(testdata.ChangeApplicationDriverLicenseBody, map[string]interface{}{
		"id":               savedRegistrationApplication.ID,
		"country":          savedRegistrationApplication.LicenseCountry,
		"number":           savedRegistrationApplication.LicenseNumber,
		"issue_date":       savedRegistrationApplication.LicenseIssueDate.Format("2006-01-02"),
		"expiry_date":      savedRegistrationApplication.LicenseExpiryDate.Format("2006-01-02"),
		"total_since_date": savedRegistrationApplication.LicenseTotalSinceDate.Format("2006-01-02"),
	})

	apitest.New("TestWebhookChangeApplicationDriverLicense").
		HandlerFunc(self.handler.ServeHTTP).
		Post("/webhook").
		Body(body).
		Expect(self.T()).
		Status(http.StatusOK).
		End()
}

func (self *JumpWebhookTests) TestWebhookChangeApplicationCar() {
	registrationApplication := entity.NewRegistrationApplication(222, time.Unix(1, 0))
	savedRegistrationApplication := *registrationApplication
	savedRegistrationApplication.SetCar(2003, "test_number", "test_license", "test_vin", "test_brand", "test_model", "test_color")

	self.registrationApplicationService.On("GetRegistrationApplication", context.Background(), registrationApplication.ID).
		Return(registrationApplication, nil)
	self.registrationApplicationService.On("SaveRegistrationApplication", context.Background(), &savedRegistrationApplication).
		Return(nil)

	body := self.parseBody(testdata.ChangeApplicationCarBody, map[string]interface{}{
		"id":      savedRegistrationApplication.ID,
		"year":    savedRegistrationApplication.CarYear,
		"number":  savedRegistrationApplication.CarNumber,
		"license": savedRegistrationApplication.CarLicense,
		"vin":     savedRegistrationApplication.CarVIN,
		"brand":   savedRegistrationApplication.CarBrand,
		"model":   savedRegistrationApplication.CarModel,
		"color":   savedRegistrationApplication.CarColor,
	})

	apitest.New("TestWebhookChangeApplicationCar").
		HandlerFunc(self.handler.ServeHTTP).
		Post("/webhook").
		Body(body).
		Expect(self.T()).
		Status(http.StatusOK).
		End()
}

func (self *JumpWebhookTests) TestWebhookChangeApplicationStatus() {
	registrationApplication := entity.NewRegistrationApplication(222, time.Unix(1, 0))
	savedRegistrationApplication := *registrationApplication
	savedRegistrationApplication.SetStatus("not_processed")

	self.registrationApplicationService.On("GetRegistrationApplication", context.Background(), registrationApplication.ID).
		Return(registrationApplication, nil)
	self.registrationApplicationService.On("SaveRegistrationApplication", context.Background(), &savedRegistrationApplication).
		Return(nil)

	body := self.parseBody(testdata.ChangeApplicationStatusBody, map[string]interface{}{
		"id":     savedRegistrationApplication.ID,
		"status": savedRegistrationApplication.Status,
	})

	apitest.New("TestWebhookChangeApplicationStatus").
		HandlerFunc(self.handler.ServeHTTP).
		Post("/webhook").
		Body(body).
		Expect(self.T()).
		Status(http.StatusOK).
		End()
}

func (self *JumpWebhookTests) parseBody(body string, args map[string]interface{}) string {
	var buf bytes.Buffer
	err := template.Must(template.New("").Parse(body)).
		Execute(&buf, args)
	if !self.NoError(err) {
		self.T().FailNow()
	}
	return buf.String()
}

func TestJumpWebhookTests(t *testing.T) {
	suite.Run(t, new(JumpWebhookTests))
}
