// Package config handles API server configuration binding and loading.
package config

import (
	"crypto/ecdsa"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Config defines configuration options structure for Volcano API server.
type Config struct {
	// AppName holds the name of the application
	AppName string `mapstructure:"app_name"`

	// Server configuration
	Server Server `mapstructure:"server"`

	// Logger configuration
	Log Log `mapstructure:"log"`

	// Node represents the node structure
	Node Node `mapstructure:"node"`

	// IPFS represents the node structure
	Ipfs Ipfs `mapstructure:"ipfs"`

	// Mongo database configuration
	Db Database `mapstructure:"db"`

	// Shared Mongo database configuration
	SharedDb Database `mapstructure:"shared_db"`

	// Cache configuration
	Cache Cache `mapstructure:"cache"`

	// Auth configuration
	Auth Auth `mapstructure:"auth"`

	// RngOracle provides configuration for random feed oracle handling
	RngOracle RandomFeedOracle `mapstructure:"rng"`

	// Notifications provides configuration for notification providers
	Notifications NotificationProviders `mapstructure:"notification"`

	// Contracts provides addresses of mandatory contracts
	Contracts Contracts `mapstructure:"contracts"`
}

// Server represents the GraphQL server configuration
type Server struct {
	BindAddress                string   `mapstructure:"bind"`
	CorsOrigin                 []string `mapstructure:"cors"`
	ReadTimeout                int64    `mapstructure:"read_timeout"`
	WriteTimeout               int64    `mapstructure:"write_timeout"`
	IdleTimeout                int64    `mapstructure:"idle_timeout"`
	HeaderTimeout              int64    `mapstructure:"header_timeout"`
	ResolverTimeout            int64    `mapstructure:"resolver_timeout"`
	MaxParserMemory            int64    `mapstructure:"mem_max"`
	MetadataWorkerThreads      int64    `mapstructure:"metadata_worker_threads"`
	AddCollectionAsAppropriate bool     `mapstructure:"add_collection_as_appropriate"`
}

// Log represents the logger configuration
type Log struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

// Node represents the Volcano Opera node access configuration
type Node struct {
	Proxy                string `mapstructure:"proxy"`
	InsecureSkipVerify   bool   `mapstructure:"insecureSkipVerify"`
	Url                  string `mapstructure:"url"`
	Token                string `mapstructure:"token"`
	BlkScannerHysteresis int64  `mapstructure:"blk_scanner_hysteresis"`
}

// Ipfs represents the IPFS node access configuration
type Ipfs struct {
	// Url of the IPFS node
	Url string `mapstructure:"url"`

	// Skip known HTTP-to-IPFS gateways and use our IPFS node instead
	SkipHttpGateways bool `mapstructure:"skip_http_gateways"`

	// Gateway to process IPFS requests instead of IPFS node (like Pinata)
	Gateway string `mapstructure:"gateway"`

	// GatewayBearer represents API key (JWT) to be used for Gateway auth
	GatewayBearer string `mapstructure:"gateway_bearer"`

	// FileCacheDir is a directory, where can be uploaded files cached
	FileCacheDir string `mapstructure:"file_cache_dir"`

	EmulateOnSharedDb bool `mapstructure:"emulate_on_shared_db"`
}

// Database represents the database access configuration.
type Database struct {
	Url    string `mapstructure:"url"`
	DbName string `mapstructure:"db"`
}

// Cache represents the cache sub-system configuration.
type Cache struct {
	Eviction time.Duration `mapstructure:"eviction"`
	MaxSize  int           `mapstructure:"size"`
}

// Auth represents the authentication configuration.
type Auth struct {
	BearerSecret string `mapstructure:"bearer_secret"`
	NonceSecret  string `mapstructure:"nonce_secret"`
}

// RandomFeedOracle configures random oracle feed.
type RandomFeedOracle struct {
	PrivateKey ecdsa.PrivateKey `mapstructure:"pk"`
	ChainID    string           `mapstructure:"chain"`
}

// NotificationProviders configures notification providers APIs.
type NotificationProviders struct {
	SendGrid *SendGridEmailProvider `mapstructure:"sendgrid,omitempty"`
}

// SendGridEmailProvider configures notification provider of SendGrid email.
type SendGridEmailProvider struct {
	ApiAddress string `mapstructure:"domain"`
	ApiKey     string `mapstructure:"key"`
}

// Contracts provides addresses of mandatory contracts
type Contracts struct {
	WrappedFTM common.Address `mapstructure:"wftm"`
}
