package data

import (
	"fmt"
	"strconv"
)

type MeetingId int32

type Meeting struct {
	Id     MeetingId
	Title  string
	UserId UserId
}

type MeetingCursor interface {
	Next() (*Meeting, error)
	Close() error
}

type Meetings interface {
	Get(id MeetingId) (*Meeting, error)
	LoadCursor() (MeetingCursor, error)
}

func MeetingIdFromString(id string) (MeetingId, error) {
	if id == "" {
		return 0, nil
	}

	int, err := strconv.Atoi(id)
	return MeetingId(int), err
}

func (id *MeetingId) String() string {
	if id == nil {
		return ""
	}

	return fmt.Sprintf("%d", *id)
}

func (id *MeetingId) Int32() int32 {
	if id == nil {
		return 0
	}

	return int32(*id)
}
