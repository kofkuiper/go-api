package services

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"

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

// Wei to Ether (float64)
func FormatEther(wei *big.Int) (*float64, error) {
	bfWei, ok := new(big.Float).SetString(wei.String()) // Big Float Wei
	if !ok {
		return nil, fmt.Errorf("can not convert %v to big.Float", wei)
	}
	bfEth := new(big.Float).Quo(bfWei, big.NewFloat(math.Pow10(18))) // Big Float Ether, divided by Big Float of (10 ^ 18 )
	balance, _ := bfEth.Float64()
	return &balance, nil
}

// Ether to Wei (Big Float)
func ParseEther(value string) (*big.Float, error) {
	bf, ok := new(big.Float).SetString(value)
	if !ok {
		return nil, fmt.Errorf("can not convert %s to big.Float", value)
	}
	wei := new(big.Float).Mul(bf, big.NewFloat(math.Pow10(18)))
	return wei, nil
}

// Int to Big Int
func BigInt(value int64) *big.Int {
	return big.NewInt(value)
}

func PraseAddress(address string) common.Address {
	address = strings.TrimSpace(address)
	address = strings.ToLower(address)
	return common.HexToAddress(address)
}
