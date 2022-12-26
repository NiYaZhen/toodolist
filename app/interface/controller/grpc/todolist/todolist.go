package todolist

import (
	"context"
	pb "todolist/todolist/app/interface/controller/grpc/protobuf"
)

func (s *Server) Create(ctx context.Context, data *pb.CreateTasksRequest) (*pb.CreateTasksResponse, error) {

	output, err := s.createUsecase.Create(ctx, convertToCreateInput(data))
	if err != nil {
		return new(pb.CreateTasksResponse), err
	}

	return convertToPbCreateResponse(output), nil
}

func (s *Server) GetTasksList(ctx context.Context, data *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	output, err := s.getUsecase.Get(ctx, convertToGetInput(data))
	if err != nil {
		return new(pb.GetTasksResponse), err
	}
	return convertToPbTaskListResponse(output), nil
}

func (s *Server) Update(ctx context.Context, data *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	output, err := s.updateUsecase.Update(ctx, convertToUpdateInput(data))
	if err != nil {
		return new(pb.UpdateResponse), err
	}
	return convertToPbUpdateResponse(output), nil
}
func (s *Server) Delete(ctx context.Context, data *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := s.deleteUsecase.Delete(ctx, convertToDeleteInput(data))
	if err != nil {
		return new(pb.DeleteResponse), err
	}
	return &pb.DeleteResponse{Err: 0, Msg: "success"}, nil
}
func (s *Server) DeleteAll(ctx context.Context, req *pb.DeleteAllRequest) (*pb.DeleteAllResponse, error) {

	return nil, nil
}
func (s *Server) Search(ctx context.Context, data *pb.SearchRequest) (*pb.SearchResponse, error) {
	output, err := s.searchUsecase.Search(ctx, converToSearchInput(data))
	if err != nil {
		return new(pb.SearchResponse), err
	}
	return convertToPbSearchResponse(output), nil
}
func (s *Server) SearchTypeTasks(ctx context.Context, data *pb.SearchTypeRequest) (*pb.SearchTypeResponse, error) {
	output, err := s.searchUsecase.SearchType(ctx, converToSearchTypeInput(data))
	if err != nil {
		return new(pb.SearchTypeResponse), err
	}
	return convertToPbSearchTypeResponse(output), nil
}
func (s *Server) Complete(ctx context.Context, req *pb.CompleteRequest) (*pb.CompleteResponse, error) {
	return nil, nil
}
func (s *Server) UpdoTasksDate(ctx context.Context, req *pb.UpdoTasksDateRequest) (*pb.UpdoTasksDateResponse, error) {
	return nil, nil
}
