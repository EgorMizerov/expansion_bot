package jump

import (
	"context"
	"net/http"

	"github.com/EgorMizerov/expansion_bot/internal/infrastructure/jump/types"
)

const (
	getPaymentsURL    = "/payments?page={{.Page}}"
	getTransactionURL = "/payments?page={{.Page}} {{range .DriverIDs }}"
)

type JumpClient struct {
	httpClient *http.Client

	hostAPI   string
	clientKey string
}

func NewJumpClient(hostAPI string, clientKey string) *JumpClient {
	return &JumpClient{httpClient: http.DefaultClient, hostAPI: hostAPI, clientKey: clientKey}
}

func (self *JumpClient) GetPayments(ctx context.Context) ([]types.Payment, error) {
	return expandPaginatedResponse[types.Payment](ctx, self, http.MethodGet, self.getURL(getPaymentsURL), self.setHeaders)
}

func (self *JumpClient) GetTransactions(ctx context.Context) ([]types.Transaction, error) {
	return expandPaginatedResponse[types.Transaction](ctx, self, http.MethodGet, self.getURL(getPaymentsURL), self.setHeaders)
}

func (self *JumpClient) Do(request *http.Request) (*http.Response, error) {
	return self.httpClient.Do(request)
}

func (self *JumpClient) getURL(method string) URL {
	return URL(self.hostAPI + method)
}

func (self *JumpClient) setHeaders(request *http.Request) {
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Client-Key", self.clientKey)
}
