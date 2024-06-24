package fleet2

import (
	"context"
	"net/http"
	"testing"

	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet_v2/request"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet_v2/testdata"
	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/fleet_v2/types"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

type FleetClientTests struct {
	suite.Suite

	fleetClient *FleetClient
	ctx         context.Context
}

func (self *FleetClientTests) SetupTest() {
	self.ctx = context.Background()
	self.fleetClient = NewFleetClient("test_host", "test_client_id", "test_park_id", "test_api_key")
	httpmock.Activate()
}

func (self *FleetClientTests) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (self *FleetClientTests) TestGetDrivers() {
	httpmock.RegisterResponder(http.MethodPost, self.fleetClient.getURL(getDrivers).String(), func(request *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(http.StatusOK, testdata.GetDriversBody)
	})

	drivers, err := self.fleetClient.GetDrivers(self.ctx, request.GetDriversRequest{
		OrderBy:        "order_by",
		OrderDirection: "test_order_direction",
		Limit:          1,
	})

	expected := &ListedResponse[types.GetDriversItem]{
		Items: []types.GetDriversItem{
			{
				DriverProfile: types.GetDriversDriverProfile{
					ID:     "52dba695081d4335b8a44e50b0b06601",
					Phones: []string{"+79223307320"},
				},
				Car: types.GetDriversCar{ID: "56726045d43093e0aa9fa50982fd68e6"},
			},
		},
		Limit:  1,
		Offset: 0,
		Total:  24,
	}
	self.NoError(err)
	self.Equal(expected, drivers)
}

func TestFleetClientTests(t *testing.T) {
	suite.Run(t, new(FleetClientTests))
}
