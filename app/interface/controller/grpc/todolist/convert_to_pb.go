package todolist

import (
	"time"
	pb "todolist/todolist/app/interface/controller/grpc/protobuf"
	dto "todolist/todolist/app/usecase/todolist"
	"todolist/todolist/app/usecase/todolist/create"
	"todolist/todolist/app/usecase/todolist/get"
	"todolist/todolist/app/usecase/todolist/search"
	"todolist/todolist/app/usecase/todolist/update"
)

func convertToPbCreateResponse(in *create.CreateOutput) *pb.CreateTasksResponse {

	if in == nil {
		return new(pb.CreateTasksResponse)
	}
	return &pb.CreateTasksResponse{
		Err: 0,
		Msg: "success",
		Data: []*pb.ToDoListEntity{
			{
				Id:          in.TaskId,
				Title:       in.Title,
				Content:     in.Content,
				Remark:      in.Remark,
				Startdateat: in.StartdateAt.Format(time.RFC3339),
				Enddateat:   in.EnddateAt.Format(time.RFC3339),
			},
		},
	}

}

func convertToPbUpdateResponse(in *update.UpdateOutput) *pb.UpdateResponse {

	if in == nil {
		return new(pb.UpdateResponse)
	}
	return &pb.UpdateResponse{
		Err: 0,
		Msg: "success",
		Data: []*pb.ToDoListEntity{
			{
				Id:          in.ID,
				Title:       in.Title,
				Content:     in.Content,
				Remark:      in.Remark,
				Startdateat: in.StartdateAt.Format(time.RFC3339),
				Enddateat:   in.EnddateAt.Format(time.RFC3339),
				Updataat:    in.UpdateAt.Format(time.RFC3339),
				Status:      in.Status,
			},
		},
	}

}

func convertToPbTaskListResponse(in *get.GetOutput) *pb.GetTasksResponse {
	if in == nil {
		return new(pb.GetTasksResponse)
	}
	return &pb.GetTasksResponse{
		Data: convertToPbTasks(in.List),
		Err:  0,
		Msg:  "success",
	}
}

func convertToPbTask(in *dto.TaskDto) *pb.ToDoListEntity {
	if in == nil {
		return new(pb.ToDoListEntity)
	}

	return &pb.ToDoListEntity{
		Title:       in.Title,
		Content:     in.Content,
		Remark:      in.Remark,
		Startdateat: in.StartdateAt,
		Enddateat:   in.EnddateAt,
	}
}

func convertToPbTasks(in []*dto.TaskDto) []*pb.ToDoListEntity {
	ans := make([]*pb.ToDoListEntity, len(in))
	for i, v := range in {
		ans[i] = convertToPbTask(v)
	}
	return ans

}

func convertToPbSearchResponse(in *search.SearchOutput) *pb.SearchResponse {
	if in == nil {
		return new(pb.SearchResponse)
	}
	task := convertToPbTask(in.Task)
	return &pb.SearchResponse{
		Err:   0,
		Msg:   "success",
		Tasks: task,
	}
}

func convertToPbSearchTypeResponse(in *search.SearchTypeOutput) *pb.SearchTypeResponse {
	if in == nil {
		return new(pb.SearchTypeResponse)
	}

	return &pb.SearchTypeResponse{
		Err:  0,
		Msg:  "success",
		Data: convertToPbTasks(in.List),
	}

}
