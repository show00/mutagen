// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synchronization/endpoint/remote/protocol.proto

package remote

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	synchronization "github.com/mutagen-io/mutagen/pkg/synchronization"
	core "github.com/mutagen-io/mutagen/pkg/synchronization/core"
	rsync "github.com/mutagen-io/mutagen/pkg/synchronization/rsync"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// InitializeSynchronizationRequest encodes a request for endpoint
// initialization.
type InitializeSynchronizationRequest struct {
	// Session is the session identifier.
	Session string `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	// Version is the session version.
	Version synchronization.Version `protobuf:"varint,2,opt,name=version,proto3,enum=synchronization.Version" json:"version,omitempty"`
	// Configuration is the session configuration.
	Configuration *synchronization.Configuration `protobuf:"bytes,3,opt,name=configuration,proto3" json:"configuration,omitempty"`
	// Root is the synchronization root path.
	Root string `protobuf:"bytes,4,opt,name=root,proto3" json:"root,omitempty"`
	// Alpha indicates whether or not the endpoint should behave as alpha (as
	// opposed to beta).
	Alpha                bool     `protobuf:"varint,5,opt,name=alpha,proto3" json:"alpha,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitializeSynchronizationRequest) Reset()         { *m = InitializeSynchronizationRequest{} }
func (m *InitializeSynchronizationRequest) String() string { return proto.CompactTextString(m) }
func (*InitializeSynchronizationRequest) ProtoMessage()    {}
func (*InitializeSynchronizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{0}
}

func (m *InitializeSynchronizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitializeSynchronizationRequest.Unmarshal(m, b)
}
func (m *InitializeSynchronizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitializeSynchronizationRequest.Marshal(b, m, deterministic)
}
func (m *InitializeSynchronizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitializeSynchronizationRequest.Merge(m, src)
}
func (m *InitializeSynchronizationRequest) XXX_Size() int {
	return xxx_messageInfo_InitializeSynchronizationRequest.Size(m)
}
func (m *InitializeSynchronizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InitializeSynchronizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InitializeSynchronizationRequest proto.InternalMessageInfo

func (m *InitializeSynchronizationRequest) GetSession() string {
	if m != nil {
		return m.Session
	}
	return ""
}

func (m *InitializeSynchronizationRequest) GetVersion() synchronization.Version {
	if m != nil {
		return m.Version
	}
	return synchronization.Version_Invalid
}

func (m *InitializeSynchronizationRequest) GetConfiguration() *synchronization.Configuration {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *InitializeSynchronizationRequest) GetRoot() string {
	if m != nil {
		return m.Root
	}
	return ""
}

func (m *InitializeSynchronizationRequest) GetAlpha() bool {
	if m != nil {
		return m.Alpha
	}
	return false
}

// InitializeSynchronizationResponse encodes initialization results.
type InitializeSynchronizationResponse struct {
	// Error is the error message (if any) resulting from initialization.
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitializeSynchronizationResponse) Reset()         { *m = InitializeSynchronizationResponse{} }
func (m *InitializeSynchronizationResponse) String() string { return proto.CompactTextString(m) }
func (*InitializeSynchronizationResponse) ProtoMessage()    {}
func (*InitializeSynchronizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{1}
}

func (m *InitializeSynchronizationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitializeSynchronizationResponse.Unmarshal(m, b)
}
func (m *InitializeSynchronizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitializeSynchronizationResponse.Marshal(b, m, deterministic)
}
func (m *InitializeSynchronizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitializeSynchronizationResponse.Merge(m, src)
}
func (m *InitializeSynchronizationResponse) XXX_Size() int {
	return xxx_messageInfo_InitializeSynchronizationResponse.Size(m)
}
func (m *InitializeSynchronizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitializeSynchronizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitializeSynchronizationResponse proto.InternalMessageInfo

func (m *InitializeSynchronizationResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// PollRequest encodes a request for one-shot polling.
type PollRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PollRequest) Reset()         { *m = PollRequest{} }
func (m *PollRequest) String() string { return proto.CompactTextString(m) }
func (*PollRequest) ProtoMessage()    {}
func (*PollRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{2}
}

func (m *PollRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollRequest.Unmarshal(m, b)
}
func (m *PollRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollRequest.Marshal(b, m, deterministic)
}
func (m *PollRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollRequest.Merge(m, src)
}
func (m *PollRequest) XXX_Size() int {
	return xxx_messageInfo_PollRequest.Size(m)
}
func (m *PollRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PollRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PollRequest proto.InternalMessageInfo

// PollCompletionRequest is paired with pollRequest and indicates a request for
// early polling completion or an acknowledgement of completion.
type PollCompletionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PollCompletionRequest) Reset()         { *m = PollCompletionRequest{} }
func (m *PollCompletionRequest) String() string { return proto.CompactTextString(m) }
func (*PollCompletionRequest) ProtoMessage()    {}
func (*PollCompletionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{3}
}

func (m *PollCompletionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollCompletionRequest.Unmarshal(m, b)
}
func (m *PollCompletionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollCompletionRequest.Marshal(b, m, deterministic)
}
func (m *PollCompletionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollCompletionRequest.Merge(m, src)
}
func (m *PollCompletionRequest) XXX_Size() int {
	return xxx_messageInfo_PollCompletionRequest.Size(m)
}
func (m *PollCompletionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PollCompletionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PollCompletionRequest proto.InternalMessageInfo

// PollResponse indicates polling completion.
type PollResponse struct {
	// Error is the error message (if any) resulting from polling.
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PollResponse) Reset()         { *m = PollResponse{} }
func (m *PollResponse) String() string { return proto.CompactTextString(m) }
func (*PollResponse) ProtoMessage()    {}
func (*PollResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{4}
}

func (m *PollResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollResponse.Unmarshal(m, b)
}
func (m *PollResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollResponse.Marshal(b, m, deterministic)
}
func (m *PollResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollResponse.Merge(m, src)
}
func (m *PollResponse) XXX_Size() int {
	return xxx_messageInfo_PollResponse.Size(m)
}
func (m *PollResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PollResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PollResponse proto.InternalMessageInfo

func (m *PollResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// ScanRequest encodes a request for a scan.
type ScanRequest struct {
	// BaseSnapshotSignature is the rsync signature to use as the base for
	// differentially transmitting snapshots.
	BaseSnapshotSignature *rsync.Signature `protobuf:"bytes,1,opt,name=baseSnapshotSignature,proto3" json:"baseSnapshotSignature,omitempty"`
	// Full indicates whether or not to force a full (warm) scan, temporarily
	// avoiding any acceleration that might be available on the endpoint.
	Full                 bool     `protobuf:"varint,2,opt,name=full,proto3" json:"full,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanRequest) Reset()         { *m = ScanRequest{} }
func (m *ScanRequest) String() string { return proto.CompactTextString(m) }
func (*ScanRequest) ProtoMessage()    {}
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{5}
}

func (m *ScanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanRequest.Unmarshal(m, b)
}
func (m *ScanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanRequest.Marshal(b, m, deterministic)
}
func (m *ScanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanRequest.Merge(m, src)
}
func (m *ScanRequest) XXX_Size() int {
	return xxx_messageInfo_ScanRequest.Size(m)
}
func (m *ScanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScanRequest proto.InternalMessageInfo

func (m *ScanRequest) GetBaseSnapshotSignature() *rsync.Signature {
	if m != nil {
		return m.BaseSnapshotSignature
	}
	return nil
}

func (m *ScanRequest) GetFull() bool {
	if m != nil {
		return m.Full
	}
	return false
}

// ScanResponse encodes the results of a scan.
type ScanResponse struct {
	// SnapshotDelta are the operations need to reconstruct the snapshot against
	// the specified base.
	SnapshotDelta []*rsync.Operation `protobuf:"bytes,1,rep,name=snapshotDelta,proto3" json:"snapshotDelta,omitempty"`
	// PreservesExecutability indicates whether or not the scan root preserves
	// POSIX executability bits.
	PreservesExecutability bool `protobuf:"varint,2,opt,name=preservesExecutability,proto3" json:"preservesExecutability,omitempty"`
	// Error is the error message (if any) resulting from scanning.
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// TryAgain indicates whether or not the error is ephermeral.
	TryAgain             bool     `protobuf:"varint,4,opt,name=tryAgain,proto3" json:"tryAgain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanResponse) Reset()         { *m = ScanResponse{} }
func (m *ScanResponse) String() string { return proto.CompactTextString(m) }
func (*ScanResponse) ProtoMessage()    {}
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{6}
}

func (m *ScanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanResponse.Unmarshal(m, b)
}
func (m *ScanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanResponse.Marshal(b, m, deterministic)
}
func (m *ScanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanResponse.Merge(m, src)
}
func (m *ScanResponse) XXX_Size() int {
	return xxx_messageInfo_ScanResponse.Size(m)
}
func (m *ScanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScanResponse proto.InternalMessageInfo

func (m *ScanResponse) GetSnapshotDelta() []*rsync.Operation {
	if m != nil {
		return m.SnapshotDelta
	}
	return nil
}

func (m *ScanResponse) GetPreservesExecutability() bool {
	if m != nil {
		return m.PreservesExecutability
	}
	return false
}

func (m *ScanResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ScanResponse) GetTryAgain() bool {
	if m != nil {
		return m.TryAgain
	}
	return false
}

// StageRequest encodes a request for staging.
type StageRequest struct {
	// Paths lists the paths that need to be staged.
	Paths []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	// Digests lists the digests for the paths that need to be staged. Its
	// length and contents correspond to that of Paths.
	Digests              [][]byte `protobuf:"bytes,2,rep,name=digests,proto3" json:"digests,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StageRequest) Reset()         { *m = StageRequest{} }
func (m *StageRequest) String() string { return proto.CompactTextString(m) }
func (*StageRequest) ProtoMessage()    {}
func (*StageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{7}
}

func (m *StageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StageRequest.Unmarshal(m, b)
}
func (m *StageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StageRequest.Marshal(b, m, deterministic)
}
func (m *StageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StageRequest.Merge(m, src)
}
func (m *StageRequest) XXX_Size() int {
	return xxx_messageInfo_StageRequest.Size(m)
}
func (m *StageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StageRequest proto.InternalMessageInfo

func (m *StageRequest) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *StageRequest) GetDigests() [][]byte {
	if m != nil {
		return m.Digests
	}
	return nil
}

// StageResponse encodes the results of staging initialization.
type StageResponse struct {
	// Paths are the paths that need to be staged (relative to the
	// synchronization root).
	Paths []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	// Signatures are the rsync signatures of the paths needing to be staged.
	Signatures []*rsync.Signature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	// Error is the error message (if any) resulting from staging
	// initialization.
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StageResponse) Reset()         { *m = StageResponse{} }
func (m *StageResponse) String() string { return proto.CompactTextString(m) }
func (*StageResponse) ProtoMessage()    {}
func (*StageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{8}
}

func (m *StageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StageResponse.Unmarshal(m, b)
}
func (m *StageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StageResponse.Marshal(b, m, deterministic)
}
func (m *StageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StageResponse.Merge(m, src)
}
func (m *StageResponse) XXX_Size() int {
	return xxx_messageInfo_StageResponse.Size(m)
}
func (m *StageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StageResponse proto.InternalMessageInfo

func (m *StageResponse) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *StageResponse) GetSignatures() []*rsync.Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *StageResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// SupplyRequest indicates a request for supplying files.
type SupplyRequest struct {
	// Paths are the paths to provide (relative to the synchronization root).
	Paths []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	// Signatures are the rsync signatures of the paths needing to be staged.
	Signatures           []*rsync.Signature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SupplyRequest) Reset()         { *m = SupplyRequest{} }
func (m *SupplyRequest) String() string { return proto.CompactTextString(m) }
func (*SupplyRequest) ProtoMessage()    {}
func (*SupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{9}
}

func (m *SupplyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupplyRequest.Unmarshal(m, b)
}
func (m *SupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupplyRequest.Marshal(b, m, deterministic)
}
func (m *SupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplyRequest.Merge(m, src)
}
func (m *SupplyRequest) XXX_Size() int {
	return xxx_messageInfo_SupplyRequest.Size(m)
}
func (m *SupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SupplyRequest proto.InternalMessageInfo

func (m *SupplyRequest) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *SupplyRequest) GetSignatures() []*rsync.Signature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// TransitionRequest encodes a request for transition application.
type TransitionRequest struct {
	// Transitions are the transitions that need to be applied.
	Transitions          []*core.Change `protobuf:"bytes,1,rep,name=transitions,proto3" json:"transitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransitionRequest) Reset()         { *m = TransitionRequest{} }
func (m *TransitionRequest) String() string { return proto.CompactTextString(m) }
func (*TransitionRequest) ProtoMessage()    {}
func (*TransitionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{10}
}

func (m *TransitionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitionRequest.Unmarshal(m, b)
}
func (m *TransitionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitionRequest.Marshal(b, m, deterministic)
}
func (m *TransitionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitionRequest.Merge(m, src)
}
func (m *TransitionRequest) XXX_Size() int {
	return xxx_messageInfo_TransitionRequest.Size(m)
}
func (m *TransitionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransitionRequest proto.InternalMessageInfo

func (m *TransitionRequest) GetTransitions() []*core.Change {
	if m != nil {
		return m.Transitions
	}
	return nil
}

// TransitionResponse encodes the results of transitioning.
type TransitionResponse struct {
	// Results are the resulting contents post-transition.
	// HACK: We have to use Archive to wrap our Entry results here because
	// Protocol Buffers won't encode a nil pointer in a repeated element in
	// certain cases, and the results of transitions may very well be nil. gob
	// also exhibits this problem.
	Results []*core.Archive `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	// Problems are any problems encountered during the transition operation.
	Problems []*core.Problem `protobuf:"bytes,2,rep,name=problems,proto3" json:"problems,omitempty"`
	// StagerMissingFiles indicates whether or not the endpoint's stager
	// indicated missing files during transitioning.
	StagerMissingFiles bool `protobuf:"varint,3,opt,name=stagerMissingFiles,proto3" json:"stagerMissingFiles,omitempty"`
	// Error is the error message (if any) resulting from the remote transition
	// method. This will always be an empty string since transition doesn't
	// return errors from local endpoints, but to match the endpoint interface
	// (which allows for transition errors due to network failures with remote
	// endpoints), we include this field.
	// TODO: Should we just remove this field? Doing so would rely on knowledge
	// of localEndpoint's transition behavior.
	Error                string   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransitionResponse) Reset()         { *m = TransitionResponse{} }
func (m *TransitionResponse) String() string { return proto.CompactTextString(m) }
func (*TransitionResponse) ProtoMessage()    {}
func (*TransitionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{11}
}

func (m *TransitionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitionResponse.Unmarshal(m, b)
}
func (m *TransitionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitionResponse.Marshal(b, m, deterministic)
}
func (m *TransitionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitionResponse.Merge(m, src)
}
func (m *TransitionResponse) XXX_Size() int {
	return xxx_messageInfo_TransitionResponse.Size(m)
}
func (m *TransitionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransitionResponse proto.InternalMessageInfo

func (m *TransitionResponse) GetResults() []*core.Archive {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *TransitionResponse) GetProblems() []*core.Problem {
	if m != nil {
		return m.Problems
	}
	return nil
}

func (m *TransitionResponse) GetStagerMissingFiles() bool {
	if m != nil {
		return m.StagerMissingFiles
	}
	return false
}

func (m *TransitionResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

// EndpointRequest is a sum type that can transmit any type of endpoint request.
// Only the sent request will be non-nil. We intentionally avoid using Protocol
// Buffers' oneof feature because it generates really ugly code and an unwieldy
// API, at least in Go. Manually checking for exclusivity is not difficult.
type EndpointRequest struct {
	// Poll represents a poll request.
	Poll *PollRequest `protobuf:"bytes,1,opt,name=poll,proto3" json:"poll,omitempty"`
	// Scan represents a scan request.
	Scan *ScanRequest `protobuf:"bytes,2,opt,name=scan,proto3" json:"scan,omitempty"`
	// Stage represents a stage request.
	Stage *StageRequest `protobuf:"bytes,3,opt,name=stage,proto3" json:"stage,omitempty"`
	// Supply represents a supply request.
	Supply *SupplyRequest `protobuf:"bytes,4,opt,name=supply,proto3" json:"supply,omitempty"`
	// Transition represents a transition request.
	Transition           *TransitionRequest `protobuf:"bytes,5,opt,name=transition,proto3" json:"transition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EndpointRequest) Reset()         { *m = EndpointRequest{} }
func (m *EndpointRequest) String() string { return proto.CompactTextString(m) }
func (*EndpointRequest) ProtoMessage()    {}
func (*EndpointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed323a11ce40f0df, []int{12}
}

func (m *EndpointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointRequest.Unmarshal(m, b)
}
func (m *EndpointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointRequest.Marshal(b, m, deterministic)
}
func (m *EndpointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointRequest.Merge(m, src)
}
func (m *EndpointRequest) XXX_Size() int {
	return xxx_messageInfo_EndpointRequest.Size(m)
}
func (m *EndpointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointRequest proto.InternalMessageInfo

func (m *EndpointRequest) GetPoll() *PollRequest {
	if m != nil {
		return m.Poll
	}
	return nil
}

func (m *EndpointRequest) GetScan() *ScanRequest {
	if m != nil {
		return m.Scan
	}
	return nil
}

func (m *EndpointRequest) GetStage() *StageRequest {
	if m != nil {
		return m.Stage
	}
	return nil
}

func (m *EndpointRequest) GetSupply() *SupplyRequest {
	if m != nil {
		return m.Supply
	}
	return nil
}

func (m *EndpointRequest) GetTransition() *TransitionRequest {
	if m != nil {
		return m.Transition
	}
	return nil
}

func init() {
	proto.RegisterType((*InitializeSynchronizationRequest)(nil), "remote.InitializeSynchronizationRequest")
	proto.RegisterType((*InitializeSynchronizationResponse)(nil), "remote.InitializeSynchronizationResponse")
	proto.RegisterType((*PollRequest)(nil), "remote.PollRequest")
	proto.RegisterType((*PollCompletionRequest)(nil), "remote.PollCompletionRequest")
	proto.RegisterType((*PollResponse)(nil), "remote.PollResponse")
	proto.RegisterType((*ScanRequest)(nil), "remote.ScanRequest")
	proto.RegisterType((*ScanResponse)(nil), "remote.ScanResponse")
	proto.RegisterType((*StageRequest)(nil), "remote.StageRequest")
	proto.RegisterType((*StageResponse)(nil), "remote.StageResponse")
	proto.RegisterType((*SupplyRequest)(nil), "remote.SupplyRequest")
	proto.RegisterType((*TransitionRequest)(nil), "remote.TransitionRequest")
	proto.RegisterType((*TransitionResponse)(nil), "remote.TransitionResponse")
	proto.RegisterType((*EndpointRequest)(nil), "remote.EndpointRequest")
}

func init() {
	proto.RegisterFile("synchronization/endpoint/remote/protocol.proto", fileDescriptor_ed323a11ce40f0df)
}

var fileDescriptor_ed323a11ce40f0df = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdf, 0x6b, 0x1b, 0x39,
	0x10, 0xc7, 0xd9, 0xd8, 0x4e, 0x9c, 0xb1, 0x7d, 0x3f, 0x74, 0xc9, 0xdd, 0x5e, 0xe0, 0x0e, 0x67,
	0xef, 0x20, 0xbe, 0x83, 0xec, 0x1e, 0x3e, 0x08, 0xe4, 0xa5, 0x90, 0x3a, 0x09, 0xf4, 0xa1, 0x34,
	0xc8, 0xa5, 0x85, 0xbe, 0xc9, 0x1b, 0x65, 0x2d, 0x2a, 0x4b, 0x5b, 0x49, 0x1b, 0xea, 0xfc, 0x51,
	0xa5, 0x7f, 0x52, 0xff, 0x8d, 0xbe, 0x95, 0x95, 0xb4, 0xf6, 0xda, 0x71, 0x5d, 0xfa, 0xa6, 0x91,
	0x3e, 0x33, 0x1a, 0x7d, 0x67, 0x67, 0x16, 0x62, 0x3d, 0x17, 0xe9, 0x54, 0x49, 0xc1, 0x1e, 0x88,
	0x61, 0x52, 0x24, 0x54, 0xdc, 0xe6, 0x92, 0x09, 0x93, 0x28, 0x3a, 0x93, 0x86, 0x26, 0xb9, 0x92,
	0x46, 0xa6, 0x92, 0xc7, 0x76, 0x81, 0x76, 0xdd, 0xf6, 0x51, 0xb4, 0xee, 0xa7, 0xca, 0x8d, 0x84,
	0x8a, 0x8c, 0x09, 0xea, 0xd8, 0xa3, 0xbf, 0xd6, 0x99, 0x54, 0x8a, 0x3b, 0x96, 0x15, 0xca, 0x5a,
	0x1e, 0xfa, 0x63, 0x1d, 0xba, 0xa7, 0x4a, 0x2f, 0x8f, 0xa3, 0xc7, 0x31, 0x14, 0x4d, 0x88, 0x4a,
	0xa7, 0xec, 0xbe, 0xba, 0xe7, 0x78, 0x23, 0x93, 0x4e, 0x89, 0xc8, 0xe8, 0xd6, 0x30, 0xb9, 0x92,
	0x13, 0x4e, 0x67, 0x8e, 0x89, 0x3e, 0x05, 0xd0, 0x7f, 0x26, 0x98, 0x61, 0x84, 0xb3, 0x07, 0x3a,
	0x5e, 0x75, 0xc0, 0xf4, 0x5d, 0x41, 0xb5, 0x41, 0x21, 0xec, 0x69, 0xaa, 0xcb, 0x04, 0xc3, 0xa0,
	0x1f, 0x0c, 0xf6, 0x71, 0x65, 0xa2, 0x21, 0xec, 0xf9, 0xd4, 0xc3, 0x9d, 0x7e, 0x30, 0xf8, 0x61,
	0x18, 0xae, 0x6b, 0x1b, 0xbf, 0x72, 0xe7, 0xb8, 0x02, 0xd1, 0x25, 0xf4, 0x56, 0x34, 0x09, 0x1b,
	0xfd, 0x60, 0xd0, 0x19, 0xfe, 0xf9, 0xc8, 0x73, 0x54, 0xa7, 0xf0, 0xaa, 0x13, 0x42, 0xd0, 0x54,
	0x52, 0x9a, 0xb0, 0x69, 0x13, 0xb2, 0x6b, 0x74, 0x00, 0x2d, 0xc2, 0xf3, 0x29, 0x09, 0x5b, 0xfd,
	0x60, 0xd0, 0xc6, 0xce, 0x88, 0xce, 0xe1, 0x78, 0xcb, 0x0b, 0x75, 0x2e, 0x85, 0xa6, 0xa5, 0x2b,
	0x55, 0x4a, 0x2a, 0xff, 0x40, 0x67, 0x44, 0x3d, 0xe8, 0xdc, 0x48, 0xce, 0xbd, 0x0e, 0xd1, 0x6f,
	0x70, 0x58, 0x9a, 0x23, 0x39, 0xcb, 0x39, 0xad, 0x09, 0x14, 0xfd, 0x0d, 0x5d, 0xc7, 0x6d, 0x8d,
	0xc6, 0xa0, 0x33, 0x4e, 0xc9, 0x42, 0xd5, 0x6b, 0x38, 0x9c, 0x10, 0x4d, 0xc7, 0x82, 0xe4, 0x7a,
	0x2a, 0xcd, 0x98, 0x65, 0x82, 0x98, 0x42, 0x51, 0xeb, 0xd4, 0x19, 0xfe, 0x14, 0xdb, 0xaf, 0x2b,
	0x5e, 0xec, 0xe3, 0xcd, 0x78, 0xa9, 0xc4, 0x5d, 0xc1, 0xb9, 0x2d, 0x40, 0x1b, 0xdb, 0x75, 0xf4,
	0x31, 0x80, 0xae, 0xbb, 0xcb, 0x67, 0x74, 0x06, 0x3d, 0xed, 0x3d, 0x2f, 0x29, 0x37, 0x24, 0x0c,
	0xfa, 0x8d, 0xda, 0x25, 0x2f, 0x72, 0x5a, 0xc9, 0xbc, 0x82, 0xa1, 0x33, 0xf8, 0x35, 0x57, 0x54,
	0x53, 0x75, 0x4f, 0xf5, 0xd5, 0x7b, 0x9a, 0x16, 0x86, 0x4c, 0x18, 0x67, 0x66, 0xee, 0xaf, 0xfb,
	0xca, 0xe9, 0x52, 0x81, 0x46, 0x4d, 0x01, 0x74, 0x04, 0x6d, 0xa3, 0xe6, 0x17, 0x19, 0x61, 0xc2,
	0x16, 0xae, 0x8d, 0x17, 0x76, 0xf4, 0x04, 0xba, 0x63, 0x43, 0x32, 0x5a, 0xc9, 0x73, 0x00, 0xad,
	0x9c, 0x98, 0xa9, 0xb6, 0x99, 0xee, 0x63, 0x67, 0x94, 0x9f, 0xe2, 0x2d, 0xcb, 0xa8, 0x36, 0x3a,
	0xdc, 0xe9, 0x37, 0x06, 0x5d, 0x5c, 0x99, 0xd1, 0x0c, 0x7a, 0xde, 0x7f, 0x59, 0x84, 0x0d, 0x01,
	0xfe, 0x03, 0xd0, 0x95, 0x74, 0x2e, 0xc6, 0x26, 0xa9, 0x6b, 0xcc, 0xe6, 0xa7, 0x44, 0xaf, 0xa1,
	0x37, 0x2e, 0xf2, 0x9c, 0xcf, 0xb7, 0xe7, 0xfb, 0xdd, 0xd7, 0x45, 0x23, 0xf8, 0xf9, 0xa5, 0x22,
	0x42, 0xb3, 0x7a, 0x07, 0xc6, 0xd0, 0x31, 0x8b, 0x4d, 0xed, 0x8b, 0xd7, 0x8d, 0xcb, 0x86, 0x8e,
	0x47, 0xb6, 0xe7, 0x71, 0x1d, 0x88, 0x3e, 0x04, 0x80, 0xea, 0x51, 0xbc, 0x24, 0x27, 0xb0, 0xa7,
	0xa8, 0x2e, 0xb8, 0xa9, 0x42, 0xf4, 0x5c, 0x88, 0x0b, 0x37, 0x5a, 0x70, 0x75, 0x8a, 0xfe, 0x81,
	0xb6, 0x9f, 0x13, 0x55, 0xd2, 0x9e, 0xbc, 0x71, 0xbb, 0x78, 0x71, 0x8c, 0x62, 0x40, 0xba, 0xd4,
	0x5d, 0x3d, 0x67, 0x5a, 0x33, 0x91, 0x5d, 0x33, 0x4e, 0xb5, 0xd5, 0xaa, 0x8d, 0x37, 0x9c, 0x2c,
	0xe5, 0x6c, 0xd6, 0xe5, 0xfc, 0x1c, 0xc0, 0x8f, 0x57, 0x7e, 0x0a, 0x57, 0x8f, 0x3e, 0x81, 0x66,
	0x2e, 0x39, 0xf7, 0xfd, 0xf0, 0x4b, 0xec, 0xa6, 0x70, 0x5c, 0xeb, 0x48, 0x6c, 0x81, 0x12, 0xd4,
	0x29, 0x71, 0x23, 0xa8, 0x06, 0xd6, 0x9a, 0x0d, 0x5b, 0x00, 0xfd, 0x0b, 0x2d, 0x9b, 0x91, 0x1f,
	0x39, 0x07, 0x0b, 0xb2, 0xf6, 0xe1, 0x61, 0x87, 0xa0, 0x53, 0xd8, 0xd5, 0xb6, 0xc0, 0x36, 0xd1,
	0xce, 0xf0, 0x70, 0x01, 0xd7, 0xcb, 0x8e, 0x3d, 0x84, 0xce, 0x01, 0x96, 0x05, 0xb0, 0x03, 0xa8,
	0x33, 0xfc, 0xbd, 0x72, 0x79, 0x54, 0x50, 0x5c, 0x83, 0x9f, 0x8e, 0xde, 0x5c, 0x64, 0xcc, 0x4c,
	0x8b, 0x49, 0x9c, 0xca, 0x59, 0x32, 0x2b, 0xca, 0xeb, 0xc5, 0x29, 0x93, 0xd5, 0x32, 0xc9, 0xdf,
	0x66, 0xc9, 0x37, 0x7e, 0x59, 0x93, 0x5d, 0x3b, 0xcf, 0xff, 0xff, 0x12, 0x00, 0x00, 0xff, 0xff,
	0x92, 0xe2, 0xc8, 0x69, 0xdc, 0x06, 0x00, 0x00,
}
