package main

import (
	"context"
	"log"
	"net"

	pb "example.com/todo/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	PORT = ":50051"
)

type TodoServer struct{
	pb.UnimplementedTodoServiceServer
	todos map[string]*pb.Todo
}

func (s *TodoServer) CreateTodo(ctx context.Context, in *pb.NewTodo) (*pb.Todo,error){
	log.Printf("Received: %v", in.GetName())
	todo := &pb.Todo{
		Name: in.GetName(),
		Description: in.GetDescription(),
		Done: false,
		Id: uuid.New().String(),
	}

	s.todos[todo.Id] = todo

	return todo, nil
}

func (s *TodoServer) DeleteTodo(ctx context.Context, in *pb.TodoId) (*pb.Empty, error){
	log.Printf("Deleting todo with ID: %s", in.GetId())

	if _, exists := s.todos[in.GetId()]; !exists {
		return nil, status.Error(codes.NotFound, "Todo not found")
	}

	delete(s.todos, in.GetId())
	return &pb.Empty{}, nil
}

func (s *TodoServer) ModifyTodo(ctx context.Context, in *pb.Todo) (*pb.Todo, error){
	log.Printf("Modifying todo with ID: %s", in.GetId())

	existingTodo, exists := s.todos[in.GetId()]

	if !exists {
		return nil, status.Error(codes.NotFound, "Todo was not found")
	}

	if in.GetName() != "" {
		existingTodo.Name = in.GetName()
	}

	if in.GetDescription() != "" {
		existingTodo.Description = in.GetDescription()
	}

	existingTodo.Done = in.GetDone()

	return existingTodo, nil
}

func (s *TodoServer) ListTodos(empty *pb.Empty, stream pb.TodoService_ListTodosServer) error {
	log.Printf("Listing all todos")

	for _, todo := range s.todos {
		if err := stream.Send(todo); err!= nil {
			return err
		}
	}
	return nil
}

func NewTodoServer() *TodoServer {
	return &TodoServer{
		todos: make(map[string]*pb.Todo),
	}
}

func main(){
	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()
	todoServer := NewTodoServer()

	pb.RegisterTodoServiceServer(s, todoServer)

	log.Printf("Server is listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}