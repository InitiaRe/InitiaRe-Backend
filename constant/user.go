package constant

const (
	USER_ROLE_NORMAL int = iota + 1
	USER_ROLE_GUEST
	USER_ROLE_CENSOR
	USER_ROLE_ADMIN
	USER_ROLE_SUPER_ADMIN
)

const (
	USER_STATUS_ACTIVE int = iota + 1
	USER_STATUS_INACTIVE
)

var USER_STATUS = map[int]string{
	USER_STATUS_ACTIVE:   "Active",
	USER_STATUS_INACTIVE: "Inactive",
}
