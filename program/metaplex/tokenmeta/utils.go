package tokenmeta

import (
	"github.com/wo4zhuzi/solana-go-sdk/common"
)

func GetTokenMetaPubkey(mint common.PublicKey) (common.PublicKey, error) {
	metadataAccount, _, err := common.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			common.MetaplexTokenMetaProgramID.Bytes(),
			mint.Bytes(),
		},
		common.MetaplexTokenMetaProgramID,
	)
	if err != nil {
		return common.PublicKey{}, err
	}
	return metadataAccount, nil
}
