// Package rpc provides high level access to the Volcano Opera blockchain
// node through RPC interface.
package rpc

import (
	"artion-api-graphql/internal/config"
	"artion-api-graphql/internal/logger"
	"artion-api-graphql/internal/repository/rpc/contracts"
	"artion-api-graphql/internal/types"
	"bytes"
	"crypto/tls"
	"embed"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	client "github.com/ethereum/go-ethereum/rpc"
)

//go:embed contracts/abi/*.json
var abiFiles embed.FS

const (
	// headerObserverCapacity represents the capacity of new headers' observer channel
	headerObserverCapacity = 5000
)

// config represents the configuration setup used by the repository
// to establish and maintain required connectivity to external services
// as needed.
var cfg *config.Config

// log represents the logger to be used by the repository.
var log logger.Logger

// Opera represents the implementation of the Blockchain interface for Volcano Opera node.
type Opera struct {
	// basics of the connection
	rpc *client.Client
	ftm *ethclient.Client

	// sync tools
	wg       *sync.WaitGroup
	sigClose chan bool

	// captured header queue
	headers chan *eth.Header

	// decode ABI structures
	abiVolcano721         *abi.ABI
	abiVolcano721Factory  *abi.ABI
	abiVolcano1155        *abi.ABI
	abiVolcano1155Factory *abi.ABI
	abiVolcano20          *abi.ABI
	abiVolcano20Factory   *abi.ABI
	abiMarketplace        *abi.ABI

	// contracts
	auctionContractsProps   map[common.Address]types.AuctionProps
	auctionContracts        map[common.Address]IAuctionContract
	marketplaceContracts    map[common.Address]IMarketplaceContract
	payTokenPriceContract   IMarketplaceContract // for token royalty or pay token price
	tokenRegistryContract   *contracts.VolcanoTokenRegistry
	royaltyRegistryContract *contracts.FantomRoyaltyRegistry
	rngFeedContract         *contracts.RandomNumberOracle

	basicContracts types.Contracts

	minBlockNumber uint64
}

// RegisterContract adds a new contract address to the RPC provider.
func (o *Opera) RegisterContract(ct string, addr *common.Address, blockNum uint64) (err error) {
	// address provided?
	if nil == addr {
		return fmt.Errorf("empty address on %s", ct)
	}

	if (o.minBlockNumber == 0 || o.minBlockNumber > blockNum) && blockNum > 0 {
		o.minBlockNumber = blockNum
	}

	// load the contract instance
	switch ct {

	case "auction":
		var ac AuctionContractV2
		ac.auctionV2Contract, err = contracts.NewVolcanoAuction(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded auction contract at %s", addr.String())
		}
		o.auctionContracts[*addr] = &ac
		o.auctionContractsProps[*addr] = types.AuctionV3Props
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}
		o.basicContracts.AuctionHall = *addr
		log.Warningf("setting basic auction to %s", addr.String())

	case "market":
		var mc MarketplaceContractV1
		mc.marketplace, err = contracts.NewVolcanoMarketplace(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}
		o.marketplaceContracts[*addr] = &mc
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}
		o.basicContracts.Marketplace = *addr
		o.payTokenPriceContract = &mc
		log.Warningf("setting basic %s to %s", ct, addr.String())

	case "rng":
		o.rngFeedContract, err = contracts.NewRandomNumberOracle(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}

	case "token_registry":
		o.tokenRegistryContract, err = contracts.NewVolcanoTokenRegistry(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}

	case "royalty_registry":
		o.royaltyRegistryContract, err = contracts.NewFantomRoyaltyRegistry(*addr, o.ftm)
		if err == nil {
			log.Noticef("loaded %s contract at %s", ct, addr.String())
		}

	default:
		err = fmt.Errorf("unknown contract type %s", ct)
	}

	return err
}

// New provides a new instance of the RPC access point.
func New() *Opera {
	con, err := connect()
	if err != nil {
		log.Criticalf("can not connect Opera node; %s", err.Error())
		return nil
	}

	op := &Opera{
		rpc: con,
		ftm: ethclient.NewClient(con),

		wg:       new(sync.WaitGroup),
		sigClose: make(chan bool, 1),
		headers:  make(chan *eth.Header, headerObserverCapacity),

		auctionContractsProps: make(map[common.Address]types.AuctionProps),
		auctionContracts:      make(map[common.Address]IAuctionContract),
		marketplaceContracts:  make(map[common.Address]IMarketplaceContract),
	}

	// load and parse ABIs
	if err := loadABI(op); err != nil {
		log.Criticalf("can not parse ABI files; %s", err.Error())
		return nil
	}

	// start the observer
	op.wg.Add(1)
	go op.observe()

	log.Notice("blockchain connection is open")
	return op
}

// connects opens RPC connection to the Opera node.
func connect() (*client.Client, error) {
	var c *client.Client
	var err error
	if len(cfg.Node.Proxy) > 0 {
		proxyUrl, _ := url.Parse(cfg.Node.Proxy)
		if strings.HasPrefix(cfg.Node.Url, "ws") {

			log.Criticalf("upgrade ethereum rpc to handle proxied wss connections")
			return nil, fmt.Errorf("upgrade ethereum rpc to handle proxied wss connections")
			/*
				rpcwssClient := client.WithWebsocketDialer(websocket.Dialer{
					HandshakeTimeout: 45 * time.Second,
					TLSClientConfig:  &tls.Config{InsecureSkipVerify: cfg.Node.InsecureSkipVerify}, // Create a custom HTTP client with InsecureSkipVerify to ignore SSL certificate errors if needed
					NetDial: (&net.Dialer{
						Timeout:   45 * time.Second,
						KeepAlive: 30 * time.Second,
					}).Dial,
				})

				c, err = client.DialOptions(context.Background(), cfg.Node.Url, rpcwssClient)
			*/
		} else {
			httpClient := &http.Client{
				Transport: &http.Transport{
					//MaxConnsPerHost : 120,
					DialContext: (&net.Dialer{
						Timeout:   30 * time.Second,
						KeepAlive: 30 * time.Second,
					}).DialContext,
					MaxIdleConnsPerHost:   25,
					MaxIdleConns:          200,
					IdleConnTimeout:       2 * time.Second,
					TLSHandshakeTimeout:   10 * time.Second,
					ExpectContinueTimeout: 1 * time.Second,
					TLSClientConfig:       &tls.Config{InsecureSkipVerify: cfg.Node.InsecureSkipVerify}, // Create a custom HTTP client with InsecureSkipVerify to ignore SSL certificate errors if needed
					Proxy:                 http.ProxyURL(proxyUrl),
				},
			}
			c, err = client.DialHTTPWithClient(cfg.Node.Url, httpClient)
		}
	} else {
		c, err = client.Dial(cfg.Node.Url)
	}
	if err != nil {
		log.Criticalf("can not connect blockchain node; %s", err.Error())
		return nil, err
	}

	log.Noticef("blockchain node connected at %s", cfg.Node.Url)
	return c, nil
}

// loadABI tries to load and parse expected ABI for contracts we need.
func loadABI(o *Opera) (err error) {
	o.abiVolcano721, err = loadABIFile("contracts/abi/VolcanoERC721Tradable.json")
	if err != nil {
		return err
	}

	o.abiVolcano721Factory, err = loadABIFile("contracts/abi/VolcanoERC721Factory.json")
	if err != nil {
		return err
	}

	o.abiVolcano1155, err = loadABIFile("contracts/abi/VolcanoERC1155Tradable.json")
	if err != nil {
		return err
	}

	o.abiVolcano1155Factory, err = loadABIFile("contracts/abi/VolcanoERC1155Factory.json")
	if err != nil {
		return err
	}

	o.abiVolcano20, err = loadABIFile("contracts/abi/VolcanoERC20Token.json")
	if err != nil {
		return err
	}

	o.abiVolcano20Factory, err = loadABIFile("contracts/abi/VolcanoERC20Factory.json")
	if err != nil {
		return err
	}

	o.abiMarketplace, err = loadABIFile("contracts/abi/VolcanoMarketplace.json")
	if err != nil {
		return err
	}

	return nil
}

// loadABIFile reads specified ABI file and returns the decoded ABI.
func loadABIFile(path string) (*abi.ABI, error) {
	// VolcanoERC721Tradable
	data, err := abiFiles.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// parse ABI
	decoded, err := abi.JSON(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return &decoded, nil
}

// Close terminates the node connection.
func (o *Opera) Close() {
	// stop block observer
	o.sigClose <- true
	o.wg.Wait()

	o.ftm.Close()
	log.Noticef("blockchain connection closed")
}

// NewHeaders provides a channel receiving new blockchain headers.
func (o *Opera) NewHeaders() chan *eth.Header {
	return o.headers
}

// BasicContracts provides addresses of basic Artion contracts.
func (o *Opera) BasicContracts() *types.Contracts {
	return &o.basicContracts
}

// IsEscrowContract test if the address is address of auction or other escrow contract.
func (o *Opera) IsEscrowContract(addr common.Address) bool {
	_, exists := o.auctionContracts[addr]
	return exists
}

// SetConfig sets the repository configuration to be used to establish
// and maintain external repository connections.
func SetConfig(c *config.Config) {
	cfg = c
}

// SetLogger sets the repository logger to be used to collect logging info.
func SetLogger(l logger.Logger) {
	log = l.ModuleLogger("rpc")
}
