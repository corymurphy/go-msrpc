package dhcpsrv2

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

// dhcpsrv2 server interface.
type Dhcpsrv2Server interface {

	// The R_DhcpEnumSubnetClientsV5 method is used to retrieve all DHCPv4 clients serviced
	// from the specified IPv4 subnet. This method returns DHCPv4 clients from all IPv4
	// subnets if the subnet address specified zero. The caller of this function can free
	// the memory pointed to by the ClientInfo parameter and its Clients member by calling
	// the function midl_user_free (see section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA      | There are more elements available to enumerate.             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS  | There are no more elements left to enumerate.               |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 0.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	EnumSubnetClientsV5(context.Context, *EnumSubnetClientsV5Request) (*EnumSubnetClientsV5Response, error)

	// The R_DhcpSetMScopeInfo method creates/modifies an IPv4 multicast subnet on the MADCAP
	// server. The behavior of this method is dependent on parameter NewScope.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                         |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist.                                        |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR           | An error occurred while accessing the MADCAP server database.                    |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E4E ERROR_DHCP_SCOPE_NAME_TOO_LONG | The specified scope name is too long. The name is limited to a maximum of 256    |
	//	|                                           | characters.                                                                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E24 ERROR_DHCP_SUBNET_EXISTS       | The specified IPv4 multicast subnet already exists.                              |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E55 ERROR_DHCP_MSCOPE_EXISTS       | The multicast scope parameters are incorrect. Either the scope already exists or |
	//	|                                           | its properties are inconsistent with the properties of another existing scope.   |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 1.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetMScopeInfo(context.Context, *SetMScopeInfoRequest) (*SetMScopeInfoResponse, error)

	// The R_DhcpGetMScopeInfo method retrieves the information of the IPv4 multicast subnet
	// managed by the MADCAP server. The caller of this function can free the memory pointed
	// by MScopeInfo by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------+
	//	|                  RETURN                  |                                           |
	//	|                VALUE/CODE                |                DESCRIPTION                |
	//	|                                          |                                           |
	//	+------------------------------------------+-------------------------------------------+
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                  |
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist. |
	//	+------------------------------------------+-------------------------------------------+
	//
	// The opnum field value for this method is 2.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	GetMScopeInfo(context.Context, *GetMScopeInfoRequest) (*GetMScopeInfoResponse, error)

	// The R_DhcpEnumMScopes method enumerates IPv4 multicast subnet names configured on
	// the MADCAP server. The caller of this function can free the memory pointed to by
	// the MScopeTable parameter by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------+-----------------------------------------------+
	//	|             RETURN             |                                               |
	//	|           VALUE/CODE           |                  DESCRIPTION                  |
	//	|                                |                                               |
	//	+--------------------------------+-----------------------------------------------+
	//	+--------------------------------+-----------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                      |
	//	+--------------------------------+-----------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate. |
	//	+--------------------------------+-----------------------------------------------+
	//
	// The opnum field value for this method is 3.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	EnumMScopes(context.Context, *EnumMScopesRequest) (*EnumMScopesResponse, error)

	// The R_DhcpAddMScopeElement method adds an IPv4 multicast subnet element (IPv4 range
	// or IPv4 exclusion range) to the IPv4 multicast subnet in the MADCAP server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist.                                        |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR          | An error occurred while accessing the MADCAP server database.                    |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E37 ERROR_DHCP_INVALID_RANGE      | The specified multicast range either overlaps an existing range or is not valid. |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E56 ERROR_MSCOPE_RANGE_TOO_SMALL  | The multicast scope range MUST have at least 256 IPv4 addresses.                 |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E35 ERROR_DHCP_IPRANGE_EXITS      | The specified multicast range already exists.                                    |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 4.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	AddMScopeElement(context.Context, *AddMScopeElementRequest) (*AddMScopeElementResponse, error)

	// The R_DhcpEnumMScopeElements method enumerates the list of specific types of IPv4
	// multicast subnet elements (IPv4 range of IPv4 exclusion) from a specific IPv4 multicast
	// subnet. The caller of this function can free the memory pointed to by EnumElementInfo
	// and its member Elements by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------------+
	//	|                  RETURN                  |                                                 |
	//	|                VALUE/CODE                |                   DESCRIPTION                   |
	//	|                                          |                                                 |
	//	+------------------------------------------+-------------------------------------------------+
	//	+------------------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                        |
	//	+------------------------------------------+-------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA               | There are more elements available to enumerate. |
	//	+------------------------------------------+-------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS           | There are no more elements left to enumerate.   |
	//	+------------------------------------------+-------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist.       |
	//	+------------------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 5.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumMScopeElements(context.Context, *EnumMScopeElementsRequest) (*EnumMScopeElementsResponse, error)

	// The R_DhcpRemoveMScopeElement method removes an IPv4 multicast subnet element (IPv4
	// multicast range or IPv4 exclusion range) from the IPv4 multicast subnet defined on
	// the MADCAP server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                         |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E27 ERROR_DHCP_ELEMENT_CANT_REMOVE | The specified IPv4 multicast subnet element cannot be removed because at least   |
	//	|                                           | one multicast IPv4 address has been leased out to a MADCAP client. The starting  |
	//	|                                           | address of the specified Multicast exclusion range is not part of any multicast  |
	//	|                                           | exclusion range configured on the server. There is an error in deleting the      |
	//	|                                           | exclusion range from the database.                                               |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR           | An error occurred while accessing the MADCAP server database.                    |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E37 ERROR_DHCP_INVALID_RANGE       | The specified IPv4 range either overlaps an existing IPv4 range or is not valid. |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 6.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemoveMScopeElement(context.Context, *RemoveMScopeElementRequest) (*RemoveMScopeElementResponse, error)

	// The R_DhcpDeleteMScope method deletes the multicast subnet from the MADCAP server.
	// The ForceFlag defines the behavior of the operation when the subnet has served a
	// MADCAP client.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                         |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist.                                        |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E27 ERROR_DHCP_ELEMENT_CANT_REMOVE | The specified IPv4 multicast scope cannot be removed because at least one        |
	//	|                                           | multicast IPv4 address has been leased out to some MADCAP client.                |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR           | An error occurred while accessing the MADCAP server database.                    |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 7.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	DeleteMScope(context.Context, *DeleteMScopeRequest) (*DeleteMScopeResponse, error)

	// The R_DhcpScanMDatabase method can be used by DHCP servers to enumerate and/or fix
	// inconsistencies between the MADCAP lease records and the bitmask representation in
	// memory (section 3.1.1.4). The caller of this function can free the memory pointed
	// to by ScanList and its member ScanItems by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+---------------------------------------------------------------+
	//	|                  RETURN                  |                                                               |
	//	|                VALUE/CODE                |                          DESCRIPTION                          |
	//	|                                          |                                                               |
	//	+------------------------------------------+---------------------------------------------------------------+
	//	+------------------------------------------+---------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                      |
	//	+------------------------------------------+---------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist.                     |
	//	+------------------------------------------+---------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR          | An error occurred while accessing the MADCAP server database. |
	//	+------------------------------------------+---------------------------------------------------------------+
	//
	// The opnum field value for this method is 8.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ScanMDatabase(context.Context, *ScanMDatabaseRequest) (*ScanMDatabaseResponse, error)

	// The R_DhcpCreateMClientInfo method creates a multicast client record on the MADCAP
	// server's database. This also marks the specified client IP address as unavailable
	// (or distributed).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	|                RETURN                 |                                                                     |
	//	|              VALUE/CODE               |                             DESCRIPTION                             |
	//	|                                       |                                                                     |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | This call was successful.                                           |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000078 ERROR_CALL_NOT_IMPLEMENTED | The method is not implemented by this version of the MADCAP server. |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//
	// The opnum field value for this method is 9.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	CreateMClientInfo(context.Context, *CreateMClientInfoRequest) (*CreateMClientInfoResponse, error)

	// The R_DhcpSetMClientInfo method sets/modifies the specific MADCAP client lease record
	// on the MADCAP server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	|                RETURN                 |                                                                     |
	//	|              VALUE/CODE               |                             DESCRIPTION                             |
	//	|                                       |                                                                     |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | This call was successful.                                           |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//	| 0x00000078 ERROR_CALL_NOT_IMPLEMENTED | The method is not implemented by this version of the MADCAP server. |
	//	+---------------------------------------+---------------------------------------------------------------------+
	//
	// The opnum field value for this method is 10.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	SetMClientInfo(context.Context, *SetMClientInfoRequest) (*SetMClientInfoResponse, error)

	// The R_DhcpGetMClientInfo method retrieves the specified MADCAP client lease record
	// information from the MADCAP server. The caller of this function can free the memory
	// pointed to by ClientInfo by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+---------------------------------------------------------------+
	//	|             RETURN              |                                                               |
	//	|           VALUE/CODE            |                          DESCRIPTION                          |
	//	|                                 |                                                               |
	//	+---------------------------------+---------------------------------------------------------------+
	//	+---------------------------------+---------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                      |
	//	+---------------------------------+---------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the MADCAP server database. |
	//	+---------------------------------+---------------------------------------------------------------+
	//
	// The opnum field value for this method is 11.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	GetMClientInfo(context.Context, *GetMClientInfoRequest) (*GetMClientInfoResponse, error)

	// The R_DhcpDeleteMClientInfo method deletes the specified MADCAP client lease record
	// from the MADCAP server. It also frees up the MADCAP client IPv4 address for redistribution.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+---------------------------------------------------------------+
	//	|             RETURN              |                                                               |
	//	|           VALUE/CODE            |                          DESCRIPTION                          |
	//	|                                 |                                                               |
	//	+---------------------------------+---------------------------------------------------------------+
	//	+---------------------------------+---------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                      |
	//	+---------------------------------+---------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the MADCAP server database. |
	//	+---------------------------------+---------------------------------------------------------------+
	//
	// The opnum field value for this method is 12.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	DeleteMClientInfo(context.Context, *DeleteMClientInfoRequest) (*DeleteMClientInfoResponse, error)

	// The R_DhcpEnumMScopeClients method enumerates all MADCAP clients serviced from the
	// specified IPv4 multicast subnet. The caller of this function can free the memory
	// pointed to by the ClientInfo parameter and other client parameters by calling the
	// function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+---------------------------------------------------------------+
	//	|                  RETURN                  |                                                               |
	//	|                VALUE/CODE                |                          DESCRIPTION                          |
	//	|                                          |                                                               |
	//	+------------------------------------------+---------------------------------------------------------------+
	//	+------------------------------------------+---------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                      |
	//	+------------------------------------------+---------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA               | There are more elements available to enumerate.               |
	//	+------------------------------------------+---------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS           | There are no more elements left to enumerate.                 |
	//	+------------------------------------------+---------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist.                     |
	//	+------------------------------------------+---------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR          | An error occurred while accessing the MADCAP server database. |
	//	+------------------------------------------+---------------------------------------------------------------+
	//
	// The opnum field value for this method is 13.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	EnumMScopeClients(context.Context, *EnumMScopeClientsRequest) (*EnumMScopeClientsResponse, error)

	// The R_DhcpCreateOptionV5 method creates an option definition of a specific option
	// for a specific user class and vendor class at the default option level. The OptionId
	// specifies the identifier of the option. If the user class or vendor class is not
	// defined, the option definition is created for the default user class and vendor class.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	|                RETURN                 |                                                                         |
	//	|              VALUE/CODE               |                               DESCRIPTION                               |
	//	|                                       |                                                                         |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | The call was successful.                                                |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x00004E29 ERROR_DHCP_OPTION_EXITS    | The specified option definition already exists on DHCP server database. |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR       | An error occurred while accessing the DHCP server database.             |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND | The class name being used is unknown or incorrect.                      |
	//	+---------------------------------------+-------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 14.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	CreateOptionV5(context.Context, *CreateOptionV5Request) (*CreateOptionV5Response, error)

	// The R_DhcpSetOptionInfoV5 method modifies the option definition of a specific option
	// for a specific user class and vendor class at the default level. If the user class
	// or vendor class is not defined, the option definition is set or modified for the
	// default user class or vendor class. This is an extension of R_DhcpSetOptionInfo (section
	// 3.1.4.10), which sets the option definition for a default user and vendor class.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                             |
	//	|                VALUE/CODE                |                                 DESCRIPTION                                 |
	//	|                                          |                                                                             |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                                    |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT | The specified option definition does not exist on the DHCP server database. |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND    | The class name being used is unknown or incorrect.                          |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 15.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetOptionInfoV5(context.Context, *SetOptionInfoV5Request) (*SetOptionInfoV5Response, error)

	// The R_DhcpGetOptionInfoV5 method retrieves the option definition of a specific option
	// for a specific user class and vendor class at the default option level. If the user
	// class or vendor class is not defined, the option definition is retrieved for the
	// default user class or vendor class. This is an extension method of R_DhcpGetOptionInfo
	// (section 3.1.4.11), which retrieves the option definition of a specific option for
	// the default user and vendor class. The caller of this function can free the memory
	// pointed to by OptionInfo by calling the function midl_user_free (see section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                         |
	//	|                VALUE/CODE                |                               DESCRIPTION                               |
	//	|                                          |                                                                         |
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                                |
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT | The specified option definition does not exist on DHCP server database. |
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND    | The class name being used is unknown or incorrect.                      |
	//	+------------------------------------------+-------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 16.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetOptionInfoV5(context.Context, *GetOptionInfoV5Request) (*GetOptionInfoV5Response, error)

	// The R_DhcpEnumOptionsV5 method enumerates the option definitions for a specific user
	// class and vendor class for the default option level. If the user class or the vendor
	// class is not defined, the option definitions are enumerated for the default user
	// class or vendor class. This method is an extension of the method in R_DhcpEnumOptions
	// (section 3.1.4.24), which enumerates the option definition for a default user and
	// vendor class. The caller of this function can free the memory pointed to by the Options
	// parameter by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------------+----------------------------------------------------+
	//	|                RETURN                 |                                                    |
	//	|              VALUE/CODE               |                    DESCRIPTION                     |
	//	|                                       |                                                    |
	//	+---------------------------------------+----------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | The call was successful.                           |
	//	+---------------------------------------+----------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA            | There are more elements available to enumerate.    |
	//	+---------------------------------------+----------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS        | There are no more elements left to enumerate.      |
	//	+---------------------------------------+----------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND | The class name being used is unknown or incorrect. |
	//	+---------------------------------------+----------------------------------------------------+
	//
	// The opnum field value for this method is 17.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	EnumOptionsV5(context.Context, *EnumOptionsV5Request) (*EnumOptionsV5Response, error)

	// The R_DhcpRemoveOptionV5 method removes the option definition of a specific option
	// for a specific user class and vendor class at the default option level. If the user
	// class or the vendor class is not specified, the option definition is removed from
	// the default user class or vendor class. The OptionID specifies the identifier of
	// the option definition.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                             |
	//	|                VALUE/CODE                |                                 DESCRIPTION                                 |
	//	|                                          |                                                                             |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                                    |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT | The specified option definition does not exist on the DHCP server database. |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND    | The class name being used is unknown or incorrect.                          |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 18.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemoveOptionV5(context.Context, *RemoveOptionV5Request) (*RemoveOptionV5Response, error)

	// The R_DhcpSetOptionValueV5 method creates the option value, when called for the first
	// time. Otherwise, it modifies the option value of a specific option on the DHCPv4
	// server for a specific user class and vendor class. ScopeInfo defines the scope on
	// which this option value is set. If the user class or vendor class is not provided,
	// a default user class or vendor class is taken.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                             |
	//	|                VALUE/CODE                 |                                 DESCRIPTION                                 |
	//	|                                           |                                                                             |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                    |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist on the DHCP server.                |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT  | The specified option definition does not exist on the DHCP server database. |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The specified DHCP client is not a reserved client.                         |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND     | The class name being used is unknown or incorrect.                          |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 19.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetOptionValueV5(context.Context, *SetOptionValueV5Request) (*SetOptionValueV5Response, error)

	// The R_DhcpSetOptionValuesV5 method creates the option value when called for the first
	// time, else it modifies it. It creates or modifies one or more options for a specific
	// user class and vendor class. If the user class or the vendor class is not specified,
	// the option values are set or modified for the default user class or vendor class.
	// ScopeInfo defines the scope on which this option value is modified.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                             |
	//	|                VALUE/CODE                 |                                 DESCRIPTION                                 |
	//	|                                           |                                                                             |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                    |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT  | The specified option definition does not exist on the DHCP server database. |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The specified DHCP client is not a reserved client.                         |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND     | The class name being used is unknown or incorrect.                          |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 20.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	SetOptionValuesV5(context.Context, *SetOptionValuesV5Request) (*SetOptionValuesV5Response, error)

	// The R_DhcpGetOptionValueV5 method retrieves the option value for a specific option
	// on the DHCPv4 server for a specific user class and vendor class. If the user class
	// or the vendor class is not specified, the option value is retrieved from the default
	// user class or vendor class. ScopeInfo defines the scope from which the option value
	// needs to be retrieved. The caller of this function can free the memory pointed to
	// by OptionValue by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                             |
	//	|                VALUE/CODE                 |                                 DESCRIPTION                                 |
	//	|                                           |                                                                             |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                    |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist on the DHCP server.                |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT  | The specified option definition does not exist on the DHCP server database. |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The specified DHCP client is not a reserved client.                         |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 21.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetOptionValueV5(context.Context, *GetOptionValueV5Request) (*GetOptionValueV5Response, error)

	// The R_DhcpEnumOptionValuesV5 method enumerates all the option values for the specific
	// user class and vendor class at a specified scope defined by ScopeInfo. If the user
	// class or the vendor class is not specified, the option values are retrieved from
	// the default user class or vendor class. The caller of this function can free the
	// memory pointed to by OptionValues and the Values member of OptionValues by calling
	// the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	|                  RETURN                   |                                                              |
	//	|                VALUE/CODE                 |                         DESCRIPTION                          |
	//	|                                           |                                                              |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                     |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA                | There are more elements available to enumerate.              |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS            | There are no more elements left to enumerate.                |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist on the DHCP server. |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The specified DHCP client is not a reserved client.          |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND     | The class name being used is unknown or incorrect.           |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 22.
	EnumOptionValuesV5(context.Context, *EnumOptionValuesV5Request) (*EnumOptionValuesV5Response, error)

	// The R_DhcpRemoveOptionValueV5 method removes the option value for a specific option
	// on the DHCPv4 server for a specific user class and vendor class. If the user class
	// or the vendor class is not specified, the option value is removed from the default
	// user class or vendor class. ScopeInfo defines the scope on which this option value
	// is removed.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                             |
	//	|                VALUE/CODE                 |                                 DESCRIPTION                                 |
	//	|                                           |                                                                             |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                    |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist on the DHCP server.                |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT  | The specified option definition does not exist on the DHCP server database. |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The specified DHCP client is not a reserved client.                         |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND     | The class name being used is unknown or incorrect.                          |
	//	+-------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 23.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemoveOptionValueV5(context.Context, *RemoveOptionValueV5Request) (*RemoveOptionValueV5Response, error)

	// The R_DhcpCreateClass method creates a user class or a vendor class definition on
	// the DHCPv4 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------------------+------------------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                              |
	//	|                 VALUE/CODE                 |                                 DESCRIPTION                                  |
	//	|                                            |                                                                              |
	//	+--------------------------------------------+------------------------------------------------------------------------------+
	//	+--------------------------------------------+------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                   | The call was successful.                                                     |
	//	+--------------------------------------------+------------------------------------------------------------------------------+
	//	| 0x00004E4D ERROR_DHCP_CLASS_ALREADY_EXISTS | The class name is already in use or the class information is already in use. |
	//	+--------------------------------------------+------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 24.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	CreateClass(context.Context, *CreateClassRequest) (*CreateClassResponse, error)

	// The R_DhcpModifyClass method modifies the user class or vendor class definition for
	// the DHCP server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------------+----------------------------------------------------+
	//	|                RETURN                 |                                                    |
	//	|              VALUE/CODE               |                    DESCRIPTION                     |
	//	|                                       |                                                    |
	//	+---------------------------------------+----------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | The call was successful.                           |
	//	+---------------------------------------+----------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND | The class name being used is unknown or incorrect. |
	//	+---------------------------------------+----------------------------------------------------+
	//
	// The opnum field value for this method is 25.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	ModifyClass(context.Context, *ModifyClassRequest) (*ModifyClassResponse, error)

	// The R_DhcpDeleteClass method deletes the user class or vendor class definition from
	// the DHCP server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------------------+----------------------------------------------------+
	//	|                   RETURN                   |                                                    |
	//	|                 VALUE/CODE                 |                    DESCRIPTION                     |
	//	|                                            |                                                    |
	//	+--------------------------------------------+----------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                   | The call was successful.                           |
	//	+--------------------------------------------+----------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND      | The class name being used is unknown or incorrect. |
	//	+--------------------------------------------+----------------------------------------------------+
	//	| 0x00004E79 ERROR_DHCP_DELETE_BUILTIN_CLASS | This class cannot be deleted.                      |
	//	+--------------------------------------------+----------------------------------------------------+
	//
	// The opnum field value for this method is 26.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	DeleteClass(context.Context, *DeleteClassRequest) (*DeleteClassResponse, error)

	// The R_DhcpGetClassInfo method retrieves the user class or vendor class information
	// configured for the DHCP server. The caller of this function can free the memory pointed
	// to by FilledClassInfo by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------------+----------------------------------------------------+
	//	|                RETURN                 |                                                    |
	//	|              VALUE/CODE               |                    DESCRIPTION                     |
	//	|                                       |                                                    |
	//	+---------------------------------------+----------------------------------------------------+
	//	+---------------------------------------+----------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | The call was successful.                           |
	//	+---------------------------------------+----------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND | The class name being used is unknown or incorrect. |
	//	+---------------------------------------+----------------------------------------------------+
	//
	// The opnum field value for this method is 27.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetClassInfo(context.Context, *GetClassInfoRequest) (*GetClassInfoResponse, error)

	// The R_DhcpEnumClasses method enumerates user classes or vendor classes configured
	// for the DHCP server. The caller of this function can free the memory pointed to by
	// ClassInfoArray and Classes by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA     | There are more elements available to enumerate. |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate.   |
	//	+--------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 28.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumClasses(context.Context, *EnumClassesRequest) (*EnumClassesResponse, error)

	// The R_DhcpGetAllOptions method retrieves all default option definitions, as well
	// as specific user class and vendor class option definitions. The caller of this function
	// can free the memory pointed to by OptionStruct, NonVendorOptions and other Options
	// by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 29.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetAllOptions(context.Context, *GetAllOptionsRequest) (*GetAllOptionsResponse, error)

	// The R_DhcpGetAllOptionValues method retrieves the option values for all the options
	// configured at the DHCPv4 server from the specific scope for all user classes and
	// vendor classes. ScopeInfo defines the scope from which this option values are retrieved.
	// The caller of this method can free the memory pointed to by Values, its Options member,
	// and the members of each element in the Options array, by calling the function midl_user_free
	// (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	|                  RETURN                   |                                                              |
	//	|                VALUE/CODE                 |                         DESCRIPTION                          |
	//	|                                           |                                                              |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                     |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified IPv4 subnet does not exist on the DHCP server. |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The specified DHCP client is not a reserved client.          |
	//	+-------------------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 30.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetAllOptionValues(context.Context, *GetAllOptionValuesRequest) (*GetAllOptionValuesResponse, error)

	// The R_DhcpGetMCastMibInfo method retrieves the multicast counter values of the MADCAP
	// server. The caller of this function can free the memory pointed to by MibInfo by
	// calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 31.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetMCastMIBInfo(context.Context, *GetMCastMIBInfoRequest) (*GetMCastMIBInfoResponse, error)

	// The R_DhcpAuditLogSetParams method sets/modifies the DHCP server setting related
	// to the audit log.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 32.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	AuditLogSetParams(context.Context, *AuditLogSetParamsRequest) (*AuditLogSetParamsResponse, error)

	// The R_DhcpAuditLogGetParams method retrieves all audit log–related settings from
	// the DHCP server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 33.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	AuditLogGetParams(context.Context, *AuditLogGetParamsRequest) (*AuditLogGetParamsResponse, error)

	// The R_DhcpServerQueryAttribute method retrieves attribute information from the DHCP
	// server. The caller of this function can free the memory pointed to by pDhcpAttrib
	// by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 34.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ServerQueryAttribute(context.Context, *ServerQueryAttributeRequest) (*ServerQueryAttributeResponse, error)

	// The R_DhcpServerQueryAttributes method retrieves one or more attributes information
	// from the DHCP server. The caller of this function can free the memory pointed to
	// by pDhcpAttribArr and pDhcpAttribs by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 35.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ServerQueryAttributes(context.Context, *ServerQueryAttributesRequest) (*ServerQueryAttributesResponse, error)

	// The R_DhcpServerRedoAuthorization method attempts to determine whether the DHCP server
	// is authorized and restores the leasing operation if the server is not authorized.
	// The rogue detection mechanism is outlined in [MS-DHCPE] (section 3.3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 36.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ServerRedoAuthorization(context.Context, *ServerRedoAuthorizationRequest) (*ServerRedoAuthorizationResponse, error)

	// The R_DhcpAddSubnetElementV5 method adds an IPv4 subnet element to the specified
	// IPv4 subnet defined on the DHCPv4 server. The subnet elements can be IPv4 reservation
	// for DHCPv4 or BOOTP clients, IPv4 range, or the IPv4 exclusion range for DHCPv4 or
	// BOOTP clients.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                       RETURN                       |                                                                                  |
	//	|                     VALUE/CODE                     |                                   DESCRIPTION                                    |
	//	|                                                    |                                                                                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                           | The call was successful.                                                         |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT           | The specified IPv4 subnet does not exist on the DHCP server.                     |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR                    | An error occurred while accessing the DHCP server database.                      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT          | The specified DHCP client is not an IPv4-reserved client.                        |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E36 ERROR_DHCP_RESERVEDIP_EXITS             | The specified IPv4 address or hardware address is being used by another DHCP     |
	//	|                                                    | client.                                                                          |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E37 ERROR_DHCP_INVALID_RANGE                | The specified IPv4 range either overlaps an existing range or is not valid.      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E51 ERROR_DHCP_IPRANGE_CONV_ILLEGAL         | Conversion of a scope to a DHCP-only scope or to a BOOTP-only scope is not       |
	//	|                                                    | allowed when DHCP and BOOTP clients both exist in the scope. Manually delete     |
	//	|                                                    | either the DHCP or the BOOTP clients from the scope, as appropriate for the type |
	//	|                                                    | of scope being created.                                                          |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E90 ERROR_SCOPE_RANGE_POLICY_RANGE_CONFLICT | There is an IP range configured for a policy in this scope. This operation on    |
	//	|                                                    | the scope IP address range cannot be performed until the policy IP address range |
	//	|                                                    | is suitably modified.                                                            |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004EA1 ERROR_DHCP_FO_IPRANGE_TYPE_CONV_ILLEGAL | Conversion of a failover scope to a scope of type BOOTP or BOTH could not be     |
	//	|                                                    | performed. Failover is supported only for DHCP scopes.                           |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 37.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	AddSubnetElementV5(context.Context, *AddSubnetElementV5Request) (*AddSubnetElementV5Response, error)

	// The R_DhcpEnumSubnetElementsV5 method enumerates the list of a specific type of IPv4
	// subnet element from the specified IPv4 subnet. The caller of this function can free
	// the memory pointed to by EnumElementInfo and the Elements field of EnumElementInfo
	// by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+--------------------------------------------------------------+
	//	|                  RETURN                  |                                                              |
	//	|                VALUE/CODE                |                         DESCRIPTION                          |
	//	|                                          |                                                              |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                     |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA               | There are more elements available to enumerate.              |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS           | There are no more elements left to enumerate.                |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist on the DHCP server. |
	//	+------------------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 38.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnetElementsV5(context.Context, *EnumSubnetElementsV5Request) (*EnumSubnetElementsV5Response, error)

	// The R_DhcpRemoveSubnetElementV5 method removes an IPv4 subnet element from the specified
	// IPv4 subnet defined on the DHCPv4 server. The subnet elements can be IPv4 reservation
	// for DHCPv4 or BOOTP clients, IPv4 range, or IPv4 exclusion range for DHCPv4 or BOOTP
	// clients.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                       RETURN                       |                                                                                  |
	//	|                     VALUE/CODE                     |                                   DESCRIPTION                                    |
	//	|                                                    |                                                                                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                           | The call was successful.                                                         |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT           | The specified IPv4 subnet does not exist.                                        |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E27 ERROR_DHCP_ELEMENT_CANT_REMOVE          | This error can occur for any of the following reasons: The specified IPv4 subnet |
	//	|                                                    | element cannot be removed because at least one IPv4 address has been leased out  |
	//	|                                                    | to a client in the subnet. The starting address of the specified IPv4 exclusion  |
	//	|                                                    | range is not part of any exclusion range configured on the server. There is an   |
	//	|                                                    | error in deleting the exclusion range from the database.                         |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR                    | An error occurred while accessing the DHCP server database.                      |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E37 ERROR_DHCP_INVALID_RANGE                | The specified IPv4 range does not match an existing IPv4 range.                  |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E90 ERROR_SCOPE_RANGE_POLICY_RANGE_CONFLICT | There is an IP address range configured for a policy in this scope. This         |
	//	|                                                    | operation on the scope IP address range cannot be performed until the policy IP  |
	//	|                                                    | address range is suitably modified.                                              |
	//	+----------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 39.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemoveSubnetElementV5(context.Context, *RemoveSubnetElementV5Request) (*RemoveSubnetElementV5Response, error)

	// The R_DhcpGetServerBindingInfo method retrieves the array of IPv4 interface binding
	// information for the DHCPv4 server. The caller of this function can free the memory
	// pointed by BindElementsInfo by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 40.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetServerBindingInfo(context.Context, *GetServerBindingInfoRequest) (*GetServerBindingInfoResponse, error)

	// The R_DhcpSetServerBindingInfo method sets/modifies the IPv4 interface bindings for
	// the DHCPv4 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	|                    RETURN                    |                                                                                  |
	//	|                  VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                              |                                                                                  |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                     | The call was successful.                                                         |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E52 ERROR_DHCP_NETWORK_CHANGED        | The network has changed. Retry this operation after checking for the network     |
	//	|                                              | changes. Network changes can be caused by interfaces that are new or no longer   |
	//	|                                              | valid, or by IPv4 addresses that are new or no longer valid.                     |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E53 ERROR_DHCP_CANNOT_MODIFY_BINDINGS | The bindings to internal IPv4 addresses cannot be modified.                      |
	//	+----------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 41.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetServerBindingInfo(context.Context, *SetServerBindingInfoRequest) (*SetServerBindingInfoResponse, error)

	// The R_DhcpQueryDnsRegCredentials method retrieves the currently set Domain Name System
	// (DNS) credentials, which are the user name and domain. These credentials are used
	// by the DHCP server for DNS dynamic registration for DHCP clients.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 42.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	QueryDNSRegCredentials(context.Context, *QueryDNSRegCredentialsRequest) (*QueryDNSRegCredentialsResponse, error)

	// The R_DhcpSetDnsRegCredentials method sets the DNS user name and credentials in the
	// DHCP server which is used for DNS registrations for the DHCP client lease record.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 43.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetDNSRegCredentials(context.Context, *SetDNSRegCredentialsRequest) (*SetDNSRegCredentialsResponse, error)

	// The R_DhcpBackupDatabase method takes backup of the configurations, settings, and
	// DHCP client lease record in the specified path.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 44.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	BackupDatabase(context.Context, *BackupDatabaseRequest) (*BackupDatabaseResponse, error)

	// The R_DhcpRestoreDatabase method sets/modifies the restore path. The DHCP server
	// uses this path to restore the configuration, settings, and DHCP client lease record
	// the next time it is restarted.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+----------------------------------------------------+
	//	|             RETURN              |                                                    |
	//	|           VALUE/CODE            |                    DESCRIPTION                     |
	//	|                                 |                                                    |
	//	+---------------------------------+----------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                           |
	//	+---------------------------------+----------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server. |
	//	+---------------------------------+----------------------------------------------------+
	//
	// The opnum field value for this method is 45.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RestoreDatabase(context.Context, *RestoreDatabaseRequest) (*RestoreDatabaseResponse, error)

	// The R_DhcpGetServerSpecificStrings method retrieves the names of the default vendor
	// class and user class. The caller of this function can free the memory pointed to
	// by ServerSpecificStrings, DefaultVendorClassName and DefaultUserClassName by calling
	// the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 46.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetServerSpecificStrings(context.Context, *GetServerSpecificStringsRequest) (*GetServerSpecificStringsResponse, error)

	// The R_DhcpCreateOptionV6 method creates an option definition for a specified user
	// class or vendor class at the default option level. The option ID specifies the identifier
	// of the option. If the user class or vendor class is not specified, the option definition
	// is created for the default user class or vendor class.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                      |                                                                                  |
	//	|                    VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                  |                                                                                  |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                         | The call was successful.                                                         |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E29 ERROR_DHCP_OPTION_EXITS               | The specified option definition already exists in the DHCP server database.      |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR                  | An error occurred while accessing the DHCP server database.                      |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E59 ERROR_DHCP_INVALID_PARAMETER_OPTION32 | The information refresh time option value is invalid, as it is less than the     |
	//	|                                                  | minimum option value.                                                            |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 47.
	CreateOptionV6(context.Context, *CreateOptionV6Request) (*CreateOptionV6Response, error)

	// The R_DhcpSetOptionInfoV6 method modifies the option definition for the specific
	// user class and vendor class at the default level. If the user class or vendor class
	// is not specified, the default user class or vendor class will be used.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                      |                                                                                  |
	//	|                    VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                  |                                                                                  |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                         | The call was successful.                                                         |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT         | The option to be modified does not exist.                                        |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR                  | An error occurred while accessing the DHCP server database.                      |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E59 ERROR_DHCP_INVALID_PARAMETER_OPTION32 | The information refresh time option value is invalid, as it is less than the     |
	//	|                                                  | minimum option value.                                                            |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 48.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetOptionInfoV6(context.Context, *SetOptionInfoV6Request) (*SetOptionInfoV6Response, error)

	// The R_DhcpGetOptionInfoV6 method retrieves the option definition of a specific option
	// for a specific user class and vendor class at the default option level. If the user
	// class or vendor class is not specified, the default vendor class or user class will
	// be taken. The caller of this function can free the memory pointed to by OptionInfo
	// by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------+
	//	|                  RETURN                  |                                           |
	//	|                VALUE/CODE                |                DESCRIPTION                |
	//	|                                          |                                           |
	//	+------------------------------------------+-------------------------------------------+
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                  |
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT | The option to be modified does not exist. |
	//	+------------------------------------------+-------------------------------------------+
	//
	// The opnum field value for this method is 49.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetOptionInfoV6(context.Context, *GetOptionInfoV6Request) (*GetOptionInfoV6Response, error)

	// The R_DhcpEnumOptionsV6 method enumerates the option definitions for a specific user
	// class and vendor class at the default option level. If the user class or vendor class
	// is not specified, the default user class or vendor class will be used. The caller
	// of this function can free the memory pointed to by Options by calling the function
	// midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA     | There are more elements available to enumerate. |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate.   |
	//	+--------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 50.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumOptionsV6(context.Context, *EnumOptionsV6Request) (*EnumOptionsV6Response, error)

	// The R_DhcpRemoveOptionV6 method removes the option definition of a specific option
	// for a specific user class or the vendor class at the default option level. If the
	// user class or the vendor class is not specified, the default user class or vendor
	// class will be used. The option id specifies the identifier of the option definition.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------+
	//	|                  RETURN                  |                                           |
	//	|                VALUE/CODE                |                DESCRIPTION                |
	//	|                                          |                                           |
	//	+------------------------------------------+-------------------------------------------+
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                  |
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT | The option to be modified does not exist. |
	//	+------------------------------------------+-------------------------------------------+
	//
	// The opnum field value for this method is 51.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemoveOptionV6(context.Context, *RemoveOptionV6Request) (*RemoveOptionV6Response, error)

	// The R_DhcpSetOptionValueV6 method creates option value when called for the first
	// time, else it modifies the option value of a specific option on the DHCPv6 server
	// for a specific user class and vendor class. ScopeInfo defines the scope on which
	// this option value is set. If the user class and vendor class is not provided, the
	// default user class and vendor class is taken.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                      RETURN                      |                                                                                  |
	//	|                    VALUE/CODE                    |                                   DESCRIPTION                                    |
	//	|                                                  |                                                                                  |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                         | The call was successful.                                                         |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT         | The option to be modified does not exist.                                        |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E59 ERROR_DHCP_INVALID_PARAMETER_OPTION32 | The information refresh time option value is invalid, as it is less than the     |
	//	|                                                  | minimum option value.                                                            |
	//	+--------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 52.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetOptionValueV6(context.Context, *SetOptionValueV6Request) (*SetOptionValueV6Response, error)

	// The R_DhcpEnumOptionValuesV6 method enumerates all the option values for the specific
	// user class or vendor class at a specified scope defined by ScopeInfo. If the user
	// class or vendor class is not specified, the default user class or vendor class will
	// be used. The caller of this function can free the memory pointed to by OptionValues
	// and the Values member of OptionValues by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA     | There are more elements available to enumerate. |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate.   |
	//	+--------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 53.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumOptionValuesV6(context.Context, *EnumOptionValuesV6Request) (*EnumOptionValuesV6Response, error)

	// The R_DhcpRemoveOptionValueV6 method deletes the option value of a specific option
	// on the DHCPv6 server for a specific user and vendor class. ScopeInfo defines the
	// scope from which this option value is removed. If the user class or vendor class
	// is not provided, the default user or vendor class is taken.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+--------------------------------------+
	//	|                  RETURN                  |                                      |
	//	|                VALUE/CODE                |             DESCRIPTION              |
	//	|                                          |                                      |
	//	+------------------------------------------+--------------------------------------+
	//	+------------------------------------------+--------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.             |
	//	+------------------------------------------+--------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT | The specified option does not exist. |
	//	+------------------------------------------+--------------------------------------+
	//
	// The opnum field value for this method is 54.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemoveOptionValueV6(context.Context, *RemoveOptionValueV6Request) (*RemoveOptionValueV6Response, error)

	// The R_DhcpGetAllOptionsV6 method retrieves all default option definitions, as well
	// as specific user class and vendor class option definitions. The caller of this function
	// can free the memory pointed to by OptionStruct, NonVendorOptions and VendorOptions
	// and by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 55.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetAllOptionsV6(context.Context, *GetAllOptionsV6Request) (*GetAllOptionsV6Response, error)

	// The R_DhcpGetAllOptionValuesV6 method returns all option values for all user classes
	// and vendor classes configured at the server, scope, or IPv6 reservation level on
	// the DHCPv6 server. The caller of this function can free the memory pointed to by
	// option Values by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 56.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetAllOptionValuesV6(context.Context, *GetAllOptionValuesV6Request) (*GetAllOptionValuesV6Response, error)

	// The R_DhcpCreateSubnetV6 method creates a new IPv6 prefix on the DHCPv6 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------------------+-------------------------------------------------------------+
	//	|                   RETURN                    |                                                             |
	//	|                 VALUE/CODE                  |                         DESCRIPTION                         |
	//	|                                             |                                                             |
	//	+---------------------------------------------+-------------------------------------------------------------+
	//	+---------------------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                    | The call was successful.                                    |
	//	+---------------------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR             | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------------------+-------------------------------------------------------------+
	//	| 0x00004E7B ERROR_DHCP_INVALID_SUBNET_PREFIX | The subnet prefix is invalid.                               |
	//	+---------------------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 57.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CreateSubnetV6(context.Context, *CreateSubnetV6Request) (*CreateSubnetV6Response, error)

	// The R_DhcpEnumSubnetsV6 method enumerates all IPv6 prefixes configured on the DHCPv6
	// server. The caller of this function can free the memory pointed to by EnumInfo by
	// calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------+-----------------------------------------------+
	//	|             RETURN             |                                               |
	//	|           VALUE/CODE           |                  DESCRIPTION                  |
	//	|                                |                                               |
	//	+--------------------------------+-----------------------------------------------+
	//	+--------------------------------+-----------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                      |
	//	+--------------------------------+-----------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate. |
	//	+--------------------------------+-----------------------------------------------+
	//
	// The opnum field value for this method is 58.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnetsV6(context.Context, *EnumSubnetsV6Request) (*EnumSubnetsV6Response, error)

	// The R_DhcpAddSubnetElementV6 method adds an IPv6 prefix element (such as IPv6 reservation
	// or IPv6 exclusion range) to the IPv6 prefix in the DHCPv6 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                 |                                                                                  |
	//	|               VALUE/CODE               |                                   DESCRIPTION                                    |
	//	|                                        |                                                                                  |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS               | The call was successful.                                                         |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER     | An invalid parameter is specified in the AddElementInfo parameter.               |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x000007DE ERROR_DUPLICATE_TAG         | The specified exclusion range conflicts with existing exclusion ranges.          |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR        | An error occurred while accessing the DHCP server database.                      |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E36 ERROR_DHCP_RESERVEDIP_EXITS | An IPv6 reservation exists for one or both of the following: the specified       |
	//	|                                        | IPv6 address the DHCPv6 client-identifier (section 2.2.1.2.5.3) and interface    |
	//	|                                        | identifier pair specified in reservation information                             |
	//	+----------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 59.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	AddSubnetElementV6(context.Context, *AddSubnetElementV6Request) (*AddSubnetElementV6Response, error)

	// The R_DhcpEnumSubnetElementsV6 method returns an enumerated list of a specific type
	// of IPv6 prefix element for a specific DHCPv6 IPv6 prefix. The caller of this function
	// can free the memory pointed to by EnumElementInfo and other Elements by calling the
	// function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA     | There are more elements available to enumerate. |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate.   |
	//	+--------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 60.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnetElementsV6(context.Context, *EnumSubnetElementsV6Request) (*EnumSubnetElementsV6Response, error)

	// The R_DhcpRemoveSubnetElementV6 method removes an IPv6 prefix element (such as IPv6
	// reservation or IPv6 exclusion range) from an IPv6 prefix defined on the DHCPv6 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 61.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemoveSubnetElementV6(context.Context, *RemoveSubnetElementV6Request) (*RemoveSubnetElementV6Response, error)

	// The R_DhcpDeleteSubnetV6 method deletes an IPv6 prefix from the DHCPv6 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                  |
	//	|                VALUE/CODE                 |                                   DESCRIPTION                                    |
	//	|                                           |                                                                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                         |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR           | An error occurred while accessing the DHCP server database.                      |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E27 ERROR_DHCP_ELEMENT_CANT_REMOVE | The specified subnet cannot be deleted because at least one IPv6 address has     |
	//	|                                           | been leased out to some client from the subnet.                                  |
	//	+-------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 62.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DeleteSubnetV6(context.Context, *DeleteSubnetV6Request) (*DeleteSubnetV6Response, error)

	// The R_DhcpGetSubnetInfoV6 method retrieves the information about a specific IPv6
	// prefix defined on the DHCPv6 server. The caller of this function can free the memory
	// pointed to by SubnetInfo by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------+
	//	|                  RETURN                  |                                           |
	//	|                VALUE/CODE                |                DESCRIPTION                |
	//	|                                          |                                           |
	//	+------------------------------------------+-------------------------------------------+
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                  |
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv6 prefix does not exist. |
	//	+------------------------------------------+-------------------------------------------+
	//
	// The opnum field value for this method is 63.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetSubnetInfoV6(context.Context, *GetSubnetInfoV6Request) (*GetSubnetInfoV6Response, error)

	// The R_DhcpEnumSubnetClientsV6 method is used to retrieve all DHCPv6 clients serviced
	// from the specified IPv6 prefix. The caller of this function can free the memory pointed
	// to by ClientInfo and other Elements by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA      | There are more elements available to enumerate.             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS  | There are no more elements left to enumerate.               |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 64.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnetClientsV6(context.Context, *EnumSubnetClientsV6Request) (*EnumSubnetClientsV6Response, error)

	// The R_DhcpServerSetConfigV6 method sets the DHCPv6 server configuration data at the
	// scope level or at the server level.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 65.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ServerSetConfigV6(context.Context, *ServerSetConfigV6Request) (*ServerSetConfigV6Response, error)

	// The R_DhcpServerGetConfigV6 method retrieves the configuration information about
	// the DHCPv6 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 66.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ServerGetConfigV6(context.Context, *ServerGetConfigV6Request) (*ServerGetConfigV6Response, error)

	// The R_DhcpSetSubnetInfoV6 method sets/modifies the information for an IPv6 prefix
	// defined on the DHCPv6 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------------------------+
	//	|                  RETURN                  |                                                             |
	//	|                VALUE/CODE                |                         DESCRIPTION                         |
	//	|                                          |                                                             |
	//	+------------------------------------------+-------------------------------------------------------------+
	//	+------------------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                    |
	//	+------------------------------------------+-------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv6 prefix does not exist.                   |
	//	+------------------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR          | An error occurred while accessing the DHCP server database. |
	//	+------------------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 67.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetSubnetInfoV6(context.Context, *SetSubnetInfoV6Request) (*SetSubnetInfoV6Response, error)

	// The R_DhcpGetMibInfoV6 method is used to retrieve the IPv6 counter values of the
	// DHCPv6 server. The caller of this function can free the memory pointed to by MibInfo
	// by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 68.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetMIBInfoV6(context.Context, *GetMIBInfoV6Request) (*GetMIBInfoV6Response, error)

	// The R_DhcpGetServerBindingInfoV6 method retrieves the array of IPv6 interface binding
	// information for the DHCPv6 server. The caller of this function can free the memory
	// pointed to by BindElementsInfo by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 69.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetServerBindingInfoV6(context.Context, *GetServerBindingInfoV6Request) (*GetServerBindingInfoV6Response, error)

	// The R_DhcpSetServerBindingInfoV6 method sets/modifies the IPv6 interface bindings
	// for the DHCPv6 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	|                   RETURN                    |                                                                                  |
	//	|                 VALUE/CODE                  |                                   DESCRIPTION                                    |
	//	|                                             |                                                                                  |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                    | The call was successful.                                                         |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E52 ERROR_DHCP_NETWORK_CHANGED       | The network has changed. Retry this operation after checking for the network     |
	//	|                                             | changes. Network changes can be caused by interfaces that are new or no longer   |
	//	|                                             | valid or by IPv6 addresses that are new or no longer valid.                      |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E53 ERROR_DHCP_CANNOT_MODIFY_BINDING | The bindings to internal IPv6 addresses cannot be modified.                      |
	//	+---------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 70.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetServerBindingInfoV6(context.Context, *SetServerBindingInfoV6Request) (*SetServerBindingInfoV6Response, error)

	// The R_DhcpSetClientInfoV6 method sets/modifies the client reservation record on the
	// DHCPv6 server database. This method is supposed to be called only after the reserved
	// DHCPv6 client is added using the R_DhcpAddSubnetElementV6 (section 3.2.4.60) method.<69>
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 71.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetClientInfoV6(context.Context, *SetClientInfoV6Request) (*SetClientInfoV6Response, error)

	// The R_DhcpGetClientInfoV6 method retrieves IPv6 address lease information of the
	// IPv6 reservation from the DHCPv6 server. The caller of this function can free the
	// memory pointed to by ClientInfo by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database or the client entry   |
	//	|                                 | is not present in the database.                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 72.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetClientInfoV6(context.Context, *GetClientInfoV6Request) (*GetClientInfoV6Response, error)

	// The R_DhcpDeleteClientInfoV6 method deletes the specified DHCPv6 client address lease
	// record from the DHCPv6 server database.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	|                 RETURN                  |                                                                                  |
	//	|               VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                         |                                                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                | The call was successful.                                                         |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR         | An error occurred while accessing the DHCP server database or the client entry   |
	//	|                                         | is not present in the database.                                                  |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E36 ERROR_DHCP_RESERVEDIP_EXISTS | There exists a reservation for the leased address.                               |
	//	+-----------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 73.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DeleteClientInfoV6(context.Context, *DeleteClientInfoV6Request) (*DeleteClientInfoV6Response, error)

	// The R_DhcpCreateClassV6 method creates an IPv6 user class or a vendor class definition
	// on the DHCPv6 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------------------+----------------------------------------------------------------------+
	//	|                   RETURN                   |                                                                      |
	//	|                 VALUE/CODE                 |                             DESCRIPTION                              |
	//	|                                            |                                                                      |
	//	+--------------------------------------------+----------------------------------------------------------------------+
	//	+--------------------------------------------+----------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                   | The call was successful.                                             |
	//	+--------------------------------------------+----------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR            | An error occurred while accessing the DHCP server database.          |
	//	+--------------------------------------------+----------------------------------------------------------------------+
	//	| 0x00004E4D ERROR_DHCP_CLASS_ALREADY_EXISTS | The vendor class or user class that is being created already exists. |
	//	+--------------------------------------------+----------------------------------------------------------------------+
	//
	// The opnum field value for this method is 74.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	CreateClassV6(context.Context, *CreateClassV6Request) (*CreateClassV6Response, error)

	// The R_DhcpModifyClassV6 method modifies the user class or vendor class definition
	// for the DHCPv6 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 75.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	ModifyClassV6(context.Context, *ModifyClassV6Request) (*ModifyClassV6Response, error)

	// The R_DhcpDeleteClassV6 method deletes the specified IPv6 user class or vendor class
	// definition from the DHCPv6 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------------+-------------------------------------------------------------+
	//	|                RETURN                 |                                                             |
	//	|              VALUE/CODE               |                         DESCRIPTION                         |
	//	|                                       |                                                             |
	//	+---------------------------------------+-------------------------------------------------------------+
	//	+---------------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS              | The call was successful.                                    |
	//	+---------------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR       | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------------+-------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND | The specified class is not defined in the DHCP server.      |
	//	+---------------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 76.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DeleteClassV6(context.Context, *DeleteClassV6Request) (*DeleteClassV6Response, error)

	// The R_DhcpEnumClassesV6 method enumerates user or vendor classes configured for the
	// DHCPv6 server. The caller of this function can free the memory pointed to by ClassInfoArray
	// and its Classes member by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA     | There are more elements available to enumerate. |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate.   |
	//	+--------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 77.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumClassesV6(context.Context, *EnumClassesV6Request) (*EnumClassesV6Response, error)

	// The R_DhcpGetOptionValueV6 method retrieves the option value for a specific option
	// on the DHCPv6 server for specific user and vendor class. ScopeInfo defines the scope
	// from which the option value needs to be retrieved. The caller of this function can
	// free the memory pointed by OptionValue by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+-------------------------------------------+--------------------------------------------------------------------------------+
	//	|                  RETURN                   |                                                                                |
	//	|                VALUE/CODE                 |                                  DESCRIPTION                                   |
	//	|                                           |                                                                                |
	//	+-------------------------------------------+--------------------------------------------------------------------------------+
	//	+-------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                  | The call was successful.                                                       |
	//	+-------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT  | The specified subnet is not defined on the DHCP server.                        |
	//	+-------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT  | The specified option is not defined at the specified level in the DHCP server. |
	//	+-------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT | The reserved IPv6 client is not defined on the DHCP server.                    |
	//	+-------------------------------------------+--------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 78.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetOptionValueV6(context.Context, *GetOptionValueV6Request) (*GetOptionValueV6Response, error)

	// The R_DhcpSetSubnetDelayOffer method sets/modifies the time delay setting on the
	// DHCPv4 server, which is used in responding to a DHCPDISCOVER message [RFC2131]. This
	// setting is configured for a specific scope.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified subnet is not defined on the DHCP server.                          |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E7C ERROR_DHCP_INVALID_DELAY      | The specified delay value is invalid, it is greater than the maximum delay of    |
	//	|                                          | 1000 milliseconds.                                                               |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 79.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetSubnetDelayOffer(context.Context, *SetSubnetDelayOfferRequest) (*SetSubnetDelayOfferResponse, error)

	// The R_DhcpGetSubnetDelayOffer method retrieves the time delay setting from the DHCPv4
	// server, which is used in responding to a DHCPDISCOVER message [RFC2131] for a specific
	// scope.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully;
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------+---------------------------------------------------------+
	//	|                  RETURN                  |                                                         |
	//	|                VALUE/CODE                |                       DESCRIPTION                       |
	//	|                                          |                                                         |
	//	+------------------------------------------+---------------------------------------------------------+
	//	+------------------------------------------+---------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                |
	//	+------------------------------------------+---------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified subnet is not defined on the DHCP server. |
	//	+------------------------------------------+---------------------------------------------------------+
	//
	// The opnum field value for this method is 80.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetSubnetDelayOffer(context.Context, *GetSubnetDelayOfferRequest) (*GetSubnetDelayOfferResponse, error)

	// The R_DhcpGetMibInfoV5 method is used to retrieve the statistics of the DHCPv4 server.
	// The caller of this function can free the memory pointed to by MibInfo and its field
	// ScopeInfo by calling the function midl_user_free (see section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 81.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetMIBInfoV5(context.Context, *GetMIBInfoV5Request) (*GetMIBInfoV5Response, error)

	// The R_DhcpAddFilterV4 method is used to add a link-layer address/pattern to allow
	// list or deny list. The DHCPv4 server allows the DHCPv4 clients whose link-layer address
	// is in the allow list to be given leases and blocks DHCPv4 clients whose link-layer
	// address is in the deny list provided the respective lists are enabled using the R_DhcpSetFilterV4
	// (section 3.2.4.85) method. This method is also used to exempt one or more hardware
	// types from filtering. However, hardware type 1 (Ethernet 10 Mb) cannot be exempted.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+------------------------------------------------------------+----------------------------------------------------------------------+
	//	|                           RETURN                           |                                                                      |
	//	|                         VALUE/CODE                         |                             DESCRIPTION                              |
	//	|                                                            |                                                                      |
	//	+------------------------------------------------------------+----------------------------------------------------------------------+
	//	+------------------------------------------------------------+----------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                                   | The call was successful.                                             |
	//	+------------------------------------------------------------+----------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR                            | An error occurred while accessing the DHCP server database.          |
	//	+------------------------------------------------------------+----------------------------------------------------------------------+
	//	| 0x00004E7E ERROR_DHCP_LINKLAYER_ADDRESS_EXISTS             | Address or Address pattern is already contained in one of the lists. |
	//	+------------------------------------------------------------+----------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER                         | Invalid input - address/pattern                                      |
	//	+------------------------------------------------------------+----------------------------------------------------------------------+
	//	| 0x00004E85 ERROR_DHCP_HARDWARE_ADDRESS_TYPE_ALREADY_EXEMPT | Hardware type already exempted from filtering.                       |
	//	+------------------------------------------------------------+----------------------------------------------------------------------+
	//
	// The opnum field value for this method is 82.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	AddFilterV4(context.Context, *AddFilterV4Request) (*AddFilterV4Response, error)

	// The R_DhcpDeleteFilterV4 method is used to delete a link-layer address/pattern from
	// allow list or deny list. This method is also used to delete an exemption of a hardware
	// type from filtering. However, hardware type 1 (Ethernet 10 Mb) cannot be exempted,
	// and this method cannot be used to delete them.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------------------------------------+-----------------------------------------------------------------+
	//	|                         RETURN                         |                                                                 |
	//	|                       VALUE/CODE                       |                           DESCRIPTION                           |
	//	|                                                        |                                                                 |
	//	+--------------------------------------------------------+-----------------------------------------------------------------+
	//	+--------------------------------------------------------+-----------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                               | The call was successful.                                        |
	//	+--------------------------------------------------------+-----------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR                        | An error occurred while accessing the DHCP Server Database.     |
	//	+--------------------------------------------------------+-----------------------------------------------------------------+
	//	| 0x00004E7F ERROR_DHCP_LINKLAYER_ADDRESS_DOES_NOT_EXIST | Address or Address pattern is not contained in any of the list. |
	//	+--------------------------------------------------------+-----------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER                     | Invalid input - address/pattern                                 |
	//	+--------------------------------------------------------+-----------------------------------------------------------------+
	//	| 0x00004E86 ERROR_DHCP_UNDEFINED_HARDWARE_ADDRESS_TYPE  | Hardware type not present in the exemption list.                |
	//	+--------------------------------------------------------+-----------------------------------------------------------------+
	//
	// The opnum field value for this method is 83.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DeleteFilterV4(context.Context, *DeleteFilterV4Request) (*DeleteFilterV4Response, error)

	// The R_DhcpSetFilterV4 method is used to enable or disable the allow and deny lists.
	// The DHCPv4 server allows the DHCPv4 clients whose link-layer address is in the allow
	// list to be given leases and blocks DHCPv4 clients whose link-layer address is in
	// the deny list, provided the respective lists are enabled using R_DhcpSetFilterV4.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 84.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetFilterV4(context.Context, *SetFilterV4Request) (*SetFilterV4Response, error)

	// The R_DhcpGetFilterV4 method is used to retrieve the enable or disable settings for
	// the allow and deny lists.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 85.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetFilterV4(context.Context, *GetFilterV4Request) (*GetFilterV4Response, error)

	// The R_DhcpEnumFilterV4 method enumerates all the filter records from either allow
	// list or deny list. It also returns a list of hardware types presently exempted from
	// filtering. These entries are present in the allow list. Exemption entries have a
	// pattern of Length 0 and IsWildCard set to TRUE; both are specified in the AddrPatt
	// field of the DHCP_FILTER_RECORD (section 2.2.1.2.92) structure.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_NO_MORE_ITEMS (0x00000103) indicates that the operation was completed
	// successfully. Otherwise, it contains a Win32 error code, as specified in [MS-ERREF].
	// This error code value can correspond to a DHCP-specific failure, which takes a value
	// between 20000 and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA      | There are more elements available to enumerate.             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS  | There are no more elements left to enumerate.               |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 86.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumFilterV4(context.Context, *EnumFilterV4Request) (*EnumFilterV4Response, error)

	// The R_DhcpSetDnsRegCredentials method sets the DNS user name and credentials in the
	// DHCP server which is used for DNS registrations for DHCP client lease record.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully,
	// else it contains a Win32 error code, as specified in [MS-ERREF]. This error code
	// value can correspond to a DHCP-specific failure, which takes a value between 20000
	// and 20099, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 87.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetDNSRegCredentialsV5(context.Context, *SetDNSRegCredentialsV5Request) (*SetDNSRegCredentialsV5Response, error)

	// The R_DhcpEnumSubnetClientsFilterStatusInfo method is used to retrieve all DHCPv4
	// clients serviced on the specified IPv4 subnet. The information also includes the
	// link-layer filter status info for the DHCPv4 client.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP specific failure, which takes a value between
	// 20000 and 20099, or any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA      | There are more elements available to enumerate.             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS  | There are no more elements left to enumerate.               |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// The opnum field value for this method is 88.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnetClientsFilterStatusInfo(context.Context, *EnumSubnetClientsFilterStatusInfoRequest) (*EnumSubnetClientsFilterStatusInfoResponse, error)

	// The R_DhcpV4FailoverCreateRelationship method is used to create a new failover relationship
	// on the DHCPv4 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates the return status.
	// A return value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed
	// successfully. Otherwise, it contains a Win32 error code, as specified in [MS-ERREF].
	// This error code value can correspond to a DHCP-specific failure, which takes a value
	// between 20000 and 20123, or to any generic failure.
	//
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                         RETURN                         |                                                                                  |
	//	|                       VALUE/CODE                       |                                   DESCRIPTION                                    |
	//	|                                                        |                                                                                  |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                               | The call was successful.                                                         |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT               | IPv4 scope does not exist on the DHCPv4 server.                                  |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E90 ERROR_DHCP_FO_SCOPE_ALREADY_IN_RELATIONSHIP | IPv4 is already part of another failover relationship.                           |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E91 ERROR_DHCP_FO_RELATIONSHIP_EXISTS           | A failover relationship already exists on the DHCPv4 server.                     |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E9D ERROR_DHCP_FO_RELATIONSHIP_NAME_TOO_LONG    | The failover relationship name in the DHCP_FAILOVER_RELATIONSHIP (section        |
	//	|                                                        | 2.2.1.2.98) structure is too long.                                               |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004EA0 ERROR_DHCP_FO_MAX_RELATIONSHIPS             | The maximum number of allowed failover relationships configured on the DHCP      |
	//	|                                                        | server has been exceeded.                                                        |
	//	+--------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 89.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	FailoverCreateRelationshipV4(context.Context, *FailoverCreateRelationshipV4Request) (*FailoverCreateRelationshipV4Response, error)

	// The R_DhcpV4FailoverSetRelationship method is used to modify an existing failover
	// relationship on the DHCPv4 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. The
	// return value ERROR_SUCCESS (0x00000000) indicates that the operation was completed
	// successfully. Otherwise, it contains a Win32 error code, as specified in [MS-ERREF].
	// This error code value can correspond to a DHCP-specific failure, which takes a value
	// between 20000 and 20123, or any generic failure.
	//
	//	+------------------------------------------------------+------------------------------------------+
	//	|                        RETURN                        |                                          |
	//	|                      VALUE/CODE                      |               DESCRIPTION                |
	//	|                                                      |                                          |
	//	+------------------------------------------------------+------------------------------------------+
	//	+------------------------------------------------------+------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                             | The call was successful.                 |
	//	+------------------------------------------------------+------------------------------------------+
	//	| 0x00004E92 ERROR_DHCP_FO_RELATIONSHIP_DOES_NOT_EXIST | The failover relationship doesn’t exist. |
	//	+------------------------------------------------------+------------------------------------------+
	//
	// The opnum field value for this method is 90.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	FailoverSetRelationshipV4(context.Context, *FailoverSetRelationshipV4Request) (*FailoverSetRelationshipV4Response, error)

	// The R_DhcpV4FailoverDeleteRelationship method is used to delete an existing failover
	// relationship on the DHCPv4 server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. The
	// return value ERROR_SUCCESS (0x00000000) indicates that the operation was completed
	// successfully. Otherwise, it contains a Win32 error code, as specified in [MS-ERREF].
	// This error code value can correspond to a DHCP-specific failure, which takes a value
	// between 20000 and 20123, or any generic failure.
	//
	//	+------------------------------------------------------+------------------------------------------+
	//	|                        RETURN                        |                                          |
	//	|                      VALUE/CODE                      |               DESCRIPTION                |
	//	|                                                      |                                          |
	//	+------------------------------------------------------+------------------------------------------+
	//	+------------------------------------------------------+------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                             | The call was successful.                 |
	//	+------------------------------------------------------+------------------------------------------+
	//	| 0x00004E92 ERROR_DHCP_FO_RELATIONSHIP_DOES_NOT_EXIST | The failover relationship doesn't exist. |
	//	+------------------------------------------------------+------------------------------------------+
	//
	// The opnum field value for this method is 91.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	FailoverDeleteRelationshipV4(context.Context, *FailoverDeleteRelationshipV4Request) (*FailoverDeleteRelationshipV4Response, error)

	// The R_DhcpV4FailoverGetRelationship method retrieves the failover relationship information
	// configured on the DHCPv4 server. The caller of this function can free the memory
	// pointed to by the pRelationship parameter by calling the function midl_user_free
	// (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20123, or any generic failure.
	//
	//	+------------------------------------------------------+-------------------------------------------+
	//	|                        RETURN                        |                                           |
	//	|                      VALUE/CODE                      |                DESCRIPTION                |
	//	|                                                      |                                           |
	//	+------------------------------------------------------+-------------------------------------------+
	//	+------------------------------------------------------+-------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                             | The call was successful.                  |
	//	+------------------------------------------------------+-------------------------------------------+
	//	| 0x00004E92 ERROR_DHCP_FO_RELATIONSHIP_DOES_NOT_EXIST | The failover relationship does not exist. |
	//	+------------------------------------------------------+-------------------------------------------+
	//
	// The opnum field value for this method is 92.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	FailoverGetRelationshipV4(context.Context, *FailoverGetRelationshipV4Request) (*FailoverGetRelationshipV4Response, error)

	// The R_DhcpV4FailoverEnumRelationship method enumerates all the failover relationships
	// on the DHCPv4 server. The caller of this function can free the memory pointed to
	// by the pRelationship parameter by calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20123, or any generic failure.
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA     | There are more elements available to enumerate. |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate.   |
	//	+--------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 93.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	FailoverEnumRelationshipV4(context.Context, *FailoverEnumRelationshipV4Request) (*FailoverEnumRelationshipV4Response, error)

	// The R_DhcpV4FailoverAddScopeToRelationship method adds scopes to an existing failover
	// relationship.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20123, or to any generic failure.
	//
	//	+--------------------------------------------------------+--------------------------------------------------------------------------------+
	//	|                         RETURN                         |                                                                                |
	//	|                       VALUE/CODE                       |                                  DESCRIPTION                                   |
	//	|                                                        |                                                                                |
	//	+--------------------------------------------------------+--------------------------------------------------------------------------------+
	//	+--------------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                               | The call was successful.                                                       |
	//	+--------------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT               | IPv4 scope does not exist on the DHCPv4 server.                                |
	//	+--------------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004E91 ERROR_DHCP_FO_SCOPE_ALREADY_IN_RELATIONSHIP | IPv4 scope is already part of another failover relationship.                   |
	//	+--------------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004E92 ERROR_DHCP_FO_RELATIONSHIP_DOES_NOT_EXIST   | Failover relationship does not exist.                                          |
	//	+--------------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004EA5 ERROR_DHCP_FO_SCOPE_SYNC_IN_PROGRESS        | Failover relationship is being re-integrated with the failover partner server. |
	//	+--------------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004E98 ERROR_DHCP_FO_STATE_NOT_NORMAL              | Failover relationship is not in the NORMAL state.                              |
	//	+--------------------------------------------------------+--------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 94.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	FailoverAddScopeToRelationshipV4(context.Context, *FailoverAddScopeToRelationshipV4Request) (*FailoverAddScopeToRelationshipV4Response, error)

	// The R_DhcpV4FailoverDeleteScopeFromRelationship method is used to delete one or more
	// scopes from an existing failover relationship.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20123, or any generic failure.
	//
	//	+------------------------------------------------------+--------------------------------------------------------------------------------+
	//	|                        RETURN                        |                                                                                |
	//	|                      VALUE/CODE                      |                                  DESCRIPTION                                   |
	//	|                                                      |                                                                                |
	//	+------------------------------------------------------+--------------------------------------------------------------------------------+
	//	+------------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                             | The call was successful.                                                       |
	//	+------------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT             | IPv4 scope doesn't exist on the DHCPv4 server.                                 |
	//	+------------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004E92 ERROR_DHCP_FO_RELATIONSHIP_DOES_NOT_EXIST | Failover relationship doesn't exist.                                           |
	//	+------------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004E94 ERROR_DHCP_FO_SCOPE_NOT_IN_RELATIONSHIP   | IPv4 subnet is not part of the failover relationship.                          |
	//	+------------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004EA5 ERROR_DHCP_FO_SCOPE_SYNC_IN_PROGRESS      | Failover relationship is being re-integrated with the failover partner server. |
	//	+------------------------------------------------------+--------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 95.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	FailoverDeleteScopeFromRelationshipV4(context.Context, *FailoverDeleteScopeFromRelationshipV4Request) (*FailoverDeleteScopeFromRelationshipV4Response, error)

	// The R_DhcpV4FailoverGetScopeRelationship method retrieves the failover relationship
	// information which is configured for a specific IPv4 subnet address. The caller of
	// this function can free the memory pointed to by the pRelationship parameter by calling
	// the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20123, or any generic failure.
	//
	//	+----------------------------------------------------+-------------------------------------------------------+
	//	|                       RETURN                       |                                                       |
	//	|                     VALUE/CODE                     |                      DESCRIPTION                      |
	//	|                                                    |                                                       |
	//	+----------------------------------------------------+-------------------------------------------------------+
	//	+----------------------------------------------------+-------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                           | The call was successful.                              |
	//	+----------------------------------------------------+-------------------------------------------------------+
	//	| 0x00004E93 ERROR_DHCP_FO_SCOPE_NOT_IN_RELATIONSHIP | IPv4 subnet is not part of the failover relationship. |
	//	+----------------------------------------------------+-------------------------------------------------------+
	//
	// The opnum field value for this method is 96.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	FailoverGetScopeRelationshipV4(context.Context, *FailoverGetScopeRelationshipV4Request) (*FailoverGetScopeRelationshipV4Response, error)

	// The R_DhcpV4FailoverGetScopeStatistics method is used to retrieve the statistics
	// of a IPv4 subnet configured for a failover relationship on the DHCPv4 server. The
	// caller of this function can free the memory pointed to by the pStats parameter by
	// calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20123, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 97.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	FailoverGetScopeStatisticsV4(context.Context, *FailoverGetScopeStatisticsV4Request) (*FailoverGetScopeStatisticsV4Response, error)

	// The R_DhcpV4FailoverGetClientInfo method retrieves DHCPv4 client lease record information
	// from the DHCPv4 server database. The caller of this function can free the memory
	// pointed to by the ClientInfo parameter, by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or any generic failure.
	//
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	|             RETURN              |                                                                                  |
	//	|           VALUE/CODE            |                                   DESCRIPTION                                    |
	//	|                                 |                                                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                                         |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database or the client entry   |
	//	|                                 | is not present in the database.                                                  |
	//	+---------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 98.
	//
	// Exceptions Thrown: No exceptions SHOULD be thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	FailoverGetClientInfoV4(context.Context, *FailoverGetClientInfoV4Request) (*FailoverGetClientInfoV4Response, error)

	// The R_DhcpV4FailoverGetSystemTime method is used to return the current time on the
	// DHCP server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20123, or any generic failure.
	//
	//	+--------------------------+--------------------------+
	//	|          RETURN          |                          |
	//	|        VALUE/CODE        |       DESCRIPTION        |
	//	|                          |                          |
	//	+--------------------------+--------------------------+
	//	+--------------------------+--------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call was successful. |
	//	+--------------------------+--------------------------+
	//
	// The opnum field value for this method is 99.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	FailoverGetSystemTimeV4(context.Context, *FailoverGetSystemTimeV4Request) (*FailoverGetSystemTimeV4Response, error)

	// The R_DhcpV4FailoverTriggerAddrAllocation method re-distributes the free addresses
	// between the primary server and secondary server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20123, or any generic failure.
	//
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                        RETURN                        |                                                                                  |
	//	|                      VALUE/CODE                      |                                   DESCRIPTION                                    |
	//	|                                                      |                                                                                  |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                             | The call was successful.                                                         |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E92 ERROR_DHCP_FO_RELATIONSHIP_DOES_NOT_EXIST | Failover relationship doesn't exit.                                              |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E94 ERROR_DHCP_FO_RELATION_IS_SECONDARY       | serverType member of failover relationship is SecondaryServer enumeration value. |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 100.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	FailoverTriggerAddrAllocationV4(context.Context, *FailoverTriggerAddrAllocationV4Request) (*FailoverTriggerAddrAllocationV4Response, error)

	// The R_DhcpV4SetOptionValue method sets the option value for a policy at the specified
	// level (scope or server).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                        RETURN                         |                                                                                  |
	//	|                      VALUE/CODE                       |                                   DESCRIPTION                                    |
	//	|                                                       |                                                                                  |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                              | The call was successful.                                                         |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT              | The specified IPv4 subnet does not exist on the DHCP server.                     |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT              | The specified option definition does not exist on the DHCP server database.      |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND                 | The class name being used is unknown or incorrect.                               |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8FL ERROR_DHCP_POLICY_NOT_FOUND               | The specified policy name does not exist.                                        |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004EA8L ERROR_DHCP_POLICY_FQDN_OPIION_UNSUPPORTED | The option value cannot be specified because the policy contains an FQDN-based   |
	//	|                                                       | condition.                                                                       |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 101.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetOptionValueV4(context.Context, *SetOptionValueV4Request) (*SetOptionValueV4Response, error)

	// The R_DhcpV4SetOptionValues method sets the specified option values for a policy
	// at the specified level.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                        RETURN                         |                                                                                  |
	//	|                      VALUE/CODE                       |                                   DESCRIPTION                                    |
	//	|                                                       |                                                                                  |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                              | The call was successful.                                                         |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT              | The specified IPv4 subnet does not exist on the DHCP server.                     |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT              | The specified option definition does not exist on the DHCP server database.      |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E32 ERROR_DHCP_NOT_RESERVED_CLIENT             | The specified DHCP client is not a reserved client.                              |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND                 | The class name being used is unknown or incorrect.                               |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8FL ERROR_DHCP_POLICY_NOT_PRESENT             | The specified policy name does not exist.                                        |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004EA8L ERROR_DHCP_POLICY_FQDN_OPTION_UNSUPPORTED | The option value cannot be specified because the policy contains an FQDN-based   |
	//	|                                                       | condition.                                                                       |
	//	+-------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 102.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetOptionValuesV4(context.Context, *SetOptionValuesV4Request) (*SetOptionValuesV4Response, error)

	// The R_DhcpV4GetOptionValue method gets the option value for the specified PolicyName
	// parameter and OptionID parameter. The memory for the OptionValue parameter is allocated
	// by this method and can be freed by the caller by calling the function midl_user_free
	// (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                             |
	//	|                VALUE/CODE                |                                 DESCRIPTION                                 |
	//	|                                          |                                                                             |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                                    |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist on the DHCP server.                |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT | The specified option definition does not exist on the DHCP server database. |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND    | The class name being used is unknown or incorrect.                          |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//	| 0x00004E8F ERROR_DHCP_POLICY_NOT_PRESENT | The specified policy name does not exist.                                   |
	//	+------------------------------------------+-----------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 103.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetOptionValueV4(context.Context, *GetOptionValueV4Request) (*GetOptionValueV4Response, error)

	// The method R_DhcpV4RemoveOptionValue removes the option value for the specified policy.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist on the DHCP server.                     |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E2A ERROR_DHCP_OPTION_NOT_PRESENT | The specified option definition does not exist on the DHCP server database, or   |
	//	|                                          | no value is set for the specified option ID on the specified policy.             |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND    | The class name being used is unknown or incorrect.                               |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8F ERROR_DHCP_POLICY_NOT_PRESENT | The specified policy name does not exist.                                        |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 104.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	RemoveOptionValueV4(context.Context, *RemoveOptionValueV4Request) (*RemoveOptionValueV4Response, error)

	// The method R_DhcpV4GetAllOptionValues gets all the server level policy or scope level
	// policy options configured. The memory for the Values parameter is allocated by this
	// method and can be freed by the caller by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or any generic failure.
	//
	//	+------------------------------------------+--------------------------------------------------------------+
	//	|                  RETURN                  |                                                              |
	//	|                VALUE/CODE                |                         DESCRIPTION                          |
	//	|                                          |                                                              |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                     |
	//	+------------------------------------------+--------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist on the DHCP server. |
	//	+------------------------------------------+--------------------------------------------------------------+
	//
	// The opnum field value for this method is 105.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetAllOptionValuesV4(context.Context, *GetAllOptionValuesV4Request) (*GetAllOptionValuesV4Response, error)

	// The R_DhcpV4QueryPolicyEnforcement method is used to retrieve the state (enabled/disabled)
	// of policy enforcement on the server or the specified IPv4 subnet.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP specific failure, which takes a value between
	// 20000 and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------+
	//	|                  RETURN                  |                                           |
	//	|                VALUE/CODE                |                DESCRIPTION                |
	//	|                                          |                                           |
	//	+------------------------------------------+-------------------------------------------+
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                  |
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist. |
	//	+------------------------------------------+-------------------------------------------+
	//
	// The opnum field value for this method is 106.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	QueryPolicyEnforcementV4(context.Context, *QueryPolicyEnforcementV4Request) (*QueryPolicyEnforcementV4Response, error)

	// The R_DhcpV4SetPolicyEnforcement method is used to set the state (enable/disable)
	// of policy enforcement of the server or the specified IPv4 subnet.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP specific failure, which takes a value between
	// 20000 and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------+
	//	|                  RETURN                  |                                           |
	//	|                VALUE/CODE                |                DESCRIPTION                |
	//	|                                          |                                           |
	//	+------------------------------------------+-------------------------------------------+
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                  |
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist. |
	//	+------------------------------------------+-------------------------------------------+
	//
	// The opnum field value for this method is 107.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetPolicyEnforcementV4(context.Context, *SetPolicyEnforcementV4Request) (*SetPolicyEnforcementV4Response, error)

	// The R_DhcpV4CreatePolicy method creates the policy according to the data specified
	// in the policy data structure.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                        RETURN                        |                                                                                  |
	//	|                      VALUE/CODE                      |                                   DESCRIPTION                                    |
	//	|                                                      |                                                                                  |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                             | The call was successful.                                                         |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT             | The specified IPv4 subnet does not exist.                                        |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8C ERROR_DHCP_RANGE_INVALID_IN_SERVER_POLICY | A policy range has been specified for a server level policy.                     |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8D ERROR_DHCP_INVALID_POLICY_EXPRESSION      | The specified conditions or expressions of the policy are invalid.               |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8B ERROR_DHCP_POLICY_RANGE_BAD               | The specified policy IP range is not contained within the IP address range of    |
	//	|                                                      | the scope, or the specified policy IP range is invalid.                          |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E89 ERROR_DHCP_POLICY_EXISTS                  | The specified policy name exists at the specified level (server or scope).       |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8A ERROR_DHCP_POLICY_RANGE_EXISTS            | The specified policy IP range overlaps the policy IP ranges of an existing       |
	//	|                                                      | policy at the specified scope.                                                   |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8E ERROR_DHCP_INVALID_PROCESSING_ORDER       | The specified processing order is greater than the maximum processing order of   |
	//	|                                                      | the existing policies at the specified level (server or scope).                  |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND                | The vendor class or user class reference in the conditions of the policy does    |
	//	|                                                      | not exist.                                                                       |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004EAC ERROR_DHCP_POLICY_FQDN_RANGE_UNSUPPORTED  | Ranges are not allowed to be set on the given policy.                            |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 108.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	CreatePolicyV4(context.Context, *CreatePolicyV4Request) (*CreatePolicyV4Response, error)

	// The R_DhcpV4GetPolicy method returns the specified policy. The memory for the Policy
	// structure is allocated by this method and can be freed by the caller by using the
	// function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP specific failure, which takes a value between
	// 20000 and 20099, or any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------+
	//	|                  RETURN                  |                                           |
	//	|                VALUE/CODE                |                DESCRIPTION                |
	//	|                                          |                                           |
	//	+------------------------------------------+-------------------------------------------+
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                  |
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist. |
	//	+------------------------------------------+-------------------------------------------+
	//
	// The opnum field value for this method is 109.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	GetPolicyV4(context.Context, *GetPolicyV4Request) (*GetPolicyV4Response, error)

	// The R_DhcpV4SetPolicy method modifies the specified DHCPv4 policy.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                        RETURN                        |                                                                                  |
	//	|                      VALUE/CODE                      |                                   DESCRIPTION                                    |
	//	|                                                      |                                                                                  |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                             | The call was successful.                                                         |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT             | The specified IPv4 subnet does not exist.                                        |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8C ERROR_DHCP_RANGE_INVALID_IN_SERVER_POLICY | A policy range has been specified for a server level policy.                     |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8D ERROR_DHCP_INVALID_POLICY_EXPRESSION      | The specified conditions or expressions of the policy are invalid.               |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8B ERROR_DHCP_POLICY_RANGE_BAD               | The specified policy range is not contained within the IP address range of the   |
	//	|                                                      | scope, or the specified policy range is invalid.                                 |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E89 ERROR_DHCP_POLICY_NOT_FOUND               | The specified policy name does not exist at the specified level (server or       |
	//	|                                                      | scope).                                                                          |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8A ERROR_DHCP_POLICY_RANGE_EXISTS            | The specified policy range overlaps the policy ranges of an existing policy at   |
	//	|                                                      | the specified scope.                                                             |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8E ERROR_DHCP_INVALID_PROCESSING_ORDER       | The specified processing order is greater than the maximum processing order of   |
	//	|                                                      | the existing policies at the specified level (server or scope).                  |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E4C ERROR_DHCP_CLASS_NOT_FOUND                | The vendor class or user class reference in the conditions of the policy does    |
	//	|                                                      | not exist.                                                                       |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004EA9 ERROR_DHCP_POLICY_EDIT_FQDN_UNSUPPORTED   | A FQDN-based condition is being added to a policy that has ranges or options     |
	//	|                                                      | configured.                                                                      |
	//	+------------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 110.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetPolicyV4(context.Context, *SetPolicyV4Request) (*SetPolicyV4Response, error)

	// The R_DhcpV4DeletePolicy method deletes the specified policy.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+------------------------------------------+-------------------------------------------+
	//	|                  RETURN                  |                                           |
	//	|                VALUE/CODE                |                DESCRIPTION                |
	//	|                                          |                                           |
	//	+------------------------------------------+-------------------------------------------+
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                  |
	//	+------------------------------------------+-------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist. |
	//	+------------------------------------------+-------------------------------------------+
	//
	// The opnum field value for this method is 111.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	DeletePolicyV4(context.Context, *DeletePolicyV4Request) (*DeletePolicyV4Response, error)

	// The method R_DhcpV4EnumPolicies returns an enumerated list of all configured server
	// level policies or scope level policies. The caller of this function can free the
	// memory pointed to by the EnumInfo parameter by calling the function midl_user_free
	// (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or any generic failure.
	//
	//	+--------------------------------+-------------------------------------------------+
	//	|             RETURN             |                                                 |
	//	|           VALUE/CODE           |                   DESCRIPTION                   |
	//	|                                |                                                 |
	//	+--------------------------------+-------------------------------------------------+
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                        |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA     | There are more elements available to enumerate. |
	//	+--------------------------------+-------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate.   |
	//	+--------------------------------+-------------------------------------------------+
	//
	// The opnum field value for this method is 112.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumPoliciesV4(context.Context, *EnumPoliciesV4Request) (*EnumPoliciesV4Response, error)

	// The R_DhcpV4AddPolicyRange method adds an IP address range to a policy.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+-----------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                       RETURN                        |                                                                                  |
	//	|                     VALUE/CODE                      |                                   DESCRIPTION                                    |
	//	|                                                     |                                                                                  |
	//	+-----------------------------------------------------+----------------------------------------------------------------------------------+
	//	+-----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                            | The call was successful.                                                         |
	//	+-----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT            | The specified IPv4 subnet does not exist.                                        |
	//	+-----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8F ERROR_DHCP_POLICY_NOT_FOUND              | The specified policy does not exist.                                             |
	//	+-----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8B ERROR_DHCP_POLICY_RANGE_BAD              | The specified policy IP range is not contained within the IP address range of    |
	//	|                                                     | the scope, or the specified policy IP range is not valid.                        |
	//	+-----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8A ERROR_DHCP_POLICY_RANGE_EXISTS           | The specified policy IP range overlaps one of the policy IP address ranges       |
	//	|                                                     | specified.                                                                       |
	//	+-----------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004EA7 ERROR_DHCP_POLICY_FQDN_RANGE_UNSUPPORTED | Ranges are not allowed to be added to the given policy.                          |
	//	+-----------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 113.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	AddPolicyRangeV4(context.Context, *AddPolicyRangeV4Request) (*AddPolicyRangeV4Response, error)

	// The R_DhcpV4RemovePolicyRange method removes the specified IP address range from
	// the list of IP address ranges of the policy.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	|                  RETURN                  |                                                                                  |
	//	|                VALUE/CODE                |                                   DESCRIPTION                                    |
	//	|                                          |                                                                                  |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call was successful.                                                         |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified IPv4 subnet does not exist.                                        |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8F ERROR_DHCP_POLICY_NOT_FOUND   | The specified policy does not exist.                                             |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00004E8B ERROR_DHCP_POLICY_RANGE_BAD   | The specified policy range is not contained within the IP address range of the   |
	//	|                                          | scope.                                                                           |
	//	+------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 114.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	RemovePolicyRangeV4(context.Context, *RemovePolicyRangeV4Request) (*RemovePolicyRangeV4Response, error)

	// The R_DhcpV4EnumSubnetClients method is used to retrieve all DHCPv4 clients serviced
	// on the specified IPv4 subnet. The information also includes the link-layer filter
	// status info for the DHCPv4 client and the policy, if any, that resulted in the specific
	// IPv4 address assignment.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+---------------------------------+-------------------------------------------------------------+
	//	|             RETURN              |                                                             |
	//	|           VALUE/CODE            |                         DESCRIPTION                         |
	//	|                                 |                                                             |
	//	+---------------------------------+-------------------------------------------------------------+
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS        | The call was successful.                                    |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA      | More client lease records are available to enumerate.       |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS  | No more client lease records are left to enumerate.         |
	//	+---------------------------------+-------------------------------------------------------------+
	//	| 0x00004E2D ERROR_DHCP_JET_ERROR | An error occurred while accessing the DHCP server database. |
	//	+---------------------------------+-------------------------------------------------------------+
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnetClientsV4(context.Context, *EnumSubnetClientsV4Request) (*EnumSubnetClientsV4Response, error)

	// The R_DhcpV6SetStatelessStoreParams method modifies the configuration settings for
	// DHCPv6 stateless client inventory at the server or scope level.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+------------------------------------------+------------------------------------------------------+
	//	|                  RETURN                  |                                                      |
	//	|                VALUE/CODE                |                     DESCRIPTION                      |
	//	|                                          |                                                      |
	//	+------------------------------------------+------------------------------------------------------+
	//	+------------------------------------------+------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call completed successfully.                     |
	//	+------------------------------------------+------------------------------------------------------+
	//	| 0x00020005 ERROR_DHCP_SUBNET_NOT_PRESENT | The IPv6 subnet does not exist on the DHCPv6 server. |
	//	+------------------------------------------+------------------------------------------------------+
	//
	// The opnum field value for this method is 116.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	SetStatelessStoreParamsV6(context.Context, *SetStatelessStoreParamsV6Request) (*SetStatelessStoreParamsV6Response, error)

	// The R_DhcpV6GetStatelessStoreParams method retrieves the current DHCPv6 stateless
	// client inventory-related configuration setting at the server or scope level. The
	// caller of this function can free the memory pointed to by the Params parameter by
	// calling the function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+------------------------------------------+------------------------------------------------------+
	//	|                  RETURN                  |                                                      |
	//	|                VALUE/CODE                |                     DESCRIPTION                      |
	//	|                                          |                                                      |
	//	+------------------------------------------+------------------------------------------------------+
	//	+------------------------------------------+------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call completed successfully.                     |
	//	+------------------------------------------+------------------------------------------------------+
	//	| 0x00020005 ERROR_DHCP_SUBNET_NOT_PRESENT | The IPv6 subnet does not exist on the DHCPv6 server. |
	//	+------------------------------------------+------------------------------------------------------+
	//
	// The opnum field value for this method is 117.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetStatelessStoreParamsV6(context.Context, *GetStatelessStoreParamsV6Request) (*GetStatelessStoreParamsV6Response, error)

	// The R_DhcpV6GetStatelessStatistics method is used to retrieve the statistics of the
	// DHCPv6 stateless server. The caller of this function can free the memory pointed
	// to by the StatelessStats parameter and its ScopeStats member array by calling the
	// function midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+--------------------------+----------------------------------+
	//	|          RETURN          |                                  |
	//	|        VALUE/CODE        |           DESCRIPTION            |
	//	|                          |                                  |
	//	+--------------------------+----------------------------------+
	//	+--------------------------+----------------------------------+
	//	| 0x00000000 ERROR_SUCCESS | The call completed successfully. |
	//	+--------------------------+----------------------------------+
	//
	// The opnum field value for this method is 118.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetStatelessStatisticsV6(context.Context, *GetStatelessStatisticsV6Request) (*GetStatelessStatisticsV6Response, error)

	// The R_DhcpV4EnumSubnetReservations method enumerates all the reservation information
	// on the DHCPv4 server for a given IPv4 subnet address. The caller of this function
	// can free the memory pointed to by the EnumElementInfo parameter by calling the function
	// midl_user_free (section 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20123, or any generic failure.
	//
	//	+--------------------------------+--------------------------------------------------+
	//	|             RETURN             |                                                  |
	//	|           VALUE/CODE           |                   DESCRIPTION                    |
	//	|                                |                                                  |
	//	+--------------------------------+--------------------------------------------------+
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS       | The call was successful.                         |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x000000EA ERROR_MORE_DATA     | There are more elements available to enumerate.  |
	//	+--------------------------------+--------------------------------------------------+
	//	| 0x00000103 ERROR_NO_MORE_ITEMS | There are no more elements left to enumerate.    |
	//	+--------------------------------+--------------------------------------------------+
	//	| ERROR_DHCP_SUBNET_NOT_PRESENT  | IPv4 subnet does not exist on the DHCPv4 server. |
	//	+--------------------------------+--------------------------------------------------+
	//
	// The opnum field value for this method is 119.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	EnumSubnetReservationsV4(context.Context, *EnumSubnetReservationsV4Request) (*EnumSubnetReservationsV4Response, error)

	// The R_DhcpV4GetFreeIPAddress method retrieves the list of IPv4 addresses available
	// to be leased out to the clients. The caller of this function can free the memory
	// pointed to by the IPAddrList parameter by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20123, or any generic failure.
	//
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                     RETURN                     |                                                                                  |
	//	|                   VALUE/CODE                   |                                   DESCRIPTION                                    |
	//	|                                                |                                                                                  |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                       | The call completed successfully.                                                 |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000103 ERROR_FILE_NOT_FOUND                | No more elements are left to enumerate.                                          |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00020126 ERROR_DHCP_REACHED_END_OF_SELECTION | The specified DHCP server has reached the end of the selected range while        |
	//	|                                                | finding the free IP addresses.                                                   |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 120.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetFreeIPAddressV4(context.Context, *GetFreeIPAddressV4Request) (*GetFreeIPAddressV4Response, error)

	// The R_DhcpV6GetFreeIPAddress method retrieves the list of IPv6 addresses available
	// to be leased out to the clients. The caller of this function can free the memory
	// pointed to by the IPAddrList parameter by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20123, or to any generic failure.
	//
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	|                     RETURN                     |                                                                                  |
	//	|                   VALUE/CODE                   |                                   DESCRIPTION                                    |
	//	|                                                |                                                                                  |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                       | The call completed successfully.                                                 |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00020005 ERROR_DHCP_SUBNET_NOT_PRESENT       | The IPv6 subnet does not exist on the DHCPv6 server.                             |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//	| 0x00020126 ERROR_DHCP_REACHED_END_OF_SELECTION | The specified DHCP server has reached the end of the selected range while        |
	//	|                                                | finding the free IP addresses.                                                   |
	//	+------------------------------------------------+----------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 121.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	GetFreeIPAddressV6(context.Context, *GetFreeIPAddressV6Request) (*GetFreeIPAddressV6Response, error)

	// The R_DhcpV4CreateClientInfo method creates a DHCPv4 client lease record on the DHCP
	// server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+------------------------------------------+----------------------------------------------------+
	//	|                  RETURN                  |                                                    |
	//	|                VALUE/CODE                |                    DESCRIPTION                     |
	//	|                                          |                                                    |
	//	+------------------------------------------+----------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call completed successfully.                   |
	//	+------------------------------------------+----------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified subnet does not exist.               |
	//	+------------------------------------------+----------------------------------------------------+
	//	| 0x00004E2E ERROR_DHCP_CLIENT_EXISTS      | The specified client already exists on the server. |
	//	+------------------------------------------+----------------------------------------------------+
	//
	// The opnum field value for this method is 122.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	CreateClientInfoV4(context.Context, *CreateClientInfoV4Request) (*CreateClientInfoV4Response, error)

	// The R_DhcpV4GetClientInfo method retrieves DHCPv4 client lease record information
	// from the DHCPv4 server database. The information also includes the link-layer filter
	// status information for the DHCPv4 client and the policy, if any, that resulted in
	// the specific IPv4 address assignment. The caller of this function can free the memory
	// pointed to by the ClientInfo parameter by calling the function midl_user_free (section
	// 3).
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+--------------------------------------+-----------------------------------------+
	//	|                RETURN                |                                         |
	//	|              VALUE/CODE              |               DESCRIPTION               |
	//	|                                      |                                         |
	//	+--------------------------------------+-----------------------------------------+
	//	+--------------------------------------+-----------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS             | The call completed successfully.        |
	//	+--------------------------------------+-----------------------------------------+
	//	| 0x00004E30 ERROR_DHCP_INVALID_CLIENT | The specified DHCP client is not valid. |
	//	+--------------------------------------+-----------------------------------------+
	//
	// The opnum field value for this method is 123.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol [MS-RPCE].
	GetClientInfoV4(context.Context, *GetClientInfoV4Request) (*GetClientInfoV4Response, error)

	// The R_DhcpV6CreateClientInfo method creates a DHCPv6 client lease record on the DHCP
	// server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+------------------------------------------+----------------------------------------------------+
	//	|                  RETURN                  |                                                    |
	//	|                VALUE/CODE                |                    DESCRIPTION                     |
	//	|                                          |                                                    |
	//	+------------------------------------------+----------------------------------------------------+
	//	+------------------------------------------+----------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                 | The call completed successfully.                   |
	//	+------------------------------------------+----------------------------------------------------+
	//	| 0x00004E25 ERROR_DHCP_SUBNET_NOT_PRESENT | The specified subnet does not exist.               |
	//	+------------------------------------------+----------------------------------------------------+
	//	| 0x00004E2E ERROR_DHCP_CLIENT_EXISTS      | The specified client already exists on the server. |
	//	+------------------------------------------+----------------------------------------------------+
	//
	// The opnum field value for this method is 124.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	CreateClientInfoV6(context.Context, *CreateClientInfoV6Request) (*CreateClientInfoV6Response, error)

	// The R_DhcpV4FailoverGetAddressStatus method queries the current address status for
	// an address belonging to a subnet that is part of a failover relationship on the DHCP
	// server.
	//
	// Return Values: A 32-bit unsigned integer value that indicates return status. A return
	// value of ERROR_SUCCESS (0x00000000) indicates that the operation was completed successfully.
	// Otherwise, it contains a Win32 error code, as specified in [MS-ERREF]. This error
	// code value can correspond to a DHCP-specific failure, which takes a value between
	// 20000 and 20099, or to any generic failure.
	//
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	|                       RETURN                       |                                                                                |
	//	|                     VALUE/CODE                     |                                  DESCRIPTION                                   |
	//	|                                                    |                                                                                |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000000 ERROR_SUCCESS                           | The call completed successfully.                                               |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00000057 ERROR_INVALID_PARAMETER                 | An invalid parameter is specified in the Address parameter.                    |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//	| 0x00004E94 ERROR_DHCP_FO_SCOPE_NOT_IN_RELATIONSHIP | The subnet associated with the address is not part of a failover relationship. |
	//	+----------------------------------------------------+--------------------------------------------------------------------------------+
	//
	// The opnum field value for this method is 125.
	//
	// Exceptions Thrown: No exceptions are thrown beyond those thrown by the underlying
	// RPC protocol specified in [MS-RPCE].
	FailoverGetAddressStatusV4(context.Context, *FailoverGetAddressStatusV4Request) (*FailoverGetAddressStatusV4Response, error)

	// The R_DhcpV4CreatePolicyEx method creates the policy according to the data specified
	// in the policy data structure.
	//
	// The R_DhcpV4CreatePolicyEx method is an extension of the R_DhcpV4CreatePolicy (Opnum
	// 108) (section 3.2.4.109) method, where a DHCP_POLICY_EX (section 2.2.1.2.121) structure
	// is specified for the pPolicy parameter, rather than a DHCP_POLICY (section 2.2.1.2.110)
	// structure. The structure contains the members of the policy to be created.
	//
	// Using the extension method, a list of DHCP_PROPERTY (section 2.2.1.2.117) elements
	// can be specified that can be associated with the given policy when creating the policy.
	//
	// Return Values: As specified in R_DhcpV4CreatePolicy (Opnum 108).
	//
	// The opnum field value for this method is 126.
	//
	// The remainder of the processing behavior for this method is as defined for the R_DhcpV4CreatePolicy
	// (Opnum 108) method.
	CreatePolicyExV4(context.Context, *CreatePolicyExV4Request) (*CreatePolicyExV4Response, error)

	// The R_DhcpV4GetPolicyEx method returns the specified policy. The memory for the Policy
	// structure is allocated by the method and can be freed by the caller by using the
	// midl_user_free function (section 3).
	//
	// The R_DhcpV4GetPolicyEx method is an extension of the R_DhcpV4GetPolicy (Opnum 109)
	// (section 3.2.4.110) method, where a DHCP_POLICY_EX (section 2.2.1.2.121) structure
	// is queried, rather than a DHCP_POLICY (section 2.2.1.2.110) structure. The structure
	// returns a list of DHCP_PROPERTY (section 2.2.1.2.117) elements that can be associated
	// with the given policy.
	//
	// Return Values: As specified in R_DhcpV4CreatePolicy (Opnum 108).
	//
	// The opnum field value for this method is 127.
	//
	// The remainder of the processing behavior for this method is as defined for the R_DhcpV4GetPolicy
	// (Opnum 109) method.
	GetPolicyExV4(context.Context, *GetPolicyExV4Request) (*GetPolicyExV4Response, error)

	// The R_DhcpV4SetPolicyEx method modifies the specified policy.
	//
	// The method is an extension of the R_DhcpV4SetPolicy (Opnum 110) (section 3.2.4.111)
	// method, where the method specifies a DHCP_POLICY_EX (section 2.2.1.2.121) structure
	// rather than a DHCP_POLICY (section 2.2.1.2.110) structure. The structure contains
	// a list of DHCP_PROPERTY (section 2.2.1.2.117) elements that can be updated for the
	// policy.
	//
	// Return Values: As specified in R_DhcpV4SetPolicy (Opnum 110).
	//
	// The opnum field value for this method is 128.
	//
	// The remainder of the processing behavior for this method is as defined for the R_DhcpV4SetPolicy
	// (Opnum 110) method, except as follows:
	//
	// * The FieldsModified parameter can also be set to the DhcpUpdatePolicyDnsSuffix value
	// of the DHCP_POLICY_FIELDS_TO_UPDATE (section 2.2.1.1.21) ( 49abe631-1c6b-4711-8337-b4b2bdf90b00
	// ) enumeration.
	//
	// * If the FieldsModified parameter is set to DhcpUpdatePolicyDnsSuffix, the *R_DhcpV4SetPolicyEx*
	// method searches for the property with an ID value of DhcpPropIdPolicyDnsSuffix and
	// Type value of DhcpPropTypeString. If such a property is located, the *R_DhcpV4SetPolicyEx*
	// method validates that the string length of the property value does not exceed 255
	// characters. If the length is exceeded, the *R_DhcpV4SetPolicyEx* method returns ERROR_INVALID_PARAMETER.
	//
	// * The *R_DhcpV4SetPolicyEx* method updates the server or scope level <DHCPv4Policy>
	// ADM element retrieved earlier according to the following:
	//
	// In addition to steps 1 through 5 specified in *R_DhcpV4SetPolicy* (Opnum 110), the
	// *R_DhcpV4SetPolicyEx* method adds the following instruction:
	//
	// * If the *DhcpUpdatePolicyDnsSuffix* enumeration value is set in the FieldsModified
	// parameter, update the DNSSuffix of the policy in the <DHCPv4Policy.DnsSuffix> ADM
	// element. If no such property exists in the list of properties with an ID value equal
	// to DhcpPropIdPolicyDnsSuffix and a Type value equal to DhcpPropTypeString, or if
	// the StringValue of the property is NULL or of zero length, then the <DHCPv4Policy.DnsSuffix>
	// ADM element is cleared; otherwise, the ADM element is set to the StringValue of the
	// property.
	//
	// * The FieldsModified parameter is set to any value other than DhcpUpdatePolicyName,
	// DhcpUpdatePolicyOrder, DhcpUpdatePolicyExpr, DhcpUpdatePolicyRanges, DhcpUpdatePolicyDescr,
	// DhcpUpdatePolicyStatus, or DhcpUpdatePolicyDnsSuffix, as defined in *DHCP_POLICY_FIELDS_TO_UPDATE*
	// enumeration.
	//
	// * The *R_DhcpV4SetPolicyEx* method returns ERROR_INVALID_PARAMETER.
	SetPolicyExV4(context.Context, *SetPolicyExV4Request) (*SetPolicyExV4Response, error)

	// The R_DhcpV4EnumPoliciesEx method returns an enumerated list of all configured server
	// level or scope level policies. The caller of this method can free the memory pointed
	// to by the EnumInfo parameter by calling the midl_user_free function (section 3).
	//
	// The R_DhcpV4EnumPoliciesEx method is an extension of the R_DhcpV4EnumPolicies (Opnum
	// 112) (section 3.2.4.130) method, where an array of DHCP_POLICY_EX (section 2.2.1.2.121)
	// structures is enumerated, rather than an array of DHCP_POLICY (section 2.2.1.2.110)
	// structures. Each DHCP_POLICY_EX structure contains a list of DHCP_PROPERTY (section
	// 2.2.1.2.117) elements that are associated with the given policy.
	//
	// Return Values: As specified in R_DhcpV4EnumPolicies (Opnum 112).
	//
	// The opnum field value for this method is 129.
	//
	// The remainder of the processing behavior for this method is as defined for the R_DhcpV4EnumPolicies
	// method, except as follows:
	//
	// * No filtering is applied to the enumerated list of configured server-level or scope-level
	// policies returned by the *R_DhcpV4EnumPoliciesEx* method.
	EnumPoliciesExV4(context.Context, *EnumPoliciesExV4Request) (*EnumPoliciesExV4Response, error)

	// The R_DhcpV4EnumSubnetClientsEx method is used to retrieve all DHCPv4 clients serviced
	// on the specified IPv4 subnet. The information retrieved also includes the link-layer
	// filter status for the DHCPv4 client and the policy, if any, that resulted in the
	// specific IPv4 address assignment.
	//
	// The R_DhcpV4EnumSubnetClientsEx method is an extension of the R_DhcpV4EnumSubnetClients
	// (Opnum 115) (section 3.2.4.116) method, where an array of  DHCP_CLIENT_INFO_EX (section
	// 2.2.1.2.119) structures is enumerated, rather than an array of DHCP_CLIENT_INFO_PB
	// (section 2.2.1.2.115) structures. Each DHCP_CLIENT_INFO_EX structure contains a list
	// of DHCP_PROPERTY (section 2.2.1.2.117) elements that are associated with the given
	// subnet client.
	//
	// Return Values: As specified in R_DhcpV4EnumSubnetClients (Opnum 115).
	//
	// The opnum field value for this method is 130.
	//
	// The remainder of the processing behavior for this method is as defined for the R_DhcpV4EnumSubnetClients
	// method.
	EnumSubnetClientsExV4(context.Context, *EnumSubnetClientsExV4Request) (*EnumSubnetClientsExV4Response, error)

	// The R_DhcpV4CreateClientInfoEx method creates a DHCPv4 client lease record on the
	// DHCP server.
	//
	// The R_DhcpV4CreateClientInfoEx method is an extension of the R_DhcpV4CreateClientInfo
	// (Opnum 122) (section 3.2.4.132) method, where a DHCP_CLIENT_INFO_EX (section 2.2.1.2.119)
	// structure is specified, rather than a DHCP_CLIENT_INFO_PB (section 2.2.1.2.119) structure.
	// The structure contains a list of DHCP_PROPERTY (section 2.2.1.2.117) elements that
	// can be associated with the given DHCPv4 client.
	CreateClientInfoExV4(context.Context, *CreateClientInfoExV4Request) (*CreateClientInfoExV4Response, error)

	// The R_DhcpV4GetClientInfoEx method retrieves DHCPv4 client lease record information
	// from the DHCPv4 server database. The retrieved information also includes the link-layer
	// filter status information for the DHCPv4 client and the policy, if any, that resulted
	// in the specific IPv4 address assignment. The caller of this method can free the memory
	// pointed to by the ClientInfo parameter by calling the midl_user_free function (section
	// 3).
	//
	// The R_DhcpV4GetClientInfoEx method is an extension of the R_DhcpV4GetClientInfo (Opnum
	// 123) (section 3.2.4.124) method, where a DHCP_CLIENT_INFO_EX (section 2.2.1.2.119)
	// structure is queried, rather than a DHCP_CLIENT_INFO_PB (section 2.2.1.2.115) structure.
	// The structure returns a list of DHCP_PROPERTY (section 2.2.1.2.117) elements that
	// can be associated with the given DHCPv4 client.
	//
	// Return Values: As specified in R_DhcpV4GetClientInfo (Opnum 123).
	//
	// The opnum field value for this method is 132.
	//
	// The remainder of the processing behavior for this method is as defined for the R_DhcpV4GetClientInfo
	// method.
	GetClientInfoExV4(context.Context, *GetClientInfoExV4Request) (*GetClientInfoExV4Response, error)
}

func RegisterDhcpsrv2Server(conn dcerpc.Conn, o Dhcpsrv2Server, opts ...dcerpc.Option) {
	conn.RegisterServer(NewDhcpsrv2ServerHandle(o), append(opts, dcerpc.WithAbstractSyntax(Dhcpsrv2SyntaxV1_0))...)
}

func NewDhcpsrv2ServerHandle(o Dhcpsrv2Server) dcerpc.ServerHandle {
	return func(ctx context.Context, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
		return Dhcpsrv2ServerHandle(ctx, o, opNum, r)
	}
}

func Dhcpsrv2ServerHandle(ctx context.Context, o Dhcpsrv2Server, opNum int, r ndr.Reader) (dcerpc.Operation, error) {
	switch opNum {
	case 0: // R_DhcpEnumSubnetClientsV5
		in := &EnumSubnetClientsV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetClientsV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 1: // R_DhcpSetMScopeInfo
		in := &SetMScopeInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMScopeInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 2: // R_DhcpGetMScopeInfo
		in := &GetMScopeInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMScopeInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 3: // R_DhcpEnumMScopes
		in := &EnumMScopesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumMScopes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 4: // R_DhcpAddMScopeElement
		in := &AddMScopeElementRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddMScopeElement(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 5: // R_DhcpEnumMScopeElements
		in := &EnumMScopeElementsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumMScopeElements(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 6: // R_DhcpRemoveMScopeElement
		in := &RemoveMScopeElementRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveMScopeElement(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 7: // R_DhcpDeleteMScope
		in := &DeleteMScopeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteMScope(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 8: // R_DhcpScanMDatabase
		in := &ScanMDatabaseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ScanMDatabase(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 9: // R_DhcpCreateMClientInfo
		in := &CreateMClientInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateMClientInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 10: // R_DhcpSetMClientInfo
		in := &SetMClientInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetMClientInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 11: // R_DhcpGetMClientInfo
		in := &GetMClientInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMClientInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 12: // R_DhcpDeleteMClientInfo
		in := &DeleteMClientInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteMClientInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 13: // R_DhcpEnumMScopeClients
		in := &EnumMScopeClientsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumMScopeClients(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 14: // R_DhcpCreateOptionV5
		in := &CreateOptionV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateOptionV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 15: // R_DhcpSetOptionInfoV5
		in := &SetOptionInfoV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOptionInfoV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 16: // R_DhcpGetOptionInfoV5
		in := &GetOptionInfoV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOptionInfoV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 17: // R_DhcpEnumOptionsV5
		in := &EnumOptionsV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumOptionsV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 18: // R_DhcpRemoveOptionV5
		in := &RemoveOptionV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveOptionV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 19: // R_DhcpSetOptionValueV5
		in := &SetOptionValueV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOptionValueV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 20: // R_DhcpSetOptionValuesV5
		in := &SetOptionValuesV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOptionValuesV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 21: // R_DhcpGetOptionValueV5
		in := &GetOptionValueV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOptionValueV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 22: // R_DhcpEnumOptionValuesV5
		in := &EnumOptionValuesV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumOptionValuesV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 23: // R_DhcpRemoveOptionValueV5
		in := &RemoveOptionValueV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveOptionValueV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 24: // R_DhcpCreateClass
		in := &CreateClassRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateClass(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 25: // R_DhcpModifyClass
		in := &ModifyClassRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ModifyClass(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 26: // R_DhcpDeleteClass
		in := &DeleteClassRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteClass(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 27: // R_DhcpGetClassInfo
		in := &GetClassInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClassInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 28: // R_DhcpEnumClasses
		in := &EnumClassesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumClasses(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 29: // R_DhcpGetAllOptions
		in := &GetAllOptionsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllOptions(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 30: // R_DhcpGetAllOptionValues
		in := &GetAllOptionValuesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllOptionValues(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 31: // R_DhcpGetMCastMibInfo
		in := &GetMCastMIBInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMCastMIBInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 32: // R_DhcpAuditLogSetParams
		in := &AuditLogSetParamsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AuditLogSetParams(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 33: // R_DhcpAuditLogGetParams
		in := &AuditLogGetParamsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AuditLogGetParams(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 34: // R_DhcpServerQueryAttribute
		in := &ServerQueryAttributeRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerQueryAttribute(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 35: // R_DhcpServerQueryAttributes
		in := &ServerQueryAttributesRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerQueryAttributes(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 36: // R_DhcpServerRedoAuthorization
		in := &ServerRedoAuthorizationRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerRedoAuthorization(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 37: // R_DhcpAddSubnetElementV5
		in := &AddSubnetElementV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddSubnetElementV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 38: // R_DhcpEnumSubnetElementsV5
		in := &EnumSubnetElementsV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetElementsV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 39: // R_DhcpRemoveSubnetElementV5
		in := &RemoveSubnetElementV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveSubnetElementV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 40: // R_DhcpGetServerBindingInfo
		in := &GetServerBindingInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServerBindingInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 41: // R_DhcpSetServerBindingInfo
		in := &SetServerBindingInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetServerBindingInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 42: // R_DhcpQueryDnsRegCredentials
		in := &QueryDNSRegCredentialsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryDNSRegCredentials(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 43: // R_DhcpSetDnsRegCredentials
		in := &SetDNSRegCredentialsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDNSRegCredentials(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 44: // R_DhcpBackupDatabase
		in := &BackupDatabaseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.BackupDatabase(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 45: // R_DhcpRestoreDatabase
		in := &RestoreDatabaseRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RestoreDatabase(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 46: // R_DhcpGetServerSpecificStrings
		in := &GetServerSpecificStringsRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServerSpecificStrings(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 47: // R_DhcpCreateOptionV6
		in := &CreateOptionV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateOptionV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 48: // R_DhcpSetOptionInfoV6
		in := &SetOptionInfoV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOptionInfoV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 49: // R_DhcpGetOptionInfoV6
		in := &GetOptionInfoV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOptionInfoV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 50: // R_DhcpEnumOptionsV6
		in := &EnumOptionsV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumOptionsV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 51: // R_DhcpRemoveOptionV6
		in := &RemoveOptionV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveOptionV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 52: // R_DhcpSetOptionValueV6
		in := &SetOptionValueV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOptionValueV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 53: // R_DhcpEnumOptionValuesV6
		in := &EnumOptionValuesV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumOptionValuesV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 54: // R_DhcpRemoveOptionValueV6
		in := &RemoveOptionValueV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveOptionValueV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 55: // R_DhcpGetAllOptionsV6
		in := &GetAllOptionsV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllOptionsV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 56: // R_DhcpGetAllOptionValuesV6
		in := &GetAllOptionValuesV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllOptionValuesV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 57: // R_DhcpCreateSubnetV6
		in := &CreateSubnetV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateSubnetV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 58: // R_DhcpEnumSubnetsV6
		in := &EnumSubnetsV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetsV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 59: // R_DhcpAddSubnetElementV6
		in := &AddSubnetElementV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddSubnetElementV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 60: // R_DhcpEnumSubnetElementsV6
		in := &EnumSubnetElementsV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetElementsV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 61: // R_DhcpRemoveSubnetElementV6
		in := &RemoveSubnetElementV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveSubnetElementV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 62: // R_DhcpDeleteSubnetV6
		in := &DeleteSubnetV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteSubnetV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 63: // R_DhcpGetSubnetInfoV6
		in := &GetSubnetInfoV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubnetInfoV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 64: // R_DhcpEnumSubnetClientsV6
		in := &EnumSubnetClientsV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetClientsV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 65: // R_DhcpServerSetConfigV6
		in := &ServerSetConfigV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerSetConfigV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 66: // R_DhcpServerGetConfigV6
		in := &ServerGetConfigV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ServerGetConfigV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 67: // R_DhcpSetSubnetInfoV6
		in := &SetSubnetInfoV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubnetInfoV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 68: // R_DhcpGetMibInfoV6
		in := &GetMIBInfoV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMIBInfoV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 69: // R_DhcpGetServerBindingInfoV6
		in := &GetServerBindingInfoV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetServerBindingInfoV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 70: // R_DhcpSetServerBindingInfoV6
		in := &SetServerBindingInfoV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetServerBindingInfoV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 71: // R_DhcpSetClientInfoV6
		in := &SetClientInfoV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetClientInfoV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 72: // R_DhcpGetClientInfoV6
		in := &GetClientInfoV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClientInfoV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 73: // R_DhcpDeleteClientInfoV6
		in := &DeleteClientInfoV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteClientInfoV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 74: // R_DhcpCreateClassV6
		in := &CreateClassV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateClassV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 75: // R_DhcpModifyClassV6
		in := &ModifyClassV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.ModifyClassV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 76: // R_DhcpDeleteClassV6
		in := &DeleteClassV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteClassV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 77: // R_DhcpEnumClassesV6
		in := &EnumClassesV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumClassesV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 78: // R_DhcpGetOptionValueV6
		in := &GetOptionValueV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOptionValueV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 79: // R_DhcpSetSubnetDelayOffer
		in := &SetSubnetDelayOfferRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetSubnetDelayOffer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 80: // R_DhcpGetSubnetDelayOffer
		in := &GetSubnetDelayOfferRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetSubnetDelayOffer(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 81: // R_DhcpGetMibInfoV5
		in := &GetMIBInfoV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetMIBInfoV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 82: // R_DhcpAddFilterV4
		in := &AddFilterV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddFilterV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 83: // R_DhcpDeleteFilterV4
		in := &DeleteFilterV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeleteFilterV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 84: // R_DhcpSetFilterV4
		in := &SetFilterV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetFilterV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 85: // R_DhcpGetFilterV4
		in := &GetFilterV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFilterV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 86: // R_DhcpEnumFilterV4
		in := &EnumFilterV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumFilterV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 87: // R_DhcpSetDnsRegCredentialsV5
		in := &SetDNSRegCredentialsV5Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetDNSRegCredentialsV5(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 88: // R_DhcpEnumSubnetClientsFilterStatusInfo
		in := &EnumSubnetClientsFilterStatusInfoRequest{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetClientsFilterStatusInfo(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 89: // R_DhcpV4FailoverCreateRelationship
		in := &FailoverCreateRelationshipV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverCreateRelationshipV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 90: // R_DhcpV4FailoverSetRelationship
		in := &FailoverSetRelationshipV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverSetRelationshipV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 91: // R_DhcpV4FailoverDeleteRelationship
		in := &FailoverDeleteRelationshipV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverDeleteRelationshipV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 92: // R_DhcpV4FailoverGetRelationship
		in := &FailoverGetRelationshipV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverGetRelationshipV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 93: // R_DhcpV4FailoverEnumRelationship
		in := &FailoverEnumRelationshipV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverEnumRelationshipV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 94: // R_DhcpV4FailoverAddScopeToRelationship
		in := &FailoverAddScopeToRelationshipV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverAddScopeToRelationshipV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 95: // R_DhcpV4FailoverDeleteScopeFromRelationship
		in := &FailoverDeleteScopeFromRelationshipV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverDeleteScopeFromRelationshipV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 96: // R_DhcpV4FailoverGetScopeRelationship
		in := &FailoverGetScopeRelationshipV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverGetScopeRelationshipV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 97: // R_DhcpV4FailoverGetScopeStatistics
		in := &FailoverGetScopeStatisticsV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverGetScopeStatisticsV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 98: // R_DhcpV4FailoverGetClientInfo
		in := &FailoverGetClientInfoV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverGetClientInfoV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 99: // R_DhcpV4FailoverGetSystemTime
		in := &FailoverGetSystemTimeV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverGetSystemTimeV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 100: // R_DhcpV4FailoverTriggerAddrAllocation
		in := &FailoverTriggerAddrAllocationV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverTriggerAddrAllocationV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 101: // R_DhcpV4SetOptionValue
		in := &SetOptionValueV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOptionValueV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 102: // R_DhcpV4SetOptionValues
		in := &SetOptionValuesV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetOptionValuesV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 103: // R_DhcpV4GetOptionValue
		in := &GetOptionValueV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetOptionValueV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 104: // R_DhcpV4RemoveOptionValue
		in := &RemoveOptionValueV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemoveOptionValueV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 105: // R_DhcpV4GetAllOptionValues
		in := &GetAllOptionValuesV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetAllOptionValuesV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 106: // R_DhcpV4QueryPolicyEnforcement
		in := &QueryPolicyEnforcementV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.QueryPolicyEnforcementV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 107: // R_DhcpV4SetPolicyEnforcement
		in := &SetPolicyEnforcementV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPolicyEnforcementV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 108: // R_DhcpV4CreatePolicy
		in := &CreatePolicyV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreatePolicyV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 109: // R_DhcpV4GetPolicy
		in := &GetPolicyV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPolicyV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 110: // R_DhcpV4SetPolicy
		in := &SetPolicyV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPolicyV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 111: // R_DhcpV4DeletePolicy
		in := &DeletePolicyV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.DeletePolicyV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 112: // R_DhcpV4EnumPolicies
		in := &EnumPoliciesV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPoliciesV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 113: // R_DhcpV4AddPolicyRange
		in := &AddPolicyRangeV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.AddPolicyRangeV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 114: // R_DhcpV4RemovePolicyRange
		in := &RemovePolicyRangeV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.RemovePolicyRangeV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 115: // R_DhcpV4EnumSubnetClients
		in := &EnumSubnetClientsV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetClientsV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 116: // R_DhcpV6SetStatelessStoreParams
		in := &SetStatelessStoreParamsV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetStatelessStoreParamsV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 117: // R_DhcpV6GetStatelessStoreParams
		in := &GetStatelessStoreParamsV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetStatelessStoreParamsV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 118: // R_DhcpV6GetStatelessStatistics
		in := &GetStatelessStatisticsV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetStatelessStatisticsV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 119: // R_DhcpV4EnumSubnetReservations
		in := &EnumSubnetReservationsV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetReservationsV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 120: // R_DhcpV4GetFreeIPAddress
		in := &GetFreeIPAddressV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFreeIPAddressV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 121: // R_DhcpV6GetFreeIPAddress
		in := &GetFreeIPAddressV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetFreeIPAddressV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 122: // R_DhcpV4CreateClientInfo
		in := &CreateClientInfoV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateClientInfoV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 123: // R_DhcpV4GetClientInfo
		in := &GetClientInfoV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClientInfoV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 124: // R_DhcpV6CreateClientInfo
		in := &CreateClientInfoV6Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateClientInfoV6(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 125: // R_DhcpV4FailoverGetAddressStatus
		in := &FailoverGetAddressStatusV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.FailoverGetAddressStatusV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 126: // R_DhcpV4CreatePolicyEx
		in := &CreatePolicyExV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreatePolicyExV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 127: // R_DhcpV4GetPolicyEx
		in := &GetPolicyExV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetPolicyExV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 128: // R_DhcpV4SetPolicyEx
		in := &SetPolicyExV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.SetPolicyExV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 129: // R_DhcpV4EnumPoliciesEx
		in := &EnumPoliciesExV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumPoliciesExV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 130: // R_DhcpV4EnumSubnetClientsEx
		in := &EnumSubnetClientsExV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.EnumSubnetClientsExV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 131: // R_DhcpV4CreateClientInfoEx
		in := &CreateClientInfoExV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.CreateClientInfoExV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	case 132: // R_DhcpV4GetClientInfoEx
		in := &GetClientInfoExV4Request{}
		if err := in.UnmarshalNDR(ctx, r); err != nil {
			return nil, err
		}
		resp, err := o.GetClientInfoExV4(ctx, in)
		return resp.xxx_ToOp(ctx), err
	}
	return nil, nil
}
