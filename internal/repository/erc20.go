// Package repository implements persistent data access and processing.
package repository

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	// MM to be checked.. are we sure about this ?!
	// Erc20BaseInterfaceID represents an identifier of the base ERC-20 interface in ERC-165 encoding.
	Erc20BaseInterfaceID = "0x36372b07"
)

// IsErc20Contract checks if the given address is a contract with ERC-20 interface support.
func (p *Proxy) IsErc20Contract(adr *common.Address) bool {
	if p.rpc.SupportsInterface(adr, Erc20BaseInterfaceID) {
		return true
	} else {
		v, err := p.rpc.CollectionName(adr)
		if err != nil || len(v) == 0 {
			return false
		}
		v, err = p.rpc.CollectionSymbol(adr)
		if err != nil || len(v) == 0 {
			return false
		}
		return true
	}
}

// Erc20StartingBlockNumber provides the first important block number for the ERC-20 contract.
func (p *Proxy) Erc20StartingBlockNumber(adr *common.Address) (uint64, error) {
	return p.rpc.Erc20StartingBlockNumber(adr)
}
