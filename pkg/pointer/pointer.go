package pointer

import "github.com/wo4zhuzi/solana-go-sdk/common"

func Uint64(v uint64) *uint64 {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func Pubkey(v common.PublicKey) *common.PublicKey {
	return &v
}
