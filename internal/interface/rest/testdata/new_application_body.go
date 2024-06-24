package testdata

const NewApplicationBody = `{
  "action": "new_application",
  "item": {
    "id": {{.id}},
    "status": {
      "id": 1,
      "slug": "draft",
      "title": "Черновик"
    },
    "date": "{{.date}}",
    "phone": "{{.phone}}",
    "referrer": null
  },
  "updated_at": "2024-06-19T23:48:55+03:00"
}`
