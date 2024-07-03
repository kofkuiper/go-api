package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

// Transfer implements PlutoService.
func (p plutoService) Transfer(value float64, to string) (*string, error) {
	bValue := ParseEther(value)
	toAddress := PraseAddress(to)

	// convert private key to transactor
	hexKey := "57026c312a291b6308625bf6b729de5a846c4e97be6d7f537a586ff75c5b9fb9"
	chainID, err := p.chainClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	auth, err := PrivateToAccount(hexKey, chainID)
	if err != nil {
		return nil, err
	}
	publicKey, err := PrivateKeyToAddress(hexKey)
	if err != nil {
		return nil, err
	}
	// Setup nonce, gas price
	nonce, err := p.chainClient.PendingNonceAt(context.Background(), *publicKey)
	if err != nil {
		return nil, err
	}
	gasPrice, err := p.chainClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(200000)
	auth.GasPrice = gasPrice

	// Transfer
	instance, err := p.plutoRepo.Instance()
	if err != nil {
		return nil, err
	}
	tx, err := instance.Transfer(auth, toAddress, bValue)
	if err != nil {
		return nil, err
	}
	transactionHash := tx.Hash().Hex()
	return &transactionHash, nil
}

// TransferEth implements PlutoService.
func (p plutoService) TransferEth(value float64, to string) (*string, error) {
	amount := ParseEther(value)
	toAddress := PraseAddress(to)

	hexKey := "57026c312a291b6308625bf6b729de5a846c4e97be6d7f537a586ff75c5b9fb9"
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, err
	}

	publicKey, err := PrivateKeyToAddress(hexKey)
	if err != nil {
		return nil, err
	}

	// Setup nonce, gas price
	nonce, err := p.chainClient.PendingNonceAt(context.Background(), *publicKey)
	if err != nil {
		return nil, err
	}

	gasPrice, err := p.chainClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	chainID, err := p.chainClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	gasLimit := uint64(200000)
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	err = p.chainClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	txHash := signedTx.Hash().Hex()
	return &txHash, nil
}
