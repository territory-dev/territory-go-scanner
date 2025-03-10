// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.1
// source: uim.proto

package index

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenType int32

const (
	TokenType_WS          TokenType = 0
	TokenType_Keyword     TokenType = 1
	TokenType_Identifier  TokenType = 2
	TokenType_Punctuation TokenType = 3
	TokenType_Comment     TokenType = 4
	TokenType_Literal     TokenType = 5
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "WS",
		1: "Keyword",
		2: "Identifier",
		3: "Punctuation",
		4: "Comment",
		5: "Literal",
	}
	TokenType_value = map[string]int32{
		"WS":          0,
		"Keyword":     1,
		"Identifier":  2,
		"Punctuation": 3,
		"Comment":     4,
		"Literal":     5,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_uim_proto_enumTypes[0].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_uim_proto_enumTypes[0]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_uim_proto_rawDescGZIP(), []int{0}
}

type NodeKind int32

const (
	NodeKind_Definition NodeKind = 0
	NodeKind_Directory  NodeKind = 1
	NodeKind_File       NodeKind = 2
	NodeKind_Structure  NodeKind = 3
	NodeKind_SourceFile NodeKind = 4
	NodeKind_Class      NodeKind = 5
)

// Enum value maps for NodeKind.
var (
	NodeKind_name = map[int32]string{
		0: "Definition",
		1: "Directory",
		2: "File",
		3: "Structure",
		4: "SourceFile",
		5: "Class",
	}
	NodeKind_value = map[string]int32{
		"Definition": 0,
		"Directory":  1,
		"File":       2,
		"Structure":  3,
		"SourceFile": 4,
		"Class":      5,
	}
)

func (x NodeKind) Enum() *NodeKind {
	p := new(NodeKind)
	*p = x
	return p
}

func (x NodeKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeKind) Descriptor() protoreflect.EnumDescriptor {
	return file_uim_proto_enumTypes[1].Descriptor()
}

func (NodeKind) Type() protoreflect.EnumType {
	return &file_uim_proto_enumTypes[1]
}

func (x NodeKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeKind.Descriptor instead.
func (NodeKind) EnumDescriptor() ([]byte, []int) {
	return file_uim_proto_rawDescGZIP(), []int{1}
}

type IndexItemKind int32

const (
	IndexItemKind_IISymbol    IndexItemKind = 0
	IndexItemKind_IIDirectory IndexItemKind = 1
	IndexItemKind_IIFile      IndexItemKind = 2
	IndexItemKind_IIMacro     IndexItemKind = 3
)

// Enum value maps for IndexItemKind.
var (
	IndexItemKind_name = map[int32]string{
		0: "IISymbol",
		1: "IIDirectory",
		2: "IIFile",
		3: "IIMacro",
	}
	IndexItemKind_value = map[string]int32{
		"IISymbol":    0,
		"IIDirectory": 1,
		"IIFile":      2,
		"IIMacro":     3,
	}
)

func (x IndexItemKind) Enum() *IndexItemKind {
	p := new(IndexItemKind)
	*p = x
	return p
}

func (x IndexItemKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndexItemKind) Descriptor() protoreflect.EnumDescriptor {
	return file_uim_proto_enumTypes[2].Descriptor()
}

func (IndexItemKind) Type() protoreflect.EnumType {
	return &file_uim_proto_enumTypes[2]
}

func (x IndexItemKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndexItemKind.Descriptor instead.
func (IndexItemKind) EnumDescriptor() ([]byte, []int) {
	return file_uim_proto_rawDescGZIP(), []int{2}
}

type BlobSliceLoc struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BlobId        uint64                 `protobuf:"varint,1,opt,name=blob_id,json=blobId,proto3" json:"blob_id,omitempty"`
	StartOffset   uint64                 `protobuf:"varint,2,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	EndOffset     uint64                 `protobuf:"varint,3,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BlobSliceLoc) Reset() {
	*x = BlobSliceLoc{}
	mi := &file_uim_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BlobSliceLoc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobSliceLoc) ProtoMessage() {}

func (x *BlobSliceLoc) ProtoReflect() protoreflect.Message {
	mi := &file_uim_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobSliceLoc.ProtoReflect.Descriptor instead.
func (*BlobSliceLoc) Descriptor() ([]byte, []int) {
	return file_uim_proto_rawDescGZIP(), []int{0}
}

func (x *BlobSliceLoc) GetBlobId() uint64 {
	if x != nil {
		return x.BlobId
	}
	return 0
}

func (x *BlobSliceLoc) GetStartOffset() uint64 {
	if x != nil {
		return x.StartOffset
	}
	return 0
}

func (x *BlobSliceLoc) GetEndOffset() uint64 {
	if x != nil {
		return x.EndOffset
	}
	return 0
}

type NodeIdWithOffsetHref struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NodeId        uint64                 `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Offset        uint32                 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NodeIdWithOffsetHref) Reset() {
	*x = NodeIdWithOffsetHref{}
	mi := &file_uim_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NodeIdWithOffsetHref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeIdWithOffsetHref) ProtoMessage() {}

func (x *NodeIdWithOffsetHref) ProtoReflect() protoreflect.Message {
	mi := &file_uim_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeIdWithOffsetHref.ProtoReflect.Descriptor instead.
func (*NodeIdWithOffsetHref) Descriptor() ([]byte, []int) {
	return file_uim_proto_rawDescGZIP(), []int{1}
}

func (x *NodeIdWithOffsetHref) GetNodeId() uint64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *NodeIdWithOffsetHref) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type UniHref struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Path          string                 `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Offset        uint32                 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UniHref) Reset() {
	*x = UniHref{}
	mi := &file_uim_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UniHref) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UniHref) ProtoMessage() {}

func (x *UniHref) ProtoReflect() protoreflect.Message {
	mi := &file_uim_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UniHref.ProtoReflect.Descriptor instead.
func (*UniHref) Descriptor() ([]byte, []int) {
	return file_uim_proto_rawDescGZIP(), []int{2}
}

func (x *UniHref) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UniHref) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Token struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	Offset uint32                 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Type   TokenType              `protobuf:"varint,2,opt,name=type,proto3,enum=territory.index.TokenType" json:"type,omitempty"`
	// Types that are valid to be assigned to Href:
	//
	//	*Token_DirectNodeLink
	//	*Token_NodeIdRef
	//	*Token_SymIdRef
	//	*Token_UniHref
	//	*Token_NodeIdWithOffsetRef
	Href isToken_Href `protobuf_oneof:"href"`
	// Deprecated: Marked as deprecated in uim.proto.
	References    *uint64   `protobuf:"varint,4,opt,name=references,proto3,oneof" json:"references,omitempty"`
	HasReferences bool      `protobuf:"varint,5,opt,name=has_references,json=hasReferences,proto3" json:"has_references,omitempty"`
	SymId         *uint64   `protobuf:"varint,8,opt,name=sym_id,json=symId,proto3,oneof" json:"sym_id,omitempty"`
	RealOffset    *uint32   `protobuf:"varint,9,opt,name=real_offset,json=realOffset,proto3,oneof" json:"real_offset,omitempty"`
	RealLine      *uint32   `protobuf:"varint,10,opt,name=real_line,json=realLine,proto3,oneof" json:"real_line,omitempty"`
	UimLocation   *Location `protobuf:"bytes,12,opt,name=uim_location,json=uimLocation,proto3,oneof" json:"uim_location,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Token) Reset() {
	*x = Token{}
	mi := &file_uim_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_uim_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_uim_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *Token) GetType() TokenType {
	if x != nil {
		return x.Type
	}
	return TokenType_WS
}

func (x *Token) GetHref() isToken_Href {
	if x != nil {
		return x.Href
	}
	return nil
}

func (x *Token) GetDirectNodeLink() uint64 {
	if x != nil {
		if x, ok := x.Href.(*Token_DirectNodeLink); ok {
			return x.DirectNodeLink
		}
	}
	return 0
}

func (x *Token) GetNodeIdRef() uint64 {
	if x != nil {
		if x, ok := x.Href.(*Token_NodeIdRef); ok {
			return x.NodeIdRef
		}
	}
	return 0
}

func (x *Token) GetSymIdRef() uint64 {
	if x != nil {
		if x, ok := x.Href.(*Token_SymIdRef); ok {
			return x.SymIdRef
		}
	}
	return 0
}

func (x *Token) GetUniHref() *UniHref {
	if x != nil {
		if x, ok := x.Href.(*Token_UniHref); ok {
			return x.UniHref
		}
	}
	return nil
}

func (x *Token) GetNodeIdWithOffsetRef() *NodeIdWithOffsetHref {
	if x != nil {
		if x, ok := x.Href.(*Token_NodeIdWithOffsetRef); ok {
			return x.NodeIdWithOffsetRef
		}
	}
	return nil
}

// Deprecated: Marked as deprecated in uim.proto.
func (x *Token) GetReferences() uint64 {
	if x != nil && x.References != nil {
		return *x.References
	}
	return 0
}

func (x *Token) GetHasReferences() bool {
	if x != nil {
		return x.HasReferences
	}
	return false
}

func (x *Token) GetSymId() uint64 {
	if x != nil && x.SymId != nil {
		return *x.SymId
	}
	return 0
}

func (x *Token) GetRealOffset() uint32 {
	if x != nil && x.RealOffset != nil {
		return *x.RealOffset
	}
	return 0
}

func (x *Token) GetRealLine() uint32 {
	if x != nil && x.RealLine != nil {
		return *x.RealLine
	}
	return 0
}

func (x *Token) GetUimLocation() *Location {
	if x != nil {
		return x.UimLocation
	}
	return nil
}

type isToken_Href interface {
	isToken_Href()
}

type Token_DirectNodeLink struct {
	DirectNodeLink uint64 `protobuf:"varint,3,opt,name=direct_node_link,json=directNodeLink,proto3,oneof"`
}

type Token_NodeIdRef struct {
	NodeIdRef uint64 `protobuf:"varint,6,opt,name=node_id_ref,json=nodeIdRef,proto3,oneof"`
}

type Token_SymIdRef struct {
	SymIdRef uint64 `protobuf:"varint,7,opt,name=sym_id_ref,json=symIdRef,proto3,oneof"`
}

type Token_UniHref struct {
	UniHref *UniHref `protobuf:"bytes,11,opt,name=uni_href,json=uniHref,proto3,oneof"`
}

type Token_NodeIdWithOffsetRef struct {
	NodeIdWithOffsetRef *NodeIdWithOffsetHref `protobuf:"bytes,13,opt,name=node_id_with_offset_ref,json=nodeIdWithOffsetRef,proto3,oneof"`
}

func (*Token_DirectNodeLink) isToken_Href() {}

func (*Token_NodeIdRef) isToken_Href() {}

func (*Token_SymIdRef) isToken_Href() {}

func (*Token_UniHref) isToken_Href() {}

func (*Token_NodeIdWithOffsetRef) isToken_Href() {}

type Location struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Line          uint32                 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	Column        uint32                 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	Offset        uint32                 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Location) Reset() {
	*x = Location{}
	mi := &file_uim_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_uim_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_uim_proto_rawDescGZIP(), []int{4}
}

func (x *Location) GetLine() uint32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *Location) GetColumn() uint32 {
	if x != nil {
		return x.Column
	}
	return 0
}

func (x *Location) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Node struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Id                  uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind                NodeKind               `protobuf:"varint,2,opt,name=kind,proto3,enum=territory.index.NodeKind" json:"kind,omitempty"`
	Path                string                 `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Container           *uint64                `protobuf:"varint,4,opt,name=container,proto3,oneof" json:"container,omitempty"`
	Start               *Location              `protobuf:"bytes,5,opt,name=start,proto3" json:"start,omitempty"`
	Text                string                 `protobuf:"bytes,6,opt,name=text,proto3" json:"text,omitempty"`
	Tokens              []*Token               `protobuf:"bytes,7,rep,name=tokens,proto3" json:"tokens,omitempty"`
	MemberOf            *string                `protobuf:"bytes,8,opt,name=member_of,json=memberOf,proto3,oneof" json:"member_of,omitempty"`
	PathId              uint32                 `protobuf:"varint,9,opt,name=path_id,json=pathId,proto3" json:"path_id,omitempty"`
	UimReferenceContext *string                `protobuf:"bytes,10,opt,name=uim_reference_context,json=uimReferenceContext,proto3,oneof" json:"uim_reference_context,omitempty"`
	UimNestLevel        *uint32                `protobuf:"varint,11,opt,name=uim_nest_level,json=uimNestLevel,proto3,oneof" json:"uim_nest_level,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Node) Reset() {
	*x = Node{}
	mi := &file_uim_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_uim_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_uim_proto_rawDescGZIP(), []int{5}
}

func (x *Node) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetKind() NodeKind {
	if x != nil {
		return x.Kind
	}
	return NodeKind_Definition
}

func (x *Node) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Node) GetContainer() uint64 {
	if x != nil && x.Container != nil {
		return *x.Container
	}
	return 0
}

func (x *Node) GetStart() *Location {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Node) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Node) GetTokens() []*Token {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *Node) GetMemberOf() string {
	if x != nil && x.MemberOf != nil {
		return *x.MemberOf
	}
	return ""
}

func (x *Node) GetPathId() uint32 {
	if x != nil {
		return x.PathId
	}
	return 0
}

func (x *Node) GetUimReferenceContext() string {
	if x != nil && x.UimReferenceContext != nil {
		return *x.UimReferenceContext
	}
	return ""
}

func (x *Node) GetUimNestLevel() uint32 {
	if x != nil && x.UimNestLevel != nil {
		return *x.UimNestLevel
	}
	return 0
}

type IndexItem struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Key   string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are valid to be assigned to Href:
	//
	//	*IndexItem_DirectNodeLink
	//	*IndexItem_Floc
	//	*IndexItem_NodeId
	//	*IndexItem_UniHref
	Href          isIndexItem_Href `protobuf_oneof:"href"`
	Kind          IndexItemKind    `protobuf:"varint,3,opt,name=kind,proto3,enum=territory.index.IndexItemKind" json:"kind,omitempty"`
	Path          *string          `protobuf:"bytes,4,opt,name=path,proto3,oneof" json:"path,omitempty"`
	Type          *string          `protobuf:"bytes,5,opt,name=type,proto3,oneof" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IndexItem) Reset() {
	*x = IndexItem{}
	mi := &file_uim_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndexItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexItem) ProtoMessage() {}

func (x *IndexItem) ProtoReflect() protoreflect.Message {
	mi := &file_uim_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexItem.ProtoReflect.Descriptor instead.
func (*IndexItem) Descriptor() ([]byte, []int) {
	return file_uim_proto_rawDescGZIP(), []int{6}
}

func (x *IndexItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *IndexItem) GetHref() isIndexItem_Href {
	if x != nil {
		return x.Href
	}
	return nil
}

func (x *IndexItem) GetDirectNodeLink() uint64 {
	if x != nil {
		if x, ok := x.Href.(*IndexItem_DirectNodeLink); ok {
			return x.DirectNodeLink
		}
	}
	return 0
}

func (x *IndexItem) GetFloc() *BlobSliceLoc {
	if x != nil {
		if x, ok := x.Href.(*IndexItem_Floc); ok {
			return x.Floc
		}
	}
	return nil
}

func (x *IndexItem) GetNodeId() uint64 {
	if x != nil {
		if x, ok := x.Href.(*IndexItem_NodeId); ok {
			return x.NodeId
		}
	}
	return 0
}

func (x *IndexItem) GetUniHref() *UniHref {
	if x != nil {
		if x, ok := x.Href.(*IndexItem_UniHref); ok {
			return x.UniHref
		}
	}
	return nil
}

func (x *IndexItem) GetKind() IndexItemKind {
	if x != nil {
		return x.Kind
	}
	return IndexItemKind_IISymbol
}

func (x *IndexItem) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *IndexItem) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

type isIndexItem_Href interface {
	isIndexItem_Href()
}

type IndexItem_DirectNodeLink struct {
	DirectNodeLink uint64 `protobuf:"varint,2,opt,name=direct_node_link,json=directNodeLink,proto3,oneof"`
}

type IndexItem_Floc struct {
	Floc *BlobSliceLoc `protobuf:"bytes,6,opt,name=floc,proto3,oneof"`
}

type IndexItem_NodeId struct {
	NodeId uint64 `protobuf:"varint,7,opt,name=node_id,json=nodeId,proto3,oneof"`
}

type IndexItem_UniHref struct {
	UniHref *UniHref `protobuf:"bytes,8,opt,name=uni_href,json=uniHref,proto3,oneof"`
}

func (*IndexItem_DirectNodeLink) isIndexItem_Href() {}

func (*IndexItem_Floc) isIndexItem_Href() {}

func (*IndexItem_NodeId) isIndexItem_Href() {}

func (*IndexItem_UniHref) isIndexItem_Href() {}

var File_uim_proto protoreflect.FileDescriptor

var file_uim_proto_rawDesc = []byte{
	0x0a, 0x09, 0x75, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x65, 0x72,
	0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x69, 0x0a, 0x0c,
	0x42, 0x6c, 0x6f, 0x62, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x12, 0x17, 0x0a, 0x07,
	0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x62,
	0x6c, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x65, 0x6e,
	0x64, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x47, 0x0a, 0x14, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x48, 0x72, 0x65, 0x66, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x35, 0x0a, 0x07, 0x55, 0x6e, 0x69, 0x48, 0x72, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x9b, 0x05, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x5f, 0x72, 0x65, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x52, 0x65, 0x66, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79, 0x6d, 0x5f, 0x69,
	0x64, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x08, 0x73,
	0x79, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x66, 0x12, 0x35, 0x0a, 0x08, 0x75, 0x6e, 0x69, 0x5f, 0x68,
	0x72, 0x65, 0x66, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x72, 0x72,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x55, 0x6e, 0x69, 0x48,
	0x72, 0x65, 0x66, 0x48, 0x00, 0x52, 0x07, 0x75, 0x6e, 0x69, 0x48, 0x72, 0x65, 0x66, 0x12, 0x5d,
	0x0a, 0x17, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x48, 0x72, 0x65, 0x66, 0x48, 0x00, 0x52, 0x13, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x57, 0x69, 0x74, 0x68, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52, 0x65, 0x66, 0x12, 0x27, 0x0a,
	0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x02, 0x18, 0x01, 0x48, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x5f, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x68, 0x61, 0x73, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x06, 0x73, 0x79, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52,
	0x05, 0x73, 0x79, 0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x72, 0x65, 0x61,
	0x6c, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x03,
	0x52, 0x0a, 0x72, 0x65, 0x61, 0x6c, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x04, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x41, 0x0a, 0x0c, 0x75, 0x69, 0x6d, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x05, 0x52, 0x0b, 0x75, 0x69, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x68, 0x72, 0x65, 0x66, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x73, 0x79, 0x6d, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x72, 0x65, 0x61, 0x6c, 0x5f,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x65, 0x61, 0x6c, 0x5f,
	0x6c, 0x69, 0x6e, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x75, 0x69, 0x6d, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xd9, 0x03, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74,
	0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x72, 0x72,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x09, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x61, 0x74, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61,
	0x74, 0x68, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x15, 0x75, 0x69, 0x6d, 0x5f, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x13, 0x75, 0x69, 0x6d, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a,
	0x0e, 0x75, 0x69, 0x6d, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x03, 0x52, 0x0c, 0x75, 0x69, 0x6d, 0x4e, 0x65, 0x73, 0x74,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x6f, 0x66, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x75, 0x69, 0x6d, 0x5f, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x42, 0x11,
	0x0a, 0x0f, 0x5f, 0x75, 0x69, 0x6d, 0x5f, 0x6e, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0xd0, 0x02, 0x0a, 0x09, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2a, 0x0a, 0x10, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0e, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x33, 0x0a,
	0x04, 0x66, 0x6c, 0x6f, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65,
	0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x42, 0x6c,
	0x6f, 0x62, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x48, 0x00, 0x52, 0x04, 0x66, 0x6c,
	0x6f, 0x63, 0x12, 0x19, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a,
	0x08, 0x75, 0x6e, 0x69, 0x5f, 0x68, 0x72, 0x65, 0x66, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x55, 0x6e, 0x69, 0x48, 0x72, 0x65, 0x66, 0x48, 0x00, 0x52, 0x07, 0x75, 0x6e, 0x69,
	0x48, 0x72, 0x65, 0x66, 0x12, 0x32, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x4b, 0x69,
	0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x68, 0x72,
	0x65, 0x66, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2a, 0x5b, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x06, 0x0a, 0x02, 0x57, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x75, 0x6e, 0x63, 0x74, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x10,
	0x05, 0x2a, 0x5d, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0e, 0x0a,
	0x0a, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x46, 0x69, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x10, 0x05,
	0x2a, 0x47, 0x0a, 0x0d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x74, 0x65, 0x6d, 0x4b, 0x69, 0x6e,
	0x64, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x49, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x49, 0x49, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x49, 0x49, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x49, 0x4d, 0x61, 0x63, 0x72, 0x6f, 0x10, 0x03, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x74, 0x65, 0x72, 0x72, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x3b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uim_proto_rawDescOnce sync.Once
	file_uim_proto_rawDescData = file_uim_proto_rawDesc
)

func file_uim_proto_rawDescGZIP() []byte {
	file_uim_proto_rawDescOnce.Do(func() {
		file_uim_proto_rawDescData = protoimpl.X.CompressGZIP(file_uim_proto_rawDescData)
	})
	return file_uim_proto_rawDescData
}

var file_uim_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_uim_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_uim_proto_goTypes = []any{
	(TokenType)(0),               // 0: territory.index.TokenType
	(NodeKind)(0),                // 1: territory.index.NodeKind
	(IndexItemKind)(0),           // 2: territory.index.IndexItemKind
	(*BlobSliceLoc)(nil),         // 3: territory.index.BlobSliceLoc
	(*NodeIdWithOffsetHref)(nil), // 4: territory.index.NodeIdWithOffsetHref
	(*UniHref)(nil),              // 5: territory.index.UniHref
	(*Token)(nil),                // 6: territory.index.Token
	(*Location)(nil),             // 7: territory.index.Location
	(*Node)(nil),                 // 8: territory.index.Node
	(*IndexItem)(nil),            // 9: territory.index.IndexItem
}
var file_uim_proto_depIdxs = []int32{
	0,  // 0: territory.index.Token.type:type_name -> territory.index.TokenType
	5,  // 1: territory.index.Token.uni_href:type_name -> territory.index.UniHref
	4,  // 2: territory.index.Token.node_id_with_offset_ref:type_name -> territory.index.NodeIdWithOffsetHref
	7,  // 3: territory.index.Token.uim_location:type_name -> territory.index.Location
	1,  // 4: territory.index.Node.kind:type_name -> territory.index.NodeKind
	7,  // 5: territory.index.Node.start:type_name -> territory.index.Location
	6,  // 6: territory.index.Node.tokens:type_name -> territory.index.Token
	3,  // 7: territory.index.IndexItem.floc:type_name -> territory.index.BlobSliceLoc
	5,  // 8: territory.index.IndexItem.uni_href:type_name -> territory.index.UniHref
	2,  // 9: territory.index.IndexItem.kind:type_name -> territory.index.IndexItemKind
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_uim_proto_init() }
func file_uim_proto_init() {
	if File_uim_proto != nil {
		return
	}
	file_uim_proto_msgTypes[3].OneofWrappers = []any{
		(*Token_DirectNodeLink)(nil),
		(*Token_NodeIdRef)(nil),
		(*Token_SymIdRef)(nil),
		(*Token_UniHref)(nil),
		(*Token_NodeIdWithOffsetRef)(nil),
	}
	file_uim_proto_msgTypes[5].OneofWrappers = []any{}
	file_uim_proto_msgTypes[6].OneofWrappers = []any{
		(*IndexItem_DirectNodeLink)(nil),
		(*IndexItem_Floc)(nil),
		(*IndexItem_NodeId)(nil),
		(*IndexItem_UniHref)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_uim_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_uim_proto_goTypes,
		DependencyIndexes: file_uim_proto_depIdxs,
		EnumInfos:         file_uim_proto_enumTypes,
		MessageInfos:      file_uim_proto_msgTypes,
	}.Build()
	File_uim_proto = out.File
	file_uim_proto_rawDesc = nil
	file_uim_proto_goTypes = nil
	file_uim_proto_depIdxs = nil
}
