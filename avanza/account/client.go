package account

import (
	"context"
	"net/http"

	"github.com/open-wallstreet/go-avanza/avanza/client"
	"github.com/open-wallstreet/go-avanza/avanza/models"
)

const (
	OverviewPath          = "/_api/account-performance/overview/total-values"
	AccountOverviewPath   = "/_mobile/account/{accountId}/overview"
	GetPositionsPath      = "/_api/position-data/positions"
	GetDealsAndOrdersPath = "/_mobile/account/dealsandorders"
	GetTransactionsPath   = "/_mobile/account/transactions/{accountId}"
	GetMyCompanyEvents    = "/_cqbe/market/my-company-events"
)

type AccountClient struct {
	*client.Client
}

// GetOverview retrieves overview of all accounts
func (a *AccountClient) GetOverview(ctx context.Context, options ...models.RequestOption) (*models.OverviewResponse, error) {
	res := &models.OverviewResponse{}
	params := &models.OverviewParams{}
	err := a.Call(ctx, http.MethodGet, OverviewPath, params, res, options...)
	return res, err
}

// GetAccountOverview get overview of a specified account.
func (a *AccountClient) GetAccountOverview(ctx context.Context, params *models.AccountOverviewParams, options ...models.RequestOption) (*models.AccountOverviewResponse, error) {
	res := &models.AccountOverviewResponse{}
	err := a.Call(ctx, http.MethodGet, AccountOverviewPath, params, res, options...)
	return res, err
}

// GetPositions gets all positions for your account
func (a *AccountClient) GetPositions(ctx context.Context, options ...models.RequestOption) (*models.GetPositionsResponse, error) {
	res := &models.GetPositionsResponse{}
	params := &models.GetPositionsParams{}
	err := a.Call(ctx, http.MethodGet, GetPositionsPath, params, res, options...)
	return res, err
}

// GetDealsAndOrders gets all deals and orders for your account that is currently active
func (a *AccountClient) GetDealsAndOrders(ctx context.Context, options ...models.RequestOption) (*models.GetDealsAndOrdersResponse, error) {
	res := &models.GetDealsAndOrdersResponse{}
	params := &models.GetDealsAndOrdersParams{}
	err := a.Call(ctx, http.MethodGet, GetDealsAndOrdersPath, params, res, options...)
	return res, err
}

// GetTransactions gets all your transactions. Can be filtered to include more or less data see models.GetTransactionsParams
func (a *AccountClient) GetTransactions(ctx context.Context, params *models.GetTransactionsParams, options ...models.RequestOption) (*models.GetTransactionsResponse, error) {
	res := &models.GetTransactionsResponse{}
	err := a.Call(ctx, http.MethodGet, GetTransactionsPath, params, res, options...)
	return res, err
}

// GetMyCompanyEvents Gets data about upcoming events companies you hold will have. (Dividends, splits etc)
func (a *AccountClient) GetMyCompanyEvents(ctx context.Context, options ...models.RequestOption) (*models.GetMyCompanyEventsResponse, error) {
	res := &models.GetMyCompanyEventsResponse{}
	type empty struct{}
	err := a.Call(ctx, http.MethodGet, GetMyCompanyEvents, empty{}, res, options...)
	return res, err
}
