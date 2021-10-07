// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"artion-api-graphql/internal/types"
	"time"
)

// nftMetadataWorker represents a service responsible for processing NFT token metadata
// update queue from the metadata updater service.
type nftMetadataWorker struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the router
	sigStop chan bool

	// inTokens represents the channel receiving tokens to be updated.
	inTokens chan *types.Token
}

// newNFTMetadataUpdater creates a new instance of the NFT metadata worker service.
func newNFTMetadataWorker(mgr *Manager) *nftMetadataWorker {
	return &nftMetadataWorker{
		mgr:     mgr,
		sigStop: make(chan bool, 1),
	}
}

// name provides the name of the service.
func (mw *nftMetadataWorker) name() string {
	return "nft metadata worker"
}

// init initializes the block scanner and registers it with the manager.
func (mw *nftMetadataWorker) init() {
	// connect queues
	mw.inTokens = mw.mgr.nftMetaUpdater.outTokens

	// add the updater to the manager
	mw.mgr.add(mw)
}

// close signals the block observer to terminate
func (mw *nftMetadataWorker) close() {
	mw.sigStop <- true
}

// run processes incoming NFT metadata update requests from the
// incoming queue.
func (mw *nftMetadataWorker) run() {
	defer func() {
		mw.mgr.closed(mw)
	}()

	for {
		// pull next token
		select {
		case <-mw.sigStop:
			return
		case tok, ok := <-mw.inTokens:
			if !ok {
				return
			}

			// process the token metadata update
			if err := mw.update(tok); err != nil {
				log.Errorf("NFT update failed; %s", err.Error())
			}
		}
	}
}

// update the given NFT metadata from external metadata source.
func (mw *nftMetadataWorker) update(tok *types.Token) error {
	// get metadata
	if tok.Uri == "" {
		log.Infof("token %d on contract %s URI not available", tok.TokenId.String(), tok.Contract.String())
		return nil
	}

	// get the metadata
	log.Debugf("loading metadata for %s / #%d", tok.Contract.String(), tok.TokenId.ToInt().Uint64())
	md, err := repo.GetTokenJsonMetadata(tok.Uri)
	if err != nil {
		log.Errorf("NFT metadata failed on %s / #%d; %s", tok.Contract.String(), tok.TokenId.ToInt().Uint64(), err.Error())
		return err
	}

	// update the data
	tok.Name = md.Name
	tok.Description = md.Description
	if md.Image != nil {
		tok.ImageURI = *md.Image
	}
	tok.MetadataAge = types.Time(time.Now())

	// update the token in persistent storage
	if err := repo.StoreToken(tok); err != nil {
		log.Errorf("failed metadata update on %s / #%d; %s", tok.Contract.String(), tok.TokenId.String(), err.Error())
		return err
	}

	log.Infof("NFT %s / %s metadata updated [%s]", tok.Contract.String(), tok.TokenId.String(), tok.Name)
	return nil
}
