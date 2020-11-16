/*
* This program is free software; you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation; either version 2 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program; if not, see <http://www.gnu.org/licenses/>.
*
* Copyright (C) echo, 2020
 */

// Package payment is a an on-chain light client micropayment.
package payment

import (
	"math/big"

	"git.atmatrix.org/infra/go-contract/payment/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//go:generate abigen --sol contract/payment.sol --pkg contract --out contract/payment.go

var (
	MainNetAddress = common.HexToAddress("0x1111111111111111111111111111111111111111")
	TestNetAddress = common.HexToAddress("0x2222222222222222222222222222222222222222")

	RecipientAddress = common.HexToAddress("0x3333333333333333333333333333333333333333")
	Duration         = big.NewInt(3600 * 24 * 30)
)

// Payment is a Go wrapper around an micropayment contract.
type Payment struct {
	*contract.MicroPaymentChannelSession
	contractBackend bind.ContractBackend
}

// NewPayment creates a struct exposing convenient high-level operations for interacting with
// the micropayment Service.
func NewPayment(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*Payment, error) {
	payment, err := contract.NewMicroPaymentChannel(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &Payment{
		&contract.MicroPaymentChannelSession{
			Contract:     payment,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

// DeployPayment deploys an instance of the micropayment.
func DeployPayment(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *types.Transaction, *Payment, error) {
	paymentAddr, tx, _, err := contract.DeployMicroPaymentChannel(transactOpts, contractBackend, RecipientAddress, Duration)
	if err != nil {
		return paymentAddr, nil, nil, err
	}

	pay, err := NewPayment(transactOpts, paymentAddr, contractBackend)
	if err != nil {
		return paymentAddr, nil, nil, err
	}

	return paymentAddr, tx, pay, nil
}
