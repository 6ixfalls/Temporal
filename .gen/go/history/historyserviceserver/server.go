// The MIT License (MIT)
// 
// Copyright (c) 2019 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by thriftrw-plugin-yarpc
// @generated

package historyserviceserver

import (
	context "context"
	history "github.com/uber/cadence/.gen/go/history"
	replicator "github.com/uber/cadence/.gen/go/replicator"
	shared "github.com/uber/cadence/.gen/go/shared"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
)

// Interface is the server-side interface for the HistoryService service.
type Interface interface {
	CloseShard(
		ctx context.Context,
		Request *shared.CloseShardRequest,
	) error

	DescribeHistoryHost(
		ctx context.Context,
		Request *shared.DescribeHistoryHostRequest,
	) (*shared.DescribeHistoryHostResponse, error)

	DescribeMutableState(
		ctx context.Context,
		Request *history.DescribeMutableStateRequest,
	) (*history.DescribeMutableStateResponse, error)

	DescribeWorkflowExecution(
		ctx context.Context,
		DescribeRequest *history.DescribeWorkflowExecutionRequest,
	) (*shared.DescribeWorkflowExecutionResponse, error)

	GetMutableState(
		ctx context.Context,
		GetRequest *history.GetMutableStateRequest,
	) (*history.GetMutableStateResponse, error)

	GetReplicationMessages(
		ctx context.Context,
		Request *replicator.GetReplicationMessagesRequest,
	) (*replicator.GetReplicationMessagesResponse, error)

	PollMutableState(
		ctx context.Context,
		PollRequest *history.PollMutableStateRequest,
	) (*history.PollMutableStateResponse, error)

	QueryWorkflow(
		ctx context.Context,
		QueryRequest *history.QueryWorkflowRequest,
	) (*history.QueryWorkflowResponse, error)

	ReapplyEvents(
		ctx context.Context,
		ReapplyEventsRequest *history.ReapplyEventsRequest,
	) error

	RecordActivityTaskHeartbeat(
		ctx context.Context,
		HeartbeatRequest *history.RecordActivityTaskHeartbeatRequest,
	) (*shared.RecordActivityTaskHeartbeatResponse, error)

	RecordActivityTaskStarted(
		ctx context.Context,
		AddRequest *history.RecordActivityTaskStartedRequest,
	) (*history.RecordActivityTaskStartedResponse, error)

	RecordChildExecutionCompleted(
		ctx context.Context,
		CompletionRequest *history.RecordChildExecutionCompletedRequest,
	) error

	RecordDecisionTaskStarted(
		ctx context.Context,
		AddRequest *history.RecordDecisionTaskStartedRequest,
	) (*history.RecordDecisionTaskStartedResponse, error)

	RemoveSignalMutableState(
		ctx context.Context,
		RemoveRequest *history.RemoveSignalMutableStateRequest,
	) error

	RemoveTask(
		ctx context.Context,
		Request *shared.RemoveTaskRequest,
	) error

	ReplicateEvents(
		ctx context.Context,
		ReplicateRequest *history.ReplicateEventsRequest,
	) error

	ReplicateEventsV2(
		ctx context.Context,
		ReplicateV2Request *history.ReplicateEventsV2Request,
	) error

	ReplicateRawEvents(
		ctx context.Context,
		ReplicateRequest *history.ReplicateRawEventsRequest,
	) error

	RequestCancelWorkflowExecution(
		ctx context.Context,
		CancelRequest *history.RequestCancelWorkflowExecutionRequest,
	) error

	ResetStickyTaskList(
		ctx context.Context,
		ResetRequest *history.ResetStickyTaskListRequest,
	) (*history.ResetStickyTaskListResponse, error)

	ResetWorkflowExecution(
		ctx context.Context,
		ResetRequest *history.ResetWorkflowExecutionRequest,
	) (*shared.ResetWorkflowExecutionResponse, error)

	RespondActivityTaskCanceled(
		ctx context.Context,
		CanceledRequest *history.RespondActivityTaskCanceledRequest,
	) error

	RespondActivityTaskCompleted(
		ctx context.Context,
		CompleteRequest *history.RespondActivityTaskCompletedRequest,
	) error

	RespondActivityTaskFailed(
		ctx context.Context,
		FailRequest *history.RespondActivityTaskFailedRequest,
	) error

	RespondDecisionTaskCompleted(
		ctx context.Context,
		CompleteRequest *history.RespondDecisionTaskCompletedRequest,
	) (*history.RespondDecisionTaskCompletedResponse, error)

	RespondDecisionTaskFailed(
		ctx context.Context,
		FailedRequest *history.RespondDecisionTaskFailedRequest,
	) error

	ScheduleDecisionTask(
		ctx context.Context,
		ScheduleRequest *history.ScheduleDecisionTaskRequest,
	) error

	SignalWithStartWorkflowExecution(
		ctx context.Context,
		SignalWithStartRequest *history.SignalWithStartWorkflowExecutionRequest,
	) (*shared.StartWorkflowExecutionResponse, error)

	SignalWorkflowExecution(
		ctx context.Context,
		SignalRequest *history.SignalWorkflowExecutionRequest,
	) error

	StartWorkflowExecution(
		ctx context.Context,
		StartRequest *history.StartWorkflowExecutionRequest,
	) (*shared.StartWorkflowExecutionResponse, error)

	SyncActivity(
		ctx context.Context,
		SyncActivityRequest *history.SyncActivityRequest,
	) error

	SyncShardStatus(
		ctx context.Context,
		SyncShardStatusRequest *history.SyncShardStatusRequest,
	) error

	TerminateWorkflowExecution(
		ctx context.Context,
		TerminateRequest *history.TerminateWorkflowExecutionRequest,
	) error
}

// New prepares an implementation of the HistoryService service for
// registration.
//
// 	handler := HistoryServiceHandler{}
// 	dispatcher.Register(historyserviceserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "HistoryService",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "CloseShard",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.CloseShard),
				},
				Signature:    "CloseShard(Request *shared.CloseShardRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "DescribeHistoryHost",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.DescribeHistoryHost),
				},
				Signature:    "DescribeHistoryHost(Request *shared.DescribeHistoryHostRequest) (*shared.DescribeHistoryHostResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "DescribeMutableState",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.DescribeMutableState),
				},
				Signature:    "DescribeMutableState(Request *history.DescribeMutableStateRequest) (*history.DescribeMutableStateResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "DescribeWorkflowExecution",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.DescribeWorkflowExecution),
				},
				Signature:    "DescribeWorkflowExecution(DescribeRequest *history.DescribeWorkflowExecutionRequest) (*shared.DescribeWorkflowExecutionResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "GetMutableState",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.GetMutableState),
				},
				Signature:    "GetMutableState(GetRequest *history.GetMutableStateRequest) (*history.GetMutableStateResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "GetReplicationMessages",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.GetReplicationMessages),
				},
				Signature:    "GetReplicationMessages(Request *replicator.GetReplicationMessagesRequest) (*replicator.GetReplicationMessagesResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "PollMutableState",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.PollMutableState),
				},
				Signature:    "PollMutableState(PollRequest *history.PollMutableStateRequest) (*history.PollMutableStateResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "QueryWorkflow",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.QueryWorkflow),
				},
				Signature:    "QueryWorkflow(QueryRequest *history.QueryWorkflowRequest) (*history.QueryWorkflowResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "ReapplyEvents",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ReapplyEvents),
				},
				Signature:    "ReapplyEvents(ReapplyEventsRequest *history.ReapplyEventsRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RecordActivityTaskHeartbeat",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RecordActivityTaskHeartbeat),
				},
				Signature:    "RecordActivityTaskHeartbeat(HeartbeatRequest *history.RecordActivityTaskHeartbeatRequest) (*shared.RecordActivityTaskHeartbeatResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RecordActivityTaskStarted",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RecordActivityTaskStarted),
				},
				Signature:    "RecordActivityTaskStarted(AddRequest *history.RecordActivityTaskStartedRequest) (*history.RecordActivityTaskStartedResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RecordChildExecutionCompleted",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RecordChildExecutionCompleted),
				},
				Signature:    "RecordChildExecutionCompleted(CompletionRequest *history.RecordChildExecutionCompletedRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RecordDecisionTaskStarted",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RecordDecisionTaskStarted),
				},
				Signature:    "RecordDecisionTaskStarted(AddRequest *history.RecordDecisionTaskStartedRequest) (*history.RecordDecisionTaskStartedResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RemoveSignalMutableState",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RemoveSignalMutableState),
				},
				Signature:    "RemoveSignalMutableState(RemoveRequest *history.RemoveSignalMutableStateRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RemoveTask",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RemoveTask),
				},
				Signature:    "RemoveTask(Request *shared.RemoveTaskRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "ReplicateEvents",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ReplicateEvents),
				},
				Signature:    "ReplicateEvents(ReplicateRequest *history.ReplicateEventsRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "ReplicateEventsV2",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ReplicateEventsV2),
				},
				Signature:    "ReplicateEventsV2(ReplicateV2Request *history.ReplicateEventsV2Request)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "ReplicateRawEvents",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ReplicateRawEvents),
				},
				Signature:    "ReplicateRawEvents(ReplicateRequest *history.ReplicateRawEventsRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RequestCancelWorkflowExecution",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RequestCancelWorkflowExecution),
				},
				Signature:    "RequestCancelWorkflowExecution(CancelRequest *history.RequestCancelWorkflowExecutionRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "ResetStickyTaskList",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ResetStickyTaskList),
				},
				Signature:    "ResetStickyTaskList(ResetRequest *history.ResetStickyTaskListRequest) (*history.ResetStickyTaskListResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "ResetWorkflowExecution",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ResetWorkflowExecution),
				},
				Signature:    "ResetWorkflowExecution(ResetRequest *history.ResetWorkflowExecutionRequest) (*shared.ResetWorkflowExecutionResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RespondActivityTaskCanceled",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RespondActivityTaskCanceled),
				},
				Signature:    "RespondActivityTaskCanceled(CanceledRequest *history.RespondActivityTaskCanceledRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RespondActivityTaskCompleted",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RespondActivityTaskCompleted),
				},
				Signature:    "RespondActivityTaskCompleted(CompleteRequest *history.RespondActivityTaskCompletedRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RespondActivityTaskFailed",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RespondActivityTaskFailed),
				},
				Signature:    "RespondActivityTaskFailed(FailRequest *history.RespondActivityTaskFailedRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RespondDecisionTaskCompleted",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RespondDecisionTaskCompleted),
				},
				Signature:    "RespondDecisionTaskCompleted(CompleteRequest *history.RespondDecisionTaskCompletedRequest) (*history.RespondDecisionTaskCompletedResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "RespondDecisionTaskFailed",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.RespondDecisionTaskFailed),
				},
				Signature:    "RespondDecisionTaskFailed(FailedRequest *history.RespondDecisionTaskFailedRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "ScheduleDecisionTask",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.ScheduleDecisionTask),
				},
				Signature:    "ScheduleDecisionTask(ScheduleRequest *history.ScheduleDecisionTaskRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "SignalWithStartWorkflowExecution",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.SignalWithStartWorkflowExecution),
				},
				Signature:    "SignalWithStartWorkflowExecution(SignalWithStartRequest *history.SignalWithStartWorkflowExecutionRequest) (*shared.StartWorkflowExecutionResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "SignalWorkflowExecution",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.SignalWorkflowExecution),
				},
				Signature:    "SignalWorkflowExecution(SignalRequest *history.SignalWorkflowExecutionRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "StartWorkflowExecution",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.StartWorkflowExecution),
				},
				Signature:    "StartWorkflowExecution(StartRequest *history.StartWorkflowExecutionRequest) (*shared.StartWorkflowExecutionResponse)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "SyncActivity",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.SyncActivity),
				},
				Signature:    "SyncActivity(SyncActivityRequest *history.SyncActivityRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "SyncShardStatus",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.SyncShardStatus),
				},
				Signature:    "SyncShardStatus(SyncShardStatusRequest *history.SyncShardStatusRequest)",
				ThriftModule: history.ThriftModule,
			},

			thrift.Method{
				Name: "TerminateWorkflowExecution",
				HandlerSpec: thrift.HandlerSpec{

					Type:  transport.Unary,
					Unary: thrift.UnaryHandler(h.TerminateWorkflowExecution),
				},
				Signature:    "TerminateWorkflowExecution(TerminateRequest *history.TerminateWorkflowExecutionRequest)",
				ThriftModule: history.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 33)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

func (h handler) CloseShard(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_CloseShard_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.CloseShard(ctx, args.Request)

	hadError := err != nil
	result, err := history.HistoryService_CloseShard_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) DescribeHistoryHost(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_DescribeHistoryHost_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.DescribeHistoryHost(ctx, args.Request)

	hadError := err != nil
	result, err := history.HistoryService_DescribeHistoryHost_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) DescribeMutableState(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_DescribeMutableState_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.DescribeMutableState(ctx, args.Request)

	hadError := err != nil
	result, err := history.HistoryService_DescribeMutableState_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) DescribeWorkflowExecution(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_DescribeWorkflowExecution_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.DescribeWorkflowExecution(ctx, args.DescribeRequest)

	hadError := err != nil
	result, err := history.HistoryService_DescribeWorkflowExecution_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) GetMutableState(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_GetMutableState_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.GetMutableState(ctx, args.GetRequest)

	hadError := err != nil
	result, err := history.HistoryService_GetMutableState_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) GetReplicationMessages(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_GetReplicationMessages_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.GetReplicationMessages(ctx, args.Request)

	hadError := err != nil
	result, err := history.HistoryService_GetReplicationMessages_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) PollMutableState(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_PollMutableState_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.PollMutableState(ctx, args.PollRequest)

	hadError := err != nil
	result, err := history.HistoryService_PollMutableState_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) QueryWorkflow(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_QueryWorkflow_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.QueryWorkflow(ctx, args.QueryRequest)

	hadError := err != nil
	result, err := history.HistoryService_QueryWorkflow_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ReapplyEvents(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_ReapplyEvents_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.ReapplyEvents(ctx, args.ReapplyEventsRequest)

	hadError := err != nil
	result, err := history.HistoryService_ReapplyEvents_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RecordActivityTaskHeartbeat(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RecordActivityTaskHeartbeat_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.RecordActivityTaskHeartbeat(ctx, args.HeartbeatRequest)

	hadError := err != nil
	result, err := history.HistoryService_RecordActivityTaskHeartbeat_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RecordActivityTaskStarted(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RecordActivityTaskStarted_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.RecordActivityTaskStarted(ctx, args.AddRequest)

	hadError := err != nil
	result, err := history.HistoryService_RecordActivityTaskStarted_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RecordChildExecutionCompleted(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RecordChildExecutionCompleted_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.RecordChildExecutionCompleted(ctx, args.CompletionRequest)

	hadError := err != nil
	result, err := history.HistoryService_RecordChildExecutionCompleted_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RecordDecisionTaskStarted(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RecordDecisionTaskStarted_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.RecordDecisionTaskStarted(ctx, args.AddRequest)

	hadError := err != nil
	result, err := history.HistoryService_RecordDecisionTaskStarted_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RemoveSignalMutableState(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RemoveSignalMutableState_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.RemoveSignalMutableState(ctx, args.RemoveRequest)

	hadError := err != nil
	result, err := history.HistoryService_RemoveSignalMutableState_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RemoveTask(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RemoveTask_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.RemoveTask(ctx, args.Request)

	hadError := err != nil
	result, err := history.HistoryService_RemoveTask_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ReplicateEvents(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_ReplicateEvents_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.ReplicateEvents(ctx, args.ReplicateRequest)

	hadError := err != nil
	result, err := history.HistoryService_ReplicateEvents_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ReplicateEventsV2(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_ReplicateEventsV2_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.ReplicateEventsV2(ctx, args.ReplicateV2Request)

	hadError := err != nil
	result, err := history.HistoryService_ReplicateEventsV2_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ReplicateRawEvents(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_ReplicateRawEvents_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.ReplicateRawEvents(ctx, args.ReplicateRequest)

	hadError := err != nil
	result, err := history.HistoryService_ReplicateRawEvents_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RequestCancelWorkflowExecution(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RequestCancelWorkflowExecution_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.RequestCancelWorkflowExecution(ctx, args.CancelRequest)

	hadError := err != nil
	result, err := history.HistoryService_RequestCancelWorkflowExecution_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ResetStickyTaskList(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_ResetStickyTaskList_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.ResetStickyTaskList(ctx, args.ResetRequest)

	hadError := err != nil
	result, err := history.HistoryService_ResetStickyTaskList_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ResetWorkflowExecution(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_ResetWorkflowExecution_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.ResetWorkflowExecution(ctx, args.ResetRequest)

	hadError := err != nil
	result, err := history.HistoryService_ResetWorkflowExecution_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RespondActivityTaskCanceled(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RespondActivityTaskCanceled_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.RespondActivityTaskCanceled(ctx, args.CanceledRequest)

	hadError := err != nil
	result, err := history.HistoryService_RespondActivityTaskCanceled_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RespondActivityTaskCompleted(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RespondActivityTaskCompleted_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.RespondActivityTaskCompleted(ctx, args.CompleteRequest)

	hadError := err != nil
	result, err := history.HistoryService_RespondActivityTaskCompleted_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RespondActivityTaskFailed(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RespondActivityTaskFailed_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.RespondActivityTaskFailed(ctx, args.FailRequest)

	hadError := err != nil
	result, err := history.HistoryService_RespondActivityTaskFailed_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RespondDecisionTaskCompleted(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RespondDecisionTaskCompleted_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.RespondDecisionTaskCompleted(ctx, args.CompleteRequest)

	hadError := err != nil
	result, err := history.HistoryService_RespondDecisionTaskCompleted_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) RespondDecisionTaskFailed(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_RespondDecisionTaskFailed_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.RespondDecisionTaskFailed(ctx, args.FailedRequest)

	hadError := err != nil
	result, err := history.HistoryService_RespondDecisionTaskFailed_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) ScheduleDecisionTask(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_ScheduleDecisionTask_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.ScheduleDecisionTask(ctx, args.ScheduleRequest)

	hadError := err != nil
	result, err := history.HistoryService_ScheduleDecisionTask_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) SignalWithStartWorkflowExecution(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_SignalWithStartWorkflowExecution_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.SignalWithStartWorkflowExecution(ctx, args.SignalWithStartRequest)

	hadError := err != nil
	result, err := history.HistoryService_SignalWithStartWorkflowExecution_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) SignalWorkflowExecution(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_SignalWorkflowExecution_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.SignalWorkflowExecution(ctx, args.SignalRequest)

	hadError := err != nil
	result, err := history.HistoryService_SignalWorkflowExecution_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) StartWorkflowExecution(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_StartWorkflowExecution_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	success, err := h.impl.StartWorkflowExecution(ctx, args.StartRequest)

	hadError := err != nil
	result, err := history.HistoryService_StartWorkflowExecution_Helper.WrapResponse(success, err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) SyncActivity(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_SyncActivity_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.SyncActivity(ctx, args.SyncActivityRequest)

	hadError := err != nil
	result, err := history.HistoryService_SyncActivity_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) SyncShardStatus(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_SyncShardStatus_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.SyncShardStatus(ctx, args.SyncShardStatusRequest)

	hadError := err != nil
	result, err := history.HistoryService_SyncShardStatus_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}

func (h handler) TerminateWorkflowExecution(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args history.HistoryService_TerminateWorkflowExecution_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, err
	}

	err := h.impl.TerminateWorkflowExecution(ctx, args.TerminateRequest)

	hadError := err != nil
	result, err := history.HistoryService_TerminateWorkflowExecution_Helper.WrapResponse(err)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
	}
	return response, err
}
