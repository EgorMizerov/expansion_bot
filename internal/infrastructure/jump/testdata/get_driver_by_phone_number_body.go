package testdata

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/jump/types"
)

func GetDriverByPhoneNumberBody() (map[string]interface{}, types.Driver) {
	driver := types.Driver{
		ID:        11328140,
		FullName:  "Мизеров Егор",
		FirstName: "Егор",
		LastName:  "Мизеров",
		Phone:     "+79926788923",
		CreatedAt: time.Unix(1000000, 0),
	}

	body := fmt.Sprintf(`{
    "items": [
        {
            "id": %d,
            "phone": "%s",
            "last_name": "%s",
            "first_name": "%s",
            "middle_name": null,
            "full_name": "%s",
            "tooltip": [],
            "balance": "0.00",
            "group": {
                "id": 34059,
                "name": "По умолчанию"
            },
            "anti_fraud_status": {
                "id": 1,
                "name": "Вывод разрешен"
            },
            "invited_by": null,
            "integrations": [
                {
                    "id": 18883,
                    "name": "YanGo"
                }
            ],
            "in_blacklist": false,
            "created_at": "%s",
            "first_trip_at": null,
            "last_trip_at": null
        }
    ],
    "meta": {
        "total": 1,
        "from": 1,
        "to": 1,
        "per_page": 20,
        "current_page": 1,
        "last_page": 1
    }
}`, driver.ID, driver.Phone, driver.LastName, driver.FirstName, driver.FullName, driver.CreatedAt.Format("2006-01-02T15:04:05-07:00"))
	sec := map[string]interface{}{}
	if err := json.Unmarshal([]byte(body), &sec); err != nil {
		panic(err)
	}
	return sec, driver
}
