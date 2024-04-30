package icatalogutils2

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

// ICatalogUtils2 server interface.
type CatalogUtils2Server interface {

	// IUnknown base class.
	iunknown.UnknownServer

	// This method is called by a client to copy one or more conglomerations from one partition
	// to another.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CopyConglomerations(context.Context, *CopyConglomerationsRequest) (*CopyConglomerationsResponse, error)

	// This method is called by a client to copy a component full configuration from one
	// conglomeration to another.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	CopyComponentConfiguration(context.Context, *CopyComponentConfigurationRequest) (*CopyComponentConfigurationResponse, error)

	// This method is called by a client to create an alias component full configuration,
	// a component full configuration of a virtual aliased component equivalent to the original
	// component except in CLSID and ProgID.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	AliasComponent(context.Context, *AliasComponentRequest) (*AliasComponentResponse, error)

	// This method is called by a client to move a component configuration from one conglomeration
	// to another.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	MoveComponentConfiguration(context.Context, *MoveComponentConfigurationRequest) (*MoveComponentConfigurationResponse, error)

	// This method is called by a client to get information about the event classes associated
	// with an IID that are configured in a specified partition.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetEventClassesForIid2(context.Context, *GetEventClassesForIid2Request) (*GetEventClassesForIid2Response, error)

	// The server method ICatalogUtils2::IsSafeToDelete (section 3.1.4.18.6) can be used
	// to determine if a file is safe to delete, but it is usually impossible for a server
	// to reliably make such a determination.
	//
	// A COMA client MUST NOT call the ICatalogUtils2::IsSafeToDelete method unless it receives
	// an explicit request from a client application to do so, and MUST return the results
	// of the call unaltered to the client application.
	IsSafeToDelete(context.Context, *IsSafeToDeleteRequest) (*IsSafeToDeleteResponse, error)

	// This method is called by a client to request that the server clear its local cache
	// of the entries in domain-controlled PartitionRoles (section 3.1.1.3.17), PartitionRoleMembers
	// (section 3.1.1.3.18), and PartitionUsers (section 3.1.1.3.16) tables, if the server
	// does such lookups with an active directory.
	//
	// This method has no parameters.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	FlushPartitionCache(context.Context, *FlushPartitionCacheRequest) (*FlushPartitionCacheResponse, error)

	// This method is called by a client to get an enumeration of software restriction policy
	// (see section 3.1.1.1.9) trust levels supported by the server.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	//
	// The server SHOULD, by some implementation-specific mechanism, attempt to translate
	// the names of the software restriction policy levels it supports to the language specified
	// by Locale, and SHOULD fall back to a default language if it cannot.
	//
	// The server then MUST attempt to set the values referenced by the out parameters as
	// follows: The server MUST attempt to set the value referenced by cLevels to the number
	// of software restriction policy levels the server supports, and the value referenced
	// by aSRPLevels to an array of SRPLevelInfo structures, each of which describes a software
	// restriction policy level, and fail the call if it cannot.
	EnumerateSRPLevels(context.Context, *EnumerateSRPLevelsRequest) (*EnumerateSRPLevelsResponse, error)

	// This method is called by a client to get a list of component full configurations
	// for a component.
	//
	// Return Values: This method MUST return S_OK (0x00000000) on success, and a failure
	// result, as specified in [MS-ERREF] section 2.1, on failure. All failure results MUST
	// be treated identically.
	GetComponentVersions(context.Context, *GetComponentVersionsRequest) (*GetComponentVersionsResponse, error)
}

func RegisterCatalogUtils2Server(conn dcerpc.Conn, o CatalogUtils2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewCatalogUtils2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(CatalogUtils2SyntaxV0_0))...)
}

func NewCatalogUtils2ServerHandle(o CatalogUtils2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return CatalogUtils2ServerHandle(ctx, o, opNum, r)
	}
}

func CatalogUtils2ServerHandle(ctx context.Context, o CatalogUtils2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	if opNum < 3 {
		// IUnknown base method.
		return iunknown.UnknownServerHandle(ctx, o, opNum, r)
	}
	switch opNum {
	case 3: // CopyConglomerations
		in := &CopyConglomerationsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CopyConglomerations(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // CopyComponentConfiguration
		in := &CopyComponentConfigurationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CopyComponentConfiguration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // AliasComponent
		in := &AliasComponentRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AliasComponent(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // MoveComponentConfiguration
		in := &MoveComponentConfigurationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.MoveComponentConfiguration(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // GetEventClassesForIID2
		in := &GetEventClassesForIid2Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetEventClassesForIid2(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // IsSafeToDelete
		in := &IsSafeToDeleteRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.IsSafeToDelete(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // FlushPartitionCache
		in := &FlushPartitionCacheRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FlushPartitionCache(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // EnumerateSRPLevels
		in := &EnumerateSRPLevelsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumerateSRPLevels(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // GetComponentVersions
		in := &GetComponentVersionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetComponentVersions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
