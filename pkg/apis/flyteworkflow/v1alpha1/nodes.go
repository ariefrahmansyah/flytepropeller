package v1alpha1

import (
	"bytes"
	"time"

	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/golang/protobuf/jsonpb"
	typesv1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var marshaler = jsonpb.Marshaler{}

type OutputVarMap struct {
	*core.VariableMap
}

func (in *OutputVarMap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if err := marshaler.Marshal(&buf, in.VariableMap); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (in *OutputVarMap) UnmarshalJSON(b []byte) error {
	in.VariableMap = &core.VariableMap{}
	return jsonpb.Unmarshal(bytes.NewReader(b), in.VariableMap)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputVarMap) DeepCopyInto(out *OutputVarMap) {
	*out = *in
	// We do not manipulate the object, so its ok
	// Once we figure out the autogenerate story we can replace this
}

type Binding struct {
	*core.Binding
}

func (in *Binding) UnmarshalJSON(b []byte) error {
	in.Binding = &core.Binding{}
	return jsonpb.Unmarshal(bytes.NewReader(b), in.Binding)
}

func (in *Binding) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if err := marshaler.Marshal(&buf, in.Binding); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Binding) DeepCopyInto(out *Binding) {
	*out = *in
	// We do not manipulate the object, so its ok
	// Once we figure out the autogenerate story we can replace this
}

// Strategy to be used to Retry a node that is in RetryableFailure state
type RetryStrategy struct {
	// MinAttempts implies the atleast n attempts to try this node before giving up. The atleast here is because we may
	// fail to write the attempt information and end up retrying again.
	// Also `0` and `1` both mean atleast one attempt will be done. 0 is a degenerate case.
	MinAttempts *int `json:"minAttempts"`
	// TODO Add retrydelay?
}

type Alias struct {
	core.Alias
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alias) DeepCopyInto(out *Alias) {
	*out = *in
	// We do not manipulate the object, so its ok
	// Once we figure out the autogenerate story we can replace this
}

type NodeMetadata struct {
	core.NodeMetadata
}

func (in *NodeMetadata) DeepCopyInto(out *NodeMetadata) {
	*out = *in
	// We do not manipulate the object, so its ok
	// Once we figure out the autogenerate story we can replace this
}

type NodeSpec struct {
	ID            NodeID                        `json:"id"`
	Name          string                        `json:"name,omitempty"`
	Resources     *typesv1.ResourceRequirements `json:"resources,omitempty"`
	Kind          NodeKind                      `json:"kind"`
	BranchNode    *BranchNodeSpec               `json:"branch,omitempty"`
	TaskRef       *TaskID                       `json:"task,omitempty"`
	WorkflowNode  *WorkflowNodeSpec             `json:"workflow,omitempty"`
	InputBindings []*Binding                    `json:"inputBindings,omitempty"`
	Config        *typesv1.ConfigMap            `json:"config,omitempty"`
	RetryStrategy *RetryStrategy                `json:"retry,omitempty"`
	OutputAliases []Alias                       `json:"outputAlias,omitempty"`

	// SecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	SecurityContext *typesv1.PodSecurityContext `json:"securityContext,omitempty" protobuf:"bytes,14,opt,name=securityContext"`
	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec.
	// If specified, these secrets will be passed to individual puller implementations for them to use. For example,
	// in the case of docker, only DockerConfig type secrets are honored.
	// More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod
	// +optional
	// +patchMergeKey=name
	// +patchStrategy=merge
	ImagePullSecrets []typesv1.LocalObjectReference `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"`
	// Specifies the hostname of the Pod
	// If not specified, the pod's hostname will be set to a system-defined value.
	// +optional
	Hostname string `json:"hostname,omitempty" protobuf:"bytes,16,opt,name=hostname"`
	// If specified, the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>".
	// If not specified, the pod will not have a domainname at all.
	// +optional
	Subdomain string `json:"subdomain,omitempty" protobuf:"bytes,17,opt,name=subdomain"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *typesv1.Affinity `json:"affinity,omitempty" protobuf:"bytes,18,opt,name=affinity"`
	// If specified, the pod will be dispatched by specified scheduler.
	// If not specified, the pod will be dispatched by default scheduler.
	// +optional
	SchedulerName string `json:"schedulerName,omitempty" protobuf:"bytes,19,opt,name=schedulerName"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []typesv1.Toleration `json:"tolerations,omitempty" protobuf:"bytes,22,opt,name=tolerations"`
	// Node execution timeout
	ExecutionDeadline *v1.Duration `json:"executionDeadline,omitempty"`
	// StartTime before the system will actively try to mark it failed and kill associated containers.
	// Value must be a positive integer. This includes time spent waiting in the queue.
	// +optional
	ActiveDeadline *v1.Duration `json:"activeDeadline,omitempty"`
	// The value set to True means task is OK with getting interrupted
	// +optional
	Interruptible *bool `json:"interruptible,omitempty"`
}

func (in *NodeSpec) GetName() string {
	return in.Name
}

func (in *NodeSpec) GetRetryStrategy() *RetryStrategy {
	return in.RetryStrategy
}

func (in *NodeSpec) GetExecutionDeadline() *time.Duration {
	if in.ExecutionDeadline != nil {
		return &in.ExecutionDeadline.Duration
	}
	return nil
}

func (in *NodeSpec) GetActiveDeadline() *time.Duration {
	if in.ActiveDeadline != nil {
		return &in.ActiveDeadline.Duration
	}
	return nil
}

func (in *NodeSpec) IsInterruptible() *bool {
	return in.Interruptible
}

func (in *NodeSpec) GetConfig() *typesv1.ConfigMap {
	return in.Config
}

func (in *NodeSpec) GetResources() *typesv1.ResourceRequirements {
	return in.Resources
}

func (in *NodeSpec) GetOutputAlias() []Alias {
	return in.OutputAliases
}

func (in *NodeSpec) GetWorkflowNode() ExecutableWorkflowNode {
	if in.WorkflowNode == nil {
		return nil
	}
	return in.WorkflowNode
}

func (in *NodeSpec) GetBranchNode() ExecutableBranchNode {
	if in.BranchNode == nil {
		return nil
	}
	return in.BranchNode
}

func (in *NodeSpec) GetTaskID() *TaskID {
	return in.TaskRef
}

func (in *NodeSpec) GetKind() NodeKind {
	return in.Kind
}

func (in *NodeSpec) GetID() NodeID {
	return in.ID
}

func (in *NodeSpec) IsStartNode() bool {
	return in.ID == StartNodeID
}

func (in *NodeSpec) IsEndNode() bool {
	return in.ID == EndNodeID
}

func (in *NodeSpec) GetInputBindings() []*Binding {
	return in.InputBindings
}
