package nspi

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

// nspi server interface.
type NspiServer interface {

	// The NspiBind method initiates a session between a client and the NSPI server.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	Bind(context.Context, *BindRequest) (*BindResponse, error)

	// The NspiUnbind method destroys the context handle. No other action is taken.
	//
	// Return Values: The server returns a DWORD value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	Unbind(context.Context, *UnbindRequest) (*UnbindResponse, error)

	// The NspiUpdateStat method updates the STAT block representing position in a table
	// to reflect positioning changes requested by the client.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	UpdateStat(context.Context, *UpdateStatRequest) (*UpdateStatResponse, error)

	// The NspiQueryRows method returns to the client a number of rows from a specified
	// table. The server MUST return no more rows than the number specified in the input
	// parameter Count. Although the protocol places no further boundary or requirements
	// on the minimum number of rows the server returns, implementations SHOULD return as
	// many rows as possible subject to this maximum limit to improve usability of the NSPI
	// server for clients.
	//
	// Return Values:  The server returns a long value specifying the return status of
	// the method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	QueryRows(context.Context, *QueryRowsRequest) (*QueryRowsResponse, error)

	// The NspiSeekEntries method searches for and sets the logical position in a specific
	// table to the first entry greater than or equal to a specified value. Optionally,
	// it might also return information about rows in the table.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	SeekEntries(context.Context, *SeekEntriesRequest) (*SeekEntriesResponse, error)

	// The NspiGetMatches method returns an Explicit Table. The rows in the table are chosen
	// based on a two possible criteria: a restriction applied to an address book container
	// or the values of a property on a single object that hold references to other objects.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetMatches(context.Context, *GetMatchesRequest) (*GetMatchesResponse, error)

	// The NspiResortRestriction method applies a sort order to the objects in a restricted
	// address book container.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ResortRestriction(context.Context, *ResortRestrictionRequest) (*ResortRestrictionResponse, error)

	// The NspiDNToMId method maps a set of DN to a set of MId.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	DNToMID(context.Context, *DNToMIDRequest) (*DNToMIDResponse, error)

	// The NspiGetPropList method returns a list of all the properties that have values
	// on a specified object.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetPropertyList(context.Context, *GetPropertyListRequest) (*GetPropertyListResponse, error)

	// The NspiGetProps method returns an address book row containing a set of the properties
	// and values that exist on an object.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetProperties(context.Context, *GetPropertiesRequest) (*GetPropertiesResponse, error)

	// The NspiCompareMIds method compares the position in an address book container of
	// two objects identified by MId and returns the value of the comparison.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	CompareMIDs(context.Context, *CompareMIDsRequest) (*CompareMIDsResponse, error)

	// The NspiModProps method is used to modify the properties of an object in the address
	// book.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ModifyProperties(context.Context, *ModifyPropertiesRequest) (*ModifyPropertiesResponse, error)

	// The NspiGetSpecialTable method returns the rows of a special table to the client.
	// The special table can be an Address Book Hierarchy Table or an Address Creation Table.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetSpecialTable(context.Context, *GetSpecialTableRequest) (*GetSpecialTableResponse, error)

	// The NspiGetTemplateInfo method returns information about template objects in the
	// address book.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetTemplateInfo(context.Context, *GetTemplateInfoRequest) (*GetTemplateInfoResponse, error)

	// The NspiModLinkAtt method modifies the values of a specific property of a specific
	// row in the address book.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ModifyLinkAttribute(context.Context, *ModifyLinkAttributeRequest) (*ModifyLinkAttributeResponse, error)

	// Opnum15NotUsedOnWire operation.
	// Opnum15NotUsedOnWire

	// The NspiQueryColumns method returns a list of all the properties the NSPI server
	// is aware of. It returns this list as an array of proptags.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	QueryColumns(context.Context, *QueryColumnsRequest) (*QueryColumnsResponse, error)

	// The NspiGetNamesFromIDs method returns a list of property names for a set of proptags.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetNamesFromIDs(context.Context, *GetNamesFromIDsRequest) (*GetNamesFromIDsResponse, error)

	// The NspiGetIDsFromNames method returns a list of proptags for a set of property names.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	GetIDsFromNames(context.Context, *GetIDsFromNamesRequest) (*GetIDsFromNamesResponse, error)

	// The NspiResolveNames method takes a set of string values in an 8-bit character set
	// and performs ANR (as specified in 3.1.1.6) on those strings. The server reports the
	// MId that are the result of the ANR process. Certain property values are returned
	// for any valid MIds identified by the ANR process.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ResolveNames(context.Context, *ResolveNamesRequest) (*ResolveNamesResponse, error)

	// The NspiResolveNamesW method takes a set of string values in the Unicode character
	// set and performs ANR (as specified in 3.1.1.6) on those strings. The server reports
	// the MId that are the result of the ANR process. Certain property values are returned
	// for any valid MIds identified by the ANR process.
	//
	// Return Values: The server returns a long value specifying the return status of the
	// method.
	//
	// # Exceptions Thrown
	//
	// No exceptions are thrown beyond those thrown by the underlying RPC protocol, as specified
	// in [MS-RPCE].
	ResolveNamesW(context.Context, *ResolveNamesWRequest) (*ResolveNamesWResponse, error)
}

func RegisterNspiServer(conn dcerpc.Conn, o NspiServer, opts ...dcerpc.Option) {
	conn.RegisterServer(NewNspiServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(NspiSyntaxV56_0))...)
}

func NewNspiServerHandle(o NspiServer) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return NspiServerHandle(ctx, o, opNum, r)
	}
}

func NspiServerHandle(ctx context.Context, o NspiServer, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // NspiBind
		in := &BindRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Bind(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // NspiUnbind
		in := &UnbindRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.Unbind(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // NspiUpdateStat
		in := &UpdateStatRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.UpdateStat(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // NspiQueryRows
		in := &QueryRowsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryRows(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // NspiSeekEntries
		in := &SeekEntriesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SeekEntries(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // NspiGetMatches
		in := &GetMatchesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMatches(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // NspiResortRestriction
		in := &ResortRestrictionRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResortRestriction(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // NspiDNToMId
		in := &DNToMIDRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DNToMID(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // NspiGetPropList
		in := &GetPropertyListRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPropertyList(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // NspiGetProps
		in := &GetPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // NspiCompareMIds
		in := &CompareMIDsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CompareMIDs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // NspiModProps
		in := &ModifyPropertiesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ModifyProperties(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // NspiGetSpecialTable
		in := &GetSpecialTableRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSpecialTable(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // NspiGetTemplateInfo
		in := &GetTemplateInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetTemplateInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // NspiModLinkAtt
		in := &ModifyLinkAttributeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ModifyLinkAttribute(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // Opnum15NotUsedOnWire
		// Opnum15NotUsedOnWire
		return nil, nil
	case 16: // NspiQueryColumns
		in := &QueryColumnsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryColumns(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // NspiGetNamesFromIDs
		in := &GetNamesFromIDsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetNamesFromIDs(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // NspiGetIDsFromNames
		in := &GetIDsFromNamesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetIDsFromNames(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // NspiResolveNames
		in := &ResolveNamesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResolveNames(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // NspiResolveNamesW
		in := &ResolveNamesWRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ResolveNamesW(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
