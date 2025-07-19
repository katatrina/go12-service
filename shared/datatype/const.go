package datatype

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusDeleted  Status = "deleted"
	StatusPending  Status = "pending"
)

type UserRole string

const (
	RoleUser    UserRole = "user"
	RoleAdmin   UserRole = "admin"
	RoleShipper UserRole = "shipper"
)

const (
	KeyRequester = "requester"
)

func (s Status) Valid() bool {
	switch s {
	case StatusActive, StatusInactive, StatusDeleted, StatusPending:
		return true
	}
	return false
}
