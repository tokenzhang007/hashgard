package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/types"
)

// Key for getting a the next available proposalID from the store
var (
	KeyDelimiter   = []byte(":")
	KeyNextIssueID = []byte("newIssueID")
)

//func BytesString(b []byte) string {
//	return *(*string)(unsafe.Pointer(&b))
//}
// Key for getting a specific issuer from the store
func KeyIssuer(issueIdStr string) []byte {
	return []byte(fmt.Sprintf("issues:%s", issueIdStr))
}

// Key for getting a specific address from the store
func KeyAddressIssues(addr string) []byte {
	return []byte(fmt.Sprintf("address:%s", addr))
}

// Key for getting a specific allowed from the store
func KeyAllowed(issueID string, sender sdk.AccAddress, spender sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("allowed:%s:%s:%s", issueID, sender.String(), spender.String()))
}
func KeySymbolIssues(symbol string) []byte {
	return []byte(fmt.Sprintf("symbol:%s", strings.ToUpper(symbol)))
}

func KeyIssueIdStr(seq uint64) string {

	return fmt.Sprintf("%s%x", types.IDPreStr, seq)
}
