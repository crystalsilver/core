package hub

import (
	"github.com/sonm-io/core/insonmnia/structs"
	pb "github.com/sonm-io/core/proto"
)

type ETH interface {
	// WaitForDealCreated waits for deal created on Buyer-side
	WaitForDealCreated(request *structs.DealRequest) (*pb.Deal, error)

	// AcceptDeal approves deal on Hub-side
	AcceptDeal(id string) error

	// CheckDealExists checks whether a given deal exists.
	CheckDealExists(id string) (bool, error)
}
