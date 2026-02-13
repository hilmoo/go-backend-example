package todo

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/hilmoo/go-backend-example/internal/model"
	"github.com/ory/herodot"
)

func getTodo(ctx context.Context, payload *GetTodoRequest, db DbRepository) (*GetTodoResponse, *herodot.DefaultError) {
	todo, err := db.GetByID(ctx, payload.ID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, herodot.ErrNotFound.WithReason(err.Error()).WithID(ErrTodoNotFound)
		}
		return nil, herodot.ErrInternalServerError.WithReasonf("Failed to get todo: %v", err)
	}

	return &GetTodoResponse{
		ID:      todo.ID,
		Title:   todo.Title,
		Details: todo.Details,
		Status:  todo.Status,
	}, nil
}

func createTodo(ctx context.Context, payload *CreateTodoRequest, db DbRepository) (*CreateTodoResponse, *herodot.DefaultError) {
	now := time.Now()
	todo := &model.Todo{
		ID:        uuid.New().String(),
		Title:     payload.Title,
		Details:   payload.Details,
		Status:    payload.Status,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := db.Create(ctx, todo); err != nil {
		return nil, herodot.ErrInternalServerError.WithReasonf("Failed to create todo: %v", err).WithID(ErrTodoCreationFailed)
	}

	return &CreateTodoResponse{
		ID:      todo.ID,
		Title:   todo.Title,
		Details: todo.Details,
		Status:  todo.Status,
	}, nil
}

func updateTodo(ctx context.Context, payload *UpdateTodoRequest, db DbRepository) (*UpdateTodoResponse, *herodot.DefaultError) {
	todo := &model.Todo{
		ID:        payload.ID,
		Title:     payload.Title,
		Details:   payload.Details,
		Status:    payload.Status,
		UpdatedAt: time.Now(),
	}

	if err := db.Update(ctx, todo); err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, herodot.ErrNotFound.WithReason(err.Error()).WithID(ErrTodoNotFound)
		}
		return nil, herodot.ErrInternalServerError.WithReasonf("Failed to update todo: %v", err).WithID(ErrTodoUpdateFailed)
	}

	return &UpdateTodoResponse{
		ID:      todo.ID,
		Title:   todo.Title,
		Details: todo.Details,
		Status:  todo.Status,
	}, nil
}

func deleteTodo(ctx context.Context, payload *DeleteTodoRequest, db DbRepository) (*DeleteTodoResponse, *herodot.DefaultError) {
	if err := db.Delete(ctx, payload.ID); err != nil {
		if herodotErr, ok := err.(*herodot.DefaultError); ok {
			return nil, herodotErr
		}
		return nil, herodot.ErrInternalServerError.WithReasonf("Failed to delete todo: %v", err).WithID(ErrTodoDeleteFailed)
	}

	return &DeleteTodoResponse{
		Message: "Todo item deleted successfully.",
	}, nil
}

func listTodos(ctx context.Context, payload *ListTodosRequest, db DbRepository) (*ListTodosResponse, *herodot.DefaultError) {
	todos, err := db.List(ctx, payload.Status)
	if err != nil {
		return nil, herodot.ErrInternalServerError.WithReasonf("Failed to list todos: %v", err)
	}

	todoResponses := make([]GetTodoResponse, 0)
	for _, todo := range todos {
		todoResponses = append(todoResponses, GetTodoResponse{
			ID:      todo.ID,
			Title:   todo.Title,
			Details: todo.Details,
			Status:  todo.Status,
		})
	}

	return &ListTodosResponse{
		Todos: todoResponses,
		Total: len(todoResponses),
	}, nil
}
