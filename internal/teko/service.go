package teko

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Service struct {
	contract *Teko
	client   bind.ContractBackend
	chainID  *big.Int
}

func NewService(client bind.ContractBackend, contractAddress common.Address, chainID *big.Int) (*Service, error) {
	contract, err := NewTeko(contractAddress, client)
	if err != nil {
		return nil, err
	}
	return &Service{
		contract: contract,
		client:   client,
		chainID:  chainID,
	}, nil
}

func (s *Service) Mint(privateKey *ecdsa.PrivateKey, to common.Address, amount *big.Int) (*types.Transaction, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, s.chainID)
	if err != nil {
		return nil, err
	}

	return s.contract.Mint(auth, to, amount)
}
