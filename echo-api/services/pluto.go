package services

type (
	PlutoService interface {
		ChainInfo() (*ChainInfo, error)
		EthBalanceOf(string) (*float64, error)
		BalanceOf(string) (*float64, error)
		Transfer(float64, string) (*string, error)
		TransferEth(float64, string) (*string, error)
	}

	ChainInfo struct {
		ChainID     uint64
		BlockNumber uint64
	}

	EthBalance struct {
		WalletAddress string `json:"walletAddress" validate:"required,len=42,eth_addr"`
	}

	TransferReq struct {
		Value float64 `json:"value" validate:"required,min=0.01"`
		To    string  `json:"to" validate:"required,len=42,eth_addr"`
	}
)
