package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleKey is the name of the module
	ModuleName = "issue"
	// StoreKey is the store key string for issue
	StoreKey = ModuleName
	// RouterKey is the message route for issue
	RouterKey = ModuleName
	// QuerierRoute is the querier route for issue
	QuerierRoute = ModuleName
	// Parameter store default namestore
	DefaultParamspace = ModuleName
)
const (
	DefaultCodespace sdk.CodespaceType = ModuleName
)
const (
	IDLength = 15
	IDPreStr = "coin"
	Custom   = "custom"
)
const (
	QueryParams = "params"
	QueryIssues = "list"
	QueryIssue  = "query"
)
const (
	TypeMsgIssue              = "issue"
	TypeMsgIssueMint          = "issueMint"
	TypeMsgIssueBurn          = "issueBurn"
	TypeMsgIssueFinishMinting = "issueFinishMinting"
)
const (
	CoinDecimalsMaxValue = uint(18)
	CoinNameMaxLength    = 15
	CoinSymbolMaxLength  = 6
)