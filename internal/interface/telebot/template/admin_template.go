package template

import (
	"bytes"
	"text/template"
)

func ParseTemplate(text string, data any) string {
	template, _ := template.New("some_template").Parse(text)
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
