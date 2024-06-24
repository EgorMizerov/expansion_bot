package testdata

import "encoding/json"

var GetDriversBody = func() map[string]interface{} {
	body := `{
    "limit": 1,
    "offset": 0,
    "total": 24,
    "driver_profiles": [
        {
            "updated_at": "2024-06-22T10:40:51.295047+00:00",
            "accounts": [
                {
                    "balance": "0.0000",
                    "balance_limit": "0.0000",
                    "currency": "RUB",
                    "id": "52dba695081d4335b8a44e50b0b06601",
                    "type": "current"
                }
            ],
            "car": {
                "amenities": [],
                "body_number": "XTAGFL110H0708806",
                "brand": "LADA (ВАЗ)",
                "callsign": "О944РУ159",
                "category": [
                    "comfort",
                    "econom",
                    "intercity",
                    "premium_suv"
                ],
                "color": "Белый",
                "id": "56726045d43093e0aa9fa50982fd68e6",
                "model": "Vesta",
                "normalized_number": "0944PY159",
                "number": "О944РУ159",
                "registration_cert": "9957178151",
                "rental": false,
                "status": "pending",
                "vin": "XTAGFL110R0708806",
                "year": 2023
            },
            "driver_profile": {
                "created_date": "2024-06-21T08:12:25.024+00:00",
                "first_name": "Сергей",
                "hire_date": "2024-06-21T00:00:00+00:00",
                "id": "52dba695081d4335b8a44e50b0b06601",
                "is_selfemployed": false,
                "has_contract_issue": false,
                "last_name": "Шилов",
                "driver_license": {
                    "birth_date": "1961-07-09T00:00:00+00:00",
                    "country": "rus",
                    "expiration_date": "2032-03-18T00:00:00+00:00",
                    "issue_date": "2022-03-18T00:00:00+00:00",
                    "normalized_number": "9928638375",
                    "number": "9928638375"
                },
                "middle_name": "Геннадиевич",
                "modified_date": "2024-06-21T08:15:07.14+00:00",
                "park_id": "c53efe4fccd149a19839dd6d5fcce124",
                "phones": [
                    "+79223307320"
                ],
                "work_rule_id": "e26a3cf21acfe01198d50030487e046b",
                "work_status": "working"
            },
            "current_status": {
                "status": "offline",
                "status_updated_at": "2024-06-22T11:00:18.236+00:00"
            }
        }
    ],
    "parks": [
        {
            "id": "c53efe4fccd149a19839dd6d5fcce124",
            "city": "Пермь",
            "name": "Экспансия"
        }
    ]
}`
	sec := map[string]interface{}{}
	if err := json.Unmarshal([]byte(body), &sec); err != nil {
		panic(err)
	}
	return sec
}()
