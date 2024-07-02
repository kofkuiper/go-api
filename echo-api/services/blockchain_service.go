package services

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

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
func ParseEther(value float64) *big.Int {
	// float64 to big float
	bf := new(big.Float)
	bf.SetFloat64(value)

	// set decimals (10^18)
	decimals := new(big.Float)
	decimals.SetInt(big.NewInt(int64(math.Pow10(18))))
	// eth to wei
	bf.Mul(bf, decimals)

	// big float to big int
	result := new(big.Int)
	bf.Int(result)
	return result
}

// Int to Big Int
func BigInt(value int64) *big.Int {
	return big.NewInt(value)
}

// Big Int to Int
func FromBigInt(value big.Int) int64 {
	return value.Int64()
}

// String to Address (Hex / '0x'+ string)
func PraseAddress(address string) common.Address {
	address = strings.TrimSpace(address)
	address = strings.ToLower(address)
	return common.HexToAddress(address)
}

// Private key (HexKey without `0x`) to Transactor
func PrivateToAccount(hexKey string, chainID *big.Int) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	return auth, err
}

// Private key (HexKey without `0x`) to Public key
func PrivateKeyToAddress(hexKey string) (*common.Address, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return &address, nil
}
