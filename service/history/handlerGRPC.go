// Copyright (c) 2019 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package history

import (
	"context"

	"github.com/temporalio/temporal/.gen/go/shared"
	"github.com/temporalio/temporal/.gen/proto/healthservice"
	"github.com/temporalio/temporal/.gen/proto/historyservice"
	"github.com/temporalio/temporal/common/adapter"
	"github.com/temporalio/temporal/common/log"
	"github.com/temporalio/temporal/common/metrics"
)

// Disable lint due to missing comments.
// Code generated by generate-adapter. DO NOT EDIT.

var _ historyservice.HistoryServiceServer = (*HandlerGRPC)(nil)

type (
	// HandlerGRPC - gRPC handler interface for historyservice
	HandlerGRPC struct {
		handlerThrift *Handler
	}
)

// NewHandlerGRPC creates a gRPC handler for the historyservice
func NewHandlerGRPC(
	handlerThrift *Handler,
) *HandlerGRPC {
	handler := &HandlerGRPC{
		handlerThrift: handlerThrift,
	}

	return handler
}

// Health is for health check
func (h *HandlerGRPC) Health(context.Context, *healthservice.HealthRequest) (_ *healthservice.HealthStatus, retError error) {
	h.handlerThrift.startWG.Wait()
	h.handlerThrift.GetLogger().Debug("Matching service health check endpoint (gRPC) reached.")
	hs := &healthservice.HealthStatus{Ok: true, Msg: "matching good"}
	return hs, nil
}

func (h *HandlerGRPC) StartWorkflowExecution(ctx context.Context, request *historyservice.StartWorkflowExecutionRequest) (_ *historyservice.StartWorkflowExecutionResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.StartWorkflowExecution(ctx, adapter.ToThriftHistoryStartWorkflowExecutionRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoHistoryStartWorkflowExecutionResponse(resp), nil
}

func (h *HandlerGRPC) GetMutableState(ctx context.Context, request *historyservice.GetMutableStateRequest) (_ *historyservice.GetMutableStateResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.GetMutableState(ctx, adapter.ToThriftGetMutableStateRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoGetMutableStateResponse(resp), nil
}

func (h *HandlerGRPC) PollMutableState(ctx context.Context, request *historyservice.PollMutableStateRequest) (_ *historyservice.PollMutableStateResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.PollMutableState(ctx, adapter.ToThriftPollMutableStateRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoPollMutableStateResponse(resp), nil
}

func (h *HandlerGRPC) ResetStickyTaskList(ctx context.Context, request *historyservice.ResetStickyTaskListRequest) (_ *historyservice.ResetStickyTaskListResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.ResetStickyTaskList(ctx, adapter.ToThriftHistoryResetStickyTaskListRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoResetStickyTaskListResponse(resp), nil
}

func (h *HandlerGRPC) RecordDecisionTaskStarted(ctx context.Context, request *historyservice.RecordDecisionTaskStartedRequest) (_ *historyservice.RecordDecisionTaskStartedResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.RecordDecisionTaskStarted(ctx, adapter.ToThriftRecordDecisionTaskStartedRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoRecordDecisionTaskStartedResponse(resp), nil
}

func (h *HandlerGRPC) RecordActivityTaskStarted(ctx context.Context, request *historyservice.RecordActivityTaskStartedRequest) (_ *historyservice.RecordActivityTaskStartedResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.RecordActivityTaskStarted(ctx, adapter.ToThriftRecordActivityTaskStartedRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoRecordActivityTaskStartedResponse(resp), nil
}

func (h *HandlerGRPC) RespondDecisionTaskCompleted(ctx context.Context, request *historyservice.RespondDecisionTaskCompletedRequest) (_ *historyservice.RespondDecisionTaskCompletedResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.RespondDecisionTaskCompleted(ctx, adapter.ToThriftHistoryRespondDecisionTaskCompletedRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoHistoryRespondDecisionTaskCompletedResponse(resp), nil
}

func (h *HandlerGRPC) RespondDecisionTaskFailed(ctx context.Context, request *historyservice.RespondDecisionTaskFailedRequest) (_ *historyservice.RespondDecisionTaskFailedResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.RespondDecisionTaskFailed(ctx, adapter.ToThriftHistoryRespondDecisionTaskFailedRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.RespondDecisionTaskFailedResponse{}, nil
}

func (h *HandlerGRPC) RecordActivityTaskHeartbeat(ctx context.Context, request *historyservice.RecordActivityTaskHeartbeatRequest) (_ *historyservice.RecordActivityTaskHeartbeatResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.RecordActivityTaskHeartbeat(ctx, adapter.ToThriftHistoryRecordActivityTaskHeartbeatRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoHistoryRecordActivityTaskHeartbeatResponse(resp), nil
}

func (h *HandlerGRPC) RespondActivityTaskCompleted(ctx context.Context, request *historyservice.RespondActivityTaskCompletedRequest) (_ *historyservice.RespondActivityTaskCompletedResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.RespondActivityTaskCompleted(ctx, adapter.ToThriftHistoryRespondActivityTaskCompletedRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.RespondActivityTaskCompletedResponse{}, nil
}

func (h *HandlerGRPC) RespondActivityTaskFailed(ctx context.Context, request *historyservice.RespondActivityTaskFailedRequest) (_ *historyservice.RespondActivityTaskFailedResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.RespondActivityTaskFailed(ctx, adapter.ToThriftHistoryRespondActivityTaskFailedRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.RespondActivityTaskFailedResponse{}, nil
}

func (h *HandlerGRPC) RespondActivityTaskCanceled(ctx context.Context, request *historyservice.RespondActivityTaskCanceledRequest) (_ *historyservice.RespondActivityTaskCanceledResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.RespondActivityTaskCanceled(ctx, adapter.ToThriftHistoryRespondActivityTaskCanceledRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.RespondActivityTaskCanceledResponse{}, nil
}

func (h *HandlerGRPC) SignalWorkflowExecution(ctx context.Context, request *historyservice.SignalWorkflowExecutionRequest) (_ *historyservice.SignalWorkflowExecutionResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.SignalWorkflowExecution(ctx, adapter.ToThriftHistorySignalWorkflowExecutionRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.SignalWorkflowExecutionResponse{}, nil
}

func (h *HandlerGRPC) SignalWithStartWorkflowExecution(ctx context.Context, request *historyservice.SignalWithStartWorkflowExecutionRequest) (_ *historyservice.SignalWithStartWorkflowExecutionResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.SignalWithStartWorkflowExecution(ctx, adapter.ToThriftHistorySignalWithStartWorkflowExecutionRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoHistorySignalWithStartWorkflowExecutionResponse(resp), nil
}

func (h *HandlerGRPC) RemoveSignalMutableState(ctx context.Context, request *historyservice.RemoveSignalMutableStateRequest) (_ *historyservice.RemoveSignalMutableStateResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.RemoveSignalMutableState(ctx, adapter.ToThriftRemoveSignalMutableStateRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.RemoveSignalMutableStateResponse{}, nil
}

func (h *HandlerGRPC) TerminateWorkflowExecution(ctx context.Context, request *historyservice.TerminateWorkflowExecutionRequest) (_ *historyservice.TerminateWorkflowExecutionResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.TerminateWorkflowExecution(ctx, adapter.ToThriftHistoryTerminateWorkflowExecutionRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.TerminateWorkflowExecutionResponse{}, nil
}

func (h *HandlerGRPC) ResetWorkflowExecution(ctx context.Context, request *historyservice.ResetWorkflowExecutionRequest) (_ *historyservice.ResetWorkflowExecutionResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.ResetWorkflowExecution(ctx, adapter.ToThriftHistoryResetWorkflowExecutionRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoHistoryResetWorkflowExecutionResponse(resp), nil
}

func (h *HandlerGRPC) RequestCancelWorkflowExecution(ctx context.Context, request *historyservice.RequestCancelWorkflowExecutionRequest) (_ *historyservice.RequestCancelWorkflowExecutionResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.RequestCancelWorkflowExecution(ctx, adapter.ToThriftHistoryRequestCancelWorkflowExecutionRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.RequestCancelWorkflowExecutionResponse{}, nil
}

func (h *HandlerGRPC) ScheduleDecisionTask(ctx context.Context, request *historyservice.ScheduleDecisionTaskRequest) (_ *historyservice.ScheduleDecisionTaskResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.ScheduleDecisionTask(ctx, adapter.ToThriftScheduleDecisionTaskRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.ScheduleDecisionTaskResponse{}, nil
}

func (h *HandlerGRPC) RecordChildExecutionCompleted(ctx context.Context, request *historyservice.RecordChildExecutionCompletedRequest) (_ *historyservice.RecordChildExecutionCompletedResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.RecordChildExecutionCompleted(ctx, adapter.ToThriftRecordChildExecutionCompletedRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.RecordChildExecutionCompletedResponse{}, nil
}

func (h *HandlerGRPC) DescribeWorkflowExecution(ctx context.Context, request *historyservice.DescribeWorkflowExecutionRequest) (_ *historyservice.DescribeWorkflowExecutionResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.DescribeWorkflowExecution(ctx, adapter.ToThriftHistoryDescribeWorkflowExecutionRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoHistoryDescribeWorkflowExecutionResponse(resp), nil
}

func (h *HandlerGRPC) ReplicateEvents(ctx context.Context, request *historyservice.ReplicateEventsRequest) (_ *historyservice.ReplicateEventsResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.ReplicateEvents(ctx, adapter.ToThriftReplicateEventsRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.ReplicateEventsResponse{}, nil
}

func (h *HandlerGRPC) ReplicateRawEvents(ctx context.Context, request *historyservice.ReplicateRawEventsRequest) (_ *historyservice.ReplicateRawEventsResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.ReplicateRawEvents(ctx, adapter.ToThriftReplicateRawEventsRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.ReplicateRawEventsResponse{}, nil
}

func (h *HandlerGRPC) ReplicateEventsV2(ctx context.Context, request *historyservice.ReplicateEventsV2Request) (_ *historyservice.ReplicateEventsV2Response, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.ReplicateEventsV2(ctx, adapter.ToThriftReplicateEventsV2Request(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.ReplicateEventsV2Response{}, nil
}

func (h *HandlerGRPC) SyncShardStatus(ctx context.Context, request *historyservice.SyncShardStatusRequest) (_ *historyservice.SyncShardStatusResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.SyncShardStatus(ctx, adapter.ToThriftSyncShardStatusRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.SyncShardStatusResponse{}, nil
}

func (h *HandlerGRPC) SyncActivity(ctx context.Context, request *historyservice.SyncActivityRequest) (_ *historyservice.SyncActivityResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.SyncActivity(ctx, adapter.ToThriftSyncActivityRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.SyncActivityResponse{}, nil
}

func (h *HandlerGRPC) DescribeMutableState(ctx context.Context, request *historyservice.DescribeMutableStateRequest) (_ *historyservice.DescribeMutableStateResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.DescribeMutableState(ctx, adapter.ToThriftDescribeMutableStateRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoDescribeMutableStateResponse(resp), nil
}

func (h *HandlerGRPC) DescribeHistoryHost(ctx context.Context, request *historyservice.DescribeHistoryHostRequest) (_ *historyservice.DescribeHistoryHostResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.DescribeHistoryHost(ctx, adapter.ToThriftHistoryDescribeHistoryHostRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoHistoryDescribeHistoryHostResponse(resp), nil
}

func (h *HandlerGRPC) CloseShard(ctx context.Context, request *historyservice.CloseShardRequest) (_ *historyservice.CloseShardResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.CloseShard(ctx, adapter.ToThriftHistoryCloseShardRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.CloseShardResponse{}, nil
}

func (h *HandlerGRPC) RemoveTask(ctx context.Context, request *historyservice.RemoveTaskRequest) (_ *historyservice.RemoveTaskResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.RemoveTask(ctx, adapter.ToThriftHistoryRemoveTaskRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.RemoveTaskResponse{}, nil
}

func (h *HandlerGRPC) GetReplicationMessages(ctx context.Context, request *historyservice.GetReplicationMessagesRequest) (_ *historyservice.GetReplicationMessagesResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.GetReplicationMessages(ctx, adapter.ToThriftHistoryGetReplicationMessagesRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoHistoryGetReplicationMessagesResponse(resp), nil
}

func (h *HandlerGRPC) GetDLQReplicationMessages(ctx context.Context, request *historyservice.GetDLQReplicationMessagesRequest) (_ *historyservice.GetDLQReplicationMessagesResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.GetDLQReplicationMessages(ctx, adapter.ToThriftHistoryGetDLQReplicationMessagesRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoHistoryGetDLQReplicationMessagesResponse(resp), nil
}

func (h *HandlerGRPC) QueryWorkflow(ctx context.Context, request *historyservice.QueryWorkflowRequest) (_ *historyservice.QueryWorkflowResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	resp, err := h.handlerThrift.QueryWorkflow(ctx, adapter.ToThriftHistoryQueryWorkflowRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return adapter.ToProtoHistoryQueryWorkflowResponse(resp), nil
}

func (h *HandlerGRPC) ReapplyEvents(ctx context.Context, request *historyservice.ReapplyEventsRequest) (_ *historyservice.ReapplyEventsResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	err := h.handlerThrift.ReapplyEvents(ctx, adapter.ToThriftHistoryReapplyEventsRequest(request))
	if err != nil {
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.ReapplyEventsResponse{}, nil
}

func (h *HandlerGRPC) ReadDLQMessages(ctx context.Context, request *historyservice.ReadDLQMessagesRequest) (_ *historyservice.ReadDLQMessagesResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	h.handlerThrift.startWG.Wait()

	scope := metrics.HistoryReadDLQMessagesScope
	h.handlerThrift.GetMetricsClient().IncCounter(scope, metrics.CadenceRequests)
	sw := h.handlerThrift.GetMetricsClient().StartTimer(scope, metrics.CadenceLatency)
	defer sw.Stop()

	engine, err := h.handlerThrift.controller.getEngineForShard(int(request.GetShardID()))
	if err != nil {
		err = h.handlerThrift.error(err, scope, "", "")
		return nil, adapter.ToProtoError(err)
	}

	resp, err := engine.ReadDLQMessages(ctx, request)
	if err != nil {
		err = h.handlerThrift.error(err, scope, "", "")
		return nil, adapter.ToProtoError(err)
	}

	return resp, nil
}

func (h *HandlerGRPC) PurgeDLQMessages(ctx context.Context, request *historyservice.PurgeDLQMessagesRequest) (_ *historyservice.PurgeDLQMessagesResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	h.handlerThrift.startWG.Wait()

	scope := metrics.HistoryPurgeDLQMessagesScope
	h.handlerThrift.GetMetricsClient().IncCounter(scope, metrics.CadenceRequests)
	sw := h.handlerThrift.GetMetricsClient().StartTimer(scope, metrics.CadenceLatency)
	defer sw.Stop()

	engine, err := h.handlerThrift.controller.getEngineForShard(int(request.GetShardID()))
	if err != nil {
		err = h.handlerThrift.error(err, scope, "", "")
		return nil, adapter.ToProtoError(err)
	}

	err = engine.PurgeDLQMessages(ctx, request)
	if err != nil {
		err = h.handlerThrift.error(err, scope, "", "")
		return nil, adapter.ToProtoError(err)
	}
	return &historyservice.PurgeDLQMessagesResponse{}, nil
}

func (h *HandlerGRPC) MergeDLQMessages(ctx context.Context, request *historyservice.MergeDLQMessagesRequest) (_ *historyservice.MergeDLQMessagesResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	h.handlerThrift.startWG.Wait()

	scope := metrics.HistoryMergeDLQMessagesScope
	h.handlerThrift.GetMetricsClient().IncCounter(scope, metrics.CadenceRequests)
	sw := h.handlerThrift.GetMetricsClient().StartTimer(scope, metrics.CadenceLatency)
	defer sw.Stop()

	engine, err := h.handlerThrift.controller.getEngineForShard(int(request.GetShardID()))
	if err != nil {
		err = h.handlerThrift.error(err, scope, "", "")
		return nil, adapter.ToProtoError(err)
	}

	resp, err := engine.MergeDLQMessages(ctx, request)
	if err != nil {
		err = h.handlerThrift.error(err, scope, "", "")
		return nil, adapter.ToProtoError(err)
	}

	return resp, nil
}

func (h *HandlerGRPC) RefreshWorkflowTasks(ctx context.Context, request *historyservice.RefreshWorkflowTasksRequest) (_ *historyservice.RefreshWorkflowTasksResponse, retError error) {
	defer log.CapturePanicGRPC(h.handlerThrift.GetLogger(), &retError)

	h.handlerThrift.startWG.Wait()

	scope := metrics.HistoryRefreshWorkflowTasksScope
	h.handlerThrift.GetMetricsClient().IncCounter(scope, metrics.CadenceRequests)
	sw := h.handlerThrift.GetMetricsClient().StartTimer(scope, metrics.CadenceLatency)
	defer sw.Stop()
	domainID := request.GetDomainUUID()
	execution := request.GetRequest().GetExecution()
	workflowID := execution.GetWorkflowId()
	engine, err := h.handlerThrift.controller.GetEngine(workflowID)
	if err != nil {
		err = h.handlerThrift.error(err, scope, domainID, workflowID)
		return nil, adapter.ToProtoError(err)
	}

	err = engine.RefreshWorkflowTasks(
		ctx,
		domainID,
		shared.WorkflowExecution{
			WorkflowId: &execution.WorkflowId,
			RunId:      &execution.RunId,
		},
	)

	if err != nil {
		err = h.handlerThrift.error(err, scope, domainID, workflowID)
		return nil, adapter.ToProtoError(err)
	}

	return &historyservice.RefreshWorkflowTasksResponse{}, nil
}
