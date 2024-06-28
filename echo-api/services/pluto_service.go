package services

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

type (
	plutoService struct {
		chainClient ethclient.Client
	}
)

func NewPlutoService(chainClient ethclient.Client) PlutoService {
	return plutoService{chainClient: chainClient}
}

// ChainInfo implements PlutoService.
func (p plutoService) ChainInfo() (*ChainInfo, error) {
	chainID, err := p.chainClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	blockNumber, err := p.chainClient.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	return &ChainInfo{
		ChainID:     uint64(chainID.Int64()),
		BlockNumber: blockNumber,
	}, nil
}
