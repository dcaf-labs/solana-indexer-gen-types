package accountmeta

import (
	"github.com/dcaf-labs/solana-indexer-gen-types/pkg/publickey"
	"github.com/gagliardetto/solana-go"
)

type AccountMeta struct {
	PublicKey  publickey.PublicKey
	IsWritable bool
	IsSigner   bool
}

func (a *AccountMeta) ToSolanaAccountMeta() solana.AccountMeta {
	return solana.AccountMeta{
		PublicKey:  a.PublicKey.PublicKey,
		IsWritable: a.IsWritable,
		IsSigner:   a.IsSigner,
	}
}
