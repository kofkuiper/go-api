package services

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kofkuiper/echo-api/repositories"
)

type (
	plutoService struct {
		chainClient ethclient.Client
		plutoRepo   repositories.PlutoRepositoryContract
	}
)

func NewPlutoService(chainClient ethclient.Client, plutoRepo repositories.PlutoRepositoryContract) PlutoService {
	return plutoService{chainClient, plutoRepo}
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

// EthBalanceOf implements PlutoService.
func (p plutoService) EthBalanceOf(walletAddress string) (*float64, error) {
	account := common.HexToAddress(walletAddress)
	wei, err := p.chainClient.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	eth, err := FormatEther(wei)
	if err != nil {
		return nil, err
	}
	return eth, nil
}

func (p plutoService) BalanceOf(walletAddress string) (*float64, error) {
	account := common.HexToAddress(walletAddress)
	instance, err := p.plutoRepo.Instance()
	if err != nil {
		return nil, err
	}

	wei, err := instance.BalanceOf(&bind.CallOpts{}, account)
	if err != nil {
		return nil, err
	}
	eth, err := FormatEther(wei)
	if err != nil {
		return nil, err
	}
	return eth, nil
}
