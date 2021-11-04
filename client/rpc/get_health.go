package rpc

import (
	"context"
)

func (s *RpcClient) GetHealth(ctx context.Context) (string, error) {
	res := struct {
		GeneralResponse
		Result string `json:"result"`
	}{}
	err := s.request(ctx, "getHealth", []interface{}{}, &res)
	if err != nil {
		return "", err
	}

	return res.Result, nil
}

