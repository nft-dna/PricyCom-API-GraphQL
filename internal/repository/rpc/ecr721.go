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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// defaultMintingTestTokenUrl is the URL we used to test NFT minting calls.
const defaultMintingTestTokenUrl = "https://minter.artion.io/default/access/minter/estimation.json"

// defaultMintingTestFee is the default fee we try on minting test (10 FTM).
//var defaultMintingTestFee = hexutil.MustDecodeBig("0x8AC7230489E80000")

func (o *Opera) Erc721IsFromFactory(contract *common.Address, block *big.Int) (*common.Address, error) {
	// prepare params
	input, err := o.Erc721Abi().Pack("factory")
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
	res, err := o.Erc721Abi().Unpack("factory", data)
	if err != nil {
		log.Errorf("can not decode contract %s name; %s", contract.String(), err.Error())
		return nil, err
	}
	return abi.ConvertType(res[0], new(common.Address)).(*common.Address), nil
}

func (o *Opera) Erc721IsPrivate(contract *common.Address, block *big.Int) (bool, error) {
	// prepare params
	input, err := o.Erc721Abi().Pack("isprivate")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return false, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return false, err
	}
	return len(data) == 32 && data[0] == 0 && data[31] > 0, nil
}

func (o *Opera) Erc721UseBaseUri(contract *common.Address, block *big.Int) (bool, error) {
	// prepare params
	input, err := o.Erc721Abi().Pack("useBaseUri")
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return false, err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, block)
	if err != nil {
		return false, err
	}
	return len(data) == 32 && data[0] == 0 && data[31] > 0, nil
}

func (o *Opera) Erc721TotalSupply(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc721Abi().Pack("totalSupply")
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

func (o *Opera) Erc721MaxSupply(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc721Abi().Pack("maxSupply")
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

func (o *Opera) Erc721MintStartTime(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc721Abi().Pack("mintStartTime")
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

func (o *Opera) Erc721MintStopTime(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc721Abi().Pack("mintStopTime")
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

func (o *Opera) Erc721RevealTime(contract *common.Address, block *big.Int) (*big.Int, error) {
	// prepare params
	input, err := o.Erc721Abi().Pack("revealTime")
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

// Erc721StartingBlockNumber provides the first important block number for the ERC-721 contract.
// We try to get the first Transfer() event on the contract,
// anything before it is irrelevant for this API.
func (o *Opera) Erc721StartingBlockNumber(adr *common.Address) (uint64, error) {
	// instantiate contract
	erc, err := contracts.NewErc721(*adr, o.ftm)
	if err != nil {
		return 0, err
	}

	// iterate over transfers from zero address (e.g. mint calls)
	iter, err := erc.FilterTransfer(nil, []common.Address{{}}, nil, nil)
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
					iter, err = erc.FilterTransfer(&filterOps, []common.Address{{}}, nil, nil)
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
							iter, err = erc.FilterTransfer(&filterOps, []common.Address{{}}, nil, nil)
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

// CanMintErc721 checks if the given user can mint a new token on the given NFT contract.
func (o *Opera) CanMintErc721(contract *common.Address, user *common.Address, fee *big.Int) (bool, error) {
	// MM: TODO.. adjust to newer Factory contract
	data, err := o.abiVolcano721.Pack("mint", *user, defaultMintingTestTokenUrl)
	if err != nil {
		return false, err
	}

	// use default fee, if not specified
	if fee == nil {
		fee = o.MustPlatformFee(contract)
		log.Infof("platform fee plus creator fee for %s is %s", contract.String(), (*hexutil.Big)(fee).String())
	}

	// try to estimate the call
	gas, err := o.ftm.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  *user,
		To:    contract,
		Data:  data,
		Value: fee,
	})
	if err != nil {
		log.Warningf("user %s can not mint on ERC-721 %s; %s", user.String(), contract.String(), err.Error())
		return false, nil
	}

	log.Infof("user %s can mint on ERC-721 %s for %d gas", user.String(), contract.String(), gas)
	return true, nil
}

// MustPlatformFee returns the platform fee for the given contract, or the default one.
func (o *Opera) MustPlatformFee(contract *common.Address) *big.Int {
	fact, err := o.Erc721IsFromFactory(contract, nil)
	if err != nil {
		return nil
	}

	input, err := o.abiVolcano721Factory.Pack("platformMintFee")
	if err != nil {
		return nil
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   fact,
		Data: input,
	}, nil)
	if err != nil {
		return nil
	}

	mfee := new(big.Int).SetBytes(data)

	input, err = o.abiVolcano721.Pack("mintCreatorFee")
	if err != nil {
		return nil
	}

	// call the contract
	data, err = o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, nil)
	if err != nil {
		return nil
	}

	mfee.Add(mfee, new(big.Int).SetBytes(data))

	return mfee
}

// Erc721TokenUri gets a token specific URI address from ERC-721 contract using tokenURI() call.
func (o *Opera) Erc721TokenUri(contract *common.Address, tokenId *big.Int) (string, error) {
	// prepare params
	input, err := o.Erc721Abi().Pack("tokenURI", tokenId)
	if err != nil {
		log.Errorf("can not pack data; %s", err.Error())
		return "", err
	}

	// call the contract
	data, err := o.ftm.CallContract(context.Background(), ethereum.CallMsg{
		From: common.Address{},
		To:   contract,
		Data: input,
	}, nil)
	res, err := o.abiVolcano721.Unpack("tokenURI", data)
	if err != nil {
		log.Errorf("can not decode response; %s", err.Error())
		return "", err
	}

	return strings.Replace(
		*abi.ConvertType(res[0], new(string)).(*string),
		"{id}",
		fmt.Sprintf("%064x", tokenId), -1), nil
}

func (o *Opera) Erc721IsApprovedForAll(contract *common.Address, owner *common.Address, operator *common.Address) (bool, error) {
	// instantiate contract
	erc, err := contracts.NewErc721(*contract, o.ftm)
	if err != nil {
		return false, err
	}
	return erc.IsApprovedForAll(nil, *owner, *operator)
}
