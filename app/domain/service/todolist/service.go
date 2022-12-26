package service

import (
	"context"
	"strconv"
	"time"
	model "todolist/todolist/app/domain/model/todolist"

	"github.com/bwmarrin/snowflake"
)

type todolistService struct {
	node *snowflake.Node
}

func NewToDoListService(node *snowflake.Node) TodoListService {
	return &todolistService{
		node: node,
	}
}

func (t *todolistService) NewTaskID(ctx context.Context) (string, error) {
	return t.node.Generate().String(), nil
}
func (t *todolistService) NewTask(ctx context.Context, taskId, titie, content, tasktype, remark, startdateAt, enddateAt string) (*model.ToDoList, error) {
	newTask := new(model.ToDoList)
	if taskId == "" {

	}
	newTask.Title = titie
	newTask.Content = content
	newTask.TaskType = tasktype
	newTask.Remark = remark
	newTask.CreatedAt = time.Time{}
	if startdateAt == "" {
		newTask.StartdateAt = time.Time{}
	} else {
		parseTime, err := time.Parse(time.RFC3339, startdateAt)
		if err != nil {
			return nil, err
		}
		newTask.StartdateAt = parseTime
	}

	parseTime, err := time.Parse(time.RFC3339, enddateAt)
	if err != nil {
		return nil, err
	}
	newTask.EnddateAt = parseTime

	return newTask, nil

}

func (t *todolistService) NewUpdateTask(ctx context.Context, titie, content, tasktype, remark, startdateAt, enddateAt string) (*model.ToDoList, error) {
	newTask := new(model.ToDoList)

	newTask.Title = titie
	newTask.Content = content
	newTask.TaskType = tasktype
	newTask.Remark = remark

	if startdateAt == "" {
		newTask.StartdateAt = time.Time{}
	} else {
		parseTime, err := time.Parse(time.RFC3339, startdateAt)
		if err != nil {
			return nil, nil
		}
		newTask.StartdateAt = parseTime
	}

	parseTime, err := time.Parse(time.RFC3339, enddateAt)
	if err != nil {
		return nil, nil
	}
	newTask.EnddateAt = parseTime
	newTask.UpdateAt = time.Now()

	return newTask, nil

}

func (t *todolistService) GetTaskStatus(ctx context.Context, startdateAt, enddateAt string) (int32, error) {
	parseEndTime, _ := time.Parse(time.RFC3339, enddateAt)
	parseStartTime, _ := time.Parse(time.RFC3339, startdateAt)

	EnddateAt := parseEndTime
	StartdateAt := parseStartTime
	EnddAfter := EnddateAt.After(StartdateAt)
	EndAfter := EnddateAt.After(time.Now())

	if EndAfter == true {
		return 3, nil
	} else if EnddAfter == true {
		return 2, nil
	} else {
		return 1, nil
	}

}

func (t *todolistService) GetTimeleft(ctx context.Context, enddateAt string) (string, error) {
	NowTime := time.Now().Unix()
	parseEndTime, err := time.Parse(time.RFC3339, enddateAt)
	if err != nil {

	}
	Timelift := NowTime - parseEndTime.Unix()
	StrTime := strconv.FormatInt((Timelift/86400), 10) + "d" + strconv.FormatInt((Timelift%86400/3600), 10) + "h" + strconv.FormatInt((Timelift%3600/60), 10) + "m" + strconv.FormatInt((Timelift%60), 10) + "s"
	return StrTime, nil
}
