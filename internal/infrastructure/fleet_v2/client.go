package fleet2

import (
	"context"
	"net/http"

	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet_v2/request"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet_v2/types"
)

const (
	getDrivers = "/v1/parks/driver-profiles/list"
)

type FleetClient struct {
	httpClient *http.Client

	hostAPI string

	clientID string
	parkID   string
	apiKey   string
}

func NewFleetClient(hostAPI string, clientID string, parkID string, apiKey string) *FleetClient {
	return &FleetClient{httpClient: http.DefaultClient, hostAPI: hostAPI, clientID: clientID, parkID: parkID, apiKey: apiKey}
}

func (self *FleetClient) GetDrivers(ctx context.Context, body request.GetDriversRequest) (*ListedResponse[types.GetDriversItem], error) {
	return getListedResponse[types.GetDriversItem](ctx, self.httpClient, http.MethodPost, self.getURL(getDrivers), body.ToBody(self.parkID), self.setHeaders, "driver_profiles")
}

func (self *FleetClient) Do(request *http.Request) (*http.Response, error) {
	return self.httpClient.Do(request)
}

func (self *FleetClient) getURL(method string) URL {
	return URL(self.hostAPI + method)
}

func (self *FleetClient) setHeaders(request *http.Request) {
	request.Header.Set("X-API-Key", self.apiKey)
	request.Header.Set("X-Client-ID", self.clientID)
	request.Header.Set("X-Park-ID", self.parkID)
}
