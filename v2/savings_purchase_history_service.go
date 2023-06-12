package binance

import (
	"context"
	"net/http"
)

// SavingsPurchaseHistoryService https://binance-docs.github.io/apidocs/spot/en/#get-purchase-record-user_data
type SavingsPurchaseHistoryService struct {
	c           *Client
	lendingType LendingType
	asset       *string
	startTime   *int64
	endTime     *int64
	current     *int32
	size        *int32
}

// LendingType sets the lendingType parameter. ("DAILY" for flexible, "ACTIVITY" for activity, "CUSTOMIZED_FIXED" for fixed)
func (s *SavingsPurchaseHistoryService) LendingType(lendingType LendingType) *SavingsPurchaseHistoryService {
	s.lendingType = lendingType
	return s
}

// Asset sets the asset parameter.
func (s *SavingsPurchaseHistoryService) Asset(asset string) *SavingsPurchaseHistoryService {
	s.asset = &asset
	return s
}

// StartTime sets the startTime parameter.
func (s *SavingsPurchaseHistoryService) StartTime(startTime int64) *SavingsPurchaseHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
func (s *SavingsPurchaseHistoryService) EndTime(endTime int64) *SavingsPurchaseHistoryService {
	s.endTime = &endTime
	return s
}

// Current sets the current parameter.
func (s *SavingsPurchaseHistoryService) Current(current int32) *SavingsPurchaseHistoryService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *SavingsPurchaseHistoryService) Size(size int32) *SavingsPurchaseHistoryService {
	s.size = &size
	return s
}

// Do sends the request.
func (s *SavingsPurchaseHistoryService) Do(ctx context.Context) (*LendingPurchaseHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/lending/union/purchaseRecord",
		secType:  secTypeSigned,
	}
	r.setParam("lendingType", s.lendingType)
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(LendingPurchaseHistory)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// LendingPurchaseHistory represents a list of lending purchase history transactions.
type LendingPurchaseHistory []LendingPurchaseHistoryTransaction

// LendingPurchaseHistoryTransaction represents a lending purchase history transaction.
type LendingPurchaseHistoryTransaction struct {
	Amount      string `json:"amount"`
	Asset       string `json:"asset"`
	CreateTime  int64  `json:"createTime"`
	LendingType string `json:"lendingType"`
	ProductName string `json:"productName"`
	PurchaseId  int64  `json:"purchaseId"`
	Status      string `json:"status"`
}
