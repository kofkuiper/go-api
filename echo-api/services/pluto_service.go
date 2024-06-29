package services

import (
	"context"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

// BalanceOf implements PlutoService.
func (p plutoService) BalanceOf(walletAddress string) (*big.Float, error) {
	account := common.HexToAddress(walletAddress)
	wei, err := p.chainClient.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	eth := formatEther(wei)
	return eth, nil
}

func formatEther(value *big.Int) *big.Float {
	fBalance := new(big.Float)
	fBalance.SetString(value.String())
	return new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
}
