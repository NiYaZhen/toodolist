package create

import "time"

type CreateInput struct {
	Title    string
	Content  string
	TaskType string
	Remark   string

	StartdateAt string
	EnddateAt   string
}

type CreateOutput struct {
	TaskId   string
	Title    string
	Content  string
	TaskType string
	Remark   string

	StartdateAt time.Time
	EnddateAt   time.Time
}
