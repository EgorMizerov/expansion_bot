package jump

import (
	"context"
	"net/http"
	"testing"

	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/jump/testdata"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
)

type JumpClientTests struct {
	suite.Suite

	jumpClient *JumpClient
	ctx        context.Context
}

func (self *JumpClientTests) SetupTest() {
	self.ctx = context.Background()
	self.jumpClient = NewJumpClient("test_host", "test_client_key")
	httpmock.Activate()
}

func (self *JumpClientTests) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (self *JumpClientTests) TestGetDriverByPhoneNumber() {
	body, driver := testdata.GetDriverByPhoneNumberBody()
	httpmock.RegisterResponder(http.MethodGet, self.jumpClient.getURL(getDriverByPhoneURL).SetParam("PhoneNumber", driver.Phone).String(), func(request *http.Request) (*http.Response, error) {
		return httpmock.NewJsonResponse(http.StatusOK, body)
	})

	result, err := self.jumpClient.GetDriverByPhoneNumber(self.ctx, driver.Phone)

	self.NoError(err)
	self.Equal(driver, result)
}

func TestFleetClientTests(t *testing.T) {
	suite.Run(t, new(JumpClientTests))
}
