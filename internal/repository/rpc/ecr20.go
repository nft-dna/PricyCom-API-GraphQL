// Package rpc provides high level access to the Volcano Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/repository/rpc/contracts"
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Erc721StartingBlockNumber provides the first important block number for the ERC-721 contract.
// We try to get the first Transfer() event on the contract,
// anything before it is irrelevant for this API.
func (o *Opera) Erc20StartingBlockNumber(adr *common.Address) (uint64, error) {
	// instantiate contract
	erc, err := contracts.NewErc20(*adr, o.ftm)
	if err != nil {
		return 0, err
	}

	// iterate over transfers from zero address (e.g. mint calls)
	iter, err := erc.FilterTransfer(nil, []common.Address{{}}, nil)
	if err != nil {

		// MM timeout on Sepolia 'open' rpc
		if strings.Contains(strings.ToLower(fmt.Sprint(err)), "timed out") {
			// do something
			chb, perr := o.CurrentHead()
			if perr == nil {
				step := uint64(20000)
				resOk := false
				b := o.minBlockNumber
				for b <= chb {
					stop := b + step
					filterOps := bind.FilterOpts{
						Context: context.Background(),
						Start:   b,
						End:     &stop,
					}
					iter, err = erc.FilterTransfer(&filterOps, []common.Address{{}}, nil)
					if err == nil && iter.Event != nil {
						resOk = true
						break
					}
					b = b + step
				}
				/*
					if !resOk {
						b = o.minBlockNumber
						for b > 0 {

							filterOps := bind.FilterOpts{
								Context: context.Background(),
								Start:   b - step,
								End:     &b,
							}
							iter, err = erc.FilterTransfer(&filterOps, []common.Address{{}}, nil)
							if err == nil && iter.Event != nil {
								resOk = true
								break
							}
							if b > step {
								b = b - step
							} else if b > 0 {
								b = 0
							} else {
								break
							}
						}
					}
				*/
				if !resOk {
					return 0, err
				}
			}
		}
	}

	var blk uint64
	if iter.Next() {
		blk = iter.Event.Raw.BlockNumber
	}

	if err := iter.Close(); err != nil {
		log.Errorf("could not close filter iterator; %s", err.Error())
	}
	return blk, nil
}

func (o *Opera) Erc20InitialReserves(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc20Abi().Pack("initialReserves")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}

func (o *Opera) Erc20MintBlocksAmount(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc20Abi().Pack("mintBlocksAmount")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}

func (o *Opera) Erc20MintBlocksFee(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc20Abi().Pack("mintBlocksFee")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}

func (o *Opera) Erc20MintBlocksMaxSupply(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc20Abi().Pack("mintBlocksMaxSupply")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return nil, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(data), nil
}
