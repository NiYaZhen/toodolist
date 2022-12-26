package todolist

import (
	pb "todolist/todolist/app/interface/controller/grpc/protobuf"
	"todolist/todolist/app/usecase/todolist/create"
	"todolist/todolist/app/usecase/todolist/delete"
	"todolist/todolist/app/usecase/todolist/get"
	"todolist/todolist/app/usecase/todolist/search"
	"todolist/todolist/app/usecase/todolist/update"
)

func convertToCreateInput(in *pb.CreateTasksRequest) *create.CreateInput {
	return &create.CreateInput{
		Title:       in.GetTitle(),
		Content:     in.GetContent(),
		Remark:      in.GetRemark(),
		TaskType:    in.GetType(),
		StartdateAt: in.GetStartdateat(),
		EnddateAt:   in.GetEnddateat(),
	}

}

func convertToUpdateInput(in *pb.UpdateRequest) *update.UpdateInput {
	return &update.UpdateInput{
		Title:       in.GetTitle(),
		Content:     in.GetContent(),
		Remark:      in.GetRemark(),
		TaskType:    in.GetType(),
		StartdateAt: in.GetStartdateat(),
		EnddateAt:   in.GetEnddateat(),
	}
}

func convertToDeleteInput(in *pb.DeleteRequest) *delete.DeleteInput {
	return &delete.DeleteInput{
		ID: in.GetId(),
	}
}

func convertToGetInput(in *pb.GetTasksRequest) *get.GetInput {
	return &get.GetInput{
		Useremail: in.Useremail,
	}
}

func converToSearchInput(in *pb.SearchRequest) *search.SearchInput {
	return &search.SearchInput{
		ID: in.GetId(),
	}
}

func converToSearchTypeInput(in *pb.SearchTypeRequest) *search.SearchTypeInput {
	return &search.SearchTypeInput{
		TaskType: in.GetType(),
	}
}
