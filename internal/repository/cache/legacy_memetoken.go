// Package cache provides in-memory shared cache.
package cache

import (
	"artion-api-graphql/internal/types"
	"encoding/json"
	"strings"

	"github.com/allegro/bigcache"
	"github.com/ethereum/go-ethereum/common"
)

const legacyMemeTokensListKey = "lmtks"

// legacyCollectionKey generates a cache key for a legacy collection.
func legacyMemeTokenKey(contract common.Address) string {
	var sb strings.Builder
	sb.WriteString("lmtk_")
	sb.Write(contract.Bytes())
	return sb.String()
}

// GetLegacyCollection try to get a legacy collection from the cache or backend.
func (c *MemCache) GetLegacyMemeToken(contract common.Address, loader func(address common.Address) (*types.LegacyCollection, error)) (*types.LegacyCollection, error) {
	key := legacyMemeTokenKey(contract)

	data, err := c.cache.Get(key)
	if err == nil {
		lc := new(types.LegacyCollection)
		if err = json.Unmarshal(data, &lc); err != nil {
			return nil, err
		}
		return lc, nil
	}

	// load slow way
	lc, err := loader(contract)
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(lc)
	if err != nil {
		log.Errorf("can not encode legacy collection into cache; %s", err)
		return lc, err
	}
	err = c.cache.Set(key, data)
	if err != nil {
		log.Errorf("could not store legacy collection into cache; %s", err.Error())
	}
	return lc, nil
}

// InvalidateLegacyCollection removes the legacy collection from the in-memory cache.
func (c *MemCache) InvalidateLegacyMemeToken(contract common.Address) {
	log.Infof("Invalidating legacy collection cache for %s", contract.String())

	err := c.cache.Delete(legacyMemeTokenKey(contract))
	if err != nil && err != bigcache.ErrEntryNotFound {
		log.Errorf("could not clear legacy collection cache; %s", err.Error())
	}

	err = c.cache.Delete(legacyMemeTokensListKey)
	if err != nil && err != bigcache.ErrEntryNotFound {
		log.Errorf("could not clear legacy collection list cache; %s", err.Error())
	}
}

// GetLegacyCollectionList try to get a list of legacy collections from cache or backend.
func (c *MemCache) GetLegacyMemeTokenList(loader func() (*types.LegacyCollectionList, error)) (*types.LegacyCollectionList, error) {
	data, err := c.cache.Get(legacyMemeTokensListKey)
	if err == nil {
		lc := new(types.LegacyCollectionList)
		if err = json.Unmarshal(data, &lc); err != nil {
			return nil, err
		}
		return lc, nil
	}

	// load slow way
	lc, err := loader()
	if err != nil {
		return nil, err
	}

	data, err = json.Marshal(lc)
	if err != nil {
		log.Errorf("can not encode legacy collection list into cache; %s", err)
		return lc, err
	}
	err = c.cache.Set(legacyMemeTokensListKey, data)
	if err != nil {
		log.Errorf("could not store legacy collection list into cache; %s", err.Error())
	}
	return lc, nil
}
