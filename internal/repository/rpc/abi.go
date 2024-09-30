// Package rpc provides high level access to the Volcano Opera blockchain
// node through RPC interface.
package rpc

import "github.com/ethereum/go-ethereum/accounts/abi"

// Erc721Abi provides access to decoded ABI of Volcano ERC-721 contract.
func (o *Opera) Erc721Abi() *abi.ABI {
	return o.abiVolcano721
}

// Erc1155Abi provides access to decoded ABI of Volcano ERC-1155 contract.
func (o *Opera) Erc1155Abi() *abi.ABI {
	return o.abiVolcano1155
}

// Erc20Abi provides access to decoded ABI of Volcano ERC-20 contract.
func (o *Opera) Erc20Abi() *abi.ABI {
	return o.abiVolcano20
}
