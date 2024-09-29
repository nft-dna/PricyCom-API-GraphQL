package repository

import (
	"artion-api-graphql/internal/types"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func (p *Proxy) GetLegacyMemeToken(address common.Address) (*types.LegacyCollection, error) {
	key := "lmtk_-" + address.String()
	user, err, _ := p.callGroup.Do(key, func() (interface{}, error) {
		return p.cache.GetLegacyMemeToken(address, p.shared.GetLegacyCollection)
	})
	return user.(*types.LegacyCollection), err
}

func (p *Proxy) ListLegacyMemeTokens(collectionFilter types.CollectionFilter, cursor types.Cursor, count int, backward bool) (out *types.LegacyCollectionList, err error) {
	if collectionFilter.IsUsed() || cursor != "" || backward {
		return p.shared.ListLegacyMemeTokens(collectionFilter, cursor, count, backward)
	}

	// cache result for the default (mostly used) params
	list, err, _ := p.callGroup.Do("lmtks", func() (interface{}, error) {
		return p.cache.GetLegacyMemeTokenList(func() (*types.LegacyCollectionList, error) {
			return p.shared.ListLegacyMemeTokens(types.CollectionFilter{}, "", 5000, false)
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

func (p *Proxy) UploadMemeTokenApplication(app types.CollectionApplication, image types.Image, owner common.Address) (err error) {
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
	collection := app.ToCollection(imageCid, &owner, cfg.Server.AddCollectionAsAppropriate, false)
	return p.shared.InsertLegacyMemeToken(collection)
}

// MustCollectionName provides a name of an Artion ERC721 and/or ERC1155 token,
// or an empty string, if the name is not available.
func (p *Proxy) MustMemeTokenName(adr *common.Address) string {
	c, err := p.shared.GetLegacyMemeToken(*adr)
	if err != nil || c == nil {
		return adr.String()
	}
	if c.Name == "" {
		return adr.String()
	}
	return c.Name
}

func (p *Proxy) ApproveMemeToken(address common.Address) (err error) {
	err = p.shared.ApproveMemeToken(address)
	p.cache.InvalidateLegacyMemeToken(address)
	return err
}

func (p *Proxy) DeclineMemeToken(address common.Address) (err error) {
	err = p.shared.DeclineMemeToken(address)
	p.cache.InvalidateLegacyMemeToken(address)
	return err
}

func (p *Proxy) BanMemeToken(address common.Address) (err error) {
	err = p.shared.BanMemeToken(address)
	p.cache.InvalidateLegacyMemeToken(address)
	return err
}

func (p *Proxy) UnbanMemeToken(address common.Address) (err error) {
	err = p.shared.UnbanMemeToken(address)
	p.cache.InvalidateLegacyMemeToken(address)
	return err
}

func (p *Proxy) ListMemeTokensWithAppropriateUpdate(after time.Time, maxAmount int64) (out []*types.LegacyCollection, err error) {
	return p.shared.ListMemeTokensWithAppropriateUpdate(after, maxAmount)
}

// IsCollectionAppropriate checks if given collection is approved/not banned collection.
func (p *Proxy) IsMemeTokenAppropriate(contract *common.Address) bool {
	return p.shared.IsMemeTokenAppropriate(contract)
}