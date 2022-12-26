package model

import "time"

type ToDoList struct {
	TaskId      string    `bson:"taskId"`
	Title       string    `bson:"title"`
	Content     string    `bson:"content"`
	TaskType    string    `bson:"tasktype"`
	Remark      string    `bson:"remark"`
	UserEmail   string    `bson:"useremail"`
	Status      int32     `bson:"status"`
	Complete    bool      `bson:"completle"`
	CreatedAt   time.Time `bson:"createdat"`
	UpdateAt    time.Time `bson:"updateat"`
	CompletedAt time.Time `bson:"completedat"`
	StartdateAt time.Time `bson:"startdateat"`
	EnddateAt   time.Time `bson:"enddateat"`
	Timeleft    string    `bson:"timeleft"`
}
