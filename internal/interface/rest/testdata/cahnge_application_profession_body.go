package testdata

const ChangeApplicationProfessionBody = `{
  "action": "change_application",
  "item": {
    "id": {{.id}},
    "profession": {
      "id": 1,
      "name": "Таксист"
    },
    "is_car_driver": true,
    "direction": {
      "id": 1,
      "slug": "taxi",
      "title": "Такси"
    }
  },
  "updated_at": "2024-06-19T23:48:57+03:00"
}`
