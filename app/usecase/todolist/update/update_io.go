package update

import "time"

type UpdateInput struct {
	ID          string
	Title       string
	Content     string
	TaskType    string
	Remark      string
	UserEmail   string
	StartdateAt string
	EnddateAt   string
}

type UpdateOutput struct {
	ID          string
	Title       string
	Content     string
	TaskType    string
	Remark      string
	UserEmail   string
	StartdateAt time.Time
	EnddateAt   time.Time
	UpdateAt    time.Time
	Status      int32
	Timeleft    string
}

type CompleteInput struct {
	ID string
}
