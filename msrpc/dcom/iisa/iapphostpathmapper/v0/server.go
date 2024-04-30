package iapphostpathmapper

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

// IAppHostPathMapper server interface.
type AppHostPathMapperServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// MapPath operation.
	MapPath(context.Context, *MapPathRequest) (*MapPathResponse, error)
}

func RegisterAppHostPathMapperServer(conn dcerpc.Conn, o AppHostPathMapperServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewAppHostPathMapperServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(AppHostPathMapperSyntaxV0_0))...)
}

func NewAppHostPathMapperServerHandle(o AppHostPathMapperServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return AppHostPathMapperServerHandle(ctx, o, opNum, r)
	}
}

func AppHostPathMapperServerHandle(ctx context.Context, o AppHostPathMapperServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // MapPath
		in := &MapPathRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MapPath(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
