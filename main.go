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

// Package main provides interacting with micropayment service
package main

import (
	"flag"
	"os"

	"git.atmatrix.org/infra/go-contract/payment"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

var (
	key, _ = crypto.HexToECDSA("")
)

var (
	logFlag = flag.Int("loglevel", 3, "Log level to use for Ethereum")
)

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.Lvl(*logFlag), log.StreamHandler(os.Stderr, log.TerminalFormat(true))))
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Crit("Failed to create rpc client", "err", err)
	}
	transactOpts := bind.NewKeyedTransactor(key)
	payAddr, tx, pay, err := payment.DeployPayment(transactOpts, client)
	if err != nil {
		log.Error("Failed to deploy payment", "err", err)
	}
	log.Info("PayAddress", "addr", payAddr.Hex())
	log.Info("Tx", "hash", tx.Hash().Hex())
	_ = pay
}
