package types

type Account interface {
	GetName() *string
	GetUsername() *string
	GetEmail() *string
	GetLink() *string
	GetCreatedAt() *string
	GetLastSeen() *string
	GetFollowerCount() *string
	GetIsAdmin() *bool
}
