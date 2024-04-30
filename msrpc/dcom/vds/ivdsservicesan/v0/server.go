package ivdsservicesan

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

// IVdsServiceSAN server interface.
type ServiceSANServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The GetSANPolicy method returns the current SAN policy setting.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT, as specified in
	// [MS-ERREF], to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetSANPolicy(context.Context, *GetSANPolicyRequest) (*GetSANPolicyResponse, error)

	// The SetSANPolicy method sets the SAN policy value.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT, as specified in
	// [MS-ERREF], to indicate success or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	SetSANPolicy(context.Context, *SetSANPolicyRequest) (*SetSANPolicyResponse, error)
}

func RegisterServiceSANServer(conn dcerpc.Conn, o ServiceSANServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewServiceSANServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(ServiceSANSyntaxV0_0))...)
}

func NewServiceSANServerHandle(o ServiceSANServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return ServiceSANServerHandle(ctx, o, opNum, r)
	}
}

func ServiceSANServerHandle(ctx context.Context, o ServiceSANServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // GetSANPolicy
		in := &GetSANPolicyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSANPolicy(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // SetSANPolicy
		in := &SetSANPolicyRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSANPolicy(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
