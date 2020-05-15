package types

import (
	"fmt"
	"strings"
)

type Account interface {
	GetName() *string
	GetUsername() *string
	GetEmail() *string
	GetLink() *string
	GetCreatedAt() *string
	GetLastSeen() *string
	GetFollowerCount() *int
	GetProjectCount() *int
	GetIsAdmin() *bool
}

func PrintAccount(account Account) string {
	var output strings.Builder
	if account.GetName() != nil {
		output.WriteString(fmt.Sprintf("name          : %s\n", *account.GetName()))
	}
	if account.GetUsername() != nil {
		output.WriteString(fmt.Sprintf("username      : %s\n", *account.GetUsername()))
	}
	if account.GetEmail() != nil {
		output.WriteString(fmt.Sprintf("email         : %s\n", *account.GetEmail()))
	}
	if account.GetLink() != nil {
		output.WriteString(fmt.Sprintf("link          : %s\n", *account.GetLink()))
	}
	if account.GetCreatedAt() != nil {
		output.WriteString(fmt.Sprintf("created at    : %s\n", *account.GetCreatedAt()))
	}
	if account.GetLastSeen() != nil {
		output.WriteString(fmt.Sprintf("last updated  : %s\n", *account.GetLastSeen()))
	}
	if account.GetProjectCount() != nil {
		output.WriteString(fmt.Sprintf("project count : %v\n", *account.GetProjectCount()))
	}
	if account.GetFollowerCount() != nil {
		output.WriteString(fmt.Sprintf("follower count: %v\n", *account.GetFollowerCount()))
	}
	return output.String()
}
