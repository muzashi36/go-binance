package binance

import (
	"context"
	"net/http"
)

type AssetTransferService struct {
	c            *Client
	transferType *string
	startTime    *int64
	endTime      *int64
	current      *int
	size         *int
	fromSymbol   *string
	toSymbol     *string
	timestamp    *int64
}

func (s *AssetTransferService) TransferType(transferType string) *AssetTransferService {
	s.transferType = &transferType
	return s
}

func (s *AssetTransferService) StartTime(startTime int64) *AssetTransferService {
	s.startTime = &startTime
	return s
}

func (s *AssetTransferService) EndTime(endTime int64) *AssetTransferService {
	s.endTime = &endTime
	return s
}

func (s *AssetTransferService) Current(current int) *AssetTransferService {
	s.current = &current
	return s
}

func (s *AssetTransferService) Size(size int) *AssetTransferService {
	s.size = &size
	return s
}

func (s *AssetTransferService) FromSymbol(fromSymbol string) *AssetTransferService {
	s.fromSymbol = &fromSymbol
	return s
}

func (s *AssetTransferService) ToSymbol(toSymbol string) *AssetTransferService {
	s.toSymbol = &toSymbol
	return s
}

func (s *AssetTransferService) Timestamp(timestamp int64) *AssetTransferService {
	s.timestamp = &timestamp
	return s
}

func (s *AssetTransferService) Do(ctx context.Context) (*AssetTransfer, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/transfer",
		secType:  secTypeSigned,
	}

	if s.transferType != nil {
		r.setParam("type", *s.transferType)
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

	if s.fromSymbol != nil {
		r.setParam("fromSymbol", *s.fromSymbol)
	}

	if s.toSymbol != nil {
		r.setParam("toSymbol", *s.toSymbol)
	}

	if s.timestamp != nil {
		r.setParam("timestamp", *s.timestamp)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := AssetTransfer{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

type AssetTransfer struct {
	Total int64                 `json:"total"`
	Rows  []AssetTransferDetail `json:"rows"`
}

type AssetTransferDetail struct {
	Asset     string `json:"asset"`
	Amount    string `json:"amount"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	TranId    int64  `json:"tranId"`
	Timestamp int64  `json:"timestamp"`
}
