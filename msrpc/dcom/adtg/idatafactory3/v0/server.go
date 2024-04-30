package idatafactory3

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	idatafactory2 "github.com/oiweiwei/go-msrpc/msrpc/dcom/adtg/idatafactory2/v0"
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
	_ = idatafactory2.GoPackage
)

// IDataFactory3 server interface.
type DataFactory3Server interface {

	// IDataFactory2 base class.
	idatafactory2.DataFactory2Server

	// Execute operation.
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)

	// Synchronize operation.
	Synchronize(context.Context, *SynchronizeRequest) (*SynchronizeResponse, error)
}

func RegisterDataFactory3Server(conn dcerpc.Conn, o DataFactory3Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDataFactory3ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(DataFactory3SyntaxV0_0))...)
}

func NewDataFactory3ServerHandle(o DataFactory3Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return DataFactory3ServerHandle(ctx, o, opNum, r)
	}
}

func DataFactory3ServerHandle(ctx context.Context, o DataFactory3Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 9 {
		// IDataFactory2 base method.
		return idatafactory2.DataFactory2ServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 9: // Execute
		in := &ExecuteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Execute(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // Synchronize
		in := &SynchronizeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Synchronize(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
