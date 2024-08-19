package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "example.com/todo/proto"
	"google.golang.org/grpc"
)

const (
	ADDRESS = "localhost:50051"
)

type TodoTask struct {
	Name string
	Description string
	Done bool
}

func main(){
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure(), grpc.WithBlock())

	var id string

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	todos := []TodoTask{
		{Name: "Code review", Description: "Review new feature code", Done: false},
		{Name: "Make Youtube vidoe", Description: "Create a new youtube video", Done: false},
		{Name: "Gym", Description: "go to the gym", Done: false},
		{Name: "Grocery", Description: "Go to the grocery", Done: false},
		{Name: "Meeting", Description: "Meet about blockers in project", Done: false},
	}

	for _, todo := range todos {
		res, err := c.CreateTodo(ctx, &pb.NewTodo{Name: todo.Name, Description: todo.Description, Done: todo.Done})

		if err != nil {
			log.Fatalf("Could not create task: %v", err)
		}

		log.Printf(`
		ID: %s
		Name: %s
		Description: %s
		Done: %v
		`, res.GetId(), res.GetName(), res.GetDescription(), res.GetDone())

		id = res.GetId()
	}

	modifiedTodo := &pb.Todo {
		Id: id,
		Name: "Updated Name",
		Description: "Updated description",
		Done: true,
	}

	modifiedRes, err := c.ModifyTodo(ctx, modifiedTodo)

	if err != nil {
		log.Printf("Todo with ID: %s", id)
		log.Fatalf("Could not modify todo: %v", err)
	}

	log.Printf("Modified Todo: %+v", modifiedRes)

	deleteRes, err := c.DeleteTodo(ctx, &pb.TodoId{Id: id})

	if err != nil {
		log.Fatalf("Could not delete todo: %v", err)
	}

	log.Printf("Deleted todo: %+v", deleteRes)

	listStream, err := c.ListTodos(ctx, &pb.Empty{})

	if err != nil {
		log.Fatalf("Could not list todos: %v", err)
	}

	for {
		todo, err := listStream.Recv()

		if err == io.EOF{
			log.Printf("end of list")
			return
		}

		if err != nil {
			log.Fatalf("Could not receive todo from stream: %v", err)
		}

		log.Printf("Listed todo: %+v", todo)
	}
}