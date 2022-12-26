package model

import "time"

type ToDoList struct {
	TaskId      string
	Title       string
	Content     string
	TaskType    string
	Remark      string
	UserEmail   string
	Status      int32
	Complete    bool
	CreatedAt   time.Time
	UpdateAt    time.Time
	CompletedAt time.Time
	StartdateAt time.Time
	EnddateAt   time.Time
	Timeleft    string
}
