package template

import (
	"bytes"
	"text/template"
)

func ParseTemplate(text string, data any) string {
	template, err := template.New("some_template").Parse(text)
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	template.Execute(&buf, data)
	return buf.String()
}

var DriverRegistrationTemplate = `{{ $defaultTime := "0001-01-01 00:00:00 +0000 UTC" }}
ID: {{ if ne 0 .UserID -}} {{ .UserID }} {{else}}{{end}}
ФИО: {{ .FullName }}
Номер телефона: {{ .PhoneNumber }}
Адрес проживания: {{ .Address }}
Статус самозанятасти: {{ if .IsSelfEmployed -}} Активен {{- else -}} Не активен {{- end }}
Выбранный тариф: {{ .WorkRule.Name }}
Номер карты: {{ .CardNumber }}

Водительский стаж с: {{ .FormattedTime .DrivingExperience }}
Серия и номер ВУ: {{ .RegistrationCertificate }}
Страна выдачи ВУ: {{ .LicenseCountry }}
Дата выдачи ВУ: {{ .FormattedTime .LicenseIssueDate }}
Дата окончания действия ВУ: {{ .FormattedTime .LicenseExpiryDate }}

Марка автомобиля: {{ .CarBrand }}
Модель автомобиля: {{ .CarModel }}
Цвет автомобиля {{ .CarColor }}
Год выпуска автомобиля: {{.FormattedYear }}
VIN автомобиля: {{ .CarVIN }}
Гос. Номер автомобиля: {{ .LicensePlateNumber }}

Реферальный ключ: {{ if eq .ReferralKey 0 }} Отсутствует {{else}} {{ .ReferralKey }} {{end}}

{{ .InputMessage }}
`

var GetCardsInfo = `Кол-во человек с картой Т-Банка {{ .TinkoffCardsCount }}
Кол-во человек с картой другого банка {{ .AnotherCardsCount }}`

var RegistrationApplicationsTemplate = `
{{- if .Items -}}
Заявки на регистрацию

{{ range .Items -}}
{{ .FormattedTime .Date }} [{{ .Fullname }}]({{ .Link }})
{{ end -}}
{{- else -}}
Список заявок пуст!
{{- end -}}`

var DriversListTemplate = `Для получения полной информации о водителе
введите его номер телефона

{{ range .Items -}}
{{ .Fullname }} {{ .PhoneNumber }}
{{ end -}}`

var DriverInfoTemplate = `Информация о водителе

ID: {{ .ID }}
ФИО: {{ .Fullname }}
Номер телефона: {{ .PhoneNumber }}
Город: {{ .City }}
Статус самозанятасти: {{ if .IsSelfEmployed -}} Активен {{- else -}} Не активен {{- end }}
Выбранный тариф: {{ .Nullable .WorkRule }}`

var DriversCarInfoTemplate = `Информация об автомобиле

Марка автомобиля: {{ .Brand }}
Модель автомобиля: {{ .Model }}
Цвет автомобиля {{ .Color }}
Год выпуска автомобиля: {{ .Year }}
VIN автомобиля: {{ .VIN }}
Гос\. Номер автомобиля: {{ .Number }}`
