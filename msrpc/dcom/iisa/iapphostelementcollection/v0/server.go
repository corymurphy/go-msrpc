package iapphostelementcollection

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	iunknown "github.com/oiweiwei/go-msrpc/msrpc/dcom/iunknown/v0"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

var (
	_ = context.Background
	_ = fmt.Errorf
	_ = utf16.Encode
	_ = strings.TrimPrefix
	_ = ndr.ZeroString
	_ = (*uuid.UUID)(nil)
	_ = (*dcerpc.SyntaxID)(nil)
	_ = (*errors.Error)(nil)
	_ = iunknown.GoPackage
)

// IAppHostElementCollection server interface.
type AppHostElementCollectionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)

	// AddElement operation.
	AddElement(context.Context, *AddElementRequest) (*AddElementResponse, error)

	// DeleteElement operation.
	DeleteElement(context.Context, *DeleteElementRequest) (*DeleteElementResponse, error)

	// Clear operation.
	Clear(context.Context, *ClearRequest) (*ClearResponse, error)

	// CreateNewElement operation.
	CreateNewElement(context.Context, *CreateNewElementRequest) (*CreateNewElementResponse, error)

	// Schema operation.
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)
}

func RegisterAppHostElementCollectionServer(conn dcerpc.Conn, o AppHostElementCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostElementCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostElementCollectionSyntaxV0_0))...)
}

func NewAppHostElementCollectionServerHandle(o AppHostElementCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostElementCollectionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostElementCollectionServerHandle(ctx context.Context, o AppHostElementCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // Count
		in := &GetCountRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetCount(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // Item
		in := &GetItemRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetItem(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // AddElement
		in := &AddElementRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddElement(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // DeleteElement
		in := &DeleteElementRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteElement(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // Clear
		in := &ClearRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Clear(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // CreateNewElement
		in := &CreateNewElementRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateNewElement(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // Schema
		in := &GetSchemaRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSchema(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
