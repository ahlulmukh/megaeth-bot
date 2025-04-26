package bepop

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Service struct {
	contract *Bepop
	client   bind.ContractBackend
	chainID  *big.Int
}

func (s *Service) Client() bind.ContractBackend {
	return s.client
}

func NewService(client bind.ContractBackend, contractAddress common.Address, chainID *big.Int) (*Service, error) {
	contract, err := NewBepop(contractAddress, client)
	if err != nil {
		return nil, err
	}
	return &Service{
		contract: contract,
		client:   client,
		chainID:  chainID,
	}, nil
}

func (s *Service) DepositETH(privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, s.chainID)
	if err != nil {
		return nil, err
	}
	auth.Value = amount

	return s.contract.Deposit(auth)
}

func (s *Service) Withdraw(privateKey *ecdsa.PrivateKey, amount *big.Int) (*types.Transaction, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, s.chainID)
	if err != nil {
		return nil, err
	}

	return s.contract.Withdraw(auth, amount)
}

func (s *Service) BalanceOf(ctx context.Context, address common.Address) (*big.Int, error) {
	return s.contract.BalanceOf(&bind.CallOpts{Context: ctx}, address)
}
