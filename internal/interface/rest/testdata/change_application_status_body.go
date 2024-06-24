package testdata

const ChangeApplicationStatusBody = `{
  "action": "change_application",
  "item": {
    "id": {{.id}},
    "status": {
      "id": 4,
      "slug": "{{.status}}",
      "title": "Не обработана"
    }
  },
  "updated_at": "2024-06-19T23:50:24+03:00"
}`
