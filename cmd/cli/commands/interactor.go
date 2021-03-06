package commands

import (
	"time"

	"github.com/sonm-io/core/insonmnia/structs"
	pb "github.com/sonm-io/core/proto"
	"github.com/sonm-io/core/util"
	"golang.org/x/net/context"
)

type CliInteractor interface {
	HubPing(context.Context) (*pb.PingReply, error)
	HubStatus(context.Context) (*pb.HubStatusReply, error)
	HubShowSlots(ctx context.Context) (*pb.SlotsReply, error)
	HubInsertSlot(ctx context.Context, slot *structs.Slot, price string) (*pb.ID, error)

	MinerList(context.Context) (*pb.ListReply, error)
	MinerStatus(minerID string, appCtx context.Context) (*pb.InfoReply, error)

	TaskList(appCtx context.Context, minerID string) (*pb.StatusMapReply, error)
	TaskLogs(appCtx context.Context, req *pb.TaskLogsRequest) (pb.Hub_TaskLogsClient, error)
	TaskPush(ctx context.Context) (pb.Hub_PushTaskClient, error)
	TaskPull(ctx context.Context, dealID, name, taskID string) (pb.Hub_PullTaskClient, error)
	TaskStart(appCtx context.Context, req *pb.HubStartTaskRequest) (*pb.HubStartTaskReply, error)
	TaskStatus(appCtx context.Context, taskID string) (*pb.TaskStatusReply, error)
	TaskStop(appCtx context.Context, taskID string) (*pb.Empty, error)
}

type grpcInteractor struct {
	hub     pb.HubClient
	timeout time.Duration
}

func (it *grpcInteractor) call(addr string) error {

	return nil
}

func (it *grpcInteractor) ctx(appCtx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(appCtx, it.timeout)
}

func (it *grpcInteractor) HubPing(appCtx context.Context) (*pb.PingReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()
	return it.hub.Ping(ctx, &pb.Empty{})
}

func (it *grpcInteractor) HubStatus(appCtx context.Context) (*pb.HubStatusReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()
	return it.hub.Status(ctx, &pb.Empty{})
}

func (it *grpcInteractor) HubDevices(c context.Context) (*pb.DevicesReply, error) {
	ctx, cancel := it.ctx(c)
	defer cancel()
	return it.hub.Devices(ctx, &pb.Empty{})
}

func (it *grpcInteractor) HubGetProperties(ctx context.Context, ID string) (*pb.GetDevicePropertiesReply, error) {
	c, cancel := it.ctx(ctx)
	defer cancel()

	req := pb.ID{Id: ID}
	return it.hub.GetDeviceProperties(c, &req)
}

func (it *grpcInteractor) HubSetProperties(ctx context.Context, ID string, properties map[string]float64) (*pb.Empty, error) {
	c, cancel := it.ctx(ctx)
	defer cancel()

	req := pb.SetDevicePropertiesRequest{
		ID:         ID,
		Properties: properties,
	}
	return it.hub.SetDeviceProperties(c, &req)
}

func (it *grpcInteractor) HubShowSlots(ctx context.Context) (*pb.SlotsReply, error) {
	c, cancel := it.ctx(ctx)
	defer cancel()

	return it.hub.Slots(c, &pb.Empty{})
}

func (it *grpcInteractor) HubInsertSlot(ctx context.Context, slot *structs.Slot, price string) (*pb.ID, error) {
	c, cancel := it.ctx(ctx)
	defer cancel()

	req := &pb.InsertSlotRequest{
		Price: price,
		Slot:  slot.Unwrap(),
	}
	return it.hub.InsertSlot(c, req)
}

func (it *grpcInteractor) MinerList(appCtx context.Context) (*pb.ListReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()
	return it.hub.List(ctx, &pb.Empty{})
}

func (it *grpcInteractor) MinerStatus(minerID string, appCtx context.Context) (*pb.InfoReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()

	return it.hub.Info(ctx, &pb.ID{Id: minerID})
}

func (it *grpcInteractor) TaskList(appCtx context.Context, minerID string) (*pb.StatusMapReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()

	req := &pb.ID{Id: minerID}
	return it.hub.MinerStatus(ctx, req)
}

func (it *grpcInteractor) TaskLogs(appCtx context.Context, req *pb.TaskLogsRequest) (pb.Hub_TaskLogsClient, error) {
	return it.hub.TaskLogs(appCtx, req)
}

func (it *grpcInteractor) TaskPush(ctx context.Context) (pb.Hub_PushTaskClient, error) {
	return it.hub.PushTask(ctx)
}

func (it *grpcInteractor) TaskPull(ctx context.Context, dealID, name, taskID string) (pb.Hub_PullTaskClient, error) {
	return it.hub.PullTask(ctx, &pb.PullTaskRequest{
		DealId: dealID,
		Name:   name,
		TaskId: taskID,
	})
}

func (it *grpcInteractor) TaskStart(appCtx context.Context, req *pb.HubStartTaskRequest) (*pb.HubStartTaskReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()
	return it.hub.StartTask(ctx, req)
}

func (it *grpcInteractor) TaskStatus(appCtx context.Context, taskID string) (*pb.TaskStatusReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()

	var req = &pb.ID{Id: taskID}
	return it.hub.TaskStatus(ctx, req)
}

func (it *grpcInteractor) TaskStop(appCtx context.Context, taskID string) (*pb.Empty, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()

	var req = &pb.ID{Id: taskID}
	return it.hub.StopTask(ctx, req)
}

func NewGrpcInteractor(addr string, to time.Duration) (CliInteractor, error) {
	cc, err := util.MakeGrpcClient(context.Background(), addr, creds)
	if err != nil {
		return nil, err
	}

	return &grpcInteractor{
		hub:     pb.NewHubClient(cc),
		timeout: to,
	}, nil
}

type NodeHubInteractor interface {
	Status() (*pb.HubStatusReply, error)

	WorkersList() (*pb.ListReply, error)
	WorkerStatus(id string) (*pb.InfoReply, error)

	GetRegisteredWorkers() (*pb.GetRegisteredWorkersReply, error)
	RegisterWorker(id string) (*pb.Empty, error)
	DeregisterWorker(id string) (*pb.Empty, error)

	DevicesList() (*pb.DevicesReply, error)
	GetDeviceProperties(id string) (*pb.GetDevicePropertiesReply, error)
	SetDeviceProperties(ID string, properties map[string]float64) (*pb.Empty, error)

	GetAskPlans() (*pb.SlotsReply, error)
	CreateAskPlan(slot *structs.Slot, price string) (*pb.ID, error)
	RemoveAskPlan(id string) (*pb.Empty, error)

	TaskList() (*pb.TaskListReply, error)
	TaskStatus(id string) (*pb.TaskStatusReply, error)
}

type hubInteractor struct {
	timeout time.Duration
	hub     pb.HubManagementClient
}

func (it *hubInteractor) Status() (*pb.HubStatusReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.hub.Status(ctx, &pb.Empty{})
}

func (it *hubInteractor) WorkersList() (*pb.ListReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.hub.WorkersList(ctx, &pb.Empty{})
}

func (it *hubInteractor) WorkerStatus(id string) (*pb.InfoReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	req := &pb.ID{Id: id}
	return it.hub.WorkerStatus(ctx, req)
}

func (it *hubInteractor) GetRegisteredWorkers() (*pb.GetRegisteredWorkersReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.hub.GetRegisteredWorkers(ctx, &pb.Empty{})
}

func (it *hubInteractor) RegisterWorker(id string) (*pb.Empty, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	req := &pb.ID{Id: id}
	return it.hub.RegisterWorker(ctx, req)
}

func (it *hubInteractor) DeregisterWorker(id string) (*pb.Empty, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	req := &pb.ID{Id: id}
	return it.hub.DeregisterWorker(ctx, req)
}

func (it *hubInteractor) DevicesList() (*pb.DevicesReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.hub.DeviceList(ctx, &pb.Empty{})
}

func (it *hubInteractor) GetDeviceProperties(id string) (*pb.GetDevicePropertiesReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	req := &pb.ID{Id: id}
	return it.hub.GetDeviceProperties(ctx, req)
}

func (it *hubInteractor) SetDeviceProperties(ID string, properties map[string]float64) (*pb.Empty, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	req := &pb.SetDevicePropertiesRequest{
		ID:         ID,
		Properties: properties,
	}

	return it.hub.SetDeviceProperties(ctx, req)
}

func (it *hubInteractor) GetAskPlans() (*pb.SlotsReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.hub.GetAskPlans(ctx, &pb.Empty{})
}

func (it *hubInteractor) CreateAskPlan(slot *structs.Slot, price string) (*pb.ID, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	req := &pb.InsertSlotRequest{
		Price: price,
		Slot:  slot.Unwrap(),
	}
	return it.hub.CreateAskPlan(ctx, req)
}

func (it *hubInteractor) RemoveAskPlan(id string) (*pb.Empty, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.hub.RemoveAskPlan(ctx, &pb.ID{Id: id})
}

func (it *hubInteractor) TaskList() (*pb.TaskListReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.hub.TaskList(ctx, &pb.Empty{})
}

func (it *hubInteractor) TaskStatus(id string) (*pb.TaskStatusReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	req := &pb.ID{Id: id}
	return it.hub.TaskStatus(ctx, req)
}

func NewHubInteractor(addr string, timeout time.Duration) (NodeHubInteractor, error) {
	cc, err := util.MakeGrpcClient(context.Background(), addr, creds)
	if err != nil {
		return nil, err
	}

	return &hubInteractor{
		timeout: timeout,
		hub:     pb.NewHubManagementClient(cc),
	}, nil
}

type NodeMarketInteractor interface {
	GetOrders(slot *structs.Slot, orderType pb.OrderType, count uint64) ([]*pb.Order, error)
	GetProcessing() (*pb.GetProcessingReply, error)
	GetOrderByID(id string) (*pb.Order, error)
	CreateOrder(order *pb.Order) (*pb.Order, error)
	CancelOrder(id string) error
}

type marketInteractor struct {
	timeout time.Duration
	market  pb.MarketClient
}

func (it *marketInteractor) GetOrders(slot *structs.Slot, orderType pb.OrderType, count uint64) ([]*pb.Order, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	req := &pb.GetOrdersRequest{
		Slot:      slot.Unwrap(),
		OrderType: orderType,
		Count:     count,
	}

	reply, err := it.market.GetOrders(ctx, req)
	if err != nil {
		return nil, err
	}

	return reply.GetOrders(), nil
}

func (it *marketInteractor) GetProcessing() (*pb.GetProcessingReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.market.GetProcessing(ctx, &pb.Empty{})
}

func (it *marketInteractor) GetOrderByID(id string) (*pb.Order, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.market.GetOrderByID(ctx, &pb.ID{Id: id})
}

func (it *marketInteractor) CreateOrder(order *pb.Order) (*pb.Order, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.market.CreateOrder(ctx, order)
}

func (it *marketInteractor) CancelOrder(id string) error {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	_, err := it.market.CancelOrder(ctx, &pb.Order{Id: id})
	return err
}

func ctx(timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

func NewMarketInteractor(addr string, timeout time.Duration) (NodeMarketInteractor, error) {
	cc, err := util.MakeGrpcClient(context.Background(), addr, creds)
	if err != nil {
		return nil, err
	}

	market := pb.NewMarketClient(cc)

	return &marketInteractor{
		timeout: timeout,
		market:  market,
	}, nil

}

type DealsInteractor interface {
	List(from string, status pb.DealStatus) ([]*pb.Deal, error)
	Status(id string) (*pb.Deal, error)
	FinishDeal(id string) error
}

type dealsInteractor struct {
	timeout time.Duration
	deals   pb.DealManagementClient
}

func (it *dealsInteractor) List(from string, status pb.DealStatus) ([]*pb.Deal, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	req := &pb.DealListRequest{Owner: from, Status: status}
	reply, err := it.deals.List(ctx, req)
	if err != nil {
		return nil, err
	}

	return reply.GetDeal(), nil
}

func (it *dealsInteractor) Status(id string) (*pb.Deal, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.deals.Status(ctx, &pb.ID{Id: id})
}

func (it *dealsInteractor) FinishDeal(id string) error {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	_, err := it.deals.Finish(ctx, &pb.ID{Id: id})
	return err
}

func NewDealsInteractor(addr string, timeout time.Duration) (DealsInteractor, error) {
	cc, err := util.MakeGrpcClient(context.Background(), addr, creds)
	if err != nil {
		return nil, err
	}

	return &dealsInteractor{
		timeout: timeout,
		deals:   pb.NewDealManagementClient(cc),
	}, nil
}

type TasksInteractor interface {
	List(hubAddr string) (*pb.TaskListReply, error)
	Start(req *pb.HubStartTaskRequest) (*pb.HubStartTaskReply, error)
	Status(id, hub string) (*pb.TaskStatusReply, error)
	Logs(req *pb.TaskLogsRequest) (pb.TaskManagement_LogsClient, error)
	Stop(id, hub string) (*pb.Empty, error)
}

type tasksInteractor struct {
	timeout time.Duration
	tasks   pb.TaskManagementClient
}

func (it *tasksInteractor) List(hubAddr string) (*pb.TaskListReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	req := &pb.TaskListRequest{HubID: hubAddr}
	return it.tasks.List(ctx, req)
}

func (it *tasksInteractor) Start(req *pb.HubStartTaskRequest) (*pb.HubStartTaskReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.tasks.Start(ctx, req)
}

func (it *tasksInteractor) Status(id, hub string) (*pb.TaskStatusReply, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.tasks.Status(ctx, &pb.TaskID{Id: id, HubAddr: hub})
}

func (it *tasksInteractor) Logs(req *pb.TaskLogsRequest) (pb.TaskManagement_LogsClient, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.tasks.Logs(ctx, req)
}

func (it *tasksInteractor) Stop(id, hub string) (*pb.Empty, error) {
	ctx, cancel := ctx(it.timeout)
	defer cancel()

	return it.tasks.Stop(ctx, &pb.TaskID{Id: id, HubAddr: hub})
}

func NewTasksInteractor(addr string, timeout time.Duration) (TasksInteractor, error) {
	cc, err := util.MakeGrpcClient(context.Background(), addr, creds)
	if err != nil {
		return nil, err
	}

	return &tasksInteractor{
		timeout: timeout,
		tasks:   pb.NewTaskManagementClient(cc),
	}, nil
}
