package windowsshutdown

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

// WindowsShutdown server interface.
type WindowsShutdownServer interface {

	// The WsdrInitiateShutdown method is used to initiate the shutdown of the remote computer.<14>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	InitiateShutdown(context.Context, *InitiateShutdownRequest) (*InitiateShutdownResponse, error)

	// The WsdrAbortShutdown method is used to terminate the shutdown of the remote computer
	// within the waiting period.<15>
	//
	// Return Values: The method returns ERROR_SUCCESS (0x00000000) on success; otherwise,
	// it returns a nonzero error code.
	AbortShutdown(context.Context, *AbortShutdownRequest) (*AbortShutdownResponse, error)
}

func RegisterWindowsShutdownServer(conn dcerpc.Conn, o WindowsShutdownServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWindowsShutdownServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WindowsShutdownSyntaxV1_0))...)
}

func NewWindowsShutdownServerHandle(o WindowsShutdownServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WindowsShutdownServerHandle(ctx, o, opNum, r)
	}
}

func WindowsShutdownServerHandle(ctx context.Context, o WindowsShutdownServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // WsdrInitiateShutdown
		in := &InitiateShutdownRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.InitiateShutdown(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // WsdrAbortShutdown
		in := &AbortShutdownRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AbortShutdown(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
