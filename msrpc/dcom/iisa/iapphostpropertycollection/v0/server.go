package iapphostpropertycollection

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

// IAppHostPropertyCollection server interface.
type AppHostPropertyCollectionServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// Count operation.
	GetCount(context.Context, *GetCountRequest) (*GetCountResponse, error)

	// Item operation.
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
}

func RegisterAppHostPropertyCollectionServer(conn dcerpc.Conn, o AppHostPropertyCollectionServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostPropertyCollectionServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostPropertyCollectionSyntaxV0_0))...)
}

func NewAppHostPropertyCollectionServerHandle(o AppHostPropertyCollectionServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostPropertyCollectionServerHandle(ctx, o, opNum, r)
	}
}

func AppHostPropertyCollectionServerHandle(ctx context.Context, o AppHostPropertyCollectionServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
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
	}
	return nil, nil
}
