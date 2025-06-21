package datatype

type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusDeleted  Status = "deleted"
)

func (s Status) Valid() bool {
	switch s {
	case StatusActive, StatusInactive, StatusDeleted:
		return true
	}
	return false
}
