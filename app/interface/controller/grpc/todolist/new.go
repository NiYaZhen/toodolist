package todolist

import (
	pb "todolist/todolist/app/interface/controller/grpc/protobuf"
	"todolist/todolist/app/usecase/todolist/create"
	"todolist/todolist/app/usecase/todolist/delete"
	"todolist/todolist/app/usecase/todolist/get"
	"todolist/todolist/app/usecase/todolist/search"
	"todolist/todolist/app/usecase/todolist/update"
)

type Server struct {
	createUsecase     create.CreateUsecase
	updateUsecase     update.UpdateUsecase
	deleteUsecase     delete.DeleteUsecase
	getUsecase        get.GetUsecase
	searchUsecase     search.SearchUsecase
	searchTypeUsecase search.SearchUsecase
}

func NewToDoListServices(createUsecase create.CreateUsecase, updateUsecase update.UpdateUsecase, deleteUsecase delete.DeleteUsecase, getUsecase get.GetUsecase, searchUsecase search.SearchUsecase, searchTypeUsecase search.SearchUsecase) pb.TasksServer {
	return &Server{
		createUsecase:     createUsecase,
		updateUsecase:     updateUsecase,
		deleteUsecase:     deleteUsecase,
		getUsecase:        getUsecase,
		searchUsecase:     searchUsecase,
		searchTypeUsecase: searchTypeUsecase,
	}
}
