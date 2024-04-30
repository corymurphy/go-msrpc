package iwamadmin

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

// IWamAdmin server interface.
type WAMAdminServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The AppCreate method creates a new application at the specified metabase path.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// The opnum field value for this method is 3.
	AppCreate(context.Context, *AppCreateRequest) (*AppCreateResponse, error)

	// The AppDelete method deletes the application from the specified metabase path.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.         |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 4.
	AppDelete(context.Context, *AppDeleteRequest) (*AppDeleteResponse, error)

	// The AppUnLoad method shuts down the specified application.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+-------------------+--------------------------+
	//	|      RETURN       |                          |
	//	|    VALUE/CODE     |       DESCRIPTION        |
	//	|                   |                          |
	//	+-------------------+--------------------------+
	//	+-------------------+--------------------------+
	//	| 0x00000000 S_OK   | The call was successful. |
	//	+-------------------+--------------------------+
	//
	// The opnum field value for this method is 5.
	AppUnload(context.Context, *AppUnloadRequest) (*AppUnloadResponse, error)

	// The AppGetStatus method retrieves the status of the application defined at the specified
	// metabase path.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 6.
	AppGetStatus(context.Context, *AppGetStatusRequest) (*AppGetStatusResponse, error)

	// The AppDeleteRecoverable method deletes the application from the specified metabase
	// path and saves external state needed to recreate the application if it is recovered.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070057 E_INVALIDARG         | One or more arguments are invalid.         |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 7.
	AppDeleteRecoverable(context.Context, *AppDeleteRecoverableRequest) (*AppDeleteRecoverableResponse, error)

	// The AppRecover method recreates an application that was deleted by the AppDeleteRecoverable
	// method.
	//
	// Return Values:  A signed 32-bit value that indicates return status. If the method
	// returns a negative value, it failed. If the 12-bit facility code (bits 16–27) is
	// set to 0x007, the value contains a Win32 error code in the lower 16 bits. Zero or
	// positive values indicate success, with the lower 16 bits in positive nonzero values
	// containing warnings or flags defined in the method implementation. For more information
	// about Win32 error codes and HRESULT values, see [MS-ERREF].
	//
	//	+---------------------------------+--------------------------------------------+
	//	|             RETURN              |                                            |
	//	|           VALUE/CODE            |                DESCRIPTION                 |
	//	|                                 |                                            |
	//	+---------------------------------+--------------------------------------------+
	//	+---------------------------------+--------------------------------------------+
	//	| 0x00000000 S_OK                 | The call was successful.                   |
	//	+---------------------------------+--------------------------------------------+
	//	| 0x80070003 ERROR_PATH_NOT_FOUND | The system cannot find the path specified. |
	//	+---------------------------------+--------------------------------------------+
	//
	// The opnum field value for this method is 8.
	AppRecover(context.Context, *AppRecoverRequest) (*AppRecoverResponse, error)
}

func RegisterWAMAdminServer(conn dcerpc.Conn, o WAMAdminServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewWAMAdminServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(WAMAdminSyntaxV0_0))...)
}

func NewWAMAdminServerHandle(o WAMAdminServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return WAMAdminServerHandle(ctx, o, opNum, r)
	}
}

func WAMAdminServerHandle(ctx context.Context, o WAMAdminServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // AppCreate
		in := &AppCreateRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AppCreate(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // AppDelete
		in := &AppDeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AppDelete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // AppUnLoad
		in := &AppUnloadRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AppUnload(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // AppGetStatus
		in := &AppGetStatusRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AppGetStatus(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // AppDeleteRecoverable
		in := &AppDeleteRecoverableRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AppDeleteRecoverable(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // AppRecover
		in := &AppRecoverRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AppRecover(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
