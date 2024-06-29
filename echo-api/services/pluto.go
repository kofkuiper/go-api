package services

import "math/big"

type (
	PlutoService interface {
		ChainInfo() (*ChainInfo, error)
		BalanceOf(string) (*big.Float, error)
	}

	ChainInfo struct {
		ChainID     uint64
		BlockNumber uint64
	}

	EthBalance struct {
		WalletAddress string `json:"walletAddress" validate:"required,len=42,eth_addr"`
	}
)
