package testdata

const ChangeApplicationDriverLicenseBody = `{
  "action": "change_application",
  "item": {
    "id": {{.id}},
    "driver_license": {
      "country": {
        "value": "{{.country}}",
        "title": "Россия"
      },
      "number": "{{.number}}",
      "issue_date": "{{.issue_date}}",
      "expiry_date": "{{.expiry_date}}",
      "expiry_date_unlimited": false,
      "total_since_date": "{{.total_since_date}}",
      "files": null
    }
  },
  "updated_at": "2024-06-19T23:49:41+03:00"
}`
