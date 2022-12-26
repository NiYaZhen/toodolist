package server

import (
	"log"
	"net"

	"google.golang.org/grpc"

	service "todolist/todolist/app/domain/service/todolist"
	tt "todolist/todolist/app/infra/persistence/mongo/todolist"
	"todolist/todolist/app/infra/snowflake"
	pb "todolist/todolist/app/interface/controller/grpc/protobuf"
	"todolist/todolist/app/interface/controller/grpc/todolist"
	"todolist/todolist/app/usecase/todolist/create"
	"todolist/todolist/app/usecase/todolist/delete"
	"todolist/todolist/app/usecase/todolist/get"
	"todolist/todolist/app/usecase/todolist/search"
	"todolist/todolist/app/usecase/todolist/update"
)

const (
	port = ":50051"
)

func RunToDoListServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 RPC 服务容器
	NewNode := snowflake.NewNode()
	grpcServer := grpc.NewServer()
	NewService := service.NewToDoListService(NewNode)
	NewRepo := tt.NewRepo()
	NewCreate := create.NewCreateUsecase(NewRepo, NewService)
	NewUpdate := update.NewUpdateUsecase(NewRepo, NewService)
	NewGet := get.NewGetUsecase(NewRepo, NewService)
	NewSearch := search.NewSearchUsecase(NewRepo)
	Newdelete := delete.NewDeleteUsecase(NewRepo)

	server := todolist.NewToDoListServices(NewCreate, NewUpdate, Newdelete, NewGet, NewSearch, NewSearch)

	//var server pb.TasksServer
	// 为 User 服务注册业务实现 将 User 服务绑定到 RPC 服务容器上
	pb.RegisterTasksServer(grpcServer, server)
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系

	//reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
