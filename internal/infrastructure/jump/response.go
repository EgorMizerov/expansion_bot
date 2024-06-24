package jump

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type ListedResponse[ItemType any] struct {
	Items []ItemType `json:"items"`
	Meta  Meta       `json:"meta"`
}

type Meta struct {
	Total       int `json:"total"`
	From        int `json:"from"`
	To          int `json:"to"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	ListPage    int `json:"list_page"`
}

func getListedResponse[ItemType any](ctx context.Context, client *JumpClient, method string, url URL, headerSetter func(request *http.Request)) (ListedResponse[ItemType], error) {
	request, _ := http.NewRequestWithContext(ctx, method, url.String(), nil)
	headerSetter(request)

	response, err := client.Do(request)
	if err != nil {
		return ListedResponse[ItemType]{}, errors.Wrap(err, "failed to do request")
	}
	var result ListedResponse[ItemType]
	bytes, err := io.ReadAll(response.Body)
	x := string(bytes)
	_ = x
	err = json.Unmarshal(bytes, &result)
	//err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return ListedResponse[ItemType]{}, errors.Wrap(err, "failed to decode response")
	}

	return result, nil
}

func expandPaginatedResponse[ItemType any](ctx context.Context, client *JumpClient, method string, url URL, headerSetter func(request *http.Request)) ([]ItemType, error) {
	var listedResponse []ItemType
	var page = 1

	for {
		response, err := getListedResponse[ItemType](ctx, client, method, url.SetPage(page), headerSetter)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get listed response")
		}
		if len(listedResponse) == 0 {
			listedResponse = make([]ItemType, 0, response.Meta.Total)
		}

		listedResponse = append(listedResponse, response.Items...)
		page++

		if response.Meta.CurrentPage >= response.Meta.ListPage {
			break
		}
	}

	return listedResponse, nil
}
