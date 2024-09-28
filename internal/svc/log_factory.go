// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"time"

	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
)

func newTokenContract(evt *eth.Log, lo *logObserver) {
	// MM todo ERC20 TokenCreated event

	if len(evt.Data) != 64 || len(evt.Topics) != 1 {
		log.Errorf("not Factory::ContractCreated() event #%d/#%d; expected 64 bytes of data, %d given; expected 1 topic, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	// make meme token address
	ca := common.Address{}
	ca.SetBytes(evt.Data[32:64])
	meme := types.Collection{
		Address:  ca,
		IsActive: true,
	}
	log.Infof("found (newTokenContract) MEME contract %s", ca.String())

	if err := extendMemeTokenDetails(&meme, evt, lo); err != nil {
		log.Criticalf("failed to load meme collection %s; %s", meme.Address.String(), err.Error())
		return
	}

	if err := repo.StoreMemeToken(&meme); err != nil {
		log.Criticalf("can not store meme collection %s; %s", meme.Address.String(), err.Error())
		return
	}

	// add observed contract based on the collection
	addObservedContract(&meme, evt)
	log.Infof("new meme collection %s found at %s", meme.Name, meme.Address.String())
}

func extendMemeTokenDetails(meme *types.Collection, evt *eth.Log, _ *logObserver) (err error) {
	// NFT contract type is derived from the factory contract type
	/*nft.Type, err = repo.NFTContractType(&evt.Address)
	if err != nil {
		log.Errorf("contract %s type not known; %s", evt.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s is %s", nft.Address.String(), nft.Type)
	*/
	meme.Type = types.ContractTypeERC20

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("header #%d not available; %s", evt.BlockNumber, err.Error())
		return err
	}
	meme.Created = types.Time(time.Unix(int64(blk.Time), 0))

	return extendMemeTokenMetadata(meme)
}

func extendMemeTokenMetadata(meme *types.Collection) (err error) {
	meme.Name, err = repo.MemeTokenName(&meme.Address)
	if err != nil {
		log.Errorf("%s %s name not known; %s", meme.Type, meme.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s name: %s", meme.Address.String(), meme.Name)

	meme.Symbol, err = repo.MemeTokenSymbol(&meme.Address)
	if err != nil {
		log.Errorf("%s %s symbol not known; %s", meme.Type, meme.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s symbol: %s", meme.Address.String(), meme.Symbol)

	/*
		legacyCollection, err := repo.GetLegacyMemeToken(nft.Address)
		if err != nil {
			log.Errorf("%s %s unable to load off-chain data; %s", nft.Type, nft.Address.String(), err.Error())
			return err
		}

		if nil != legacyCollection {
			nft.Categories, err = legacyCollection.CategoriesAsInt()
			if err != nil {
				log.Errorf("%s %s unable to decode categories; %s", nft.Type, nft.Address.String(), err.Error())
				return err
			}
		}
	*/

	return nil
}

// newNFTContract handles log event for new factory deployed ERC721/ERC1155 contract.
// Factory::event ContractCreated(address creator, address nft)
// MM Factory::event ContractCreated(address creator, address nft, bool isPrivate)
func newNFTContract(evt *eth.Log, lo *logObserver) {
	// sanity check: no additional topics; 2 x Address = 2 x 32 bytes
	//if len(evt.Data) != 64 || len(evt.Topics) != 1 {
	// sanity check: no additional topics; 3 x Address = 3 x 32 bytes
	if len(evt.Data) != 96 || len(evt.Topics) != 1 {
		log.Errorf("not Factory::ContractCreated() event #%d/#%d; expected 96 bytes of data, %d given; expected 1 topic, %d given",
			evt.BlockNumber, evt.Index, len(evt.Data), len(evt.Topics))
		return
	}

	// make NFT address
	ca := common.Address{}
	ca.SetBytes(evt.Data[32:64])
	nft := types.Collection{
		Address:  ca,
		IsActive: true,
	}
	log.Infof("found (newNFTContract) NFT contract %s", ca.String())

	// load NFT details
	if err := extendNFTCollectionDetails(&nft, evt, lo); err != nil {
		log.Criticalf("failed to load NFT collection %s; %s", nft.Address.String(), err.Error())
		return
	}

	// add the collection to persistent storage
	if err := repo.StoreCollection(&nft); err != nil {
		log.Criticalf("can not store NFT collection %s; %s", nft.Address.String(), err.Error())
		return
	}

	// add observed contract based on the collection
	addObservedContract(&nft, evt)
	log.Infof("new NFT collection %s found at %s", nft.Name, nft.Address.String())
}

// extendNFTCollectionDetails collects details of an NFT contract.
func extendNFTCollectionDetails(nft *types.Collection, evt *eth.Log, _ *logObserver) (err error) {
	// NFT contract type is derived from the factory contract type
	nft.Type, err = repo.NFTContractType(&evt.Address)
	if err != nil {
		log.Errorf("contract %s type not known; %s", evt.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s is %s", nft.Address.String(), nft.Type)

	blk, err := repo.GetHeader(evt.BlockNumber)
	if err != nil {
		log.Errorf("header #%d not available; %s", evt.BlockNumber, err.Error())
		return err
	}
	nft.Created = types.Time(time.Unix(int64(blk.Time), 0))

	return extendCollectionMetadata(nft)
}

// extendCollectionMetadata adds metadata to the provided collection structure.
func extendCollectionMetadata(nft *types.Collection) (err error) {
	nft.Name, err = repo.CollectionName(&nft.Address)
	if err != nil {
		log.Errorf("%s %s name not known; %s", nft.Type, nft.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s name: %s", nft.Address.String(), nft.Name)

	nft.Symbol, err = repo.CollectionSymbol(&nft.Address)
	if err != nil {
		log.Errorf("%s %s symbol not known; %s", nft.Type, nft.Address.String(), err.Error())
		return err
	}
	log.Debugf("NFT contract %s symbol: %s", nft.Address.String(), nft.Symbol)

	legacyCollection, err := repo.GetLegacyCollection(nft.Address)
	if err != nil {
		log.Errorf("%s %s unable to load off-chain data; %s", nft.Type, nft.Address.String(), err.Error())
		return err
	}

	if nil != legacyCollection {
		nft.Categories, err = legacyCollection.CategoriesAsInt()
		if err != nil {
			log.Errorf("%s %s unable to decode categories; %s", nft.Type, nft.Address.String(), err.Error())
			return err
		}
	}

	return nil
}

// addObservedContract adds new observed contract into repository and log observer.
func addObservedContract(nft *types.Collection, evt *eth.Log) {
	ca := common.Address{}
	if nil != evt.Data && 32 <= len(evt.Data) {
		ca.SetBytes(evt.Data[:32])
	}

	oc := types.ObservedContract{
		Address:     nft.Address,
		Type:        nft.Type,
		Created:     nft.Created,
		Creator:     ca,
		BlockNumber: evt.BlockNumber,
		DeployedBy:  evt.TxHash,
	}

	// store observed contract into the repository
	repo.AddObservedContract(&oc)
}
