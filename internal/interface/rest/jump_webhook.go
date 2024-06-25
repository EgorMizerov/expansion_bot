package rest

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/EgorMizerov/expansion_bot/internal/application/interfaces"
	"github.com/EgorMizerov/expansion_bot/internal/common"
	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
)

type JumpWebhook struct {
	registrationApplicationService interfaces.RegistrationApplicationService
}

func NewJumpWebhook(registrationApplicationService interfaces.RegistrationApplicationService) *JumpWebhook {
	return &JumpWebhook{registrationApplicationService: registrationApplicationService}
}

func (self *JumpWebhook) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	slog.Info(
		"",
		slog.String("method", request.Method),
		slog.String("path", request.URL.Path))

	if request.Method != http.MethodPost || request.URL.Path != "/webhook" {
		return
	}

	var callback Callback
	err := json.NewDecoder(request.Body).Decode(&callback)
	if err != nil {
		s := err.Error()
		_ = s
		return
	}

	switch callback.Action {
	case NewApplicationAction:
		application := entity.NewRegistrationApplication(callback.Item.ID, *callback.Item.Date)
		application.SetPhone(*callback.Item.Phone)

		err = self.registrationApplicationService.SaveRegistrationApplication(request.Context(), application)
		if err != nil {
			return
		}
	// TODO: Переделать с PUT на PATCH.
	case ChangeApplicationAction:
		application, err := self.registrationApplicationService.GetRegistrationApplication(request.Context(), entity.RegistrationApplicationID(callback.Item.ID))
		if err != nil {
			return
		}

		if callback.Item.Status != nil {
			application.SetStatus(callback.Item.Status.Slug)
			if callback.Item.Status.Slug == "registered" {
				err = self.registrationApplicationService.ConfirmRegistrationApplication(request.Context(), application)
				return
			}
		}
		if personInfo := callback.Item.PersonInfo; personInfo != nil {
			application.SetPersonInfo(
				personInfo.LastName,
				personInfo.FirstName,
				personInfo.MiddleName,
				personInfo.City,
			)
		}
		if license := callback.Item.DriverLicense; license != nil {
			application.SetDriverLicense(
				license.Country.Value,
				license.Number,
				common.ParseDate(license.IssueDate),
				common.ParseDate(license.ExpiryDate),
				common.ParseDate(license.TotalSinceDate),
			)
		}
		if car := callback.Item.Car; car != nil {
			application.SetCar(
				car.Year,
				car.Number,
				car.License,
				*car.VIN,
				car.Brand,
				car.Model.Name,
				car.Color.Name,
			)
		}

		err = self.registrationApplicationService.SaveRegistrationApplication(request.Context(), application)
		if err != nil {
			return
		}
	}

	writer.WriteHeader(http.StatusOK)
}
