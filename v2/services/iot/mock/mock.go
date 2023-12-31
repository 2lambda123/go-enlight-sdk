package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	iot_grpcapi "github.com/SKF/proto/v2/iot"

	"github.com/SKF/go-enlight-sdk/v2/services/iot"
)

type client struct {
	mock.Mock
}

func Create() *client { // nolint: golint
	return new(client)
}

var _ iot.IoTClient = &client{}

func (mock *client) Dial(host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(host, port, opts)
	return args.Error(0)
}
func (mock *client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error {
	args := mock.Called(ctx, host, port, opts)
	return args.Error(0)
}

func (mock *client) Close() {
	mock.Called()
}
func (mock *client) DeepPing() error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *client) DeepPingWithContext(ctx context.Context) error {
	args := mock.Called(ctx)
	return args.Error(0)
}

func (mock *client) CreateTask(task iot_grpcapi.InitialTaskDescription) (string, error) {
	args := mock.Called(task)
	return args.String(0), args.Error(1)
}
func (mock *client) CreateTaskWithContext(ctx context.Context, task iot_grpcapi.InitialTaskDescription) (string, error) {
	args := mock.Called(ctx, task)
	return args.String(0), args.Error(1)
}

func (mock *client) DeleteTask(userID, taskID string) error {
	args := mock.Called(userID, taskID)
	return args.Error(0)
}
func (mock *client) DeleteTaskWithContext(ctx context.Context, userID, taskID string) error {
	args := mock.Called(ctx, userID, taskID)
	return args.Error(0)
}

func (mock *client) SetTaskCompleted(userID, taskID string) error {
	args := mock.Called(userID, taskID)
	return args.Error(0)
}
func (mock *client) SetTaskCompletedWithContext(ctx context.Context, userID, taskID string) error {
	args := mock.Called(ctx, userID, taskID)
	return args.Error(0)
}

func (mock *client) GetAllTasks(userID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetAllTasksWithContext(ctx context.Context, userID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}

func (mock *client) GetUncompletedTasks(userID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetUncompletedTasksWithContext(ctx context.Context, userID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}

func (mock *client) GetUncompletedTasksByHierarchy(nodeID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(nodeID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetUncompletedTasksByHierarchyWithContext(ctx context.Context, nodeID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(ctx, nodeID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}

func (mock *client) GetActiveTasks(userID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(userID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetActiveTasksWithContext(ctx context.Context, userID string) ([]iot_grpcapi.TaskDescription, error) {
	args := mock.Called(ctx, userID)
	return args.Get(0).([]iot_grpcapi.TaskDescription), args.Error(1)
}

func (mock *client) SetTaskStatus(input iot_grpcapi.SetTaskStatusInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) SetTaskStatusWithContext(ctx context.Context, input iot_grpcapi.SetTaskStatusInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) IngestNodeData(input iot_grpcapi.IngestNodeDataInput) error {
	args := mock.Called(input)
	return args.Error(0)
}

func (mock *client) IngestNodeDataWithContext(ctx context.Context, input iot_grpcapi.IngestNodeDataInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) IngestNodesData(input iot_grpcapi.IngestNodesDataInput) error {
	args := mock.Called(input)
	return args.Error(0)
}

func (mock *client) IngestNodesDataWithContext(ctx context.Context, input iot_grpcapi.IngestNodesDataInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetLatestNodeData(input iot_grpcapi.GetLatestNodeDataInput) (*iot_grpcapi.NodeData, error) {
	args := mock.Called(input)
	return args.Get(0).(*iot_grpcapi.NodeData), args.Error(1)
}

func (mock *client) GetLatestNodeDataWithContext(ctx context.Context, input iot_grpcapi.GetLatestNodeDataInput) (*iot_grpcapi.NodeData, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*iot_grpcapi.NodeData), args.Error(1)
}

func (mock *client) GetNodeData(input iot_grpcapi.GetNodeDataInput) ([]iot_grpcapi.NodeData, error) {
	args := mock.Called(input)
	return args.Get(0).([]iot_grpcapi.NodeData), args.Error(1)
}
func (mock *client) GetNodeDataWithContext(ctx context.Context, input iot_grpcapi.GetNodeDataInput) ([]iot_grpcapi.NodeData, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).([]iot_grpcapi.NodeData), args.Error(1)
}

func (mock *client) GetMedia(input iot_grpcapi.GetMediaInput) (iot_grpcapi.Media, error) {
	args := mock.Called(input)
	return args.Get(0).(iot_grpcapi.Media), args.Error(1)
}
func (mock *client) GetMediaWithContext(ctx context.Context, input iot_grpcapi.GetMediaInput) (media iot_grpcapi.Media, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(iot_grpcapi.Media), args.Error(1)
}

func (mock *client) RequestGetMediaSignedURLWithContext(ctx context.Context, in *iot_grpcapi.GetMediaSignedUrlInput) (*iot_grpcapi.GetMediaSignedUrlOutput, error) {
	args := mock.Called(ctx, in)
	return args.Get(0).(*iot_grpcapi.GetMediaSignedUrlOutput), args.Error(1)
}
func (mock *client) RequestGetMediaSignedURL(in *iot_grpcapi.GetMediaSignedUrlInput) (*iot_grpcapi.GetMediaSignedUrlOutput, error) {
	args := mock.Called(in)
	return args.Get(0).(*iot_grpcapi.GetMediaSignedUrlOutput), args.Error(1)
}

func (mock *client) RequestPutMediaSignedURLWithContext(ctx context.Context, in *iot_grpcapi.PutMediaSignedUrlInput) (*iot_grpcapi.PutMediaSignedUrlOutput, error) {
	args := mock.Called(ctx, in)
	return args.Get(0).(*iot_grpcapi.PutMediaSignedUrlOutput), args.Error(1)
}
func (mock *client) RequestPutMediaSignedURL(in *iot_grpcapi.PutMediaSignedUrlInput) (*iot_grpcapi.PutMediaSignedUrlOutput, error) {
	args := mock.Called(in)
	return args.Get(0).(*iot_grpcapi.PutMediaSignedUrlOutput), args.Error(1)
}

func (mock *client) GetTasksByStatus(input iot_grpcapi.GetTasksByStatusInput) ([]*iot_grpcapi.TaskDescription, error) {
	args := mock.Called(input)
	return args.Get(0).([]*iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetTasksByStatusWithContext(ctx context.Context, input iot_grpcapi.GetTasksByStatusInput) ([]*iot_grpcapi.TaskDescription, error) {
	args := mock.Called(ctx, input)
	return args.Get(0).([]*iot_grpcapi.TaskDescription), args.Error(1)
}

func (mock *client) GetTaskByUUID(input string) (output *iot_grpcapi.TaskDescription, err error) {
	args := mock.Called(input)
	return args.Get(0).(*iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetTaskByUUIDWithContext(ctx context.Context, input string) (output *iot_grpcapi.TaskDescription, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetTaskByLongId(input int64) (output *iot_grpcapi.TaskDescription, err error) { // nolint: golint
	args := mock.Called(input)
	return args.Get(0).(*iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) GetTaskByLongIdWithContext(ctx context.Context, input int64) (output *iot_grpcapi.TaskDescription, err error) { // nolint: golint
	args := mock.Called(ctx, input)
	return args.Get(0).(*iot_grpcapi.TaskDescription), args.Error(1)
}
func (mock *client) DeleteNodeData(input iot_grpcapi.DeleteNodeDataInput) error {
	args := mock.Called(input)
	return args.Error(0)
}
func (mock *client) DeleteNodeDataWithContext(ctx context.Context, input iot_grpcapi.DeleteNodeDataInput) error {
	args := mock.Called(ctx, input)
	return args.Error(0)
}

func (mock *client) GetTasksModifiedSinceTimestamp(input iot_grpcapi.GetTasksModifiedSinceTimestampInput) (*iot_grpcapi.GetTasksModifiedSinceTimestampOutput, error) {
	args := mock.Called(input)
	return args.Get(0).(*iot_grpcapi.GetTasksModifiedSinceTimestampOutput), args.Error(1)
}
func (mock *client) GetTasksModifiedSinceTimestampWithContext(ctx context.Context, input iot_grpcapi.GetTasksModifiedSinceTimestampInput) (output *iot_grpcapi.GetTasksModifiedSinceTimestampOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*iot_grpcapi.GetTasksModifiedSinceTimestampOutput), args.Error(1)
}

func (mock *client) GetNodeEventLog(input iot_grpcapi.GetNodeEventLogInput) (output *iot_grpcapi.GetNodeEventLogOutput, err error) {
	args := mock.Called(input)
	return args.Get(0).(*iot_grpcapi.GetNodeEventLogOutput), args.Error(1)
}
func (mock *client) GetNodeEventLogWithContext(ctx context.Context, input iot_grpcapi.GetNodeEventLogInput) (output *iot_grpcapi.GetNodeEventLogOutput, err error) {
	args := mock.Called(ctx, input)
	return args.Get(0).(*iot_grpcapi.GetNodeEventLogOutput), args.Error(1)
}
