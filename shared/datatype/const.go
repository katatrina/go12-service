package datatype

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusDeleted  Status = "deleted"
	StatusPending  Status = "pending"
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
