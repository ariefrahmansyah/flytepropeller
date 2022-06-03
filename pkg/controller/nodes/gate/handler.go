package gate

import (
	"context"
	"fmt"
	"time"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/service"

	"github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	"github.com/flyteorg/flytepropeller/pkg/controller/config"
	"github.com/flyteorg/flytepropeller/pkg/controller/nodes/errors"
	"github.com/flyteorg/flytepropeller/pkg/controller/nodes/handler"
	//"github.com/flyteorg/flytepropeller/pkg/controller/nodes/recovery"

	"github.com/flyteorg/flytestdlib/logger"
	"github.com/flyteorg/flytestdlib/promutils"
	//"github.com/flyteorg/flytestdlib/promutils/labeled"
	"github.com/flyteorg/flytestdlib/storage"
)

type gateNodeHandler struct {
	signalClient service.SignalServiceClient
	metrics      metrics
}

type metrics struct {
	//CacheError labeled.Counter
}

func newMetrics(scope promutils.Scope) metrics {
	return metrics{
		//CacheError: labeled.NewCounter("cache_err", "workflow handler failed to store or load from data store.", scope),
	}
}

func (g *gateNodeHandler) FinalizeRequired() bool {
	return false
}

func (g *gateNodeHandler) Setup(_ context.Context, _ handler.SetupContext) error {
	return nil
}

func (g *gateNodeHandler) Handle(ctx context.Context, nCtx handler.NodeExecutionContext) (handler.Transition, error) {
	gateNode := nCtx.Node().GetGateNode()

	// retrieve gate node status?
	gateNodeState := nCtx.NodeStateReader().GetGateNodeState()

	if gateNodeState.Phase == v1alpha1.GateNodePhaseUndefined {
		// Using GateNodeState to store a StartedAt timestamp because the NodeStatus on
		// NodeExecutionContext is deprecated
		gateNodeState.Phase = v1alpha1.GateNodePhaseExecuting
		gateNodeState.StartedAt = time.Now()
	}

	switch gateNode.GetKind() {
	case v1alpha1.ConditionKindSignal:
		// retrieve signal duration
		signalCondition := gateNode.GetSignal()
		if signalCondition == nil {
			errMsg := "gateNode signal condition is nil"
			return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoFailure(core.ExecutionError_SYSTEM,
				errors.BadSpecificationError, errMsg, nil)), nil
		}

		// use admin client to query for signal
		request := &admin.SignalGetOrCreateRequest{
			Id: &core.SignalIdentifier{
				ExecutionId: nCtx.ExecutionContext().GetExecutionID().WorkflowExecutionIdentifier,
				SignalId:    signalCondition.SignalId,
			},
			Type: signalCondition.Type,
		}

		signal, err := g.signalClient.GetOrCreateSignal(ctx, request)
		if err != nil {
			return handler.UnknownTransition, err
		}

		// if signal has value then write to output and transition to success
		if signal.Value.Value != nil {
			outputs := &core.LiteralMap{
				Literals: map[string]*core.Literal{
					// TODO - please verify
					// I believe the TypedInterface proto is only designed for tasks and workflows.
					// Therefore, I think we need to use a constant variable name here for signal
					// outputs so that it is translatable between flytekit and the backend.
					"o0": signal.Value,
				},
			}

			outputFile := v1alpha1.GetOutputsFile(nCtx.NodeStatus().GetOutputDir())

			so := storage.Options{}
			if err := nCtx.DataStore().WriteProtobuf(ctx, outputFile, so, outputs); err != nil {
				logger.Errorf(ctx, "Failed to write signal outputs. Error [%v]", err)
				return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoFailure(core.ExecutionError_SYSTEM, "WriteOutputsFailed",
					fmt.Sprintf("Failed to write signal value to [%v]. Error: %s", outputFile, err.Error()), nil)), nil
			}

			o := &handler.OutputInfo{OutputURI: outputFile}
			return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoSuccess(&handler.ExecutionInfo{
				OutputInfo: o,
			})), nil
		}
	case v1alpha1.ConditionKindSleep:
		// retrieve sleep duration
		sleepCondition := gateNode.GetSleep()
		if sleepCondition == nil {
			errMsg := "gateNode sleep condition is nil"
			return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoFailure(core.ExecutionError_SYSTEM,
				errors.BadSpecificationError, errMsg, nil)), nil
		}

		sleepDuration := sleepCondition.GetDuration().AsDuration()

		// check duration of node sleep
		now := time.Now()
		if sleepDuration <= now.Sub(gateNodeState.StartedAt) {
			return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoSuccess(&handler.ExecutionInfo{})), nil
		}
	default:
		errMsg := "gateNode does not have a supported condition reference"
		return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoFailure(core.ExecutionError_SYSTEM,
			errors.BadSpecificationError, errMsg, nil)), nil
	}

	// update gate node status
	if err := nCtx.NodeStateWriter().PutGateNodeState(gateNodeState); err != nil {
		logger.Errorf(ctx, "Failed to store TaskNode state, err :%s", err.Error())
		return handler.UnknownTransition, err
	}

	return handler.DoTransition(handler.TransitionTypeEphemeral, handler.PhaseInfoRunning(&handler.ExecutionInfo{})), nil
}

func (b *gateNodeHandler) Abort(ctx context.Context, nCtx handler.NodeExecutionContext, reason string) error {
	/*wfNode := nCtx.Node().GetWorkflowNode()
	if wfNode.GetSubWorkflowRef() != nil {
		return w.subWfHandler.HandleAbort(ctx, nCtx, reason)
	}

	if wfNode.GetLaunchPlanRefID() != nil {
		return w.lpHandler.HandleAbort(ctx, nCtx, reason)
	}*/
	return nil
}

func (w *gateNodeHandler) Finalize(ctx context.Context, _ handler.NodeExecutionContext) error {
	//logger.Warnf(ctx, "Subworkflow finalize invoked. Nothing to be done")
	return nil
}

func New(eventConfig *config.EventConfig, signalClient service.SignalServiceClient, scope promutils.Scope) handler.Node {
	gateScope := scope.NewSubScope("gate")
	return &gateNodeHandler{
		signalClient: signalClient,
		metrics:      newMetrics(gateScope),
	}
}
