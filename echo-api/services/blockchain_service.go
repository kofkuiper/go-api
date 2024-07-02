package services

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
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
