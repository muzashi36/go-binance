package binance

import (
	"context"
	"net/http"
)

// SavingsInterestHistoryService https://binance-docs.github.io/apidocs/spot/en/#get-interest-history-user_data-2
type SavingsInterestHistoryService struct {
	c           *Client
	lendingType LendingType
	asset       *string
	startTime   *int64
	endTime     *int64
	current     *int32
	size        *int32
}

// LendingType sets the lendingType parameter.[ ("DAILY" for flexible, "ACTIVITY" for activity, "CUSTOMIZED_FIXED" for fixed)]
func (s *SavingsInterestHistoryService) LendingType(lendingType LendingType) *SavingsInterestHistoryService {
	s.lendingType = lendingType
	return s
}

// Asset sets the asset parameter.
func (s *SavingsInterestHistoryService) Asset(asset string) *SavingsInterestHistoryService {
	s.asset = &asset
	return s
}

// StartTime sets the startTime parameter.
func (s *SavingsInterestHistoryService) StartTime(startTime int64) *SavingsInterestHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
func (s *SavingsInterestHistoryService) EndTime(endTime int64) *SavingsInterestHistoryService {
	s.endTime = &endTime
	return s
}

// Current sets the current parameter.
func (s *SavingsInterestHistoryService) Current(current int32) *SavingsInterestHistoryService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *SavingsInterestHistoryService) Size(size int32) *SavingsInterestHistoryService {
	s.size = &size
	return s
}

// Do sends the request.
func (s *SavingsInterestHistoryService) Do(ctx context.Context) (*LendingInterestHistory, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/lending/union/interestHistory",
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
	res := new(LendingInterestHistory)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// LendingInterestHistory represents a list of lending Interest history transactions.
type LendingInterestHistory []LendingInterestHistoryTransaction

// LendingInterestHistoryTransaction represents a lending Interest history transaction.
type LendingInterestHistoryTransaction struct {
	Asset       string `json:"asset"`
	Interest    string `json:"interest"`
	LendingType string `json:"lendingType"`
	ProductName string `json:"productName"`
	Time        int64  `json:"time"`
}
