package hub

import (
	"crypto/ecdsa"

	"github.com/sonm-io/core/blockchain"
	pb "github.com/sonm-io/core/proto"
	"github.com/sonm-io/core/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
)

// options for building hub instance
type options struct {
	version       string
	ctx           context.Context
	ethKey        *ecdsa.PrivateKey
	ethAddr       string
	bcr           blockchain.Blockchainer
	market        pb.MarketClient
	locator       pb.LocatorClient
	creds         credentials.TransportCredentials
	rot           util.HitlessCertRotator
	cluster       Cluster
	clusterEvents <-chan ClusterEvent
}

func defaultHubOptions() *options {
	return &options{
		ctx: context.Background(),
	}
}

// Option func is for applying any params to hub options
type Option func(options *options)

func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

func WithBlockchain(bcr blockchain.Blockchainer) Option {
	return func(o *options) {
		o.bcr = bcr
	}
}

func WithMarket(mp pb.MarketClient) Option {
	return func(o *options) {
		o.market = mp
	}
}

func WithLocator(loc pb.LocatorClient) Option {
	return func(o *options) {
		o.locator = loc
	}
}

func WithPrivateKey(k *ecdsa.PrivateKey) Option {
	return func(o *options) {
		o.ethKey = k
		o.ethAddr = util.PubKeyToAddr(k.PublicKey)
	}
}

func WithVersion(v string) Option {
	return func(o *options) {
		o.version = v
	}
}

func WithCreds(creds credentials.TransportCredentials) Option {
	return func(o *options) {
		o.creds = creds
	}
}

func WithCertRotator(rot util.HitlessCertRotator) Option {
	return func(o *options) {
		o.rot = rot
	}
}

func WithCluster(cl Cluster, evt <-chan ClusterEvent) Option {
	return func(o *options) {
		o.cluster = cl
		o.clusterEvents = evt
	}
}
