package rpc

import (
	"context"
	"github.com/wo4zhuzi/solana-go-sdk/client"
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

	err = client.CheckRpcResult(res.GeneralResponse, err)

	if err != nil {
		return "", err
	}

	return res.Result, nil
}

