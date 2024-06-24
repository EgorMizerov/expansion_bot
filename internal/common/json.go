package common

import "encoding/json"

func MustMarshal(data any) []byte {
	result, _ := json.Marshal(data)
	return result
}
