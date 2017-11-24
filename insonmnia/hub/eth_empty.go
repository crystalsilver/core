// +build infra

package hub

import (
	"crypto/ecdsa"

	"github.com/pborman/uuid"
	"github.com/sonm-io/core/insonmnia/structs"
	pb "github.com/sonm-io/core/proto"
	"golang.org/x/net/context"
)

type emptyEth struct{}

func (e *emptyEth) WaitForDealCreated(request *structs.DealRequest) (*pb.Deal, error) {
	return &pb.Deal{
		Id:                uuid.New(),
		Status:            pb.DealStatus_PENDING,
		SpecificationHash: request.SpecHash,
	}, nil
}

func (e *emptyEth) AcceptDeal(id string) error {
	return nil
}

func (e *emptyEth) CheckDealExists(id string) (bool, error) {
	return true, nil
}

// NewETH constructs a new Ethereum client.
func NewETH(ctx context.Context, key *ecdsa.PrivateKey) (ETH, error) {
	return &emptyEth{}, nil
}
