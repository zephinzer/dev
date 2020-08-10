package types

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AccountTests struct {
	suite.Suite
}

type accountMock struct{}

func (am accountMock) GetName() *string {
	m := "__name"
	return &m
}
func (am accountMock) GetUsername() *string {
	m := "__username"
	return &m
}
func (am accountMock) GetEmail() *string {
	m := "__email"
	return &m
}
func (am accountMock) GetLink() *string {
	m := "__link"
	return &m
}
func (am accountMock) GetCreatedAt() *string {
	m := "__created_at"
	return &m
}
func (am accountMock) GetLastSeen() *string {
	m := "__last_seen"
	return &m
}
func (am accountMock) GetFollowerCount() *int {
	m := -123456
	return &m
}
func (am accountMock) GetProjectCount() *int {
	m := -12345
	return &m
}
func (am accountMock) GetIsAdmin() *bool {
	m := false
	return &m
}

func TestAccount(t *testing.T) {
	suite.Run(t, &AccountTests{})
}

func (s *AccountTests) Test_PrintAccount() {
	accImpl := accountMock{}
	acc := Account(accImpl)
	output := PrintAccount(acc)
	s.Regexp(regexp.MustCompile("name.+__name"), output)
	s.Regexp(regexp.MustCompile("username.+__username"), output)
	s.Regexp(regexp.MustCompile("email.+__email"), output)
	s.Regexp(regexp.MustCompile("link.+__link"), output)
	s.Regexp(regexp.MustCompile("created at.+__created_at"), output)
	s.Regexp(regexp.MustCompile("last updated.+__last_seen"), output)
	s.Regexp(regexp.MustCompile("project count.+-12345"), output)
	s.Regexp(regexp.MustCompile("follower count.+-123456"), output)
	s.Regexp(regexp.MustCompile("is admin.+false"), output)
}
