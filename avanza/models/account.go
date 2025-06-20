package models

type GetMyCompanyEventsResponse struct {
	Dividends struct {
		PreliminaryDividends []struct {
			Position struct {
				OrderbookID    string `json:"orderbookId"`
				Name           string `json:"name"`
				FlagCode       string `json:"flagCode"`
				InstrumentName string `json:"instrumentName"`
			} `json:"position"`
			PaymentDate  string `json:"paymentDate"`
			ExDate       string `json:"exDate"`
			DividendType string `json:"dividendType"`
			Quantity     struct {
				Value            float64 `json:"value"`
				Unit             string  `json:"unit"`
				UnitType         string  `json:"unitType"`
				DecimalPrecision int     `json:"decimalPrecision"`
			} `json:"quantity"`
			SekQuantity struct {
				Value            float64 `json:"value"`
				Unit             string  `json:"unit"`
				UnitType         string  `json:"unitType"`
				DecimalPrecision int     `json:"decimalPrecision"`
			} `json:"sekQuantity"`
		} `json:"preliminaryDividends"`
		ConfirmedDividends []struct {
			Position struct {
				OrderbookID    string `json:"orderbookId"`
				Name           string `json:"name"`
				FlagCode       string `json:"flagCode"`
				InstrumentName string `json:"instrumentName"`
			} `json:"position"`
			PaymentDate  string `json:"paymentDate"`
			ExDate       string `json:"exDate"`
			DividendType string `json:"dividendType"`
			Quantity     struct {
				Value            float64 `json:"value"`
				Unit             string  `json:"unit"`
				UnitType         string  `json:"unitType"`
				DecimalPrecision int     `json:"decimalPrecision"`
			} `json:"quantity"`
			SekQuantity struct {
				Value            float64 `json:"value"`
				Unit             string  `json:"unit"`
				UnitType         string  `json:"unitType"`
				DecimalPrecision int     `json:"decimalPrecision"`
			} `json:"sekQuantity"`
		} `json:"confirmedDividends"`
	} `json:"dividends"`
	Reports []struct {
		Date     string `json:"date"`
		Type     string `json:"type"`
		Position struct {
			OrderbookID    string `json:"orderbookId"`
			Name           string `json:"name"`
			FlagCode       string `json:"flagCode"`
			InstrumentName string `json:"instrumentName"`
		} `json:"position"`
	} `json:"reports"`
}

type GetTransactionsResponse struct {
	Transactions              []Transaction `json:"transactions"`
	TotalNumberOfTransactions int           `json:"totalNumberOfTransactions"`
}

type Transaction struct {
	Account struct {
		Type string `json:"type"`
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"account"`
	TransactionType  string  `json:"transactionType"`
	VerificationDate string  `json:"verificationDate"`
	Description      string  `json:"description"`
	Currency         string  `json:"currency"`
	Amount           float64 `json:"amount"`
	ID               string  `json:"id"`
	Sum              float64 `json:"sum,omitempty"`
	Commission       float64 `json:"commission,omitempty"`
	NoteID           string  `json:"noteId,omitempty"`
	CurrencyRate     float64 `json:"currencyRate,omitempty"`
	Orderbook        struct {
		Isin     string `json:"isin"`
		Currency string `json:"currency"`
		Name     string `json:"name"`
		FlagCode string `json:"flagCode"`
		ID       string `json:"id"`
		Type     string `json:"type"`
	} `json:"orderbook,omitempty"`
	Volume float64 `json:"volume,omitempty"`
	Price  float64 `json:"price,omitempty"`
}

type GetTransactionsParams struct {
	AccountID    string    `path:"accountId" validate:"required"`
	FromDate     *Date     `query:"from"`
	ToDate       *Date     `query:"to"`
	OrderBookIds *[]string `query:"orderbookId"`
	MinAmount    *float64  `query:"minAmount"`
	MaxAmount    *float64  `query:"maxAmount"`
}
type AccountOverviewParams struct {
	AccountID string `path:"accountId" validate:"required"`
}

type AccountOverviewResponse struct {
	CourtageClass                      string  `json:"courtageClass"`
	Depositable                        bool    `json:"depositable"`
	AccountType                        string  `json:"accountType"`
	Withdrawable                       bool    `json:"withdrawable"`
	ClearingNumber                     string  `json:"clearingNumber"`
	InstrumentTransferPossible         bool    `json:"instrumentTransferPossible"`
	InternalTransferPossible           bool    `json:"internalTransferPossible"`
	JointlyOwned                       bool    `json:"jointlyOwned"`
	AccountId                          string  `json:"accountId"`
	AccountTypeName                    string  `json:"accountTypeName"`
	InterestRate                       float64 `json:"interestRate"`
	NumberOfOrders                     int     `json:"numberOfOrders"`
	NumberOfDeals                      int     `json:"numberOfDeals"`
	PerformanceSinceOneWeek            float64 `json:"performanceSinceOneWeek"`
	PerformanceSinceOneMonth           float64 `json:"performanceSinceOneMonth"`
	PerformanceSinceThreeMonths        float64 `json:"performanceSinceThreeMonths"`
	PerformanceSinceSixMonths          float64 `json:"performanceSinceSixMonths"`
	PerformanceSinceOneYear            float64 `json:"performanceSinceOneYear"`
	PerformanceSinceThreeYears         float64 `json:"performanceSinceThreeYears"`
	PerformanceSinceOneWeekPercent     float64 `json:"performanceSinceOneWeekPercent"`
	PerformanceSinceOneMonthPercent    float64 `json:"performanceSinceOneMonthPercent"`
	PerformanceSinceThreeMonthsPercent float64 `json:"performanceSinceThreeMonthsPercent"`
	PerformanceSinceSixMonthsPercent   float64 `json:"performanceSinceSixMonthsPercent"`
	PerformanceSinceOneYearPercent     float64 `json:"performanceSinceOneYearPercent"`
	PerformanceSinceThreeYearsPercent  float64 `json:"performanceSinceThreeYearsPercent"`
	AvailableSuperLoanAmount           float64 `json:"availableSuperLoanAmount"`
	AllowMonthlySaving                 bool    `json:"allowMonthlySaving"`
	TotalProfit                        float64 `json:"totalProfit"`
	CurrencyAccounts                   []struct {
		Currency string  `json:"currency"`
		Balance  float64 `json:"balance"`
	} `json:"currencyAccounts"`
	CreditLimit               float64 `json:"creditLimit"`
	ForwardBalance            float64 `json:"forwardBalance"`
	ReservedAmount            float64 `json:"reservedAmount"`
	TotalCollateralValue      float64 `json:"totalCollateralValue"`
	TotalPositionsValue       float64 `json:"totalPositionsValue"`
	BuyingPower               float64 `json:"buyingPower"`
	TotalProfitPercent        float64 `json:"totalProfitPercent"`
	Overdrawn                 bool    `json:"overdrawn"`
	Performance               float64 `json:"performance"`
	AccruedInterest           float64 `json:"accruedInterest"`
	CreditAfterInterest       float64 `json:"creditAfterInterest"`
	PerformancePercent        float64 `json:"performancePercent"`
	OverMortgaged             bool    `json:"overMortgaged"`
	TotalBalance              float64 `json:"totalBalance"`
	OwnCapital                float64 `json:"ownCapital"`
	NumberOfTransfers         int     `json:"numberOfTransfers"`
	NumberOfIntradayTransfers int     `json:"numberOfIntradayTransfers"`
	StandardDeviation         float64 `json:"standardDeviation"`
	SharpeRatio               float64 `json:"sharpeRatio"`
}

type OverviewParams struct{}

type OverviewResponse struct {
	Accounts []struct {
		AccountType        string  `json:"accountType"`
		InterestRate       float64 `json:"interestRate"`
		Depositable        bool    `json:"depositable"`
		Attorney           bool    `json:"attorney"`
		Active             bool    `json:"active"`
		AccountId          string  `json:"accountId"`
		AccountPartlyOwned bool    `json:"accountPartlyOwned"`
		Tradable           bool    `json:"tradable"`
		TotalBalance       float64 `json:"totalBalance"`
		TotalBalanceDue    float64 `json:"totalBalanceDue"`
		OwnCapital         float64 `json:"ownCapital"`
		BuyingPower        float64 `json:"buyingPower"`
		TotalProfitPercent float64 `json:"totalProfitPercent"`
		Performance        float64 `json:"performance"`
		TotalProfit        float64 `json:"totalProfit"`
		PerformancePercent float64 `json:"performancePercent"`
		Name               string  `json:"name"`
		SparkontoPlusType  string  `json:"sparkontoPlusType,omitempty"`
	} `json:"accounts"`
	NumberOfOrders            int     `json:"numberOfOrders"`
	NumberOfDeals             int     `json:"numberOfDeals"`
	TotalBuyingPower          float64 `json:"totalBuyingPower"`
	TotalOwnCapital           float64 `json:"totalOwnCapital"`
	TotalPerformancePercent   float64 `json:"totalPerformancePercent"`
	TotalPerformance          float64 `json:"totalPerformance"`
	TotalBalance              float64 `json:"totalBalance"`
	NumberOfTransfers         int     `json:"numberOfTransfers"`
	NumberOfIntradayTransfers int     `json:"numberOfIntradayTransfers"`
}

type GetPositionsParams struct{}

type Value struct {
	Value            float64 `json:"value"`
	Unit             string  `json:"unit"`
	UnitType         string  `json:"unitType"`
	DecimalPrecision int     `json:"decimalPrecision"`
}

type Account struct {
	Id             string `json:"id"`
	Type           string `json:"type"`
	Name           string `json:"name"`
	UrlParameterId string `json:"urlParameterId"`
	HasCredit      bool   `json:"hasCredit"`
}

type GetPositionsResponse struct {
	WithOrderBook []struct {
		Account    Account `json:"account"`
		Instrument struct {
			Id        string `json:"id"`
			Type      string `json:"type"`
			Name      string `json:"name"`
			OrderBook struct {
				Id          string
				FlagCode    string
				Name        string
				Type        string
				TradeStatus string
				Quote       struct {
					Highest       Value  `json:"highest"`
					Lowest        Value  `json:"lowest"`
					Buy           Value  `json:"buy"`
					Sell          Value  `json:"sell"`
					Latest        Value  `json:"latest"`
					Change        Value  `json:"change"`
					ChangePercent Value  `json:"changePercent"`
					Updated       string `json:"updated"`
				} `json:"quote"`
				Turnover struct {
					Volume Value `json:"volume"`
					Value  Value `json:"value"`
				} `json:"turnover"`
				LastDeal struct {
					Date string
					Time string
				}
			} `json:"orderBook"`
			Currency     string  `json:"currency"`
			Isin         string  `json:"isin"`
			VolumeFactor float64 `json:"volumeFactor"`
		} `json:"instrument"`
		LastTradingDayPerformance struct {
			Absolute Value `json:"Absolute"`
			Relative Value `json:"relative"`
		} `json:"lastTradingDayPerformance"`
		Id                                     string `json:"id"`
		SuperInterestApproved                  bool   `json:"superInterestApproved"`
		Volume                                 Value  `json:"volume"`
		Value                                  Value  `json:"value"`
		AverageAcquiredPrice                   Value  `json:"averageAcquiredPrice"`
		AverageAcquiredPriceInstrumentCurrency Value  `json:"averageAcquiredPriceInstrumentCurrency"`
		AcquiredValue                          Value  `json:"acquiredValue"`
		CollateralFactor                       Value  `json:"collateralFactor"`
	} `json:"withOrderbook"`
	WithoutOrderBook []struct{} `json:"withoutOrderbook,omitempty"`
	CashPositions    []struct {
		Account      Account `json:"account"`
		TotalBalance Value   `json:"totalBalance"`
	} `json:"cashPositions"`
	WithCreditAccount bool `json:"withCreditAccount"`
}

type GetDealsAndOrdersParams struct{}

type GetDealsAndOrdersResponse struct {
	Orders []interface{} `json:"orders"`
	Deals  []struct {
		Account struct {
			Type string `json:"type"`
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"account"`
		DealID            string `json:"deadId"`
		DealTime          string `json:"dealTime"`
		MarketTransaction bool   `json:"marketTransaction"`
		Orderbook         struct {
			// TODO
		} `json:"orderbook"`
		OrderId string  `json:"orderId"`
		Price   float64 `json:"price"`
		Sum     float64 `json:"sum"`
		Type    string  `json:"type"`
		Volume  int     `json:"Volume"`
	} `json:"deals"`
	Accounts []struct {
		Type string `json:"type"`
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"accounts"`
	ReservedAmount float64 `json:"reservedAmount"`
}
