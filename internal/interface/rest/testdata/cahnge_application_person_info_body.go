package testdata

const ChangeApplicationPersonInfoBody = `{
  "action": "change_application",
  "item": {
    "id": {{.id}},
    "person_info": {
      "last_name": "{{.last_name}}",
      "first_name": "{{.first_name}}",
      "middle_name": "{{.middle_name}}",
      "birthdate": null,
      "city": "{{.city}}",
      "driver_license_taxi": null,
      "comment": "",
      "files": null
    }
  },
  "updated_at": "2024-06-19T23:49:11+03:00"
}`
