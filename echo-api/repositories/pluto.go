package repositories

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type (
	PlutoRepositoryContract struct {
		address string
		client  *ethclient.Client
	}
)

func NewPlutoRepoContract(address string, client *ethclient.Client) PlutoRepositoryContract {
	return PlutoRepositoryContract{address, client}
}

func (p PlutoRepositoryContract) Instance() (*Pluto, error) {
	tokenAddress := common.HexToAddress(p.address)
	pluto, err := NewPluto(tokenAddress, p.client)
	if err != nil {
		return nil, err
	}
	return pluto, nil
}
