package testdata

const ChangeApplicationCarBody = `{
  "action": "change_application",
  "item": {
    "id": {{.id}},
    "car": {
      "year": {{.year}},
      "number": "{{.number}}",
      "license": "{{.license}}",
      "vin": "{{.vin}}",
      "no_vin": false,
      "brand": "{{.brand}}",
      "model": {
        "id": 167,
        "name": "{{.model}}"
      },
      "color": {
        "id": 3,
        "name": "{{.color}}"
      },
      "files": null
    }
  },
  "updated_at": "2024-06-19T23:50:24+03:00"
}`
