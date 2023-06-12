package binance

import (
	"context"
	"net/http"
)

// SavingsRedemptionHistoryService https://binance-docs.github.io/apidocs/spot/en/#get-redemption-record-user_data
type SavingsRedemptionHistoryService struct {
	c           *Client
	lendingType LendingType
	asset       *string
	startTime   *int64
	endTime     *int64
	current     *int32
	size        *int32
}

// LendingType sets the lendingType parameter.[ ("DAILY" for flexible, "ACTIVITY" for activity, "CUSTOMIZED_FIXED" for fixed)]
func (s *SavingsRedemptionHistoryService) LendingType(lendingType LendingType) *SavingsRedemptionHistoryService {
	s.lendingType = lendingType
	return s
}

// Asset sets the asset parameter.
func (s *SavingsRedemptionHistoryService) Asset(asset string) *SavingsRedemptionHistoryService {
	s.asset = &asset
	return s
}

// StartTime sets the startTime parameter.
func (s *SavingsRedemptionHistoryService) StartTime(startTime int64) *SavingsRedemptionHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
func (s *SavingsRedemptionHistoryService) EndTime(endTime int64) *SavingsRedemptionHistoryService {
	s.endTime = &endTime
	return s
}

// Current sets the current parameter.
func (s *SavingsRedemptionHistoryService) Current(current int32) *SavingsRedemptionHistoryService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *SavingsRedemptionHistoryService) Size(size int32) *SavingsRedemptionHistoryService {
	s.size = &size
	return s
}

// Do sends the request.
func (s *SavingsRedemptionHistoryService) Do(ctx context.Context) (*LendingRedemptionHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/lending/union/redemptionRecord",
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
	res := new(LendingRedemptionHistory)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// LendingRedemptionHistory represents a list of lending Redemption history transactions.
type LendingRedemptionHistory []LendingRedemptionHistoryTransaction

// LendingRedemptionHistoryTransaction represents a lending Redemption history transaction.
type LendingRedemptionHistoryTransaction struct {
	Amount      string  `json:"amount"`
	Asset       string  `json:"asset"`
	CreateTime  int64   `json:"createTime"`
	Principle   string  `json:"principle"`
	Interest    *string `json:"interest"`
	ProjectId   string  `json:"projectId"`
	ProjectName string  `json:"projectName"`
	StartTime   *int    `json:"startTime"`
	Status      string  `json:"status"`
	Type        *string `json:"type"`
}
