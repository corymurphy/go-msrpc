package ivdsvdprovider

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

// IVdsVdProvider server interface.
type VDiskProviderServer interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// The QueryVDisks method returns a list of virtual disks that are managed by the provider.
	//
	// Return Values: The method MUST return zero or a nonerror HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	QueryVDisks(context.Context, *QueryVDisksRequest) (*QueryVDisksResponse, error)

	// The CreateVDisk method defines a new virtual disk. This method creates a virtual
	// disk file to be used as the backing store for the virtual disk.
	//
	// Return Values: The method MUST return zero to indicate success, or return an implementation-specific
	// nonzero error code to indicate failure. For the HRESULT values predefined by the
	// Virtual Disk Service Remote Protocol, see section 2.2.3.
	CreateVDisk(context.Context, *CreateVDiskRequest) (*CreateVDiskResponse, error)

	// The AddVDisk method creates a virtual disk object representing the specified virtual
	// disk and adds it to the list of virtual disks managed by the provider. This method
	// returns an IVdsVDisk (section 3.1.15.1) interface pointer to the specified virtual
	// disk object.
	//
	// Return Values: The method MUST return zero or a nonerror HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	AddVDisk(context.Context, *AddVDiskRequest) (*AddVDiskResponse, error)

	// The GetDiskFromVDisk method returns an IVdsDisk (section 3.1.12.1) interface pointer
	// for a virtual disk given an IVdsVDisk (section 3.1.15.1) interface pointer.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetDiskFromVDisk(context.Context, *GetDiskFromVDiskRequest) (*GetDiskFromVDiskResponse, error)

	// The GetVDiskFromDisk method returns an IVdsVDisk (section 3.1.15.1) interface pointer
	// for the virtual disk given an IVdsDisk (section 3.1.12.1) interface pointer.
	//
	// Return Values: The method MUST return zero or a non-error HRESULT (as specified in
	// [MS-ERREF]) to indicate success, or return an implementation-specific nonzero error
	// code to indicate failure. For the HRESULT values predefined by the Virtual Disk Service
	// Remote Protocol, see section 2.2.3.
	GetVDiskFromDisk(context.Context, *GetVDiskFromDiskRequest) (*GetVDiskFromDiskResponse, error)
}

func RegisterVDiskProviderServer(conn dcerpc.Conn, o VDiskProviderServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewVDiskProviderServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(VDiskProviderSyntaxV0_0))...)
}

func NewVDiskProviderServerHandle(o VDiskProviderServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return VDiskProviderServerHandle(ctx, o, opNum, r)
	}
}

func VDiskProviderServerHandle(ctx context.Context, o VDiskProviderServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // QueryVDisks
		in := &QueryVDisksRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryVDisks(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // CreateVDisk
		in := &CreateVDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateVDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // AddVDisk
		in := &AddVDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddVDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // GetDiskFromVDisk
		in := &GetDiskFromVDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetDiskFromVDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // GetVDiskFromDisk
		in := &GetVDiskFromDiskRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetVDiskFromDisk(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
