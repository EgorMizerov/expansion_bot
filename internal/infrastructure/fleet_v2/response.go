package fleet2

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/EgorMizerov/expansion_bot/internal/common"
	"github.com/mitchellh/mapstructure"
	"github.com/perimeterx/marshmallow"
	"github.com/pkg/errors"
)

type HTTPClient interface {
	Do(request *http.Request) (*http.Response, error)
}

type ListedResponse[ItemType any] struct {
	Items  []ItemType `json:"-"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Total  int        `json:"total"`
}

func getListedResponse[ItemType any](ctx context.Context, client HTTPClient, method string, url URL, body any, headerSetter func(request *http.Request), itemTag string) (*ListedResponse[ItemType], error) {
	var requestBody io.Reader
	if body != common.Default[any]() {
		requestBody = bytes.NewReader(common.MustMarshal(body))
	}

	request, _ := http.NewRequestWithContext(ctx, method, url.String(), requestBody)
	headerSetter(request)

	response, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to do request")
	}

	var result ListedResponse[ItemType]
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, errors.Wrap(err, "failed to read response body")
	}
	x := string(responseBytes)
	_ = x
	unmarshalled, err := marshmallow.Unmarshal(responseBytes, &result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response body")
	}

	items := unmarshalled[itemTag]
	var i []ItemType
	mapstructure.Decode(items, &i)
	result.Items = i

	return &result, nil
}

//func expandPaginatedResponse[ItemType any](ctx context.Context, client *FleetClient, method string, url URL, headerSetter func(request *http.Request)) ([]ItemType, error) {
//	var listedResponse []ItemType
//	var page = 1
//
//	for {
//		response, err := getListedResponse[ItemType](ctx, client, method, url.SetPage(page), headerSetter, nil)
//		if err != nil {
//			return nil, errors.Wrap(err, "failed to get listed response")
//		}
//		if len(listedResponse) == 0 {
//			listedResponse = make([]ItemType, 0, response.Total)
//		}
//
//		//listedResponse = append(listedResponse, response.Items...)
//		page++
//
//		if response.Offset >= response.Total {
//			break
//		}
//	}
//
//	return listedResponse, nil
//}
