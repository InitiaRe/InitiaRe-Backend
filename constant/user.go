package constant

const (
	USER_STATUS_ACTIVE int = iota + 1
	USER_STATUS_INACTIVE
)

const (
	USER_SYSTEM_ID = 1
)

var USER_STATUS = map[int]string{
	USER_STATUS_ACTIVE:   "Active",
	USER_STATUS_INACTIVE: "Inactive",
}