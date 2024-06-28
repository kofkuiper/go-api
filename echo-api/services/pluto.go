package services

type (
	PlutoService interface {
		ChainInfo() (*ChainInfo, error)
	}

	ChainInfo struct {
		ChainID     uint64
		BlockNumber uint64
	}
)
