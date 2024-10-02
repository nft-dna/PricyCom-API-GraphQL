package repository

import (
	"artion-api-graphql/internal/types"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (p *Proxy) GetLegacyCollection(address common.Address) (*types.LegacyCollection, error) {
	key := "lcol_-" + address.String()
	user, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.cache.GetLegacyCollection(address, p.shared.GetLegacyCollection)
	})
	return user.(*types.LegacyCollection), err
}

func (p *Proxy) ListLegacyCollections(collectionFilter types.CollectionFilter, cursor types.Cursor, count int, backward bool) (out *types.LegacyCollectionList, err error) {
	if collectionFilter.IsUsed() || cursor != "" || backward {
		return p.shared.ListLegacyCollections(collectionFilter, cursor, count, backward)
	}

	// cache result for the default (mostly used) params
	list, err, _ := p.callGroup.Do("lcols", func() (interface{}, error) {
		return p.cache.GetLegacyCollectionList(func() (*types.LegacyCollectionList, error) {
			return p.shared.ListLegacyCollections(types.CollectionFilter{}, "", 5000, false)
		})
	})

	colList := list.(*types.LegacyCollectionList)
	if colList != nil && len(colList.Collection) > count {
		newColList := *colList // clone callGroup output before modification
		newColList.Collection = newColList.Collection[:count]
		newColList.HasNext = true
		return &newColList, err
	}
	return colList, err
}

func (p *Proxy) UploadCollectionApplication(app types.CollectionApplication, image types.Image, owner common.Address) (err error) {
	// MM
	var imageCid string
	if p.pinner.EmulateOnSharedDb() {
		err = p.shared.StoreImage(&image)
		if err != nil {
			return fmt.Errorf("saving collection image on share_db failed; %s", err)
		}
		imageCid = image.ID().Hex()
	} else {
		imageCid, err = p.pinner.PinFile("collection-image-"+app.Contract.String(), image.Data)
		if err != nil {
			return fmt.Errorf("uploading collection image failed; %s", err)
		}
	}

	isOwnerOnly := false
	// check if it is created by 'our factory' contract
	// TODO.. implement an 'interface' or check for contract creator address...
	isInternal := true

	if p.IsErc1155Contract(&app.Contract) {
		isOwnerOnly, err = p.CollectionErc1155IsPrivate(&app.Contract)
		if err != nil {
			isInternal = false
		}
		if isInternal {
			_, err = p.CollectionErc1155MaxSupply(&app.Contract)
			if err != nil {
				isInternal = false
			}
		}
		if isInternal {
			_, err = p.CollectionErc1155MintStartTime(&app.Contract)
			if err != nil {
				isInternal = false
			}
		}
		if isInternal {
			_, err = p.CollectionErc1155MintStopTime(&app.Contract)
			if err != nil {
				isInternal = false
			}
		}
		if isInternal {
			_, err = p.CollectionErc1155RevealTime(&app.Contract)
			if err != nil {
				//isInternal = false
			}
		}
		if isInternal {
			_, err = p.CollectionErc1155MaxItemSupply(&app.Contract)
			if err != nil {
				isInternal = false
			}
		}
	} else {
		isOwnerOnly, err = p.CollectionErc721IsPrivate(&app.Contract)
		if err != nil {
			isInternal = false
		}
		if isInternal {
			_, err = p.CollectionErc721MaxSupply(&app.Contract)
			if err != nil {
				isInternal = false
			}
		}
		if isInternal {
			_, err = p.CollectionErc721MintStartTime(&app.Contract)
			if err != nil {
				isInternal = false
			}
		}
		if isInternal {
			_, err = p.CollectionErc721MintStopTime(&app.Contract)
			if err != nil {
				isInternal = false
			}
		}
		if isInternal {
			_, err = p.CollectionErc721RevealTime(&app.Contract)
			if err != nil {
				//isInternal = false
			}
		}
		if isInternal {
			_, err = p.CollectionErc721UseBaseUri(&app.Contract)
			if err != nil {
				//isInternal = false
			}
		}
	}

	collection := app.ToCollection(imageCid, &owner, cfg.Server.AddCollectionAsAppropriate, isInternal, !isInternal || isOwnerOnly)
	return p.shared.InsertLegacyCollection(collection)
}

// MustCollectionName provides a name of an Artion ERC721 and/or ERC1155 token,
// or an empty string, if the name is not available.
func (p *Proxy) MustCollectionName(adr *common.Address) string {
	c, err := p.shared.GetLegacyCollection(*adr)
	if err != nil || c == nil {
		return adr.String()
	}
	if c.Name == "" {
		return adr.String()
	}
	return c.Name
}

func (p *Proxy) ApproveCollection(address common.Address) (err error) {
	err = p.shared.ApproveCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) DeclineCollection(address common.Address) (err error) {
	err = p.shared.DeclineCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) BanCollection(address common.Address) (err error) {
	err = p.shared.BanCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) UnbanCollection(address common.Address) (err error) {
	err = p.shared.UnbanCollection(address)
	p.cache.InvalidateLegacyCollection(address)
	return err
}

func (p *Proxy) ListCollectionsWithAppropriateUpdate(after time.Time, maxAmount int64) (out []*types.LegacyCollection, err error) {
	return p.shared.ListCollectionsWithAppropriateUpdate(after, maxAmount)
}

// IsCollectionAppropriate checks if given collection is approved/not banned collection.
func (p *Proxy) IsCollectionAppropriate(contract *common.Address) bool {
	return p.shared.IsCollectionAppropriate(contract)
}
